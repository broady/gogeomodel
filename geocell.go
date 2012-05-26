// Copyright 2012 Chris Broadfoot (chris@chrisbroadfoot.id.au). All rights reserved.
// Licensed under Apache 2.
// Based on the geomodel project (http://code.google.com/p/geomodel)
package geocell

import (
	"math"
	"strings"
)

const (
	precision = 13
	base16    = "0123456789abcdef"
)

var (
	latSpan = [2]float64{-90, 90}
	lngSpan = [2]float64{-180, 180}
)

type Cell string

func Encode(latlng LatLng) Cell {
	latbits := worldToCoord(latSpan, latlng.Lat, precision*2)
	lngbits := worldToCoord(lngSpan, latlng.Lng, precision*2)
	return fromBits(latbits, lngbits, precision)
}

func worldToCoord(span [2]float64, coord float64, numbits int) int {
	numcells := math.Pow(2, float64(numbits))
	w := (span[1] - span[0]) / numcells
	n := (coord - span[0]) / w
	if n >= numcells {
		// out of bounds
		return int(numcells - 1)
	}
	return int(n)
}

func (cell Cell) Decode() LatLngBox {
	s, n := coordToWorld(latSpan, cell.latbits(), cell.Precision()*2)
	w, e := coordToWorld(lngSpan, cell.lngbits(), cell.Precision()*2)
	return LatLngBox{
		South: s,
		North: n,
		West:  w,
		East:  e,
	}
}

func coordToWorld(span [2]float64, coord int, numbits int) (min, max float64) {
	numcells := math.Pow(2, float64(numbits))
	w := (span[1] - span[0]) / numcells
	n := float64(coord)*w + span[0]
	return n, n + w
}

func (cell Cell) Precision() int {
	return len(cell)
}

func (cell Cell) East() Cell {
	return fromBits(cell.latbits(), cell.lngbits()+1, cell.Precision())
}

func (cell Cell) West() Cell {
	return fromBits(cell.latbits(), cell.lngbits()-1, cell.Precision())
}

func (cell Cell) North() (c Cell, ok bool) {
	north := cell.latbits() + 1
	if getbit(north, uint(cell.Precision()*2)) == 1 {
		return cell, false
	}
	return fromBits(north, cell.lngbits(), cell.Precision()), true
}

func (cell Cell) South() (c Cell, ok bool) {
	latbits := cell.latbits()
	if latbits == 0 {
		return cell, false
	}
	return fromBits(latbits-1, cell.lngbits(), cell.Precision()), true
}

func fromBits(lat, lng, precision int) Cell {
	cell := ""
	for i := 0; i < precision; i++ {
		// interleave bits
		n := 0
		n |= lng & 1
		n |= lat & 1 << 1
		n |= lng & 2 << 1
		n |= lat & 2 << 2

		lat >>= 2
		lng >>= 2
		cell = string(base16[n]) + cell
	}
	return Cell(cell)
}

func (cell Cell) lngbits() int {
	return cell.deinterleave(0)
}

func (cell Cell) latbits() int {
	return cell.deinterleave(1)
}

func (cell Cell) deinterleave(offset uint) int {
	// may fail if int is 32 bit?
	bits := 0
	for _, r := range cell {
		i := strings.Index(base16, string(r))
		bits <<= 1
		bits |= getbit(i, 2+offset)
		bits <<= 1
		bits |= getbit(i, offset)
	}
	return bits
}

// getbit returns 0 or 1 for a given bit position of an int.
func getbit(b int, pos uint) int {
	return (b >> pos) & 1
}
