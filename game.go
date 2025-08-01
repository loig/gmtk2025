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

import "log"

type game struct {
	state            int
	soundEngine      soundEngine
	sequencer        sequencer
	character        character
	cursor           cursor
	buttonSet        buttonSet
	level            int
	title            title
	intro            intro
	evolutionStep    int
	evolutionSubStep int
}

// Possible game states
const (
	stateSetupSequence int = iota
	statePlaySequence
	stateTitle
	stateIntro
)

func newGame() (g game) {
	loadFonts()
	loadImages()
	initLevels()
	g.soundEngine = newSoundEngine()
	g.sequencer = newSequencer(110, 4)
	g.reset()
	log.Print(g.state)
	return
}

func (g *game) reset() {
	g.intro = setupIntro()
	g.level = 0
	g.evolutionStep = 1
	g.evolutionSubStep = 0
	g.setLevel()
	g.state = stateTitle
}

func (g *game) setLevel() {
	if g.evolutionStep-1 < len(levelSteps) {
		if g.level >= levelSteps[g.evolutionStep-1] {
			g.evolutionStep++
			g.evolutionSubStep = 0
		}
	}
	g.character.reset(levelSet[g.level], true)
	g.state = stateSetupSequence
	g.buttonSet.setupButtons(len(g.character.moveSequence))
}
