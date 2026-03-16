# (0x20) Teleport Player

Sent by the server to teleport the player. Should only be applied to the player whose session receives the packet. 19 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x20`
`0x01` | `UINT` | `serial` | Player serial
`0x05` | `USHORT` | `bodyType` | Body type
`0x07` | `BYTE` | `unknown1` | Always `0x00`
`0x08` | `USHORT` | `skinColor` | Skin hue
`0x0A` | `BYTE` | `status` | Player status flags (see below)
`0x0B` | `USHORT` | `xLoc` | X coordinate
`0x0D` | `USHORT` | `yLoc` | Y coordinate
`0x0F` | `BYTE[2]` | `unknown2` | Always `0x0000`
`0x11` | `BYTE` | `direction` | Direction
`0x12` | `SBYTE` | `zLoc` | Z coordinate

## Status flags

Bit | Name | Description
--- | --- | ---
0x00 | Normal | No flags set
0x01 | Unknown1 | Unknown
0x02 | CanAlterPaperdoll | Player can alter paperdoll
0x04 | Poisoned | Player is poisoned
0x08 | GoldenHealth | Player has golden health bar
0x10 | Unknown2 | Unknown
0x20 | Unknown3 | Unknown
0x40 | WarMode | Player is in war mode
0x80 | Hidden | Player is hidden
