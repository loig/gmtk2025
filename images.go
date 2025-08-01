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
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/tiles.png
var tilesBytes []byte
var tilesImage *ebiten.Image

//go:embed assets/buttons.png
var buttonsBytes []byte
var buttonsImage *ebiten.Image

//go:embed assets/cursor.png
var cursorBytes []byte
var cursorImage *ebiten.Image

// load all images
func loadImages() {
	decoded, _, err := image.Decode(bytes.NewReader(tilesBytes))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(buttonsBytes))
	if err != nil {
		log.Fatal(err)
	}
	buttonsImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(cursorBytes))
	if err != nil {
		log.Fatal(err)
	}
	cursorImage = ebiten.NewImageFromImage(decoded)
}
