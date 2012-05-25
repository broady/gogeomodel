// Copyright 2012 Chris Broadfoot (chris@chrisbroadfoot.id.au). All rights reserved.
// Licensed under Apache 2.
package geocell

import "testing"

const (
	lat  = 37.4
	lng  = -152.34
	cell = "8b274e45b9e19"
)

func Test_Encode_1(t *testing.T) {
	result := Encode(lat, lng)
	if result != cell {
		t.Error(result, cell)
	}
}
func Test_Decode_1(t *testing.T) {
	lats, lngs := Decode(cell)
	if lat < lats[0] || lat > lats[1] {
		t.Error("lat out of range", lat, lats)
	}
	if lng < lngs[0] || lng > lngs[1] {
		t.Error("lng out of range", lng, lngs)
	}
}

func Test_Decode_2(t *testing.T) {
	bad := []byte(cell)
	bad[3] = 'e'
	lats, lngs := Decode(string(bad))
	if lat > lats[0] && lat < lats[1] {
		t.Error("lat out of range", lat, lats)
	}
	if lng > lngs[0] && lng < lngs[1] {
		t.Error("lng out of range", lng, lngs)
	}
}
