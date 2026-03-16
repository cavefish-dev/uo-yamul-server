# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

This `doc/` folder contains reference documentation for two aspects of the Ultima Online data formats used by this project:

- **`mulHexPat/`** — [ImHex](https://imhex.werwolv.net/) pattern files for inspecting UO binary MUL/UOP data files
- **`ultima/transport/`** — Markdown documentation for UO client↔server network packet formats

## Transport Packet Docs (`ultima/transport/`)

Each file documents one UO protocol packet. Use `template.md` as the template when adding new packet docs.

Reference sites for packet formats include:
- [WolfPAck UO Protocol Documentation](https://www.hoogi.de/wolfpack/wiki/doku.php?id=uo_protocol)

### Template format

```markdown
# (0xNN) PACKET_NAME

Description of the packet and when it is used.

## Message format

Offset | Type | Name | Description
--- | --- | --- | ---
0x00 | BYTE | cmd | Command code, always 0xNN
...
```

File naming convention: `0xNN_description.md` (e.g., `0x01_disconnect.md`).

### Type names

Packet docs use the **UO wire type names** from the WolfPAck reference site (BYTE, USHORT, UINT, etc.). See [`types.md`](ultima/transport/types.md) for the mapping to Go types and the `ClientConnection` read/write methods.

## MUL Hex Patterns (`mulHexPat/`)

These are ImHex `.hexpat` pattern files for visually inspecting UO binary files:

- **`tiledata.hexpat`** — Parses `tiledata.mul`: 512 land tile groups followed by static tile groups. Each group has 32 entries preceded by a 4-byte unknown header. Pattern limit is set to `0x30A800`.
- **`uop.hexpat`** — Parses UOP container files (newer UO format): `Header → Table → TableEntry[]`, where each entry contains offset, size, decompressed size, filename hash, and compression flag, with `MapData` blocks at `offset + headerLength`.

These patterns correspond to the MUL parsing code in `yamul-backend/src/main/kotlin/.../game/controller/domain/mul/`.
