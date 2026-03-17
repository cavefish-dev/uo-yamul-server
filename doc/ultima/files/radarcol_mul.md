# RADARCOL.MUL

Radar-view color lookup table. Fixed size: 131,072 bytes (65,536 × 2).

Reference: [Heptazane — Section 3.12](https://uo.stratics.com/heptazane/fileformats.shtml#3.12)

## File structure

65,536 sequential 16-bit color entries. Index directly by the tile ID from `MAP0.MUL` (`Graphic` field) or the object ID from `STATICS0.MUL` (`Color` field) to get the radar display color.

### Color entry

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`tileID * 2` | `2` | `UWORD` | `Color` | 16-bit color in UO format (see [`types.md`](types.md))
