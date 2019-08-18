/*
 * objecthave.go
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
	"github.com/isangeles/flame/core/module/flag"
)

// objecthave handles objecthave command.
func objecthave(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s: no active game", ObjectShow)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ObjectShow)
	}
	switch cmd.OptionArgs()[0] {
	case "flag":
		return objecthaveFlag(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s",
			ObjectShow, cmd.OptionArgs()[0])
	}
}

// objecthaveFlag handles flag option for objecthave.
func objecthaveFlag(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectHave)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no args", ObjectHave)
	}
	objects := make([]flag.Flagger, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectHave, arg)
		}
		flagger, ok := ob.(flag.Flagger)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no flags",
				ObjectHave, ob.ID(), ob.Serial())
		}
		objects = append(objects, flagger)
	}
	for _, ob := range objects {
		have := false
		for _, f := range ob.Flags() {
			if f.ID() == cmd.Args()[0] {
				have = true
				break
			}
		}
		if !have {
			return 0, "false"
		}
	}
	return 0, "true"
}
