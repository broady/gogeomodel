// Copyright 2012 Chris Broadfoot (chris@chrisbroadfoot.id.au). All rights reserved.
// Licensed under Apache 2.
// Based on the geomodel project (http://code.google.com/p/geomodel)
package geocell

import (
	"strings"
)

const (
	precision = 13
	base16    = "0123456789abcdef"
)

func Encode(latlng LatLng) Cell {
	cell := ""
	lats := [2]float64{-90, 90}
	lngs := [2]float64{-180, 180}

	for len(cell) < precision {
		n := 0
		// interleave bits
		n |= constrict(&lats, latlng.Lat) << 3
		n |= constrict(&lngs, latlng.Lng) << 2
		n |= constrict(&lats, latlng.Lat) << 1
		n |= constrict(&lngs, latlng.Lng)
		cell += string(base16[n])
	}
	return Cell(cell)
}

func (cell Cell) Decode() LatLngBox {
	// South, North
	lats := [2]float64{-90, 90}
	// West, East
	lngs := [2]float64{-180, 180}

	for _, r := range cell {
		i := strings.Index(base16, string(r))
		refine(&lats, (i>>3)&1)
		refine(&lngs, (i>>2)&1)
		refine(&lats, (i>>1)&1)
		refine(&lngs, i&1)
	}
	return LatLngBox{
		South: lats[0],
		North: lats[1],
		West:  lngs[0],
		East:  lngs[1],
	}
}

func constrict(pair *[2]float64, coord float64) int {
	m := mid(pair)
	if coord > m {
		pair[0] = m
		return 1
	}
	pair[1] = m
	return 0
}

func refine(pair *[2]float64, bit int) {
	if bit != 0 {
		pair[0] = mid(pair)
	} else {
		pair[1] = mid(pair)
	}
}

func mid(pair *[2]float64) float64 {
	return (pair[0] + pair[1]) / 2
}
