# (0x1C) Send Text

Sent by the server to display character names, system messages, object messages, etc. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x1C`
`0x01` | `USHORT` | `blockSize` | Packet size
`0x03` | `UINT` | `itemID` | Serial of the speaker; `0xFFFFFFFF` = system message
`0x07` | `USHORT` | `model` | Model of the speaker; `0xFFFF` = system message
`0x09` | `BYTE` | `type` | Message type (see values below)
`0x0A` | `USHORT` | `hue` | Text color
`0x0C` | `USHORT` | `font` | Font identifier
`0x0E` | `BYTE[30]` | `name` | Name of the speaker or object
`0x2C` | `BYTE[]` | `message` | Null-terminated message text

## Type values

Type | Description
--- | ---
`0x00` | Normal
`0x01` | Broadcast
`0x02` | Emote
`0x06` | System
`0x07` | Message
`0x08` | Whisper
`0x09` | Yell

## Example packet (hex)

```
-- "John Doe" says "Hello World!"
1C 00 39 00 00 00 01 00 02 03 00 04 00 05 4A 6F
68 6E 20 44 6F 65 00 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 48 65 6C 6C
6F 20 57 6F 72 6C 64 21 00
```
