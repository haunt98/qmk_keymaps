# QMK custom keymaps

## [dztech/dz60rgb_wkl](https://github.com/qmk/qmk_firmware/tree/master/keyboards/dztech/dz60rgb_wkl)

> V2.1: ATmega32U4 (AVR), takes .bin files Started shipping in June 2021 as a slightly modified version of v2.

### Stock

Stock firmware is `dztech_dz60rgb_wkl/stock/new dztech_dz60rgb_wkl_v3_via.bin`.
Can be flashed with QMK Toolbox.

Links:

- [DZ60RGB-WKL Hot-Swap PCB](https://kbdfans.com/products/dz60rgb-wkl-hot-swap-pcb)
- [USB Drive Flash Manual (New Flash Manual)](https://docs.google.com/document/d/111qx6Qec4JqtIhWaZlMND-VuRnFtn9a-gJaHN8fsL7M/edit?usp=sharing)

### Compile

```sh
make dztech_dz60rgb_wkl
# Output is dztech_dz60rgb_wkl_v2_1_haunt98.bin
```

Currently firmware size:

```txt
* The firmware size is fine - 16080/26624 (60%, 10544 bytes free)
```

## Install

Locate firmware file after compile successfully.

Put keyboard into bootloader mode in 2 ways:

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

Blog

- [Measure and reduce keyboard input latency with QMK on the Kinesis Advantage (2021)](https://michael.stapelberg.ch/posts/2021-05-08-keyboard-input-latency-qmk-kinesis/)

## TODO

- [x] How to flash
- [x] Add stock firmware/docs
- [x] Add qmk lint
