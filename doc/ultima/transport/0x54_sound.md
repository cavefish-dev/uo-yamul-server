# (0x54) SOUND

This is sent by the server to play a sound effect at a specified 3D location. Fixed size of 12 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x54`
`0x01` | `BOOL` | `singular` | If false, the sound repeats indefinitely
`0x02` | `USHORT` | `effect` | Sound effect number
`0x04` | `USHORT` | `volume` | Volume level (changing this appears to have no effect in the client)
`0x06` | `USHORT` | `x` | X coordinate of the sound origin
`0x08` | `USHORT` | `y` | Y coordinate of the sound origin
`0x0A` | `USHORT` | `z` | Z coordinate of the sound origin
