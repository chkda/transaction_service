package service

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/chkda/transaction_service/internal/app"
	"github.com/chkda/transaction_service/internal/interfaces/healthcheck"
	"github.com/chkda/transaction_service/internal/interfaces/sum"
	"github.com/chkda/transaction_service/internal/interfaces/transactions/create"
	"github.com/chkda/transaction_service/internal/interfaces/transactions/read"
	typ "github.com/chkda/transaction_service/internal/interfaces/types"
	badgercache "github.com/chkda/transaction_service/pkg/datastores/cache/badger"
	rediscache "github.com/chkda/transaction_service/pkg/datastores/cache/redis"
	"github.com/chkda/transaction_service/pkg/datastores/database/mysqlstore"
	"github.com/chkda/transaction_service/pkg/metrics/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	HTTPPort         string             `json:"http_port"`
	RedisConfig      *rediscache.Config `json:"redis"`
	MySQLConfig      *mysqlstore.Config `json:"mysql"`
	PrometheusConfig *prometheus.Config `json:"prometheus"`
}

func Start(configFile string) {
	file, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	config := &Config{}
	err = json.Unmarshal(fileBytes, config)
	if err != nil {
		panic(err)
	}

	redisClient := rediscache.New(config.RedisConfig)
	badgerClient, err := badgercache.New()
	if err != nil {
		panic(err)
	}
	mysqlClient, err := mysqlstore.New(config.MySQLConfig)
	if err != nil {
		panic(err)
	}

	appHandler := app.New(badgerClient, redisClient, mysqlClient)
	healthcheckController := healthcheck.New()
	createTransactionController := create.New(appHandler)
	readTransactionController := read.New(appHandler)
	sumController := sum.New(appHandler)
	typesController := typ.New(appHandler)
	log.Println("[INFO]:Starting server")
	serv := echo.New()
	serv.Use(middleware.Recover())
	serv.GET(healthcheckController.GetRoute(), healthcheckController.Handler)
	serv.PUT(createTransactionController.GetRoute(), createTransactionController.Handler)
	serv.GET(readTransactionController.GetRoute(), readTransactionController.Handler)
	serv.GET(sumController.GetRoute(), sumController.Handler)
	serv.GET(typesController.GetRoute(), typesController.Handler)
	serv.Logger.Fatal(serv.Start(":" + config.HTTPPort))
}
