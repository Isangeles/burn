/*
 * objectadd.go
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
	"github.com/isangeles/flame/core/module/object/item"
)

// objectadd handles objectadd command.
func objectadd(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s:no active game", ObjectAdd)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s:no option args", ObjectAdd)
	}
	switch cmd.OptionArgs()[0] {
	case "item":
		return objectaddItem(cmd)
	default:
		return 2, fmt.Sprintf("%s:no_such_option:%s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
}

// objectaddItem handles item option for objectadd.
func objectaddItem(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", ObjectAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s:no enought args for: %s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
	objects := make([]item.Container, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s:object not found: %s", ObjectAdd, arg)
		}
    con, ok := ob.(item.Container)
    if !ok {
      return 3, fmt.Sprintf("%s: object: %s#%s: no inventory",
        ObjectAdd, ob.ID(), ob.Serial())
    }
		objects = append(objects, con)
	}
	id := cmd.Args()[0]
	item, err := data.Item(id)
	if err != nil {
		return 3, fmt.Sprintf("%s:fail to retrieve item: %v", ObjectAdd, err)
	}
	for _, ob := range objects {
		err = ob.Inventory().AddItem(item)
		if err != nil {
			return 3, fmt.Sprintf("%s: fail to add item: %v",
				ObjectAdd, err)
		}
	}
	return 0, ""
}
