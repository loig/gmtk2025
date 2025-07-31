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
	"bytes"
	_ "embed"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//go:embed assets/BD07.WAV
var kickBytes []byte

//go:embed assets/SNR07.WAV
var snareBytes []byte

// A sound engine is responsible for playing sounds
// so that the same sound is not played twice at the same
// frame.
type soundEngine struct {
	audioContext *audio.Context
	nextSounds   [numSounds]bool
	sounds       [numSounds][]byte
}

// The list of existing sounds.
const (
	soundKick int = iota
	soundSnare
	numSounds
)

// Initialisation of the sound engine (sound decoding).
func newSoundEngine() (engine soundEngine) {

	var err error
	var sound *wav.Stream
	engine.audioContext = audio.NewContext(44100)

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(kickBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundKick], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(snareBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundSnare], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	return
}

// Play all sounds that have been registered for
// playing in e.nextSounds.
func (e *soundEngine) playNow() {
	for soundID, play := range e.nextSounds {
		if play {
			e.playSound(soundID)
			e.nextSounds[soundID] = false
		}
	}
}

// Play one sound by generating a new player.
func (e soundEngine) playSound(ID int) {
	soundPlayer := e.audioContext.NewPlayerFromBytes(e.sounds[ID])
	soundPlayer.Play()
}
