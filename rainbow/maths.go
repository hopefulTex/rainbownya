package rainbow

import (
	"math"
	"strconv"
)

type Variables struct {
	Freq   float64
	Spread float64
	Width  float64
	Center float64
	Seed   float64
}

func DefaultVars() Variables {
	return Variables{
		Freq:   0.1,
		Spread: 3.0,
		Width:  127.5,
		Center: 127.5,
		Seed:   -1.0,
	}
}

func (v *Variables) calcColor(index int) string {

	var comb float64 = v.Freq * (float64(index)/v.Spread + v.Seed)

	r := int64(math.Sin(comb)*v.Width + v.Center)
	g := int64(math.Sin(comb+(2*math.Pi/3.0))*v.Width + v.Center)
	b := int64(math.Sin(comb+(4*math.Pi/3.0))*v.Width + v.Center)

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
