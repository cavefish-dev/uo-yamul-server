# UO Wire Types

This file maps UO protocol type names (as used in the packet documentation on this site and the [WolfPack reference](https://www.hoogi.de/wolfpack/wiki/doku.php?id=uo_protocol)) to their Go equivalents in the gateway.

All multi-byte values use **big-endian (network) byte order**.

## Type mapping

| UO Type | Size (bytes) | Go Type | Notes |
|---------|-------------|---------|-------|
| BYTE | 1 | `byte` | Unsigned 8-bit integer |
| SBYTE | 1 | `int8` | Signed 8-bit integer |
| BOOL | 1 | `bool` | 0x00 = false, 0xFF = true |
| CHAR | 1 | `byte` | Single ASCII character |
| CHAR[n] | n | `string` | Fixed-length null-padded ASCII string; read via `ReadFixedString(n)` |
| UNI | 2 | `rune` | 16-bit Unicode character, big-endian |
| SHORT | 2 | `int16` | Signed 16-bit integer, big-endian |
| USHORT | 2 | `uint16` | Unsigned 16-bit integer, big-endian; read via `ReadUShort()` |
| INT | 4 | `int32` | Signed 32-bit integer, big-endian |
| UINT | 4 | `uint32` | Unsigned 32-bit integer, big-endian; read via `ReadUInt()` |

## Go `ClientConnection` interface

Defined in `yamul-gateway/internal/interfaces/client_connection.go`:

```go
ReadByte() byte
WriteByte(value byte)

ReadUShort() uint16
WriteUShort(value uint16)

ReadUInt() uint32
WriteUInt(value uint32)

ReadFixedString(length int) string
WriteFixedString(length int, value string)

ReadFixedBytes(length int) []byte
```
