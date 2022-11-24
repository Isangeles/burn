/*
 * areashow_test.go
 *
 * Copyright 2022 Dariusz Sikora <ds@isangeles.dev>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 *
 *
 */

package burn

import (
	"fmt"
	"testing"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/area"
	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
)

var (
	charData = res.CharacterData{ID: "char"}
	areaData = res.AreaData{ID: "area"}
)

// TestAreaShowObjects tests objects option for areashow.
func TestAreaShowObjects(t *testing.T) {
	// Create module with area and char.
	Module = flame.NewModule(res.ModuleData{})
	area := area.New()
	area.Apply(areaData)
	Module.Chapter().AddAreas(area)
	char := character.New(charData)
	area.AddObject(char)
	// Create command.
	cmd := testCommand{
		tool: AreaShow,
		optionArgs: []string{"objects"},
		targetArgs: []string{area.ID()},
	}
	// Test.
	res, out := areashow(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	expOut := fmt.Sprintf("%s%s%s", char.ID(), IDSerialSep, char.Serial())
	if out != expOut {
		t.Errorf("Command output invalid: '%s' != '%s'", out, expOut)
	}
}
