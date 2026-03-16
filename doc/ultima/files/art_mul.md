# ART.MUL

Tile graphics for ground tiles and static objects. Accessed via `ARTIDX.MUL`. Two sub-formats exist: **Raw** (land tiles) and **Run** (static/object tiles).

Reference: [Heptazane — Section 3.4](https://uo.stratics.com/heptazane/fileformats.shtml#3.4)

## File structure

Each tile begins at the offset from `ARTIDX.MUL` and starts with a 4-byte header flag.

### Header

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Flag` | Format selector: `Flag > 0xFFFF` or `Flag == 0` → Raw; otherwise → Run

---

### Raw tile (land tiles, always 44×44 pixels)

Pixels follow a diamond pattern rotated 45°. Row pixel counts: 2, 4, 6 … 44, 44, 42, 40 … 2.

Each pixel is a `WORD` (16-bit color, see [`types.md`](types.md)).

---

### Run tile (static/object tiles)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `2` | `UWORD` | `Width` | Tile width in pixels
`0x02` | `2` | `UWORD` | `Height` | Tile height in pixels
`0x04` | `2*H` | `UWORD[H]` | `LStart` | Per-row start offset table (relative to DataStart)
`0x04+2H` | var | — | `Data` | Run-length encoded rows (see below)

**Run tile decompression:**

```
DataStart = &Data[0]
X = 0
Y = 0
Seek to DataStart + LStart[0] * 2

while Y < Height:
  Read UWORD XOffs
  Read UWORD Run

  if (XOffs + Run) >= 2048: return   // sentinel / bounds guard
  if (XOffs + Run) != 0:
    X += XOffs
    // read Run pixels (WORD each) at column X
    X += Run
  else:
    X = 0
    Y++
    if Y < Height:
      Seek to DataStart + LStart[Y] * 2
```
