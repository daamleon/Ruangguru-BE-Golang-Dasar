package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	// "fmt"
	"errors"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	filebasedDb *filebased.Data
}

func NewCategoryRepo(filebasedDb *filebased.Data) *categoryRepository {
	return &categoryRepository{filebasedDb}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	c.filebasedDb.StoreCategory(*Category)
	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	existingCategory, err := c.filebasedDb.GetCategoryByID(id)
	if err != nil {
		return errors.New("category not found")
	}

	existingCategory.Name = category.Name
	err = c.filebasedDb.UpdateCategory(id, *existingCategory) // Tambahkan `id` sebagai parameter pertama
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) Delete(id int) error {
	return c.filebasedDb.DeleteCategory(id)
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	return c.filebasedDb.GetCategoryByID(id)
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	return c.filebasedDb.GetCategories() // Gunakan `GetCategories` yang ada di filebased
}