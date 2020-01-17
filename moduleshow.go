/*
 * moduleshow.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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
	"strings"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/module/area"
)

// moduleshow handles moduleshow command.
func moduleshow(cmd Command) (int, string) {
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ModuleShow)
	}
	switch cmd.OptionArgs()[0] {
	case "id":
		return 0, flame.Mod().Conf().ID
	case "area-chars":
		return moduleshowAreaChars(cmd)
	case "area-objects":
		return moduleshowAreaObjects(cmd)
	case "areas":
		return moduleshowAreas(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", ModuleShow,
			cmd.OptionArgs()[0])
	}
}

// moduleshowAreaChars handles area-chars option for moduleshow.
func moduleshowAreaChars(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 3, fmt.Sprintf("%s: no game loaded", ModuleShow)
	}
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought target args for: %s",
			ModuleShow, cmd.OptionArgs()[0])
	}
	areas := flame.Mod().Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.TargetArgs()[0]
	var area *area.Area
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		area = a
		break
	}
	if area == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			ModuleShow, areaID)
	}
	out := ""
	for _, c := range area.Characters() {
		out = fmt.Sprintf("%s%s%s%s ", out, c.ID(),
			IDSerialSep, c.Serial())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// moduleshowAreaObjects handles area-objects
// option for moduleshow.
func moduleshowAreaObjects(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 3, fmt.Sprintf("%s: no game loaded", ModuleShow)
	}
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ModuleShow, cmd.Args()[0])
	}
	areas := flame.Mod().Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.TargetArgs()[0]
	var area *area.Area
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		area = a
		break
	}
	if area == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			ModuleShow, areaID)
	}
	out := ""
	for _, o := range area.Objects() {
		out = fmt.Sprintf("%s%s%s%s ", out, o.ID(),
			IDSerialSep, o.Serial())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// moduleshowAreas handles areas option for moduleshow.
func moduleshowAreas(cmd Command) (int, string) {
	chapter := flame.Mod().Chapter()
	if chapter == nil {
		return 3, fmt.Sprintf("%s: no active chapter",
			ModuleShow)
	}
	areas := chapter.Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	out := ""
	for _, a := range areas {
		out = fmt.Sprintf("%s%s ", out, a.ID())
	}
	out = strings.TrimSpace(out)
	return 0, out
}
