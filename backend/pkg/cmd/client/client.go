package main

import (
	//"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	//"os"
	"google.golang.org/grpc"
	gamepb "github.com/mehranghajari/Menchastic/backend/pkg/api/v2"
)

var roomID = flag.String("roomID", "1", "room id for game")
var senderName = flag.String("sender", "Mehran", "Senders name")
var tcpServer = flag.String("server", ":5400", "Tcp server")

func joinRoom(ctx context.Context, client gamepb.MenchasticServiceClient) {
	
	requestJoinRoom := gamepb.RequestJoinRoom{
		Id: -1,
		SecretKey: "123",
		Member: &gamepb.Member{
			Username: *senderName,
			DisplayName: "mehranghajari",
		},
	}
	stream, err := client.JoinRoom(ctx, &requestJoinRoom)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}

	fmt.Printf("Joined room\n")

	for {
		responseRoom, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive message from channel joining. \nErr: %v", err)
		}

		fmt.Printf("updated room: (%v) -> %v \n", responseRoom.GetRoom().Name, responseRoom.GetRoom().Id)
		fmt.Printf("number of members : %d , isFull : %v \n" , len(responseRoom.GetRoom().GetMembers().Members) , responseRoom.GetIsFull())
	}
}

func createRoom(ctx context.Context, client gamepb.MenchasticServiceClient) {
	
	requestCreateRoom := gamepb.RequestCreateRoom{
		Name: "Mench",
		IsPrivate: "no",
		SecretKey: "123",
		Owner: &gamepb.Member{
			Username: *senderName,
			DisplayName: "mehranghajari",
		},
	}
	stream, err := client.CreateRoom(ctx, &requestCreateRoom)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}

	fmt.Printf("created room: %v \n", *roomID)

	for {
		responseRoom, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("Failed to receive message from channel joining. \nErr: %v", err)
		}

		fmt.Printf("updated room: (%v) -> %v \n", responseRoom.GetRoom().Name, responseRoom.GetRoom().Id)
		fmt.Printf("number of members : %d \n" , len(responseRoom.GetRoom().GetMembers().Members))
	}
}

// func sendMessage(ctx context.Context, client chatpb.ChatServiceClient, message string) {
// 	stream, err := client.SendMessage(ctx)
// 	if err != nil {
// 		log.Printf("Cannot send message: error: %v", err)
// 	}
// 	msg := chatpb.Message{
// 		Channel: &chatpb.Channel{
// 			Name:        *channelName,
// 			SendersName: *senderName},
// 		Message: message,
// 		Sender:  *senderName,
// 	}
// 	stream.Send(&msg)
//
// 	ack, err := stream.CloseAndRecv()
// 	fmt.Printf("Message sent: %v \n", ack)
// }

func main() {

	flag.Parse()

	fmt.Println("--- CLIENT APP ---")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dail: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()
	client := gamepb.NewMenchasticServiceClient(conn)

	//createRoom(ctx, client)
	joinRoom(ctx , client)

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	go sendMessage(ctx, client, scanner.Text())
	// }

}