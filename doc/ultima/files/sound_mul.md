# SOUND.MUL

Sampled audio data: 16-bit mono PCM at 22,050 Hz. Accessed via `SOUNDIDX.MUL`.

Reference: [Heptazane — Section 3.15](https://uo.stratics.com/heptazane/fileformats.shtml#3.15)

## File structure

Each sound entry starts at the byte position given by `SOUNDIDX.MUL`.

### Sound entry

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `16` | `CHAR[16]` | `Filename` | Original source filename (null-padded)
`0x10` | `4` | `DWORD` | `Unknown1` | Purpose unknown
`0x14` | `4` | `DWORD` | `Unknown2` | Purpose unknown
`0x18` | `4` | `DWORD` | `Unknown3` | Purpose unknown
`0x1C` | `4` | `DWORD` | `Unknown4` | Purpose unknown
`0x20` | var | `BYTE[]` | `AudioData` | Raw PCM audio (16-bit mono, 22,050 Hz); length = `SOUNDIDX.Length - 32`
