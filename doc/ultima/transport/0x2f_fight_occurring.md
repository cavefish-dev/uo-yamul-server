# (0x2F) Fight Occurring

Sent by the server during on-screen combat to signal the client to execute an attack swing animation. 10 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x2F`
`0x01` | `BYTE` | `unknown1` | Always `0x00`
`0x02` | `UINT` | `attackerID` | Serial of the attacking entity
`0x06` | `UINT` | `attackedID` | Serial of the target entity
