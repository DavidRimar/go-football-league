package interfaces

type EventPublisher interface {
	PublishEvent(event string) error
}
