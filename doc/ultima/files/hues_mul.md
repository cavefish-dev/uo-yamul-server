# HUES.MUL

Color hue definitions used to recolor artwork. The file begins with a 768-byte header followed by hue group entries.

Reference: [Heptazane — Section 3.7](https://uo.stratics.com/heptazane/fileformats.shtml#3.7)

## File structure

The file is a sequence of 544-byte `HueGroup` records (after the 768-byte header).

### HueGroup (544 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Header` | Group header value
`0x04` | `32*8` | `HueEntry[8]` | `Entries` | 8 hue entries per group

### HueEntry (32 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `WORD[2]` | `ColorTable` | Not 32 entries — 2 colors only per spec (see note)
`0x04` | `2` | `WORD` | `TableStart` | First palette index this hue replaces
`0x06` | `2` | `WORD` | `TableEnd` | Last palette index this hue replaces
`0x08` | `20` | `CHAR[20]` | `Name` | Hue name (null-terminated)

> **Note:** The reference gives `WORD[32]` ColorTable (64 bytes) making each HueEntry 88 bytes and HueGroup 708 bytes — inconsistency in the source. Use empirical file size to confirm.

## Hue application methods

1. **Range mapping** — replace pixels whose color falls between `TableStart` and `TableEnd` with the corresponding `ColorTable` entry.
2. **Full grayscale** — convert all non-black pixels using the hue's color table.
