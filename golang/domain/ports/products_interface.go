package ports

import (
	"github.com/Akiles94/mytheresa-test/domain/dto"
	"github.com/Akiles94/mytheresa-test/infrastructure/httputils"
)

type IProducts interface {
	GetProducts(dto.QueryParams) (*[]dto.ProductResp, *httputils.Exception)
}
