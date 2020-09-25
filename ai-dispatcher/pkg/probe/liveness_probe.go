package probe

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	utils "prophetstor.com/alameda/ai-dispatcher/pkg/utils"
)

type LivenessProbeConfig struct {
	DatahubAddr string
	QueueURL    string
}

func checkRabbitmqNotBlock(url string) error {
	conn, err := amqp.Dial(url)
	if conn != nil {
		defer conn.Close()
	}

	if err != nil {
		fmt.Println(err)
		return err
	}
	ch, err := conn.Channel()
	q, err := ch.QueueDeclare(
		"test_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		amqp.Table{
			"x-message-deduplication": true,
		}, // arguments
	)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte("{'test': '123'}"),
			DeliveryMode: 2, // 2 means persistent
			Headers: amqp.Table{
				//"x-deduplication-header": sender.getMessageHash(msgID),
				"x-deduplication-header": "1000",
			},
		})
	if err != nil {
		return err
	}
	return nil
}

func watchDogProbe() error {
	granularities := viper.GetStringSlice("serviceSetting.granularities")
	delaySec := viper.GetInt64("watchdog.delayedSec")

	for _, granularity := range granularities {
		granularitySec := int64(viper.GetInt(
			fmt.Sprintf("granularities.%s.dataGranularitySec", granularity)))
		if granularitySec == 0 {
			return fmt.Errorf("granularity %v is not defined or set incorrect", granularitySec)
		}

		for _, watchFile := range []string{
			fmt.Sprintf("%s/%v", viper.GetString("watchdog.model.directory"), granularitySec),
			fmt.Sprintf("%s/%v", viper.GetString("watchdog.predict.directory"), granularitySec),
		} {
			updatedTimeDiff, err := utils.FileNotUpdateTimeSec(watchFile)
			if err != nil {
				if !os.IsNotExist(err) {
					return err
				} else {
					continue
				}
			}

			if updatedTimeDiff > granularitySec+delaySec {
				return fmt.Errorf("granularity %s of watchdog file %s is not updated for %v seconds",
					granularity, watchFile, updatedTimeDiff)
			}
		}
	}
	return nil
}
