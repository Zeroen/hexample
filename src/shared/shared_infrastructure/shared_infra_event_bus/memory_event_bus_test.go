package shared_infra_event_bus

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"hexample.com/src/shared/shared_domain/shared_domain_event_bus"
	"testing"
	"time"
)

// ------------------------------------------------------
//                  Test Definition
// ------------------------------------------------------

func TestMemoryEventBus(t *testing.T) {

	tests := []struct {
		info   string
		events []shared_domain_event_bus.DomainEvent
	}{
		{
			info: "Check that subscriber receives three emitted events",
			events: []shared_domain_event_bus.DomainEvent{
				NewGenericMemoryEventTest("Event1", 1, 2),
				NewGenericMemoryEventTest("Event2", 3, 4),
				NewGenericMemoryEventTest("Event3", 5, 6),
			},
		},
	}
	for _, test := range tests {

		subscriber := NewGenericTestMemoryEventSubscriber(t)
		subscribers := []DomainEventSubscriber{
			subscriber,
		}

		reb, err := NewMemoryEventBus(
			subscribers)
		require.NoError(t, err, "Error creating the MemoryEventBUS")

		err = reb.Publish(test.events)
		require.NoError(t, err, "Error publishing events")

		time.Sleep(1 * time.Second)

		data := subscriber.RetrieveReceivedData()
		require.Equal(t, len(data), 3)
	}
}

// ------------------------------------------------------
//                  GenericMemoryEventTest Definition
// ------------------------------------------------------
//
// Note:
//       This event will be used in the rest of the test.
//

type GenericMemoryEventTest struct {
	ID     string
	Field1 int `json:"field1"`
	Field2 int `json:"field2"`
}

func NewGenericMemoryEventTest(ID string, field1 int, field2 int) *GenericMemoryEventTest {
	return &GenericMemoryEventTest{
		ID:     ID,
		Field1: field1,
		Field2: field2,
	}
}

func (e *GenericMemoryEventTest) FullQualifiedEventName() string {
	return e.ID
}

// ------------------------------------------------------
//              EventSubscriber Definition
// ------------------------------------------------------
//
// Note:
//       This subscriber will be used in the rest of the test.
//

type GenericTestMemoryEventSubscriber struct {
	events []*GenericMemoryEventTest
	t      *testing.T
}

func NewGenericTestMemoryEventSubscriber(t *testing.T) *GenericTestMemoryEventSubscriber {
	return &GenericTestMemoryEventSubscriber{
		t:      t,
		events: make([]*GenericMemoryEventTest, 0),
	}
}

func (g *GenericTestMemoryEventSubscriber) SubscribedTo() shared_domain_event_bus.DomainEvent {
	return &GenericMemoryEventTest{}
}

func (g *GenericTestMemoryEventSubscriber) ChannelName() string {
	return "EventTestForGenericTestSubscriber"
}

func (g *GenericTestMemoryEventSubscriber) Consume(data []byte) {
	var event GenericMemoryEventTest
	err := json.Unmarshal(data, &event)
	require.NoError(g.t, err, "There should not be an error unmarshaling events.")
	g.events = append(g.events, &event)
}

func (g *GenericTestMemoryEventSubscriber) RetrieveReceivedData() []*GenericMemoryEventTest {
	return g.events
}
