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
	latbits := constrict([2]float64{-90, 90}, latlng.Lat, precision*2)
	lngbits := constrict([2]float64{-180, 180}, latlng.Lng, precision*2)
	return fromBits(latbits, lngbits, precision)
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

func (cell Cell) East() Cell {
	return fromBits(cell.latbits(), cell.lngbits()+1, len(cell))
}

func (cell Cell) West() Cell {
	return fromBits(cell.latbits(), cell.lngbits()-1, len(cell))
}

func (cell Cell) North() Cell {
	return fromBits(cell.latbits()+1, cell.lngbits(), len(cell))
}

func (cell Cell) South() Cell {
	return fromBits(cell.latbits()-1, cell.lngbits(), len(cell))
}

func fromBits(lat, lng, precision int) Cell {
	cell := ""
	// interleave
	for i := 0; i < precision; i++ {
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
	bits := 0
	for _, r := range cell {
		i := strings.Index(base16, string(r))
		bits <<= 1
		bits |= (i >> (2 + offset)) & 1
		bits <<= 1
		bits |= (i >> offset) & 1
	}
	return bits
}

func constrict(span [2]float64, coord float64, numbits int) int {
	bits := 0
	for i := numbits; i > 0; i-- {
		// subdivide
		m := mid(span)
		if coord > m {
			span[0] = m
			bits |= 1 << uint(i-1)
		} else {
			span[1] = m
		}
	}
	return bits
}

func refine(span *[2]float64, bit int) {
	if bit&1 != 0 {
		span[0] = mid(*span)
	} else {
		span[1] = mid(*span)
	}
}

func mid(pair [2]float64) float64 {
	return (pair[0] + pair[1]) / 2
}
