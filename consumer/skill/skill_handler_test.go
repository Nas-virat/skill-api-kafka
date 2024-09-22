package skill

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/IBM/sarama"
)

type MockSkillRepository struct {
	skill     Skill
	err       error
	wasCalled bool
}

func (mockRepo *MockSkillRepository) CreateSkill(skill Skill) (*Skill, error) {
	mockRepo.wasCalled = true
	return &mockRepo.skill, mockRepo.err
}
func (mockRepo *MockSkillRepository) UpdateSkill(skill Skill) (*Skill, error) {
	mockRepo.wasCalled = true
	return &mockRepo.skill, mockRepo.err
}
func (mockRepo *MockSkillRepository) UpdateSkillNameByKey(key string, name string) (*Skill, error) {
	mockRepo.wasCalled = true
	return &mockRepo.skill, mockRepo.err
}
func (mockRepo *MockSkillRepository) UpdateSkillDescriptionByKey(key string, description string) (*Skill, error) {
	mockRepo.wasCalled = true
	return &mockRepo.skill, mockRepo.err
}
func (mockRepo *MockSkillRepository) UpdateSkillLogoByKey(key string, logo string) (*Skill, error) {
	mockRepo.wasCalled = true
	return &mockRepo.skill, mockRepo.err
}
func (mockRepo *MockSkillRepository) UpdateSkillTagsByKey(key string, tags []string) (*Skill, error) {
	mockRepo.wasCalled = true
	return &mockRepo.skill, mockRepo.err
}
func (mockRepo *MockSkillRepository) DeleteSkillByKey(key string) error {
	mockRepo.wasCalled = true
	return mockRepo.err
}

func TestCreateSkill(t *testing.T) {
	t.Run("should not return error when skill is created successfully", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{skill: skill, err: nil}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(CreateSkillAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
	t.Run("should return error when can repo return error", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{err: errors.New("")}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(CreateSkillAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
	t.Run("should return error when can not json marshal to skill", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}

		fakeSkill := `{"key":"1","name":"test","description":"test","logo":"test","tags":["tag"]}}`
		mockSkillRepo := &MockSkillRepository{
			skill:     skill,
			err:       nil,
			wasCalled: false,
		}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(fakeSkill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(CreateSkillAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be false")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})

}

func TestUpdateSkill(t *testing.T) {
	t.Run("should not return error when skill is created successfully", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{skill: skill, err: nil}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateSkillAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})

	t.Run("should return error when can repo return error", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{err: errors.New("")}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateSkillAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
	t.Run("should return error when can not json marshal to skill", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}

		fakeSkill := `{"key":"1","name":"test","description":"test","logo":"test","tags":["tag"]}}`
		mockSkillRepo := &MockSkillRepository{
			skill:     skill,
			err:       nil,
			wasCalled: false,
		}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(fakeSkill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateSkillAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be false")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
}

func TestUpdateName(t *testing.T) {
	t.Run("should not return error when skill is created successfully", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{skill: skill, err: nil}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateNameAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})

	t.Run("should return error when can repo return error", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{err: errors.New("")}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateNameAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
	t.Run("should return error when can not json marshal to skill", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}

		fakeSkill := `{"key":"1","name":"test","description":"test","logo":"test","tags":["tag"]}}`
		mockSkillRepo := &MockSkillRepository{
			skill:     skill,
			err:       nil,
			wasCalled: false,
		}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(fakeSkill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateNameAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be false")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
}

func TestUpdateDesc(t *testing.T) {
	t.Run("should not return error when skill is created successfully", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{skill: skill, err: nil}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateDescAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})

	t.Run("should return error when can repo return error", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{err: errors.New("")}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateDescAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
	t.Run("should return error when can not json marshal to skill", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}

		fakeSkill := `{"key":"1","name":"test","description":"test","logo":"test","tags":["tag"]}}`
		mockSkillRepo := &MockSkillRepository{
			skill:     skill,
			err:       nil,
			wasCalled: false,
		}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(fakeSkill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateDescAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be false")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
}

func TestUpdateLogo(t *testing.T) {
	t.Run("should not return error when skill is created successfully", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{skill: skill, err: nil}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateLogoAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})

	t.Run("should return error when can repo return error", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{err: errors.New("")}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateLogoAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
	t.Run("should return error when can not json marshal to skill", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}

		fakeSkill := `{"key":"1","name":"test","description":"test","logo":"test","tags":["tag"]}}`
		mockSkillRepo := &MockSkillRepository{
			skill:     skill,
			err:       nil,
			wasCalled: false,
		}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(fakeSkill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateLogoAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be false")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
}

func TestUpdateTag(t *testing.T) {
	t.Run("should not return error when skill is created successfully", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{skill: skill, err: nil}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateTagsAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})

	t.Run("should return error when can repo return error", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mockSkillRepo := &MockSkillRepository{err: errors.New("")}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(skill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateTagsAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if !mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be true")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
	t.Run("should return error when can not json marshal to skill", func(t *testing.T) {

		//arange
		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}

		fakeSkill := `{"key":"1","name":"test","description":"test","logo":"test","tags":["tag"]}}`
		mockSkillRepo := &MockSkillRepository{
			skill:     skill,
			err:       nil,
			wasCalled: false,
		}
		skillEventHandler := NewSkillEventHandler(mockSkillRepo)

		value, _ := json.Marshal(fakeSkill)

		msg := &sarama.ConsumerMessage{
			Key:       []byte(UpdateTagsAction),
			Value:     value,
			Topic:     "skills",
			Partition: 0,
			Offset:    123456,
		}

		//act
		err := skillEventHandler.ProcessMessage(msg)

		//assert
		if mockSkillRepo.wasCalled {
			t.Error("expected wasCalled to be false")
		}
		if err == nil {
			t.Errorf("expected error to be nil but got %v", err)
		}
	})
}
