# (0xBC) Seasonal Information

Server message that communicates seasonal changes to the client.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xBC`
`0x01` | `BYTE` | `id1` | Season identifier (used when id2 == 1)
`0x02` | `BYTE` | `id2` | If value equals 1, triggers a season change

## Season values (id1)

Value | Description
--- | ---
`0` | Spring
`1` | Summer
`2` | Fall
`3` | Winter
`4` | Desolation
