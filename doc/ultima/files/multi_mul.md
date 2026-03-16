# MULTI.MUL

Multi-object data: grouped art pieces composing houses, castles, ships, etc. Accessed via `MULTI.IDX`.

Reference: [Heptazane — Section 3.10](https://uo.stratics.com/heptazane/fileformats.shtml#3.10)

## File structure

Each multi entry is a sequence of 12-byte component records at the offset from `MULTI.IDX`.

### Component record (12 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `WORD` | `BlockNum` | Art ID; add 16384 to get the actual art tile number
`0x02` | `2` | `WORD` | `X` | X offset relative to multi origin
`0x04` | `2` | `WORD` | `Y` | Y offset relative to multi origin
`0x06` | `2` | `WORD` | `Alt` | Altitude offset
`0x08` | `4` | `UDWORD` | `Flags` | Item flags (see `tiledata_mul.md` for flag values)

## Drawing

```
DrawX = LeftMostX + (X - Y) * 22 - (Width  / 2)
DrawY = TopMostY  + (X + Y) * 22 - Height - Alt * 4
```
