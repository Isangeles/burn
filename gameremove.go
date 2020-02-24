/*
 * gameremove.go
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

	"github.com/isangeles/flame/core/module/character"
)

// gameremove handles gameremove command.
func gameremove(cmd Command) (int, string) {
	if Game == nil {
		return 2, fmt.Sprintf("%s: no game set", GameRemove)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", GameRemove)
	}
	switch cmd.OptionArgs()[0] {
	case "area-character", "area-char":
		return gameremoveAreaCharacter(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", GameRemove,
			cmd.OptionArgs()[0])
	}
}

// gameremoveAreaCharacter handles area-character option for gameremove.
func gameremoveAreaCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			GameAdd, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := Game.Module().Object(id, serial)
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
			a.RemoveCharacter(ob)
		}
		return 0, ""
	}
	return 3, fmt.Sprintf("%s: fail to found area: %s",
		GameAdd, areaID)
}
