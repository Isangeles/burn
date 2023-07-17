/*
 * areaset_test.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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
	"time"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/area"
	"github.com/isangeles/flame/data/res"
)

// TestAreaSetNoOptions tests handling of the command without any options/args.
func TestAreaSetNoOptions(t *testing.T) {
	// Create module.
	Module = flame.NewModule(res.ModuleData{})
	// Create command.
	cmd := testCommand{
		tool: AreaSet,
	}
	// Test.
	res, out := areaset(cmd)
	if res != 2 {
		t.Errorf("Command result invalid: %d != 2", res)
	}
	expOut := fmt.Sprintf("%s: no option args", AreaSet)
	if out != expOut {
		t.Errorf("Command output invalid: '%s' != '%s'", out, expOut)
	}
}

// TestAreaSetWeather tests weather option for areaset.
func TestAreaSetWeather(t *testing.T) {
	// Create module with area.
	Module = flame.NewModule(res.ModuleData{})
	modArea := area.New(areaData)
	Module.Chapter().AddAreas(modArea)
	// Create command.
	cmd := testCommand{
		tool:       AreaSet,
		optionArgs: []string{"weather"},
		targetArgs: []string{modArea.ID()},
		args:       []string{"weatherRain"},
	}
	// Test.
	res, out := areaset(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) != 0 {
		t.Errorf("Command output invalid: '%s' != ''", out)
	}
	if modArea.Weather().Conditions != area.Rain {
		t.Errorf("Area weather not changed: '%s' != '%s'", modArea.Weather().Conditions,
			area.Rain)
	}
}

// TestAreaSetTime tests time option for areaset.
func TestAreaSetTime(t *testing.T) {
	// Create module with area.
	Module = flame.NewModule(res.ModuleData{})
	modArea := area.New(areaData)
	Module.Chapter().AddAreas(modArea)
	// Create command.
	cmd := testCommand{
		tool:       AreaSet,
		optionArgs: []string{"time"},
		targetArgs: []string{modArea.ID()},
		args:       []string{"10:00AM"},
	}
	// Test.
	res, out := areaset(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if len(out) != 0 {
		t.Errorf("Command output invalid: '%s' != ''", out)
	}
	if modArea.Time.Format(time.Kitchen) != "10:00AM" {
		t.Errorf("Area time not changed: '%s' != '10:00AM'",
			modArea.Time.Format(time.Kitchen))
	}
	// Test invalid time.
	cmd.args[0] = "invalid"
	res, out = areaset(cmd)
	if res != 3 {
		t.Errorf("Invalid error result code: %d != 3", res)
	}
	if len(out) < 1 {
		t.Errorf("No error message")
	}
}
