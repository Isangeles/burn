/*
 * moduleadd.go
 *
 * Copyright 2019-2023 Dariusz Sikora <ds@isangeles.dev>
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
	"strconv"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/serial"
)

// moduleadd handles moduleadd command.
func moduleadd(cmd Command) (int, string) {
	if Module == nil {
		return 2, fmt.Sprintf("%s: no module set", ModuleAdd)
	}
	if cmd.OptionArgs() == nil || len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ModuleAdd)
	}
	switch cmd.OptionArgs()[0] {
	case "character", "char":
		return moduleaddCharacter(cmd)
	case "area-character", "area-char":
		return moduleaddAreaCharacter(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", ModuleAdd,
			cmd.OptionArgs()[0])
	}
}

// moduleaddCharacter handles character option for moduleadd.
func moduleaddCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 2 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ModuleAdd, cmd.OptionArgs()[0])
	}
	id := cmd.Args()[0]
	posX, posY := 0.0, 0.0
	if len(cmd.Args()) > 4 {
		var err error
		posX, err = strconv.ParseFloat(cmd.Args()[3], 64)
		if err != nil {
			return 3, fmt.Sprintf("%s: unable to parse x position: %v",
				ModuleAdd, err)
		}
		posY, err = strconv.ParseFloat(cmd.Args()[4], 64)
		if err != nil {
			return 3, fmt.Sprintf("%s: unable to parse y position: %v",
				ModuleAdd, err)
		}
	}
	data := res.Character(id, "")
	if data == nil {
		return 3, fmt.Sprintf("%s: character data not found",
			ModuleAdd)
	}
	char := character.New(*data)
	char.SetPosition(posX, posY)
	areaID := cmd.Args()[1]
	areas := Module.Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		a.AddObject(char)
		return 0, ""
	}
	return 3, fmt.Sprintf("%s: unable to found area: %s",
		ModuleAdd, areaID)
}

// moduleaddAreaCharacter handles area-character option for moduleadd.
func moduleaddAreaCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ModuleAdd, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s", ModuleAdd, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not a character",
				ModuleAdd, ob.ID(), ob.Serial())
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
			a.AddObject(ob)
		}
		return 0, ""
	}
	return 3, fmt.Sprintf("%s: unable to found area: %s", ModuleAdd, areaID)
}
