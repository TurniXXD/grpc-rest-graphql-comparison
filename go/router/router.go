package router

import (
	"net/http"

	rest "github.com/turnixxd/grpc-rest-graphql-comparison/go/rest"
)

func Serve() {
	http.HandleFunc("/get", rest.HandleGetImages)
	http.HandleFunc("/update", rest.HandleUpdateImages)
	//http.HandleFunc("/graphql", HandleUpdateImages)
	//http.HandleFunc("/query", HandleUpdateImages)

	http.ListenAndServe(":80", nil)
}
