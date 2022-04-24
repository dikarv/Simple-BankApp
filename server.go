package main

import (
	"enigmacamp.com/bank/config"
	"enigmacamp.com/bank/delivery/api"
	"enigmacamp.com/bank/delivery/tokenauth"
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
	p.routerEngine.Use(tokenauth.AuthTokenMiddleware())
	customerApiGroup := p.routerEngine.Group("/bank")
	api.NewCustomerApi(customerApiGroup, p.cfg.UseCaseManager.LoginUseCase(), p.cfg.UseCaseManager.TransferUseCase(), p.cfg.UseCaseManager.LogoutUseCase())
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
