package main

import (
	"canvas"
	"image"

)

// AnimateSystem takes a slice of Universe objects along with a canvas width
// parameter and a frequency parameter.
// Every frequency steps, it generates a slice of images corresponding to drawing each Universe
// on a canvasWidth x canvasWidth canvas.
// A scaling factor is a final input that is used to scale the stars big enough to see them.
func AnimateSystem(timePoints []*Pond, canvasWidth, frequency int, scalingFactor float64) []image.Image {
	images := make([]image.Image, 0)

	if len(timePoints) == 0 {
		panic("Error: no Pond objects present in AnimateSystem.")
	}

	// for every universe, draw to canvas and grab the image
	for i := range timePoints {
		if i%frequency == 0 {
			images = append(images, timePoints[i].DrawToCanvas(canvasWidth, scalingFactor))
		}
	}

	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing a Universe
// object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels.
// A scaling factor is needed to make the stars big enough to see them.
func (p *Pond) DrawToCanvas(canvasWidth int, scalingFactor float64) image.Image {
	if p == nil {
		panic("Can't Draw a nil pond.")
	}

	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// range over all the bodies and draw them.
	for _, b := range p.swimbots {
		if b != nil {
			var sliceOfSegments []*Segment
			sliceOfSegments = b.mainSegment.RecursiveFindSegment(sliceOfSegments)
			for i := range sliceOfSegments {
				index := sliceOfSegments[i].index
				red := uint8(b.segGenes[index][0])
				green := uint8(b.segGenes[index][1])
				blue := uint8(b.segGenes[index][2])
				c.SetFillColor(canvas.MakeColor(red, green, blue))
				length := b.segGenes[index][4]
				width := b.segGenes[index][5]
				// segGenes[i][4] = (rand.Float64() * 15.0) + 5.0 // should take on values from 5.0 to 20.0
				// // width
				// segGenes[i][5] = (rand.Float64() * 3.7) + 0.3 // should take on values from 0.3 to 4.0
				cx := (sliceOfSegments[i].position.x / p.width) * float64(canvasWidth)
				cy := (sliceOfSegments[i].position.y / p.width) * float64(canvasWidth)

				// r := scalingFactor * (1 / p.width) * float64(canvasWidth)
				c.Segment(cx, cy, width, length, sliceOfSegments[i].angle)
				c.Fill()
			}
		}
	}

	// range over all the bodies and draw them.
	for _, f := range p.foodBits {
		if f != nil {
			c.SetFillColor(canvas.MakeColor(255, 255, 255))
			cx := (f.position.x / p.width) * float64(canvasWidth)
			cy := (f.position.y / p.width) * float64(canvasWidth)
			r := scalingFactor * (1 / p.width) * float64(canvasWidth)
			c.Circle(cx, cy, r)
			c.Fill()
		}
	}
	// we want to return an image!
	return c.GetImage()
}

func (currSegment *Segment) RecursiveFindSegment(sliceOfSegments []*Segment) []*Segment {
    if currSegment.subSegments == nil {
		sliceOfSegments = append(sliceOfSegments, currSegment)
        return sliceOfSegments
    } else {
        // if we haven't reach the end segments
		sliceOfSegments = append(sliceOfSegments, currSegment)
		for i := range currSegment.subSegments {
			sliceOfSegments = append(sliceOfSegments, currSegment.subSegments[i])
		}
    }
    return sliceOfSegments
}
