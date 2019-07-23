package notification

import (
	"encoding/json"
	"strings"

	"github.com/influxdata/influxdb"
)

// StatusRule includes parametes of status rules.
type StatusRule struct {
	CurrentLevel  CheckLevel  `json:"currentLevel"`
	PreviousLevel *CheckLevel `json:"previousLevel"`
	// Alert when >= Count per Period
	Count  int               `json:"count"`
	Period influxdb.Duration `json:"period"`
}

// CheckLevel is the enum value of status levels.
type CheckLevel int

// consts of CheckStatusLevel
const (
	Unknown CheckLevel = iota
	Ok
	Info
	Critical
	Warn
)

var checkLevels = []string{
	"UNKNOWN",
	"OK",
	"INFO",
	"CRIT",
	"WARN",
}

var checkLevelMaps = map[string]CheckLevel{
	"UNKNOWN": Unknown,
	"OK":      Ok,
	"INFO":    Info,
	"CRIT":    Critical,
	"WARN":    Warn,
}

// MarshalJSON implements json.Marshaller.
func (cl CheckLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(cl.String())
}

// Pointer returns the reference of the CheckLevel.
func (cl CheckLevel) Pointer() *CheckLevel {
	s := new(CheckLevel)
	*s = cl
	return s
}

// UnmarshalJSON implements json.Unmarshaller.
func (cl *CheckLevel) UnmarshalJSON(b []byte) error {
	var ss string
	if err := json.Unmarshal(b, &ss); err != nil {
		return err
	}
	*cl = ParseCheckLevel(strings.ToUpper(ss))
	return nil
}

// String returns the string value, invalid CheckLevel will return Unknown.
func (cl CheckLevel) String() string {
	if cl < Unknown || cl > Warn {
		cl = Unknown
	}
	return checkLevels[cl]
}

// ParseCheckLevel will parse the string to checkLevel
func ParseCheckLevel(s string) CheckLevel {
	if cl, ok := checkLevelMaps[s]; ok {
		return cl
	}
	return Unknown
}
