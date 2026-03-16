# (0x05) Attack Request

Client message sent when user tries to attack someone.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x05`
`0x01` | `UINT` | `serial` | Serial of the target to be attacked
