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
	"log"
)

func (g *game) Update() error {

	g.cursor.update()

	newBeat, halfBeat := g.sequencer.update(&g.soundEngine)

	if newBeat {
		g.buttonSet.setBeat()
	}

	if halfBeat {
		g.buttonSet.setHalfBeat()
	}

	clicked, buttonKind := g.buttonSet.update(g.cursor.x, g.cursor.y)

	if clicked && buttonKind == buttonReset {
		g.character.reset(levelSet[g.level])
		g.state = stateSetupSequence
	} else {

		// Setup a sequence
		if g.state == stateSetupSequence {
			if clicked && buttonKind == buttonPlay {
				g.state = statePlaySequence
				g.buttonSet.setFirstLoop()
			}
		} else if g.state == statePlaySequence {

			// Run a sequence
			if newBeat {
				playSound, soundID := g.character.updateOnBeat()
				if playSound {
					g.soundEngine.nextSounds[soundID] = true
				}
			}

			if halfBeat {
				g.character.updateOnHalfBeat()
			}

			if g.character.checkGoal() {
				log.Print("The end")
			}
		}

	}

	g.soundEngine.playNow()

	return nil
}
