package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	db "github.com/turnixxd/grpc-rest-graphql-comparison/go/database"
)

func HandleGetImages(w http.ResponseWriter, r *http.Request) {
	id, valid := r.URL.Query()["id"]

	if valid {
		imageID := id[0]
		image := db.GetImage(imageID)

		// if ID is empty
		if image.Id == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("{}"))
			return
		}

		imageBytes, err := json.Marshal(image)

		if err != nil {
			panic(err)
		}

		w.Write(imageBytes)
		return
	}

	images := db.GetImages()
	imageBytes, err := json.Marshal(images)

	if err != nil {
		panic(err)
	}

	w.Write(imageBytes)
}

func HandleUpdateImages(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		_, valid := r.URL.Query()["id"]

		if valid {
			var image db.Image
			err = json.Unmarshal(body, &image)
			if err != nil {
				w.WriteHeader(400)
				fmt.Fprintf(w, "Bad request")
			}

			db.SaveImage(image)
			return
		}

		var images []db.Image
		err = json.Unmarshal(body, &images)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "Bad request")
		}

		db.SaveImages(images)
		return

	} else {
		w.WriteHeader(405)
		fmt.Fprintf(w, "Method not Supported!")
	}
}
