# MULTI.MUL

Multi-object data: grouped art pieces composing houses, castles, ships, etc. Accessed via `MULTI.IDX`.

Reference: [Heptazane — Section 3.10](https://uo.stratics.com/heptazane/fileformats.shtml#3.10) · [SphereServer — uo_files/CUOMulti.h](https://github.com/Sphereserver/Source-X/blob/master/src/game/uo_files/CUOMulti.h)

## File structure

Each multi entry is a sequence of component records at the offset from `MULTI.IDX`. Record size depends on client version (see variants below).

> **High Seas note:** Pre-HS clients use 12-byte records; HS+ clients use 14-byte records with an extra ship-access field.

### Component record (12 bytes)

Source: `CUOMultiItemRec` (`m_wTileID`, `m_dx`, `m_dy`, `m_dz`, `m_visible`)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `WORD` | `BlockNum` | Art ID; add 16384 to get the actual art tile number
`0x02` | `2` | `WORD` | `X` | X offset relative to multi origin
`0x04` | `2` | `WORD` | `Y` | Y offset relative to multi origin
`0x06` | `2` | `WORD` | `Alt` | Altitude offset
`0x08` | `4` | `UDWORD` | `Flags` | Item flags (see `tiledata_mul.md` for flag values)

### Component record — High Seas variant (14 bytes)

Source: `CUOMultiItemRec_HS`. Adds a 2-byte ship-access field at the end.

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `WORD` | `BlockNum` | Art ID; add 16384 to get the actual art tile number
`0x02` | `2` | `WORD` | `X` | X offset relative to multi origin
`0x04` | `2` | `WORD` | `Y` | Y offset relative to multi origin
`0x06` | `2` | `WORD` | `Alt` | Altitude offset
`0x08` | `4` | `UDWORD` | `Flags` | Item flags (see `tiledata_mul.md` for flag values)
`0x0C` | `2` | `WORD` | `ShipAccess` | High Seas ship access flags (`m_shipAccess`)

## Drawing

```
DrawX = LeftMostX + (X - Y) * 22 - (Width  / 2)
DrawY = TopMostY  + (X + Y) * 22 - Height - Alt * 4
```
