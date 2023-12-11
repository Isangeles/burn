/*
 * moduleshow_test.go
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

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/area"
	"github.com/isangeles/flame/data/res"
)

// TestModuleShowNoOptions tests handling moduleshow command
// with no options/args.
func TestModuleShowNoOptions(t *testing.T) {
	Module = flame.NewModule(res.ModuleData{})
	res, out := objectadd(testCommand{})
	if res != 2 {
		t.Errorf("Command result invalid: %d != 2", res)
	}
	expOut := fmt.Sprintf("%s: no option args", ObjectAdd)
	if out != expOut {
		t.Errorf("Command output invalid: '%s' != '%s'", out, expOut)
	}
}

// TestModuleShowID tests id option for moduleshow.
func TestModuleShowID(t *testing.T) {
	modConf := make(map[string][]string)
	modConf["id"] = []string{"test"}
	Module = flame.NewModule(res.ModuleData{Config: modConf})
	cmd := testCommand{
		optionArgs: []string{"id"},
	}
	res, out := moduleshow(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if out != Module.Conf().ID {
		t.Errorf("Command output invalid: '%s' != '%s'", out,
			Module.Conf().ID)
	}
}

// TestModuleShowChapter tests chapter option for moduleshow.
func TestModuleShowChapter(t *testing.T) {
	modConf := make(map[string][]string)
	modConf["chapter"] = []string{"test"}
	Module = flame.NewModule(res.ModuleData{Config: modConf})
	cmd := testCommand{
		optionArgs: []string{"chapter"},
	}
	res, out := moduleshow(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	if out != Module.Conf().Chapter {
		t.Errorf("Command output invalid: '%s' != '%s'", out,
			Module.Conf().Chapter)
	}
}

// TestModuleShowAreas tests areas option for moduleshow.
func TestModuleShowAreas(t *testing.T) {
	area1 := area.New(res.AreaData{ID: "area1"})
	area2 := area.New(res.AreaData{ID: "area2"})
	Module = flame.NewModule(res.ModuleData{})
	Module.Chapter().AddAreas(area1, area2)
	cmd := testCommand{
		optionArgs: []string{"areas"},
	}
	res, out := moduleshow(cmd)
	if res != 0 {
		t.Errorf("Command result invalid: %d != 0", res)
	}
	expOut := fmt.Sprintf("%s %s", area1.ID(), area2.ID())
	if out != expOut {
		t.Errorf("Command output invalid: '%s' != '%s'", out,
			expOut)
	}
}
