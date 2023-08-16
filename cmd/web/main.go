package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/niteshchandra7/url_shortner/pkg/config"
	"github.com/niteshchandra7/url_shortner/pkg/drivers"
	"github.com/niteshchandra7/url_shortner/pkg/handlers"
	"github.com/niteshchandra7/url_shortner/pkg/renders"
	"github.com/niteshchandra7/url_shortner/pkg/repository/dbrepo"
)

var appConfig *config.AppConfig
var session *scs.SessionManager

func main() {
	config.LoadEnvironment()
	appConfig = config.New(os.Getenv("PORT"))
	appConfig.InProduction = false

	session = scs.New()
	session.Lifetime = 10 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	appConfig.Session = session

	setupApplication()
	defer appConfig.DB.SQL.Close()
	log.Println("Connected to database!!!")

	server := &http.Server{
		Addr:    appConfig.Addr,
		Handler: GetRoutes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

func setupApplication() {
	db, err := drivers.ConnectSQL(os.Getenv("DSN"))
	if err != nil {
		log.Panicln(err)
	}
	appConfig.DB = db
	renders.SetNewRepo(renders.GetNewRepo(appConfig))
	renders.CreateTemplateCache()
	handlers.SetNewRepo(handlers.GetNewRepo(appConfig, dbrepo.NewPostgresRepo(appConfig)))
}
