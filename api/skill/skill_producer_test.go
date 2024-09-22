package skill

import (
	"errors"
	"testing"

	"github.com/IBM/sarama"
)

type mockSyncProcuer struct {
	partion   int32
	offset    int64
	err       error
	wasCalled bool
}

func (m *mockSyncProcuer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	m.wasCalled = true
	return m.partion, m.offset, m.err
}

func (m *mockSyncProcuer) SendMessages(msgs []*sarama.ProducerMessage) error {
	return m.err
}

func (m *mockSyncProcuer) Close() error {
	return m.err
}

func (m *mockSyncProcuer) TxnStatus() sarama.ProducerTxnStatusFlag {
	return sarama.ProducerTxnFlagUninitialized
}

func (m *mockSyncProcuer) IsTransactional() bool {
	return true
}

func (m *mockSyncProcuer) BeginTxn() error {
	return m.err
}

func (m *mockSyncProcuer) CommitTxn() error {
	return m.err
}

func (m *mockSyncProcuer) AbortTxn() error {
	return m.err
}

func (m *mockSyncProcuer) AddOffsetsToTxn(offsets map[string][]*sarama.PartitionOffsetMetadata, groupId string) error {
	return m.err
}

func (m *mockSyncProcuer) AddMessageToTxn(msg *sarama.ConsumerMessage, groupId string, metadata *string) error {
	return m.err
}

func TestPublishMessage(t *testing.T) {

	t.Run("should not return error when producer publish message successfully", func(t *testing.T) {
		//arrange
		mockSyncProcuer := &mockSyncProcuer{
			partion: 3,
			offset:  3453466,
			err:     nil,
		}
		producer := NewProducer(mockSyncProcuer)

		action := CreateSkillAction
		payload := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}

		//act
		err := producer.PublishMessage(action, payload)

		//assert
		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		if !mockSyncProcuer.wasCalled {
			t.Errorf("Expected mockSyncProducer to send message to call be not call")
		}
	})
	t.Run("should return error when producer publish message unsuccessfully", func(t *testing.T) {
		//arrange
		mockSyncProcuer := &mockSyncProcuer{
			partion: 3,
			offset:  3453466,
			err:     errors.New("error"),
		}
		producer := NewProducer(mockSyncProcuer)

		action := CreateSkillAction
		payload := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}

		//act
		err := producer.PublishMessage(action, payload)

		//assert
		if err == nil {
			t.Errorf("expected no error but got %v", err)
		}

		if !mockSyncProcuer.wasCalled {
			t.Errorf("Expected mockSyncProducer to send message to call be not call")
		}
	})
}
