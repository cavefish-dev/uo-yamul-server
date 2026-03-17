# DIF Files (stadifl, stadifi, stadif, mapdifl, mapdif)

Patch files used by the 3D (LBR+) client to override terrain and static data without replacing the base MUL files. Each set of DIF files corresponds to a facet number (`#`).

Reference: [Heptazane — Section 3.23](https://uo.stratics.com/heptazane/fileformats.shtml#3.23)

## File set

| File | Role |
| --- | --- |
| `stadifl#.mul` | List of affected static block numbers |
| `stadifi#.mul` | Index into `stadif#.mul` (mirrors `staidx#.mul` format) |
| `stadif#.mul` | Replacement static data (mirrors `statics#.mul` format) |
| `mapdifl#.mul` | List of affected map block numbers |
| `mapdif#.mul` | Replacement map data (mirrors `map#.mul` block format) |

## File structures

### stadifl#.mul / mapdifl#.mul — affected block list

Sequential array of `DWORD` block numbers. Each entry is one block that has been patched.

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`N * 4` | `4` | `DWORD` | `BlockNum` | Map/static block number that is overridden

### stadifi#.mul — static patch index

Same 12-byte record format as `staidx#.mul`. Entry N corresponds to the block number at `stadifl#.mul[N]`.

Offset | Size | Type | Name | Description
--- | --- | --- | --- | ---
`0x00` | `4` | `DWORD` | `Start` | Byte position in `stadif#.mul`
`0x04` | `4` | `DWORD` | `Length` | Size of statics data for this block
`0x08` | `4` | `DWORD` | `Unknown` | Purpose unknown

### stadif#.mul — static patch data

Same 7-byte record format as `statics#.mul`. See [`statics0_mul.md`](statics0_mul.md).

### mapdif#.mul — map patch data

Same 196-byte block format as `map#.mul`. Entry N corresponds to the block number at `mapdifl#.mul[N]`. See [`map0_mul.md`](map0_mul.md).
