# (0x0b) Unknown (Edit Area?)

A client packet of unknown purpose. Potentially related to area editing; possibly obsolete.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x0b`
`0x01` | `BYTE[9]` | `unknown` | Nine unknown bytes
