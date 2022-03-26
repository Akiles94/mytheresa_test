package ports

import (
	"github.com/Akiles94/mytheresa-test/domain/dto"
	"github.com/Akiles94/mytheresa-test/domain/models"
)

type IRepository interface {
	Init(string) error
	GetProducts(dto.QueryParams) (*[]models.Product, error)
}
