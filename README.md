# QMK for dztech/dz60rgb_wkl

From [QMK dztech/dz60rgb_wkl](https://github.com/qmk/qmk_firmware/tree/master/keyboards/dztech/dz60rgb_wkl):

> V2.1: ATmega32U4 (AVR), takes .bin files Started shipping in June 2021 as a slightly modified version of v2.

## Install

Run `make` to build the firmware.
If success, firmware file `dztech_dz60rgb_wkl_v2_1_haunt98.bin` is located inside repo.

Currently firmware size:

```txt
* The firmware size is fine - 16074/26624 (60%, 10550 bytes free)
```

Put keyboard into bootloader mode in 2 ways (only choose one):

- Press `QK_BOOT`.
- Unplug cable, then hold `ESC` and plug cable.

Use [QMK Toolbox](https://github.com/qmk/qmk_toolbox) to flash.

## Features

Basic

- [Keymap FAQ](https://github.com/qmk/qmk_firmware/blob/master/docs/faq_keymap.md)
- [Keycodes Overview](https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md)
- [Modifier Keys](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_advanced_keycodes.md)
- [Quantum Keycodes](https://github.com/qmk/qmk_firmware/blob/master/docs/quantum_keycodes.md)
- [Layers](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_layers.md)

Do

- [Grave Escape](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_grave_esc.md)
- [Contact bounce / contact chatter](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_debounce_type.md)
- [Bootmagic Lite](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_bootmagic.md)

Don't

- [Magic Keycodes](https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes_magic.md)
- [Command](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_command.md)
- [Mod-Tap](https://github.com/qmk/qmk_firmware/blob/master/docs/mod_tap.md)
- [Mouse keys](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_mouse_keys.md)
- [Space Cadet: The Future, Built In](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_space_cadet.md)
- [One Shot Keys](https://github.com/qmk/qmk_firmware/blob/master/docs/one_shot_keys.md)
- [RGB Matrix Lighting](https://docs.qmk.fm/#/feature_rgb_matrix)

Build/Optimize

- [Configuring QMK](https://github.com/qmk/qmk_firmware/blob/master/docs/config_options.md)
- [Squeezing the most out of AVR](https://github.com/qmk/qmk_firmware/blob/master/docs/squeezing_avr.md)
- [Reducing firmware size](https://get.vial.today/docs/firmware-size.html)
- [Configuring QMK](https://www.caniusevia.com/docs/configuring_qmk)

## Stock

Stock firmware is `stock/dztech_dz60rgb_wkl_v2_1.bin`.
Can be flashed with QMK Toolbox.

Links:

- [DZ60RGB-WKL Hot-Swap PCB](https://kbdfans.com/products/dz60rgb-wkl-hot-swap-pcb)
- [USB Drive Flash Manual (New Flash Manual)](https://docs.google.com/document/d/111qx6Qec4JqtIhWaZlMND-VuRnFtn9a-gJaHN8fsL7M/edit?usp=sharing)

## TODO

- [x] How to flash
- [x] Add stock firmware/docs
