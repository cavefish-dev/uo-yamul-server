# (0xA9) Characters / Starting Locations

Sent by the game server after login to provide the client with the character selection screen data. Contains character slots and available starting city locations. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA9`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `BYTE` | `charCount` | Number of character slot entries (typically 5 or 7)
`0x04` | `CHAR[30]` | `charName[0]` | First character slot name, null-terminated and zero-padded
`0x22` | `CHAR[30]` | `charPassword[0]` | First character slot password field (unused, always zeros)
`...` | `...` | `...` | Remaining character slot entries (60 bytes each)
`0x04 + 60*charCount` | `BYTE` | `startLocationCount` | Number of starting location entries
`...` | `...` | `...` | Starting location entries (legacy or extended format depending on client version)
`last-3` | `UINT` | `flags` | Client feature flags affecting UI behavior
