# (0x21) Character Move Reject

Sent by the server to reject a movement request. Sending this packet resets the client's movement sequence to zero; the client must jump back to the given position. 8 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x21`
`0x01` | `BYTE` | `sequence` | Sequence number of the rejected move
`0x02` | `USHORT` | `xLoc` | Correct X coordinate
`0x04` | `USHORT` | `yLoc` | Correct Y coordinate
`0x06` | `BYTE` | `direction` | Correct direction
`0x07` | `SBYTE` | `zLoc` | Correct Z coordinate
