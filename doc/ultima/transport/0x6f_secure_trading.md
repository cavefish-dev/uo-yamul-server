# (0x6F) SECURE_TRADING

Handles secure trading functionality between characters. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x6F`
`0x01` | `USHORT` | `blockSize` | Size of the packet
`0x03` | `BYTE` | `action` | Action identifier
`0x04` | `UINT` | `id1` | First identifier
`0x08` | `UINT` | `id2` | Second identifier
`0x0C` | `UINT` | `id3` | Third identifier
`0x10` | `BYTE` | `nameFollowing` | Flag indicating a character name follows (0=no, 1=yes)
`--` | `BYTE[]` | `charName` | Character name (present only if nameFollowing = 1)
