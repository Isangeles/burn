/*
 * objectremove.go
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
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/core/module/object/quest"
	"github.com/isangeles/flame/core/module/object/skill"
)

// objecteremove handles objectremove command.
func objectremove(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s: no active game", ObjectRemove)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ObjectRemove)
	}
	switch cmd.OptionArgs()[0] {
	case "item":
		return objectremoveItem(cmd)
	case "effect":
		return objectremoveEffect(cmd)
	case "skill":
		return objectremoveSkill(cmd)
	case "quest":
		return objectremoveQuest(cmd)
	case "flag":
		return objectremoveFlag(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s",
			ObjectRemove, cmd.OptionArgs()[0])
	}
}

// objectremoveItem handles item option for objectremove.
func objectremoveItem(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectRemove)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectRemove, cmd.OptionArgs()[0])
	}
	objects := make([]item.Container, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s", ObjectRemove, arg)
		}
		con, ok := ob.(item.Container)
		if !ok {
			return 3, fmt.Sprintf("%s: %s#%s: no inventory", ObjectRemove,
				ob.ID(), ob.Serial())
		}
		objects = append(objects, con)
	}
	itemID, itemSerial := argSerialID(cmd.Args()[0])
	for i, ob := range objects {
		var item item.Item
		for _, it := range ob.Inventory().Items() {
			if it.ID() != itemID || it.Serial() != itemSerial {
				continue
			}
			item = it
			break
		}
		if item == nil {
			return 3, fmt.Sprintf("%s: object %d: have no item: %s#%s",
				ObjectRemove, i, itemID, itemSerial)
		}
		ob.Inventory().RemoveItem(item)
	}
	return 0, ""
}

// objectremoveEffect handles effect option for objectremove.
func objectremoveEffect(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectRemove)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectRemove, cmd.OptionArgs()[0])
	}
	objects := make([]effect.Target, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectRemove, arg)
		}
		tar, ok := ob.(effect.Target)
		if !ok {
			return 3, fmt.Sprintf("%s: %s#%s: no effects",
				ObjectRemove, ob.ID(), ob.Serial())
		}
		objects = append(objects, tar)
	}
	effectID, effectSerial := argSerialID(cmd.Args()[0])
	for i, ob := range objects {
		var eff *effect.Effect
		for _, e := range ob.Effects() {
			if effectID != e.ID() || effectSerial != e.Serial() {
				continue
			}
			eff = e
			break
		}
		if eff == nil {
			return 3, fmt.Sprintf("%s: object %d: have no effect: %s#%s", ObjectRemove,
				i, effectID, effectSerial)
		}
		ob.RemoveEffect(eff)
	}
	return 0, ""
}

// objectremovewSkill handles skill option for objectremove.
func objectremoveSkill(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectRemove)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectRemove, cmd.OptionArgs()[0])
	}
	objects := make([]skill.SkillUser, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectRemove, arg)
		}
		user, ok := ob.(skill.SkillUser)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no skills",
				ObjectRemove, ob.ID(), ob.Serial())
		}
		objects = append(objects, user)
	}
	skillID, skillSerial := argSerialID(cmd.Args()[0])
	for i, ob := range objects {
		var skill *skill.Skill
		for _, s := range ob.Skills() {
			if s.ID() != skillID || s.Serial() != skillSerial {
				continue
			}
			skill = s
			break
		}
		if skill == nil {
			return 3, fmt.Sprintf("%s: object %d: have no skill: %s#s",
				ObjectRemove, i, skillID, skillSerial)
		}
		ob.RemoveSkill(skill)
	}
	return 0, ""
}

// objectremoveQuest handles quest option for objectremove.
func objectremoveQuest(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectRemove)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectRemove, cmd.OptionArgs()[0])
	}
	objects := make([]quest.Quester, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectRemove, arg)
		}
		quester, ok := ob.(quest.Quester)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no quests",
				ObjectRemove, ob.ID(), ob.Serial())
		}
		objects = append(objects, quester)
	}
	questID := cmd.Args()[0]
	for i, ob := range objects {
		var quest *quest.Quest
		for _, q := range ob.Journal().Quests() {
			if q.ID() != questID {
				continue
			}
			quest = q
			break
		}
		if quest == nil {
			return 3, fmt.Sprintf("%s: object %d: have no quest: %s",
				ObjectRemove, i, questID)
		}
		ob.Journal().RemoveQuest(quest)
	}
	return 0, ""
}

// objectremoveFlag handles flag option for objectremove.
func objectremoveFlag(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectRemove)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectRemove, cmd.OptionArgs()[0])
	}
	objects := make([]flag.Flagger, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectRemove, arg)
		}
		flagger, ok := ob.(flag.Flagger)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no flags",
				ObjectRemove, ob.ID(), ob.Serial())
		}
		objects = append(objects, flagger)
	}
	for _, ob := range objects {
		ob.RemoveFlag(flag.Flag(cmd.Args()[0]))
	}
	return 0, ""
}
