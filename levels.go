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
	_ "embed"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var levelSet []level
var levelSteps [3]int
var levelStepReset int

// A level is an area (a matrix of things such
// as floor, walls, etc), the number of moves
// in the loop for the character, a starting
// position and a goal position.
type level struct {
	area           [][]int
	sequenceLen    int // The sequence length should never be over 8
	startX, startY int
	goalX, goalY   int
}

// The type of things that can be found in a level.
const (
	levelFloor int = iota
	levelCeiling
	levelUpBox
	levelRightBox
	levelDownBox
	levelLeftBox
	levelResetBox
	levelNothingBox
	levelUp
	levelRight
	levelDown
	levelLeft
	levelReset
	levelWall
	levelEmpty
)

//go:embed levels/test
var testLevelBytes []byte

//go:embed levels/test1
var testLevel1Bytes []byte

//go:embed levels/learn
var learnLevelBytes []byte

//go:embed levels/learnautomove
var learnautomoveLevelBytes []byte

//go:embed levels/learnblock
var learnblockLevelBytes []byte

//go:embed levels/learnreset
var learnresetLevelBytes []byte

// Set up the levels
func initLevels() {

	// For tests, to be removed
	levelSet = append(levelSet, readLevel(learnblockLevelBytes))

	// First level
	levelSet = append(levelSet, readLevel(learnLevelBytes))

	// From there auto moves can be used
	levelSteps[0] = len(levelSet)
	levelSet = append(levelSet, readLevel(learnautomoveLevelBytes))

	// From there replace blocks can be used
	levelSteps[1] = len(levelSet)
	levelSet = append(levelSet, readLevel(learnblockLevelBytes))

	// From there reset can be used (must be after learning auto moves)
	levelSteps[2] = len(levelSet)
	levelSet = append(levelSet, readLevel(learnresetLevelBytes))
	levelStepReset = len(levelSet)

}

// Read a text file representing a level
func readLevel(levelBytes []byte) (l level) {
	x, y := 0, -1
	for _, b := range levelBytes {
		switch b {
		case '.':
			l.area[y] = append(l.area[y], levelFloor)
			x++
		case 's':
			l.area[y] = append(l.area[y], levelFloor)
			l.startX = x
			l.startY = y
			x++
		case 'g':
			l.area[y] = append(l.area[y], levelFloor)
			l.goalX = x
			l.goalY = y
			x++
		case '#':
			l.area[y] = append(l.area[y], levelWall)
			x++
		case '\n':
			l.area = append(l.area, make([]int, 0))
			y++
			x = 0
		case '1', '2', '3', '4', '5', '6', '7', '8':
			l.sequenceLen = int(b) - 48
		case 'u':
			l.area[y] = append(l.area[y], levelUp)
			x++
		case 'U':
			l.area[y] = append(l.area[y], levelUpBox)
			x++
		case 'd':
			l.area[y] = append(l.area[y], levelDown)
			x++
		case 'D':
			l.area[y] = append(l.area[y], levelDownBox)
			x++
		case 'l':
			l.area[y] = append(l.area[y], levelLeft)
			x++
		case 'L':
			l.area[y] = append(l.area[y], levelLeftBox)
			x++
		case 'r':
			l.area[y] = append(l.area[y], levelRight)
			x++
		case 'R':
			l.area[y] = append(l.area[y], levelRightBox)
			x++
		case 'b':
			l.area[y] = append(l.area[y], levelReset)
			x++
		case 'B':
			l.area[y] = append(l.area[y], levelResetBox)
			x++
		case 'n':
			l.area[y] = append(l.area[y], levelNothingBox)
			x++
		default:
			l.area[y] = append(l.area[y], levelEmpty)
			x++
		}
	}

	simplifyLevelArea(l.area)

	return l
}

// Set up a level for better display
func simplifyLevelArea(area [][]int) {

	// Remove floor around outer walls
	for y := 0; y < len(area); y++ {
		reached := 0
		for x := 0; x < len(area[y]) && area[y][x] != levelWall; x++ {
			area[y][x] = levelEmpty
			reached = x
		}
		for x := len(area[y]) - 1; x > reached && area[y][x] != levelWall; x-- {
			area[y][x] = levelEmpty
		}
	}

	// Put ceiling when needed
	for y := 0; y < len(area)-1; y++ {
		for x := 0; x < len(area[y]); x++ {
			if area[y][x] == levelWall && area[y+1][x] == levelWall {
				area[y][x] = levelCeiling
			}
		}
	}
}

// Draw an area on screen.
func drawLevelArea(area [][]int, startX, startY float64, screen *ebiten.Image) {

	// draw floor
	for y := range area {
		for x := range area[y] {
			if area[y][x] != levelWall &&
				area[y][x] != levelCeiling &&
				area[y][x] != levelEmpty {

				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(
					startX+float64(x*globalTileSize)-globalTileMargin,
					startY+float64(y*globalTileSize)-globalTileMargin)

				subImageX := 0

				if (x+y)%2 == 0 {
					subImageX = (globalTileSize + 2*globalTileMargin)
				}

				screen.DrawImage(tilesImage.SubImage(
					image.Rect(subImageX, 0,
						subImageX+(globalTileSize+2*globalTileMargin),
						globalTileSize+2*globalTileMargin)).(*ebiten.Image),
					options)
			}
		}
	}

	// draw walls
	for y := range area {
		for x := range area[y] {
			if area[y][x] == levelWall {

				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(
					startX+float64(x*globalTileSize)-globalTileMargin,
					startY+float64(y*globalTileSize)-globalTileMargin)

				subImageX := (globalTileSize + 2*globalTileMargin) * (levelWall + 2)

				screen.DrawImage(tilesImage.SubImage(
					image.Rect(subImageX, 0,
						subImageX+(globalTileSize+2*globalTileMargin),
						globalTileSize+2*globalTileMargin)).(*ebiten.Image),
					options)
			}
		}
	}

	// draw other things
	for y := range area {
		for x := range area[y] {
			if area[y][x] != levelFloor && area[y][x] != levelEmpty {

				options := &ebiten.DrawImageOptions{}
				options.GeoM.Translate(
					startX+float64(x*globalTileSize)-globalTileMargin,
					startY+float64(y*globalTileSize)-globalTileMargin)

				subImageX := (globalTileSize + 2*globalTileMargin) * (area[y][x] + 1)

				screen.DrawImage(tilesImage.SubImage(
					image.Rect(subImageX, 0,
						subImageX+(globalTileSize+2*globalTileMargin),
						globalTileSize+2*globalTileMargin)).(*ebiten.Image),
					options)
			}
		}
	}

}

// Draw a goal on screen.
func drawGoal(x, y int, startX, startY float64, onBeat bool, screen *ebiten.Image) {

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		startX+float64(x*globalTileSize)-globalTileMargin,
		startY+float64(y*globalTileSize)-globalTileMargin)

	increment := 4
	if !onBeat {
		increment++
	}
	subImageX := (globalTileSize + 2*globalTileMargin) * (levelEmpty + increment)

	screen.DrawImage(tilesImage.SubImage(
		image.Rect(subImageX, 0,
			subImageX+(globalTileSize+2*globalTileMargin),
			globalTileSize+2*globalTileMargin)).(*ebiten.Image),
		options)

}
