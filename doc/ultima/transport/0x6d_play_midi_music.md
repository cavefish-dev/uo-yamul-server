# (0x6D) PLAY_MIDI_MUSIC

Transmitted by the server to initiate music playback.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x6D`
`0x01` | `USHORT` | `musicID` | Music identifier

**Total size:** 3 bytes
