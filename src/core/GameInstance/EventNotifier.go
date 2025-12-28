package gameinstance

type EventNotification struct {
	Type    string `json:"string"`
	Tick    int    `json:"tick"`
	Payload any    `json:"payload"`
}

type IEventNotifier interface {
	Notify(e *EventNotification)
}
