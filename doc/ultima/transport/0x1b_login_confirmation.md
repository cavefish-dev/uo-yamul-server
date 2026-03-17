# (0x1B) Login Confirmation

Sent by the server to confirm a login on a shard, providing the character location, body type, and other data. 37 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x1B`
`0x01` | `UINT` | `playerID` | Player serial
`0x05` | `UINT` | `unknown1` | Unknown
`0x09` | `USHORT` | `bodyType` | Character body type
`0x0B` | `USHORT` | `xLoc` | X coordinate
`0x0D` | `USHORT` | `yLoc` | Y coordinate
`0x0F` | `BYTE` | `unknown2` | Always `0x00`
`0x10` | `SBYTE` | `zLoc` | Z coordinate
`0x11` | `BYTE` | `direction` | Character direction
`0x12` | `BYTE[2]` | `unknown3` | Unknown
`0x14` | `BYTE[4]` | `unknown4` | Usually `0xFF`
`0x18` | `BYTE[4]` | `unknown5` | Unknown
`0x1C` | `BYTE` | `flagByte` | Possibly MapWidth - 8
`0x1D` | `BYTE` | `highlightColor` | Possibly MapHeight
`0x1E` | `BYTE` | `unknown6` | Unknown
`0x1F` | `USHORT` | `unknown7` | Unknown
`0x21` | `UINT` | `unknown8` | Unknown

## Example packet (hex)

```
-- CharID=0x1FEF, body=0x0190, x=0x041D, y=0x0598, z=0xFFAB, direction=East
1B 00 00 1F EF 00 00 00 00 01 90 04 1D 05 98 FF
AB 07 00 FF FF FF FF 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00
```
