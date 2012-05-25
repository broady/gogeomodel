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

func Encode(lat, lng float64) string {
	cell := ""
	lats := [2]float64{-90, 90}
	lngs := [2]float64{-180, 180}

	for len(cell) < precision {
		n := 0
		// interleave bits
		n |= constrict(&lats, lat) << 3
		n |= constrict(&lngs, lng) << 2
		n |= constrict(&lats, lat) << 1
		n |= constrict(&lngs, lng)
		cell += string(base16[n])
	}
	return cell
}

func Decode(cell string) ([2]float64, [2]float64) {
	lats := [2]float64{-90, 90}
	lngs := [2]float64{-180, 180}

	for _, r := range cell {
		i := strings.Index(base16, string(r))
		refine(&lats, (i>>3)&1)
		refine(&lngs, (i>>2)&1)
		refine(&lats, (i>>1)&1)
		refine(&lngs, i&1)
	}
	return lats, lngs
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
