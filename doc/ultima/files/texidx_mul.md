# TEXIDX.MUL

Index to `TEXMAPS.MUL` ground texture entries. Each record is 12 bytes; seek to `TexNum * 12`.

Reference: [Heptazane — Section 3.20](https://uo.stratics.com/heptazane/fileformats.shtml#3.20)

## File structure

Flat array of 12-byte index records.

### Index record

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Start` | Byte position in `TEXMAPS.MUL`
`0x04` | `4` | `DWORD` | `Length` | Data segment size in bytes (determines texture dimensions)
`0x08` | `4` | `DWORD` | `Unknown` | Purpose unknown
