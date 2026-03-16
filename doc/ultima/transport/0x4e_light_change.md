# (0x4E) LIGHT_CHANGE

This is sent by the server to change the personal light level of an object. Fixed size of 6 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x4E`
`0x01` | `UINT` | `creature_id` | Serial of the target object
`0x05` | `BYTE` | `level` | Personal light level for the target
