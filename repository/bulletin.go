package repository

import "github.com/ptdrpg/efidy/entity"

func (r *Repository) FindAllBulletin() ([]entity.Bulletin, error) {
	var bulletins []entity.Bulletin
	if err := r.DB.Model(&entity.Bulletin{}).Find(&bulletins).Error; err != nil {
		return []entity.Bulletin{}, err
	}
	return bulletins, nil
}

func (r *Repository) FindBulletintById(id int) (entity.Bulletin, error) {
	var bulletin entity.Bulletin
	if err := r.DB.Where("id = ?", id).Find(&bulletin).Error; err != nil {
		return entity.Bulletin{}, err
	}
	return bulletin, nil
}

func(r *Repository) FindBulletinByNum(num int) (entity.Bulletin,error) {
	var bulletin entity.Bulletin
	if err := r.DB.Where("num_bulletin = ?", num).Find(&bulletin); err != nil {
		return entity.Bulletin{}, err.Error
	}

	return bulletin, nil
}

func (r *Repository) CreateBulletin(bulletin *entity.Bulletin) error {
	if err := r.DB.Create(bulletin).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateBulletin(bulletin *entity.Bulletin) error {
	if err := r.DB.Model(bulletin).Updates(&bulletin).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteBulletin(id int) error {
	var candidat entity.Bulletin
	if err := r.DB.Where("id = ?", id).Delete(&candidat).Error; err != nil {
		return err
	}

	return nil
}
