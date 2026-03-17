# (0x6C) TARGETING_CURSOR

The server sends this packet to bring up a targeting cursor. The client sends it back after targeting something or pressing the Escape key.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x6C`
`0x01` | `BYTE` | `type` | Targeting type (see values below)
`0x02` | `UINT` | `serial` | Targeting cursor's serial number
`0x06` | `BYTE` | `flag` | Flag value (see values below)
`0x07` | `UINT` | `target` | Target object's serial (client response only)
`0x0B` | `USHORT` | `xLoc` | X coordinate (client response only)
`0x0D` | `USHORT` | `yLoc` | Y coordinate (client response only)
`0x0F` | `SHORT` | `zLoc` | Z coordinate (client response only)
`0x11` | `USHORT` | `model` | Artwork number (client response only)

**Total size:** 19 bytes

## Type values

Value | Meaning
--- | ---
`0x00` | Select Object
`0x01` | Select X, Y, Z

## Flag values

Value | Meaning
--- | ---
`0x00` | Normal
`0x01` | Criminal Action
`0x02` | Unknown
`0x03` | Cancel Target (server-side)

## Notes

- The model number field should not be relied upon for accuracy.
