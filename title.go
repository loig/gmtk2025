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
	onBeat bool
	upChar int
}

func (t title) draw(screen *ebiten.Image) {

	for charNum := 0; charNum < 4; charNum++ {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(charNum)*200, 100)
		if t.onBeat && t.upChar == charNum {
			options.GeoM.Translate(0, -10)
		}

		screen.DrawImage(charsImage.SubImage(
			image.Rect(charNum*200, 0, (charNum+1)*200, 260)).(*ebiten.Image), options)
	}

	y := 450
	if t.onBeat {
		y -= 5
	}
	drawTextAt("Click to start", 300, float64(y), screen)

}

func (t *title) updateOnBeat() {
	t.onBeat = !t.onBeat
	if !t.onBeat {
		t.upChar = (t.upChar + 1) % 4
	}
}

func (t *title) update() (done bool) {
	done = inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	return
}
