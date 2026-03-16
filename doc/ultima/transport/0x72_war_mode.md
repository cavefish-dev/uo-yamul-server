# (0x72) WAR_MODE

Used by the client to request a war mode change, and by the server to send the current war mode status.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x72`
`0x01` | `BYTE` | `flag` | War mode state (`0x00`=Normal, `0x01`=Fighting)
`0x02` | `BYTE` | `unknown1` | Always `0x00`
`0x03` | `BYTE` | `unknown2` | Always `0x32`
`0x04` | `BYTE` | `unknown3` | Always `0x00`

**Total size:** 5 bytes

## Flag values

Value | Meaning
--- | ---
`0x00` | Normal mode
`0x01` | Fighting / War mode enabled
