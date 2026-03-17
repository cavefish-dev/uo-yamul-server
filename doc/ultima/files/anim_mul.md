# ANIM.MUL

Raw animation data for monsters, people, and equipment. Accessed via `ANIM.IDX`.

Reference: [Heptazane — Section 3.2](https://uo.stratics.com/heptazane/fileformats.shtml#3.2)

## File structure

Each animation group begins at the offset given by `ANIM.IDX`.

### Animation group

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x000` | `512` | `WORD[256]` | `Palette` | 256 palette color entries
`0x200` | `4` | `DWORD` | `FrameCount` | Number of frames
`0x204` | `4*N` | `DWORD[N]` | `FrameOffsets` | Per-frame offset from start of group

### Frame

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `WORD` | `ImageCenterX` | Horizontal center point
`0x02` | `2` | `WORD` | `ImageCenterY` | Vertical center point
`0x04` | `2` | `WORD` | `Width` | Frame width in pixels
`0x06` | `2` | `WORD` | `Height` | Frame height in pixels
`0x08` | var | — | `Data` | Run-length encoded pixel stream (see below)

### Frame data decompression

```
PrevLineNum = 0xFF
Y = CenterY - ImageCenterY - Height

while not EOF:
  Read UWORD RowHeader
  Read UWORD RowOfs

  if (RowHeader == 0x7FFF) or (RowOfs == 0x7FFF): break

  RunLength = RowHeader & 0x0FFF
  LineNum   = RowHeader >> 12
  RowOfs    = RowOfs >> 6          // upper 10 bits = signed X offset from CenterX
  X         = CenterX + RowOfs

  if (PrevLineNum != 0xFF) and (LineNum != PrevLineNum): Y++
  PrevLineNum = LineNum

  if Y >= 0 and Y < MaxHeight:
    for I in 0 .. RunLength-1:
      Read UBYTE B
      Pixel[X+I, Y] = Palette[B]
```
