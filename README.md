# QMK for dztech/dz60rgb_wkl

## Install

Run `make` to build the firmware.

If success, firmware file is `dztech_dz60rgb_wkl_v2_1_haunt98.bin`

Currently firmware size:

```txt
* The firmware size is fine - 22364/26624 (83%, 4260 bytes free)
```

Put keyboard into bootloader mode:

- Press `QK_BOOT`.
- Unplug cable, then hold `ESC` and plug cable.

Use [QMK Toolbox](https://github.com/qmk/qmk_toolbox) to flash.

## Features

- [Keymap FAQ](https://github.com/qmk/qmk_firmware/blob/master/docs/faq_keymap.md)
  - [Keycodes Overview](https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md)
  - Do
    - [Grave Escape](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_grave_esc.md)
  - Don't
    - [Space Cadet: The Future, Built In](https://github.com/qmk/qmk_firmware/blob/master/docs/feature_space_cadet.md)
    - [Magic Keycodes](https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes_magic.md)
    - [One Shot Keys](https://github.com/qmk/qmk_firmware/blob/master/docs/one_shot_keys.md)
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
