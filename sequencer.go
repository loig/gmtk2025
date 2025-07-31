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

// A sequence is a set of steps, each one associated to a time
// the step steps[k] is associated to the time stepsTime[k]
// the time associated to a step is expressed as a float in [0, 1[
// meaning that this step should occur at the corresponding
// proportion of the full sequence duration.
type sequence struct {
	steps       []int
	stepsTime   []float64
	currentStep int
}

// Create a new sequence which steps are given as a string
// of the form "x--x". The time is divided equally by the
// number of characters in the string and a character different
// from '-' means that a sound should be played at that point.
func newSequence(steps string, soundID int) (s sequence) {
	numSteps := len(steps)
	for position, step := range steps {
		if step != '-' {
			s.steps = append(s.steps, soundID)
			s.stepsTime = append(s.stepsTime, float64(position)/float64(numSteps))
		}
	}
	log.Print(s)
	return
}

// A sequencer is responsible for playing a set of sequences
// The full duration of each sequence is in frames (60 frames
// per second) and its value is numBeats*framesPerBeat.
type sequencer struct {
	framesPerBeat int
	numBeats      int
	sequences     []sequence
	currentBeat   int
	currentFrame  int
}

// Create a new sequencer given a bpm (beats per minute)
// and a number of beats per sequencer cycle.
// The actual bpm is an approximation of the requested bpm
// has everything is counted in frames.
func newSequencer(bpm, beats int) (s sequencer) {
	s.framesPerBeat = 3600 / bpm
	s.numBeats = beats
	s.sequences = []sequence{
		newSequence("x-----x-x---", soundKick),
		newSequence("-x-x", soundSnare),
	}
	return
}

// Update a sequencer by checking in each sequence if a sound shall
// be played and restarting the sequencer cycle if needed.
// Also, returns a boolean that tells if a newBeat just started, for
// use by the calling function.
func (s *sequencer) update(soundEngine *soundEngine) (newBeat bool) {

	if s.currentFrame >= s.framesPerBeat {
		s.currentBeat = (s.currentBeat + 1) % s.numBeats
		s.currentFrame = 0
		newBeat = true
	}

	reset := newBeat && s.currentBeat == 0

	timePosition := float64(s.currentBeat*s.framesPerBeat+s.currentFrame) / float64(s.numBeats*s.framesPerBeat)

	for sequencePosition := 0; sequencePosition < len(s.sequences); sequencePosition++ {
		s.sequences[sequencePosition].update(timePosition, reset, soundEngine)
	}

	s.currentFrame++

	return
}

// Update a sequence given a timePosition (that is a float in [0,1[
// telling at which proportion of the cycle the sequencer is)
func (s *sequence) update(timePosition float64, hasReset bool, soundEngine *soundEngine) {

	if hasReset {
		s.currentStep = 0
	}

	if s.currentStep < len(s.steps) && s.stepsTime[s.currentStep] <= timePosition {
		soundEngine.nextSounds[s.steps[s.currentStep]] = true
		s.currentStep++
	}
}
