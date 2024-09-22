package skill

import (
	"encoding/json"
	"log"
	"os"

	"github.com/IBM/sarama"
)

type SkillAction string

const (
	CreateSkillAction SkillAction = "create"
	UpdateSkillAction SkillAction = "update"
	DeleteSkillAction SkillAction = "delete"
	UpdateNameAction  SkillAction = "update_name"
	UpdateDescAction  SkillAction = "update_desc"
	UpdateLogoAction  SkillAction = "update_logo"
	UpdateTagsAction  SkillAction = "update_tags"
)

type skillProcuer struct {
	producer sarama.SyncProducer
}

type SkillProcuer interface {
	PublishMessage(action SkillAction, payload interface{}) error
}

func NewProducer(producer sarama.SyncProducer) skillProcuer {
	return skillProcuer{producer: producer}
}

func (p skillProcuer) PublishMessage(action SkillAction, payload interface{}) error {

	objBytes, _ := json.Marshal(payload)

	msg := &sarama.ProducerMessage{Topic: os.Getenv("TOPIC"), Key: sarama.StringEncoder(action), Value: sarama.ByteEncoder(objBytes)}
	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
		return err
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		return err
	}
}
