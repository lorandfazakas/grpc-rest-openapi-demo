package main

import (
	"context"
	"fmt"

	pb "github.com/lorandfazakas/grpc-rest-openapi-demo/pb"
	"google.golang.org/grpc"
)

var empty pb.Empty

func main() {
	serverAddr := ":10000"
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("error connecting: ", err)
	}
	defer conn.Close()
	client := pb.NewToDoServiceClient(conn)
	toDos, err := client.GetAllToDos(context.Background(), &empty)
	if err != nil {
		fmt.Println("error in grpc call:", err)
	}
	for _, t := range toDos.ToDo {
		fmt.Println(t)
	}

}
