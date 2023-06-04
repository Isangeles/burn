/*
 * chaptershow.go
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
)

// chaptershow handles chaptershow command.
func chaptershow(cmd Command) (int, string) {
	if Module == nil {
		return 2, fmt.Sprintf("%s: no module set", ChapterShow)
	}
	if Module.Chapter() == nil {
		return 2, fmt.Sprintf("%s: no active chapter", ChapterShow)
	}
	if cmd.OptionArgs() == nil || len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ChapterShow)
	}
	switch cmd.OptionArgs()[0] {
	case "lang-path":
		return 0, Module.Chapter().Conf().LangPath()
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", ChapterShow,
			cmd.OptionArgs()[0])
	}
}
