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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// The set of buttons, can change at each level
type buttonSet struct {
	content   []button
	onBeat    bool
	firstLoop bool
}

// A button has a position and a size
type button struct {
	drawX, drawY        float64
	x, y, width, height int
	hover               bool
	kind                int
	positionInSequence  int
}

// Kinds of buttons
const (
	buttonPlay int = iota
	buttonReset
	buttonSequence
)

// Initialize the button set for a given level
func (bSet *buttonSet) setupButtons(sequenceLen int) {
	buttonSet := make([]button, sequenceLen+2)

	// Play button
	buttonSet[0] = button{
		drawX: 0, drawY: float64(globalScreenHeight - globalButtonHeight),
		x: 0, y: globalScreenHeight - globalButtonHeight,
		width: globalButtonWidth, height: globalButtonHeight,
		kind: buttonPlay,
	}

	// Reset button
	buttonSet[1] = button{
		drawX: globalScreenWidth - globalButtonWidth,
		drawY: float64(globalScreenHeight - globalButtonHeight),
		x:     globalScreenWidth - globalButtonWidth,
		y:     globalScreenHeight - globalButtonHeight,
		width: globalButtonWidth, height: globalButtonHeight,
		kind: buttonReset,
	}

	// Sequence buttons
	x := (globalScreenWidth - sequenceLen*globalButtonWidth) / 2
	for pos := 0; pos < sequenceLen; pos++ {
		buttonSet[pos+2] = button{
			drawX: float64(x), drawY: float64(globalScreenHeight - globalButtonHeight),
			x: x, y: globalScreenHeight - globalButtonHeight,
			width: globalButtonWidth, height: globalButtonHeight,
			kind:               buttonSequence,
			positionInSequence: pos,
		}
		x += globalButtonWidth
	}

	bSet.content = buttonSet
}

// Record if it is beat or half beat time
func (bSet *buttonSet) setBeat() {
	bSet.onBeat = true
	bSet.firstLoop = false
}

func (bSet *buttonSet) setHalfBeat() {
	bSet.onBeat = false
}

func (bSet *buttonSet) setFirstLoop() {
	bSet.firstLoop = true
}

// Update the buttons
func (bSet *buttonSet) update(cursorX, cursorY int) (click bool, clickKind int, positionInSequence int) {

	hoveredPos := -1

	for pos, button := range bSet.content {
		bSet.content[pos].hover = cursorX >= button.x && cursorX < button.x+button.width &&
			cursorY >= button.y && cursorY < button.y+button.height
		if bSet.content[pos].hover {
			hoveredPos = pos
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && hoveredPos != -1 {
		return true, bSet.content[hoveredPos].kind, bSet.content[hoveredPos].positionInSequence
	}

	return
}

// Draw the buttons
func (buttonSet buttonSet) draw(sequence []int, currentPosition int, inPlay bool, screen *ebiten.Image) {

	for _, button := range buttonSet.content {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(button.drawX, button.drawY)

		imageNum := 0
		directionNum := nothing
		switch button.kind {
		case buttonPlay:
			imageNum = 11
			if button.hover {
				imageNum += 2
			}
		case buttonReset:
			imageNum = 12
			if button.hover {
				imageNum += 2
			}
		case buttonSequence:
			imageNum = 8
			directionNum = sequence[button.positionInSequence]
			if (!button.hover && buttonSet.onBeat && !inPlay && directionNum == nothing) ||
				(button.hover && !inPlay) ||
				(button.hover && inPlay && (currentPosition != button.positionInSequence || !buttonSet.onBeat)) {
				imageNum = 9
			} else if currentPosition == button.positionInSequence &&
				buttonSet.onBeat && inPlay && !buttonSet.firstLoop {
				imageNum = 10
				directionNum += nothing
			}
		}

		screen.DrawImage(buttonsImage.SubImage(
			image.Rect(imageNum*globalButtonWidth, 0,
				(imageNum+1)*globalButtonWidth,
				globalButtonHeight)).(*ebiten.Image),
			options)

		if button.kind == buttonSequence &&
			sequence[button.positionInSequence] != nothing {
			screen.DrawImage(buttonsImage.SubImage(
				image.Rect(directionNum*globalButtonWidth, 0,
					(directionNum+1)*globalButtonWidth,
					globalButtonHeight)).(*ebiten.Image),
				options)
		}
	}

	/*
		startY := float64(globalScreenHeight - globalButtonHeight)

		// draw play and rewind
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(0, startY)
		screen.DrawImage(buttonsImage.SubImage(
			image.Rect(10*globalButtonWidth, 0,
				11*globalButtonWidth,
				globalButtonHeight)).(*ebiten.Image),
			options)

		options = &ebiten.DrawImageOptions{}
		options.GeoM.Translate(globalScreenWidth-globalButtonWidth, startY)
		screen.DrawImage(buttonsImage.SubImage(
			image.Rect(11*globalButtonWidth, 0,
				12*globalButtonWidth,
				globalButtonHeight)).(*ebiten.Image),
			options)

		//draw the sequence buttons
		startX := float64(globalScreenWidth-len(sequence)*globalButtonWidth) / 2
		for position, direction := range sequence {
			options := &ebiten.DrawImageOptions{}
			options.GeoM.Translate(startX, startY)
			// button
			buttonImage := 8
			if position == currentPosition && drawBeat {
				buttonImage++
			}
			screen.DrawImage(buttonsImage.SubImage(
				image.Rect(buttonImage*globalButtonWidth, 0,
					(buttonImage+1)*globalButtonWidth,
					globalButtonHeight)).(*ebiten.Image),
				options)
			// move
			if direction != nothing {
				if position == currentPosition && drawBeat {
					direction += nothing
				}

				screen.DrawImage(buttonsImage.SubImage(
					image.Rect(direction*globalButtonWidth, 0,
						(direction+1)*globalButtonWidth,
						globalButtonHeight)).(*ebiten.Image),
					options)
			}
			startX += globalButtonWidth
		}
	*/

}
