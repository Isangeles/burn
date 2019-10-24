/*
 * gameremove.go
 *
 * Copyright 2019 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/module/object/character"
)

// gameremove handles gameremove command.
func gameremove(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s: no game started", GameRemove)
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
	if len(cmd.Args()) < 2 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			GameAdd, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
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
	scenID := cmd.Args()[0]
	areaID := cmd.Args()[1]
	for _, s := range flame.Mod().Chapter().Scenarios() {
		if s.ID() != scenID {
			continue
		}
		if s.Mainarea().ID() == areaID {
			for _, ob := range objects {
				s.Mainarea().RemoveCharacter(ob)
			}
			return 0, ""
		}
		for _, a := range s.Mainarea().Subareas() {
			if a.ID() != areaID {
				continue
			}
			for _, ob := range objects {
				a.RemoveCharacter(ob)
			}
			return 0, ""
		}
	}
	return 3, fmt.Sprintf("%s: fail to found scenario area: %s: %s",
		GameAdd, scenID, areaID)
}