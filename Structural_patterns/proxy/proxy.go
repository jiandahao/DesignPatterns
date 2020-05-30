package proxy

import "fmt"

type ServiceInterface interface {
	ListResource()
}

type Service struct{}

func (s *Service) ListResource() {
	fmt.Println("list resource")
}

func NewService() *Service {
	return &Service{}
}

type Proxy struct {
	Service ServiceInterface
}

func (p *Proxy) ListResource() {
	fmt.Println("Proxy: Do something before service")
	p.Service.ListResource()
	fmt.Println("Proxy: Do something after service")
}

func NewProxy(service ServiceInterface) *Proxy {
	return &Proxy{Service: service}
}

//type ServiceManager struct {
//	Service ServiceInterface
//}
//
//func NewServiceManager(service ServiceInterface) *ServiceManager {
//	return &ServiceManager{Service:service}
//}
