# (0x15) Follow

Sent by the server to notify a player that one entity is following another. 9 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x15`
`0x01` | `UINT` | `unknown` | Serial of the entity being followed
`0x05` | `UINT` | `follower` | Serial of the character initiating the follow
