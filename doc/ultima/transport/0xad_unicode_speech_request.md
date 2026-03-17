# (0xAD) Unicode Speech Request

Clients send this packet when talking.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xAD`
`0x01` | `USHORT` | `size` | Packet size
`0x03` | `BYTE` | `mode` | Speech mode (see below)
`0x04` | `USHORT` | `hue` | Hue/color value
`0x06` | `USHORT` | `font` | Font number
`0x08` | `CHAR[4]` | `language` | Preferred language code (e.g. "ENU")
`0x0C` | `BYTE[]` | `KeywordInfo` | Present only if (mode & `0xC0`); packed 12-bit keyword integers
`—` | `UNI[]` | `text` | Spoken text, null-terminated (`0x0000`)

## Speech modes

Value | Description
--- | ---
`0x00` | Regular
`0x01` | Broadcast
`0x02` | Emote
`0x06` | System
`0x07` | Message
`0x08` | Whisper
`0x09` | Yell

The flag 0xC0 is set in mode if the message contains keyword information.

## KeywordInfo structure

KeywordInfo contains 12-bit integers packed into a byte array. The first 12-bit value is the keyword count: `(KeywordInfo[0] << 4) + (KeywordInfo[1] >> 4)`. Subsequent 12-bit values are individual keyword IDs, extracted sequentially using bit-shifting.

## Example packet (hex)

```
-- mode=0x0D, hue=0x0044, font=3, lang="IVL", text="asd" (unicode)
AD 00 14 0D 00 44 00 03 49 56 4C 00 00 61 00 73
00 64 00 00
```
