# (0xA6) Tip/Notice Window

Sent by the server to display a scroll-style pop-up window containing text (used for tips of the day, game notices, or other informational messages). Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA6`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `BYTE` | `scrollType` | Scroll window type: `0x00`=tip of the day, `0x01`=notice
`0x04` | `UINT` | `context` | Context identifier (tip index or notice ID)
`0x08` | `USHORT` | `textLen` | Length of the text that follows
`0x0A` | `CHAR[textLen]` | `text` | Message text (not null-terminated)
