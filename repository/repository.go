package repository

import (
	db "time-wise/database"

	"gorm.io/gorm"
)

type Repository struct {
	Table string
}

func (r *Repository) Super(Table string) {
	r.Table = Table
}

func (r *Repository) Update(condition string, values map[string]interface{}, T any) {
	result := db.GetDatabase().Table(r.Table).Where(condition, values).Updates(T)
	if result.Error != nil {
		T = nil
	}
}

func (r *Repository) Create(T any) {
	result := db.GetDatabase().Table(r.Table).Create(T)

	if result.Error != nil {
		T = nil
	}
}

func (r *Repository) FindOne(condition string, values map[string]interface{}, T any) {
	result := db.GetDatabase().Table(r.Table).First(T, condition, values)

	if result.Error != nil {
		T = nil
	}
}

func (r *Repository) Raw() *gorm.DB {
	return db.GetDatabase().Table(r.Table)
}

func (r *Repository) FindAll(condition string, values map[string]interface{}, T any) {
	if condition == "" {
		db.GetDatabase().Table(r.Table).Find(&T)
		return
	}

	result := db.GetDatabase().Table(r.Table).Find(&T, condition, values)
	if result.Error != nil {
		T = nil
	}
}
