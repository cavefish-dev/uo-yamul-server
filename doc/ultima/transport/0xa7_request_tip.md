# (0xA7) Request Tip/Notice

Sent by the client to request the previous or next tip of the day (or notice). The server responds with packet 0xA6. Fixed size of 4 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA7`
`0x01` | `USHORT` | `tipIndex` | Index of the currently displayed tip
`0x03` | `BYTE` | `direction` | Navigation direction: `0x00`=previous tip, `0x01`=next tip

## Example packet (hex)

```
-- Request next tip after tip index 1
A7 00 01 01

-- Request previous tip before tip index 3
A7 00 03 00
```
