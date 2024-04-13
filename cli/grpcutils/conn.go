package grpcutils

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcConnection() (*grpc.ClientConn, error) {
	addr := "localhost:8082"
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
		return nil, err
	}

	return conn, nil
}

func CloseGrpcConnection(conn *grpc.ClientConn) {
	if conn != nil {
		conn.Close()
	}
}
