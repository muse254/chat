package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat/proto"
)

// chatServer implements the persistance for the chat service(for this implmentation).
type chatServer struct {
	users    map[string]chan string
	channels map[string]*channel

	proto.UnimplementedChatServiceServer
}

type channel struct {
	participants []string
	messages     chan string
}

func (c *chatServer) Connect(req *proto.ConnectRequest, stream proto.ChatService_ConnectServer) error {
	// check if username already exists, return error if true
	if _, ok := c.users[req.Username]; ok {
		return errors.New("the username already exists")
	}

	// add username to chatServer and create
	c.users[req.Username] = make(chan string)

	// when message for user is received stream it back to the client
	ctx := stream.Context()
	for {
		select {
		case msg := <-c.users[req.Username]:
			stream.Send(&proto.ConnectResponse{
				Message: msg,
			})

		case <-ctx.Done():
			// remove user from all group channels
			for channel, chatroom := range c.channels {
				users, ok := removeUser(chatroom.participants, req.Username)
				if !ok {
					return fmt.Errorf("unable to remove %q from channel %s", req.Username, channel)
				}
				chatroom.participants = users
				fmt.Printf("%s removed from channel %s: %s\n", req.Username, channel, users)
			}

			// remove user message chan from server
			delete(c.users, req.Username)
			fmt.Printf("%q removed from server\n", req.Username)

			return ctx.Err()
		}
	}
}

func (c *chatServer) JoinGroupChat(ctx context.Context, req *proto.JoinGroupChatRequest) (*emptypb.Empty, error) {
	// check if channel name exists, if not return error
	channel, ok := c.channels[req.Channel]
	if !ok {
		return nil, errors.New("the channel does not exist")
	}

	username, err := getUsername(ctx)
	if err != nil {
		return nil, err
	}

	// add participant to channel
	channel.participants = append(channel.participants, username)
	fmt.Printf("%q has joined channel %q\n", username, req.Channel)

	return &emptypb.Empty{}, nil
}

func (c *chatServer) LeftGroupChat(ctx context.Context, req *proto.LeftGroupChatRequest) (*emptypb.Empty, error) {
	// check if channel name exists, if not return error
	channel, ok := c.channels[req.Channel]
	if !ok {
		return nil, errors.New("the channel does not exist")
	}

	username, err := getUsername(ctx)
	if err != nil {
		return nil, err
	}

	// remove user from the channel
	channel.participants, ok = removeUser(channel.participants, username)
	if !ok {
		return nil, errors.New("user was not in channel")
	}
	//channel.participants = users

	fmt.Printf("%q has left channel %s\n", username, req.Channel)

	// if no users in channel, delete channel
	if len(channel.participants) == 0 {
		delete(c.channels, req.Channel)
		fmt.Printf("%q channel deleted, no members\n", req.Channel)
	}

	return &emptypb.Empty{}, nil
}

func (c *chatServer) CreateGroupChat(ctx context.Context, req *proto.CreateGroupChatRequest) (*emptypb.Empty, error) {
	// if channel already exists, return an error
	if _, ok := c.channels[req.Channel]; ok {
		return nil, errors.New("the channel already exists")
	}

	username, err := getUsername(ctx)
	if err != nil {
		return nil, err
	}

	// create channel and set fist participant as current client
	c.channels[req.Channel] = &channel{
		participants: []string{username},
		messages:     make(chan string),
	}

	fmt.Printf("%q channel created by %q\n", req.Channel, username)

	return &emptypb.Empty{}, nil
}

func (c *chatServer) SendMessage(ctx context.Context, req *proto.SendMessageRequest) (*emptypb.Empty, error) {
	var recipient chan string

	if user, ok := c.users[req.UsernameOrChannel]; ok {
		// check if the recipient is a user
		recipient = user
	} else if channel, ok := c.channels[req.UsernameOrChannel]; ok {
		// check if the recipient is a channle
		recipient = channel.messages
	} else {
		// no recipient was found
		return nil, errors.New("the recipient was not found in usernames or channels")
	}

	recipient <- req.Message

	return &emptypb.Empty{}, nil
}

func (c *chatServer) ListChannels(ctx context.Context, _ *emptypb.Empty) (*proto.ListChannelsResponse, error) {
	channels := make([]string, 0, len(c.channels))
	for channel := range c.channels {
		channels = append(channels, channel)
	}

	return &proto.ListChannelsResponse{
		Channels: channels,
	}, nil
}

func removeUser(users []string, username string) ([]string, bool) {
	fmt.Printf("%q to remove %q\n", users, username)
	for i, user := range users {
		if user == username {
			users[i] = users[len(users)-1]
			return users[:len(users)-1], true
		}
	}
	return nil, false
}

func getUsername(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("could not retrieve metadata from context")
	}

	if res, ok := md["username"]; ok {
		return res[0], nil
	}

	return "", errors.New("the client did not provide a username")
}

func main() {
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("error occured on net.Listen(): %s", err)
	}

	srv := grpc.NewServer()
	proto.RegisterChatServiceServer(srv, &chatServer{
		users:    map[string]chan string{},
		channels: map[string]*channel{},
	})

	if err = srv.Serve(listener); err != nil {
		log.Fatalf("err occurred on srv.Serve(): %s", err)
	}
}
