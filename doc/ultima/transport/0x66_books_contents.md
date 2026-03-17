# (0x66) BOOKS_CONTENTS

Sent by the server to display book contents, and by the client to edit book contents. Variable length.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x66`
`0x01` | `USHORT` | `blockSize` | Total size of the packet
`0x03` | `UINT` | `bookID` | The book's serial
`0x07` | `USHORT` | `pages` | Number of pages being sent
`0x09` | `USHORT` | `pageno` | Page number (1-based)
`--` | `SHORT` | `lines` | Number of lines on this page
`--` | `CHAR[]` | `text` | Line text, null terminated (repeats per line, per page)

## Notes

- The `lines` and `text` fields repeat within a loop for each page.
- A `lines` value of -1 indicates no lines (server usage) or requests a page (client usage, deprecated).
- Client values greater than 0 write page content.
