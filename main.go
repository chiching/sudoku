package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Grid [9][9]byte

// check if the value can be set at this coordinate point
func (g *Grid) canSet(row, col int, value byte) bool {
	for i := 0; i < 9; i++ {
		// row has the value already
		if g[row][i] == value {
			return false
		}
		// column has the value already
		if g[i][col] == value {
			return false
		}
		// sub grid has the value already
		if g[(row/3)*3+i/3][(col/3)*3+i%3] == value {
			return false
		}
	}
	return true
}

// back-track algorithm
func (g *Grid) trySetValue(row, col int) bool {
	// if the last column, move to next row
	max := 9
	if col == max {
		return g.trySetValue(row+1, 0)
	}
	// if more than max row, means finished, return.
	if row == max {
		return true
	}

	// if current point is not default value 0, move to try next point
	if g[row][col] != byte(0) {
		return g.trySetValue(row, col+1)
	}

	// try this point from 1 to 9
	for i := 1; i <= 9; i++ {
		value := byte(i)
		if !g.canSet(row, col, value) {
			continue
		}

		g[row][col] = value
		// if try finished, return.
		if g.trySetValue(row, col+1) {
			return true
		}
		// rollback.
		g[row][col] = byte(0)
	}
	// didn't find the solution.
	return false
}

// load the data from txt file to be an array.
func loadData() []*Grid {
	// read data
	f, err := os.Open("sudoku.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)
	grids := []*Grid{}
	var oneGrid *Grid
	row := 0
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		// start a new grid
		if string(line[0:4]) == "Grid" {
			oneGrid = &Grid{}
			grids = append(grids, oneGrid)
			row = 0
			continue
		}
		if row > 8 || len(line) > 9 {
			panic("data error!")
		}
		for col, v := range line {
			bv, _ := strconv.Atoi(string(v))
			if bv > 9 || bv < 0 {
				panic("data error!")
			}
			oneGrid[row][col] = byte(bv)
		}
		row++
	}
	return grids
}

func main() {
	grids := loadData()
	for idx, g := range grids {
		g.trySetValue(0, 0)
		fmt.Printf("Grid %d, The sum of top left 3-digit numbers: %d", idx+1, g[0][0]+g[0][1]+g[0][2])
		fmt.Println()
	}
}
