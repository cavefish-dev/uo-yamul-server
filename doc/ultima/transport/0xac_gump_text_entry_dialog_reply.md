# (0xAC) Gump Text Entry Dialog Reply

This is the client's response to the String Query packet (0xAB).

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xAC`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `UINT` | `serial` | Query serial
`0x07` | `BYTE` | `unknown1` | Type?
`0x08` | `BYTE` | `unknown2` | Index?
`0x09` | `USHORT` | `length` | The length of the response, including the null terminator
`0x0B` | `CHAR[]` | `text` | The client's text response
