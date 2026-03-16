# (0xBD) Client Version Message

Sent by the client to report its version string to the server.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xBD`
`0x01` | `USHORT` | `length` | Packet length
`0x03` | `CHAR[]` | `client` | Client version string, null-terminated (e.g. "1.26.4")
