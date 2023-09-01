package notification

type Message struct {
	Title string
	Text  string
}

type Interface interface {
	Send(Message) error
}

type Notification struct {
	services []Interface
}

func NewNotification() *Notification {
	return &Notification{
		services: []Interface{
			NewTelegram(),
		},
	}
}

func (sn Notification) Send(message Message) error {
	for _, service := range sn.services {
		go func(p Interface) {
			_ = p.Send(message)
		}(service)
	}

	return nil
}
