# (0xBB) Ultima Messenger

Both client and server message. Purpose not fully documented.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xBB`
`0x01` | `UINT` | `id1` | Identifier 1
`0x05` | `UINT` | `id2` | Identifier 2
