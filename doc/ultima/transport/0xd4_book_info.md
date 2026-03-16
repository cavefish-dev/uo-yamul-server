# (0xD4) Book Info

Sent by the server to open a book, and by the client to change the title or author. Older clients use the Book Header packet (0x93) instead.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xD4`
`0x01` | `UINT` | `bookID` | Book serial
`0x05` | `BOOL` | `iseditable` | Whether the book can be edited by the player
`0x06` | `BOOL` | `isnew` | Whether this is a new (blank) book
`0x07` | `USHORT` | `pages` | Number of pages in the book
`0x09` | `USHORT` | `title_length` | Length of the title string
`0x0B` | `CHAR[]` | `title` | Book title, null-terminated (`0x00`)
`—` | `USHORT` | `author_length` | Length of the author string
`—` | `CHAR[]` | `author` | Book author, null-terminated (`0x00`)
