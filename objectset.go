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
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/effect"
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
	case "health", "hp":
		return objectsetHealth(cmd)
	case "mana":
		return objectsetMana(cmd)
	case "experience", "exp":
		return objectsetExperience(cmd)
	case "target":
		return objectsetTarget(cmd)
	case "destination", "dest":
		return objectsetDestination(cmd)
	case "position", "pos":
		return objectsetPosition(cmd)
	case "chat":
		return objectsetChat(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
}

// objectsetHealth handles health option for objectset.
func objectsetHealth(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	objects := make([]effect.Target, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		tar, ok := ob.(effect.Target)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not targetable",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, tar)
	}
	val, err := strconv.Atoi(cmd.Args()[0])
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.Args()[0])
	}
	for _, o := range objects {
		o.SetHealth(val)
	}
	return 0, ""
}

// objectsetExperience handles experience option for objectset.
func objectsetExperience(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	objects := make([]effect.Target, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		tar, ok := ob.(effect.Target)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not character",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, tar)
	}
	val, err := strconv.Atoi(cmd.Args()[0])
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.Args()[0])
	}
	for _, o := range objects {
		o.SetExperience(val)
	}
	return 0, ""
}

// objectsetMana handles mana option for objectset.
func objectsetMana(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	objects := make([]effect.Target, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		tar, ok := ob.(effect.Target)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not targetable",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, tar)
	} 
	val, err := strconv.Atoi(cmd.Args()[0])
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.Args()[0])
	}
	for _, o := range objects {
		o.SetMana(val)
	}
	return 0, ""
}

// objectsetTarget handles target option for objectset.
func objectsetTarget(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not charatcer",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, char)
	} 
	id, serial := argSerialID(cmd.Args()[0])
	ob := flame.Game().Module().Object(id, serial)
	if ob == nil {
		return 8, fmt.Sprintf("%s: object not found: %s",
			ObjectSet, cmd.Args()[1])
	}
	tar, ok := ob.(effect.Target)
	if !ok {
		return 3, fmt.Sprintf("%s: object: %s#%s: not targetable",
			ObjectSet, ob.ID(), ob.Serial())
	}
	for _, o := range objects {
		o.SetTarget(tar)
	}
	return 0, ""
}

// objectsetDestination handles destionation option for objectset.
func objectsetDestination(cmd Command) (int, string) {
	if len(cmd.Args()) < 2 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not charatcer",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, char)
	}
	x, err := strconv.ParseFloat(cmd.Args()[1], 64)
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.OptionArgs()[0])
	}
	y, err := strconv.ParseFloat(cmd.Args()[2], 64)
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.OptionArgs()[1])
	}
	for _, o := range objects {
		o.SetDestPoint(x, y)
	}
	return 0, ""
}

// objectsetPosition handles position option for objectset.
func objectsetPosition(cmd Command) (int, string) {
	if len(cmd.Args()) < 2 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	objects := make([]object.Positioner, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		posOb, ok := ob.(object.Positioner)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not positioner",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, posOb)
	}
	x, err := strconv.ParseFloat(cmd.Args()[0], 64)
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.OptionArgs()[0])
	}
	y, err := strconv.ParseFloat(cmd.Args()[1], 64)
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.OptionArgs()[1])
	}
	for _, ob := range objects {
		ob.SetPosition(x, y)
	}
	return 0, ""
}

// objectsetChat handles chat option for objectset.
func objectsetChat(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not charatcer",
				ObjectSet, ob.ID(), ob.Serial())
		}
		objects = append(objects, char)
	}
	for _, o := range objects {
		o.SendChat(cmd.Args()[1])
	}
	return 0, ""
}
