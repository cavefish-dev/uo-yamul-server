# (0x83) Delete Character

Sent by the client to request deletion of a character slot. The password and client IP fields are present for historical reasons but are ignored by the server. Fixed size of 39 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x83`
`0x01` | `CHAR[30]` | `password` | Account password (ignored by server), null-terminated and zero-padded
`0x1F` | `UINT` | `slot` | Zero-based character slot index to delete
`0x23` | `UINT` | `clientIP` | Client IP address (ignored by server)

## Example packet (hex)

```
-- Delete character in slot 0, client IP=127.0.0.1
83 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 7F 00 00 01
```
