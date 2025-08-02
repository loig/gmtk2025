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

type title struct {
	onBeat    bool
	upChar    int
	upSubChar int
}

func (t title) draw(screen *ebiten.Image) {

	// Title
	for charNum := 0; charNum < 4; charNum++ {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(charNum)*200, 70)
		if t.onBeat && t.upChar == charNum {
			options.GeoM.Translate(0, -10)
		}

		screen.DrawImage(charsImage.SubImage(
			image.Rect(charNum*200, 0, (charNum+1)*200, 260)).(*ebiten.Image), options)
	}

	// Subtitle
	charPositions := [7]int{0, 1, 5, 2, 5, 3, 4}
	charSizes := [7]int{50, 50, 20, 50, 20, 50, 50}
	charX := 255
	for pos, charPos := range charPositions {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(charX), 330)

		if t.upSubChar == pos+1 {
			options.GeoM.Translate(0, -5)
		}

		screen.DrawImage(subtitleImage.SubImage(
			image.Rect(charPos*50, 0, charPos*50+charSizes[pos], 55)).(*ebiten.Image), options)

		charX += charSizes[pos]
	}

	// Click text
	y := 470
	if t.onBeat {
		y -= 5
	}
	drawTextAt("Click to start", 300, float64(y), screen)

	// Info text
	text := "A game for GMTK game jam 2025"
	drawTextAt(text, 20, 10, screen)

}

func (t *title) updateOnBeat() {
	t.onBeat = !t.onBeat
	if !t.onBeat {
		t.upChar = (t.upChar + 1) % 4
	}
	t.upSubChar = (t.upSubChar + 1) % 8
}

func (t *title) update() (done bool) {
	done = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	return
}
