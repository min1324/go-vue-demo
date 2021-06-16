package repository

import (
	"demo/common"
	"demo/model"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{DB: common.GetDB()}
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}
	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) Update(cg model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&cg).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &cg, nil
}

func (c CategoryRepository) SelectById(id int) (*model.Category, error) {
	category := model.Category{}
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (c CategoryRepository) DeleteById(id int) error {
	if err := c.DB.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
