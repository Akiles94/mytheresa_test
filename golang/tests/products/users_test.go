package products

import (
	"strconv"
	"testing"

	"github.com/Akiles94/mytheresa-test/domain/adapters"
	"github.com/Akiles94/mytheresa-test/domain/dto"
	"github.com/Akiles94/mytheresa-test/domain/ports"
	datasetimpl "github.com/Akiles94/mytheresa-test/infrastructure/repository/adapters/dataset_impl"
	"github.com/stretchr/testify/assert"
)

type ProductsTest struct {
	ProductsImpl ports.IProducts
}

func TestGetProducts(t *testing.T) {
	//Normally for the repo we could have mocked data repo, but for this case
	//is not necessary becouse we are not doing a DB connection.
	repo := &datasetimpl.DatasetRepo{}
	repo.Init("./products.json")

	productsImpl := adapters.ProductsImpl{
		Repo: repo,
	}

	queryParams := dto.QueryParams{}

	products, _ := productsImpl.GetProducts(queryParams)
	for _, item := range *products {
		//useCase 1: boots category have a 30% discount
		if item.Category == "boots" {
			assert.Equal(t, "30%", *item.Price.DiscountPercentage)
		}
		//useCase 2: The product with sku = 000003 has a 15% discount
		if item.Category != "boots" && item.Sku == "000003" {
			assert.Equal(t, "15%", *item.Price.DiscountPercentage)
		}
		//useCase 3: When multiple discounts collide, the bigger discount must be applied
		if item.Category == "boots" && item.Sku == "000003" {
			assert.Equal(t, "30%", *item.Price.DiscountPercentage)
		}
		//useCase 4: price.currency is always EUR
		assert.Equal(t, "EUR", item.Price.Currency)

		//useCase 5: When a product does not have a discount, price.final and price.original should be the same number and discount_percentage should be null
		if item.Category != "boots" && item.Sku != "000003" {
			assert.Nil(t, item.Price.DiscountPercentage)
			assert.Equal(t, float32(item.Price.Original), item.Price.Final)
		}

		//useCase 6: When a product has a discount price.original is the original price,
		//price.final is the amount with the discount applied and discount_percentage
		//represents the applied discount with the % sign.
		if item.Price.DiscountPercentage != nil {
			stringDiscount := *item.Price.DiscountPercentage
			discountPercentage, _ := strconv.Atoi(stringDiscount[:len(stringDiscount)-1])
			floatPercentage := float32(discountPercentage)
			floatPercentage = floatPercentage / 100
			mustDiscount := float32(item.Price.Original) * floatPercentage
			mustFinal := float32(item.Price.Original) - mustDiscount
			assert.Equal(t, mustFinal, item.Price.Final)
		}
	}
}
