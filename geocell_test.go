// Copyright 2012 Chris Broadfoot (chris@chrisbroadfoot.id.au). All rights reserved.
// Licensed under Apache 2.
package geocell

import "testing"

const (
	lat  = 37.4
	lng  = -152.34
	cell = Cell("8b274e45b9e19")
)

var latlng = LatLng{lat, lng}

func Test_Encode_1(t *testing.T) {
	result := Encode(latlng)
	if result != cell {
		t.Error(result, cell)
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
