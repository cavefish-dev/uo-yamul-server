# (0x08) Drop Item(s)

This is sent by the client when the player drops an item.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x08`
`0x01` | `UINT` | `serial` | Item ID
`0x05` | `USHORT` | `locX` | X coordinate
`0x07` | `USHORT` | `locY` | Y coordinate
`0x09` | `BYTE` | `locZ` | Z coordinate (signed)
`0x0A` | `UINT` | `containerId` | Target container ID (`0xFFFFFFFF` for world placement)
