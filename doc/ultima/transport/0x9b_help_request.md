# (0x9B) Help Request

Sent by the client when the player clicks the Help button. Triggers the server-side help system. The packet body contains no meaningful fields beyond the command byte; the remainder is zero-padded. Fixed size of 258 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x9B`
`0x01` | `BYTE[257]` | `padding` | Zero-padded; no meaningful fields
