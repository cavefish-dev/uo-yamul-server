# (0x0d) Unknown

A client packet of unknown purpose. Possibly related to NPC data; potentially obsolete.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x0d`
`0x01` | `BYTE[2]` | `unknown` | Two unknown bytes
