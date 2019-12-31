package imagework

import (
	"image"
	"image/color"
)

const (
	maxBuffForX     = 6
	maxBuffForY     = 5
	alphaRoundCodes = 128
)

func coordOfCode(img image.Image) (int, int, int, int) {

	imar := imgInArr(img)
	x1, x2, y1, y2 := analiseImar(imar)

	return x1, x2, y1, y2
}

func imgInArr(img image.Image) [][]int {

	dar := make([][]int, img.Bounds().Max.Y)

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		dar[y] = make([]int, img.Bounds().Max.X)
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			if img.At(x, y).(color.NRGBA).A == alphaRoundCodes {
				dar[y][x] = 1
			} else {
				dar[y][x] = 0
			}
		}
	}
	return dar
}

func analiseImar(imar [][]int) (int, int, int, int) {

	var (
		buff int
		x1s  []int
		x2s  []int
		y1s  []int
		y2s  []int
	)

	for y := 0; y < len(imar); y++ {
		for x := 0; x < len(imar[y]); x++ {
			x1s, x2s, buff = findExtremus(x1s, x2s, imar[y][x], x, buff, maxBuffForX)
		}
		if buff >= maxBuffForX {
			x2s = append(x2s, len(imar[y])-1)
		}
		buff = 0
	}

	for x := 0; x < len(imar[0]); x++ {
		for y := 0; y < len(imar); y++ {
			y1s, y2s, buff = findExtremus(y1s, y2s, imar[y][x], y, buff, maxBuffForY)
		}
		if buff >= maxBuffForY {
			y2s = append(y2s, len(imar)-1)
		}
		buff = 0
	}

	minx1, maxx2 := minmax(x1s, x2s)
	miny1, maxy2 := minmax(y1s, y2s)

	return minx1, miny1, maxx2, maxy2
}

func findExtremus(x1s, x2s []int, value, index, buff, maxBuff int) ([]int, []int, int) {

	if value == 1 {
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
