package keycodes

import (
	"encoding/json"
	"github.com/containers-ai/alameda/pkg/database/influxdb"
	"github.com/containers-ai/alameda/pkg/database/ldap"
	"github.com/containers-ai/alameda/pkg/utils/kubernetes"
	"github.com/containers-ai/alameda/pkg/utils/log"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sync"
	"time"
)

// Configuration
var (
	CliPath            = defaultCliPath
	AesKey             = []byte("")
	Duration     int64 = defaultRefreshInterval
	InfluxConfig *influxdb.Config
	LdapConfig   *ldap.Config
	K8SClient    client.Client
)

// Global variable
var (
	scope                                     = log.RegisterScope("account-mgt", "keycode", 0)
	KeycodeStatus                             = KeycodeStatusUnknown
	KeycodeTimestamp        int64             = 0
	KeycodeGracePeriod      int64             = 0
	KeycodeList             []*Keycode        = nil
	KeycodeSummary          *Keycode          = nil
	KeycodeCapacityOccupied *CapacityOccupied = nil
	KeycodeTM               time.Time
	KeycodeMutex            sync.Mutex
)

type Keycode struct {
	Keycode          string           `json:"keycode"          example:"A5IMH-KBAFI-XTEDK-G4OQM-QMM67-4TEST"`
	KeycodeType      string           `json:"keycodeType"      example:"Regular/Trial"`
	KeycodeVersion   int              `json:"keycodeVersion"   example:"2"`
	ApplyTimestamp   int64            `json:"applyTimestamp"   example:"1546271999"`
	ExpireTimestamp  int64            `json:"expireTimestamp"  example:"1546271999"`
	LicenseState     string           `json:"licenseState"     example:"Valid/Invalid/Expired"`
	Registered       bool             `json:"registered"       example:"false"`
	Capacity         Capacity         `json:"capacity"         example:"capacity"`
	Functionality    Functionality    `json:"functionality"    example:"functionality"`
	Retention        Retention        `json:"retention"        example:"retention"`
	ServiceAgreement ServiceAgreement `json:"serviceAgreement" example:"service agreement"`
	Description      string           `json:"description"      example:"your-description"`
	Message          string           `json:"message"          example:"your-message"`
}

type Capacity struct {
	Users int   `json:"users"  example:"-1"`
	Hosts int   `json:"hosts"  example:"20"`
	Disks int   `json:"disks"  example:"200"`
	CPUs  int64 `json:"cpus"   examples:"2"`
}

type Functionality struct {
	Diskprophet bool `json:"diskprophet" example:"true"`
	Workload    bool `json:"workload"    example:"true"`
}

type Retention struct {
	ValidMonth int `json:"validMonth" example:"0"`
	Years      int `json:"years"      example:"1"`
}

type ServiceAgreement struct {
}

func NewKeycode(keycode string) *Keycode {
	key := Keycode{}
	if keycode != "" {
		err := json.Unmarshal([]byte(keycode), &key)
		if err != nil {
			scope.Errorf("failed to unmarshal keycode: %v", err)
			return nil
		}
	}
	return &key
}

func KeycodeInit(config *Config) error {
	CliPath = config.CliPath
	Duration = config.RefreshInterval
	AesKey = config.AesKey
	InfluxConfig = config.InfluxDB
	LdapConfig = config.Ldap

	k8sClient, err := kubernetes.NewK8SClient()
	if err != nil {
		return err
	}
	K8SClient = k8sClient

	return nil
}
