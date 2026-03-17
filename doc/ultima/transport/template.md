# (0xNN) PACKET_NAME

Description of the packet and when it is used.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0xNN`
`0x01` | `USHORT` | `field2` | Description of field2

## Example packet (hex)

(Note: Format must be in hex bytes, 8 columns, with spaces between bytes. Example: `NN 01 00 ...`)

```
-- A comment describing the packet contents
NN 01 00 ...

-- Another packet example (add an empty line between examples)
NN 02 00 ...
```