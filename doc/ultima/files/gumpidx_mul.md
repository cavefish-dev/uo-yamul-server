# GUMPIDX.MUL

Index to `GUMPART.MUL` gump entries. Each record is 12 bytes; seek to `GumpNum * 12`.

Reference: [Heptazane — Section 3.6](https://uo.stratics.com/heptazane/fileformats.shtml#3.6)

## File structure

Flat array of 12-byte index records.

### Index record

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Lookup` | Byte offset in `GUMPART.MUL`; `0xFFFFFFFF` = undefined
`0x04` | `4` | `DWORD` | `Size` | Size of the gump data block in bytes
`0x08` | `2` | `UWORD` | `Height` | Image height in pixels
`0x0A` | `2` | `UWORD` | `Width` | Image width in pixels
