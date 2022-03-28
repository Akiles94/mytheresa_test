<?php

namespace App\Domain\Dto;


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
