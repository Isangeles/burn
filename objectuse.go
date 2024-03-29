/*
 * objectuse.go
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

	"github.com/isangeles/flame/serial"
	"github.com/isangeles/flame/skill"
)

// objectuse handles objectuse command.
func objectuse(cmd Command) (int, string) {
	if Module == nil {
		return 2, fmt.Sprintf("%s: no game set", ObjectUse)
	}
	if cmd.OptionArgs() == nil || len(cmd.OptionArgs()[0]) < 1 {
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
	objects := make([]skill.User, 0)
	for _, arg := range cmd.TargetArgs() {
		id, ser := argSerialID(arg)
		ob := serial.Object(id, ser)
		if ob == nil {
			return 3, fmt.Sprintf("%s: object not found: %s#%s",
				ObjectUse, id, ser)
		}
		user, ok := ob.(skill.User)
		if !ok {
			return 3, fmt.Sprintf("%s: object: %s#%s: have no skills",
				ObjectUse, ob.ID(), ob.Serial())
		}
		objects = append(objects, user)
	}
	for _, o := range objects {
		id := cmd.Args()[0]
		var skill *skill.Skill
		for _, s := range o.Skills() {
			if s.ID() == id {
				skill = s
			}
		}
		if skill == nil {
			return 3, fmt.Sprintf("%s: object: skill not known: %s",
				ObjectUse, id)
		}
		o.Use(skill)
	}
	return 0, ""
}
