# (0x0c) Tile Data

Sent by the god client to edit tiledata.mul. Variable length packet. If tileNum includes the flag 0x8000 then it is a map tile.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x0c`
`0x01` | `USHORT` | `size` | Packet size in bytes
`0x03` | `USHORT` | `tileNum` | Tile identifier to modify (`0x8000` flag = map tile)
`0x05` | `UINT` | `flags` | Tiledata flags for the item
`0x09` | `BYTE` | `weight` | Item weight
`0x0A` | `BYTE` | `quality` | Item quality
`0x0B` | `USHORT` | `unknown1` | Unknown field
`0x0D` | `BYTE` | `unknown2` | Unknown field
`0x0E` | `BYTE` | `quantity` | Item quantity
`0x0F` | `USHORT` | `anim_frame` | Animation frame number
`0x11` | `BYTE` | `unknown3` | Unknown field
`0x12` | `BYTE` | `hue` | Color/hue designation
`0x13` | `BYTE` | `unknown4` | Unknown field
`0x14` | `BYTE` | `value` | Item value
`0x15` | `BYTE` | `height` | Item height
`0x16` | `CHAR[20]` | `name` | Item name string
