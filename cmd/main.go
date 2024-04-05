package main

import (
	"fmt"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Menu/pkg/config"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Menu/pkg/db"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Menu/pkg/services"
	"lendral3n/KulinerKlasik-Microservices-gRPC-Menu/pkg/pb"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
    c, err := config.LoadConfig()

    if err != nil {
        log.Fatalln("Failed at config", err)
    }

   h := db.Init(
		c.DB_USERNAME,
		c.DB_PASSWORD,
		c.DB_HOSTNAME,
		c.DB_PORT,
		c.DB_NAME,
	)

    lis, err := net.Listen("tcp", c.Port)

    if err != nil {
        log.Fatalln("Failed to listing:", err)
    }

    fmt.Println("Menu Svc on", c.Port)

    s := services.Server{
        H:   h,
    }

    grpcServer := grpc.NewServer()

    pb.RegisterMenuServiceServer(grpcServer, &s)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalln("Failed to serve:", err)
    }
}