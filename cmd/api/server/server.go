package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/OvictorVieira/transact.ease/internal/config"
	"github.com/OvictorVieira/transact.ease/internal/constants"
	"github.com/OvictorVieira/transact.ease/internal/routes"
	AccountRoutes "github.com/OvictorVieira/transact.ease/internal/routes/accounts"
	TransactionRoutes "github.com/OvictorVieira/transact.ease/internal/routes/transactions"
	"github.com/OvictorVieira/transact.ease/pkg/logger"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	HttpServer *http.Server
}

func NewApp() (*App, error) {
	conn, err := config.SetupPostgresConnection()

	if err != nil {
		return nil, err
	}

	// setup router
	router := routes.SetupRouter()

	// API Routes
	api := router.Group("api")
	AccountRoutes.NewAccountsRoute(api, conn).Routes()
	TransactionRoutes.NewTransactionsRoute(api, conn).Routes()

	// setup http server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.AppConfig.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &App{
		HttpServer: server,
	}, nil
}

func (a *App) Run() (err error) {
	// Gracefull Shutdown
	go func() {
		logger.InfoF("success to listen and serve on :%d", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer}, config.AppConfig.Port)
		if err := a.HttpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// make blocking channel and waiting for a signal
	<-quit
	logger.Info("shutdown server ...", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.HttpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error when shutdown server: %v", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	logger.Info("timeout of 5 seconds.", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	logger.Info("server exiting", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryServer})
	return
}
