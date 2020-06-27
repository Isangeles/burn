/*
 * objectshow.go
 *
 * Copyright 2019-2020 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/module/character"
	"github.com/isangeles/flame/module/dialog"
	"github.com/isangeles/flame/module/effect"
	"github.com/isangeles/flame/module/flag"
	"github.com/isangeles/flame/module/item"
	"github.com/isangeles/flame/module/objects"
	"github.com/isangeles/flame/module/quest"
	"github.com/isangeles/flame/module/serial"
	"github.com/isangeles/flame/module/skill"
)

// objectshow handles objectshow command.
func objectshow(cmd Command) (int, string) {
	if Game == nil {
		return 2, fmt.Sprintf("%s: no game set", ObjectShow)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ObjectShow)
	}
	switch cmd.OptionArgs()[0] {
	case "id":
		return objectshowID(cmd)
	case "serial":
		return objectshowSerial(cmd)
	case "equipment", "eq":
		return objectshowEquipment(cmd)
	case "effects":
		return objectshowEffects(cmd)
	case "dialogs":
		return objectshowDialogs(cmd)
	case "quests":
		return objectshowQuests(cmd)
	case "flags":
		return objectshowFlags(cmd)
	case "recipes":
		return objectshowRecipes(cmd)
	case "position", "pos":
		return objectshowPosition(cmd)
	case "items":
		return objectshowItems(cmd)
	case "skills":
		return objectshowSkills(cmd)
	case "health", "hp":
		return objectshowHealth(cmd)
	case "max-health", "max-hp":
		return objectshowMaxHealth(cmd)
	case "mana":
		return objectshowMana(cmd)
	case "range":
		return objectshowRange(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s",
			ObjectShow, cmd.OptionArgs()[0])
	}
}

// objectshowID handles id option for objectshow.
func objectshowID(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]objects.Object, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obs = append(obs, ob)
	}
	out := ""
	for _, ob := range obs {
		out = fmt.Sprintf("%s%s ", out, ob.ID())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowSerial handles serial option for objectshow.
func objectshowSerial(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]objects.Object, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obs = append(obs, ob)
	}
	out := ""
	for _, ob := range obs {
		out = fmt.Sprintf("%s%s ", out, ob.Serial())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowEquipment handles equipment option for objectshow.
func objectshowEquipment(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: is not character",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, char)
	}
	out := ""
	for _, ob := range obs {
		for _, it := range ob.Equipment().Items() {
			out += fmt.Sprintf("%s#%s:", it.ID(), it.Serial())
			for _, s := range it.Slots() {
				out += fmt.Sprintf("%s ", s)
			}
			out += "\n"
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowEffects handles effects option for objectshow.
func objectshowEffects(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]effect.Target, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		tar, ok := ob.(effect.Target)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no effects ",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, tar)
	}
	out := ""
	for _, o := range obs {
		for _, e := range o.Effects() {
			out += fmt.Sprintf("%s#%s ", e.ID(), e.Serial())
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowDialogs handles dialogs option for objectshow.
func objectshowDialogs(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]dialog.Talker, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		talker, ok := ob.(dialog.Talker)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no dialogs",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, talker)
	}
	out := ""
	for _, o := range obs {
		for _, d := range o.Dialogs() {
			out += fmt.Sprintf("%s ", d.ID())
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowQuests handles quests option for objectshow.
func objectshowQuests(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]quest.Quester, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		quester, ok := ob.(quest.Quester)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no quests",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, quester)
	}
	out := ""
	for _, o := range obs {
		for _, q := range o.Journal().Quests() {
			out += fmt.Sprintf("%s ", q.ID())
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowFlags handles flags option for objectshow.
func objectshowFlags(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]flag.Flagger, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		flagger, ok := ob.(flag.Flagger)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no flags",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, flagger)
	}
	out := ""
	for _, o := range obs {
		for _, f := range o.Flags() {
			out += fmt.Sprintf("%s ", f.ID())
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowRecipes handles recipes option for objectshow.
func objectshowRecipes(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: is not character",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, char)
	}
	out := ""
	for _, o := range obs {
		for _, r := range o.Crafting().Recipes() {
			out += fmt.Sprintf("%s ", r.ID())
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowPosition handles position option for objectshow.
func objectshowPosition(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]objects.Positioner, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		posOb, ok := ob.(objects.Positioner)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not positioner",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, posOb)
	}
	out := ""
	for _, ob := range obs {
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
	obs := make([]item.Container, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s", ObjectAdd, arg)
		}
		con, ok := ob.(item.Container)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not container",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		obs = append(obs, con)
	}
	out := ""
	for _, ob := range obs {
		for _, it := range ob.Inventory().Items() {
			out = fmt.Sprintf("%s%s#%s ", out, it.ID(), it.Serial())
		}
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowSkills handles items option for objectshow.
func objectshowSkills(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectShow)
	}
	obs := make([]skill.User, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s", ObjectAdd, arg)
		}
		user, ok := ob.(skill.User)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: have no skills",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, user)
	}
	out := ""
	for _, ob := range obs {
		for _, s := range ob.Skills() {
			out = fmt.Sprintf("%s%s ", out, s.ID())
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
	obs := make([]objects.Killable, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obHP, ok := ob.(objects.Killable)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not killable",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, obHP)
	}
	out := ""
	for _, ob := range obs {
		out = fmt.Sprintf("%s%d ", out, ob.Health())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowMaxHealth handles max-health option for objectshow.
func objectshowMaxHealth(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", ObjectShow)
	}
	obs := make([]objects.Killable, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obHP, ok := ob.(objects.Killable)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: not killable",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, obHP)
	}
	out := ""
	for _, ob := range obs {
		out = fmt.Sprintf("%s%d ", out, ob.MaxHealth())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowMana handles mana option for objectshow.
func objectshowMana(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", ObjectShow)
	}
	obs := make([]objects.Magician, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obMana, ok := ob.(objects.Magician)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no mana",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, obMana)
	}
	out := ""
	for _, ob := range obs {
		out = fmt.Sprintf("%s%d ", out, ob.Mana())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// objectshowMaxMana handles max-mana option for objectshow.
func objectshowMaxMana(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s:no target args", ObjectShow)
	}
	obs := make([]objects.Magician, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obMana, ok := ob.(objects.Magician)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no mana",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, obMana)
	}
	out := ""
	for _, ob := range obs {
		out = fmt.Sprintf("%s%d ", out, ob.MaxMana())
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
	obs := make([]objects.Positioner, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectShow, arg)
		}
		obPos, ok := ob.(objects.Positioner)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no position",
				ObjectShow, ob.ID(), ob.Serial())
		}
		obs = append(obs, obPos)
	}
	id, ser := argSerialID(cmd.Args()[0])
	tar := serial.Object(id, ser)
	if tar == nil {
		return 3, fmt.Sprintf("%s: object not found: %s",
			ObjectShow, cmd.Args()[0])
	}
	tarPos, ok := tar.(objects.Positioner)
	if !ok {
		return 3, fmt.Sprintf("%s: target: %s#%s: no position",
			ObjectShow, tar.ID(), tar.Serial())
	}
	out := ""
	for _, ob := range obs {
		out += fmt.Sprintf("%f ", objects.Range(ob, tarPos))
	}
	out = strings.TrimSpace(out)
	return 0, out
}
