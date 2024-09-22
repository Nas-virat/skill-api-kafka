package skill

import (
	"bytes"
	"encoding/json"
	"gokafka/errs"
	"gokafka/response"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetSkillByKey(t *testing.T) {
	t.Run("should response skills by key", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "key", Value: "1"})

		skill := Skill{
			Key:         "1",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status: "success",
			Data:   skill,
		})

		//act
		handler.GetSkillByKey(c)

		//assert

		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.Bytes(), want)

	})

	t.Run("should response error when skill not found by key", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "key", Value: "1"})
		mock := &mockRepo{err: errs.NewError(http.StatusNotFound, "Skill not found")}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Skill not found",
		})

		//act
		handler.GetSkillByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusNotFound)
		assert.Equal(t, w.Body.Bytes(), want)
	})
}

func TestGetSkills(t *testing.T) {
	t.Run("should response skills from repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skills := []Skill{
			{
				Key:         "test-key",
				Name:        "test",
				Description: "test",
				Logo:        "test",
				Tags:        []string{"tag"},
			},
		}
		mock := &mockRepo{skills: skills}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status: "success",
			Data:   skills,
		})

		//act
		handler.GetSkills(c)
		//assert
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should response error when skill not found by key", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "key", Value: "1"})
		mock := &mockRepo{err: errs.NewError(http.StatusNotFound, "Skill not found")}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Skill not found",
		})

		//act
		handler.GetSkills(c)
		//assert
		assert.Equal(t, w.Code, http.StatusNotFound)
		assert.Equal(t, w.Body.Bytes(), want)
	})
}

func TestCreateSkill(t *testing.T) {
	t.Run("should response created skill from repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status: "success",
			Data:   skill,
		})
		//act
		handler.CreateSkill(c)
		//assert
		assert.Equal(t, w.Code, http.StatusCreated)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should response error when skill not found by key", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't bind payload",
		})

		//act
		handler.CreateSkill(c)
		//assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should response created skill from repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{err: errs.NewError(http.StatusInternalServerError, "could not create skill")}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "could not create skill",
		})
		//act
		handler.CreateSkill(c)
		//assert
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, w.Body.Bytes(), want)
	})
}

func TestUpdateSkill(t *testing.T) {
	t.Run("should response updated skill from repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPut, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status: "success",
			Data:   skill,
		})
		//act
		handler.UpdateSkill(c)
		//assert
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.Bytes(), want)
	})

	t.Run("should response error when no payload", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't bind payload",
		})

		//act
		handler.UpdateSkill(c)
		//assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should response err when repository return error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{err: errs.NewError(http.StatusInternalServerError, "could not update skill")}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPut, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "could not update skill",
		})
		//act
		handler.UpdateSkill(c)
		//assert
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, w.Body.Bytes(), want)
	})
}

func TestUpdateNameByKey(t *testing.T) {
	t.Run("should response updated skill from repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status: "success",
			Data:   skill,
		})
		//act
		handler.UpdateSkillNameByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should response error when no payload", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't bind payload",
		})

		//act
		handler.UpdateSkillNameByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should err updated when repository return error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{err: errs.NewError(http.StatusInternalServerError, "Can't update skill")}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't update skill",
		})
		//act
		handler.UpdateSkillNameByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, w.Body.Bytes(), want)
	})
}

func TestUpdateDescriptionByKey(t *testing.T) {
	t.Run("should response updated skill from repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status: "success",
			Data:   skill,
		})
		//act
		handler.UpdateSkillDescriptionByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should response error when no payload", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't bind payload",
		})

		//act
		handler.UpdateSkillDescriptionByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should err updated when repository return error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{err: errs.NewError(http.StatusInternalServerError, "Can't update skill")}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't update skill",
		})
		//act
		handler.UpdateSkillDescriptionByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, w.Body.Bytes(), want)
	})
}
func TestUpdateTagsByKey(t *testing.T) {
	t.Run("should response updated skill from repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status: "success",
			Data:   skill,
		})
		//act
		handler.UpdateSkillTagsByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should response error when no payload", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't bind payload",
		})

		//act
		handler.UpdateSkillTagsByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should err updated when repository return error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{err: errs.NewError(http.StatusInternalServerError, "Can't update skill")}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't update skill",
		})
		//act
		handler.UpdateSkillTagsByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, w.Body.Bytes(), want)
	})
}

func TestUpdateLogoByKey(t *testing.T) {
	t.Run("should response updated skill from repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/", bytes.NewReader(body))
		want, _ := json.Marshal(response.Response{
			Status: "success",
			Data:   skill,
		})
		//act
		handler.UpdateSkillLogoByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should response error when no payload", func(t *testing.T) {
		//arrange
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{skill: skill}
		handler := NewSkillHandler(mock)

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't bind payload",
		})

		//act
		handler.UpdateSkillLogoByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusBadRequest)
		assert.Equal(t, w.Body.Bytes(), want)
	})
	t.Run("should err updated when repository return error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		skill := Skill{
			Key:         "test-key",
			Name:        "test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"tag"},
		}
		mock := &mockRepo{err: errs.NewError(http.StatusInternalServerError, "Can't update skill")}
		handler := NewSkillHandler(mock)
		body, _ := json.Marshal(skill)
		c.Request, _ = http.NewRequest(http.MethodPatch, "/", bytes.NewReader(body))

		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't update skill",
		})
		//act
		handler.UpdateSkillLogoByKey(c)
		//assert
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, w.Body.Bytes(), want)
	})
}

func TestDeleteSkill(t *testing.T) {
	t.Run("should delete skill", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "test-key", Value: "test"}}
		want, _ := json.Marshal(response.Response{
			Status:  "success",
			Message: "Skill deleted",
		})
		mock := &mockRepo{}
		handler := NewSkillHandler(mock)
		//act
		handler.DeleteSkill(c)
		//assert
		assert.Equal(t, w.Code, http.StatusOK)
		assert.Equal(t, want, w.Body.Bytes())
	})

	t.Run("should return error when delete skill repo fail", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = []gin.Param{{Key: "test-key", Value: "test"}}
		want, _ := json.Marshal(response.Response{
			Status:  "error",
			Message: "Can't delete skill",
		})
		mock := &mockRepo{err: errs.NewError(http.StatusInternalServerError, "Can't delete skill")}
		handler := NewSkillHandler(mock)
		//act
		handler.DeleteSkill(c)
		//assert
		assert.Equal(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, want, w.Body.Bytes())
	})

}
