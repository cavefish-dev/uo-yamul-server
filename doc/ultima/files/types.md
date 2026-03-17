# UO File Format Data Types

Primitive types used across all UO binary file format documentation.

Reference: [Heptazane — Section 1.1](https://uo.stratics.com/heptazane/fileformats.shtml#1.1)

## Primitive types

| Type | Size | Signed | Range |
| --- | --- | --- | --- |
| `CHAR` | 1 byte | — | Character value |
| `BYTE` | 1 byte | yes | -128 to 127 |
| `UBYTE` | 1 byte | no | 0 to 255 |
| `WORD` | 2 bytes | yes | -32,768 to 32,767 |
| `UWORD` | 2 bytes | no | 0 to 65,535 |
| `DWORD` | 4 bytes | yes | -2,147,483,648 to 2,147,483,647 |
| `UDWORD` | 4 bytes | no | 0 to 4,294,967,295 |

All multi-byte values are **little-endian**.

## Color format

Reference: [Heptazane — Section 1.2](https://uo.stratics.com/heptazane/fileformats.shtml#1.2)

All colors are 16-bit `UWORD` values packed as:

```
Bit: 15  14 13 12 11 10  9  8  7  6  5  4  3  2  1  0
      U   R  R  R  R  R  G  G  G  G  G  B  B  B  B  B
```

- Bit 15: unused
- Bits 14–10: Red (5 bits)
- Bits 9–5: Green (5 bits)
- Bits 4–0: Blue (5 bits)

**32-bit conversion:**

```
R32 = ((Color16 >> 10) & 0x1F) * 0xFF / 0x1F
G32 = ((Color16 >>  5) & 0x1F) * 0xFF / 0x1F
B32 = ( Color16        & 0x1F) * 0xFF / 0x1F
Color32 = R32 | (G32 << 8) | (B32 << 16)
```
