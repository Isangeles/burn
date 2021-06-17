/*
 * areashow.go
 *
 * Copyright 2021 Dariusz Sikora <dev@isangeles.pl>
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
	"time"

	"github.com/isangeles/flame/area"
)

// areashow handles areashow command.
func areashow(cmd Command) (int, string) {
	if Module == nil {
		return 2, fmt.Sprintf("%s: no module set", AreaShow)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", AreaShow)
	}
	switch cmd.OptionArgs()[0] {
	case "chars":
		return areashowChars(cmd)
	case "objects":
		return areashowObjects(cmd)
	case "time":
		return areashowTime(cmd)
	case "weather":
		return areashowWeather(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", AreaShow,
			cmd.OptionArgs()[0])
	}
}

// areashowChars handles chars option for areashow.
func areashowChars(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought target args for: %s",
			AreaShow, cmd.OptionArgs()[0])
	}
	areas := Module.Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.TargetArgs()[0]
	var area *area.Area
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		area = a
		break
	}
	if area == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			AreaShow, areaID)
	}
	out := ""
	for _, c := range area.Characters() {
		out = fmt.Sprintf("%s%s%s%s ", out, c.ID(),
			IDSerialSep, c.Serial())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// areashowObjects handles objects option for areashow.
func areashowObjects(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			AreaShow, cmd.Args()[0])
	}
	areas := Module.Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.TargetArgs()[0]
	var area *area.Area
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		area = a
		break
	}
	if area == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			AreaShow, areaID)
	}
	out := ""
	for _, o := range area.Objects() {
		out = fmt.Sprintf("%s%s%s%s ", out, o.ID(),
			IDSerialSep, o.Serial())
	}
	out = strings.TrimSpace(out)
	return 0, out
}

// areashowTime handles time option for areashow.
func areashowTime(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			AreaShow, cmd.Args()[0])
	}
	areas := Module.Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.TargetArgs()[0]
	var area *area.Area
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		area = a
		break
	}
	if area == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			AreaShow, areaID)
	}
	return 0, area.Time().Format(time.Kitchen)
}

// areashowWeather handles weather option for areashow.
func areashowWeather(cmd Command) (int, string) {
	if len(cmd.TargetArgs()) < 1 {
		return 3, fmt.Sprintf("%s: no enought args for: %s",
			AreaShow, cmd.Args()[0])
	}
	areas := Module.Chapter().Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	areaID := cmd.TargetArgs()[0]
	var area *area.Area
	for _, a := range areas {
		if a.ID() != areaID {
			continue
		}
		area = a
		break
	}
	if area == nil {
		return 3, fmt.Sprintf("%s: area not found: %s",
			AreaShow, areaID)
	}
	return 0, fmt.Sprintf("%s", area.Weather().Conditions())
}
