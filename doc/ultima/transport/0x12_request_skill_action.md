# (0x12) Request Skill / Action / Magic Usage

Sent by the client to request skill use, spell casting, or other actions. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x12`
`0x01` | `USHORT` | `blockSize` | Packet size
`0x03` | `BYTE` | `type` | Command type (see values below)
`0x04` | `BYTE[]` | `data` | Null-terminated payload; content depends on type

## Type values

Type | Name | Description
--- | --- | ---
0x00 | God Mode Teleport | Teleport in god mode
0x24 ($) | Skill Usage | Followed by null-terminated skill identifier (e.g. "1 0" for anatomy, "21 0" for hiding)
0x56 (V) | Macro'd Spell | Followed by spell number (2–64)
0x58 (X) | Open Door | Null termination byte only
0x6B | God Mode Command | God client command
0xC7 | Action | Null-terminated string: "bow" or "salute"
