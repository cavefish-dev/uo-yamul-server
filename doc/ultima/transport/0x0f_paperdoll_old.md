# (0x0f) Paperdoll (Old)

Legacy paperdoll packet. The structure and purpose of the payload bytes are undocumented.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x0f`
`0x01` | `BYTE[60]` | `unknown` | 60 bytes of undocumented paperdoll data
