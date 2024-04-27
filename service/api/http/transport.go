package http

import (
	"fmt"

	"github.com/aerogo/aero"
)

func MakeHandlers() *aero.Application {
	app := aero.New()
	app.Load()
	//routing
	app.Get("/", home)

	app.Post("/hello", hello)
	app.Get("/hello", getHello)
	app.Get("/hello/:person", helloPerson)
	app.Get("/goodmorning/:person", goodmorningPerson)

	app.Get("/images/*file", imageFile)
	app.Get("/streamhello", streamHello)
	app.Get("/streamhelloslow", streamHelloSlow)

	//middleware
	app.Use(storeSession)
	app.Use(elapse)
	//app.Use(one, two, three)

	app.OnStart(func() {
		fmt.Printf("Started server\n")
	})
	app.OnEnd(func() {
		fmt.Printf("Stop server\n")
	})

	return app
}
