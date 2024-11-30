package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	scales := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		scales[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			scales[i][j] = uint8(j)
		}
	}
	return scales
}

func main() {
	pic.Show(Pic)
}