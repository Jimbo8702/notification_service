package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Jimbo8702/notification_service/types"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	client, err := NewMongoClient(os.Getenv("MONGO_CONNECT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	fbcm, err := NewFirebaseMessageClient(os.Getenv("FIREBASE_SERVICE_ACCOUNT_ID"))
	if err != nil {
		log.Fatal(err)
	}
	var (
		store = NewMongoNotificationStore(client, os.Getenv("MONGO_DBNAME"))
		svc   = NewFCMNotificationService(fbcm, store)
		svcwl = NewLogMiddleware(svc)
	)
	log.Fatal(makeGRPCTransport(os.Getenv("GRPC_LISTEN_ADDR"), svcwl))
}

func makeGRPCTransport(listenAddr string, svc NotificationService) error {
	fmt.Println("GRPC transport running on port:", listenAddr)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	server := grpc.NewServer([]grpc.ServerOption{}...)
	types.RegisterNotificationServiceServer(server, NewGRPCNotificationServer(svc))
	return server.Serve(ln)
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}