## Introduction
Burn is a command interpreter for the [Flame](https://github.com/Isangeles/flame) engine.

Burn enables easy testing and debugging of Flame games and serves as a base for [Ash](https://github.com/Isangeles/burn/tree/master/ash) scripting language.

## Syntax
  ### Commands
  Standard Burn command syntax:
```
  [tool name] -o [option] -t [targets...] -a [arguments...]
```
  Beside many arguments Burn handles also many targets, e.g. one command can be executed on many game characters, which results with a combined output.

  Example command:
```
  objectshow -o position -t player_test#0 player_test#1
```
  Example output:
```
  0.0 0.0 420.0 100.0
```
  Shows positions of game objects with serial IDs 'player_test#0' and 'player_test#1'.

  To push argument with multiple word to command use quotes:
```
  objectset -o chat -t testchar#0 -a 'hey you!'
```

  ### Expressions
  Commands can also be joined into expressions.

  Target pipe expression:
```
  [command] |t [command]
```
  Executes the first command and uses the output as target arguments(-t) to execute next command.

  Example expression:
```
  moduleshow -o area-chars -t area1_test |t objectshow -o position
```
  Example output:
```
  0.0 0.0 12.0 131.0 130.0 201.0
```
  Shows positions of all game characters in the area with ID 'area1_test'.

  Argument pipe expression:
```
  [command] |a [command]
```
  Executes the first command and uses the output as arguments(-a) to execute next command.
  
  Example expression:
```
  objectshow -o pos -t player_a#1 |a gameadd -o char -a testchar area1
```
  Description: spawns new character with ID 'testchar' in area with ID 'area1' on position
  of existing character with serial ID 'player_a#1'.

## Ash
[Ash](https://github.com/Isangeles/burn/tree/master/ash) is scripting language that allows running Burn commands under conditional loop.

## Commands
Set target:
```
  objectset -o target -t [character ID]#[character serial] -a [target ID]#[target serial]
```
Description: sets object with specified serial ID(-a) as target of character with specified serial ID(-t).

Export game character:
```
  engineexport -o char -t [char ID]#[char serial]
```
Description: exports game character with specified ID to XML file in
data/modules/[module]/characters directory.

Load module:
```
  engineimport -t module -a [module name] [module path](optional)
```
Description: loads module with the specified name(module directory name) and with a specified path,
if no path provided, the engine will search default modules directory(data/modules).

Save game:
```
  engineexport -t game -a [save file name]
```
Description: saves current game to 'savegames/[module]' directory.

Add item:
```
  objectadd -o item -a [item ID] -t [character ID]#[character serial]
```
Description: adds item with specified ID to inventory of game character with specified serial ID.

Equip item:
```
  objectadd -o equipment -t [character ID]#[character serial] -a [slot ID] [item serial ID]
```
Description: equips item with specified ID for game character with specified serial ID.

Add effect:
```
  objectadd -o effect -t [character ID]#[character serial] -a [effect ID]
```
Description: puts effect with specified ID on game character with specified serial ID

Spawn NPC:
```
  gameadd -o character -a [character ID] [areaID] [posX](optional) [posY](optional)
```
Description: spawns new chapter NPC with specified ID in specified area at given position(0, 0 if not specified).

Add character to scenario area:
```
  gameadd -o area-char -t [character ID]#[character serial] -a [areaID]
```
Description: add character with specified serial ID to specified area.

Show translation text for specified ID:
```
  resshow -o lang-text -a [id]
```
Description: shows translation text for specified ID.

Show area objects data:
```
  resshow -o objects 
```
Description: shows IDs of all loaded objects data.

For complete documentation of available commands check `doc/command` directory.
## Documentation
Source code documentation can be easily browsed with `go doc` command.

Documentation for Burn commands in form of Troff pages is available under `doc/command` directory.

You can easily view documentation pages with `man` command.

For example to display documentation page for gameadd command:
```
$ man doc/command/gameadd
```

Note that documentation of Burn commands is still incomplete.

## Contributing
You are welcome to contribute to project development.

If you looking for things to do, then check TODO file or contact me(dev@isangeles.pl).

When you find something to do, create new branch for your feature.
After you finish, open pull request to merge your changes with master branch.

## Contact
* Isangeles <<dev@isangeles.pl>>

## License
Copyright 2018-2020 Dariusz Sikora <<dev@isangeles.pl>>

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
MA 02110-1301, USA.
