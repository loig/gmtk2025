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
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

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
	levelWall
	levelEmpty
)

// A level for testing purpose (will not be used in the final game).
var testLevel level = level{
	area: [][]int{
		{levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelWall, levelWall, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelWall, levelWall, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelCeiling, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelWall, levelWall, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor},
		{levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor, levelFloor},
	},
	sequenceLen: 8,
	startX:      0, startY: 0,
	goalX: 3, goalY: 3,
}

// Draw an area on screen.
func drawLevelArea(area [][]int, startX, startY float64, screen *ebiten.Image) {

	// draw floor
	for y := range area {
		for x := range area[y] {
			if area[y][x] == levelFloor {

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
func drawGoal(x, y int, startX, startY float64, screen *ebiten.Image) {

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		startX+float64(x*globalTileSize)-globalTileMargin,
		startY+float64(y*globalTileSize)-globalTileMargin)

	subImageX := (globalTileSize + 2*globalTileMargin) * (levelEmpty + 3)

	screen.DrawImage(tilesImage.SubImage(
		image.Rect(subImageX, 0,
			subImageX+(globalTileSize+2*globalTileMargin),
			globalTileSize+2*globalTileMargin)).(*ebiten.Image),
		options)

}
