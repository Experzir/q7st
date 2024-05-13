package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"server/services"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func main() {

	//grpc server
	go func() {
		server := grpc.NewServer()
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal(err)
		}

		services.RegisterBeefServer(server, services.NewBeefServer())

		fmt.Println("Server (grpc) listening on port :8080")
		err = server.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}
	}()

	//http server
	gofiber()

}

// httpserve
func gofiber() {
	app := fiber.New()

	app.Get("/beef/summary", func(c *fiber.Ctx) error {
		data, err := services.Count()
		if err != nil {
			return err
		}
		mapData := map[string]map[string]int{"beef": data}
		jsonBytes, err := json.Marshal(mapData)
		if err != nil {
			return err
		}
		return c.Send(jsonBytes)
	})

	fmt.Println("Server (http) listening on port :8081")
	app.Listen(":8081")
}
