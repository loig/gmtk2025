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
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{R: 0xca, G: 0xa0, B: 0x5a, A: 255})

	if g.state == stateTitle {
		g.title.draw(screen)
	} else if g.state == stateIntro {
		g.intro.draw(screen)
	} else if g.state == stateEnd {
		g.end.draw(screen)
	} else {
		if g.level == 0 {
			//drawTuto(screen)
		}

		g.character.draw(screen)

		g.buttonSet.draw(
			g.character.moveSequence,
			g.character.currentMovePosition,
			g.character.HideMove,
			g.state == statePlaySequence,
			!g.soundEngine.mute,
			screen)

		g.boxSwitcher.draw(screen)

		g.drawLevelInfo(screen)
	}

	g.cursor.draw(screen)

	//drawTextAt(fmt.Sprintf("TPS: %f, FPS:%f", ebiten.ActualTPS(), ebiten.ActualFPS()), 0, 0, screen)

}

func (g game) drawLevelInfo(screen *ebiten.Image) {

	text := fmt.Sprintf("Cybernetic Unit Benchmark ver. 0.%d", g.evolutionStep)
	if g.evolutionSubStep > 0 {
		text = fmt.Sprintf("%s.%d", text, g.evolutionSubStep)
	}
	drawTextAt(text, 20, 10, screen)

	text = fmt.Sprintf("Freq. %d", g.bpm)
	drawTextAt(text, 650, 10, screen)

}

func drawTuto(screen *ebiten.Image) {
	drawTextAt("Change speed →", 505, 50, screen)
	drawTextAt("← Stop sound", 50, 50, screen)
	drawTextAt("This is C.U.B →", 50, 150, screen)
	drawTextAt("← Move C.U.B here", 535, 290, screen)
	drawTextAt("Create loop\n     ↓", 320, 420, screen)
	drawTextAt("Start loop\n  ↓", 10, 435, screen)
	drawTextAt("Restart level\nor erase loop\n          ↓", 600, 400, screen)
}
