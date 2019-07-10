/*
 * charset.go
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
	"github.com/isangeles/flame/core/module/object/character"
)

// charset handles charset command.
func charset(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s:no active game", CharSet)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s:no option args", CharSet)
	}
	switch cmd.OptionArgs()[0] {
	case "experience", "exp":
		return charsetExperience(cmd)
	default:
		return 2, fmt.Sprintf("%s:no_such_option:%s",
			CharSet, cmd.OptionArgs()[0])
	}
}

// charsetExperience handles experience option for charset.
func charsetExperience(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s:no_enought_args_for:%s",
			CharSet, cmd.OptionArgs()[0])
	}
	chars := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		char := flame.Game().Module().Chapter().Character(id, serial)
		if char == nil {
			return 3, fmt.Sprintf("%s:character_not_found:%s",
				CharSet, arg)
		}
		chars = append(chars, char)
	}
	val, err := strconv.Atoi(cmd.Args()[0])
	if err != nil {
		return 3, fmt.Sprintf("%s:invalid_argument:%s", CharSet,
			cmd.Args()[0])
	}
	for _, char := range chars {
		char.SetExperience(val)
	}
	return 0, ""
}
