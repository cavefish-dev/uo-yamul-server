# (0xB9) Client Features

Enables various client-side features. Should be sent before the Character List packet (0xA9).

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB9`
`0x01` | `USHORT` | `flags` | Feature flag bitmask (see below)

## Feature flags

Value | Description
--- | ---
`0x0001` | Chat
`0x0002` | LBR (Lord Blackthorn's Revenge) animations
`0x0010` | Create Paladin/Necromancer
`0x0020` | Sixth character slot
`0x8000` | More features — must be present to enable Age of Shadows features or sixth character slot
