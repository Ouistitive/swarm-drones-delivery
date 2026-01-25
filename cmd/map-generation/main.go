package main

import (
	"os"
	"strconv"
	"swarm-drones-delivery/internal/mapgen"
)

func getParameters(args []string) (width, height int) {
	width = 50
	height = 50

	if len(args) >= 2 {
		if w, err := strconv.Atoi(args[0]); err == nil {
			width = w
		}
		if h, err := strconv.Atoi(args[1]); err == nil {
			height = h
		}
	}

	return
}

func main() {
	argsWithoutProg := os.Args[1:]
	width, height := getParameters(argsWithoutProg)
	
	m := mapgen.WaveFunctionCollapse(width, height)
	mapgen.WriteMap(&m)
}