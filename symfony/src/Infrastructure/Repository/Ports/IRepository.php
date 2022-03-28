<?php

namespace App\Infrastructure\Repository\Ports;

use App\Domain\Dto\QueryParams;
use App\Domain\Models\Product;

interface IRepository {
    /**
    * @return Product[]
    */
    public function GetProducts(QueryParams $params);
}