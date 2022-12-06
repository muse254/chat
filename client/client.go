package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

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
			fmt.Printf("%s received: %s\n", username, streamResp.Message)
		}
	}
}

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

	// allow all users to join in first
	time.Sleep(1 * time.Second)

	// send message from John to Jane
	johnCtx := metadata.AppendToOutgoingContext(ctx, "username", "John")
	client.SendMessage(johnCtx, &proto.SendMessageRequest{
		UsernameOrChannel: "Jane",
		Message:           "Hi, how are you?",
	})

	// reply from Jane
	janeCtx := metadata.AppendToOutgoingContext(ctx, "username", "Jane")
	client.SendMessage(janeCtx, &proto.SendMessageRequest{
		UsernameOrChannel: "John",
		Message:           "I'm doing great thanks for asking",
	})

	// create group chat
	client.CreateGroupChat(janeCtx, &proto.CreateGroupChatRequest{
		Channel: "Study",
	})

	// let John join the Study channel
	client.JoinGroupChat(johnCtx, &proto.JoinGroupChatRequest{
		Channel: "Study",
	})

	// list the channels
	resp, err := client.ListChannels(johnCtx, &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("All channels: %s\n", strings.Join(resp.Channels, ", "))

	// all members leave channel: Study
	if _, err = client.LeftGroupChat(johnCtx, &proto.LeftGroupChatRequest{
		Channel: "Study",
	}); err != nil {
		log.Fatal(err)
	}
	if _, err = client.LeftGroupChat(janeCtx, &proto.LeftGroupChatRequest{
		Channel: "Study",
	}); err != nil {
		log.Fatal(err)
	}

	<-janeCtx.Done()
	<-johnCtx.Done()
	<-ctx.Done()
}
