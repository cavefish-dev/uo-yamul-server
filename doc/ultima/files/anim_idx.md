# ANIM.IDX

Index file for accessing animation frames in `ANIM.MUL`. Each record is 12 bytes; seek to `BNum * 12` to reach a record, then use `Lookup` as the file offset into `ANIM.MUL`.

Reference: [Heptazane — Section 3.1](https://uo.stratics.com/heptazane/fileformats.shtml#3.1)

## File structure

The file is a flat array of 12-byte index records.

### Index record

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Lookup` | Byte offset in `ANIM.MUL`; `0xFFFFFFFF` = undefined
`0x04` | `4` | `DWORD` | `Size` | Size of the animation data block
`0x08` | `4` | `DWORD` | `Unknown` | Purpose unknown
