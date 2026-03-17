# (0x2C) Death / Resurrection

Sent by the server to inform the player of character death or to process a resurrection choice. The resurrection menu has been removed from modern UO clients. 2 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x2C`
`0x01` | `BYTE` | `action` | Action code (see values below)

## Action values

Value | Description
--- | ---
`0x00` | Notify client of character death
`0x01` | Client selected resurrection with penalties
`0x02` | Client opted to continue as ghost
