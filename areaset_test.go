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
	"testing"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/area"
	"github.com/isangeles/flame/data/res"
)

// TestAreaSetWeather tests weather option for areaset.
func TestAreaSetWeather(t *testing.T) {
	// Create module with area.
	Module = flame.NewModule(res.ModuleData{})
	modArea := area.New(areaData)
	Module.Chapter().AddAreas(modArea)
	// Create command.
	cmd := testCommand{
		tool: AreaShow,
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
