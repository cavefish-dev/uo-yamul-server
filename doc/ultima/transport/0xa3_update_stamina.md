# (0xA3) Update Current Stamina

Sent by the server to update a character's stamina values on the client. Used to keep the stamina bar current during movement and combat. Fixed size of 9 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA3`
`0x01` | `UINT` | `charUID` | Serial of the character whose stamina is being updated
`0x05` | `USHORT` | `max` | Maximum stamina
`0x07` | `USHORT` | `current` | Current stamina

## Example packet (hex)

```
-- charUID=0x00000001, max stamina=100, current stamina=100 (full)
A3 00 00 00 01 00 64 00
64
```
