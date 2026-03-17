# (0x86) Resend Characters

Sent by the server to re-send the character list to the client. Used after a failed character deletion or other events that require refreshing the character selection screen. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x86`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `BYTE` | `unknown` | Unknown; typically `0x00`
`0x04` | `BYTE` | `charCount` | Number of character entries (typically 5 or 7)
`0x05` | `CHAR[30]` | `charName[0]` | First character slot name, null-terminated and zero-padded
`0x23` | `CHAR[30]` | `charPassword[0]` | First character slot password field (unused, always zeros)
`...` | `...` | `...` | Remaining character entries (60 bytes each)
`0x05 + 60*charCount` | `BYTE` | `startLocationCount` | Number of starting location entries
`...` | `...` | `...` | Starting location entries (variable format)
