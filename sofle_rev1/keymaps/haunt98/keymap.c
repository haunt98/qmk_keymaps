// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include QMK_KEYBOARD_H
#include "haunt98.h"

enum layers {
    LAYER_QWERTY,
    LAYER_LOWER,
    LAYER_UPPER,
    LAYER_ADJUST,
};

// Inspire from https://github.com/qmk/qmk_firmware/blob/master/keyboards/sofle/keymaps/default/keymap.c
// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/features/space_cadet.md
// clang-format off
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    [LAYER_QWERTY] = LAYOUT(
        KC_GRV,   KC_1,     KC_2,    KC_3,     KC_4,     KC_5,     KC_6,     KC_7,     KC_8,     KC_9,    KC_0,     KC_MINS,
        KC_TAB,   KC_Q,     KC_W,    KC_E,     KC_R,     KC_T,     KC_Y,     KC_U,     KC_I,     KC_O,    KC_P,     KC_BSPC,
        CTL_ESC,  KC_A,     KC_S,    KC_D,     KC_F,     KC_G,     KC_H,     KC_J,     KC_K,     KC_L,    KC_SCLN,  KC_ENT,
        SC_LSPO,  KC_Z,     KC_X,    KC_C,     KC_V,     KC_B,     KC_MUTE,  XXXXXXX,  KC_N,     KC_M,    KC_COMM,  KC_DOT,   KC_SLSH,  SC_RSPC,
        KC_LOPT,  KC_LCMD,  KC_SPC,  TL_LOWR,  KC_LBRC,  KC_RBRC,  TL_UPPR,  KC_SPC,   KC_RCMD,  KC_ROPT
    ),
    [LAYER_LOWER] = LAYOUT(
        _______,  KC_F1,    KC_F2,    KC_F3,    KC_F4,    KC_F5,    KC_F6,    KC_F7,    KC_F8,    KC_F9,    KC_F10,   KC_F11,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_LBRC,  KC_RBRC,  KC_7,     KC_8,     KC_9,     KC_0,     XXXXXXX,  KC_F12,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_4,     KC_5,     KC_6,     XXXXXXX,  XXXXXXX,  XXXXXXX,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  _______,  _______,  KC_1,     KC_2,     KC_3,     XXXXXXX,  XXXXXXX,  _______,
        _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______
    ),
    [LAYER_UPPER] = LAYOUT(
        _______,  KC_F1,    KC_F2,    KC_F3,    KC_F4,    KC_F5,    KC_F6,    KC_F7,    KC_F8,    KC_F9,    KC_F10,   KC_F11,
        _______,  _______,  _______,  KC_F7,    KC_F8,    KC_F9,    KC_LCBR,  KC_RCBR,  KC_EQL,   KC_PLUS,  XXXXXXX,  KC_F12,
        _______,  XXXXXXX,  XXXXXXX,  KC_F4,    KC_F5,    KC_F6,    KC_LEFT,  KC_DOWN,  KC_UP,    KC_RGHT,  KC_COLN,  _______,
        _______,  XXXXXXX,  XXXXXXX,  KC_F1,    KC_F2,    KC_F3,    _______,  _______,  KC_BSLS,  KC_PIPE,  KC_QUOT,  KC_DQUO,  XXXXXXX,  _______,
        _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______
    ),
    [LAYER_ADJUST] = LAYOUT(
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_WH_U,  KC_WH_D,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        KC_CAPS,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX
    )
};
// clang-format on

// https://github.com/qmk/qmk_firmware/blob/master/docs/features/oled_driver.md
#ifdef OLED_ENABLE

static void print_status_narrow(void) {
    oled_write_P(PSTR("\n\n"), false);

    // Print current layer
    // Each line has 5 char width
    oled_write_ln_P(PSTR("LAYER"), false);
    switch (get_highest_layer(layer_state)) {
        case LAYER_QWERTY:
            oled_write_ln_P(PSTR("BASE"), false);
            break;
        case LAYER_LOWER:
            oled_write_ln_P(PSTR("LOW"), false);
            break;
        case LAYER_UPPER:
            oled_write_ln_P(PSTR("UP"), false);
            break;
        case LAYER_ADJUST:
            oled_write_ln_P(PSTR("ADJ"), false);
            break;
        default:
            oled_write_ln_P(PSTR("???"), false);
    }

    oled_write_P(PSTR("\n\n"), false);

    // Print capslock status
    led_t led_usb_state = host_keyboard_led_state();
    oled_write_ln_P(PSTR("CAPS"), led_usb_state.caps_lock);
}

oled_rotation_t oled_init_user(oled_rotation_t rotation) {
    if (is_keyboard_master()) {
        return OLED_ROTATION_270;
    }
    return rotation;
}

bool oled_task_user(void) {
    if (is_keyboard_master()) {
        print_status_narrow();
    } else {
        render_logo_cyberpunk();
    }
    return false;
}

#endif

// https://github.com/qmk/qmk_firmware/blob/master/docs/features/encoders.md
#ifdef ENCODER_ENABLE

bool encoder_update_user(uint8_t index, bool clockwise) {
    if (index == 0) {
        switch (get_highest_layer(layer_state)) {
            case LAYER_UPPER:
                if (clockwise) {
                    tap_code(KC_BRMU);
                } else {
                    tap_code(KC_BRMD);
                }
                break;
            default:
                if (clockwise) {
                    tap_code(KC_VOLU);
                } else {
                    tap_code(KC_VOLD);
                }
        }
    } else if (index == 1) {
        switch (get_highest_layer(layer_state)) {
            case LAYER_LOWER:
                if (clockwise) {
                    tap_code(KC_BRMU);
                } else {
                    tap_code(KC_BRMD);
                }
                break;
            default:
                if (clockwise) {
                    tap_code(KC_PGDN);
                } else {
                    tap_code(KC_PGUP);
                }
        }
    }
    return false;
}

#endif
