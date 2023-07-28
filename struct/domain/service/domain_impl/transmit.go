package domain_impl

import (
	"zjx/project/struct/domain/service"
)

type transmitDomainServiceImpl struct {
}

func (p *transmitDomainServiceImpl) Transmit(request string) string {
	return "transmit"
}

func NewTransmitDomainServiceImpl() service.TransmitDomainService {
	return &transmitDomainServiceImpl{}
}
