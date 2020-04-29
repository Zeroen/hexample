package events

type DatastoreCreatedEvent struct {
	DatastoreId string `json:"datastoreId"`
	UserId      string `json:"UserId"`
}

func NewDatastoreCreatedEvent(idDs string, idUser string) DatastoreCreatedEvent {
	return DatastoreCreatedEvent{
		DatastoreId: idDs,
		UserId:      idUser,
	}
}

func (dsce DatastoreCreatedEvent) FullQualifiedEventName() string {
	return "DatastoreCreatedEvent"
}
