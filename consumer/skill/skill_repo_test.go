package skill_test

import (
	"database/sql"
	"savedb/skill"
	"testing"

	"github.com/lib/pq"
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

func getCount(db *sql.DB) int {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM skill").Scan(&count)
	return count
}

func getData(db *sql.DB, key string) skill.Skill {
	rows := db.QueryRow("SELECT key, name, description, logo, tags FROM skill WHERE key = $1", key)
	var skill skill.Skill
	rows.Scan(&skill.Key, &skill.Name, &skill.Description, &skill.Logo, pq.Array(&skill.Tags))
	return skill
}

func TestCreateSkillRepo(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()
	mockRepo := skill.NewSkillRepo(db)
	// Act
	mockRepo.CreateSkill(skill.Skill{
		Key:         "key",
		Name:        "name",
		Description: "description",
		Logo:        "logo",
		Tags:        []string{"tag1", "tag2"},
	})
	// Assert
	if getCount(db) != 1 {
		t.Errorf("expected 1 row, got %d\n", getCount(db))
	}
}

func TestUpdateSkillRepo(t *testing.T) {
	// Arrange
	db := newMockDB()
	defer db.Close()
	mockRepo := skill.NewSkillRepo(db)

	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('key', 'name', 'description', 'logo', '{tag2,tag3}')")

	want := skill.Skill{
		Key:         "key",
		Name:        "name",
		Description: "description",
		Logo:        "logo",
		Tags:        []string{"tag3", "tag4"},
	}

	// Act
	mockRepo.UpdateSkill(want)

	// Assert
	if getCount(db) != 1 {
		t.Errorf("expected 1 row, got %d\n", getCount(db))
	}

	result := getData(db, "key")

	if result.Name != "name" {
		t.Errorf("")
	}
	assert.Equal(t, want.Key, result.Key)
	assert.Equal(t, want.Name, result.Name)
	assert.Equal(t, want.Description, result.Description)
	assert.Equal(t, want.Logo, result.Logo)
	assert.Equal(t, want.Tags, result.Tags)
}

func TestUpdateSkillNameByKey(t *testing.T) {
	//arange
	db := newMockDB()
	defer db.Close()
	mockRepo := skill.NewSkillRepo(db)

	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('key', 'name', 'description', 'logo', '{tag2,tag3}')")

	want := "nameupdate"
	//act
	mockRepo.UpdateSkillNameByKey("key", want)

	//assert
	result := getData(db, "key")
	if getCount(db) != 1 {
		t.Errorf("expected 1 row, got %d\n", getCount(db))
	}

	assert.Equal(t, want, result.Name)
}

func TestUpdateSkillDescriptionByKey(t *testing.T) {
	//arange
	db := newMockDB()
	defer db.Close()
	mockRepo := skill.NewSkillRepo(db)

	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('key', 'name', 'description', 'logo', '{tag2,tag3}')")

	want := "descriptionupdate"
	//act
	mockRepo.UpdateSkillDescriptionByKey("key", want)

	//assert
	result := getData(db, "key")
	if getCount(db) != 1 {
		t.Errorf("expected 1 row, got %d\n", getCount(db))
	}

	assert.Equal(t, want, result.Description)
}

func TestUpdateSkillLogoByKey(t *testing.T) {
	//arange
	db := newMockDB()
	defer db.Close()
	mockRepo := skill.NewSkillRepo(db)

	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('key', 'name', 'description', 'logo', '{tag2,tag3}')")

	want := "Logoupdate"
	//act
	mockRepo.UpdateSkillLogoByKey("key", want)

	//assert
	result := getData(db, "key")
	if getCount(db) != 1 {
		t.Errorf("expected 1 row, got %d\n", getCount(db))
	}

	assert.Equal(t, want, result.Logo)
}

func TestUpdateSkillTagsByKey(t *testing.T) {
	//arange
	db := newMockDB()
	defer db.Close()
	mockRepo := skill.NewSkillRepo(db)

	db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('key', 'name', 'description', 'logo', '{tag2,tag3}')")

	want := []string{"tag1", "tag4"}
	//act
	mockRepo.UpdateSkillTagsByKey("key", want)

	//assert
	result := getData(db, "key")
	if getCount(db) != 1 {
		t.Errorf("expected 1 row, got %d\n", getCount(db))
	}

	assert.Equal(t, want, result.Tags)
}

func TestDeleteSkill(t *testing.T) {
	t.Run("should delete skill", func(t *testing.T) {
		//arange
		db := newMockDB()
		defer db.Close()
		mockRepo := skill.NewSkillRepo(db)

		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('key', 'name', 'description', 'logo', '{tag2,tag3}')")

		//act
		mockRepo.DeleteSkillByKey("key")

		if getCount(db) != 0 {
			t.Errorf("expected 1 row, got %d\n", getCount(db))
		}
	})
	t.Run("should error when no delete skill", func(t *testing.T) {
		//arange
		db := newMockDB()
		defer db.Close()
		mockRepo := skill.NewSkillRepo(db)

		db.Exec("INSERT INTO skill (key, name, description, logo, tags) VALUES ('key', 'name', 'description', 'logo', '{tag2,tag3}')")

		//act
		mockRepo.DeleteSkillByKey("notkey")

		if getCount(db) != 1 {
			t.Errorf("expected 1 row, got %d\n", getCount(db))
		}
	})
}
