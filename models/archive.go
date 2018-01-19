package models

import "container/list"

const (
	EVENT_JOIN    = iota
	EVEN_LEAVE
	EVENT_MESSAGE
)

type EventType int //事件类型
type Event struct {
	Type      EventType //join leave message
	User      string
	Timestamp int
	Content   string
}

const archiveSize = 20   //档案文件大小
var archive = list.New() //档案文件
//新的文件要保存新的事件，通过文件集合list
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}
func GetEvents(lastReceived int) []Event  {
	events:=make([]Event,0,archive.Len())
	for event:=archive.Front();event!=nil;event=event.Next(){
		e:=event.Value.(Event)
		if e.Timestamp>int(lastReceived){
			events=append(events,e)
		}
	}
	return events
}
