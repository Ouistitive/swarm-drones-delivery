package mapgen

type TileState int

const (
	TileChargingPoint TileState = iota
	TileDestination
	TileRooftop
	TileSpawner
	TileWarehouse
	TileEmpty
)

var AllTileStates = []TileState{
	TileChargingPoint,
	TileDestination,
	TileRooftop,
	TileSpawner,
	TileWarehouse,
	TileEmpty,
}

type Tile struct {
	FinalState 		*TileState
	PossibleStates 	[]TileState
}

type Map struct {
	Width, Height 	int
	Tiles 			[][]Tile
}

func NewTile() Tile {
	return Tile{
		FinalState:		nil,
		PossibleStates: AllTileStates,
	}
}

func NewMap(w, h int) Map {
	tiles := make([][]Tile, 0)

	for range h {
		row := []Tile{}
		for range w {
			row = append(row, NewTile())
		}
		tiles = append(tiles, row)
	}

	return Map{
		Tiles: 	tiles,
		Width: 	w,
		Height: h,
	}
}

type Direction string

const (
	Left  Direction = "left"
	Right Direction = "right"
	Up    Direction = "up"
	Down  Direction = "down"
)

type RuleSet map[
	TileState]map[
		Direction]map[
			TileState]float64

func ReadRuleSet() RuleSet {
	return RuleSet{}
}