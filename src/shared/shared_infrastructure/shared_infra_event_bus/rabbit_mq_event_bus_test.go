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

func TestRabbitMQEventBus(t *testing.T) {

	tests := []struct {
		info   string
		events []shared_domain_event_bus.DomainEvent
	}{
		{
			info: "Check that subscriber receives three emitted events",
			events: []shared_domain_event_bus.DomainEvent{
				NewGenericEventTest("Event1", 1, 2),
				NewGenericEventTest("Event2", 3, 4),
				NewGenericEventTest("Event3", 5, 6),
			},
		},
	}
	for _, test := range tests {

		subscriber := NewGenericTestEventSubscriber(t)
		subscribers := []DomainEventSubscriber{
			subscriber,
		}

		reb, err := NewRabbitEventBus(
			"rabbitmq",
			"rabbitmq",
			"localhost",
			5672,
			subscribers)
		require.NoError(t, err, "Error connecting to rabbitMQ")

		err = reb.Publish(test.events)
		require.NoError(t, err, "Error publishing events")

		time.Sleep(1 * time.Second)

		data := subscriber.RetrieveReceivedData()
		require.Equal(t, len(data), 3)
		require.Equal(t, test.events[0], data[0])
		require.Equal(t, test.events[1], data[1])
		require.Equal(t, test.events[2], data[2])
	}
}

// ------------------------------------------------------
//                  GenericEventTest Definition
// ------------------------------------------------------
//
// Note:
//       This event will be used in the rest of the test.
//

type GenericEventTest struct {
	ID     string
	Field1 int `json:"field1"`
	Field2 int `json:"field2"`
}

func NewGenericEventTest(ID string, field1 int, field2 int) *GenericEventTest {
	return &GenericEventTest{
		ID:     ID,
		Field1: field1,
		Field2: field2,
	}
}

func (e *GenericEventTest) FullQualifiedEventName() string {
	return e.ID
}

// ------------------------------------------------------
//              EventSubscriber Definition
// ------------------------------------------------------
//
// Note:
//       This subscriber will be used in the rest of the test.
//

type GenericTestEventSubscriber struct {
	events []*GenericEventTest
	t      *testing.T
}

func NewGenericTestEventSubscriber(t *testing.T) *GenericTestEventSubscriber {
	return &GenericTestEventSubscriber{
		t:      t,
		events: make([]*GenericEventTest, 0),
	}
}

func (g *GenericTestEventSubscriber) SubscribedTo() shared_domain_event_bus.DomainEvent {
	return &GenericEventTest{}
}

func (g *GenericTestEventSubscriber) ChannelName() string {
	return "EventTestForGenericTestSubscriber"
}

func (g *GenericTestEventSubscriber) Consume(data []byte) {
	var event GenericEventTest
	err := json.Unmarshal(data, &event)
	require.NoError(g.t, err, "There should not be an error unmarshaling events.")
	g.events = append(g.events, &event)
}

func (g *GenericTestEventSubscriber) RetrieveReceivedData() []*GenericEventTest {
	return g.events
}
