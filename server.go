package main

import (
	"./src/martini"
        "./src/render"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()

  //martini-contrib renderer
  m.Use(render.Renderer())

  //Searches on the "assets" folder for static files
  m.Use(martini.Static("assets"))
  m.Get("/", func(r render.Render) {
    r.HTML(200, "index","")
  })
  
  log.Fatal(http.ListenAndServe(":3002", m))
}
