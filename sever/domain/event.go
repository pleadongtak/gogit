package domain

import "time"

type Event struct {
	UserId         string
	CTN            string
	Email          string
	RegDateTime    time.Time
	AccessDatetime time.Time
	Msg            string
	AccPoint       int
	Point          int
	ServiceType    int
}
type EventSaver interface {
	Save(Event) error
}
