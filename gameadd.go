/*
 * gameadd.go
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
	"strconv"

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/module/serial"
)

// gameadd handles gameadd command.
func gameadd(cmd Command) (int, string) {
	if Game == nil {
		return 2, fmt.Sprintf("%s: no game set", GameAdd)
	}

	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", GameAdd)
	}
	switch cmd.OptionArgs()[0] {
	case "character", "char":
		return gameaddCharacter(cmd)
	case "area-character", "area-char":
		return gameaddAreaCharacter(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", GameAdd,
			cmd.OptionArgs()[0])
	}
}

// gameaddCharacter handles character option for gameadd.
func gameaddCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 2 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			GameAdd, cmd.OptionArgs()[0])
	}
	id := cmd.Args()[0]
	posX, posY := 0.0, 0.0
	if len(cmd.Args()) > 4 {
		var err error
		posX, err = strconv.ParseFloat(cmd.Args()[3], 64)
		if err != nil {
			return 3, fmt.Sprintf("%s: unable to parse x position: %v",
				GameAdd, err)
		}
		posY, err = strconv.ParseFloat(cmd.Args()[4], 64)
		if err != nil {
			return 3, fmt.Sprintf("%s: unable to parse y position: %v",
				GameAdd, err)
		}
	}
	data := res.Character(id, "")
	if data == nil {
		return 3, fmt.Sprintf("%s: character data not found",
			GameAdd)
	}
	char := character.New(*data)
	char.SetPosition(posX, posY)
	areaID := cmd.Args()[1]
	areas := Game.Module().Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		a.AddCharacter(char)
		return 0, ""
	}
	return 3, fmt.Sprintf("%s: unable to found area: %s",
		GameAdd, areaID)
}

// gameaddAreaCharacter handles area-character option for gameadd.
func gameaddAreaCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			GameAdd, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s", GameAdd, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not a character",
				GameAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, char)
	}
	areas := Game.Module().Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.Args()[0]
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		for _, ob := range objects {
			a.AddCharacter(ob)
		}
		return 0, ""
	}
	return 3, fmt.Sprintf("%s: unable to found area: %s", GameAdd, areaID)
}
