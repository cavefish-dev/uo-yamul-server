# (0x6E) CHARACTER_ANIMATION

Sent by the server to display an animation on a character.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x6E`
`0x01` | `UINT` | `serial` | Character ID
`0x05` | `USHORT` | `animationId` | Animation identifier (`0x00`–`0x24`)
`0x07` | `BYTE` | `unknown1` | Unknown field (always `0x00`)
`0x08` | `BYTE` | `direction` | Direction value
`0x09` | `USHORT` | `repeat` | Repetition mode (0=infinite, 1=once, 2=twice)
`0x0B` | `BYTE` | `fw_bw` | Playback direction (`0x00`=forward, `0x01`=backwards)
`0x0C` | `BYTE` | `repeatFlag` | Repeat control (0=don't repeat, 1=repeat)
`0x0D` | `BYTE` | `frameDelay` | Speed control (`0x00`=fastest, `0xFF`=slowest)

**Total size:** 14 bytes

## Animation ID categories

The animation ID field supports 37 types (0x00–0x24), including movement animations (walk, run), combat actions (attacks, blocks), emotes (bow, salute), mounted actions, and death sequences.
