# (0x2B) Toggle GodMode Reply

Sent by the server to enable or disable god mode. Should be sent in response to a Request God Mode (0x04) packet from the client. 2 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x2B`
`0x01` | `BOOL` | `godmode` | If true, god mode is enabled and access to god client features is allowed
