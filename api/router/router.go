package router

import (
	"database/sql"
	"gokafka/skill"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB, producerConfig sarama.SyncProducer) *gin.Engine {

	router := gin.Default()

	producer := skill.NewProducer(producerConfig)
	skillrepo := skill.NewSkillRepo(db, producer)
	skillHandler := skill.NewSkillHandler(skillrepo)

	v1 := router.Group("/api/v1")
	v1.GET("/skills/:key", skillHandler.GetSkillByKey)
	v1.GET("/skills", skillHandler.GetSkills)
	v1.POST("/skills", skillHandler.CreateSkill)
	v1.PUT("/skills/:key", skillHandler.UpdateSkill)
	v1.PATCH("/skills/:key/actions/name", skillHandler.UpdateSkillNameByKey)
	v1.PATCH("/skills/:key/actions/description", skillHandler.UpdateSkillDescriptionByKey)
	v1.PATCH("/skills/:key/actions/logo", skillHandler.UpdateSkillLogoByKey)
	v1.PATCH("/skills/:key/actions/tags", skillHandler.UpdateSkillTagsByKey)
	v1.DELETE("/skills/:key", skillHandler.DeleteSkill)

	return router
}
