# (0x34) GET_PLAYER_STATUS

This is sent by the client to query various information types, including verdata, tiledata, character status, and character skills. Fixed size of 10 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x34`
`0x01` | `UINT` | `pattern` | Always `0xEDEDEDED`
`0x05` | `BYTE` | `type` | Query type (see values below)
`0x06` | `UINT` | `serial` | Character serial, or verdata type identifier when type is `0x00`

## Type values

Value | Description
--- | ---
`0x00` | Verdata query
`0x02` | Tiledata query
`0x03` | Unknown
`0x04` | Character status (server responds with packet `0x11`)
`0x05` | Character skills (server responds with packet `0x3A`)
