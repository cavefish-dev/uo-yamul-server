# (0x75) RENAME_MOB

Sent by the client to alter the name of a mobile.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x75`
`0x01` | `UINT` | `id` | Target character's serial number
`0x05` | `CHAR[30]` | `name` | Target character's new name

**Total size:** 35 bytes
