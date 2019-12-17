package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid(t *testing.T) {
	grid := NewGrid()
	fmt.Println(grid)
	grid.Set(Point{0, 0}, true)
	grid.Set(Point{0, 1}, true)
	grid.Set(Point{0, -1}, true)
	fmt.Println(grid)

	assert.Equal(t, true, grid.Get(Point{0, 0}))
	assert.Equal(t, true, grid.Get(Point{0, 1}))
	assert.Equal(t, true, grid.Get(Point{0, -1}))
	assert.Equal(t, nil, grid.Get(Point{100, 100}))
}
