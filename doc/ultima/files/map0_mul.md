# MAP0.MUL

Base terrain data for facet 0 (Britannia). Total size: 37,748,736 bytes.

Reference: [Heptazane — Section 3.8](https://uo.stratics.com/heptazane/fileformats.shtml#3.8)

## File structure

The file is 393,216 sequential blocks arranged in a 768×512 grid of map blocks. Each block covers an 8×8 tile area.

**Block location formula:**
```
XBlock   = XPos / 8
YBlock   = YPos / 8
BlockNum = (XBlock * 512) + YBlock
FileOffset = BlockNum * 196
```

### Map block (196 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Header` | Unknown header value
`0x04` | `3*64` | `Cell[64]` | `Cells` | 8×8 grid of terrain cells (row-major)

### Cell (3 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `UWORD` | `Graphic` | Terrain tile ID; look up color in `RADARCOL.MUL`
`0x02` | `1` | `BYTE` | `Altitude` | Height in game units (-128 to 127); 0 = sea level
