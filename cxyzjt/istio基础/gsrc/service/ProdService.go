package service

import (
	"context"
	"mypro/gsrc/pbfiles"
)

type ProdService struct {

}
func NewProdService() *ProdService {
	return &ProdService{}
}
func(this *ProdService) GetProd(ctx context.Context, req *pbfiles.ProdRequest) (rsp *pbfiles.ProdResponse,err error){
	model:=&pbfiles.ProdModel{Id:req.ProdId,Name:"测试商品"}
	rsp=&pbfiles.ProdResponse{
		Result:model,
	}
	return
}