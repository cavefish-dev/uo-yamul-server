# yamul-gateway TODO

## 1. Configuration (currently hardcoded)

- TCP listen port (`2593` in `main.go`)
- gRPC backend endpoints (`localhost:8087`, `localhost:8088`, `localhost:8089`)
- Log level and format
- Starting city/tavern/coordinates (hardcoded in character creation response)
- Character slot count

## 2. In-code TODOs

- `x06_double_click.go` — target serial is read but not forwarded to backend
- `xc8_view_range.go` — view range value is read but not used
- `xbf19_extended_stats.go` — case 5 is unknown; currently kills the connection
- `xbf18_world_patches.go` — returns placeholder data, not real world patch info
- `xb5_open_chat_window.go` — handler is incomplete
- `on_open_chat_window.go` — stub body, no implementation
- `game_service.go:66` — server-initiated stream close is not handled

## 3. Unimplemented event processors (`events/main.go`)

- `TypeCharacterSelection`
- `TypeClientDoubleClick`
- `TypeClientMoveRequest`
- `TypeSkillUpdateClient`
- `TypeUnicodeSpeechSelected`

## 4. Stub packet handlers (read bytes, then kill connection or no-op)

- `0x09` — single click: reads 4 bytes, kills connection
- `0x34` — get player status: reads 9 bytes, kills connection
- `0x9B` — help request: reads 257 bytes, sends no response

## 5. Missing UO protocol packets

Outbound (server→client) packets not yet implemented:

| Packet | Description |
|--------|-------------|
| `0x01` | Disconnect/logout |
| `0x03` | Talk/speech |
| `0x05` | Attack OK |
| `0x07` | Enter world |
| `0x08` | Drop item failed |
| `0x12` | Trigger on/off |
| `0x13` | Equip item |
| `0x24` | Draw game player |
| `0x25` | Draw item in container |
| `0x2C` | Death |
| `0x2E` | Equip item |
| `0x2F` | Swing |
| `0x3C` | Container contents |
| `0x4E` | Personal light level |
| `0x53` | War mode |
| `0x6C` | Target cursor |
| `0x6E` | Character animation |
| `0x74` | Open buy window |
| `0x77` | Update character |
| `0x7C` | Open menu |
| `0x83` | Delete character |
| `0x99` | Move item |
| `0xAE` | Unicode speech |
| `0xB0` | Generic gump |
| `0xB8` | Character profile |
| `0xBE` | Quest arrow |
| `0xC0` | Sound effect |
| `0xCC` | Cliloc message affix |
| `0xD6` | Mega cliloc |
| `0xD7` | Custom house info |
| `0xDD` | Compressed gump |
| `0xF3` | Draw object |

## 6. Test coverage gaps

- 28+ packet handlers have no unit tests
- No integration test covering a full login → character select → enter game flow
