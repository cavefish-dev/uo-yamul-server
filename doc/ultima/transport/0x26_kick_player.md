# (0x26) Kick Player

Sent by the server to kick a player from the game. Old client format. 5 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x26`
`0x01` | `UINT` | `id` | Serial of the GM who issued the kick
