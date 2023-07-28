package service

type TransmitDomainService interface {
	Transmit(request string) string
}

type EditDomainService interface {
	Edit(request string) error
}
