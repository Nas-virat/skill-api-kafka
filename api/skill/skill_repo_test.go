package skill_test

import (
	"database/sql"
	"errors"
	"gokafka/skill"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func newMockDB() *sql.DB {
	db, _ := sql.Open("sqlite", "file:skill?mode=memory&cache=shared")
	q := `
		CREATE TABLE IF NOT EXISTS skill (
		key TEXT PRIMARY KEY,
		name TEXT NOT NULL DEFAULT '',
		description TEXT NOT NULL DEFAULT '',
		logo TEXT NOT NULL DEFAULT '',
		tags TEXT [] NOT NULL DEFAULT '{}'
	);
	`
	db.Exec(q)
	return db
}

type MockProducer struct {
	err error
}

func (p *MockProducer) PublishMessage(action skill.SkillAction, payload interface{}) error {
	return p.err
}
func TestGetSkillByKeyRepo(t *testing.T) {

	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'go', 'description', 'logo', '{tag2,tag3}')")

		producer := &MockProducer{}

		repo := skill.NewSkillRepo(db, producer)

		//act
		skill, err := repo.GetSkillByKey("go")

		//assert
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}

		assert.Equal(t, skill.Key, "go")
		assert.Equal(t, skill.Name, "go")
		assert.Equal(t, skill.Description, "description")
		assert.Equal(t, skill.Logo, "logo")

	})
}

func TestGetSkillsRepo(t *testing.T) {

	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'go', 'description', 'logo', '{tag2,tag3}')")
		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('python', 'python', 'dec', 'lo', '{tag2,tag3}')")

		producer := &MockProducer{}

		repo := skill.NewSkillRepo(db, producer)

		//act
		skills, err := repo.GetSkills()

		//assert
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}

		assert.Equal(t, skills[0].Key, "go")
		assert.Equal(t, skills[0].Name, "go")
		assert.Equal(t, skills[0].Description, "description")
		assert.Equal(t, skills[0].Logo, "logo")

		assert.Equal(t, skills[1].Key, "python")
		assert.Equal(t, skills[1].Name, "python")
		assert.Equal(t, skills[1].Description, "dec")
		assert.Equal(t, skills[1].Logo, "lo")

	})
}
func TestCreateSkillRepo(t *testing.T) {

	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		producer := &MockProducer{err: nil}

		repo := skill.NewSkillRepo(db, producer)

		skill := skill.Skill{
			Key:         "go",
			Name:        "go",
			Description: "description",
			Logo:        "logo",
		}
		//act
		resultSkill, err := repo.CreateSkill(skill)

		//assert
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}

		assert.Equal(t, "go", resultSkill.Key)
		assert.Equal(t, "go", resultSkill.Name)
		assert.Equal(t, "description", resultSkill.Description)
		assert.Equal(t, "logo", resultSkill.Logo)

	})
	t.Run("should return error when fail publish data", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		producer := &MockProducer{err: errors.New("error")}

		repo := skill.NewSkillRepo(db, producer)

		skill := skill.Skill{
			Key:         "go",
			Name:        "go",
			Description: "description",
			Logo:        "logo",
		}
		//act
		resultSkill, err := repo.CreateSkill(skill)

		//assert
		if err == nil {
			t.Errorf("expected nil but got %v", err)
		}

		if resultSkill != nil {
			t.Errorf("expected resultSkill to be nil but got %v", resultSkill)
		}

	})
}

func TestUpdateSkillRepo(t *testing.T) {

	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		producer := &MockProducer{err: nil}

		repo := skill.NewSkillRepo(db, producer)

		skill := skill.Skill{
			Key:         "go",
			Name:        "go",
			Description: "description",
			Logo:        "logo",
		}
		//act
		resultSkill, err := repo.UpdateSkill("python", skill)

		//assert
		if err == nil {
			t.Errorf("expected to be nil but got %v", err)
		}

		if resultSkill != nil {
			t.Errorf("expected resultSkill to be nil but got %v", resultSkill)
		}

	})
	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		producer := &MockProducer{err: nil}

		repo := skill.NewSkillRepo(db, producer)

		skill := skill.Skill{
			Key:         "go",
			Name:        "go",
			Description: "description",
			Logo:        "logo",
		}
		//act
		resultSkill, err := repo.UpdateSkill("go", skill)

		//assert
		if err != nil {
			t.Errorf("expected error to be nil but got %v", err)
		}

		assert.Equal(t, "go", resultSkill.Key)
		assert.Equal(t, "go", resultSkill.Name)
		assert.Equal(t, "description", resultSkill.Description)
		assert.Equal(t, "logo", resultSkill.Logo)

	})
	t.Run("should return error when fail publish data", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		producer := &MockProducer{err: errors.New("error")}

		repo := skill.NewSkillRepo(db, producer)

		skill := skill.Skill{
			Key:         "go",
			Name:        "go",
			Description: "description",
			Logo:        "logo",
		}
		//act
		resultSkill, err := repo.UpdateSkill("go", skill)

		//assert
		if err == nil {
			t.Errorf("expected nil but got %v", err)
		}

		if resultSkill != nil {
			t.Errorf("expected resultSkill to be nil but got %v", resultSkill)
		}

	})
}

