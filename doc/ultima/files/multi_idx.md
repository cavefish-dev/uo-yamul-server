# MULTI.IDX

Index to `MULTI.MUL` multi-object groups (houses, castles, ships, etc.). Each record is 12 bytes; seek to `MultiNum * 12`.

Reference: [Heptazane — Section 3.9](https://uo.stratics.com/heptazane/fileformats.shtml#3.9)

## File structure

Flat array of 12-byte index records.

### Index record

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Lookup` | Byte offset in `MULTI.MUL`; `0xFFFFFFFF` = undefined
`0x04` | `4` | `DWORD` | `Size` | Size of the multi block in bytes
`0x08` | `4` | `DWORD` | `Unknown` | Purpose unknown
