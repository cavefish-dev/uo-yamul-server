# (0x76) NEW_SUBSERVER

New subserver packet.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x76`
`0x01` | `BYTE[2]` | `xloc` | X location coordinates
`0x03` | `BYTE[2]` | `yloc` | Y location coordinates
`0x05` | `BYTE[2]` | `zloc` | Z location coordinates
`0x07` | `BYTE[9]` | `unknown` | Reserved / unknown data

**Total size:** 16 bytes
