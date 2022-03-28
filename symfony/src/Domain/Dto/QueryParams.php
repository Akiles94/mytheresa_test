<?php

namespace App\Domain\Dto;

class QueryParams {
    
    public ?string $category;
    public ?int $priceLessThan;
    public int $limit;
    public int $offset;

    public function __construct(
        ?string $category,
        ?int $priceLessThan,
        int $limit,
        int $offset,
    )
    {
        $this->category = $category;
        $this->priceLessThan = $priceLessThan;
        $this->limit = $limit;
        $this->offset = $offset;
    }
}