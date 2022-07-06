package grpc

// import (
// 	"context"
// 	"log"
// 	"net"
// 	"os"

// 	"github.com/turnixxd/grpc-rest-graphql-comparison/go/database"
// 	"github.com/turnixxd/grpc-rest-graphql-comparison/go/env"
// 	pb "github.com/turnixxd/grpc-rest-graphql-comparison/go/grpc/proto"
// 	"google.golang.org/grpc"
// )

// // Constructor
// func NewUserServiceServer() *UserServiceServer {
// 	return &UserServiceServer{}
// }

// type UserServiceServer struct {
// 	pb.UserServiceServer
// }

// var db database.Database

// func (server *UserServiceServer) Run() error {
// 	lis, err := net.Listen("tcp", ":"+env.Process("SERVER_PORT"))
// 	if err != nil {
// 		log.Fatal("Failed to listen on port %v, %v", env.Process("SERVER_PORT"), err)
// 	}

// 	s := grpc.NewServer()
// 	databaseImplementation := os.Args[1]
// 	db, err = database.Factory(databaseImplementation)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// cannot be type of &UserManagementServer{} because we are using the server that Run() receives
// 	pb.RegisterUserServiceServer(s, server)
// 	log.Printf("gRPC server listening at %v", lis.Addr())
// 	return s.Serve(lis)
// }

// func (s *UserServiceServer) Set(ctx context.Context, in *pb.NewUser) (*pb.ServerResponse, error) {
// 	value, err := db.Set(in.GetUsername(), in.GetPassword(), in.GetRole())
// 	return generateResponse(value, err)
// }

// func (s *UserServiceServer) Get(ctx context.Context, in *pb.User) (*pb.ServerResponse, error) {
// 	value, err := db.Get(in.GetId(), in.GetUsername(), in.GetPassword(), in.GetRole())
// 	return generateResponse(value, err)
// }

// func (s *UserServiceServer) Delete(ctx context.Context, in *pb.DeleteUser) (*pb.ServerResponse, error) {
// 	value, err := db.Delete(in.GetId(), in.GetUsername(), in.GetPassword(), in.GetRole())
// 	return generateResponse(value, err)
// }

// func generateResponse(id int32, err error) (*pb.ServerResponse, error) {
// 	if err != nil {
// 		return &pb.ServerResponse{State: false, Id: id, Error: err.Error()}, nil
// 	}
// 	return &pb.ServerResponse{State: true, Id: id, Error: ""}, nil
// }

// func Serve() {
// 	var basic_service_server *UserServiceServer = NewUserServiceServer()
// 	if err := basic_service_server.Run(); err != nil {
// 		log.Fatal("Failed to serve: %v", err)
// 	}
// }
