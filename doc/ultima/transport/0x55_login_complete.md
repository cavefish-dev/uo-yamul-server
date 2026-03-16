# (0x55) LOGIN_COMPLETE

This is sent by the server to notify the client that the login process is complete. Causes the game window to display and the game to start. Fixed size of 1 byte.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x55`
