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
	"github.com/isangeles/flame/core/module/flag"
	"github.com/isangeles/flame/core/module/object/character"
	"github.com/isangeles/flame/core/module/object/effect"
	"github.com/isangeles/flame/core/module/object/item"
	"github.com/isangeles/flame/core/module/object/quest"
	"github.com/isangeles/flame/core/module/object/skill"
)

// objectadd handles objectadd command.
func objectadd(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s: no active game", ObjectAdd)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ObjectAdd)
	}
	switch cmd.OptionArgs()[0] {
	case "item":
		return objectaddItem(cmd)
	case "flag":
		return objectaddFlag(cmd)
	case "effect":
		return objectaddEffect(cmd)
	case "skill":
		return objectaddSkill(cmd)
	case "quest":
		return objectaddQuest(cmd)
	case "recipe":
		return objectaddRecipe(cmd)
	case "equipment", "eq":
		return objectaddEquipment(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
}

// objectaddItem handles item option for objectadd.
func objectaddItem(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectAdd, cmd.OptionArgs()[0])
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
			return 3, fmt.Sprintf("%s: object: %s#%s: no inventory",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, con)
	}
	id := cmd.Args()[0]
	item, err := data.Item(id)
	if err != nil {
		return 3, fmt.Sprintf("%s: fail to retrieve item: %v", ObjectAdd, err)
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

// objectaddFlag handles add option for objectadd.
func objectaddFlag(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
	objects := make([]flag.Flagger, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectAdd, arg)
		}
		flagger, ok := ob.(flag.Flagger)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no flags",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, flagger)
	}
	for _, ob := range objects {
		ob.AddFlag(flag.Flag(cmd.Args()[0]))
	}
	return 0, ""
}

// objectaddEffect handles add option for objectadd.
func objectaddEffect(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
	objects := make([]effect.Target, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectAdd, arg)
		}
		tar, ok := ob.(effect.Target)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no effects",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, tar)
	}
	effectID := cmd.Args()[0]
	for _, ob := range objects {
		effect, err := data.Effect(flame.Mod(), effectID)
		if err != nil {
			return 3, fmt.Sprintf("%s: fail to retrieve effect: %v", ObjectAdd, err)
		}
		ob.AddEffect(effect)
	}
	return 0, ""
}

// objectaddSkill handles skill option for objectadd.
func objectaddSkill(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
	objects := make([]skill.SkillUser, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectAdd, arg)
		}
		user, ok := ob.(skill.SkillUser)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no skills",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, user)
	}
	skillID := cmd.Args()[0]
	for _, ob := range objects {
		skill, err := data.Skill(skillID)
		if err != nil {
			return 3, fmt.Sprintf("%s: fail to retrieve skill: %v", ObjectAdd, err)
		}
		ob.AddSkill(skill)
	}
	return 0, ""
}

// objectaddQuest handles quest option for objectadd.
func objectaddQuest(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
	objects := make([]quest.Quester, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectAdd, arg)
		}
		quester, ok := ob.(quest.Quester)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: no quests",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, quester)
	}
	questID := cmd.Args()[0]
	for _, ob := range objects {
		quest, err := data.Quest(questID)
		if err != nil {
			return 3, fmt.Sprintf("%s: fail to retrieve quest: %v", ObjectAdd, err)
		}
		ob.Journal().AddQuest(quest)
	}
	return 0, ""
}

// objectaddRecipe handles recipe option for objectadd.
func objectaddRecipe(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectAdd, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: is not character",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, char)
	}
	recipeID := cmd.Args()[0]
	for _, ob := range objects {
		recipe, err := data.Recipe(recipeID)
		if err != nil {
			return 3, fmt.Sprintf("%s: fail to retrieve recipe: %v", ObjectAdd, err)
		}
		ob.AddRecipe(recipe)
	}
	return 0, ""
}

// objectaddEquipment handles equipment option for objectadd.
func objectaddEquipment(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no target args", ObjectAdd)
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectAdd, cmd.OptionArgs()[0])
	}
	objects := make([]*character.Character, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s",
				ObjectAdd, arg)
		}
		char, ok := ob.(*character.Character)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: is not character",
				ObjectAdd, ob.ID(), ob.Serial())
		}
		objects = append(objects, char)
	}
	switch cmd.Args()[0] {
	case "hand-right":
		for _, o := range objects {
			id, serial := argSerialID(cmd.Args()[1])
			it := o.Inventory().Item(id, serial)
			if it == nil {
				return 3, fmt.Sprintf("%s: %s#%s: fail to retrieve item from inventory: %s#%s",
					ObjectAdd, o.ID(), o.Serial(), id, serial)
			}
			eit, ok := it.(item.Equiper)
			if !ok {
				return 3, fmt.Sprintf("%s: %s#%s: item not equipable: %s#%s",
					ObjectAdd, o.ID(), o.Serial(), id, serial)
			}
			for _, s := range o.Equipment().Slots() {
				if s.Type() == character.Hand_right {
					break
				}
				s.SetItem(eit)
			}
		}
		return 0, ""
	default:
		return 3, fmt.Sprintf("%s: no vaild target for %s: '%s'", ObjectAdd,
			cmd.OptionArgs()[0], cmd.Args()[0])
	}

}
