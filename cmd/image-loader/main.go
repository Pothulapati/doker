package main

import (
	"fmt"
	"image-loader/pkg/docker"
	"image-loader/pkg/handlers"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

type Logger struct {
	handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Info(r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Write([]byte("hi"))

	})
	router.POST("/load", handlers.LoadImage)
	router.GET("/list", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		imagesResp, err := docker.ListDockerImages(r.Context())
		if err != nil {
			log.Fatal(err)
		}
		w.Write(imagesResp)

	})
	router.GET("/prune", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		imagesResp, err := docker.DockerPruneImages(r.Context())
		if err != nil {
			log.Fatal(err)
		}
		w.Write(imagesResp)

	})
	fmt.Println("Starting Server at 3000")
	err := http.ListenAndServe("0.0.0.0:3000", Logger{handler: router})
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Log Everything
	log.SetLevel(log.InfoLevel)
}
