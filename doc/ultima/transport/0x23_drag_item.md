# (0x23) Drag Item

Sent by the server to display an animation of an item being dragged from one location to another. This packet does not actually move the item; it only triggers the visual animation. 26 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x23`
`0x01` | `USHORT` | `modelID` | Item graphic (artwork)
`0x03` | `BYTE[3]` | `unknown1` | Always `0x000000`
`0x06` | `USHORT` | `amount` | Item quantity
`0x08` | `UINT` | `srcContainer` | Source container serial; `0xFFFFFFFF` = ground
`0x0C` | `USHORT` | `srcX` | Source X coordinate
`0x0E` | `USHORT` | `srcY` | Source Y coordinate
`0x10` | `SBYTE` | `srcZ` | Source Z coordinate
`0x11` | `UINT` | `dstContainer` | Destination container serial; `0xFFFFFFFF` = ground
`0x15` | `USHORT` | `dstX` | Destination X coordinate
`0x17` | `USHORT` | `dstY` | Destination Y coordinate
`0x19` | `SBYTE` | `dstZ` | Destination Z coordinate
