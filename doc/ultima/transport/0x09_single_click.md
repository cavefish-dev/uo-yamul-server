# (0x09) Single Click

This is sent by the client when the player single clicks an object.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x09`
`0x01` | `UINT` | `serial` | ID of single clicked object
