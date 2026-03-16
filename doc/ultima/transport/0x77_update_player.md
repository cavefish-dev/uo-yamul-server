# (0x77) UPDATE_PLAYER

Sent to update a character on screen when it is walking or changing in some way.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x77`
`0x01` | `UINT` | `playerId` | Player identifier
`0x05` | `USHORT` | `model` | Character model / appearance
`0x07` | `USHORT` | `xLoc` | X coordinate
`0x09` | `USHORT` | `yLoc` | Y coordinate
`0x0B` | `SBYTE` | `zLoc` | Z coordinate
`0x0C` | `BYTE` | `direction` | Character facing direction
`0x0D` | `USHORT` | `hue` | Hue / skin colour value
`0x0F` | `BYTE` | `flag` | Character status indicator
`0x10` | `BYTE` | `notoriety` | Notoriety level / alignment

**Total size:** 17 bytes
