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
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type intro struct {
	beat int
	step int
	text []string
}

func setupIntro() (i intro) {
	i.beat = 0
	i.step = 0
	i.text = []string{
		"July 22, 2010.",
		"07:56 am.",
		"",
		"Initiating Cybernetic Unit Benchmark.",
		"",
		"Objective 1: Test looping system.",
		"Objective 2: Evaluate sentience.",
		"",
		"Loading...",
		"Setting up..",
		"...",
		"..",
		"....",
		"Ready.",
		"",
		"Click to start.",
	}
	return
}

func (l *intro) update() (done bool) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		l.step++
		return l.step > len(l.text)
	}

	return
}

func (l *intro) updateOnBeat() (playSound bool) {
	if l.step < len(l.text) {
		playSound = l.beat < len(l.text[l.step]) && l.text[l.step][l.beat] != ' '
		l.beat++
		if l.beat >= len(l.text[l.step])+2 {
			l.beat = 0
			l.step++
		}
	}
	return
}

func (l intro) draw(screen *ebiten.Image) {
	x := 10
	y := 20
	for pos := 0; pos < l.step && pos < len(l.text); pos++ {
		drawTextAt(l.text[pos], float64(x), float64(y), screen)
		y += 25
	}

	if l.step < len(l.text) {
		upTo := l.beat
		if upTo > len(l.text[l.step]) {
			upTo = len(l.text[l.step])
		}
		drawTextAt(l.text[l.step][:upTo], float64(x), float64(y), screen)
	}
}
