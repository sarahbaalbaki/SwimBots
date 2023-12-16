package main
// Pond
type Pond struct {
	preference	int
	swimbots 	[]*Swimbot
	foodBits 	[]*Food
	width    	float64
}

type OrderedPair struct {
	x float64
	y float64
}

type Swimbot struct {
	goal                             Goal
	age                              float64
	energy                           float64
	position, velocity, acceleration OrderedPair
	mass                             float64
	family                           []int
	botGene                          CommonGene
	segGenes                         []SegmentGene
	mainSegment                      *Segment
}

type Goal struct {
	isBot bool
	index int
}

type CommonGene struct {
	angularMovement       float64
	translationalMovement float64
	numSegments           int
}

type Segment struct {
	position 	OrderedPair
	angle		float64
	index       int
	subSegments []*Segment
}

type SegmentGene []float64

// 0 red, 1 blue, 2 green float64
// 3 angleToParent float64
// 4 length, 5 width float64

type Food struct {
	position OrderedPair
}
