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
	"github.com/isangeles/flame/core/module/object/item"
)

// objectshow handles objectshow command.
func objectshow(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s: no active game", ObjectShow)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ObjectShow)
	}
	switch cmd.OptionArgs()[0] {
	case "position", "pos":
		return objectshowPosition(cmd)
	case "items":
		return objectshowItems(cmd)
	case "health", "hp":
		return objectshowHealth(cmd)
	case "range":
		return objectshowRange(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s",
			ObjectShow, cmd.OptionArgs()[0])
	}
}

// objectshowPosition handles position option for objectshow.
func objectshowPosition(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	objects := make([]object.Positioner, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		posOb, ok := ob.(object.Positioner)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not positioner",
				ObjectShow, ob.ID(), ob.Serial())
		}
		objects = append(objects, posOb)
	}
	out := ""
	for _, ob := range objects {
		x, y := ob.Position()
		out = fmt.Sprintf("%s%f %f ", out, x, y)
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowItems handles items option for objectshow.
func objectshowItems(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	objects := make([]item.Container, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s", ObjectAdd, arg)
		}
		con, ok := ob.(item.Container)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not container",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, con)
	}
	out := ""
	for _, ob := range objects {
		for _, it := range ob.Inventory().Items() {
			out = fmt.Sprintf("%s%s#%s ", out, it.ID(), it.Serial())
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowHealth handles health option for objectshow.
func objectshowHealth(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", ObjectShow)
	}
	objects := make([]object.Killable, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obHP, ok := ob.(object.Killable)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not killable",
				ObjectShow, ob.ID(), ob.Serial())
		}
		objects = append(objects, obHP)
	}
	out := ""
	for _, ob := range objects {
		out = fmt.Sprintf("%s%d ", out, ob.Health())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowRange handles range option for objectshow.
func objectshowRange(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", ObjectShow)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no args", ObjectShow)
	}
	objects := make([]object.Positioner, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obPos, ok := ob.(object.Positioner)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no position",
				ObjectShow, ob.ID(), ob.Serial())
		}
		objects = append(objects, obPos)
	}
	id, serial := argSerialID(cmd.Args()[0])
	tar := flame.Game().Module().Object(id, serial)
	if tar == nil {
		return 3, fmt.Sprintf("%s: object not found: %s",
			ObjectShow, cmd.Args()[0])
	}
	tarPos, ok := tar.(object.Positioner)
	if !ok {
		return 3, fmt.Sprintf("%s: target: %s#%s: no position",
			ObjectShow, tar.ID(), tar.Serial())
	}
	out := ""
	for _, ob := range objects {
		out += fmt.Sprintf("%f ", object.Range(ob, tarPos))
	}
	out = strings.TrimSpace(out)
	return 0, out
}
