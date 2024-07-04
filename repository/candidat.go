package repository

import "github.com/ptdrpg/efidy/entity"

func (r *Repository) FindAllCandidat() ([]entity.Candidat, error) {
	var candidats []entity.Candidat
	if err := r.DB.Model(&entity.Candidat{}).Order("created_at").Find(&candidats).Error; err != nil {
		return []entity.Candidat{}, err
	}
	return candidats, nil
}

func (r *Repository) FindCandidatById(id int) (entity.Candidat, error) {
	var candidat entity.Candidat
	if err := r.DB.Where("id = ?", id).Find(&candidat).Error; err != nil {
		return entity.Candidat{}, err
	}
	return candidat, nil
}

func (r *Repository) FindCandidatBynum(id int) (entity.Candidat, error) {
	var candidat entity.Candidat
	if err := r.DB.Where("num = ?", id).Find(&candidat).Error; err != nil {
		return entity.Candidat{}, err
	}
	return candidat, nil
}

func (r *Repository) CreateCandidat(candidat *entity.Candidat) error {
	if err := r.DB.Create(candidat).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateCandidat(candidat *entity.Candidat) error {
	if err := r.DB.Where("id = ?", candidat.ID).Updates(&candidat).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteCandidat(id int) error {
	var candidat entity.Candidat
	if err := r.DB.Where("id = ?", id).Delete(&candidat).Error; err != nil {
		return err
	}

	return nil
}
