package repo_impl

import (
	"zjx/project/struct/infra/repo"
	"zjx/project/struct/infra/repo/mysql"
	"zjx/project/struct/infra/repo/redis"
	"zjx/project/struct/infra/rpc"
)

type editRepoImpl struct {
}

func NewEditRepoImpl() repo.EditRepo {
	return &editRepoImpl{}
}

func (p *editRepoImpl) Update() {
	rpc.CallDownStream()
	mysql.UpdateRecord()
	redis.UpdateRecord()
}
