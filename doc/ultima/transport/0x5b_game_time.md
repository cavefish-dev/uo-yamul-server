# (0x5B) GAME_TIME

This is sent by the server to set the current in-game time. Fixed size of 4 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x5B`
`0x01` | `BYTE` | `hour` | Hour value
`0x02` | `BYTE` | `minute` | Minute value
`0x03` | `BYTE` | `second` | Second value
