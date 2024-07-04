package entity

type Candidat struct {
	ID         uint   `gorm:"primary_key" json:"uid"`
	Num        int    `json:"num"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Gender     string `json:"gender"`
	VoteNumber int    `json:"vote_number"`
}
