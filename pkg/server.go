package pkg

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	databaseBehavior "github.com/motojouya/ddd_go/pkg/database/behavior"
	localBehavior "github.com/motojouya/ddd_go/pkg/local/behavior"
	localCore "github.com/motojouya/ddd_go/pkg/local/core"
)

type ServeCmd struct{}

func (srv *ServeCmd) Run() error {
	serverConf, err := localBehavior.GetEnv[localCore.Server]()
	if err != nil {
		fmt.Println("failed to get server config:", err)
	}

	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	Route(e)

	// db connectionを取得してserver停止時にclose
	dbGetter := databaseBehavior.NewDatabaseGet()
	db, err := dbGetter.GetDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	// SIGINT/SIGTERMで停止する
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(serverConf.GetEchoPort()); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
		fmt.Println("start server!")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}

func Route(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	e.GET("/heartbeat", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// routeを追加していく
}
