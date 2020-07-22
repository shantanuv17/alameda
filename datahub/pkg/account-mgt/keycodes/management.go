package keycodes

import (
	"encoding/json"
	"fmt"
	ClusterStatusEntity "github.com/containers-ai/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	RepoInflux "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb"
	RepoClusterStatus "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/clusterstatus"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	ApiEvents "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/events"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"math"
	"strings"
	"time"
)

type KeycodeMgt struct {
	Executor      *KeycodeExecutor
	Status        *KeycodeStatusObject
	KeycodeStatus int
	InfluxCfg     *InternalInflux.Config
	InvalidReason string
}

func NewKeycodeMgt(config *InternalInflux.Config) *KeycodeMgt {
	keycodeMgt := KeycodeMgt{}
	keycodeMgt.Executor = NewKeycodeExecutor()
	keycodeMgt.Status = NewKeycodeStatusObject()
	if keycodeMgt.InfluxCfg == nil {
		keycodeMgt.InfluxCfg = config
	}
	if KeycodeSummary != nil {
		keycodeMgt.KeycodeStatus = keycodeMgt.GetStatus()
	}
	return &keycodeMgt
}

func (c *KeycodeMgt) AddKeycode(keycode string) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.Executor.AddKeycode(keycode)

	if err != nil {
		scope.Errorf("failed to add keycode(%s)", keycode)
		return err
	}

	c.refresh(true)

	return nil
}

func (c *KeycodeMgt) DeleteKeycode(keycode string) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.Executor.DeleteKeycode(keycode)

	if err != nil {
		scope.Errorf("failed to delete keycode(%s)", keycode)
		return err
	}

	c.refresh(true)

	return nil
}

func (c *KeycodeMgt) GetKeycode(keycode string) (*Keycode, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.refresh(false)

	if err != nil {
		scope.Errorf("failed to get keycode(%s)", keycode)
		return nil, err
	}

	stripped := strings.Replace(keycode, "-", "", -1)

	for _, keycodeObj := range KeycodeList {
		if keycodeObj.Keycode == stripped {
			return keycodeObj, nil
		}
	}

	return nil, nil
}

func (c *KeycodeMgt) GetKeycodeSummary() (*Keycode, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.refresh(false)

	if err != nil {
		scope.Error("failed to get keycode summary")
		return nil, err
	}

	return KeycodeSummary, nil
}

func (c *KeycodeMgt) GetKeycodes(keycodes []string) ([]*Keycode, *Keycode, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.refresh(false)

	keycodeList := make([]*Keycode, 0)

	if err != nil {
		scope.Error("failed to get keycodes")
		return nil, nil, err
	}

	for _, keycode := range keycodes {
		stripped := strings.Replace(keycode, "-", "", -1)
		for _, keycodeObj := range KeycodeList {
			if keycodeObj.Keycode == stripped {
				keycodeList = append(keycodeList, keycodeObj)
			}
		}
	}

	return keycodeList, KeycodeSummary, nil
}

func (c *KeycodeMgt) GetAllKeycodes() ([]*Keycode, *Keycode, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.refresh(false)

	if err != nil {
		scope.Error("failed to get all keycodes")
		return make([]*Keycode, 0), nil, err
	}

	return KeycodeList, KeycodeSummary, nil
}

func (c *KeycodeMgt) GetRegistrationData() (string, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	registrationData, err := c.Executor.GetRegistrationData()

	if err != nil {
		scope.Error("failed to get registration data")
		return "", err
	}

	return registrationData, nil
}

func (c *KeycodeMgt) PutSignatureData(signatureData string) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.Executor.PutSignatureData(signatureData)

	if err != nil {
		return err
	}

	c.refresh(true)

	return nil
}

func (c *KeycodeMgt) PutSignatureDataFile(filePath string) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.Executor.PutSignatureDataFile(filePath)

	if err != nil {
		return err
	}

	c.refresh(true)

	return nil
}

func (c *KeycodeMgt) Refresh(force bool) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	return c.refresh(force)
}

func (c *KeycodeMgt) IsValid() bool {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	err := c.refresh(false)

	if err != nil {
		scope.Errorf("failed to check if keycode is valid: %s", err.Error())
		return false
	}

	switch c.GetStatus() {
	case KeycodeStatusNoKeycode:
		return false
	case KeycodeStatusInvalid:
		return false
	case KeycodeStatusExpired:
		return false
	case KeycodeStatusNotActivated:
		return true
	case KeycodeStatusValid:
		return true
	case KeycodeStatusUnknown:
		return false
	default:
		return false
	}
}

func (c *KeycodeMgt) IsExpired() bool {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	summary, err := c.GetKeycodeSummary()

	if err != nil {
		scope.Error("failed to check if keycode is expired")
		return false
	}

	if summary.LicenseState == "Valid" {
		return false
	}

	return true
}

// NOTE: DO Refresh() before GetStatus() if necessary
func (c *KeycodeMgt) GetStatus() int {
	status := c.Status.GetStatus()
	if status != KeycodeStatusValid && c.Status.reason != "" {
		c.InvalidReason = c.Status.GetReason()
	}
	return status
}

