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

// The character position, sequence of moves (that will
// be played in loop), and the current level (an array
// of things that can be floor, walls, etc).
type character struct {
	x, y                   int
	moveSequence           []int
	nextMovePosition       int
	currentMovePosition    int
	levelArea              [][]int
	levelGoalX, levelGoalY int
	displayX, displayY     float64
	onBeat                 bool
}

// The possible moves of the character.
const (
	moveUp int = iota
	moveRight
	moveDown
	moveLeft
	moveReset
	nothing
)

// Reset a given level by emptying the sequence
// of moves, moving the character to the start,
// copying the area (in case consumables were
// used).
func (c *character) reset(level level) {
	c.x = level.startX
	c.y = level.startY
	c.moveSequence = make([]int, level.sequenceLen)
	for pos := 0; pos < len(c.moveSequence); pos++ {
		c.moveSequence[pos] = nothing
	}
	c.nextMovePosition = 0
	c.currentMovePosition = 0
	c.levelArea = make([][]int, len(level.area))
	for linePos, line := range level.area {
		c.levelArea[linePos] = make([]int, len(line))
		copy(c.levelArea[linePos], line)
	}
	c.levelGoalX = level.goalX
	c.levelGoalY = level.goalY
	c.displayY = float64(globalScreenHeight-globalButtonHeight-len(level.area)*globalTileSize) / 2
	if len(level.area) > 0 {
		c.displayX = float64(globalScreenWidth-len(level.area[0])*globalTileSize) / 2
	}
}

// The character performs one step of its
// sequence of moves at each beat. If the
// step is not "do nothing" then a sound is
// played on the beat.
func (c *character) updateOnBeat() (playSound bool, soundID int) {
	if c.applyMove(c.moveSequence[c.nextMovePosition]) {
		playSound, soundID = getMoveSoundId(c.moveSequence[c.nextMovePosition])
	} else {
		playSound = true
		soundID = soundBlip
	}
	c.currentMovePosition = c.nextMovePosition
	if c.moveSequence[c.nextMovePosition] == moveReset {
		c.nextMovePosition = 0
	} else {
		c.nextMovePosition = (c.nextMovePosition + 1) % len(c.moveSequence)
	}
	return
}

func (c *character) setBeat() {
	c.onBeat = true
}

// Get the effect of a given move on the character
// depending of the area and the character current
// position.
func (c *character) applyMove(move int) (success bool) {

	xTo := c.x
	yTo := c.y
	switch move {
	case moveUp:
		yTo--
	case moveRight:
		xTo++
	case moveDown:
		yTo++
	case moveLeft:
		xTo--
	}

	success = c.isAccessible(xTo, yTo)

	if success {
		c.x = xTo
		c.y = yTo
	}

	return
}

// Check if a given position in the area is
// suitable for the character to stay on.
func (c character) isAccessible(x, y int) bool {
	return x >= 0 && y >= 0 &&
		y < len(c.levelArea) && x < len(c.levelArea[y]) &&
		c.levelArea[y][x] != levelWall &&
		c.levelArea[y][x] != levelCeiling
}

// Given a move, get the corresponding sound ID.
func getMoveSoundId(move int) (playSound bool, soundID int) {
	if move == nothing {
		return false, 0
	}

	switch move {
	case moveUp:
		soundID = soundC3
	case moveRight:
		soundID = soundE3
	case moveDown:
		soundID = soundG3
	case moveLeft:
		soundID = soundC4
	case moveReset:
		soundID = soundC2
	}

	return true, soundID
}

// On each half beat consumables are consumed
// and their effects are applied. This produces
// a sound on the half beat.
func (c *character) updateOnHalfBeat() (playSound bool, soundID int) {

	effect := c.levelArea[c.y][c.x]

	switch effect {
	case levelUp, levelLeft, levelDown, levelRight:
		if c.applyMove(c.levelArea[c.y][c.x] - levelUp) {
			playSound, soundID = true, soundG4
		}
	case levelUpBox, levelLeftBox, levelDownBox, levelRightBox, levelResetBox, levelNothingBox:
		c.levelArea[c.y][c.x], c.moveSequence[c.currentMovePosition] =
			c.moveSequence[c.currentMovePosition]+levelUpBox,
			c.levelArea[c.y][c.x]-levelUpBox
	}

	return
}

func (c *character) setHalfBeat() {
	c.onBeat = false
}

// If the character has reached the goal position
// the level is complete.
func (c character) checkGoal() bool {
	return c.x == c.levelGoalX && c.y == c.levelGoalY
}

// Draw the character and the area on screen.
func (c character) draw(screen *ebiten.Image) {

	drawLevelArea(c.levelArea, c.displayX, c.displayY, screen)

	drawGoal(c.levelGoalX, c.levelGoalY, c.displayX, c.displayY, c.onBeat, screen)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		c.displayX+float64(c.x*globalTileSize)-globalTileMargin,
		c.displayY+float64(c.y*globalTileSize)-globalTileMargin)

	increment := 2
	if !c.onBeat {
		increment++
	}
	subImageX := (globalTileSize + 2*globalTileMargin) * (levelEmpty + increment)

	screen.DrawImage(tilesImage.SubImage(
		image.Rect(subImageX, 0,
			subImageX+(globalTileSize+2*globalTileMargin),
			globalTileSize+2*globalTileMargin)).(*ebiten.Image),
		options)
}
