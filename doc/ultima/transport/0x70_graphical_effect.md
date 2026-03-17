# (0x70) GRAPHICAL_EFFECT

Displays a visual effect, supporting various animation behaviours including directional effects, lightning strikes, and stationary animations.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x70`
`0x01` | `BYTE` | `direction` | Effect behaviour type (see values below)
`0x02` | `UINT` | `characterId` | Source character identifier
`0x06` | `UINT` | `targetId` | Destination character identifier
`0x0A` | `USHORT` | `model` | Starting animation frame
`0x0C` | `USHORT` | `xLoc` | Source X coordinate
`0x0E` | `USHORT` | `yLoc` | Source Y coordinate
`0x10` | `SBYTE` | `zLoc` | Source Z coordinate
`0x11` | `USHORT` | `xLoc_target` | Destination X coordinate
`0x13` | `USHORT` | `yLoc_target` | Destination Y coordinate
`0x15` | `SBYTE` | `zLoc_target` | Destination Z coordinate
`0x16` | `BYTE` | `speed` | Animation playback rate
`0x17` | `BYTE` | `duration` | Effect length (0=extended, 1=minimal)
`0x18` | `USHORT` | `unknown1` | Reserved field (always 0)
`0x1A` | `BYTE` | `adjust` | Rotation behaviour (1=disabled)
`0x1B` | `BYTE` | `explode` | Impact detonation flag

**Total size:** 28 bytes

## Direction values

Value | Meaning
--- | ---
`0x00` | Vector animation from source to target
`0x01` | Electrical discharge at origin
`0x02` | Stationary effect at current location
`0x03` | Effect follows source character
