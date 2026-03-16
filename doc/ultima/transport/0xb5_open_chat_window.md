# (0xB5) Open Chat Window

Opens the chat window. This message is very incomplete. From the server, just know that it is 0xB5 len len, and pass the data through as is appropriate.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB5`
`0x01` | `BYTE[63]` | `chatname` | Chat name, if known by client (all `0x00` if unknown); name in unicode
