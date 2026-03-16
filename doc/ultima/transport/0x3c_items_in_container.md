# (0x3C) ITEMS_IN_CONTAINER

This is sent by the server to display the contents of a container. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x3C`
`0x01` | `USHORT` | `blockSize` | Total packet size in bytes
`0x03` | `USHORT` | `itemCount` | Number of item entries that follow

Each item entry repeats for every item in the container:

Offset (per item) | Type | Name | Description
--- | --- | --- | ---
`+0x00` | `UINT` | `serial` | Item serial identifier
`+0x04` | `USHORT` | `artwork` | Item graphic ID
`+0x06` | `BYTE` | `stackId` | Stack identifier
`+0x07` | `USHORT` | `amount` | Quantity in the stack
`+0x09` | `USHORT` | `xLoc` | Horizontal position within the container
`+0x0B` | `USHORT` | `yLoc` | Vertical position within the container
`+0x0D` | `UINT` | `containerID` | Serial of the parent container
`+0x11` | `USHORT` | `color` | Item hue/tint
