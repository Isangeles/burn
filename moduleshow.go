/*
 * moduleshow.go
 *
 * Copyright 2019-2021 Dariusz Sikora <dev@isangeles.pl>
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
)

// moduleshow handles moduleshow command.
func moduleshow(cmd Command) (int, string) {
	if Module == nil {
		return 2, fmt.Sprintf("%s: no module set", ModuleShow)
	}
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ModuleShow)
	}
	switch cmd.OptionArgs()[0] {
	case "id":
		return 0, Module.Conf().ID
	case "areas":
		return moduleshowAreas(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", ModuleShow,
			cmd.OptionArgs()[0])
	}
}


// moduleshowAreas handles areas option for moduleshow.
func moduleshowAreas(cmd Command) (int, string) {
	chapter := Module.Chapter()
	if chapter == nil {
		return 3, fmt.Sprintf("%s: no active chapter",
			ModuleShow)
	}
	areas := chapter.Areas()
	for _, a := range areas {
		areas = append(areas, a.AllSubareas()...)
	}
	out := ""
	for _, a := range areas {
		out = fmt.Sprintf("%s%s ", out, a.ID())
	}
	out = strings.TrimSpace(out)
	return 0, out
}
