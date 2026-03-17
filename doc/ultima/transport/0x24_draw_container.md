# (0x24) Draw Container

Sent by the server to open a container or game board on the client. 7 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x24`
`0x01` | `UINT` | `serial` | Serial of the container
`0x05` | `USHORT` | `model` | Model number of the container (e.g. `0x003C` = backpack)
