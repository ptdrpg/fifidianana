package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/efidy/entity"
	"github.com/ptdrpg/efidy/lib"
)

func (c *Controller) FindAllCandidat(ctx *gin.Context) {
	candidats, err := c.R.FindAllCandidat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, candidats)
}

type DataLists struct {
	Data []entity.Candidat `json:"data"`
}

type ByNum struct {
	Data entity.Candidat `json:"data"`
}

// @Summary find all men
// @Schemes
// @Description find all men candidat
// @Tags candidat
// @Accept json
// @Produce json
// @Success 200 {object} DataLists
// @Router /candidat/men [get]
func (c *Controller) FindAllMen(ctx *gin.Context) {
	var allMen []entity.Candidat
	allcandidat, _ := c.R.FindAllCandidat()
	for i := 0; i < len(allcandidat); i++ {
		if allcandidat[i].Gender == "H" || allcandidat[i].Gender == "h" {
			allMen = append(allMen, allcandidat[i])
		}
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": allMen,
	})
}

// @Summary find all women
// @Schemes
// @Description find all woman candidat
// @Tags candidat
// @Accept json
// @Produce json
// @Success 200 {object} DataLists
// @Router /candidat/woman [get]
func (c *Controller) FindAllWoman(ctx *gin.Context) {
	var allWoman []entity.Candidat
	allcandidat, _ := c.R.FindAllCandidat()
	for i := 0; i < len(allcandidat); i++ {
		if allcandidat[i].Gender == "F" || allcandidat[i].Gender == "f" {
			allWoman = append(allWoman, allcandidat[i])
		}
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": allWoman,
	})
}

// @Summary find specific candidat
// @Schemes
// @Description find some specific candidat
// @Tags candidat
// @Accept json
// @Produce json
// @Success 200 {object} ByNum
// @Router /candidat/:id [get]
func (c *Controller) FindCandidatByNum(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	candidat, err := c.R.FindCandidatBynum(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, candidat)
}

type CandidatInput struct {
	Num    int    `json:"num"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type Createresponse struct {
	ID         uint   `json:"id"`
	Num        int    `json:"num"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Gender     string `json:"gender"`
	VoteNumber int    `json:"vote_number"`
}

// @Summary create candidat
// @Schemes
// @Description create candidat
// @Tags candidat
// @Accept json
// @Produce json
// @Param body body CandidatInput true " "
// @Success 201 {object} Createresponse
// @Router /candidat [post]
func (c *Controller) CreateCandidat(ctx *gin.Context) {
	var input CandidatInput
	var candidat entity.Candidat
	binding := ctx.ShouldBindJSON(&input)
	if binding != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": binding.Error(),
		})
		return
	}

	candidat.Name = input.Name
	candidat.Num = input.Num

	if save := c.R.CreateCandidat(&candidat); save != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": save.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, candidat)
}

// @Summary update specific candidat
// @Schemes
// @Description update some specific candidat
// @Tags candidat
// @Accept json
// @Produce json
// @Success 200 {object} Createresponse
// @Router /candidat/:id [update]
func (c *Controller) UpdateCandidat(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)
	var candidat entity.Candidat
	binding := ctx.ShouldBindJSON(&candidat)
	if binding != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": binding.Error(),
		})
		return
	}
	candidat.ID = uint(id)

	if err := c.R.UpdateCandidat(&candidat); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, candidat)
}

func (c *Controller) DeleteCandidat(ctx *gin.Context) {
	getId := ctx.Param("id")
	id, _ := strconv.Atoi(getId)

	if err := c.R.DeleteCandidat(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, "candidat succefuly deleted")
}

// @Summary create candidat
// @Schemes
// @Description create candidat
// @Tags candidat
// @Accept json
// @Produce json
// @Param body body CandidatInput true " "
// @Success 201 {object} Createresponse
// @Router /candidat [post]
func (c *Controller) UploadCandidatAvatar(ctx *gin.Context) {
	candidatId := ctx.Param("id")
	id, _ := strconv.Atoi(candidatId)

	avatar, err := ctx.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	path, pathErr := lib.CreateImage(avatar, ctx)
	if pathErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": pathErr.Error(),
		})
		return
	}

	candidat, getCandidat := c.R.FindCandidatBynum(id)
	if getCandidat != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"erro": getCandidat.Error(),
		})
		return
	}

	candidat.Avatar = path
	c.R.UpdateCandidat(&candidat)

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, candidat)
}
