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

type game struct {
	soundEngine soundEngine
	sequencer   sequencer
	character   character
}

func newGame() (g game) {
	g.soundEngine = newSoundEngine()
	g.sequencer = newSequencer(110, 4)
	g.character.reset(testLevel)
	g.character.moveSequence = []int{moveDown, moveRight, nothing, moveRight, moveRight, moveLeft, moveDown, moveUp}
	return
}
