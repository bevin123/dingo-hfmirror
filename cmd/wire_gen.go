// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"dingo-hfmirror/internal/dao"
	"dingo-hfmirror/internal/handler"
	"dingo-hfmirror/internal/router"
	"dingo-hfmirror/internal/service"
	"dingo-hfmirror/pkg/config"
	"dingo-hfmirror/pkg/server"
)

// Injectors from wire.go:

func wireApp(configConfig *config.Config) (*App, func(), error) {
	echo := server.NewEngine()
	fileDao := dao.NewFileDao()
	fileService := service.NewFileService(fileDao)
	fileHandler := handler.NewFileHandler(fileService)
	metaDao := dao.NewMetaDao(fileDao)
	metaService := service.NewMetaService(fileDao, metaDao)
	metaHandler := handler.NewMetaHandler(metaService)
	httpRouter := router.NewHttpRouter(echo, fileHandler, metaHandler)
	serverServer := server.NewServer(configConfig, echo, httpRouter)
	app := newApp(serverServer)
	return app, func() {
	}, nil
}
