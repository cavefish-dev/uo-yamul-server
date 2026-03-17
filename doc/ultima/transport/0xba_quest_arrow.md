# (0xBA) Quest Arrow

Server message that shows or hides the quest arrow pointing to a map location.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xBA`
`0x01` | `BYTE` | `active` | 1 = on, 0 = off
`0x02` | `USHORT` | `xloc` | X coordinate of the target location
`0x04` | `USHORT` | `yloc` | Y coordinate of the target location
