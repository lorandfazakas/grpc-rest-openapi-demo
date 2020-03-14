package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/lorandfazakas/grpc-rest-openapi-demo/pb"
	"github.com/lorandfazakas/grpc-rest-openapi-demo/store"
	"google.golang.org/grpc"
)

const grpcPort = ":10000"
const httpPort = ":10001"

func main() {
	log.Println("Starting application")
	store.ReadJson()
	log.Println("Starting grpc server on port " + grpcPort)
	go serveGRPC()

	log.Println("Starting HTTP server on port " + httpPort)
	serveHTTP()
}

func serveGRPC() {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterToDoServiceServer(server, new(ToDoService))
	server.Serve(listen)
}

func serveHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterToDoServiceHandlerFromEndpoint(ctx, gw, "localhost"+grpcPort, opts)
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./swaggerui"))
	mux.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))
	mux.Handle("/api/", gw)

	return http.ListenAndServe(httpPort, mux)
}

// ToDoService implements the available services for todos
type ToDoService struct{}

// GetAllToDos returns all the ToDos
func (s *ToDoService) GetAllToDos(ctx context.Context, req *pb.Empty) (*pb.ToDos, error) {
	toDos := store.GetAllToDos()
	log.Println("GetAllToDos called")
	return &toDos, nil
}

// GetToDoByID returns a ToDo by ID
func (s *ToDoService) GetToDoByID(ctx context.Context, id *pb.ID) (*pb.ToDo, error) {
	toDo := store.GetToDoByID(id.Id)
	return toDo, nil
}

// CreateToDo creates a new ToDo
func (s *ToDoService) CreateToDo(ctx context.Context, toDo *pb.ToDo) (*pb.Empty, error) {
	store.CreateToDo(*toDo)
	return &pb.Empty{}, nil
}

// FinishToDo marks a ToDo as done
func (s *ToDoService) FinishToDo(ctx context.Context, id *pb.ID) (*pb.Empty, error) {
	store.FinishToDo(id.Id)
	return &pb.Empty{}, nil
}
