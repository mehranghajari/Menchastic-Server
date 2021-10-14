package v1

import (
	"fmt"
	"io"
	"sync"
	chatpb "github.com/mehranghajari/Menchastic/backend/pkg/api/v1"
)

type chatServiceServer struct {
	chatpb.UnimplementedChatServiceServer
	mu      sync.Mutex
	channel map[string] map[string] chan *chatpb.Message 
}

func (s *chatServiceServer) JoinChannel(ch *chatpb.Channel, msgStream chatpb.ChatService_JoinChannelServer) error {
	
	msgChannel := make(chan *chatpb.Message)

	if s.channel[ch.Name] == nil{
		fmt.Printf("create map for this room: [%v][chan]\n", ch.SendersName)
		s.channel[ch.Name] = map[string] chan *chatpb.Message {}
	}


	s.channel[ch.Name][ch.SendersName] = msgChannel
	// doing this never closes the stream
	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-msgChannel:
			fmt.Printf("GO ROUTINE (got message): %v \n", msg)
			msgStream.Send(msg)
		}
	}
}

func (s *chatServiceServer) SendMessage(msgStream chatpb.ChatService_SendMessageServer) error {
	msg, err := msgStream.Recv()

	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}

	ack := chatpb.MessageAck{Status: "SENT"}
	msgStream.SendAndClose(&ack)

	go func() {
		streams := s.channel[msg.Channel.Name]
		for senderName, msgChan := range streams {
			if senderName != msg.Sender{
				msgChan <- msg
			}
		}
	}()

	return nil
}

func NewServer() *chatServiceServer {
	s := &chatServiceServer{
		channel: make(map[string] map[string] chan *chatpb.Message ),
	}
	fmt.Println(s)
	return s
}
