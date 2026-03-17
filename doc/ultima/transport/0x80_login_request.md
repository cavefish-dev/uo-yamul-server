# (0x80) Login Request

Sent by the client to the login server to authenticate. Contains the account username and password in plaintext. Fixed size of 61 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x80`
`0x01` | `CHAR[30]` | `username` | Account username, null-terminated and zero-padded
`0x1F` | `CHAR[30]` | `password` | Account password, null-terminated and zero-padded

## Example packet (hex)

```
-- Login: username="admin", password="hunter2"
80 61 64 6D 69 6E 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 68
75 6E 74 65 72 32 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00
```
