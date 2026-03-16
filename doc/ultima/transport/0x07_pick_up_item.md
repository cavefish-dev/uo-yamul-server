# (0x07) Pick Up Item(s)

This is sent by the client when the player picks up an item.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x07`
`0x01` | `UINT` | `serial` | Item identifier
`0x05` | `USHORT` | `amount` | Number of items in stack
