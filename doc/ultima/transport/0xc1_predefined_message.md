# (0xC1) Predefined Message

Server message. Transmits a predefined system message. The packet length is always 0x32 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xC1`
`0x01` | `USHORT` | `len` | Packet length, always `0x32`
`0x03` | `UINT` | `id` | Entity identifier
`0x07` | `USHORT` | `body` | Body appearance value
`0x09` | `BYTE` | `type` | Display type: 6 = lower left, 7 = on player
`0x0A` | `USHORT` | `hue` | Color value
`0x0C` | `USHORT` | `font` | Font identifier
`0x0E` | `USHORT` | `msgtype` | Message type, `0x0007` observed
`0x10` | `USHORT` | `messagenumber` | Message number (`0xA120` base)
`0x12` | `CHAR[32]` | `speaker` | Speaker's name

## Example message numbers

Value | Message
--- | ---
`0xA12D` | "You cannot use skills."
`0xA2E2` | "But that's not dead!"
`0xA5F3` | "Help request aborted."
`0xA5F0` | Help prompt message
