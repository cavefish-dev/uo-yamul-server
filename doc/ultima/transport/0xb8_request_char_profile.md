# (0xB8) Request Character Profile

Handles character profile requests and responses. The structure varies based on direction and mode.

## Client Request (read profile)

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB8`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `BYTE` | `mode` | CLIENT ONLY — does not exist in server message
`0x04` | `UINT` | `id` | Character serial

## Client Update Request (write profile)

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB8`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `BYTE` | `mode` | CLIENT ONLY
`0x04` | `BYTE[2]` | `cmdType` | `0x0001` to indicate update
`0x06` | `BYTE[2]` | `msglen` | Number of unicode characters in the profile
`0x08` | `BYTE[]` | `profile` | Updated profile text in unicode (not null-terminated)

## Server Response

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB8`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `UINT` | `id` | Character serial
`—` | `CHAR[]` | `charName` | Character name, null-terminated (not unicode)
`—` | `BYTE[2]` | `padding` | `0x0000` separator
`—` | `UNI[]` | `profile` | Profile text in unicode (may be empty)
`—` | `BYTE[2]` | `padding` | `0x0000` separator
`—` | `BYTE[2]` | `terminator` | `0x3300` end marker
