# (0xA0) Select Server

Sent by the client to choose a game server from the server list (received via packet 0xA8). The server responds with a relay packet (0x8C) directing the client to the chosen game server. Fixed size of 3 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA0`
`0x01` | `USHORT` | `serverIndex` | Zero-based index of the selected server from the server list

## Example packet (hex)

```
-- Select first server (index 0)
A0 00 00

-- Select second server (index 1)
A0 00 01
```
