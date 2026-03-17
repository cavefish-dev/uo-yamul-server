# PALETTE.MUL

8-bit color palette used with certain graphic formats. Fixed size: 768 bytes (256 × 3).

Reference: [Heptazane — Section 3.11](https://uo.stratics.com/heptazane/fileformats.shtml#3.11)

## File structure

256 sequential 3-byte palette entries.

### Palette entry (3 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `1` | `UBYTE` | `Red` | Red component (0–255)
`0x01` | `1` | `UBYTE` | `Green` | Green component (0–255)
`0x02` | `1` | `UBYTE` | `Blue` | Blue component (0–255)
