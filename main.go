package main

import (
    "os"
    "fmt"
    "github.com/codegangsta/negroni"
    "github.com/julienschmidt/httprouter"
    "net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "WelWWWW111come1111!\n")
}

// Handler
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    // Middleware
    n := negroni.Classic()
    n.UseHandler(router)

    // Start server
    n.Run(":" + os.Getenv("PORT"))
}
