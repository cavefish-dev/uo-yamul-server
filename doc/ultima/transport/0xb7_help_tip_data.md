# (0xB7) Help/Tip Data

Contains the content of a help or tip page sent from the server.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB7`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `UINT` | `id` | Help/tip identifier
`0x07` | `UNI[]` | `message` | The message in unicode, null-terminated (`0x0000`)
`—` | `USHORT` | `terminator` | Message terminator, `0x3300`
