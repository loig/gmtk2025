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

import "math/rand/v2"

func (g *game) Update() error {

	defer g.soundEngine.playNow()

	g.cursor.update()

	newBeat, halfBeat := g.sequencer.update(&g.soundEngine)

	if newBeat {
		g.buttonSet.setBeat()
		g.character.setBeat()
		g.title.updateOnBeat()
		if g.state == stateIntro {
			if g.intro.updateOnBeat() {
				g.soundEngine.nextSounds[rand.IntN(3)+soundBlip2] = true
			}
		}
		if g.state == stateEnd {
			if g.end.updateOnBeat() {
				g.soundEngine.nextSounds[rand.IntN(3)+soundBlip2] = true
			}
		}
	}

	if halfBeat {
		g.buttonSet.setHalfBeat()
		g.character.setHalfBeat()
		g.title.updateOnBeat()
		if g.state == stateIntro {
			if g.intro.updateOnBeat() {
				g.soundEngine.nextSounds[rand.IntN(3)+soundBlip2] = true
			}
		}
		if g.state == stateEnd {
			if g.end.updateOnBeat() {
				g.soundEngine.nextSounds[rand.IntN(3)+soundBlip2] = true
			}
		}
	}

	if g.state == stateTitle {
		if g.title.update() {
			g.level = 0
			g.state = stateIntro
			g.soundEngine.nextSounds[soundGo] = true
		}
		return nil
	}

	if g.state == stateIntro {
		if g.intro.update() {
			g.setLevel()
			g.soundEngine.nextSounds[soundGo] = true
		}
		return nil
	}

	if g.state == stateEnd {
		if g.end.update() {
			g.reset()
			g.soundEngine.nextSounds[soundGo] = true
		}
		return nil
	}

	clicked, buttonKind, positionInSequence, smallPosition :=
		g.buttonSet.update(g.cursor.x, g.cursor.y, g.state == stateSetupSequence, g.level >= levelStepReset)

	if clicked && buttonKind == buttonIncBPM {
		g.bpm += 5
		if g.bpm > globalMaxBPM {
			g.bpm = globalMaxBPM
		} else {
			g.sequencer.setBpm(g.bpm)
		}
	}

	if clicked && buttonKind == buttonDecBPM {
		g.bpm -= 5
		if g.bpm < globalMinBPM {
			g.bpm = globalMinBPM
		} else {
			g.sequencer.setBpm(g.bpm)
		}
	}

	if clicked && buttonKind == buttonToggleSound {
		g.soundEngine.toggleSound()
	}

	if clicked && buttonKind == buttonReset {
		g.character.reset(levelSet[g.level], g.state == stateSetupSequence)
		g.state = stateSetupSequence
		g.soundEngine.nextSounds[soundBack] = true
	} else {

		// Setup a sequence
		if g.state == stateSetupSequence {
			if clicked && buttonKind == buttonPlay {
				g.soundEngine.nextSounds[soundGo] = true
				g.state = statePlaySequence
				g.buttonSet.setFirstLoop()
			} else if clicked && buttonKind == buttonSelectMove {
				g.character.moveSequence[positionInSequence] =
					getMoveFromChoice(smallPosition, g.character.moveSequence[positionInSequence], g.level >= levelStepReset)
			}
		} else if g.state == statePlaySequence {
			// Run a sequence

			if newBeat && g.character.checkGoal() {
				g.level++
				g.evolutionSubStep++
				g.soundEngine.nextSounds[soundSuccess] = true
				g.setLevel()
				return nil
			}

			if newBeat {
				playSound, soundID := g.character.updateOnBeat()
				if playSound {
					g.soundEngine.nextSounds[soundID] = true
				}
			}

			if halfBeat {
				playSound, soundID := g.character.updateOnHalfBeat()
				if playSound {
					g.soundEngine.nextSounds[soundID] = true
				}
			}

		}

	}

	return nil
}

// Retrieve the move chosen from the number of a small choice
// button and the current move in the sequence of moves
func getMoveFromChoice(choice, currentMove int, withReset bool) (newMove int) {
	newMove = choice
	if newMove >= currentMove {
		newMove++
	}
	if !withReset && newMove >= moveReset {
		newMove++
	}
	return
}
