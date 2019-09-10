/*
 * moduleadd.go
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
	"strconv"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module/serial"
)

// moduleadd handles moduleadd command.
func moduleadd(cmd Command) (int, string) {
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ModuleAdd)
	}
	switch cmd.OptionArgs()[0] {
	case "character", "char":
		return moduleaddCharacter(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", ModuleAdd,
			cmd.OptionArgs()[0])
	}
}

// moduleaddCharacter handles character option for moduleadd.
func moduleaddCharacter(cmd Command) (int, string) {
	if len(cmd.Args()) < 3 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ModuleAdd, cmd.OptionArgs()[0])
	}
	id := cmd.Args()[0]
	scenID := cmd.Args()[1]
	areaID := cmd.Args()[2]
	posX, posY := 0.0, 0.0
	if len(cmd.Args()) > 4 {
		var err error
		posX, err = strconv.ParseFloat(cmd.Args()[3], 64)
		if err != nil {
			return 3, fmt.Sprintf("%s: fail to parse x position: %v",
				ModuleAdd, err)
		}
		posY, err = strconv.ParseFloat(cmd.Args()[4], 64)
		if err != nil {
			return 3, fmt.Sprintf("%s: fail to parse y position: %v",
				ModuleAdd, err)
		}
	}
	char, err := data.Character(flame.Mod(), id)
	if err != nil {
		return 3, fmt.Sprintf("%s: fail to retrieve character: %v",
			ModuleShow, err)
	}
	char.SetPosition(posX, posY)
	for _, s := range flame.Mod().Chapter().Scenarios() {
		if s.ID() != scenID {
			continue
		}
		for _, a := range s.Areas() {
			if a.ID() != areaID {
				continue
			}
			serial.AssignSerial(char)
			a.AddCharacter(char)
			return 0, ""
		}
	}
	return 3, fmt.Sprintf("%s: fail to found scenario area: %s: %s",
		ModuleShow, scenID, areaID)
}
