/*
 * objectshow.go
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
	"github.com/isangeles/flame/core/module/object"
)

// objectshow handles objectshow command.
func objectshow(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s:no active game", ObjectShow)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s:no option args", ObjectShow)
	}
	switch cmd.OptionArgs()[0] {
	case "position", "pos":
		return objectshowPosition(cmd)
	default:
		return 2, fmt.Sprintf("%s:no_such_option:%s",
			ObjectShow, cmd.OptionArgs()[0])
	}
}

// objectshowPosition handles position option for objectshow.
func objectshowPosition(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", ObjectShow)
	}
	objects := make([]object.Positioner, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s:object_not_found:%s",
				ObjectShow, arg)
		}
		posOb, ok := ob.(object.Positioner)
		if !ok {
			return 3, fmt.Sprintf("%s:object:%s#%s:not positioner",
				ObjectShow, ob.ID(), ob.Serial())
		}
		objects = append(objects, posOb)
	}
	out := ""
	for _, ob := range objects {
		x, y := ob.Position()
		out = fmt.Sprintf("%s%fx%f ", out, x, y)
	}
	out = strings.TrimSpace(out)
	return 0, out
}
