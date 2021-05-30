// Getir Challenge
//
//
//     Schemes: http, https
//     Host: fufu-getir-challange.herokuapp.com
//     BasePath: /
//     Version: 0.0.1
//     Contact: Furkan Aksoy<aksoyfurkan@pm.me> https://github.com/fufuceng
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"github.com/fufuceng/getir-challange/config"
	"github.com/fufuceng/getir-challange/inmem"
	"github.com/fufuceng/getir-challange/mongo"
	"github.com/fufuceng/getir-challange/server"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.Initialize(); err != nil {
		logrus.WithError(err).Error("error while initializing config")
		return
	}

	if err := mongo.Initialize(config.Get().Mongo); err != nil {
		logrus.WithError(err).Error("error while initializing mongo")
		return
	}

	if err := inmem.Initialize(); err != nil {
		logrus.WithError(err).Error("error while initializing in-memory db")
	}

	server.Initialize(config.Get().Server)
	server.Run() // blocking
}
