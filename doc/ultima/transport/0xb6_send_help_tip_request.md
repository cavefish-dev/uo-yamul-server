# (0xB6) Send Help/Tip Request

Request for a help or tip page.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xB6`
`0x01` | `UINT` | `id` | Help/tip identifier
`0x05` | `BYTE` | `language` | Language code; 1 for English
`0x06` | `BYTE[3]` | `languagename` | Language name string (e.g. "ENU" for English/United States)
