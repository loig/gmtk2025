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
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Visual effect when switching moves with box on the floor
type boxSwitcher struct {
	floor, move          int
	floorBoxX, floorBoxY float64
	moveBoxX, moveBoxY   float64
	frame, numFrames     int
	dx, dy               float64
}

func (b *boxSwitcher) reset() {
	b.numFrames = 0
}

func (b *boxSwitcher) setUp(
	charX, charY int, displayX, displayY float64,
	numMoves int, switchMoveNum int,
	bpm int, floorMove, seqMove int) {

	b.floorBoxX = displayX + float64(charX*globalTileSize) + 6
	b.floorBoxY = displayY + float64(charY*globalTileSize) + 16

	b.moveBoxX = float64(globalScreenWidth-numMoves*globalButtonWidth)/2 +
		float64(switchMoveNum*globalButtonWidth) +
		globalButtonWidth/2 - globalTileSize/2
	b.moveBoxY = globalScreenHeight - 3*globalButtonHeight/4

	b.frame = 0
	b.numFrames = 3600 / (2 * bpm)

	distanceX := b.floorBoxX - b.moveBoxX
	distanceY := b.floorBoxY - b.moveBoxY

	b.dx = distanceX / float64(b.numFrames)
	b.dy = distanceY / float64(b.numFrames)

	b.floor = seqMove
	b.move = floorMove

	log.Printf("Set up with floor %d and player %d", b.floor, b.move)

}

func (b *boxSwitcher) update() {
	if b.frame < b.numFrames {
		b.frame++
		b.moveBoxX += b.dx
		b.moveBoxY += b.dy
		b.floorBoxX -= b.dx
		b.floorBoxY -= b.dy
	}
}

func (b boxSwitcher) draw(screen *ebiten.Image) {

	if b.frame < b.numFrames {

		// Box coming from floor
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(b.floorBoxX-16, b.floorBoxY-6)
		imageNum := b.floor + levelUpBox + 1

		screen.DrawImage(tilesImage.SubImage(
			image.Rect(imageNum*(globalTileSize+2*globalTileMargin), 0,
				(imageNum+1)*(globalTileSize+2*globalTileMargin),
				globalTileSize+2*globalTileMargin)).(*ebiten.Image),
			options)

		// Box coming from move sequence
		options = &ebiten.DrawImageOptions{}
		options.GeoM.Translate(b.moveBoxX-16, b.moveBoxY-6)
		imageNum = b.move + levelUpBox + 1

		screen.DrawImage(tilesImage.SubImage(
			image.Rect(imageNum*(globalTileSize+2*globalTileMargin), 0,
				(imageNum+1)*(globalTileSize+2*globalTileMargin),
				globalTileSize+2*globalTileMargin)).(*ebiten.Image),
			options)
	}
}
