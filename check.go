package influxdb

import "encoding/json"

// Check represents the information required to generate a periodic check task.
type Check struct {
	ID                    ID         `json:"id"`
	OrganizationID        ID         `json:"orgID,omitempty"`
	TaskID                ID         `json:"-"` // the generated task
	Tags                  []CheckTag `json:"tags"`
	StatusMessageTemplate string     `json:"statusMessageTemplate"`
	// AuthorizationID ID     `json:"authorizationID"`
	// Name            string `json:"name"`
	// Description     string `json:"description,omitempty"`
	// Status          string `json:"status"`
	// Query           string `json:"flux"`
	// Every           string `json:"every,omitempty"`
	// Cron            string `json:"cron,omitempty"`
	// Offset          string `json:"offset,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	UpdatedAt  string `json:"updatedAt,omitempty"`
	Properties CheckProperties
}

// CheckTag is a tag k/v pair used when a check writes to the system bucket.
type CheckTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Valid returns an error if the checktag is missing fields
func (t *CheckTag) Valid() error {
	if t.Key == "" || t.Value == "" {
		return &Error{
			Code: EInvalid,
			Msg:  "checktag must contain a key and a value",
		}
	}
	return nil
}

func (c *Check) UnmarshalJSON(data []byte) error {
	return nil
}

func (c *Check) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func UnmarshalCheckPropertiesJSON(b []byte) (CheckProperties, error) {
	var v struct {
		B json.RawMessage `json:"properties"`
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return nil, err
	}

	// if len(v.B)

	// var t struct {
	//   Type: string `json:"type"`
	// }
	return nil, nil
}

type CheckProperties interface {
	GetType() string
}

// CheckUpdate updates a check
type CheckUpdate struct {
}
