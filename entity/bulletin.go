package entity

type Bulletin struct {
	ID          uint   `gorm:"primary_key" json:"uid"`
	Operateur   string `json:"operateur"`
	NumBulletin int    `json:"num_bulletin"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
