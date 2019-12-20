package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
)

const imagePathFormat = "words_4_ordered/%d.jpg"

func viewer() {
	window := gocv.NewWindow("Hello")

	images := window.CreateTrackbar("Image", 184)
	images.SetPos(1)
	cannyT1TB := window.CreateTrackbar("Canny T1", 1000)
	cannyT1TB.SetPos(30)
	houghRhoTB := window.CreateTrackbar("Hough_Rho", 100)
	houghRhoTB.SetPos(1)
	houghTresholdTB := window.CreateTrackbar("Hough_Treshold", 100)
	houghTresholdTB.SetPos(6)
	houghMinLineLengthTB := window.CreateTrackbar("Hough_MinLineLength", 1000)
	houghMinLineLengthTB.SetPos(200)
	houghMaxLineGapTB := window.CreateTrackbar("Hough_MaxLineGap", 500)
	houghMaxLineGapTB.SetPos(6)

	offsetTB := window.CreateTrackbar("Offset", 40)
	offsetTB.SetPos(1)

	t := &Test{}

	for {
		cannyT1 := float32(cannyT1TB.GetPos())
		iamge := int(images.GetPos())
		houghRho := float32(houghRhoTB.GetPos())
		houghTreshold := houghTresholdTB.GetPos()
		houghMinLineLength := float32(houghMinLineLengthTB.GetPos())
		houghMaxLineGap := float32(houghMaxLineGapTB.GetPos())
		offset := int(offsetTB.GetPos())

		if t.Image != iamge || t.Offset != offset || t.Canny_T1 != cannyT1 || t.Hough_Rho != houghRho || t.Hough_Treshold != houghTreshold || t.Hough_MinLineLength != houghMinLineLength || t.Hough_MaxLineGap != houghMaxLineGap {
			t.Canny_T1 = cannyT1
			t.Hough_Rho = houghRho
			t.Hough_Treshold = houghTreshold
			t.Hough_MinLineLength = houghMinLineLength
			t.Hough_MaxLineGap = houghMaxLineGap
			t.Hough_MaxLineGap = houghMaxLineGap
			t.Image = iamge
			t.Offset = offset

			window.IMShow(t.update())
		}

		if window.WaitKey(1) == 27 {
			fmt.Println("update")
		}
	}
}

type Test struct {
	Canny_T1            float32
	Canny_T2            float32
	Hough_Rho           float32
	Hough_Treshold      int
	Hough_MinLineLength float32
	Hough_MaxLineGap    float32
	Image               int
	Offset              int
}

func (t *Test) update() gocv.Mat {
	fmt.Println(fmt.Printf("%+v", t))
	edges := gocv.NewMat()
	lines := gocv.NewMat()
	img := gocv.IMRead(fmt.Sprintf(imagePathFormat, t.Image), gocv.IMReadColor)
	gocv.Canny(img, &edges, t.Canny_T1, t.Canny_T1*3)
	gocv.HoughLinesPWithParams(edges, &lines, t.Hough_Rho, math.Pi/180.0, t.Hough_Treshold, t.Hough_MinLineLength, t.Hough_MaxLineGap)
	r, g, b := rainbow(50, 30)
	for j := 0; j < lines.Rows(); j++ {
		vec := lines.GetVeciAt(j, 0)
		gocv.Line(&img, image.Point{X: int(vec[0]), Y: int(vec[1])}, image.Point{X: int(vec[2]), Y: int(vec[3])}, color.RGBA{uint8(r), uint8(g), uint8(b), 0}, 3)
	}
	return img
}

func rainbow(numOfSteps, step float64) (int, int, int) {

	var r, g, b float64

	h := step / numOfSteps
	i := math.Floor(h * 6)
	f := h*6 - i
	q := 1 - f

	os := math.Remainder(i, 6)
	//fmt.Println(os, h, i, f, q)

	switch os {
	case 0:
		r = 1
		g = f
		b = 0
	case 1:
		r = q
		g = 1
		b = 0
	case 2:
		r = 0
		g = 1
		b = f
	case 3:
		r = 0
		g = q
		b = 1
	case 4:
		r = f
		g = 0
		b = 1
	case 5:
		r = 1
		g = 0
		b = q
	}
	r = r * 255
	g = g * 255
	b = b * 255

	return int(r), int(g), int(b)
}
