# (0x99) Arrow

Sent by the server to display or hide a directional arrow on the client's minimap pointing toward a target location. Fixed size of 6 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x99`
`0x01` | `BYTE` | `display` | `0x01` to show the arrow, `0x00` to hide it
`0x02` | `USHORT` | `x` | Target X coordinate
`0x04` | `USHORT` | `y` | Target Y coordinate
