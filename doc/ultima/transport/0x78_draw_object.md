# (0x78) DRAW_OBJECT

Server packet used to display a character on screen. Variable length. Note: this page may be incomplete and/or inaccurate.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x78`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `UINT` | `serial` | Item ID / Player ID
`0x07` | `USHORT` | `body` | Model number
`--` | `USHORT` | `corpseModel` | Corpse model number (present if serial & `0x80000000`)
`0x09` | `USHORT` | `xLoc` | X location (15-bit value; bit 15 signals direction presence)
`0x0B` | `USHORT` | `yLoc` | Y location
`--` | `BYTE` | `direction` | Direction (present if xLoc & `0x8000`)
`0x0D` | `SBYTE` | `zLoc` | Z location
`0x0E` | `BYTE` | `direction` | Direction
`0x0F` | `USHORT` | `hue` | Dye / skin colour
`0x11` | `BYTE` | `status` | Status byte
`0x12` | `BYTE` | `notoriety` | Notoriety

## Equipped items loop

Repeated for each equipped item until a 0x00000000 terminator is read.

Offset | Type | Name | Description
--- | --- | --- | ---
`--` | `UINT` | `itemserial` | Item serial number
`--` | `USHORT` | `artwork` | Graphic ID (15-bit; bit 15 indicates hue follows)
`--` | `BYTE` | `layer` | Clothing layer
`--` | `USHORT` | `hue` | Colour (present if artwork & `0x8000`)
