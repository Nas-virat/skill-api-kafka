package skill

import (
	"log/slog"

	"github.com/IBM/sarama"
)

type SkillConsumer struct {
	ready             chan struct{}
	skillEventHandler SkillEventHandler
}

func NewConsumerGroup(skillEventHandler SkillEventHandler) *SkillConsumer {
	return &SkillConsumer{
		ready:             make(chan struct{}),
		skillEventHandler: skillEventHandler,
	}
}

func (consumer *SkillConsumer) Setup(_ sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}
func (*SkillConsumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (s *SkillConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The ConsumeClaim itself is called within a goroutine, see:
	// https://github.com/IBM/sarama/blob/main/consumer_group.go#L27-L29
consume:
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				slog.Info("message channel was closed")
				break consume
			}
			s.skillEventHandler.ProcessMessage(msg)
			sess.MarkMessage(msg, "")
		// Should return when session.Context() is done.
		// If not, will raise ErrRebalanceInProgress or read tcp <ip>:<port>: i/o timeout when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-sess.Context().Done():
			break consume
		}
	}
	return sess.Context().Err()
}

func (c *SkillConsumer) NewReady() {
	c.ready = make(chan struct{})
}

func (c *SkillConsumer) Ready() <-chan struct{} {
	return c.ready
}
