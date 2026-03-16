# (0x02) Move Request

This packet is sent by the client when the player tries to walk, run, or change his direction. The sequence counter increments from 0 to 255, then resets to 1 (not back to 0). The fastwalk prevention key is a server-supplied value sent in the Move Ack response.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x02`
`0x01` | `BYTE` | `direction` | Movement direction
`0x02` | `BYTE` | `sequence` | Sequence number (0-255, resets to 1 not 0)
`0x03` | `UINT` | `key` | Fastwalk prevention key from server
