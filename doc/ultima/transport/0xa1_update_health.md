# (0xA1) Update Current Health

Sent by the server to update a character's hit point values on the client. Used to keep the health bar current during combat and healing. Fixed size of 9 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA1`
`0x01` | `UINT` | `charUID` | Serial of the character whose health is being updated
`0x05` | `USHORT` | `max` | Maximum hit points
`0x07` | `USHORT` | `current` | Current hit points

## Example packet (hex)

```
-- charUID=0x00000001, max HP=100, current HP=75
A1 00 00 00 01 00 64 00
4B
```
