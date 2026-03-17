# (0xAF) Death Animation

This is used to display a death animation.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xAF`
`0x01` | `UINT` | `serial` | The serial of the dying character
`0x05` | `UINT` | `corpse` | The serial of the corpse item
`0x09` | `UINT` | `unknown` | Unknown
