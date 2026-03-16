# (0xB0) Send Gump Menu Dialog

This is sent by the server and used to display a gump.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB0`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `UINT` | `player` | Serial of the player receiving the gump
`0x07` | `UINT` | `serial` | Gump instance serial
`0x0B` | `UINT` | `top` | The x position in pixels at which the gump will appear
`0x0F` | `UINT` | `left` | The y position in pixels at which the gump will appear
`0x13` | `USHORT` | `length_data` | Size of the layout data section
`0x15` | `CHAR[]` | `data` | Gump layout commands, null-terminated (`0x00`)
`—` | `USHORT` | `lines` | Number of text lines in the gump

Repeated for each text line (loops `lines` times):

Offset | Type | Name | Description
--- | --- | --- | ---
`—` | `USHORT` | `length` | Length of this text entry in characters
`—` | `UNI[]` | `text` | Unicode text string
