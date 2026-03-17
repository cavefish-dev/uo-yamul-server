# (0x0a) God Client Edit

This packet is used by the god client to perform various tasks, most of them related to adding things.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x0a`
`0x01` | `BYTE` | `subcmd` | Subcommand byte (see values below)
`0x02` | `USHORT` | `x` | X coordinate
`0x04` | `USHORT` | `y` | Y coordinate
`0x06` | `USHORT` | `itemid` | Item identifier
`0x08` | `BYTE` | `z` | Z coordinate
`0x09` | `BYTE[2]` | `extra` | Additional data

### Subcommands

Value | Description
--- | ---
`0x04` | Add New Dynamic Item
`0x06` | Hackmove Request
`0x07` | Add New NPC
`0x0A` | Add New Static Item