func TestUpdateSkillNameByKeyRepo(t *testing.T) {
	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'go', 'description', 'logo', '{tag2,tag3}')")

		producer := &MockProducer{err: nil}

		repo := skill.NewSkillRepo(db, producer)

		//act
		resultSkill, err := repo.UpdateSkillNameByKey("go", "gopher")

		//assert
		if err != nil {
			t.Errorf("expected to be nil but got %v", err)
		}

		assert.Equal(t, "go", resultSkill.Key)
		assert.Equal(t, "gopher", resultSkill.Name)
		assert.Equal(t, "description", resultSkill.Description)
		assert.Equal(t, "logo", resultSkill.Logo)

	})
}

func TestUpdateSkillDescriptionByKeyRepo(t *testing.T) {
	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'go', 'description', 'logo', '{tag2,tag3}')")

		producer := &MockProducer{err: nil}

		repo := skill.NewSkillRepo(db, producer)

		//act
		resultSkill, err := repo.UpdateSkillDescriptionByKey("go", "gopher description")

		//assert
		if err != nil {
			t.Errorf("expected to be nil but got %v", err)
		}

		assert.Equal(t, "go", resultSkill.Key)
		assert.Equal(t, "go", resultSkill.Name)
		assert.Equal(t, "gopher description", resultSkill.Description)
		assert.Equal(t, "logo", resultSkill.Logo)

	})
}

func TestUpdateSkillLogoByKeyRepo(t *testing.T) {
	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'go', 'description', 'logo', '{tag2,tag3}')")

		producer := &MockProducer{err: nil}

		repo := skill.NewSkillRepo(db, producer)

		//act
		resultSkill, err := repo.UpdateSkillLogoByKey("go", "gopher logo")

		//assert
		if err != nil {
			t.Errorf("expected to be nil but got %v", err)
		}

		assert.Equal(t, "go", resultSkill.Key)
		assert.Equal(t, "go", resultSkill.Name)
		assert.Equal(t, "description", resultSkill.Description)
		assert.Equal(t, "gopher logo", resultSkill.Logo)

	})
}

func TestUpdateSkillTagsByKeyRepo(t *testing.T) {
	t.Run("should return skill when key is exist", func(t *testing.T) {

		//arange
		db := newMockDB()
		defer db.Close()

		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'go', 'description', 'logo', '{tag2,tag3}')")

		producer := &MockProducer{err: nil}

		repo := skill.NewSkillRepo(db, producer)

		//act
		resultSkill, err := repo.UpdateSkillTagsByKey("go", []string{"gopher logo"})

		//assert
		if err != nil {
			t.Errorf("expected to be nil but got %v", err)
		}

		assert.Equal(t, "go", resultSkill.Key)
		assert.Equal(t, "go", resultSkill.Name)
		assert.Equal(t, "description", resultSkill.Description)
		assert.Equal(t, "logo", resultSkill.Logo)
		assert.Equal(t, []string{"gopher logo"}, resultSkill.Tags)

	})
}

func TestDeleteSkillRepo(t *testing.T) {
	//arange
	db := newMockDB()
	defer db.Close()

	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('go', 'go', 'description', 'logo', '{tag2,tag3}')")

	producer := &MockProducer{err: nil}

	repo := skill.NewSkillRepo(db, producer)

	//act
	err := repo.DeleteSkillByKey("go")

	//assert
	if err != nil {
		t.Errorf("expected to be nil but got %v", err)
	}
}
