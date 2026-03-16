# (0x33) PAUSE_RESUME

This is sent by the client to pause or resume client operations. This could be used to implement a "sleep" functionality.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x33`
`0x01` | `BYTE` | `action` | 0 = pause, 1 = resume
