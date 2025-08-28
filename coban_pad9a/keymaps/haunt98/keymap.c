// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include QMK_KEYBOARD_H
#include "haunt98.h"

// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
// clang-format off
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    [LAYER_QWERTY] = LAYOUT(
        XXXXXXX,  XXXXXXX,
        KC_ESC,   KC_UP,    KC_F12,
        TL_LOWR,  KC_DOWN,  KC_ENT
    ),
    [LAYER_LOWER] = LAYOUT(
        XXXXXXX,  XXXXXXX,
        _______,  KC_PGUP,  _______,
        _______,  KC_PGDN,  _______
    ),
};
// clang-format on

// https://github.com/qmk/qmk_firmware/blob/master/docs/features/encoders.md
#ifdef ENCODER_ENABLE

#    if defined(ENCODER_MAP_ENABLE)
const uint16_t PROGMEM encoder_map[][NUM_ENCODERS][NUM_DIRECTIONS] = {
    [LAYER_QWERTY] = {ENCODER_CCW_CW(KC_VOLD, KC_VOLU), ENCODER_CCW_CW(KC_PGUP, KC_PGDN)},
    [LAYER_LOWER]  = {ENCODER_CCW_CW(KC_VOLD, KC_VOLU), ENCODER_CCW_CW(KC_BRMD, KC_BRMU)},
};
#    endif

#endif
