# Septim bounded-decoder patch

1. Based on upstream v0.1.2, commit `410e34b936dc3c65e1409ef9ef9565c9fd4217e8`.
2. `TryDecodeWithBudget` limits visited cells and peeling steps across calls. Discard the decoder after an error. Existing `TryDecode` remains source compatible.
3. Invalid peeling degrees return a typed error in the bounded API.
4. Mapping indices always advance and saturate instead of overflowing. This prevents a zero floating-point increment from stalling generation.
5. Applications still bound input inventories, coded-symbol count, integer counts, concurrency and bytes. Decoding does not authenticate records or establish remote completeness. No Septim transport, record schema or application policy belongs here.
