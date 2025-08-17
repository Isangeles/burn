/*
 * objecthave_test.go
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
	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/flag"
)

// TestObjectHaveNoOptions tests handling of the command without any options/args.
func TestObjectHaveNoOptions(t *testing.T) {
	// Create module
	Module = flame.NewModule(res.ModuleData{})
	// Create command
	cmd := testCommand{
		tool: ObjectHave,
	}
	// Test
	res, out := objecthave(cmd)
	if res != 2 {
		t.Errorf("Command result invalid: %d != 2", res)
	}
	expOut := fmt.Sprintf("%s: no option args", ObjectHave)
	if out != expOut {
		t.Errorf("Command output invalid: '%s' != '%s'", out, expOut)
	}
}

// TestObjectHaveFlag tests handling objecthave flag option.
func TestObjectHaveFlag(t *testing.T) {
	// Create module
	Module = flame.NewModule(res.ModuleData{})
	char := character.New(charData)
	// Create command
	cmd := testCommand{
		tool:       ObjectHave,
		optionArgs: []string{"flag"},
		targetArgs: []string{fmt.Sprintf("%s#%s", char.ID(), char.Serial())},
		args:       []string{"testFlag"},
	}
	// Test - flag not present
	res, out := objecthave(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if out != "false" {
		t.Errorf("Command output invalid: %s != 'false'", out)
	}
	// Test - flag present
	char.AddFlag(flag.Flag("testFlag"))
	res, out = objecthave(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if out != "true" {
		t.Errorf("Command output invalid: %s != 'true'", out)
	}
}
