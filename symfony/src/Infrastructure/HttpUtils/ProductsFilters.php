<?php

namespace App\Infrastructure\HttpUtils;

use App\Domain\Dto\QueryParams;
use Symfony\Component\HttpFoundation\Request;

class ProductsFilters {

    public static function GetProductsParams(Request $request):QueryParams{
        $response = null;
        $category = $request->get('category');
        $priceLessThan = $request->get('priceLessThan');        
        $limit = $request->get('limit',20);        
        $offset = $request->get('offset',0);        
        $response = new QueryParams(
            category: $category,
            priceLessThan: $priceLessThan,
            limit: $limit,
            offset: $offset,
        );
        return $response;
    }

    public static function GetProductsFiltered($products, QueryParams $params){
        $response = [];
        $count = 0;
        foreach($products as $product){            
            $count++;
            $catFlag = false;
            $priceFlag = false;
            if(!is_null($params->category)){
                if($params->category == $product->category){
                    $catFlag = true;
                }                
            }else{
                $catFlag = true;
            }
            if(!is_null($params->priceLessThan)){
                if($product->price < $params->priceLessThan){
                    $priceFlag = true;
                }                
            }else{
                $priceFlag = true;
            }
            if($catFlag === true && $priceFlag === true && $count <= $params->limit && $count > $params->offset){
                array_push($response, $product);
            }
        }
        return $response;
    }    
}