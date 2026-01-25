package world

import (
	"os"
	"strings"
)

func ReadMap(mapPath, addrPath string) (*Map, error) {
	content, err := os.ReadFile(mapPath)
	if err != nil {
		return nil, err
	}
	addr, err := ReadAddresses(addrPath)
	if err != nil {
		return nil, err
	}

	return loadWorld(string(content), addr), nil
}

func loadWorld(content string, addr []Address) *Map {
	lines := strings.Split(content, "\n")
	cells := make([][]rune, 0)
	rooftops := make([]Position, 0)
	spawners := make([]Position, 0)
	deliveryDestinations := make([]DeliveryDestination, 0)
	warehouses := make([]Position, 0)
	chargingPoints := make([]ChargingPoint, 0)

	addrIdx := 0
	y := 0.0
	for _, line := range lines {
		row := []rune(line)
		cells = append(cells, row)
		for x, r := range line {
			switch r {
			case 'R':
				rooftops = append(rooftops, NewPosition(float64(x), y))
			case 'S':
				spawners = append(spawners, NewPosition(float64(x), y))
			case 'D':
				deliveryDestinations = append(deliveryDestinations, NewDeliveryDestination(NewPosition(float64(x), y), addr[addrIdx]))
				// addrIdx++
			case 'W':
				warehouses = append(warehouses, NewPosition(float64(x), y))
			case 'C':
				chargingPoints = append(chargingPoints, NewChargingPoint(NewPosition(float64(x), y)))
			}
		}

		y++
	}

	return &Map{Width: len(lines[0]), Height: len(lines), Cells: cells, Rooftops: rooftops, Spawners: spawners, DeliveryDests: deliveryDestinations, Warehouses: warehouses, ChargingPoints: chargingPoints}
}
