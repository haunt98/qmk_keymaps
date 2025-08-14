// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include QMK_KEYBOARD_H
#include "haunt98.h"

// Inspire from https://github.com/qmk/qmk_firmware/blob/master/keyboards/crkbd/keymaps/default/keymap.c
// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/features/space_cadet.md
// clang-format off
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    [LAYER_QWERTY] = LAYOUT_split_3x6_3(
        KC_TAB,   KC_Q,   KC_W,    KC_E,    KC_R,   KC_T,    KC_Y,  KC_U,  KC_I,     KC_O,    KC_P,     KC_BSPC,
        KC_LCTL,  KC_A,   KC_S,    KC_D,    KC_F,   KC_G,    KC_H,  KC_J,  KC_K,     KC_L,    KC_SCLN,  KC_QUOT,
        KC_LSFT,  KC_Z,   KC_X,    KC_C,    KC_V,   KC_B,    KC_N,  KC_M,  KC_COMM,  KC_DOT,  KC_SLSH,  KC_ESC,
        KC_LGUI,  MO(1),  KC_SPC,  KC_ENT,  MO(2),  KC_RALT
    ),
    [LAYER_LOWER] = LAYOUT_split_3x6_3(
        KC_TAB,   KC_1,     KC_2,     KC_3,     KC_4,     KC_5,     KC_6,     KC_7,     KC_8,            KC_9,     KC_0,     KC_BSPC,
        KC_LCTL,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_LEFT,  KC_DOWN,  KC_UP,KC_RIGHT,  XXXXXXX,  XXXXXXX,
        KC_LSFT,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,         XXXXXXX,  XXXXXXX,  XXXXXXX,
        KC_LGUI,  _______,  KC_SPC,   KC_ENT,   MO(3),    KC_RALT
    ),
    [LAYER_UPPER] = LAYOUT_split_3x6_3(
        KC_TAB,   KC_EXLM,  KC_AT,    KC_HASH,  KC_DLR,   KC_PERC,  KC_CIRC,  KC_AMPR,  KC_ASTR,  KC_LPRN,  KC_RPRN,  KC_BSPC,
        KC_LCTL,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_MINS,  KC_EQL,   KC_LBRC,  KC_RBRC,  KC_BSLS,  KC_GRV,
        KC_LSFT,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_UNDS,  KC_PLUS,  KC_LCBR,  KC_RCBR,  KC_PIPE,  KC_TILD,
        KC_LGUI,  MO(3),    KC_SPC,   KC_ENT,   _______,  KC_RALT
    ),
    [LAYER_ADJUST] = LAYOUT_split_3x6_3(
        QK_BOOT,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        RM_TOGG,  RM_HUEU,  RM_SATU,  RM_VALU,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        RM_NEXT,  RM_HUED,  RM_SATD,  RM_VALD,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        KC_LGUI,  _______,  KC_SPC,   KC_ENT,   _______,  KC_RALT
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
