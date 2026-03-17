# (0x8C) Connect to Game Server

Sent by the login server after successful authentication to redirect the client to the appropriate game server. Contains the game server's IP address, port, and an auth token used for the subsequent game server login. Fixed size of 11 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x8C`
`0x01` | `BYTE` | `ip0` | First octet of the game server IP address
`0x02` | `BYTE` | `ip1` | Second octet of the game server IP address
`0x03` | `BYTE` | `ip2` | Third octet of the game server IP address
`0x04` | `BYTE` | `ip3` | Fourth octet of the game server IP address
`0x05` | `USHORT` | `port` | Game server TCP port
`0x07` | `UINT` | `authID` | Authentication token to present to the game server in packet 0x91

## Example packet (hex)

```
-- Redirect to game server at 192.168.1.100:2593, authID=0x12345678
8C C0 A8 01 64 0A 21 12
34 56 78
```
