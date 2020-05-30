package main

import "github.com/jiandahao/DesignPatterns/Structural_patterns/proxy"

func main() {
	service := proxy.NewService()
	serviceProxy := proxy.NewProxy(service)

	serviceProxy.ListResource()
}
