# (0xA5) Open Web Browser

Sent by the server to instruct the client to open a URL in the system web browser. Variable length depending on URL length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA5`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `CHAR[...]` | `url` | URL to open, null-terminated
