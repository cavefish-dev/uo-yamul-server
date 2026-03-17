# STATICS0.MUL

Static map objects (buildings, trees, obstacles, decorations) for facet 0. Accessed via `STAIDX0.MUL`.

Reference: [Heptazane — Section 3.18](https://uo.stratics.com/heptazane/fileformats.shtml#3.18) · [SphereServer — uo_files/CUOStaticItem.h](https://github.com/Sphereserver/Source-X/blob/master/src/game/uo_files/CUOStaticItem.h)

## File structure

Each block is a sequence of 7-byte static records at the position and length given by `STAIDX0.MUL`. Record count = `Length / 7`.

### Static record (7 bytes)

Source: `CUOStaticItemRec` (`m_wTileID`, `m_x`, `m_y`, `m_z`, `m_wHue`)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `UWORD` | `TileID` | Tile/item ID; add 16384 for `RADARCOL.MUL` lookup
`0x02` | `1` | `UBYTE` | `X` | Cell X offset within the 8×8 block (0–7)
`0x03` | `1` | `UBYTE` | `Y` | Cell Y offset within the 8×8 block (0–7)
`0x04` | `1` | `BYTE` | `Alt` | Altitude (-128 to 127)
`0x05` | `2` | `UWORD` | `Hue` | Hue/color modifier applied to the item
