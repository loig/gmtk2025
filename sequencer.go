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
	"math/rand"
)

// A sequence is a set of steps, each one associated to a time
// the step steps[k] is associated to the time stepsTime[k]
// the time associated to a step is expressed as a float in [0, 1[
// meaning that this step should occur at the corresponding
// proportion of the full sequence duration.
// Each step also has a probability to make a sound or not.
type sequence struct {
	steps       []int
	stepsTime   []float64
	stepsProba  []float64
	currentStep int
}

// Create a new sequence which steps are given as a string
// of the form "x--x". The time is divided equally by the
// number of characters in the string.
// '-' means that no sound should be played at that point.
// Anything else means that a sound could be played.
// A number (1 to 9) X gives a probability 0.X to play a sound.
func newSequence(steps string, soundID int) (s sequence) {
	numSteps := len(steps)
	for position, step := range steps {
		if step != '-' {
			s.steps = append(s.steps, soundID)
			s.stepsTime = append(s.stepsTime, float64(position)/float64(numSteps))
			if int(step) >= 49 && int(step) <= 57 {
				proba := float64(int(step)-48) / 10
				s.stepsProba = append(s.stepsProba, proba)
			} else {
				s.stepsProba = append(s.stepsProba, 1)
			}
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

// Set the bpm of a given sequencer while keeping the state of
// the sequence currently playing
func (s *sequencer) setBpm(bpm int) {
	newFramesPerBeat := 3600 / (bpm * 2)
	if s.framesPerBeat != 0 {
		s.currentFrame = (s.currentFrame * newFramesPerBeat) / s.framesPerBeat
	} else {
		s.currentFrame = 0
	}
	s.framesPerBeat = newFramesPerBeat
}

// Create a new sequencer given a bpm (beats per minute)
// and a number of beats per sequencer cycle.
// The actual bpm is an approximation of the requested bpm
// has everything is counted in frames.
func newSequencer(bpm, beats int) (s sequencer) {
	s.setBpm(bpm)
	s.numBeats = beats * 2
	s.sequences = []sequence{
		newSequence("x-2-----x-x-----x-----1xx-----x-x-------x-x-----x1x---x---x---5-", soundKick),
		newSequence("----x-11----x-------x-------x-------x--8----x-------x--2----x---", soundSnare),
		newSequence("--xx--x-xxx--xx--5xx--x-xxx--xx5--xx--x-xxx-5xx--5xx--x5xxx-5xx-", soundHats),
	}
	return
}

// Update a sequencer by checking in each sequence if a sound shall
// be played and restarting the sequencer cycle if needed.
// Also, returns a boolean that tells if a newBeat just started, for
// use by the calling function.
func (s *sequencer) update(soundEngine *soundEngine) (newBeat, halfBeat bool) {

	if s.currentFrame >= s.framesPerBeat {
		s.currentBeat = (s.currentBeat + 1) % s.numBeats
		s.currentFrame = 0
	}

	reset := s.currentFrame == 0 && s.currentBeat == 0

	timePosition := float64(s.currentBeat*s.framesPerBeat+s.currentFrame) / float64(s.numBeats*s.framesPerBeat)

	for sequencePosition := 0; sequencePosition < len(s.sequences); sequencePosition++ {
		s.sequences[sequencePosition].update(timePosition, reset, soundEngine)
	}

	newBeat = s.currentFrame == 0
	halfBeat = s.currentFrame == s.framesPerBeat/2

	s.currentFrame++

	return
}

// Update a sequence given a timePosition (that is a float in [0,1[
// telling at which proportion of the cycle the sequencer is).
func (s *sequence) update(timePosition float64, hasReset bool, soundEngine *soundEngine) {

	if hasReset {
		s.currentStep = 0
	}

	if s.currentStep < len(s.steps) && s.stepsTime[s.currentStep] <= timePosition {
		if s.stepsProba[s.currentStep] >= rand.Float64() {
			soundEngine.nextSounds[s.steps[s.currentStep]] = true
		}
		s.currentStep++
	}
}
