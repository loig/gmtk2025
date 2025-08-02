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

//go:embed assets/CLHAT1.WAV
var hatsBytes []byte

//go:embed assets/ArpBC2.wav
var c2Bytes []byte

//go:embed assets/ArpBC3.wav
var c3Bytes []byte

//go:embed assets/ArpBC4.wav
var c4Bytes []byte

//go:embed assets/ArpBE3.wav
var e3Bytes []byte

//go:embed assets/ArpBG3.wav
var g3Bytes []byte

//go:embed assets/ArpBG4.wav
var g4Bytes []byte

//go:embed assets/RolandEBlip22.wav
var blipBytes []byte

//go:embed assets/RolandEBlip17.wav
var blip2Bytes []byte

//go:embed assets/RolandEBlip18.wav
var blip3Bytes []byte

//go:embed assets/RolandEBlip20.wav
var blip4Bytes []byte

//go:embed assets/RolandEBlip10.wav
var successBytes []byte

//go:embed assets/RolandEBlip11.wav
var goBytes []byte

//go:embed assets/RolandEBlip06.wav
var backBytes []byte

// A sound engine is responsible for playing sounds
// so that the same sound is not played twice at the same
// frame.
type soundEngine struct {
	audioContext *audio.Context
	nextSounds   [numSounds]bool
	sounds       [numSounds][]byte
	mute         bool
}

// The list of existing sounds.
const (
	soundKick int = iota
	soundSnare
	soundHats
	soundC2
	soundC3
	soundC4
	soundE3
	soundG3
	soundG4
	soundBlip
	soundBlip2
	soundBlip3
	soundBlip4
	soundSuccess
	soundGo
	soundBack
	numSounds
)

// Toggle sound
func (s *soundEngine) toggleSound() {
	s.mute = !s.mute
}

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

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(hatsBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundHats], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(c2Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundC2], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(c3Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundC3], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(c4Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundC4], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(e3Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundE3], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(g3Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundG3], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(g4Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundG4], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(blipBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundBlip], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(blip2Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundBlip2], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(blip3Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundBlip3], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(blip4Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundBlip4], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(successBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundSuccess], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(goBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundGo], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(backBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundBack], err = io.ReadAll(sound)
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
	if !e.mute {
		soundPlayer := e.audioContext.NewPlayerFromBytes(e.sounds[ID])
		soundPlayer.Play()
	}
}
