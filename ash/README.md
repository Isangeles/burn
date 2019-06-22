## Introduction
Ash is scripting language that executes Burn commands in conditional loop.

## Examples
Infinite loop script:
```
# Every 5 seconds, sends text from first argument on chat
# channel of game character with serial ID from argument 2
true {
    charman -o set -a chat @1 -t @2;
    wait(5);
}
```
Script with inner block:
```
# Script that sets position of character with serial ID from arg 1
# to 0x0 if range between him and character with serial ID from arg 2
# is less than 50.
true {
    rawdis(@1, @2) < 50 {
      charman -o set -a position 0 0 -t @1;
	    wait(5);
    };
}
```
