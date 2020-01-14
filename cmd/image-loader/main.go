package main

import (
	"fmt"
	"image-loader/pkg/docker"
	"image-loader/pkg/handlers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.Write([]byte("hi"))

	})
	router.PUT("/image_load", handlers.LoadImage)
	router.GET("/list", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		imagesResp, err := docker.GetDockerImages(r.Context())
		if err != nil {
			log.Fatal(err)
		}
		w.Write(imagesResp)

	})

	fmt.Printf("Starting Server at 3000")
	err := http.ListenAndServe("0.0.0.0:3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
