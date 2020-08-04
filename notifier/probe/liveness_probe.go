package probe

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"time"
)

const (
	lastEventDataFile = "/tmp/.event.req.latest"
	currentEventDataFile = "/tmp/.event.req.current"
	queueName = "event"
)

type LivenessProbeConfig struct{}

/* CheckRBMQEventQueueAccess check the rabbitmq event queue is consuming event normally.

This function is try to get event from event queue and compare with the message which got
last time.
Conditions:
1. queue length == 0: no message in the queue, assume the status is ok
2. queue length > 0 and no last event message exists: queue status is normal
3. queue length > 0 and last event message exist: compare message, return error if they are same

Current event which get from queue will saved for compare in the next round
 */
func CheckRBMQEventQueueAccess(url string) error {
	var (
		conn *amqp.Connection
		ch *amqp.Channel
		q amqp.Queue
		msgs <- chan amqp.Delivery
		eventBytes []byte
		eventBytesLast []byte
		err error
		f *os.File
		n int
	)
	os.Remove(lastEventDataFile)
	conn, err = amqp.Dial(url)
	if err != nil {
		scope.Errorf("Failed to dial RabbitMQ %s: %s", url, err.Error())
		return err
	}
	ch, err = conn.Channel()
	if err != nil {
		scope.Errorf("Failed to get AMQP channel: %s", err.Error())
		return err
	}
	defer ch.Close()
	q, err = ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		scope.Errorf("Failed to decleare queue \"%s\": %s", queueName, err.Error())
		return err
	}
	totalMsg := q.Messages
	if totalMsg == 0 {
		scope.Debugf("No event queued")
		os.Remove(currentEventDataFile)
		return nil
	}
	msgs, err = ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		scope.Errorf("Failed to config consumer for queue \"%s\": %s", queueName, err.Error())
		return err
	}
	go func(){
		for d := range msgs {
			eventBytes = d.Body
			d.Reject(true)
			break
		}
	}()
	for i := 0; i < 6; i++ {
		if len(eventBytes) > 0 {
			scope.Debug("event received, break from sleep loop")
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	f, err = os.Open(currentEventDataFile)
	if err != nil {
		if os.IsNotExist(err) == false {
			scope.Errorf("Failed to open file %s: %s", currentEventDataFile, err.Error())
			return err
		}
	} else {
		st, err := f.Stat()
		if err != nil {
			scope.Errorf("Failed to stat file %s: %s", currentEventDataFile, err.Error())
			return err
		}
		eventBytesLast = make([]byte, st.Size())
		n, err = f.Read(eventBytesLast)
		if err != nil {
			scope.Errorf("Failed to read last event data from file %s: %s", currentEventDataFile, err.Error())
			return err
		}
		scope.Debugf("read last event data, len: %d", n)
		f.Close()
		if totalMsg > 0 {
			err = os.Rename(currentEventDataFile, lastEventDataFile)
			if err != nil {
				scope.Errorf("Failed to rename file %s to %s: %s", currentEventDataFile, lastEventDataFile, err.Error())
				return err
			}
		}
	}

	f, err = os.Create(currentEventDataFile)
	if err != nil {
		scope.Errorf("Failed to open file %s for write event data: %s", currentEventDataFile, err.Error())
		return err
	}
	_, err = f.Write(eventBytes)
	f.Close()
	if err != nil {
		scope.Errorf("Failed to write event data to file %s: %s", currentEventDataFile, err.Error())
		return err
	}

	if len(eventBytesLast) > 0 && len(eventBytes) > 0 {
		if bytes.Equal(eventBytes, eventBytesLast) {
			err = fmt.Errorf("last event does not consumed")
			scope.Errorf("%s", err.Error())
			return err
		} else {
			scope.Debugf("Last event is consumed")
		}
	} else {
		scope.Debugf("Event queue and consumer is running properly")
	}
	return nil
}