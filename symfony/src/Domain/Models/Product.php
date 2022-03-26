<?php

namespace Models;

class Product{
    
    public string $sku;
    public string $name;
    public string $category;
    public float $price;

    public function __construct(
        string $sku,
        string $name,
        string $category,
        float $price
    )
    {
        $this->sku = $sku;
        $this->name = $name;
        $this->category = $category;
        $this->price = $price;
    }

}