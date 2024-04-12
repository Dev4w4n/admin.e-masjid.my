package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Dev4w4n/admin.e-masjid.my/api/pb"
	"github.com/Dev4w4n/e-masjid.my/api/core/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	env, err := env.GetEnvironment()
	if err != nil {
		log.Fatalf("Error getting environment: %v", err)
	}

	addr := "localhost:" + env.ServerPort
	cred := insecure.NewCredentials()

	// Need this if want to detect not able to connect to server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(cred),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatalf("Error dialing: %v", err)
	}
	defer conn.Close()

	log.Printf("Client connected to %s", addr)
	client := pb.NewTenantsClient(conn)

	fmt.Println(client)
}
