package interfaces

type EventPublisher interface {
	PublishEvent(event interface{}) error
}
