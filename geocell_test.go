// Copyright 2012 Chris Broadfoot (chris@chrisbroadfoot.id.au). All rights reserved.
// Licensed under Apache 2.
package geocell

import (
	"testing"
	"fmt"
)

const (
	lat  = 37.4
	lng  = -152.34
	cell = Cell("8b274e45b9e19")
)

var latlng = LatLng{lat, lng}

func Test_Encode_1(t *testing.T) {
	result := Encode(latlng)
	if result != cell {
		t.Error(cell, fmt.Sprintf("%b", cell.latbits()), fmt.Sprintf("%b", cell.lngbits()))
		t.Error(result, fmt.Sprintf("%b", result.latbits()), fmt.Sprintf("%b", result.lngbits()))
	}
}
func Test_Decode_1(t *testing.T) {
	box := cell.Decode()
	if !box.Contains(latlng) {
		t.Error("latlng not in decoded box")
	}
}

func Test_Decode_Bad_1(t *testing.T) {
	bad := []byte(cell)
	bad[3] = 'e'
	box := Cell(bad).Decode()
	if box.Contains(latlng) {
		t.Error("latlng in decoded box")
	}
}

func Test_Decode_Bad_2(t *testing.T) {
	bad := []byte(cell)
	bad[5] = 'f'
	box := Cell(bad).Decode()
	if box.Contains(latlng) {
		t.Error("latlng in decoded box")
	}
}

func Test_East_1(t *testing.T) {
	if Cell("8").East() != Cell("9") {
		t.Error("8->9")
	}
	if Cell("9").East() != Cell("c") {
		t.Error("9->c")
	}
	if Cell("c").East() != Cell("d") {
		t.Error("c->d")
	}
}

// test wrapping at different levels
func Test_East_2(t *testing.T) {
	if Cell("7").East() != Cell("2") {
		t.Error("7->2")
	}
	if Cell("57df").East() != Cell("028a") {
		t.Error("57df->028a")
	}
}

func Test_LatBits_1(t *testing.T) {
	if Cell("fc30").latbits() != 228 {
		t.Error("f3c0")
	}
	if Cell("fc30").latbits() != Cell("fd75").latbits() {
		t.Error("f3c0==fd75")
	}
	if Cell("fd75").latbits() != Cell("b931").latbits() {
		t.Error("fd75==b931")
	}
}

func Test_LngBits_1(t *testing.T) {
	if Cell("fc30").lngbits() != 228 {
		t.Error(Cell("f3c0").latbits())
	}
	if Cell("fc30").lngbits() != Cell("feba").lngbits() {
		t.Error("f3c0==feba")
	}
	if Cell("feba").lngbits() != Cell("5410").lngbits() {
		t.Error("feba==5410")
	}
}

func Test_FromBits_1(t *testing.T) {
	cell := Cell("0b274f45b9e10")
	lats := cell.latbits()
	lngs := cell.lngbits()

	if fromBits(lats, lngs, len(cell)) != cell {
		t.Error("fail", lats, lngs, fromBits(lats, lngs, len(cell)))
	}
}
