# (0x13) Equip Item Request

Sent by the client to equip an item on a character. The layer byte should not be trusted.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x13`
`0x01` | `UINT` | `itemID` | Item to equip
`0x05` | `BYTE` | `layer` | Equipment layer (do not trust)
`0x06` | `UINT` | `playerID` | Target character
