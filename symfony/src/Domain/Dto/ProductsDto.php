<?php

namespace App\Domain\Dto;

class QueryParams {
    
    public ?string $category;
    public ?int $priceLessThan;

    public function __construct(
        ?string $category,
        ?int $priceLessThan
    )
    {
        $this->category = $category;
        $this->priceLessThan = $priceLessThan;
    }
}

class PriceResp {
    public int $original;
    public float $final;
    public ?string $discountPercentage;
    public string $currency;

    public function __construct(
        string $original,
        string $final,
        ?string $discountPercentage,
        string $currency
    )
    {
        $this->original = $original;
        $this->final = $final;
        $this->discountPercentage = $discountPercentage;
        $this->currency = $currency;
    }
}

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