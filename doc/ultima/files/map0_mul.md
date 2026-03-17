# MAP0.MUL

Base terrain data for facet 0 (Britannia). Total size: 37,748,736 bytes.

Reference: [Heptazane — Section 3.8](https://uo.stratics.com/heptazane/fileformats.shtml#3.8) · [SphereServer — uo_files/CUOMapList.h](https://github.com/Sphereserver/Source-X/blob/master/src/game/uo_files/CUOMapList.h)

## Facets

UO ships multiple map files for different game regions. Each has different dimensions:

| File | Facet | Name | Width × Height (tiles) |
| --- | --- | --- | --- |
| `MAP0.MUL` | 0 | Britannia (Felucca) | 7168 × 4096 |
| `MAP1.MUL` | 1 | Trammel | 7168 × 4096 |
| `MAP2.MUL` | 2 | Ilshenar | 2304 × 1600 |
| `MAP3.MUL` | 3 | Malas | 2560 × 2048 |
| `MAP4.MUL` | 4 | Tokuno Islands | 1448 × 1448 |
| `MAP5.MUL` | 5 | Termur (HS) | 1280 × 4096 |

## File structure

The file is sequential blocks arranged in a grid of map blocks. Each block covers an 8×8 tile area.

**Block location formula:**
```
XBlock   = XPos / 8
YBlock   = YPos / 8
BlockNum = (XBlock * (MapHeight / 8)) + YBlock
FileOffset = BlockNum * 196
```

### Map block (196 bytes)

Source: `CUOMapBlock` (`m_wID1`, `m_wID2`, `m_Cells`)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `UWORD` | `ID1` | Block header ID 1 (`m_wID1`)
`0x02` | `2` | `UWORD` | `ID2` | Block header ID 2 (`m_wID2`)
`0x04` | `3*64` | `Cell[64]` | `Cells` | 8×8 grid of terrain cells (row-major)

### Cell (3 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `UWORD` | `Graphic` | Terrain tile ID; look up color in `RADARCOL.MUL`
`0x02` | `1` | `BYTE` | `Altitude` | Height in game units (-128 to 127); 0 = sea level
