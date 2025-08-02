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
	content        []button
	onBeat         bool
	firstLoop      bool
	activePosition int
	hasActive      bool
}

// A button has a position and a size
type button struct {
	drawX, drawY        float64
	x, y, width, height int
	hover               bool
	kind                int
	positionInSequence  int
	smallPosition       int
	smallReset          bool
}

// Kinds of buttons
const (
	buttonPlay int = iota
	buttonReset
	buttonSequence
	buttonSelectMove
	buttonIncBPM
	buttonDecBPM
	buttonToggleSound
)

// Add small move buttons to a set
func (bSet *buttonSet) addButtons(withReset bool) {

	numButtons := 4
	if withReset {
		numButtons++
	}

	buttonWidth := 28
	buttonHeight := 38 + 3 // 3 is the shift when hovering
	buttonSep := 4

	x := bSet.content[bSet.activePosition].x - 6 + (globalButtonWidth-numButtons*(buttonWidth+buttonSep))/2
	y := bSet.content[bSet.activePosition].y - 6 - buttonHeight
	position := bSet.content[bSet.activePosition].positionInSequence

	for num := 0; num < numButtons; num++ {
		bSet.content = append(bSet.content, button{
			drawX: float64(x), drawY: float64(y),
			x: x, y: y - globalTileMargin/4,
			width: buttonWidth, height: buttonHeight,
			kind:               buttonSelectMove,
			positionInSequence: position,
			smallPosition:      num,
			smallReset:         withReset,
		})
		x += buttonWidth + buttonSep
	}

}

// Remove small move buttons from a set
func (bSet *buttonSet) removeButtons() {

	for bSet.content[len(bSet.content)-1].kind == buttonSelectMove {
		bSet.content = bSet.content[:len(bSet.content)-1]
	}

}

