# (0x56) MAP_COMMAND

This is sent by the client and server to manage map pins and plot courses, primarily used for ship navigation. Fixed size of 11 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x56`
`0x01` | `UINT` | `serial` | Map item serial number
`0x05` | `BYTE` | `command` | Operation type (see values below)
`0x06` | `BYTE` | `pin` | 0-based pin number; also used as a boolean for request/toggle edit mode commands
`0x07` | `USHORT` | `pinX` | Pixel distance from the left edge of the map
`0x09` | `USHORT` | `pinY` | Pixel distance from the top edge of the map

## Command values

Value | Description
--- | ---
`0x01` | Add pin
`0x03` | Move pin
`0x04` | Delete pin
`0x05` | Open map
`0x06` | Request edit mode
`0x07` | Toggle edit mode
