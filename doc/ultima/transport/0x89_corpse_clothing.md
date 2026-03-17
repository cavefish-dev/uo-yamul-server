# (0x89) Corpse Clothing

Sent by the server to inform the client which items are on a corpse, organized by body layer. Each entry specifies the layer and the item's serial. The list is terminated by a zero byte. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x89`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `UINT` | `corpseUID` | Serial of the corpse container
`0x07` | `BYTE` | `layer` | Body layer of first equipped item (non-zero)
`0x08` | `UINT` | `itemUID` | Serial of the first item
`...` | `...` | `...` | Additional `layer(BYTE) + itemUID(UINT)` pairs
`last` | `BYTE` | `terminator` | Always `0x00`, marks end of item list
