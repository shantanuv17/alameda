package keycodes

import (
	"encoding/json"
	"fmt"
	"github.com/containers-ai/alameda/datahub/pkg/account-mgt/licenses"
	ClusterStatusEntity "github.com/containers-ai/alameda/datahub/pkg/dao/entities/influxdb/clusterstatus"
	RepoInflux "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb"
	RepoClusterStatus "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb/clusterstatus"
	InfluxDB "github.com/containers-ai/alameda/pkg/database/influxdb"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"math"
	"time"
)

type KeycodeMgt struct {
	Executor  *KeycodeExecutor
	Status    *KeycodeStatusObject
	InfluxCfg *InfluxDB.Config
}

func NewKeycodeMgt(config *InfluxDB.Config) *KeycodeMgt {
	keycodeMgt := KeycodeMgt{}
	keycodeMgt.Executor = NewKeycodeExecutor()
	keycodeMgt.Status = NewKeycodeStatusObject()
	if keycodeMgt.InfluxCfg == nil {
		keycodeMgt.InfluxCfg = config
	}
	return &keycodeMgt
}

func (c *KeycodeMgt) AddKeycode(keycode string) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.Executor.AddKeycode(keycode); err != nil {
		scope.Errorf("failed to add keycode(%s)", keycode)
		return err
	}

	return c.refresh(true)
}

func (c *KeycodeMgt) DeleteKeycode(keycode string) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.Executor.DeleteKeycode(keycode); err != nil {
		scope.Errorf("failed to delete keycode(%s)", keycode)
		return err
	}

	return c.refresh(true)
}

func (c *KeycodeMgt) GetKeycode(keycode string) (*Keycode, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
		scope.Errorf("failed to get keycode(%s)", keycode)
		return nil, err
	}

	for _, keycodeObj := range KeycodeList {
		if keycodeObj.Keycode == keycode {
			return keycodeObj, nil
		}
	}

	return nil, nil
}

func (c *KeycodeMgt) GetKeycodeSummary() (*Keycode, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
		scope.Error("failed to get keycode summary")
		return nil, err
	}

	return KeycodeSummary, nil
}

func (c *KeycodeMgt) GetKeycodes(keycodes []string) ([]*Keycode, *Keycode, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
		scope.Error("failed to get keycodes")
		return nil, nil, err
	}

	keycodeList := make([]*Keycode, 0)
	for _, keycode := range keycodes {
		for _, keycodeObj := range KeycodeList {
			if keycodeObj.Keycode == keycode {
				keycodeList = append(keycodeList, keycodeObj)
			}
		}
	}

	return keycodeList, KeycodeSummary, nil
}

func (c *KeycodeMgt) GetAllKeycodes() ([]*Keycode, *Keycode, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
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

func (c *KeycodeMgt) GetCPUCoresOccupied() (int64, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
		scope.Error("failed to get cpu cores occupied")
		return 0, err
	}

	if KeycodeCapacityOccupied == nil {
		return 0, nil
	}

	return KeycodeCapacityOccupied.CPUCores, nil
}

func (c *KeycodeMgt) GetMemoryBytesOccupied() (int64, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
		scope.Error("failed to get memory bytes occupied")
		return 0, err
	}

	if KeycodeCapacityOccupied == nil {
		return 0, nil
	}

	return KeycodeCapacityOccupied.MemoryBytes, nil
}

func (c *KeycodeMgt) GetKeycodeGracePeriod() (int64, error) {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
		scope.Error("failed to get keycode grace period")
		return 0, err
	}

	return KeycodeGracePeriod, nil
}

func (c *KeycodeMgt) PutSignatureData(signatureData string) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.Executor.PutSignatureData(signatureData); err != nil {
		scope.Error("failed to put signature data")
		return err
	}

	return c.refresh(true)
}

func (c *KeycodeMgt) PutSignatureDataFile(filePath string) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.Executor.PutSignatureDataFile(filePath); err != nil {
		scope.Error("failed to put signature data file")
		return err
	}

	return c.refresh(true)
}

func (c *KeycodeMgt) Refresh(force bool) error {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	return c.refresh(force)
}

func (c *KeycodeMgt) IsValid() bool {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
		scope.Errorf("failed to check if keycode is valid: %s", err.Error())
		return false
	}

	switch KeycodeStatus {
	case KeycodeStatusUnknown:
		return false
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
	case KeycodeStatusCapacityCPUCoresGracePeriod:
		return true
	case KeycodeStatusCapacityCPUCoresExceeded:
		return false
	default:
		return false
	}
}

// NOTE: DO Refresh() before GetStatus() if necessary
func (c *KeycodeMgt) GetStatus() int {
	KeycodeMutex.Lock()
	defer KeycodeMutex.Unlock()

	if err := c.refresh(false); err != nil {
		scope.Errorf("failed to get keycode status: %s", err.Error())
		return KeycodeStatusUnknown
	}

	return KeycodeStatus
}

func (c *KeycodeMgt) PostEvent() error {
	if err := PostEvent(KeycodeEventLevelMap[KeycodeStatus], KeycodeStatusMessage[KeycodeStatus]); err != nil {
		scope.Errorf("failed to post keycode event: %s", err.Error())
		return err
	}
	return nil
}

