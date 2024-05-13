package main

import (
	"client/services"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	creds := insecure.NewCredentials()
	grpc, err := grpc.Dial("localhost:8080",
		grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer grpc.Close()

	Beefclient := services.NewBeefClient(grpc)

	res, err := Beefclient.CountBeef(context.Background(), &services.BeefRequest{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
