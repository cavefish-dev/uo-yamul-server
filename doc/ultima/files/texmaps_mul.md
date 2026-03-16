# TEXMAPS.MUL

Ground texture data used for isometric terrain rendering. Accessed via `TEXIDX.MUL`.

Reference: [Heptazane — Section 3.21](https://uo.stratics.com/heptazane/fileformats.shtml#3.21)

## File structure

Each texture entry starts at the byte position given by `TEXIDX.MUL`. The entry is raw 16-bit pixel data; the image dimensions are determined by `Length`:

| `Length` | Dimensions |
| --- | --- |
| `0x2000` (8,192 bytes) | 64 × 64 pixels |
| `0x8000` (32,768 bytes) | 128 × 128 pixels |

Each pixel is a `UWORD` in the standard UO 16-bit color format (see [`types.md`](types.md)). Pixels are stored row-major, left-to-right, top-to-bottom.
