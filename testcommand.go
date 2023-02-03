/*
 * testcommand.go
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

import "fmt"

// Simple command struct for testing.
type testCommand struct {
	tool                         string
	targetArgs, optionArgs, args []string
}

// Tool returns command tool.
func (tc testCommand) Tool() string {
	return tc.tool
}

// TargetArgs returns command target args.
func (tc testCommand) TargetArgs() []string {
	return tc.targetArgs
}

// OptionArgs returns command option args.
func (tc testCommand) OptionArgs() []string {
	return tc.optionArgs
}

// Args returns command args.
func (tc testCommand) Args() []string {
	return tc.args
}

// AddArgs adds specified arguments to command.
func (tc testCommand) AddArgs(args ...string) {
	tc.args = append(tc.args, args...)
}

// AddTargetArgs adds specified target arguments to command.
func (tc testCommand) AddTargetArgs(args ...string) {
	tc.targetArgs = append(tc.targetArgs, args...)
}

// String returns text for command.
func (tc testCommand) String() string {
	return fmt.Sprintf("Tool: %s, Target args: %v, Option args: %v, Args: %v",
		tc.tool, tc.targetArgs, tc.optionArgs, tc.args)
}

