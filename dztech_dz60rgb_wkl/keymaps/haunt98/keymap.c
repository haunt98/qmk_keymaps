// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include QMK_KEYBOARD_H
#include "haunt98.h"

enum layers {
    LAYER_QWERTY,
    LAYER_NAV,
    LAYER_OPT,
    LAYER_ADJ,
};

// Inspire from https://github.com/qmk/qmk_firmware/blob/master/keyboards/dztech/dz60rgb_wkl/keymaps/hhkb/keymap.c
// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_grave_esc.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/mod_tap.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_space_cadet.md
// clang-format off
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    [LAYER_QWERTY] = LAYOUT_60_tsangan_hhkb(
        QK_GESC,        KC_1,     KC_2,     KC_3,    KC_4,     KC_5,           KC_6,    KC_7,  KC_8,     KC_9,    KC_0,     KC_MINS,  KC_EQL,         KC_BSLS,  KC_GRV,
        KC_TAB,         KC_Q,     KC_W,     KC_E,    KC_R,     KC_T,           KC_Y,    KC_U,  KC_I,     KC_O,    KC_P,     KC_LBRC,  KC_RBRC,        KC_BSPC,
        CTL_T(KC_ESC),  KC_A,     KC_S,     KC_D,    KC_F,     KC_G,           KC_H,    KC_J,  KC_K,     KC_L,    KC_SCLN,  KC_QUOT,  KC_ENT,
        SC_LSPO,        KC_Z,     KC_X,     KC_C,    KC_V,     KC_B,           KC_N,    KC_M,  KC_COMM,  KC_DOT,  KC_SLSH,  SC_RSPC,  MO(LAYER_NAV),
        XXXXXXX,        KC_LOPT,  KC_LCMD,  KC_SPC,  KC_RCMD,  MO(LAYER_OPT),  XXXXXXX
    ),
    [LAYER_NAV] = LAYOUT_60_tsangan_hhkb(
        KC_GRV,   KC_F1,    KC_F2,    KC_F3,    KC_F4,    KC_F5,    KC_F6,    KC_F7,    KC_F8,    KC_F9,    KC_F10,   KC_F11,   KC_F12,   XXXXXXX,  XXXXXXX,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_UP,    XXXXXXX,  KC_DEL,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_LEFT,  KC_RGHT,  XXXXXXX,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_DOWN,  _______,  XXXXXXX,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  _______,  XXXXXXX
    ),
    [LAYER_OPT] = LAYOUT_60_tsangan_hhkb(
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,  KC_BRMD,  KC_BRMU,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        KC_CAPS,  KC_VOLD,  KC_VOLU,  KC_MUTE,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        _______,  RGB_HUI,  RGB_SAI,  RGB_VAI,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  _______,  _______,
        XXXXXXX,  RGB_TOG,  RGB_MOD,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX
    ),
    [LAYER_ADJ] = LAYOUT_60_tsangan_hhkb(
        TD(TD_SAFE_BOOT),  EE_CLR,   XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,           XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,           XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,           XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,           XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX
    ),
};
// clang-format on


// https://github.com/qmk/qmk_firmware/blob/master/docs/ref_functions.md
layer_state_t layer_state_set_user(layer_state_t state) {
    return update_tri_layer_state(state, LAYER_NAV, LAYER_OPT, LAYER_ADJ);
}
