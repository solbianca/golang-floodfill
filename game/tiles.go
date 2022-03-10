package game

import "floodfill/utils"

type Tile struct {
	utils.Address
	IsBlock bool
	RoomId  int
}

func newTile(column, row, roomId int) *Tile {
	return &Tile{RoomId: roomId, Address: utils.NewAddress(column, row), IsBlock: false}
}

type TileCollection struct {
	tiles         map[int]map[int]*Tile
	Columns, Rows int
}

func NewTileCollection(columns, rows int) *TileCollection {
	tiles := &TileCollection{tiles: map[int]map[int]*Tile{}, Columns: columns, Rows: rows}

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			tiles.Set(newTile(column, row, 0))
		}
	}

	return tiles
}

func (c *TileCollection) Set(t *Tile) {
	rows, ok := c.tiles[t.Column]
	if !ok {
		rows = map[int]*Tile{}
	}

	rows[t.Row] = t

	c.tiles[t.Column] = rows
}

func (c *TileCollection) Has(column, row int) bool {
	rows, ok := c.tiles[column]
	if !ok {
		return false
	}

	_, ok = rows[row]

	if !ok {
		return false
	}

	return true
}

func (c *TileCollection) Get(column, row int) *Tile {
	rows, ok := c.tiles[column]
	if !ok {
		return nil
	}

	tile, ok := rows[row]

	if !ok {
		return nil
	}

	return tile
}

func (c *TileCollection) All() []*Tile {
	tiles := []*Tile{}

	for _, rows := range c.tiles {
		for _, tile := range rows {
			tiles = append(tiles, tile)
		}
	}

	return tiles
}
