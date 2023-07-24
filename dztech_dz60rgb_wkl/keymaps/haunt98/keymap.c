// Copyright 2023 haunt98

#include QMK_KEYBOARD_H

// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_grave_esc.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/mod_tap.md
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    [0] = LAYOUT_60_tsangan_hhkb(
        QK_GESC, KC_1, KC_2, KC_3, KC_4, KC_5, KC_6, KC_7, KC_8, KC_9, KC_0, KC_MINS, KC_EQL, KC_BSLS, KC_GRV,
        KC_TAB, KC_Q, KC_W, KC_E, KC_R, KC_T, KC_Y, KC_U, KC_I, KC_O, KC_P, KC_LBRC, KC_RBRC, KC_BSPC,
        CTL_T(KC_ESC), KC_A, KC_S, KC_D, KC_F, KC_G, KC_H, KC_J, KC_K, KC_L, KC_SCLN, KC_QUOT, KC_ENT,
        KC_LSFT, KC_Z, KC_X, KC_C, KC_V, KC_B, KC_N, KC_M, KC_COMM, KC_DOT, KC_SLSH, KC_RSFT, MO(1),
        XXXXXXX, KC_LALT, KC_LGUI, KC_SPC, KC_RGUI, KC_RALT, XXXXXXX),
    [1] = LAYOUT_60_tsangan_hhkb(
        KC_GRV, KC_F1, KC_F2, KC_F3, KC_F4, KC_F5, KC_F6, KC_F7, KC_F8, KC_F9, KC_F10, KC_F11, KC_F12, _______, TG(2),
        _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, KC_UP, _______, _______,
        _______, KC_VOLD, KC_VOLU, KC_MUTE, _______, _______, _______, _______, _______, _______, KC_LEFT, KC_RGHT, _______,
        _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, KC_DOWN, _______, _______,
        XXXXXXX, _______, _______, _______, _______, _______, XXXXXXX),
    [2] = LAYOUT_60_tsangan_hhkb(
        QK_BOOT, EE_CLR, _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, TO(0),
        _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, _______, _______,
        KC_CAPS, RGB_HUI, RGB_SAI, RGB_VAI, _______, _______, _______, _______, KC_HOME, KC_PGUP, _______, _______, _______,
        _______, _______, _______, _______, _______, _______, _______, _______, KC_END, KC_PGDN, _______, _______, _______,
        XXXXXXX, _______, _______, _______, RGB_TOG, RGB_MOD, XXXXXXX)};
