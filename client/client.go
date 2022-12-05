package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	"chat/proto"
)

// create a user
func createClient(ctx context.Context, c proto.ChatServiceClient, username string) error {
	// create user
	stream, err := c.Connect(ctx, &proto.ConnectRequest{
		Username: username,
	})
	if err != nil {
		log.Fatalf("error on client.Connect(): %s", err)
	}

	fmt.Printf("%s has joined chat server\n", username)

	// listen and write messages to stdout
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			streamResp, err := stream.Recv()
			if err != nil {
				return err
			}

			// print out to stdout: [username: message]
			fmt.Printf("%s: %s\n", username, streamResp.Message)
		}

	}
}

func sendMessage()

func main() {
	// dial server
	conn, err := grpc.Dial(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error on grpc.Dial(): %s", err)
	}

	// create a client
	client := proto.NewChatServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	// create a user: John Doe
	go func() {
		createClient(ctx, client, "John")
	}()

	// create another user: Jane Doe
	go func() {
		createClient(ctx, client, "Jane")
	}()

	// give 2 seconds for users to join
	time.Sleep(2 * time.Second)

	// send message from John to Jane
	// FIXEME: lines below not complete
	johnCtx := metadata.AppendToOutgoingContext(ctx, "username", "John")
	client.SendMessage(johnCtx, &proto.SendMessageRequest{
		UsernameOrChannel: "Jane",
	})

	<-ctx.Done()
}
