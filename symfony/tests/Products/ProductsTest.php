<?php

namespace Tests\Products;

use App\Domain\Adapters\ProductsImpl;
use App\Domain\Dto\QueryParams;
use Symfony\Bundle\FrameworkBundle\Test\KernelTestCase;

class ProductsTest extends KernelTestCase{
    public function testSomething(){
        self::bootKernel();        
        $container = static::getContainer();
        $productsImpl = $container->get(ProductsImpl::class);
        $params = new QueryParams(null, null,20,0);
        $products = $productsImpl->GetProducts($params);
        foreach($products as $product){
            //useCase 1: boots category have a 30% discount
            if($product->category === "boots"){
                $this->assertEquals("30%", $product->price->discountPercentage);
            }
            //useCase 2: The product with sku = 000003 has a 15% discount            
            if($product->category !== "boots" && $product->sku === "000003"){
                $this->assertEquals("15%", $product->price->discountPercentage);
            }
            //useCase 3: When multiple discounts collide, the bigger discount must be applied
            if($product->category === "boots" && $product->sku === "000003"){
                $this->assertEquals("30%", $product->price->discountPercentage);
            }
            //useCase 4: price.currency is always EUR
            $this->assertEquals("EUR", $product->price->currency);
            //useCase 5: When a product does not have a discount, price.final and price.original should be the same number and discount_percentage should be null
            if($product->category !== "boots" && $product->sku !== "000003"){
                $this->assertNull($product->price->discountPercentage);
                $this->assertEquals(floatval($product->price->original), $product->price->final);
            }
            //useCase 6: When a product has a discount price.original is the original price,
            //price.final is the amount with the discount applied and discount_percentage
            //represents the applied discount with the % sign.
            if(!is_null($product->price->discountPercentage)){
                $discountFloat = floatval(substr($product->price->discountPercentage,0,-1));
                $discountDec = $discountFloat/100;
                $mustDiscount = floatval($product->price->original) * $discountDec;
                $mustFinal = floatval($product->price->original) - $mustDiscount;                
                $this->assertEquals($mustFinal,$product->price->final);
            }
        }
    }
}