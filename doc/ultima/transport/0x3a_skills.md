# (0x3A) SKILLS

This is sent by the server to display and update skills. The client also sends this packet to alter skill locks. Variable length.

## Server-to-client: Send Skills

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x3A`
`0x01` | `USHORT` | `blockSize` | Total packet size in bytes
`0x03` | `BYTE` | `type` | Format type (see values below)

Each skill entry follows immediately after the header and repeats for every skill:

Offset (per skill) | Type | Name | Description
--- | --- | --- | ---
`+0x00` | `USHORT` | `skillId` | 1-based skill number
`+0x02` | `USHORT` | `value` | Current skill value
`+0x04` | `USHORT` | `basevalue` | Value shown as "real" skill
`+0x06` | `BYTE` | `status` | Lock state: 0 = up, 1 = down, 2 = locked
`+0x07` | `USHORT` | `maxvalue` | Skill cap (only present when type includes cap, see below)

## Type values

Value | Description
--- | ---
`0x00` | Basic (no cap)
`0x01` | God view (no cap)
`0x02` | Basic + skill cap
`0x03` | God view + skill cap
`0xDF` | Single skill update + skill cap
`0xFF` | Single skill update (no cap)

## Client-to-server: Set Skill Lock

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x3A`
`0x01` | `USHORT` | `blockSize` | Total packet size in bytes
`0x03` | `USHORT` | `skillId` | Target skill identifier
`0x05` | `BYTE` | `status` | Lock state: 0 = up, 1 = down, 2 = locked
