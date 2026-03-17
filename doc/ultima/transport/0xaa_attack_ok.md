# (0xAA) Attack Request Reply

This tells the client which target they are fighting.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xAA`
`0x01` | `UINT` | `serial` | The serial of the current attack target; serial is set to `0x00000000` when attack is refused
