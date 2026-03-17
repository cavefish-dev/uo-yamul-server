# (0x95) Hue Picker

Bidirectional packet used for the hue (color) selection interface. The server sends it to open the hue picker gump; the client sends back the same structure with the chosen hue. Fixed size of 9 bytes in both directions.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x95`
`0x01` | `UINT` | `serial` | Serial of the item being recolored
`0x05` | `USHORT` | `hue` | Selected hue value (`0x0000` when server opens picker)
`0x07` | `USHORT` | `graphic` | Graphic ID of the item being displayed in the picker

## Example packet (hex)

```
-- Server opens hue picker for item 0x40000001, graphic 0x1F03
95 40 00 00 01 00 00 1F
03

-- Client response: selected hue 0x0384
95 40 00 00 01 03 84 1F
03
```
