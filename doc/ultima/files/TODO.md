# File Format Documentation — TODO

Tracks documentation coverage for UO data file formats in `ultima/files/`.

## Documented

These formats have documentation confirmed against [SphereServer Source-X](https://github.com/Sphereserver/Source-X/tree/master/src/game/uo_files):

| File | Doc | Notes |
| --- | --- | --- |
| `tiledata.mul` | `tiledata_mul.md` | Pre-HS and HS variants documented |
| `statics*.mul` | `statics0_mul.md` | Field names corrected from SphereServer `CUOStaticItemRec` |
| `staidx*.mul` | `staidx0_mul.md` | |
| `multi.mul` | `multi_mul.md` | Pre-HS and HS variants documented |
| `multi.idx` | `multi_idx.md` | |
| `map*.mul` | `map0_mul.md` | Facet table added; header fields clarified |
| `hues.mul` | `hues_mul.md` | HueEntry corrected to 88 bytes (`WORD[34]` + `CHAR[20]`) |
| `verdata.mul` | `verdata_mul.md` | |
| `mobtypes.mul` | `mobtypes_mul.md` | New |
| `art.mul` / `artidx.mul` | `art_mul.md`, `art_idx.md` | |
| `anim.mul` / `animidx.mul` | `anim_mul.md`, `anim_idx.md` | |
| `gumpart.mul` / `gumpidx.mul` | `gumpart_mul.md`, `gumpidx_mul.md` | |
| `texmaps.mul` / `texidx.mul` | `texmaps_mul.md`, `texidx_mul.md` | |
| `sound.mul` / `soundidx.mul` | `sound_mul.md`, `soundidx_mul.md` | |
| `palette.mul` | `palette_mul.md` | |
| `radarcol.mul` | `radarcol_mul.md` | |
| `skills.mul` / `skills.idx` | `skills_mul.md`, `skills_idx.md` | |

## Not Yet Documented

These formats are present in UO clients but lack documentation here:

| File | Notes |
| --- | --- |
| `animdata.mul` | Animation timing and frame sequencing data |
| `bodyconv.def` | Body ID conversion table (maps bodies to animation file sets) |
| `cliloc*.mul` | Client localization strings (language-specific) |
| `fonts.mul` | Font glyph bitmap data |
| `light*.mul` | Light pattern shapes for dynamic lighting |
| `speech.mul` | NPC speech trigger keyword lists |
| `doors.mul` | Door animation and state data |
| `grdvirtue.mul` | Virtue/honor system data |
