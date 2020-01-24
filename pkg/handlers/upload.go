package handlers

import (
	"image-loader/pkg/docker"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// LoadImage takes a httpRequest decodes it and loads the multipart content into docker
func LoadImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Save file locally
	log.Info("Recieved a Load Request")
	reader, err := GetFileReader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	log.Info("Talking to docker to load the image")
	// Make docker load the image
	resp, err := docker.LoadDockerImage(r.Context(), reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
	}

	w.Write(resp)
}

// GetFileReader returns a file reader from a http request
func GetFileReader(r *http.Request) (io.Reader, error) {

	r.ParseMultipartForm(32 << 20)
	log.Info("Reading MultiPart form data")

	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}

	return file, nil
}
