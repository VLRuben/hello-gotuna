package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alcalbg/gotdd/app"
	"github.com/alcalbg/gotdd/models"
	"github.com/alcalbg/gotdd/session"
	"github.com/alcalbg/gotdd/static"
	"github.com/alcalbg/gotdd/util"
	"github.com/gorilla/sessions"
)

func main() {

	port := ":8888"

	app := app.NewApp(util.Options{
		Logger:         log.New(os.Stdout, "", 0),
		FS:             static.EmbededStatic,
		Session:        session.NewSession(sessions.NewCookieStore([]byte(os.Getenv("APP_KEY")))),
		UserRepository: models.NewInMemoryUserRepository(),
	})

	fmt.Printf("starting server at http://localhost%s \n", port)

	if err := http.ListenAndServe(port, app); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
