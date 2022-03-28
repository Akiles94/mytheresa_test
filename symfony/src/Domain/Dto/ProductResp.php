<?php

namespace App\Domain\Dto;

class ProductResp {
    public string $sku;
    public string $name;
    public string $category;
    public PriceResp $price;

    public function __construct(
        string $sku,
        string $name,
        string $category,
        PriceResp $price
    )
    {
        $this->sku = $sku;
        $this->name = $name;
        $this->category = $category;
        $this->price = $price;
    }
}