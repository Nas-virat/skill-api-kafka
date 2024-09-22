package skill

import (
	"database/sql"

	"github.com/lib/pq"
)

type skillRepo struct {
	db *sql.DB
}

type SkillRepo interface {
	CreateSkill(skill Skill) (*Skill, error)
	UpdateSkill(skill Skill) (*Skill, error)
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

func NewSkillRepo(db *sql.DB) SkillRepo {
	return &skillRepo{db: db}
}

func (r *skillRepo) CreateSkill(skill Skill) (*Skill, error) {
	createdSkill := Skill{}
	query := "INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5) RETURNING key, name, description, logo, tags"
	record := r.db.QueryRow(query, skill.Key, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags))
	err := ScanSkill(record, &createdSkill)
	return &createdSkill, err
}

func (r *skillRepo) UpdateSkill(skill Skill) (*Skill, error) {
	updateSkill := Skill{}
	query := "UPDATE skill SET name=$1, description=$2, logo=$3, tags=$4 WHERE key=$5 RETURNING key, name, description, logo, tags"
	record := r.db.QueryRow(query, skill.Name, skill.Description, skill.Logo, pq.Array(skill.Tags), skill.Key)
	err := ScanSkill(record, &updateSkill)
	return &updateSkill, err
}

func (r *skillRepo) UpdateSkillNameByKey(key string, name string) (*Skill, error) {
	updateSkill := Skill{}
	query := "UPDATE skill SET name=$1 WHERE key=$2 RETURNING key, name, description, logo, tags"
	record := r.db.QueryRow(query, name, key)
	err := ScanSkill(record, &updateSkill)
	return &updateSkill, err
}

func (r *skillRepo) UpdateSkillDescriptionByKey(key string, description string) (*Skill, error) {
	updatedSkill := Skill{}
	query := "UPDATE skill SET description=$1 WHERE key=$2 RETURNING key, name, description, logo, tags"
	record := r.db.QueryRow(query, description, key)
	err := ScanSkill(record, &updatedSkill)
	return &updatedSkill, err
}

func (r *skillRepo) UpdateSkillLogoByKey(key string, logo string) (*Skill, error) {
	updatedSkill := Skill{}
	query := "UPDATE skill SET logo=$1 WHERE key=$2 RETURNING key, name, description, logo, tags"
	record := r.db.QueryRow(query, logo, key)
	err := ScanSkill(record, &updatedSkill)
	return &updatedSkill, err
}

func (r *skillRepo) UpdateSkillTagsByKey(key string, tags []string) (*Skill, error) {
	updatedSkill := Skill{}
	query := "UPDATE skill SET tags=$1 WHERE key=$2 RETURNING key, name, description, logo, tags"
	record := r.db.QueryRow(query, pq.Array(tags), key)
	err := ScanSkill(record, &updatedSkill)
	return &updatedSkill, err
}

func (r *skillRepo) DeleteSkillByKey(key string) error {
	query := "DELETE FROM skill WHERE key=$1"
	result, err := r.db.Exec(query, key)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}
