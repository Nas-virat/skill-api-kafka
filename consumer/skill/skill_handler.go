package skill

import (
	"encoding/json"
	"log"

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

type SkillEventHandler interface {
	ProcessMessage(msg *sarama.ConsumerMessage) error
	createSkillHandler(msg *sarama.ConsumerMessage) error
	updateSkillHandler(msg *sarama.ConsumerMessage) error
	updateNameHandler(msg *sarama.ConsumerMessage) error
	updateDescriptionHandler(msg *sarama.ConsumerMessage) error
	updateLogoHandler(msg *sarama.ConsumerMessage) error
	updateTagHandler(msg *sarama.ConsumerMessage) error
}

type skillEventHandler struct {
	skillRepo SkillRepo
}

func NewSkillEventHandler(skillRepo SkillRepo) *skillEventHandler {
	return &skillEventHandler{skillRepo: skillRepo}
}

func (s *skillEventHandler) ProcessMessage(msg *sarama.ConsumerMessage) error {
	var err error = nil
	log.Printf("Message key:%s topic:%q partition:%d offset:%d \n", string(msg.Key), msg.Topic, msg.Partition, msg.Offset)

	switch string(msg.Key) {
	case string(CreateSkillAction):
		err = s.createSkillHandler(msg)
	case string(UpdateSkillAction):
		err = s.updateSkillHandler(msg)
	case string(UpdateNameAction):
		err = s.updateNameHandler(msg)
	case string(UpdateDescAction):
		err = s.updateDescriptionHandler(msg)
	case string(UpdateLogoAction):
		err = s.updateLogoHandler(msg)
	case string(UpdateTagsAction):
		err = s.updateTagHandler(msg)
	default:
		log.Println("Unknown")
	}
	return err
}

func (s *skillEventHandler) createSkillHandler(msg *sarama.ConsumerMessage) error {
	skill := Skill{}
	err := json.Unmarshal(msg.Value, &skill)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	_, err = s.skillRepo.CreateSkill(skill)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}

func (s *skillEventHandler) updateSkillHandler(msg *sarama.ConsumerMessage) error {
	skill := Skill{}
	err := json.Unmarshal(msg.Value, &skill)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	_, err = s.skillRepo.UpdateSkill(skill)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}

func (s *skillEventHandler) updateNameHandler(msg *sarama.ConsumerMessage) error {
	nameUpdateMessage := NameUpdateMessage{}
	err := json.Unmarshal(msg.Value, &nameUpdateMessage)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	_, err = s.skillRepo.UpdateSkillNameByKey(nameUpdateMessage.Key, nameUpdateMessage.Name)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}

func (s *skillEventHandler) updateDescriptionHandler(msg *sarama.ConsumerMessage) error {
	descriptionUpdateMessage := DescriptionUpdateMessage{}
	err := json.Unmarshal(msg.Value, &descriptionUpdateMessage)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	_, err = s.skillRepo.UpdateSkillDescriptionByKey(descriptionUpdateMessage.Key, descriptionUpdateMessage.Description)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}

func (s *skillEventHandler) updateLogoHandler(msg *sarama.ConsumerMessage) error {
	logoUpdateMessage := LogoUpdateMessage{}
	err := json.Unmarshal(msg.Value, &logoUpdateMessage)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	_, err = s.skillRepo.UpdateSkillLogoByKey(logoUpdateMessage.Key, logoUpdateMessage.Logo)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}

func (s *skillEventHandler) updateTagHandler(msg *sarama.ConsumerMessage) error {
	tagsUpdateMessage := TagsUpdateMessage{}
	err := json.Unmarshal(msg.Value, &tagsUpdateMessage)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	_, err = s.skillRepo.UpdateSkillTagsByKey(tagsUpdateMessage.Key, tagsUpdateMessage.Tags)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return err
	}
	return nil
}
