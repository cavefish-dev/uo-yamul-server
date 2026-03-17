# (0x25) Add Item to Container

Sent by the server to add a single item to a container. Should only be used when a player moves an item into a container, not for populating initial container contents. 20 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x25`
`0x01` | `UINT` | `serial` | Serial of the item to add
`0x05` | `USHORT` | `artwork` | Item graphic
`0x07` | `BYTE` | `unknown1` | Always `0x00`
`0x08` | `USHORT` | `amount` | Number of items in the stack
`0x0A` | `USHORT` | `dx` | Pixel offset from the left edge of the container gump
`0x0C` | `USHORT` | `dy` | Pixel offset from the top edge of the container gump
`0x0E` | `UINT` | `containerID` | Serial of the container receiving the item
`0x12` | `USHORT` | `hue` | Item color/hue
