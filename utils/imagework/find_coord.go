package imagework

import (
	"image"
	"image/color"
)

const (
	maxBuffForX     = 3
	maxBuffForY     = 5
	alphaRoundCodes = 128
)

func coordOfCode(img image.Image) (int, int, int, int) {

	max := img.Bounds().Max

	var (
		buff int
		x1s  []int
		x2s  []int
		y1s  []int
		y2s  []int
	)

	for y := 0; y < max.Y; y++ {
		for x := 0; x < max.X; x++ {
			x1s, x2s, buff = findExtremus(x1s, x2s, img.At(x, y), x, buff, maxBuffForX)
		}
		if buff >= maxBuffForX {
			x2s = append(x2s, max.X-1)
		}
		buff = 0
	}

	for x := 0; x < max.X; x++ {
		for y := 0; y < max.Y; y++ {
			y1s, y2s, buff = findExtremus(y1s, y2s, img.At(x, y), y, buff, maxBuffForY)
		}
		if buff >= maxBuffForY {
			y2s = append(y2s, max.Y-1)
		}
		buff = 0
	}

	minx1, maxx2 := minmax(x1s, x2s)
	miny1, maxy2 := minmax(y1s, y2s)

	return minx1, miny1, maxx2, maxy2
}

func findExtremus(x1s, x2s []int, value color.Color, index, buff, maxBuff int) ([]int, []int, int) {

	if value.(color.NRGBA).A == alphaRoundCodes {
		buff++
		if buff == maxBuff {
			x1s = append(x1s, index-(maxBuff-1))
		}
	} else {
		if buff >= maxBuff {
			x2s = append(x2s, index-1)
		}
		buff = 0
	}

	return x1s, x2s, buff
}

func minmax(minOf []int, maxOf []int) (int, int) {

	min := minOf[0]
	max := maxOf[0]
	for i := 1; i < len(minOf); i++ {
		if min > minOf[i] {
			min = minOf[i]
		}
		if max < maxOf[i] {
			max = maxOf[i]
		}
	}
	return min, max
}
