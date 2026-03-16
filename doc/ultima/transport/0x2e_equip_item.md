# (0x2E) Equip Item

Sent by the server to equip a single item on a character. 15 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x2E`
`0x01` | `UINT` | `serial` | Serial of the item to equip (always starts `0x40` in observed data)
`0x05` | `USHORT` | `model` | Item model number
`0x07` | `BYTE` | `unknown1` | Always `0x00`
`0x08` | `BYTE` | `layer` | Equipment layer
`0x09` | `UINT` | `playerID` | Serial of the character on which the item will be equipped
`0x0D` | `USHORT` | `hue` | Item color/hue
