# (0x9A) ASCII Prompt

Bidirectional packet for requesting and returning text input from the player. The server sends it to display a text input prompt; the client sends back the player's response. Variable length.

## Message format

### Server → Client

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x9A`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `UINT` | `contextUID1` | Context identifier part 1
`0x07` | `UINT` | `contextUID2` | Context identifier part 2
`0x0B` | `UINT` | `reserved` | Reserved; typically `0x00000000`
`0x0F` | `CHAR[...]` | `text` | Prompt text to display, null-terminated

### Client → Server

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x9A`
`0x01` | `USHORT` | `size` | Total packet length in bytes
`0x03` | `UINT` | `contextUID1` | Context identifier part 1 (echoed from server packet)
`0x07` | `UINT` | `contextUID2` | Context identifier part 2 (echoed from server packet)
`0x0B` | `UINT` | `responseType` | Response type: `0x00000000`=cancel, `0x00000001`=respond
`0x0F` | `CHAR[...]` | `userInput` | Player's text response, null-terminated (max 128 characters)
