# (0x71) BULLETIN_BOARD

Manages bulletin board communications. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x71`
`0x01` | `USHORT` | `blockSize` | Total size of the packet
`0x03` | `BYTE` | `action` | Subcommand / action identifier (see values below)

## Subcommands

Value | Meaning
--- | ---
`0x00` | Display Bulletin Board
`0x01` | Message Summary
`0x02` | Message
`0x03` | Request Message
`0x04` | Request Message Summary
`0x05` | Post a Message
`0x06` | Remove Posted Message

No documentation available for individual subcommand field layouts.
