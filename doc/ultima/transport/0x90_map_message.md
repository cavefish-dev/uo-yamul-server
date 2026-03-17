# (0x90) Map Message

Sent by the server to open or update a map gump on the client. Specifies the map item serial, the gump graphic, the map bounds in world coordinates, and the displayed map dimensions. Fixed size of 19 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x90`
`0x01` | `UINT` | `mapUID` | Serial of the map item
`0x05` | `USHORT` | `gumpID` | Gump graphic ID used to display the map
`0x07` | `USHORT` | `left` | Left bound of the mapped area in world coordinates
`0x09` | `USHORT` | `top` | Top bound of the mapped area in world coordinates
`0x0B` | `USHORT` | `right` | Right bound of the mapped area in world coordinates
`0x0D` | `USHORT` | `bottom` | Bottom bound of the mapped area in world coordinates
`0x0F` | `USHORT` | `mapWidth` | Width of the map display in pixels
`0x11` | `USHORT` | `mapHeight` | Height of the map display in pixels

## Example packet (hex)

```
-- Map item 0x40000001, gump 0x0139, area (0,0)-(400,400), display 400x400
90 40 00 00 01 01 39 00
00 00 00 01 90 01 90 01
90 01 90
```
