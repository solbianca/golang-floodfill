package floodfill

import (
	"floodfill/game"
	"floodfill/utils"
)

func FloodFill(tiles *game.TileCollection) *RoomCollection {
	rooms := newRoomCollection()
	addresses := utils.NewAddressCollection()

	for _, tile := range tiles.All() {
		addresses.Set(tile.Address)
	}

	for column := 0; column < tiles.Columns; column++ {
		for row := 0; row < tiles.Rows; row++ {
			tile := tiles.Get(column, row)
			if tile == nil || tile.IsBlock {
				addresses.Remove(tile)
				continue
			}

			room := resolveRoom(utils.NewAddress(column, row), addresses, tiles)

			if room.Len() > 0 {
				rooms.add(room)
			}
		}
	}

	for _, room := range rooms.Rooms {
		for _, address := range room.addresses {
			tile := tiles.Get(address.GetAddress())

			tile.RoomId = room.RoomId
			tiles.Set(tile)
		}
	}

	return rooms
}

func resolveRoom(
	address utils.Address,
	addresses *utils.AddressCollection,
	tiles *game.TileCollection,
) *room {
	room := newRoom()

	if !addresses.Has(address.GetAddress()) {
		return room
	}

	toVisit := utils.NewAddressStack()
	toVisit.Push(address)
	addresses.Remove(address)
	toVisit = findNeighbors(address, toVisit, addresses)

	for {
		if toVisit.Empty() {
			room.resolve()
			return room
		}

		address := toVisit.Pop()
		addresses.Remove(address)
		tile := tiles.Get(address.GetAddress())

		if tile == nil || tile.IsBlock == true {
			continue
		}

		room.add(tile.RoomId, tile)
		toVisit = findNeighbors(address, toVisit, addresses)

		if toVisit.Empty() {
			room.resolve()
			return room
		}
	}
}

func findNeighbors(
	address utils.Addressed,
	toVisit *utils.AddressStack,
	addresses *utils.AddressCollection,
) *utils.AddressStack {
	column, row := address.GetAddress()

	if addresses.Has(column+1, row) {
		address := utils.NewAddress(column+1, row)
		toVisit.Push(address)
		addresses.Remove(address)
	}
	if addresses.Has(column-1, row) {
		address := utils.NewAddress(column-1, row)
		toVisit.Push(address)
		addresses.Remove(address)
	}
	if addresses.Has(column, row+1) {
		address := utils.NewAddress(column, row+1)
		toVisit.Push(address)
		addresses.Remove(address)
	}
	if addresses.Has(column, row-1) {
		address := utils.NewAddress(column, row-1)
		toVisit.Push(address)
		addresses.Remove(address)
	}
	if addresses.Has(column+1, row+1) {
		address := utils.NewAddress(column+1, row+1)
		toVisit.Push(address)
		addresses.Remove(address)
	}
	if addresses.Has(column+1, row-1) {
		address := utils.NewAddress(column+1, row-1)
		toVisit.Push(address)
		addresses.Remove(address)
	}
	if addresses.Has(column-1, row-1) {
		address := utils.NewAddress(column-1, row-1)
		toVisit.Push(address)
		addresses.Remove(address)
	}
	if addresses.Has(column-1, row+1) {
		address := utils.NewAddress(column-1, row+1)
		toVisit.Push(address)
		addresses.Remove(address)
	}

	return toVisit
}
