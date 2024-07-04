package repository

import "github.com/ptdrpg/efidy/entity"

func (r *Repository) FindAllvote() ([]entity.Vote, error) {
	var candidats []entity.Vote
	if err := r.DB.Model(&entity.Candidat{}).Order("created_at").Find(&candidats).Error; err != nil {
		return []entity.Vote{}, err
	}
	return candidats, nil
}

func (r *Repository) FindVoteByNumB(id int) (entity.Vote, error) {
	var candidat entity.Vote
	if err := r.DB.Where("id = ?", id).Find(&candidat).Error; err != nil {
		return entity.Vote{}, err
	}
	return candidat, nil
}

func (r *Repository) CreateVote(vote *entity.Vote) error {
	if err := r.DB.Create(vote).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateVote(vote *entity.Vote) error {
	if err := r.DB.Model(vote).Updates(&vote).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteVote(id int) error {
	var vote entity.Vote
	if err := r.DB.Where("id = ?", id).Delete(&vote).Error; err != nil {
		return err
	}

	return nil
}
