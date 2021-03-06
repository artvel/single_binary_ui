package doctmpl

import (
	"github.com/labstack/echo"
	"fmt"
	"net/http"
	"github.com/labstack/echo/middleware"
	"errors"
	"flag"
	"github.com/labstack/echo-contrib/session"
	"github.com/gorilla/sessions"
	"os"
	"os/signal"
	"io/ioutil"
	"context"
	"time"
	"encoding/gob"
	"log"
)

var (
	ReadAllFile = func(path string) ([]byte, error) {
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		return ioutil.ReadAll(f)
	}
)

type MyServer struct {
	quit chan os.Signal
}

func (ms *MyServer) Close() {
	if ms.quit != nil {
		ms.quit <- os.Interrupt
	}
}

func StartServer(e *echo.Echo, addr string, afterStart func(), onShutdown func()) (*MyServer, error) {
	e.HideBanner = true
	var err error
	ms := &MyServer{quit: make(chan os.Signal)}
	// Start server
	go func() {
		fmt.Println("starting at", addr)
		if err = e.Start(addr) /*e.StartAutoTLS(":443")*/ ; err != nil {
			fmt.Println("shutting down the server cause of: ", err)
			ms.quit <- os.Interrupt
		}
	}()
	if afterStart != nil {
		go afterStart()
	}
	signal.Notify(ms.quit, os.Interrupt)
	<-ms.quit
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	shutdownWebListener(e, &ctx)
	if onShutdown != nil {
		onShutdown()
	}
	return ms, err
}

func shutdownWebListener(e *echo.Echo, ctx *context.Context) {
	if err := e.Server.Shutdown(*ctx); err != nil {
		e.Logger.Fatal(err)
	}
	if err := e.TLSServer.Shutdown(*ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func SetupServer(beforeStart func(ec *echo.Echo)) (*MyServer, error) {
	port := flag.String("p", "58082", "Port")
	host := flag.String("h", "127.0.0.1", "Host")
	flag.Parse()
	var err error

	if err != nil {
		panic(err)
	}
	e := echo.New()
	gob.Register(map[string]interface{}{})
	sessionStore := sessions.NewCookieStore([]byte("secret_Dummy_1234"), []byte("12345678901234567890123456789012"))
	e.Use(session.Middleware(sessionStore))

	e.Pre(middleware.Secure())

	e.GET("/", func(c echo.Context) error {
		b, err := ReadAllFile("view/dev.html")
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}
		return c.HTMLBlob(http.StatusOK, b)
	})
	e.GET("/btn", func(c echo.Context) error {
		log.Println("button clicked")
		return c.NoContent(http.StatusOK)
	})

	if beforeStart == nil {
		panic(errors.New("beforeStart can't be nil!"))
	}
	beforeStart(e)
	//web.StaticSetup(system, e, "/static", "static", moreAssetDirs...)

	return StartServer(e, (*host)+":"+(*port), nil, func() {
		fmt.Println("shutting down the server")
	})
}
