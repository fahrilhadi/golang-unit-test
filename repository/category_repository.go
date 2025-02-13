package repository

import "github.com/fahrilhadi/golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}