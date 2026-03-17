# ARTIDX.MUL

Index to `ART.MUL` tiles. Each record is 12 bytes; seek to `TileNum * 12`, read the record, then use `Lookup` as the byte offset into `ART.MUL`.

Reference: [Heptazane — Section 3.3](https://uo.stratics.com/heptazane/fileformats.shtml#3.3)

## File structure

Flat array of 12-byte index records.

### Index record

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Lookup` | Byte offset in `ART.MUL`; `0xFFFFFFFF` = undefined
`0x04` | `4` | `DWORD` | `Size` | Size of the art tile data
`0x08` | `4` | `DWORD` | `Unknown` | Purpose unknown
