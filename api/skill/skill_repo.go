package skill

import (
	"database/sql"
	"gokafka/errs"
	"net/http"

	"github.com/lib/pq"
)

type skillRepo struct {
	db       *sql.DB
	producer SkillProcuer
}

type SkillRepo interface {
	GetSkillByKey(key string) (*Skill, error)
	GetSkills() ([]Skill, error)
	CreateSkill(skill Skill) (*Skill, error)
	UpdateSkill(key string, skill Skill) (*Skill, error)
	UpdateSkillNameByKey(key string, name string) (*Skill, error)
	UpdateSkillDescriptionByKey(key string, description string) (*Skill, error)
	UpdateSkillLogoByKey(key string, logo string) (*Skill, error)
	UpdateSkillTagsByKey(key string, tags []string) (*Skill, error)
	DeleteSkillByKey(key string) error
}

func ScanSkill(rows *sql.Row, skill *Skill) error {
	err := rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
	return err
}

func NewSkillRepo(db *sql.DB, producer SkillProcuer) *skillRepo {
	return &skillRepo{db: db, producer: producer}
}

func (r *skillRepo) GetSkillByKey(key string) (*Skill, error) {
	skill := Skill{}
	query := "SELECT key, name, description, logo, tags FROM skill WHERE key=$1"
	record := r.db.QueryRow(query, key)
	err := ScanSkill(record, &skill)
	if err != nil {
		return &Skill{}, errs.NewError(http.StatusNotFound, "Skill not found")
	}
	return &skill, nil
}

func (r *skillRepo) GetSkills() ([]Skill, error) {

	skills := []Skill{}
	query := "SELECT key, name, description, logo, tags FROM skill"
	records, err := r.db.Query(query)
	if err != nil {
		return []Skill{}, errs.NewError(http.StatusInternalServerError, err.Error())
	}
	for records.Next() {
		skill := Skill{}
		err := records.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
		if err != nil {
			return []Skill{}, errs.NewError(http.StatusInternalServerError, err.Error())
		}
		skills = append(skills, skill)
	}

	return skills, nil
}

func (r *skillRepo) CreateSkill(skill Skill) (*Skill, error) {

	if err := r.producer.PublishMessage(CreateSkillAction, skill); err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	return &skill, nil
}

func (r *skillRepo) UpdateSkill(key string, skill Skill) (*Skill, error) {

	if skill.Key != key {
		return nil, errs.NewError(http.StatusBadRequest, "Key does not match")
	}

	if err := r.producer.PublishMessage(UpdateSkillAction, skill); err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	return &skill, nil
}

func (r *skillRepo) UpdateSkillNameByKey(key string, name string) (*Skill, error) {

	nameUpdateMessage := NameUpdateMessage{
		Key:  key,
		Name: name,
	}

	err := r.producer.PublishMessage(UpdateNameAction, nameUpdateMessage)

	if err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	updateSkill, err := r.GetSkillByKey(key)
	if err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	updateSkill.Name = name

	return updateSkill, nil
}

func (r *skillRepo) UpdateSkillDescriptionByKey(key string, description string) (*Skill, error) {

	descriptionUpdateMessage := DescriptionUpdateMessage{
		Key:         key,
		Description: description,
	}

	err := r.producer.PublishMessage(UpdateDescAction, descriptionUpdateMessage)

	if err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	updateSkill, err := r.GetSkillByKey(key)

	if err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	updateSkill.Description = description

	return updateSkill, nil
}

func (r *skillRepo) UpdateSkillLogoByKey(key string, logo string) (*Skill, error) {

	logoUpdateMessage := LogoUpdateMessage{
		Key:  key,
		Logo: logo,
	}

	err := r.producer.PublishMessage(UpdateLogoAction, logoUpdateMessage)

	if err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	updateSkill, err := r.GetSkillByKey(key)

	if err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	updateSkill.Logo = logo

	return updateSkill, nil
}

func (r *skillRepo) UpdateSkillTagsByKey(key string, tags []string) (*Skill, error) {

	tagsUpdateMessage := TagsUpdateMessage{
		Key:  key,
		Tags: tags,
	}
	err := r.producer.PublishMessage(UpdateTagsAction, tagsUpdateMessage)

	if err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	updateSkill, err := r.GetSkillByKey(key)

	if err != nil {
		return nil, errs.NewError(http.StatusInternalServerError, err.Error())
	}

	updateSkill.Tags = tags

	return updateSkill, nil
}

func (r *skillRepo) DeleteSkillByKey(key string) error {
	query := "DELETE FROM skill WHERE key=$1"
	result, err := r.db.Exec(query, key)
	if err != nil {
		return errs.NewError(http.StatusInternalServerError, "not be able to delete skill")
	}

	resultRows, err := result.RowsAffected()
	if err != nil {
		return errs.NewError(http.StatusInternalServerError, "not be able to delete skill")
	}
	if resultRows == 0 {
		return errs.NewError(http.StatusInternalServerError, "not be able to delete skill")
	}

	return nil
}
