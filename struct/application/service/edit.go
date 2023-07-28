package service

import (
	"zjx/project/struct/domain/service/domain_impl"
)

func RpcEdit(request string) error {
	domainImpl := domain_impl.NewEditDomainServiceImpl()
	return domainImpl.Edit(request)
}
