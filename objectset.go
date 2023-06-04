/*
 * objectset.go
 *
 * Copyright 2019-2023 Dariusz Sikora <ds@isangeles.dev>
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

	"github.com/isangeles/flame/area"
	"github.com/isangeles/flame/character"
	"github.com/isangeles/flame/effect"
	"github.com/isangeles/flame/objects"
	"github.com/isangeles/flame/serial"
)

// objectset handles objectset command.
func objectset(cmd Command) (int, string) {
	if Module == nil {
		return 2, fmt.Sprintf("%s: no game set", ObjectSet)
	}
	if cmd.OptionArgs() == nil || len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ObjectSet)
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
	case "area":
		return objectsetArea(cmd)
	case "guild":
		return objectsetGuild(cmd)
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
	obs := make([]objects.Killable, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		tar, ok := ob.(objects.Killable)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not targetable",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, tar)
	}
	val, err := strconv.Atoi(cmd.Args()[0])
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.Args()[0])
	}
	for _, o := range obs {
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
	obs := make([]objects.Experiencer, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		tar, ok := ob.(objects.Experiencer)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not character",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, tar)
	}
	val, err := strconv.Atoi(cmd.Args()[0])
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.Args()[0])
	}
	for _, o := range obs {
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
	obs := make([]objects.Magician, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		tar, ok := ob.(objects.Magician)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not targetable",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, tar)
	}
	val, err := strconv.Atoi(cmd.Args()[0])
	if err != nil {
		return 3, fmt.Sprintf("%s: invalid argument: %s", ObjectSet,
			cmd.Args()[0])
	}
	for _, o := range obs {
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
	obs := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not charatcer",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, char)
	}
	id, ser := argSerialID(cmd.Args()[0])
	ob := serial.Object(id, ser)
	if ob == nil {
		return 8, fmt.Sprintf("%s: object not found: %s",
			ObjectSet, cmd.Args()[1])
	}
	tar, ok := ob.(effect.Target)
	if !ok {
		return 3, fmt.Sprintf("%s: object: %s#%s: not targetable",
			ObjectSet, ob.ID(), ob.Serial())
	}
	for _, o := range obs {
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
	obs := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not charatcer",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, char)
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
	for _, o := range obs {
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
	obs := make([]objects.Positioner, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		posOb, ok := ob.(objects.Positioner)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not positioner",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, posOb)
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
	for _, ob := range obs {
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
	obs := make([]objects.Logger, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(objects.Logger)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not charatcer",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, char)
	}
	for _, o := range obs {
		msg := objects.NewMessage(cmd.Args()[0], true)
		o.ChatLog().Add(msg)
	}
	return 0, ""
}

// objectsetArea handles area option for objectset.
func objectsetArea(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	obs := make([]area.Object, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(area.Object)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not area object",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, char)
	}
	for _, o := range obs {
		o.SetAreaID(cmd.Args()[0])
	}
	return 0, ""
}

// objectsetGuild handles guild option for objectset.
func objectsetGuild(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectSet, cmd.OptionArgs()[0])
	}
	obs := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectSet, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not a character",
				ObjectSet, ob.ID(), ob.Serial())
		}
		obs = append(obs, char)
	}
	for _, o := range obs {
		o.SetGuild(character.NewGuild(cmd.Args()[0]))
	}
	return 0, ""
}
