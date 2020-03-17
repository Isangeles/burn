/*
 * resshow.go
 *
 * Copyright 2020 Dariusz Sikora <dev@isangeles.pl>
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

	"github.com/isangeles/flame/data/res"
	"github.com/isangeles/flame/data/res/lang"
)

// resshow handles resshow command.
func resshow(cmd Command) (int, string) {
	if len(cmd.OptionArgs()[0]) < 1 {
		return 2, fmt.Sprintf("%s: no option args", ResShow)
	}
	switch cmd.OptionArgs()[0] {
	case "objects":
		out := ""
		for _, obd := range res.Objects() {
			out = fmt.Sprintf("%s%s ", out, obd.BasicData.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "dialogs":
		out := ""
		for _, dl := range res.Dialogs() {
			out = fmt.Sprintf("%s%s ", out, dl.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "quests":
		out := ""
		for _, qr := range res.Quests() {
			out = fmt.Sprintf("%s%s ", out, qr.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "armors":
		out := ""
		for _, ad := range res.Armors() {
			out = fmt.Sprintf("%s%s ", out, ad.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "weapons":
		out := ""
		for _, wd := range res.Weapons() {
			out = fmt.Sprintf("%s%s ", out, wd.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "miscs":
		out := ""
		for _, md := range res.MiscItems() {
			out = fmt.Sprintf("%s%s ", out, md.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "recipes":
		out := ""
		for _, rd := range res.Recipes() {
			out = fmt.Sprintf("%s%s ", out, rd.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "races":
		out := ""
		for _, rd := range res.Races() {
			out = fmt.Sprintf("%s%s ", out, rd.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "translations":
		out := ""
		for _, td := range res.Translations() {
			out = fmt.Sprintf("%s%s ", out, td.ID)
		}
		out = strings.TrimSpace(out)
		return 0, out
	case "lang-text":
		return resshowLangText(cmd)
	default:
		return 2, fmt.Sprintf("%s: no such option: %s", ResShow,
			cmd.OptionArgs()[0])
	}
}

// resshowLangText handles lang-text option for resshow.
func resshowLangText(cmd Command) (int, string) {
	if len(cmd.Args()) < 1 {
		return 3, fmt.Sprintf("%s: no enught args for: %s",
			ResShow, cmd.OptionArgs()[0])
	}
	out := fmt.Sprintf("%v", lang.Text(cmd.Args()[0]))
	return 0, out
}
