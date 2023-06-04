/*
 * engineshow.go
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

	"github.com/isangeles/flame"
)

// engineshow handles engineshow command.
func engineshow(cmd Command) (int, string) {
	if cmd.OptionArgs() == nil || len(cmd.OptionArgs()) < 1 {
		return 2, fmt.Sprintf("%s: no option args", EngineShow)
	}
	switch cmd.OptionArgs()[0] {
	case "version":
		return 0, flame.Version
	case "echo":
		if len(cmd.Args()) < 1 {
			return 3, fmt.Sprintf("%s: no args", EngineShow)
		}
		return 0, cmd.Args()[0]
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", EngineShow,
			cmd.OptionArgs()[0])
	}
}
