package v2

import (
	"fmt"
	"sync"
	"context"
	gamepb "github.com/mehranghajari/Menchastic/backend/pkg/api/v2"
)

type roomID int64;
type userneme string;

type MenchasticServiceServer struct {
	gamepb.UnimplementedMenchasticServiceServer
	mu      sync.Mutex

	rooms gamepb.Rooms
	roomsChannels map[roomID] map[userneme] chan *gamepb.ResponseRoom
}

func (s *MenchasticServiceServer) CreateRoom( req *gamepb.RequestCreateRoom, msgStream gamepb.MenchasticService_CreateRoomServer) error {

	id := 1

	// Create Room
	room := &gamepb.Room {
		Id: int64(id),
		Name: req.Name,
		IsPrivate: req.IsPrivate,
		SecretKey: req.SecretKey,
		Owner: req.GetOwner(),
		Members: &gamepb.Members{
			Member: []*gamepb.Member{
				req.GetOwner(),
			},
		},
	}
	s.mu.Lock()
	s.rooms.Rooms = append(s.rooms.Rooms, room)
	s.mu.Unlock()

	msgChannel := make(chan *gamepb.ResponseRoom )
	// Create Room channel
	fmt.Printf("create romm channel for this room: [%v][chan]\n", req.GetOwner().GetUsername())
	s.roomsChannels[roomID(id)] = map[userneme] chan *gamepb.ResponseRoom {}
	s.roomsChannels[roomID(id)][userneme(req.GetOwner().GetUsername())] = msgChannel

	for {
		select {
		case <- msgStream.Context().Done():
			return nil
		case msg := <- msgChannel:
			msgStream.Send(msg)
		}
	}
}

func (s *MenchasticServiceServer) ListRoom( C context.Context,req *gamepb.RequestGame)  ( *gamepb.Rooms, error) {
	
	for _ , room := range s.rooms.Rooms{
		fmt.Println(room.GetName() , room.GetId())
	}

	return &s.rooms, nil

}


func (s *MenchasticServiceServer) JoinRoom( req *gamepb.RequestJoinRoom, msgStream gamepb.MenchasticService_JoinRoomServer) error {

	msgChannel := make(chan *gamepb.ResponseRoom )

	s.roomsChannels[roomID(req.GetId())][userneme(req.Member.GetUsername())] = msgChannel
	
	// doing this never closes the stream
	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-msgChannel:
			fmt.Printf("GO ROUTINE (Joined): %v \n", msg)
			msgStream.Send(msg)
		}
	}
}


func NewServer() *MenchasticServiceServer {
	s := &MenchasticServiceServer{
		roomsChannels: make(map[roomID] map[userneme] chan *gamepb.ResponseRoom ),
		rooms: gamepb.Rooms{},
	}
	fmt.Println(s)
	return s
}
