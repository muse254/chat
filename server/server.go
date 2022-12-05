package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"chat/proto"
)

// chatServer implements the persistance for the chat service(for this implmentation).
type chatServer struct {
	users    map[string]chan string
	channels map[string]channel

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
				removeUser(chatroom.participants, req.Username)
				fmt.Printf("%s removed from channel %s\n", req.Username, channel)
			}

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

	// the context carries the username; extract username
	// if username is not present, account was not created
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, errors.New("the client does not have a username")
	}

	// add participant to channel
	channel.participants = append(channel.participants, username)
	fmt.Printf("%s has joined channel %s", username, req.Channel)

	return &emptypb.Empty{}, nil
}

func (c *chatServer) LeftGroupChat(ctx context.Context, req *proto.LeftGroupChatRequest) (*emptypb.Empty, error) {
	// check if channel name exists, if not return error
	channel, ok := c.channels[req.Channel]
	if !ok {
		return nil, errors.New("the channel does not exist")
	}

	// the context carries the username; extract username
	// if username is not present, account was not created
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, errors.New("the client does not have a username")
	}

	// remove user from the channel
	if removeUser(channel.participants, username) {
		return nil, errors.New("user was not in channel")
	}

	return &emptypb.Empty{}, nil
}

func (c *chatServer) CreateGroupChat(ctx context.Context, req *proto.CreateGroupChatRequest) (*emptypb.Empty, error) {
	// if channel already exists, return an error
	if _, ok := c.channels[req.Channel]; ok {
		return nil, errors.New("the channel already exists")
	}

	// the context carries the username; extract username
	// if username is not present, account was not created
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, errors.New("the client does not have a username")
	}

	// create channel and set fist participant as current client
	c.channels[req.Channel] = channel{
		participants: []string{username},
		messages:     make(chan string),
	}

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

func removeUser(users []string, username string) bool {
	for i, user := range users {
		if user == username {
			users[i] = users[len(users)-1]
			users = users[:len(users)-1]
			return true
		}
	}
	return false
}

func main() {
	listener, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("error occured on net.Listen(): %s", err)
	}

	srv := grpc.NewServer()
	proto.RegisterChatServiceServer(srv, &chatServer{
		users: map[string]chan string{},
	})

	if err = srv.Serve(listener); err != nil {
		log.Fatalf("err occurred on srv.Serve(): %s", err)
	}
}
