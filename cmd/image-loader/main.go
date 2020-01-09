package main

import (
	"fmt"
	execute "github.com/alexellis/go-execute/pkg/v1"
	"github.com/julienschmidt/httprouter"
	"image-loader/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.PUT("/image_load", handlers.LoadImage)
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		cmd := execute.ExecTask{
			Command:      "docker",
			Args:         []string{"images"},
			StreamStdio:  false,
		}

		res, err := cmd.Execute()
		if err!=nil{
			w.Write([]byte("Couldn't execute docker:%s"))
		}
		if res.ExitCode != 0 {
			panic("Non-zero exit code: " + res.Stderr)
		}

		w.Write([]byte(res.Stdout))

	})

	fmt.Printf("Starting Server at 3000")
	err := http.ListenAndServe("0.0.0.0:3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
