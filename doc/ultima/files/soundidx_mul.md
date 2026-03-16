# SOUNDIDX.MUL

Index to `SOUND.MUL` audio entries. Each record is 12 bytes; seek to `SoundNum * 12`.

Reference: [Heptazane — Section 3.16](https://uo.stratics.com/heptazane/fileformats.shtml#3.16)

## File structure

Flat array of 12-byte index records.

### Index record

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Start` | Byte position in `SOUND.MUL`
`0x04` | `4` | `DWORD` | `Length` | Size of data segment in bytes (includes 32-byte header)
`0x08` | `2` | `UWORD` | `Index` | Sound index number
`0x0A` | `2` | `UWORD` | `Reserved` | Reserved; ignore
