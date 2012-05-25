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

type Cell string

func Encode(latlng LatLng) Cell {
	latbits := constrict([2]float64{90, -90}, latlng.Lat, precision*2)
	lngbits := constrict([2]float64{180, -180}, latlng.Lng, precision*2)
	return fromBits(latbits, lngbits, precision)
}

func constrict(span [2]float64, coord float64, numbits int) int {
	bits := 0
	for i := uint(numbits); i > 0; i-- {
		// subdivide
		m := mid(span)
		if coord < m {
			span[0] = m
		} else {
			span[1] = m
			bits |= 1 << (i - 1)
		}
	}
	return bits
}

func (cell Cell) Decode() LatLngBox {
	lats := refine([2]float64{90, -90}, cell.latbits(), cell.Precision()*2)
	lngs := refine([2]float64{180, -180}, cell.lngbits(), cell.Precision()*2)
	return LatLngBox{
		South: lats[1],
		North: lats[0],
		West:  lngs[1],
		East:  lngs[0],
	}
}

func refine(span [2]float64, bits int, numbits int) [2]float64 {
	for i := uint(numbits); i > 0; i-- {
		span[getbit(bits, i-1)] = mid(span)
	}
	return span
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

func (cell Cell) North() Cell {
	// TODO: avoid wrapping vertically
	return fromBits(cell.latbits()+1, cell.lngbits(), cell.Precision())
}

func (cell Cell) South() Cell {
	// TODO: avoid wrapping vertically
	return fromBits(cell.latbits()-1, cell.lngbits(), cell.Precision())
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

func mid(pair [2]float64) float64 {
	return (pair[0] + pair[1]) / 2
}

// getbit returns 0 or 1 for a given bit position of an int.
func getbit(b int, pos uint) int {
	return (b >> pos) & 1
}
