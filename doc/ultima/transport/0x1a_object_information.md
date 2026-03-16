# (0x1A) Object Information

Sent by the server to display an item on the ground. Variable length. Several fields are conditional based on flag bits in earlier fields.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x1A`
`0x01` | `USHORT` | `blockSize` | Packet size
`0x03` | `UINT` | `itemID` | Item serial; if bit 31 set, item count field follows
`0x07` | `USHORT` | `model` | Model number; if bit 15 set, an increment counter byte follows
`—` | `USHORT` | `itemCount` | (conditional) Item count; present if itemID & `0x80000000`
`—` | `BYTE` | `incrCounter` | (conditional) Increment to model; present if model & `0x8000`
`—` | `USHORT` | `xLoc` | X location (bits 0–14); if bit 15 set, direction byte follows
`—` | `USHORT` | `yLoc` | Y location; if bit 15 set, dye field follows; if bit 14 set, flag byte follows
`—` | `BYTE` | `direction` | (conditional) Direction; present if xLoc & `0x8000`
`—` | `BYTE` | `zLoc` | Z location
`—` | `USHORT` | `dye` | (conditional) Hue/dye; present if yLoc & `0x8000`
`—` | `BYTE` | `flagByte` | (conditional) Flags; present if yLoc & `0x4000`
