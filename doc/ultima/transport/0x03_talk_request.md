# (0x03) Talk Request

This is sent by older clients instead of the Speech Unicode packet.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x03`
`0x01` | `USHORT` | `blockSize` | Length of message plus 8 bytes
`0x03` | `BYTE` | `speechType` | Speech mode/type (see values below)
`0x04` | `USHORT` | `hue` | Color value
`0x06` | `USHORT` | `font` | Font identifier
`0x08` | `CHAR[]` | `msg` | Null-terminated message string

### Speech Type Values

Value | Description
--- | ---
`0x00` | Regular
`0x01` | Broadcast
`0x02` | Emote
`0x06` | System
`0x07` | Message
`0x08` | Whisper
`0x09` | Yell
