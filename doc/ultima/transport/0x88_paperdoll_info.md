# (0x88) Paperdoll Info

Sent by the server when the client opens a paperdoll (character or mobile equipment view). Contains the character's serial, name/title string, and display flags. Fixed size of 66 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x88`
`0x01` | `UINT` | `charUID` | Serial of the character whose paperdoll is being shown
`0x05` | `CHAR[60]` | `nameTitle` | Character name and optional title, null-terminated and zero-padded
`0x41` | `BYTE` | `flags` | Display flags: `0x01`=war mode active, `0x02`=can lift items

## Example packet (hex)

```
-- Paperdoll for charUID=0x00000001, name="Player", not in war mode
88 00 00 00 01 50 6C 61
79 65 72 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00
```