// Initialize the button set for a given level
func (bSet *buttonSet) setupButtons(sequenceLen int) {
	buttonSet := make([]button, sequenceLen+5, sequenceLen+10)

	// Play button
	buttonSet[0] = button{
		drawX: 0, drawY: float64(globalScreenHeight - globalButtonHeight),
		x: 8, y: globalScreenHeight - globalButtonHeight + 37,
		width: 65, height: 72,
		kind: buttonPlay,
	}

	// Reset button
	buttonSet[1] = button{
		drawX: globalScreenWidth - globalButtonWidth,
		drawY: float64(globalScreenHeight - globalButtonHeight),
		x:     globalScreenWidth - globalButtonWidth + 8,
		y:     globalScreenHeight - globalButtonHeight + 37,
		width: 65, height: 72,
		kind: buttonReset,
	}

	// Sequencer control buttons
	smallButtonX := 724
	buttonSet[2] = button{
		drawX: float64(smallButtonX), drawY: 50,
		x: smallButtonX, y: 50,
		width: globalSmallButtonWidth, height: globalSmallButtonHeight,
		kind: buttonDecBPM,
	}
	smallButtonX += 35
	buttonSet[3] = button{
		drawX: float64(smallButtonX), drawY: 50,
		x: smallButtonX, y: 50,
		width: globalSmallButtonWidth, height: globalSmallButtonHeight,
		kind: buttonIncBPM,
	}

	// Toggle sound
	buttonSet[4] = button{
		drawX: 10, drawY: 50,
		x: 10, y: 50,
		width: globalSmallButtonWidth, height: globalSmallButtonHeight,
		kind: buttonToggleSound,
	}

	// Sequence buttons
	x := (globalScreenWidth - sequenceLen*globalButtonWidth) / 2
	for pos := 0; pos < sequenceLen; pos++ {
		buttonSet[pos+5] = button{
			drawX: float64(x), drawY: float64(globalScreenHeight - globalButtonHeight),
			x: x + 6, y: globalScreenHeight - globalButtonHeight + 21,
			width: 68, height: 93,
			kind:               buttonSequence,
			positionInSequence: pos,
		}
		x += globalButtonWidth
	}

	bSet.content = buttonSet
	bSet.hasActive = false
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
func (bSet *buttonSet) update(cursorX, cursorY int, inSetUp bool, withReset bool) (click bool, clickKind int, positionInSequence int, smallPosition int) {

	hoveredPos := -1

	for pos, button := range bSet.content {
		bSet.content[pos].hover = cursorX >= button.x && cursorX < button.x+button.width &&
			cursorY >= button.y && cursorY < button.y+button.height
		if bSet.content[pos].hover {
			hoveredPos = pos
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && hoveredPos != -1 {

		click = true
		clickKind = bSet.content[hoveredPos].kind
		positionInSequence = bSet.content[hoveredPos].positionInSequence
		smallPosition = bSet.content[hoveredPos].smallPosition

		if inSetUp {
			if bSet.content[hoveredPos].kind == buttonSequence {
				if bSet.activePosition == hoveredPos && bSet.hasActive {
					bSet.hasActive = false
					bSet.removeButtons()
				} else {
					if bSet.hasActive {
						bSet.removeButtons()
					}
					bSet.activePosition = hoveredPos
					bSet.hasActive = true
					bSet.addButtons(withReset)
				}
			} else if bSet.content[hoveredPos].kind == buttonSelectMove {
				bSet.hasActive = false
				bSet.removeButtons()
			} else {
				bSet.hasActive = false
				bSet.removeButtons()
			}
		}
	}

	return
}

// Draw the buttons
func (buttonSet buttonSet) draw(sequence []int, currentPosition int, inPlay bool, musicOn bool, screen *ebiten.Image) {

	for buttonNum, button := range buttonSet.content {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(button.drawX, button.drawY)

		imageNum := 0
		directionNum := nothing
		switch button.kind {
		case buttonPlay:
			imageNum = 18
			if button.hover {
				imageNum += 2
			}
		case buttonReset:
			imageNum = 19
			if button.hover {
				imageNum += 2
			}
		case buttonSequence:
			imageNum = 15
			directionNum = sequence[button.positionInSequence]
			if buttonSet.hasActive && buttonSet.activePosition == buttonNum {
				imageNum = 17
				directionNum += 2 * nothing
			} else if (!button.hover && buttonSet.onBeat && !inPlay && directionNum == nothing) ||
				(button.hover && !inPlay) ||
				(button.hover && inPlay && (currentPosition != button.positionInSequence || !buttonSet.onBeat)) {
				imageNum = 16
				directionNum += nothing
			} else if currentPosition == button.positionInSequence &&
				buttonSet.onBeat && inPlay && !buttonSet.firstLoop {
				imageNum = 17
				directionNum += 2 * nothing
			}
		case buttonSelectMove:
			imageNum = button.smallPosition
			if imageNum >= sequence[button.positionInSequence] {
				imageNum++
			}
			if !button.smallReset && imageNum >= moveReset {
				imageNum++
			}
			imageNum += levelUpBox + 1
		}

		if button.kind == buttonSelectMove {
			options.GeoM.Translate(-16, -6)
			if button.hover {
				options.GeoM.Translate(0, 3)
			}
			screen.DrawImage(tilesImage.SubImage(
				image.Rect(imageNum*(globalTileSize+2*globalTileMargin), 0,
					(imageNum+1)*(globalTileSize+2*globalTileMargin),
					globalTileSize+2*globalTileMargin)).(*ebiten.Image),
				options)
			continue
		}

		if button.kind == buttonIncBPM ||
			button.kind == buttonDecBPM ||
			button.kind == buttonToggleSound {

			mult := button.kind - buttonIncBPM

			if button.kind == buttonToggleSound && !musicOn {
				mult++
			}

			if button.hover {
				options.GeoM.Translate(0, 1)
			}
			screen.DrawImage(smallbuttonsImage.SubImage(
				image.Rect(mult*globalSmallButtonWidth, 0,
					(mult+1)*(globalSmallButtonWidth),
					globalSmallButtonHeight)).(*ebiten.Image),
				options)
			continue
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
