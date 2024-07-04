package entity

type Vote struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	NumCandidat int    `json:"num_candidat"`
	NumBulletin int    `json:"num_bulletin"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
