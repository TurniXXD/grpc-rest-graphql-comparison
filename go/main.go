package main

// import (
// 	server "github.com/turnixxd/grpc-rest-graphql-comparison/go/grpc"
// )

import (
	db "github.com/turnixxd/grpc-rest-graphql-comparison/go/database"
	r "github.com/turnixxd/grpc-rest-graphql-comparison/go/router"
)

func main() {
	db.Database()
	r.Serve()
}
