# (0x00) Create Character

This message is sent from UO client when user selects to create a new character.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x00`
`0x01` | `UINT` | `pattern1` | Always `0xedededed`
`0x05` | `UINT` | `pattern2` | Always `0xffffffff`
`0x09` | `BYTE` | `pattern3` | Always `0x00`
`0x10` | `CHAR[30]` | `char_name` | Character name
`0x40` | `CHAR[30]` | `char_password` | Character password
`0x70` | `BYTE` | `gender` | 0=male, 1=female
`0x71` | `BYTE` | `str` | Strength attribute
`0x72` | `BYTE` | `dex` | Dexterity attribute
`0x73` | `BYTE` | `int` | Intelligence attribute
`0x74` | `BYTE` | `skill1` | First skill index
`0x75` | `BYTE` | `skillVal1` | First skill value
`0x76` | `BYTE` | `skill2` | Second skill index
`0x77` | `BYTE` | `skillVal2` | Second skill value
`0x78` | `BYTE` | `skill3` | Third skill index
`0x79` | `BYTE` | `skillVal3` | Third skill value
`0x80` | `USHORT` | `skinColor` | Character skin tone
`0x82` | `USHORT` | `hairStyle` | Hair style selection
`0x84` | `USHORT` | `hairColor` | Hair color
`0x86` | `USHORT` | `facialhairStyle` | Facial hair style
`0x88` | `USHORT` | `facialhairColor` | Facial hair color
`0x90` | `USHORT` | `location` | Starting location
`0x92` | `USHORT` | `unknown1` | Unknown data
`0x94` | `USHORT` | `slot` | Character slot (max 6)
`0x96` | `UINT` | `clientIP` | Client IP address
`0x9A` | `USHORT` | `shirtColor` | Shirt color
`0x9C` | `USHORT` | `pantsColor` | Pants color
