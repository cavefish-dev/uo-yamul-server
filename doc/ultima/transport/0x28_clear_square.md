# (0x28) Clear Square

Server message. Purpose not fully documented. 5 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x28`
`0x01` | `USHORT` | `xLoc` | X coordinate of the square to clear
`0x03` | `USHORT` | `yLoc` | Y coordinate of the square to clear
