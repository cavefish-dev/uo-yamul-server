# (0x93) Open Book

Sent by the server to open a book gump on the client. Contains the book's serial, write permissions, page count, title, and author. Fixed size of 99 bytes.

Note: this packet is sometimes listed as "Update Map Pins" in older references, but that is incorrect — 0x93 is the Open Book packet.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x93`
`0x01` | `UINT` | `bookUID` | Serial of the book item
`0x05` | `BYTE` | `writable` | `0x01` if the player can write in this book, `0x00` otherwise
`0x06` | `BYTE` | `writable2` | Duplicate writable flag (legacy field)
`0x07` | `USHORT` | `pageCount` | Number of pages in the book
`0x09` | `CHAR[60]` | `title` | Book title, null-terminated and zero-padded
`0x45` | `CHAR[30]` | `author` | Book author name, null-terminated and zero-padded
