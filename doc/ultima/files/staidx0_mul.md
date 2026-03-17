# STAIDX0.MUL

Index to `STATICS0.MUL` static object data. Fixed size: 4,718,592 bytes (393,216 × 12). One index record per map block, in 1:1 correspondence with `MAP0.MUL` blocks.

Reference: [Heptazane — Section 3.17](https://uo.stratics.com/heptazane/fileformats.shtml#3.17)

## File structure

Flat array of 12-byte index records. Use the same block-number formula as `MAP0.MUL` to locate a record:

```
BlockNum   = (XBlock * 512) + YBlock
FileOffset = BlockNum * 12
```

### Index record

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Start` | Byte position in `STATICS0.MUL`; `0xFFFFFFFF` = empty block
`0x04` | `4` | `DWORD` | `Length` | Size of statics data for this block in bytes
`0x08` | `4` | `DWORD` | `Unknown` | Purpose unknown
