package main

import (
	"github.com/chevelleboyer/go-spaceadventure-walkthrough/internal/spaceadventure"
)

func main() {
	var ps = spaceadventure.PlanetarySystem{"Solare System"}
	spaceadventure.Start(ps)
}


