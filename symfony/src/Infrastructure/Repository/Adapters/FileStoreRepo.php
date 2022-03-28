<?php

namespace App\Infrastructure\Repository\Adapters;

use App\Infrastructure\Repository\Ports\IRepository;
use App\Domain\Dto\QueryParams;
use App\Infrastructure\HttpUtils\ProductsFilters;
use App\Domain\Models\Product;

class FileStoreRepo implements IRepository{
    
    private string $filePath;

    public function __construct(string $filePath)
    {   
        $this->filePath = $filePath;
    }

    public function GetProducts(QueryParams $params){
        //Get base path
        $basePath = dirname(__DIR__,4);
        $strJsonFileContents = file_get_contents($basePath.$this->filePath);
        // Convert to array 
        $array = json_decode($strJsonFileContents);
        $products = [];
        foreach($array as $item){
            $product = new Product(
                sku: $item->sku,
                name: $item->name,
                category: $item->category,
                price: $item->price
            );
            array_push($products, $product);
        }

        $products = ProductsFilters::GetProductsFiltered($products, $params);

        return $products;
    }

}