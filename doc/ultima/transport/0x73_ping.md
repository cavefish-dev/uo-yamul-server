# (0x73) PING

Sent by the client to ping the server, and by the server to ping the client.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x73`
`0x01` | `BYTE` | `seq` | Ping sequence value (usually `0x00`)

**Total size:** 2 bytes
