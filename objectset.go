/*
 * objectset.go
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
	"github.com/isangeles/flame/core/module/object"
)

// objectset handles objectset command.
func objectset(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s:no active game", ObjectSet)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s:no option args", ObjectSet)
	}
	switch cmd.OptionArgs()[0] {
	case "position", "pos":
		return objectsetPosition(cmd)
	default:
		return 2, fmt.Sprintf("%s:no_such_option:%s",
			ObjectSet, cmd.OptionArgs()[0])
	}
}

// objectsetPosition handles position option for objectset.
func objectsetPosition(cmd Command) (int, string) {
	if len(cmd.Args()) < 2 {
		return 3, fmt.Sprintf("%s:no_enought_args_for:%s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	objects := make([]object.Positioner, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s:object_not_found:%s",
				ObjectSet, arg)
		}
		posOb, ok := ob.(object.Positioner)
		if !ok {
			return 3, fmt.Sprintf("%s:object:%s#%s:not positioner",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, posOb)
	}
	x, err := strconv.ParseFloat(cmd.Args()[0], 64)
	if err != nil {
		return 3, fmt.Sprintf("%s:invalid_argument:%s", ObjectSet,
			cmd.OptionArgs()[0])
	}
	y, err := strconv.ParseFloat(cmd.Args()[1], 64)
	if err != nil {
		return 3, fmt.Sprintf("%s:invalid_argument:%s", ObjectSet,
			cmd.OptionArgs()[1])
	}
	for _, ob := range objects {
		ob.SetPosition(x, y)
	}
	return 0, ""
}
