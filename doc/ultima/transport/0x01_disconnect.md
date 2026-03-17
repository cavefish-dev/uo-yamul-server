# (0x01) DISCONNECT

This is a message sent from client to server when the user chooses to return to the main menu from the character select menu. Since the character select menu no longer has a main menu button, this message is no longer used.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x01`
`0x01` | `BYTE[4]` | `pattern` | Pattern used for validation, always `0xFF` `0xFF` `0xFF` `0xFF`