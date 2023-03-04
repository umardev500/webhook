package config

import (
	"log"
	"os"
	"webhook/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection struct{}

func (c *Connection) getConn(port string) (conn *grpc.ClientConn) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithNoProxy())
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (c *Connection) OrderConn() (client pb.OrderServiceClient) {
	port := os.Getenv("ORDER_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewOrderServiceClient(conn)

	return
}

func (c *Connection) CustomerConn() (client pb.CustomerServiceClient) {
	port := os.Getenv("CUSTOMER_RPC_PORT")
	conn := c.getConn(port)
	client = pb.NewCustomerServiceClient(conn)

	return
}
