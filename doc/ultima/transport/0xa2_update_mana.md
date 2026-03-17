# (0xA2) Update Current Mana

Sent by the server to update a character's mana values on the client. Used to keep the mana bar current after casting spells or meditating. Fixed size of 9 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xA2`
`0x01` | `UINT` | `charUID` | Serial of the character whose mana is being updated
`0x05` | `USHORT` | `max` | Maximum mana
`0x07` | `USHORT` | `current` | Current mana

## Example packet (hex)

```
-- charUID=0x00000001, max mana=80, current mana=80 (full)
A2 00 00 00 01 00 50 00
50
```
