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
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed assets/OfficeCodePro-Bold.ttf
var ocpRegular_ttf []byte
var ocpFaceSource *text.GoTextFaceSource
var ocpFace *text.GoTextFace

func loadFonts() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(ocpRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}
	ocpFaceSource = s

	ocpFace = &text.GoTextFace{
		Source: ocpFaceSource,
		Size:   24,
	}
}

func drawTextAt(theText string, x, y float64, screen *ebiten.Image) (height float64) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(x, y+2)
	op.ColorScale.ScaleWithColor(color.RGBA{R: 0x54, G: 0x33, B: 0x44, A: 255})
	op.LineSpacing = ocpFace.Size * 1.5
	text.Draw(screen, theText, ocpFace, op)

	op = &text.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(color.RGBA{R: 0x8b, G: 0x40, B: 0x49, A: 255})
	op.LineSpacing = ocpFace.Size * 1.5
	text.Draw(screen, theText, ocpFace, op)

	_, height = text.Measure(theText, ocpFace, ocpFace.Size*1.5)
	return
}
