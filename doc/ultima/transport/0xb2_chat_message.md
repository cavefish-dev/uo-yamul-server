# (0xB2) Chat Message

Chat messaging packet. The structure varies depending on the message type.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB2`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `USHORT` | `messagetype` | Message type identifier (see subtypes below)

## Subtype 0x03EB — Display Enter Username Window

Offset | Type | Name | Description
--- | --- | --- | ---
`0x05` | `BYTE[8]` | `unknown` | All `0x00`

Alternative format (with username):

Offset | Type | Name | Description
--- | --- | --- | ---
`0x05` | `UNI[30]` | `username` | Unicode username
`0x65` | `UINT` | `unknown` | All `0x00` (string terminator)

## Subtype 0x03ED — Username Accepted / Display Window

No detailed field structure documented.

## Subtype 0x03E8 — Channel Join

Offset | Type | Name | Description
--- | --- | --- | ---
`0x05` | `BYTE[4]` | `unknown` | All `0x00`
`0x09` | `UNI[]` | `channel` | Unicode channel name (variable length)
`—` | `BYTE[2]` | `unknown` | `0x0000`
`—` | `BYTE[2]` | `unknown` | `0x0030`
`—` | `BYTE[2]` | `unknown` | `0x0000`
