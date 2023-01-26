/*
 * areaset.go
 *
 * Copyright 2023 Dariusz Sikora <ds@isangeles.dev>
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
	"time"

	"github.com/isangeles/flame/area"
)

// areaset handles areaset command.
func areaset(cmd Command) (int, string) {
	if Module == nil {
		return 2, fmt.Sprintf("%s: no module set", AreaSet)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", AreaSet)
	}
	switch cmd.OptionArgs()[0] {
	case "weather":
		return areasetWeather(cmd)
	case "time":
		return areasetTime(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", AreaSet,
			cmd.OptionArgs()[0])
	}
}

// areasetWeather handles weather option for areaset.
func areasetWeather(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			AreaSet, cmd.Args()[0])
	}
	areas := Module.Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.TargetArgs()[0]
	var ar *area.Area
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		ar = a
		break
	}
	if ar == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			AreaSet, areaID)
	}
	ar.Weather().Conditions = area.Conditions(cmd.Args()[0])
	return 0, ""
}

// areasetTime handles time option for areaset.
func areasetTime(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			AreaSet, cmd.Args()[0])
	}
	areas := Module.Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.TargetArgs()[0]
	var ar *area.Area
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		ar = a
		break
	}
	if ar == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			AreaSet, areaID)
	}
	t, err := time.Parse(time.Kitchen, cmd.Args()[0])
	if err != nil {
		return 3, fmt.Sprintf("%s: unable to parse time: %v",
			AreaSet, err)
	}
	ar.Time = t
	return 0, ""
}
