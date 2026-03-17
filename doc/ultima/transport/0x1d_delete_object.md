# (0x1D) Delete Object

Sent by the server to remove an item from the player's sight. Also sent by the god client to delete a dynamic item. 5 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x1D`
`0x01` | `UINT` | `serial` | Serial of the object to remove
