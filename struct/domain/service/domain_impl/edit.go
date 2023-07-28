package domain_impl

import (
	"zjx/project/struct/domain/service"
	"zjx/project/struct/infra/repo/service/repo_impl"
)

type editDomainServiceImpl struct {
}

func (p *editDomainServiceImpl) Edit(request string) error {
	impl := repo_impl.NewEditRepoImpl()
	impl.Update()
	return nil
}

func NewEditDomainServiceImpl() service.EditDomainService {
	return &editDomainServiceImpl{}
}
