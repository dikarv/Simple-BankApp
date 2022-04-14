package main

import (
	"enigmacamp.com/bank/config"
	"enigmacamp.com/bank/delivery/api"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine *gin.Engine
	cfg          config.Config
}

func (p *appServer) initHandlers() {
	p.v1()
}

func (p *appServer) v1() {
	customerApiGroup := p.routerEngine.Group("/users")
	api.NewCustomerApi(customerApiGroup, p.cfg.UseCaseManager.LoginUseCase(), p.cfg.UseCaseManager.TransferUseCase())
}

func (p *appServer) Run() {
	p.initHandlers()
	err := p.routerEngine.Run(p.cfg.ApiConfig.Url)
	if err != nil {
		panic(err)
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig(".", "config")
	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
