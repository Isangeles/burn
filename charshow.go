/*
 * charshow.go
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
	"strings"

	"github.com/isangeles/flame"
	"github.com/isangeles/flame/core/module/object/character"
)

// charshow handles charshow command.
func charshow(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s:no active game", CharShow)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s:no option args", CharShow)
	}
	switch cmd.OptionArgs()[0] {
	case "items":
		return charshowItems(cmd)
	default:
		return 2, fmt.Sprintf("%s:no_such_option:%s",
			CharShow, cmd.OptionArgs()[0])
	}
}

// charshowItems handles items option for charshow.
func charshowItems(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", CharAdd)
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
	out := ""
	for _, char := range chars {
		for _, it := range char.Inventory().Items() {
			out = fmt.Sprintf("%s%s#%s ", out, it.ID(), it.Serial())
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}
