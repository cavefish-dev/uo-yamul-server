# (0xAB) Gump Text Entry Dialog

This is used to elicit a string or numerical response from the client.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xAB`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `UINT` | `serial` | The query serial
`0x07` | `BYTE` | `parent_id` | Parent identifier
`0x08` | `BYTE` | `button_id` | Button identifier
`0x09` | `USHORT` | `length1` | The length of the text that follows
`0x0B` | `CHAR[]` | `text1` | Text to appear at the top of the gump
`—` | `BOOL` | `cancel` | If true, the client is able to cancel the query
`—` | `BYTE` | `type` | The type of query: `0x01` - String; `0x02` - Number
`—` | `UINT` | `max` | Maximum length (for string) or maximum number (for number query)
`—` | `USHORT` | `length2` | The length of the text that follows
`—` | `CHAR[]` | `text2` | Text to appear above the entry box
