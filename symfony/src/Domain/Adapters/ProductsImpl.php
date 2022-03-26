<?php

namespace App\Domain\Adapters;

use App\Domain\Ports\IProducts;
use App\Infrastructure\Repository\Ports\IRepository;
use App\Domain\Dto\PriceResp;
use App\Domain\Dto\ProductResp;
use App\Domain\Dto\QueryParams;

class ProductsImpl implements IProducts{

    private IRepository $repo;

    public function __construct(IRepository $repo)
    {
        $this->repo = $repo;
    }

    public function GetProducts(QueryParams $params)
    {
        $response = [];        
        $repoRes = $this->repo->GetProducts($params);
        foreach($repoRes as $item){
            $respProduct = null;
            $original = $item->price;
            $final = floatval($item->price);
            $percentageDiscount = null;
            //Products in the boots category have a 30% discount
            if($item->category === "boots"){
                $percentageDiscount = "30%";
                $aux = floatval($item->price) * 0.3;
                $final = floatval($original) - $aux;
            }
            //Product with sku = 000003 has a 15% discount
		    //When multiple discounts colide, the bigger discount must be applied
            if($item->sku === "000003" && $item->category !== "boots"){
                $percentageDiscount = "15%";
                $aux = floatval($item->price) * 0.15;
                $final = floatval($original) - $aux;
            }
            $respProduct = new ProductResp(
                sku: $item->sku,
                name: $item->name,
                category: $item->category,
                price: new PriceResp(
                    original: $original,
                    final: $final,
                    discountPercentage: $percentageDiscount,
                    currency: "EUR"
                )
            );
            //var_dump($respProduct);
            array_push($response, $respProduct);
        }        
        return $response;
    }
}