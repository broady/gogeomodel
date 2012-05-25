// Copyright 2012 Chris Broadfoot (chris@chrisbroadfoot.id.au). All rights reserved.
// Licensed under Apache 2.
package geocell

import "testing"

func Test_Bounds_Contains_1(t *testing.T) {
	b := LatLngBox{
		South: -1,
		North: 1,
		West: 9,
		East: 11,
	}
	if !b.Contains(LatLng{0, 10}) {
		t.Error("box should contain latlng")
	}
	if b.Contains(LatLng{2, 10}) {
		t.Error("box should not contain latlng")
	}
	if b.Contains(LatLng{-2, 10}) {
		t.Error("box should not contain latlng")
	}
	if b.Contains(LatLng{0, -10}) {
		t.Error("box should not contain latlng")
	}
}
