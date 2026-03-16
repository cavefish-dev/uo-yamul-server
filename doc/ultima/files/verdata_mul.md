# VERDATA.MUL

Patch override file. Entries here take precedence over the corresponding records in other MUL files. Used to distribute post-release patches without replacing entire data files.

Reference: [Heptazane — Section 3.22](https://uo.stratics.com/heptazane/fileformats.shtml#3.22)

## File structure

Starts with a 4-byte entry count, followed by that many 20-byte index entries, then the raw patch data blocks.

### File header

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `EntryCount` | Number of patch index entries

### Index entry (20 bytes)

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `FileID` | Source file identifier (see table below)
`0x04` | `4` | `DWORD` | `Block` | Item or block number within that file
`0x08` | `4` | `DWORD` | `Position` | Byte offset of the patch data within `VERDATA.MUL`
`0x0C` | `4` | `DWORD` | `Size` | Patch block size in bytes
`0x10` | `4` | `DWORD` | `Various` | Format-specific data (see notes)

### FileID values

| FileID | File |
| --- | --- |
| `0x00` | `map0.mul` |
| `0x01` | `staidx0.mul` |
| `0x02` | `statics0.mul` |
| `0x03` | `artidx.mul` |
| `0x04` | `art.mul` ✱ |
| `0x05` | `anim.idx` |
| `0x06` | `anim.mul` |
| `0x07` | `soundidx.mul` |
| `0x08` | `sound.mul` |
| `0x09` | `texidx.mul` |
| `0x0A` | `texmaps.mul` |
| `0x0B` | `gumpidx.mul` |
| `0x0C` | `gumpart.mul` ✱ |
| `0x0D` | `multi.idx` |
| `0x0E` | `multi.mul` ✱ |
| `0x0F` | `skills.idx` |
| `0x10` | `skills.mul` |
| `0x1E` | `tiledata.mul` ✱ |
| `0x1F` | `animdata.mul` ✱ |

✱ = actively used in VERDATA patches

### Notes

- **FileID `0x0C` (gumpart.mul):** The `Various` field encodes `WORD Height | (WORD Width << 16)`.
- Patch data for a given entry is located at byte offset `Position` within `VERDATA.MUL` and has the same format as the corresponding record in the target file.
