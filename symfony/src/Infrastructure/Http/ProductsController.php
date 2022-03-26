<?php

namespace App\Controller;

use App\Domain\Ports\IProducts;
use Exception;
use HttpUtils\HttpUtils;
use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\Routing\Annotation\Route;
use Psr\Log\LoggerInterface;

class ProductsController extends AbstractController{

    private LoggerInterface $logger;
    private IProducts $productsImpl;

    public function __construct(
        LoggerInterface $logger,
        IProducts $productsImpl
    )
    {
        $this->logger = $logger;
        $this->productsImpl = $productsImpl;
    }

    /**
     * @Route("/products/list", name="products_list")
     * 
     */
    public function list(Request $request):JsonResponse{
        try{
            $queryParams = HttpUtils::GetProductsParams($request);

            $result = $this->productsImpl->GetProducts($queryParams);

            $response = new JsonResponse();
            $response->setData($result);
    
            return $response;
        }catch(Exception $error){
            $this->logger->error(json_encode($error));
            return json_encode($error);
        }        
    }

}