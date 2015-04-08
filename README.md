### yac8vm ###

Yet Another Chip-8 Virtual Machine (initial development, 2015)

### Installation ###

```
#!bash
$ sudo apt-get install libsdl2-dev   # or equivalent
$ go get bitbucket.com/qx89l4/yac8vm/cmd/yac8vm
```

### Usage ###

```
#!bash
$ ROMPATH=$GOPATH/src/bitbucket.com/qx89l4/yac8vm/roms
$ yac8vm run $ROMPATH/romfile  # execute ROM
$ yac8vm dis $ROMPATH/romfile  # disassemble ROM
```

### Keyboard ###

```
#!bash

[1][2][3][4]      [1][2][3][C]
[Q][W][E][R]  =>  [4][5][6][D]
[A][S][D][F]      [7][8][9][E]        
[Z][X][C][V]      [A][0][B][F]

P       => Pause execution
Escape  => Break execution
```

### Links ###

- http://devernay.free.fr/hacks/chip8/C8TECH10.HTM - Cowgod's Chip-8 Technical Reference 
- http://www.pong-story.com/chip8/ - David Winter's CHIP-8 emulation page

### Screenshots ###

![Hidden](https://bitbucket.org/qx89l4/yac8vm/raw/master/misc/shot-hidden.png)

![Space Invaders](https://bitbucket.org/qx89l4/yac8vm/raw/master/misc/shot-invaders.png)

![Pong](https://bitbucket.org/qx89l4/yac8vm/raw/master/misc/shot-pong.png)
