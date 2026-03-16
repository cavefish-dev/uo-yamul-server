# (0x04) God Mode

This is a client packet used in the God client to toggle god mode. Little information is available about this special client.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x04`
`0x01` | `BYTE` | `mode` | Toggle state (0=off, 1=on)
