# (0x06) Double Click

This is sent by the client when it double clicks an object. When the serial includes the flag 0x80000000, the target represents the player character itself.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x06`
`0x01` | `UINT` | `serial` | ID of double clicked object (`0x80000000` flag = self)
