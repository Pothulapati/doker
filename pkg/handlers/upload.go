package handlers

import (
	"fmt"
	"image-loader/pkg/docker"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func LoadImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Save file locally
	fmt.Print("Loading Image")
	reader, err := saveFileLocally(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	// Make docker load the image
	resp, err := docker.LoadDockerImage(r.Context(), reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	w.Write(resp)
}

func saveFileLocally(r *http.Request) (io.Reader, error) {

	r.ParseMultipartForm(32 << 20)

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}

	return file, nil
}
