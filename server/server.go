package server

import (
	"errors"
	"github.com/fufuceng/getir-challange/config"
	"github.com/fufuceng/getir-challange/memrecords"
	"github.com/fufuceng/getir-challange/records"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

var server *http.Server

func InitializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./static/docs")))

	records.NewDefaultRouter().Register(mux)
	memrecords.NewDefaultRouter().Register(mux)

	return mux
}

func InitializeRoutesForTest() *http.ServeMux {
	mux := http.NewServeMux()

	records.NewDefaultRouter().Register(mux)
	memrecords.NewDefaultRouter().Register(mux)

	return mux
}

func Initialize(config config.Server) {
	server = &http.Server{
		Addr:         strings.Join([]string{config.Host, config.Port}, ":"),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      InitializeRoutes(),
	}

}

//Run responsible to run server. It's blocking operation, please use carefully
func Run() {
	logger := logrus.WithField("location", "Run")

	logger.WithFields(logrus.Fields{"addr": server.Addr}).Info("running server..")
	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logrus.Info("server closed..")
		}

		logger.WithError(err)
	}
}
