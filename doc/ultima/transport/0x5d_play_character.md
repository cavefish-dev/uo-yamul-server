# (0x5D) PLAY_CHARACTER

This is sent by the client after the user picks a character to login with. Fixed size of 0x49 (73) bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x5D`
`0x01` | `UINT` | `unknown` | Always `0xEDEDEDED`
`0x05` | `CHAR[30]` | `name` | Character name
`0x23` | `CHAR[30]` | `password` | Character password
`0x41` | `UINT` | `slot` | Character slot position in the character list
`0x45` | `UINT` | `key` | User's encryption key

## Example packet (hex)

```
-- name="asdf", password="test", slot=0x01020304, key=0x0100007F
5D ED ED ED ED 61 73 64 66 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 00 00 74 65 73 74 00 00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
00 01 02 03 04 01 00 00 7F
```
