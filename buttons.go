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

func drawButtons(sequence []int, currentPosition int, drawBeat bool, screen *ebiten.Image) {

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

}
