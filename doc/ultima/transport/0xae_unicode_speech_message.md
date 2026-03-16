# (0xAE) Unicode Speech Message

Sent to tell the client that someone is talking.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xAE`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `UINT` | `serial` | The serial of the character that is talking; `0xFFFFFFFF` is used for system messages
`0x07` | `UINT` | `body` | The body number of the character that is talking; `0xFFFF` is used for system
`0x0B` | `BYTE` | `mode` | Speech mode
`0x0C` | `USHORT` | `hue` | The message's hue value
`0x0E` | `USHORT` | `font` | Font number
`0x10` | `CHAR[4]` | `language` | The client's language preference
`0x14` | `CHAR[30]` | `name` | The name of the character that is talking
`0x32` | `UNI[]` | `text` | The message in unicode, null-terminated (`0x0000`)
