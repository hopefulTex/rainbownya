package main

import (
	"math"
	"strconv"
)

type variables struct {
	freq       float64
	spread     float64
	width      float64
	center     float64
	seed       float64
	forceColor bool
}

// Defaults
func newVars() variables {
	return variables{
		freq:   0.1,
		spread: 3.0,
		width:  127.5,
		center: 127.5,
		seed:   -1.0,
	}
}

func (v *variables) forceTTY() {
	v.forceColor = true
}

func (v *variables) setFrequency(in float64) {
	v.freq = in
}

func (v *variables) setSpread(in float64) {
	v.spread = in
}

func (v *variables) setWidth(in float64) {
	v.width = in
}

func (v *variables) setCenter(in float64) {
	v.center = in
}

func (v *variables) setSeed(in float64) {
	v.seed = in
}

func (v *variables) calcColor(in string, index int) string {

	var comb float64 = v.freq * (float64(index)/v.spread + v.seed)

	r := int64(math.Sin(comb)*v.width + v.center)
	g := int64(math.Sin(comb+(2*math.Pi/3.0))*v.width + v.center)
	b := int64(math.Sin(comb+(4*math.Pi/3.0))*v.width + v.center)

	return toHex(r, g, b)
}

func toHex(r, g, b int64) string {
	red := strconv.FormatInt(r, 16)
	if len(red) < 2 {
		red = "0" + red
	}
	green := strconv.FormatInt(g, 16)
	if len(green) < 2 {
		green = "0" + green
	}
	blue := strconv.FormatInt(b, 16)
	if len(blue) < 2 {
		blue = "0" + blue
	}
	return "#" + red + green + blue
}
