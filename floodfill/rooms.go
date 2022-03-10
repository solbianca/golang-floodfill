package floodfill

import "floodfill/utils"

var idGenerator *RoomIdGenerator

func init() {
	idGenerator = NewRoomIdGenerator()
}

type RoomIdGenerator struct {
	lastId int
}

func NewRoomIdGenerator() *RoomIdGenerator {
	return &RoomIdGenerator{lastId: 1}
}

func (g *RoomIdGenerator) Next() int {
	g.lastId++

	return g.lastId
}

type room struct {
	RoomId    int
	addresses []utils.Addressed
	counters  map[int]int
}

func newRoom() *room {
	return &room{addresses: []utils.Addressed{}, counters: map[int]int{}, RoomId: -1}
}

func (r *room) add(roomId int, address utils.Addressed) {
	r.addresses = append(r.addresses, address)

	counter, ok := r.counters[roomId]
	if !ok {
		counter = 0
	}
	counter++
	r.counters[roomId] = counter
}

func (r *room) resolve() {
	targetRoomId, maxValue := -1, 0

	for roomId, counter := range r.counters {
		if counter > maxValue {
			maxValue = counter
			targetRoomId = roomId
		}
	}

	r.RoomId = targetRoomId
}

func (r *room) Len() int {
	return len(r.addresses)
}

type RoomCollection struct {
	Rooms     map[int]*room
	addresses map[utils.Addressed]int
}

func newRoomCollection() *RoomCollection {
	return &RoomCollection{Rooms: map[int]*room{}, addresses: map[utils.Addressed]int{}}
}

func (rc *RoomCollection) add(addedRoom *room) {
	room, ok := rc.Rooms[addedRoom.RoomId]

	if !ok {
		rc.Rooms[addedRoom.RoomId] = addedRoom
		return
	}

	if addedRoom.Len() > room.Len() {
		delete(rc.Rooms, room.RoomId)
		room.RoomId = idGenerator.Next()
		rc.Rooms[room.RoomId] = room
	} else {
		addedRoom.RoomId = idGenerator.Next()
	}

	rc.Rooms[addedRoom.RoomId] = addedRoom
}

func (rc *RoomCollection) GetRoomIdByAddress(address utils.Addressed) int {
	roomId, ok := rc.addresses[address]

	if ok {
		return roomId
	}

	return 0
}
