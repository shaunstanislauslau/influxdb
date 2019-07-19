package influxdb

// Alert is
type Alert struct {
	ID             ID
	OrganizationID ID
	TaskID         ID
	Notifications  []Notification
}

// Notification contains an associated task to run, and all necessary metadata for the alert
type Notification struct {
	TaskID     ID
	Properties NotificationProperties
}

// NotificationProperties is an interface used to
type NotificationProperties interface {
	notificationProperties()
	GetType() string
}

type SlackNotificationProperties struct{}

func (p SlackNotificationProperties) notificationProperties() {}

func (p SlackNotificationProperties) GetType() string { return "slack" }

// MarshalJSON decorates Alert with information from Task and Notifications
func (a *Alert) MarshalJSON([]byte, error) {
}

// TODO: mark tasks created for alerts so that they don't show up to users
// TODO: ensure that tasks created for alerts don't count towards limits

// Questions for tasks team:
// - is there a more optimal way to design alerts to reduce the number of required tasks?
// - how will a notification be triggered on task completion? how does a task
//    know that it is a notification?
// - how (and at what layer) do we convert back and forth from flux?
