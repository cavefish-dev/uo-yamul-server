# (0x11) Stat Window Info

Sent by the server to provide character stat information displayed in the status window. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x11`
`0x01` | `USHORT` | `size` | Packet length (varies)
`0x03` | `UINT` | `playerID` | Player identifier
`0x07` | `CHAR[30]` | `playerName` | Character name
`0x25` | `USHORT` | `currentHitpoints` | Current hit points
`0x27` | `USHORT` | `maxHitpoints` | Maximum hit points
`0x29` | `BOOL` | `flagName` | `0xFF` = allowed, `0x00` = not allowed
`0x2A` | `BYTE` | `flagDisplay` | Display flag; if non-zero, additional fields follow
`0x2B` | `BYTE` | `gender` | 0=male, 1=female
`0x2C` | `USHORT` | `str` | Strength
`0x2E` | `USHORT` | `dex` | Dexterity
`0x30` | `USHORT` | `int` | Intelligence
`0x32` | `USHORT` | `currentStamina` | Current stamina
`0x34` | `USHORT` | `maxStamina` | Maximum stamina
`0x36` | `USHORT` | `currentMana` | Current mana
`0x38` | `USHORT` | `maxMana` | Maximum mana
`0x3A` | `UINT` | `gold` | Gold amount
`0x3E` | `USHORT` | `resistPhysical` | Physical resistance / armor class
`0x40` | `USHORT` | `weight` | Current weight
`0x42` | `USHORT` | `maxWeight` | Maximum weight capacity
`0x44` | `BYTE` | `race` | Character race
`0x45` | `USHORT` | `statCap` | Total allowable sum of str + dex + int
`0x47` | `BYTE` | `currentFollowers` | Follower slots currently used
`0x48` | `BYTE` | `maximumFollowers` | Maximum follower slots
`0x49` | `USHORT` | `resistFire` | Fire resistance
`0x4B` | `USHORT` | `resistCold` | Cold resistance
`0x4D` | `USHORT` | `resistPoison` | Poison resistance
`0x4F` | `USHORT` | `resistEnergy` | Energy resistance
`0x51` | `USHORT` | `luck` | Luck
`0x53` | `USHORT` | `minDamage` | Minimum damage dealt
`0x55` | `USHORT` | `maxDamage` | Maximum damage dealt
`0x57` | `UINT` | `tithingPoints` | Tithing points balance
