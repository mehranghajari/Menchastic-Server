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
			Members: []*gamepb.Member{
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

func (s *MenchasticServiceServer) CreateNewRoom( id int) {

	// Create Room
	room := &gamepb.Room {
		Id: int64(id),
		Name: "random",
		IsPrivate: "none",
		SecretKey: "none",
		Owner: nil,
		Members: &gamepb.Members{
			Members: []*gamepb.Member{},
		},
	}
	s.mu.Lock()
	s.rooms.Rooms = append(s.rooms.Rooms, room)
	s.mu.Unlock()

}

func (s *MenchasticServiceServer) ListRoom( C context.Context,req *gamepb.RequestGame)  ( *gamepb.Rooms, error) {
	
	for _ , room := range s.rooms.Rooms{
		fmt.Println(room.GetName() , room.GetId())
	}

	return &s.rooms, nil

}


func (s *MenchasticServiceServer) JoinRoom( req *gamepb.RequestJoinRoom, msgStream gamepb.MenchasticService_JoinRoomServer) error {
	var id int ;
	var lastRoom *gamepb.Room; 
	var isFull bool ;
	if req.GetId() == -1 { // join randomly to a room 
		if len(s.rooms.Rooms) == 0 {
			s.CreateNewRoom(1)
		}
		lastRoom = s.rooms.Rooms[len(s.rooms.Rooms) - 1] 
		id = int(lastRoom.Id)
	}else {
		id = int(req.GetId())
	}

	
	msgChannel := make(chan *gamepb.ResponseRoom )

	if s.roomsChannels[roomID(id)] == nil {
		s.roomsChannels[roomID(id)] = map[userneme] chan *gamepb.ResponseRoom {}
	}
	s.roomsChannels[roomID(id)][userneme(req.Member.GetUsername())] = msgChannel
	
	fmt.Println(req.Member.GetUsername() , "joined the room " , id)

	lastRoom.Members.Members = append(lastRoom.Members.Members, req.GetMember())

	isFull = len(lastRoom.Members.Members) == 4
	if isFull{
		s.CreateNewRoom(int(lastRoom.GetId())+1)
	}
	
	// send message to other members of room
	go func() {
		streams := s.roomsChannels[roomID(id)]
		for senderName, msgChan := range streams {
			if string(senderName) != req.Member.GetUsername(){
				msgChan <- &gamepb.ResponseRoom{
					Room: lastRoom,
					IsFull: isFull,
				}
			}
		}
	}()
	
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
