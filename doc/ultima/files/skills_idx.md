# SKILLS.IDX

Index to `SKILLS.MUL` skill name entries. Each record is 12 bytes; seek to `SkillNum * 12`.

Reference: [Heptazane — Section 3.13](https://uo.stratics.com/heptazane/fileformats.shtml#3.13)

## File structure

Flat array of 12-byte index records.

### Index record

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Start` | Byte position in `SKILLS.MUL`
`0x04` | `4` | `DWORD` | `Length` | Size of data segment in bytes
`0x08` | `4` | `DWORD` | `Unknown` | Purpose unknown
