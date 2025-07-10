/*
 * moduleadd_test.go
 *
 * Copyright 2025 Dariusz Sikora <ds@isangeles.dev>
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

// TestModuleAddNoOptions tests handling of the command without any options/args.
func TestModuleAddNoOptions(t *testing.T) {
	// Create module
	Module = flame.NewModule(res.ModuleData{})
	// Create command
	cmd := testCommand{
		tool: ModuleAdd,
	}
	// Test
	res, out := moduleadd(cmd)
	if res != 2 {
		t.Errorf("Command result invalid: %d != 2", res)
	}
	expOut := fmt.Sprintf("%s: no option args", ModuleAdd)
	if out != expOut {
		t.Errorf("Command output invalid: '%s' != '%s'", out, expOut)
	}
}

// TestModuleAddCharacter tests character option for moduleadd.
func TestModuleAddCharacter(t *testing.T) {
	// Create test data
	Module = flame.NewModule(res.ModuleData{})
	modArea := area.New(res.AreaData{ID: "area"})
	Module.Chapter().AddAreas(modArea)
	res.Characters = append(res.Characters, charData)
	// Create command
	cmd := testCommand{
		optionArgs: []string{"character"},
		args:       []string{charData.ID, modArea.ID(), "10", "20"},
	}
	// Test
	res, out := moduleadd(cmd)
	if res != 0 {
		t.Fatalf("Command result invalid: %d != 0", res)
	}
	if len(out) > 0 {
		t.Fatalf("Command output invalid: '%s'", out)
	}
	var ob area.Object
	for _, o := range modArea.Objects() {
		if o.ID() != cmd.args[0] {
			continue
		}
		ob = o
		break
	}
	char, _ := ob.(*character.Character)
	if char == nil {
		t.Fatalf("New character not found in the area")
	}
	posX, posY := char.Position()
	if posX != 10 || posY != 20 {
		t.Fatalf("New character position invalid: %fx%f != 10x20", posX, posY)
	}
}

// TestModuleAddAreaCharacter tests area-character option for moduleadd.
func TestModuleAddAreaCharacter(t *testing.T) {
	// Create test data
	Module = flame.NewModule(res.ModuleData{})
	area1 := area.New(res.AreaData{ID: "area1"})
	area2 := area.New(res.AreaData{ID: "area2"})
	char := character.New(res.CharacterData{ID: "object"})
	area1.AddObject(char)
	Module.Chapter().AddAreas(area1, area2)
	res.Characters = append(res.Characters, charData)
	// Create command
	cmd := testCommand{
		optionArgs: []string{"area-character"},
		targetArgs: []string{fmt.Sprintf("%s#%s", char.ID(), char.Serial())},
		args:       []string{area2.ID()},
	}
	// Test
	res, out := moduleadd(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) > 0 {
		t.Fatalf("Command output invalid: '%s'", out)
	}
	var ob area.Object
	for _, o := range area2.Objects() {
		if o.ID()+o.Serial() != char.ID()+char.Serial() {
			continue
		}
		ob = o
		break
	}
	char, _ = ob.(*character.Character)
	if char == nil {
		t.Fatalf("Character not found in the target area")
	}
}
