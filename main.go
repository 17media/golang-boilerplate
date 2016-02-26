package main

import (
        "os"
        "fmt"
        "github.com/codegangsta/negroni"
        "github.com/julienschmidt/httprouter"
        "net/http"
        "github.com/facebookgo/grace/gracehttp"
        config "github.com/spf13/viper"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        fmt.Fprint(w, "WelWWWW111come1111!\n")
}

// Hello function
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
        router := httprouter.New()
        router.GET("/", Index)
        router.GET("/hello/:name", Hello)

        User.hashPassword
        // Middleware
        n := negroni.Classic()
        n.UseHandler(router)

        // Start server
        gracehttp.Serve(
                &http.Server{
                        Addr: ":" + os.Getenv("PORT"),
                        Handler: n,
                },
        )
}
