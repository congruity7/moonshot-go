package main

import (
	"context"
	"database/sql"
	"net/http"
	"sync"
	"time"

	"github.com/congruity7/moonshot-go/pkg/api"
	"github.com/congruity7/moonshot-go/pkg/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func Setup() {
	dbInstance, err := sql.Open("mysql", "fred:asdfzxcv@cloudsql(phonic-ceremony-317217:asia-northeast1:divvy-db)/moonshot")

	if err != nil {
		logrus.Fatal("connecting to database : ", err)
	}
	defer dbInstance.Close()

	var wg sync.WaitGroup
	var stopChan <-chan struct{}

	dbInstance.SetConnMaxLifetime(4 * time.Minute)
	dbInstance.SetConnMaxIdleTime(4 * time.Minute)
	dbInstance.SetMaxIdleConns(5)
	dbInstance.SetMaxOpenConns(5)

	ds := service.NewDatabaseService(dbInstance)
	rs := service.NewRedisService(nil)
	logger := &logrus.Logger{}

	wg.Add(1)

	go StartAPI(&wg, ds, rs, logger, stopChan)

	wg.Wait()
}

func StartAPI(wg *sync.WaitGroup, ds *service.DatabaseService, rs *service.RedisService, logger *logrus.Logger, stopChan <-chan struct{}) {
	defer wg.Done()

	router := httprouter.New()

	ac := api.NewAPIContext(ds, rs, logger)
	router.GET("/moonshot/v1/users/:user_id/", ac.GetUserByID)
	router.GET("/moonshot/v1/users", ac.GetUsers)
	router.POST("/moonshot/v1/user", ac.CreateUser)
	router.PUT("/moonshot/v1/user", ac.UpdateUser)
	router.DELETE("/moonshot/v1/user/:user_id/", ac.DeleteUserByID)

	n := negroni.New(negroni.NewRecovery())
	n.UseHandler(router)

	api := &http.Server{Addr: "127.0.0.1:8000", Handler: n}

	go func() {
		if err := api.ListenAndServe(); err != nil {
			logger.Error("starting api server", err)
		}
	}()

	<-stopChan

	logger.Info("shutting down api server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	api.Shutdown(ctx)

}

func main() {
	Setup()
}
