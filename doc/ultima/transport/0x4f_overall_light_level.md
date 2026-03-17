# (0x4F) OVERALL_LIGHT_LEVEL

This is sent by the server to change the user's overall sunlight level. Fixed size of 2 bytes.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
`0x00` | `BYTE` | `cmd` | Command code, always `0x4F`
`0x01` | `BYTE` | `level` | Light intensity value

## Level values

Value | Description
--- | ---
`0x00` | Day (full light)
`0x09` | OSI night
`0x1F` | Black (maximum darkness)

## Example packet (hex)

```
-- level=0x18 (dim, not full dark)
4F 18
```
