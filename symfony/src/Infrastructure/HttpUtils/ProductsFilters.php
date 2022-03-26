<?php

namespace HttpUtils;

use App\Domain\Dto\QueryParams;
use Models\Product;
use Symfony\Component\HttpFoundation\Request;

class HttpUtils {

    public static function GetProductsParams(Request $request):QueryParams{
        $response = null;
        $category = $request->get('category');
        $priceLessThan = $request->get('priceLessThan');        
        if(!is_null($category) || !is_null($priceLessThan)){
            $response = new QueryParams(
                category: $category,
                priceLessThan: $priceLessThan
            );
        }else{
            $response = new QueryParams(
                category: null,
                priceLessThan: null
            );
        }
        return $response;
    }

    public static function GetProductsFiltered($products, QueryParams $params){
        $response = [];
        foreach($products as $product){            
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
            if($catFlag === true && $priceFlag === true){
                array_push($response, $product);
            }
        }
        return $response;
    }    
}