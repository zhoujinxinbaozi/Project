package service

import (
	"zjx/project/struct/domain/service/domain_impl"
)

func RpcTransmit(request string) string {
	domainImpl := domain_impl.NewTransmitDomainServiceImpl()
	return domainImpl.Transmit(request)
}