// NOTE: DO GET KeycodeMutex lock before using this function
func (c *KeycodeMgt) refresh(force bool) error {
	tm := time.Now()
	tmUnix := tm.Unix()

	if (force == true) || (int64(math.Abs(float64(tmUnix-KeycodeTimestamp))) >= Duration) {
		keycodeList, keycodeSummary, err := c.Executor.GetAllKeycodes()
		if err != nil {
			scope.Error("failed to refresh keycodes information")
			return err
		}

		// Get federatorai capacity occupied from influxdb
		capacityOccupied, err := GetFederatoraiCapacityOccupied(c.InfluxCfg)
		if err != nil {
			scope.Errorf("failed to get federatorai capacity occupied: %s", err.Error())
			return err
		}
		scope.Infof("licensed CPU cores capacity: %d, CPU cores occupied: %d", keycodeSummary.Capacity.CPUs, capacityOccupied.CPUCores)

		gracePeriod, err := GetKeycodeGracePeriod(c.InfluxCfg)
		if err != nil {
			scope.Errorf("failed to get grace period: %s", err.Error())
			return err
		}

		// If everything goes right, refresh the global variables
		KeycodeList = keycodeList
		KeycodeSummary = keycodeSummary
		KeycodeCapacityOccupied = capacityOccupied
		KeycodeGracePeriod = gracePeriod
		KeycodeTimestamp = tmUnix
		KeycodeTM = tm
		KeycodeStatus = c.Status.GetStatus()
		KeycodeSummary.LicenseState = c.licenseState()
		KeycodeSummary.Message = KeycodeStatusMessage[KeycodeStatus]

		// Update influxdb data
		if err := c.updateInflux(); err != nil {
			scope.Errorf("failed to update influx entries: %s", err.Error())
			return err
		}

		scope.Info("keycode cache data refreshed for CUD OP")
	}

	return nil
}

func (c *KeycodeMgt) licenseState() string {
	if KeycodeStatus == KeycodeStatusCapacityCPUCoresExceeded {
		return "Invalid"
	}
	return KeycodeSummary.LicenseState
}

func (c *KeycodeMgt) updateInflux() error {
	if err := c.updateKeycodeGracePeriod(); err != nil {
		scope.Errorf("failed to update keycode grace period: %s", err.Error())
		return err
	}
	if err := c.deleteKeycodeEntries(); err != nil {
		scope.Errorf("failed to delete keycode entries: %s", err.Error())
		return err
	}
	if err := c.updateKeycodeEntries(KeycodeInfluxTargetMap[KeycodeStatus], KeycodeStatusName[KeycodeStatus]); err != nil {
		scope.Errorf("failed to update keycode entries: %s", err.Error())
		return err
	}
	return nil
}

func (c *KeycodeMgt) deleteKeycodeEntries() error {
	client := InfluxDB.NewClient(InfluxConfig)

	cmd := fmt.Sprintf("DROP SERIES FROM \"%s\"", RepoClusterStatus.Keycode)

	_, err := client.QueryDB(cmd, string(RepoInflux.ClusterStatus))
	if err != nil {
		scope.Errorf(err.Error())
		return err
	}

	return nil
}

func (c *KeycodeMgt) updateKeycodeEntries(keycode, status string) error {
	points := make([]*InfluxClient.Point, 0)
	client := InfluxDB.NewClient(InfluxConfig)

	// Generate keycode summary influxdb point
	jsonStr, _ := json.Marshal(KeycodeSummary)
	tags := map[string]string{
		string(ClusterStatusEntity.Keycode): keycode,
	}
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

	// Generate keycode influxdb points
	if KeycodeList != nil {
		for _, keyInfo := range KeycodeList {
			jsonStr, _ := json.Marshal(keyInfo)
			tags := map[string]string{
				string(ClusterStatusEntity.Keycode): keyInfo.Keycode,
			}
			fields := map[string]interface{}{
				string(ClusterStatusEntity.KeycodeStatus):          keyInfo.LicenseState,
				string(ClusterStatusEntity.KeycodeType):            keyInfo.KeycodeType,
				string(ClusterStatusEntity.KeycodeState):           keyInfo.LicenseState,
				string(ClusterStatusEntity.KeycodeRegistered):      keyInfo.Registered,
				string(ClusterStatusEntity.KeycodeExpireTimestamp): keyInfo.ExpireTimestamp,
				string(ClusterStatusEntity.KeycodeRawdata):         string(jsonStr[:]),
			}
			pt, err := InfluxClient.NewPoint(string(RepoClusterStatus.Keycode), tags, fields, time.Unix(0, 0))
			if err != nil {
				scope.Error(err.Error())
			}
			points = append(points, pt)
		}
	}

	err = client.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.ClusterStatus),
	})

	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (c *KeycodeMgt) updateKeycodeGracePeriod() error {
	if KeycodeStatus == KeycodeStatusValid {
		KeycodeGracePeriod = 0
		if err := DeleteKeycodeGracePeriod(c.InfluxCfg); err != nil {
			return err
		}
		return nil
	}
	if KeycodeStatus == KeycodeStatusCapacityCPUCoresGracePeriod || KeycodeStatus == KeycodeStatusCapacityCPUCoresExceeded {
		if KeycodeGracePeriod == 0 {
			ts := time.Unix(time.Now().Unix()+licenses.CPUCapacityGracePeriod, 0)
			KeycodeGracePeriod = ts.Unix()
			if err := WriteKeycodeGracePeriod(ts.Unix(), c.InfluxCfg); err != nil {
				return err
			}
		}
		// If grace period is greater than ZERO and less than expire timestamp of summary, we have to writ it to summary
		if KeycodeGracePeriod > 0 && KeycodeGracePeriod < KeycodeSummary.ExpireTimestamp {
			KeycodeSummary.ExpireTimestamp = KeycodeGracePeriod
		}
	}
	return nil
}
