package skill

import (
	"gokafka/errs"
	"gokafka/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type skillHandler struct {
	skillrepo SkillRepo
}

func NewSkillHandler(skillrepo SkillRepo) *skillHandler {
	return &skillHandler{skillrepo: skillrepo}
}

func (h *skillHandler) GetSkillByKey(ctx *gin.Context) {
	key := ctx.Param("key")
	skill, err := h.skillrepo.GetSkillByKey(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, skill)
}

func (h *skillHandler) GetSkills(ctx *gin.Context) {
	skills, err := h.skillrepo.GetSkills()
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, skills)
}

func (h *skillHandler) CreateSkill(ctx *gin.Context) {
	skill := Skill{}
	err := ctx.BindJSON(&skill)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}

	skillResult, err := h.skillrepo.CreateSkill(skill)

	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusCreated, skillResult)
}

func (h *skillHandler) UpdateSkill(ctx *gin.Context) {
	skill := Skill{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&skill)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}

	skillResult, err := h.skillrepo.UpdateSkill(key, skill)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.Success(ctx, http.StatusOK, skillResult)
}

func (h *skillHandler) UpdateSkillNameByKey(ctx *gin.Context) {
	req := NameUpdateRequest{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&req)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}
	skill, err := h.skillrepo.UpdateSkillNameByKey(key, req.Name)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, http.StatusOK, skill)
}

func (h *skillHandler) UpdateSkillDescriptionByKey(ctx *gin.Context) {
	req := DescriptionUpdateRequest{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&req)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}
	skill, err := h.skillrepo.UpdateSkillDescriptionByKey(key, req.Description)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, http.StatusOK, skill)
}

func (h *skillHandler) UpdateSkillLogoByKey(ctx *gin.Context) {
	req := LogoUpdateRequest{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&req)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}
	skill, err := h.skillrepo.UpdateSkillLogoByKey(key, req.Logo)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, http.StatusOK, skill)
}

func (h *skillHandler) UpdateSkillTagsByKey(ctx *gin.Context) {
	req := TagsUpdateRequest{}
	key := ctx.Param("key")
	err := ctx.BindJSON(&req)
	if err != nil {
		response.Error(ctx, errs.NewError(http.StatusBadRequest, "Can't bind payload"))
		return
	}
	skill, err := h.skillrepo.UpdateSkillTagsByKey(key, req.Tags)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, http.StatusOK, skill)
}

func (h *skillHandler) DeleteSkill(ctx *gin.Context) {
	key := ctx.Param("key")
	err := h.skillrepo.DeleteSkillByKey(key)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	response.SuccessMsg(ctx, http.StatusOK, "Skill deleted")
}
