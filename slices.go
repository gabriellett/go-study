package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for x := 0; x<dx; x++ {
		var line []uint8 
		for y := 0; y < dy; y++ {
			line = append(line, uint8((x+y)/2))
		}
		pic = append(pic, line)
	}
	
	return pic
}

func main() {
	pic.Show(Pic)
}
