package adapters

import (
	"github.com/Akiles94/mytheresa-test/domain/dto"
	"github.com/Akiles94/mytheresa-test/infrastructure/httputils"
	"github.com/Akiles94/mytheresa-test/infrastructure/repository/ports"
	log "github.com/sirupsen/logrus"
)

type ProductsImpl struct {
	Repo ports.IRepository
}

func (p *ProductsImpl) GetProducts(params dto.QueryParams) (*[]dto.ProductResp, *httputils.Exception) {
	response := []dto.ProductResp{}
	repoRes, err := p.Repo.GetProducts(params)
	if err != nil {
		log.Error("Error getting data from repo: ", err)
		return nil, httputils.NewException(httputils.GeneralException, err)
	}

	for _, item := range *repoRes {
		var respProduct dto.ProductResp
		var original int
		var final float32
		var percentageDiscount *string
		original = item.Price
		final = float32(item.Price)
		//Products in the boots category have a 30% discount
		if item.Category == "boots" {
			auxPercentageDiscount := "30%"
			percentageDiscount = &auxPercentageDiscount
			tmpDiscount := float32(item.Price) * 0.3
			final = float32(original) - tmpDiscount
		}
		//Product with sku = 000003 has a 15% discount
		//When multiple discounts colide, the bigger discount must be applied
		if item.Sku == "000003" && item.Category != "boots" {
			auxPercentageDiscount := "15%"
			percentageDiscount = &auxPercentageDiscount
			tmpDiscount := float32(item.Price) * 0.15
			final = float32(original) - tmpDiscount
		}
		respProduct = dto.ProductResp{
			Sku:      item.Sku,
			Name:     item.Name,
			Category: item.Category,
			Price: dto.PriceResp{
				Original:           original,
				Final:              final,
				DiscountPercentage: percentageDiscount,
				Currency:           "EUR",
			},
		}
		response = append(response, respProduct)
	}

	return &response, nil
}
