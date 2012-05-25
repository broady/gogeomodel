// Copyright 2012 Chris Broadfoot (chris@chrisbroadfoot.id.au). All rights reserved.
// Licensed under Apache 2.
package geocell

type Cell string

type LatLng struct {
	Lat, Lng float64
}

type LatLngBox struct {
	South, North, West, East float64
}

func (b LatLngBox) NorthWest() LatLng {
	return LatLng{b.North, b.West}
}

func (b LatLngBox) SouthEast() LatLng {
	return LatLng{b.South, b.East}
}

func (b LatLngBox) Contains(l LatLng) bool {
	if l.Lat < b.South || l.Lat > b.North {
		return false
	}
	if l.Lng < b.West || l.Lng > b.East {
		return false
	}
	return true
}
