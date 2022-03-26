<?php

namespace App\Domain\Ports;

use App\Domain\Dto\QueryParams;

interface IProducts {
    public function GetProducts(QueryParams $params);
}