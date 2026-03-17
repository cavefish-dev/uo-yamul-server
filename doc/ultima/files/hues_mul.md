# HUES.MUL

Color hue definitions used to recolor artwork. The file begins with a 768-byte header followed by hue group entries.

Reference: [Heptazane — Section 3.7](https://uo.stratics.com/heptazane/fileformats.shtml#3.7) · [SphereServer — uo_files/CUOHues.h](https://github.com/Sphereserver/Source-X/blob/master/src/game/uo_files/CUOHues.h)

## File structure

The file is a sequence of 708-byte `HueGroup` records (after the 768-byte header).

### HueGroup (708 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Header` | Group header value
`0x04` | `88*8` | `HueEntry[8]` | `Entries` | 8 hue entries per group (8 × 88 = 704 bytes)

### HueEntry (88 bytes)

Source: `CUOHuesRec` (`m_color[34]`, `m_name[20]`)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `68` | `WORD[34]` | `ColorTable` | 34 color entries for hue mapping (16-bit colors)
`0x44` | `20` | `CHAR[20]` | `Name` | Hue name (null-padded)

## Hue application methods

1. **Range mapping** — replace pixels whose palette index maps into the hue's `ColorTable` with the corresponding entry.
2. **Full grayscale** — convert all non-black pixels using the hue's color table.
