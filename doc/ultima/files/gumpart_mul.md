# GUMPART.MUL

UI gump and paperdoll graphics stored as run-length encoded rows. Accessed via `GUMPIDX.MUL`.

Reference: [Heptazane — Section 3.5](https://uo.stratics.com/heptazane/fileformats.shtml#3.5)

## File structure

Each gump entry starts at the offset in `GUMPIDX.MUL`. Width, Height, and Size come from the index record.

### Gump entry

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4*H` | `DWORD[H]` | `RowLookup` | Per-row block offset (in DWORD units from start of RowLookup)
`0x04*H` | var | — | `RowData` | Run-length encoded pixel blocks

### Row data

Each row is a sequence of `(Value, Run)` pairs:

| Field | Size | Type | Description |
| --- | --- | --- | --- |
| `Value` | 2 | `WORD` | 16-bit pixel color |
| `Run` | 2 | `WORD` | Number of consecutive pixels with this color |

**Block count for row Y:**
- If `Y < Height - 1`: `BlockCount = RowLookup[Y+1] - RowLookup[Y]`
- If `Y == Height - 1`: `BlockCount = (Size / 4) - RowLookup[Y]`
