package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
)

func LoadImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Save file locally
	fmt.Print("Loading Image")
	err := saveFileLocally(r)
	if err != nil {
		http.Error(w, "Couldn't save file locally", http.StatusBadGateway)
	}

	// Make docker load the image



}

func saveFileLocally(r *http.Request) (error) {

	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("file")
	if err!=nil {
		return err
	}

	defer file.Close()

	f, err := os.OpenFile("images/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err !=nil{
		return err
	}

	defer f.Close()
	io.Copy(f, file)

	return nil
}