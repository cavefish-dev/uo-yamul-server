# (0x91) Game Server Login

Sent by the client to the game server immediately after connecting, presenting the auth token received from the login server (via packet 0x8C) along with credentials. Fixed size of 65 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x91`
`0x01` | `UINT` | `authID` | Authentication token received from the login server in packet 0x8C
`0x05` | `CHAR[30]` | `username` | Account username, null-terminated and zero-padded
`0x23` | `CHAR[30]` | `password` | Account password, null-terminated and zero-padded

## Example packet (hex)

```
-- Game server login: authID=0x12345678, username="admin", password="hunter2"
91 12 34 56 78 61 64 6D
69 6E 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 68
75 6E 74 65 72 32 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
00 00 00 00 00 00 00 00
01
```
