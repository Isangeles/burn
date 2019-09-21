/*
 * objectuse.go
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
	"github.com/isangeles/flame/core/module/object/skill"
)

// objectuse handles objectuse command.
func objectuse(cmd Command) (int, string) {
	if flame.Game() == nil {
		return 2, fmt.Sprintf("%s: no active game", ObjectUse)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ObjectUse)
	}
	switch cmd.OptionArgs()[0] {
	case "skill":
		return objectuseSkill(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s",
			ObjectUse, cmd.OptionArgs()[0])
	}
}

// objectuseSkill handles skill option for objectuse.
func objectuseSkill(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought target args for: %s",
			ObjectUse, cmd.OptionArgs()[0])
	}
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			ObjectUse, cmd.OptionArgs()[0])
	}
	objects := make([]skill.SkillUser, 0)
	for _, arg := range cmd.TargetArgs() {
		id, serial := argSerialID(arg)
		ob := flame.Game().Module().Object(id, serial)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s#%s",
				ObjectUse, id, serial)
		}
		user, ok := ob.(skill.SkillUser)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: have no skills",
				ObjectUse, ob.ID(), ob.Serial())
		}
		objects = append(objects, user)
	}
	for _, o := range objects {
		id, serial := argSerialID(cmd.Args()[0])
		var skill *skill.Skill
		for _, s := range o.Skills() {
			if s.ID() == id && s.Serial() == serial {
				skill = s
			}
		}
		if skill == nil {
			return 3, fmt.Sprintf("%s: object: skill not known: %s#%s",
				ObjectUse, id, serial)
		}
		o.UseSkill(skill)
	}
	return 0, ""
}
