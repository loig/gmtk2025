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
)

func (g *game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{R: 0xca, G: 0xa0, B: 0x5a, A: 255})

	g.character.draw(screen)

	g.buttonSet.draw(
		g.character.moveSequence,
		(g.character.nextMovePosition+len(g.character.moveSequence)-1)%len(g.character.moveSequence),
		g.state == statePlaySequence,
		screen)

	g.cursor.draw(screen)

}
