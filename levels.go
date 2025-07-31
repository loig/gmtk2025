/*
CUB 2: Origins, a game for GMTK Game Jam 2025
Copyright (C) 2025 Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// A level is an area (a matrix of things such
// as floor, walls, etc), the number of moves
// in the loop for the character, a starting
// position and a goal position.
type level struct {
	area           [][]int
	sequenceLen    int
	startX, startY int
	goalX, goalY   int
}

// The type of things that can be found in a level.
const (
	levelFloor int = iota
	levelWall
)

// A level for testing purpose (will not be used in the final game).
var testLevel level = level{
	area: [][]int{
		{levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelWall, levelWall, levelFloor},
		{levelFloor, levelFloor, levelFloor, levelFloor},
	},
	sequenceLen: 8,
	startX:      0, startY: 0,
	goalX: 3, goalY: 3,
}

// Draw an area on screen.
func drawLevelArea(area [][]int, screen *ebiten.Image) {

	for y := range area {
		for x := range area[y] {
			odd := (x+y)%2 == 0
			floorColor := color.RGBA{G: 255, A: 255}
			if odd {
				floorColor = color.RGBA{G: 120, A: 255}
			}
			switch area[y][x] {
			case levelFloor:
				vector.DrawFilledRect(screen, float32(x*globalTileSize), float32(y*globalTileSize), globalTileSize, globalTileSize, floorColor, false)
			case levelWall:
				vector.DrawFilledRect(screen, float32(x*globalTileSize), float32(y*globalTileSize), globalTileSize, globalTileSize, color.Black, false)
			}
		}
	}

}

// Draw a goal on screen.
func drawGoal(x, y int, screen *ebiten.Image) {

	vector.DrawFilledRect(screen, float32(x*globalTileSize), float32(y*globalTileSize), globalTileSize, globalTileSize, color.RGBA{R: 255, A: 255}, false)

}
