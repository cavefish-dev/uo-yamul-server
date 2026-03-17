# (0xA8) Server List

Sent by the login server after successful authentication to present the list of available game servers. The client uses this list to display the server selection screen. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA8`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `BYTE` | `flags` | Server list flags; always `0xFF`
`0x04` | `USHORT` | `serverCount` | Number of server entries that follow
`0x06` | `USHORT` | `index` | Index of first server entry (zero-based)
`0x08` | `CHAR[32]` | `name` | Server name for first entry, null-terminated and zero-padded
`0x28` | `BYTE` | `percentFull` | Server load as a percentage (0–100)
`0x29` | `BYTE` | `timezone` | Server timezone offset
`0x2A` | `UINT` | `ip` | Server IP address as a packed 32-bit integer
`...` | `...` | `...` | Additional server entries repeating from `index` (40 bytes each)
