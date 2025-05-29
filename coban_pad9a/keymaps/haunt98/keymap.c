// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include QMK_KEYBOARD_H

// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
// clang-format off
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    [0] = LAYOUT(
        KC_MUTE,  KC_MUTE,
        KC_ESC,   KC_F2,    KC_F12,
        KC_WAKE,  KC_LCMD,  KC_ENT
    ),
};
// clang-format on

// https://github.com/qmk/qmk_firmware/blob/master/docs/features/encoders.md
#ifdef ENCODER_ENABLE

#    if defined(ENCODER_MAP_ENABLE)
const uint16_t PROGMEM encoder_map[][NUM_ENCODERS][NUM_DIRECTIONS] = {
    [0] = {ENCODER_CCW_CW(KC_VOLD, KC_VOLU), ENCODER_CCW_CW(KC_PGUP, KC_PGDN)},
};
#    endif

#endif
