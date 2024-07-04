package controller

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/efidy/entity"
)

type BulletinOutput struct {
	ID          uint          `json:"uid"`
	Operateur   string        `json:"operateur"`
	NumBulletin int           `json:"num_bulletin"`
	Vote        []entity.Vote `json:"vote"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

func (c *Controller) CreateOutput(numBulletin entity.Bulletin, wg *sync.WaitGroup, temp chan BulletinOutput) {
	defer wg.Done()
	var bulletin BulletinOutput
	allVotes, _ := c.R.FindAllvote()
	for i := 0; i < len(allVotes); i++ {
		if allVotes[i].NumBulletin == numBulletin.NumBulletin {
			bulletin.Vote = append(bulletin.Vote, allVotes[i])
		}
	}

	bulletin.ID = numBulletin.ID
	bulletin.CreatedAt = numBulletin.CreatedAt
	bulletin.UpdatedAt = numBulletin.UpdatedAt
	bulletin.NumBulletin = numBulletin.NumBulletin
	bulletin.Operateur = numBulletin.Operateur

	temp <- bulletin
}

func (c *Controller) FindAllBulletin(ctx *gin.Context) {
	var bulletins []BulletinOutput
	var wg sync.WaitGroup
	chanel := make(chan BulletinOutput)
	votes, _ := c.R.FindAllBulletin()

	for i := 0; i < len(votes); i++ {
		wg.Add(1)
		go c.CreateOutput(votes[i], &wg, chanel)
	}

	go func() {
		wg.Wait()
		close(chanel)
	}()

	for i := range chanel {
		bulletins = append(bulletins, i)
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, bulletins)
}

type BulletinInput struct {
	Operateur   string `json:"operateur"`
	NumBulletin int    `json:"num_bulletin"`
	Vote        []int  `json:"vote"`
}

func (c *Controller) SaveVote(ctx *gin.Context, numCandidat int, bId int, wg *sync.WaitGroup, temp chan entity.Vote) {
	defer wg.Done()
	var vote entity.Vote
	vote.CreatedAt = time.Now().String()
	vote.NumBulletin = bId
	vote.NumCandidat = numCandidat
	vote.UpdatedAt = time.Now().String()

	c.R.CreateVote(&vote)
	candidat, _:= c.R.FindCandidatBynum(numCandidat)

	candidat.VoteNumber = candidat.VoteNumber + 1
	update := c.R.UpdateCandidat(&candidat)
	if update != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": update.Error(),
		})
		return
	}

	temp <- vote
}

func removeDuplicates(number []int) []int {
	unique := make(map[int]bool)
	for _, nombre := range number {
			if _, exists := unique[nombre]; !exists {
					unique[nombre] = true
			}
	}
	result := make([]int, 0, len(unique))
	for nombre := range unique {
			result = append(result, nombre)
	}

	return result
}

func (c *Controller) SaveBulletin(ctx *gin.Context) {
	var input BulletinInput
	var bulletin entity.Bulletin
	var output BulletinOutput
	if binding := ctx.ShouldBindJSON(&input); binding != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": binding.Error(),
		})
		return
	}

	checkBulletin, checkerror := c.R.FindBulletinByNum(input.NumBulletin)
	if checkerror != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": checkerror.Error(),
		})
		return
	}

	if checkBulletin.ID == uint(input.NumBulletin) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "bulletin already exist",
		})
		return
	}

	var wg sync.WaitGroup
	chanel := make(chan entity.Vote)

	bulletin.CreatedAt = time.Now().String()
	bulletin.NumBulletin = input.NumBulletin
	bulletin.Operateur = input.Operateur
	bulletin.UpdatedAt = time.Now().String()

	if err := c.R.CreateBulletin(&bulletin); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	newVote := removeDuplicates(input.Vote)

	for i:= 0; i < len(newVote); i++ {
		wg.Add(1)
		go c.SaveVote(ctx, input.Vote[i], int(input.NumBulletin), &wg, chanel)
	}

	go func() {
		wg.Wait()
		close(chanel)
	}()
		
	output.ID = bulletin.ID
	output.UpdatedAt = bulletin.UpdatedAt
	output.CreatedAt = bulletin.CreatedAt
	output.Operateur = bulletin.Operateur

	for i := range chanel {
		output.Vote = append(output.Vote, i)
	}
	
	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, output)
}

func (c *Controller) FindBulletinByNum(ctx *gin.Context) {
	getId := ctx.Param("id")
	num, _ := strconv.Atoi(getId)
	bulletin, err := c.R.FindBulletinByNum(num)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Header("content-Type", "application/json")
	ctx.JSON(http.StatusOK, bulletin)
}
