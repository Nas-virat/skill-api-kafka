package skill_test

import (
	"testing"

	"github.com/IBM/sarama"
)

type MockEventHandler struct {
	err error
}

func (m *MockEventHandler) ProcessMessage(msg *sarama.ConsumerMessage) error {
	return m.err
}
func (m *MockEventHandler) createSkillHandler(msg *sarama.ConsumerMessage) error {
	return m.err
}
func (m *MockEventHandler) updateSkillHandler(msg *sarama.ConsumerMessage) error {

	return m.err
}
func (m *MockEventHandler) updateNameHandler(msg *sarama.ConsumerMessage) error {

	return m.err
}
func (m *MockEventHandler) updateDescriptionHandler(msg *sarama.ConsumerMessage) error {

	return m.err
}
func (m *MockEventHandler) updateLogoHandler(msg *sarama.ConsumerMessage) error {

	return m.err
}
func (m *MockEventHandler) updateTagHandler(msg *sarama.ConsumerMessage) error {

	return m.err
}

func TestConsumer(t *testing.T) {

}
