/*
 * charadd.go
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
	"github.com/isangeles/flame/core/data"
	"github.com/isangeles/flame/core/module/object/character"
)

// charadd handles charadd command.
func charadd(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s:no active game", CharAdd)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s:no option args", CharAdd)
	}
	switch cmd.OptionArgs()[0] {
	case "item":
		return charaddItem(cmd)
	default:
		return 2, fmt.Sprintf("%s:no_such_option:%s",
			CharAdd, cmd.OptionArgs()[0])
	}
}

// charaddItem handles item option for charadd.
func charaddItem(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", CharAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s:no_enought_args_for:%s",
			CharAdd, cmd.OptionArgs()[0])
	}
	chars := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		char := flame.Game().Module().Chapter().Character(id, serial)
		if char == nil {
			return 3, fmt.Sprintf("%s:character_not_found:%s", CharAdd, arg)
		}
		chars = append(chars, char)
	}
	id := cmd.Args()[0]
	item, err := data.Item(id)
	if err != nil {
		return 3, fmt.Sprintf("%s:fail_to_retrieve_item:%v", CharAdd, err)
	}
	for _, char := range chars {
		err = char.Inventory().AddItem(item)
		if err != nil {
			return 3, fmt.Sprintf("%s:char:%s#%s:fail_to_add_item:%v",
				CharAdd, char.ID(), char.Serial(), err)
		}
	}
	return 0, ""
}
