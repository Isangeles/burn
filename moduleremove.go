/*
 * moduleremove.go
 *
 * Copyright 2021-2025 Dariusz Sikora <ds@isangeles.dev>
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

	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/serial"
)

// moduleremove handles moduleremove command.
func moduleremove(cmd Command) (int, string) {
	if Module == nil {
		return 2, fmt.Sprintf("%s: no module set", ModuleRemove)
	}
	if cmd.OptionArgs() == nil || len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ModuleRemove)
	}
	switch cmd.OptionArgs()[0] {
	case "area-character", "area-char":
		return moduleremoveAreaCharacter(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", ModuleRemove,
			cmd.OptionArgs()[0])
	}
}

// moduleremoveAreaCharacter handles area-character option for moduleremove.
func moduleremoveAreaCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enough args for: %s",
			ModuleRemove, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s", ModuleRemove, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s %s: not a character",
				ModuleRemove, ob.ID(), ob.Serial())
		}
		objects = append(objects, char)
	}
	areas := Module.Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.Args()[0]
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		for _, ob := range objects {
			a.RemoveObject(ob)
		}
		return 0, ""
	}
	return 3, fmt.Sprintf("%s: module area not found: %s",
		ModuleRemove, areaID)
}