// NOTE: DO GET KeycodeMutex lock before using this function
func (c *KeycodeMgt) refresh(force bool) error {
	tm := time.Now()
	tmUnix := tm.Unix()
	refreshed := false
	keycode := "N/A"

	if (force == true) || (int64(math.Abs(float64(tmUnix-KeycodeTimestamp))) >= KeycodeDuration) {
		c.InvalidReason = ""
		keycodeList, keycodeSummary, err := c.Executor.GetAllKeycodes()
		if err != nil {
			scope.Error("failed to refresh keycodes information")
			return err
		}
		// get cluster CPU info from influxdb
		scope.Infof("Licensed CPU cores capacity: %d", keycodeSummary.Capacity.CPUs)
		cores, err := GetAlamedaClusterCPUs(c.InfluxCfg)
		if err != nil {
			scope.Errorf("Failed to refresh keycode CPU info, unable to get CPU info: %s", err.Error())
			ClusterCPUCores = 0
		} else {
			ClusterCPUCores = cores
		}

		// If everything goes right, refresh the global variables
		KeycodeTimestamp = tmUnix
		KeycodeList = keycodeList
		KeycodeSummary = keycodeSummary
		KeycodeTM = tm
		refreshed = true
	}

	if len(KeycodeList) > 0 {
		// log the first keycode in KeycodeList
		keycode = KeycodeList[0].Keycode
	}
	if force == false {
		if refreshed == true {
			scope.Infof("keycode cache data refreshed, keycode: %s", keycode)
		} else {
			scope.Debugf("cached keycode (@ %s): %s", KeycodeTM.Format(time.RFC3339), keycode)
		}
	} else {
		scope.Infof("keycode cache data refreshed for CUD OP, keycode: %s", keycode)
	}

	if c.KeycodeStatus != c.GetStatus() {
		KeycodeStatus = c.GetStatus()
		c.KeycodeStatus = c.GetStatus()

		// Update InfluxDB and post event
		switch KeycodeStatus {
		case KeycodeStatusNoKeycode:
			c.writeInfluxEntry("N/A", KeycodeStatusName[KeycodeStatusNoKeycode])
			c.deleteInfluxEntry("Summary")
			PostEvent(ApiEvents.EventLevel_EVENT_LEVEL_ERROR, KeycodeStatusMessage[KeycodeStatusNoKeycode])
		case KeycodeStatusInvalid:
			msg := KeycodeStatusMessage[KeycodeStatusInvalid]
			if c.InvalidReason != "" {
				msg = c.InvalidReason
			}
			c.writeInfluxEntry("Summary", KeycodeStatusName[KeycodeStatusInvalid])
			c.deleteInfluxEntry("N/A")
			PostEvent(ApiEvents.EventLevel_EVENT_LEVEL_ERROR, msg)
		case KeycodeStatusExpired:
			c.writeInfluxEntry("Summary", KeycodeStatusName[KeycodeStatusExpired])
			c.deleteInfluxEntry("N/A")
			PostEvent(ApiEvents.EventLevel_EVENT_LEVEL_ERROR, KeycodeStatusMessage[KeycodeStatusExpired])
		case KeycodeStatusNotActivated:
			c.writeInfluxEntry("Summary", KeycodeStatusName[KeycodeStatusNotActivated])
			c.deleteInfluxEntry("N/A")
			PostEvent(ApiEvents.EventLevel_EVENT_LEVEL_WARNING, KeycodeStatusMessage[KeycodeStatusNotActivated])
		case KeycodeStatusValid:
			c.writeInfluxEntry("Summary", KeycodeStatusName[KeycodeStatusValid])
			c.deleteInfluxEntry("N/A")
			PostEvent(ApiEvents.EventLevel_EVENT_LEVEL_INFO, KeycodeStatusMessage[KeycodeStatusValid])
		default:
			c.writeInfluxEntry("Summary", KeycodeStatusName[KeycodeStatusUnknown])
			c.deleteInfluxEntry("N/A")
			PostEvent(ApiEvents.EventLevel_EVENT_LEVEL_ERROR, KeycodeStatusMessage[KeycodeStatusUnknown])
		}
	}

	return nil
}

func (c *KeycodeMgt) writeInfluxEntry(keycode, status string) error {
	points := make([]*InfluxClient.Point, 0)
	client := InternalInflux.NewClient(InfluxConfig)

	tags := map[string]string{
		string(ClusterStatusEntity.Keycode): keycode,
	}

	jsonStr, _ := json.Marshal(KeycodeSummary)
	fields := map[string]interface{}{
		string(ClusterStatusEntity.KeycodeStatus):          status,
		string(ClusterStatusEntity.KeycodeType):            KeycodeSummary.KeycodeType,
		string(ClusterStatusEntity.KeycodeState):           KeycodeSummary.LicenseState,
		string(ClusterStatusEntity.KeycodeRegistered):      KeycodeSummary.Registered,
		string(ClusterStatusEntity.KeycodeExpireTimestamp): KeycodeSummary.ExpireTimestamp,
		string(ClusterStatusEntity.KeycodeRawdata):         string(jsonStr[:]),
	}

	pt, err := InfluxClient.NewPoint(string(RepoClusterStatus.Keycode), tags, fields, time.Unix(0, 0))
	if err != nil {
		scope.Error(err.Error())
	}
	points = append(points, pt)

	err = client.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.ClusterStatus),
	})

	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (c *KeycodeMgt) deleteInfluxEntry(keycode string) error {
	if keycode != "" {
		client := InternalInflux.NewClient(InfluxConfig)

		cmd := fmt.Sprintf("DROP SERIES FROM %s WHERE \"%s\"='%s'", RepoClusterStatus.Keycode, ClusterStatusEntity.Keycode, keycode)
		scope.Debugf("delete keycode in influxdb command: %s", cmd)
		_, err := client.QueryDB(cmd, string(RepoInflux.ClusterStatus))
		if err != nil {
			scope.Errorf(err.Error())
			return nil
		}
	}
	return nil
}
