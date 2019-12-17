package main

// Point ...
type Point struct {
	X, Y int
}

// Grid ...
type Grid struct {
	Values map[int]map[int]interface{}
}

// NewGrid ...
func NewGrid() *Grid {
	return &Grid{
		Values: make(map[int]map[int]interface{}),
	}

}

// Get ...
func (g *Grid) Get(p Point) interface{} {
	return g.Values[p.X][p.Y]
}

// Set ...
func (g *Grid) Set(p Point, v interface{}) {
	if _, ok := g.Values[p.X]; !ok {
		g.Values[p.X] = make(map[int]interface{})
	}
	g.Values[p.X][p.Y] = v
}
