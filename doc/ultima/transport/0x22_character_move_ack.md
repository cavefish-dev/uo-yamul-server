# (0x22) Character Move ACK / Resync Request

Sent by the server to accept a movement request. The client also sends this packet to request a resync; the proper server response to a resync request is a Teleport (0x20) packet. 3 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x22`
`0x01` | `BYTE` | `sequence` | Sequence number matching the accepted move request
`0x02` | `BYTE` | `status` | Player status flags
