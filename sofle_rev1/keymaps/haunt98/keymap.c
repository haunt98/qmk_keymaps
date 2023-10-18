// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include QMK_KEYBOARD_H
#include "haunt98.h"

enum layers {
    LAYER_QWERTY,
    LAYER_SYM,
    LAYER_NAV,
    LAYER_ADJ,
};

#define LT_SYM_SPC LT(LAYER_SYM, KC_SPC)
#define MO_NAV MO(LAYER_NAV)

// Inspire from https://github.com/qmk/qmk_firmware/blob/master/keyboards/sofle/keymaps/default/keymap.c
// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_grave_esc.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/mod_tap.md
// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_space_cadet.md
// clang-format off
const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {
    [LAYER_QWERTY] = LAYOUT(
        QK_GESC,  KC_1,      KC_2,     KC_3,        KC_4,    KC_5,    KC_6,     KC_7,     KC_8,      KC_9,    KC_0,     KC_GRV,
        KC_TAB,   KC_Q,      KC_W,     KC_E,        KC_R,    KC_T,    KC_Y,     KC_U,     KC_I,      KC_O,    KC_P,     KC_BSPC,
        CTL_ESC,  KC_A,      KC_S,     KC_D,        KC_F,    KC_G,    KC_H,     KC_J,     KC_K,      KC_L,    KC_SCLN,  KC_QUOT,
        SC_LSPO,  KC_Z,      KC_X,     KC_C,        KC_V,    KC_B,    KC_MUTE,  XXXXXXX,  KC_N,      KC_M,    KC_COMM,  KC_DOT,   KC_SLSH,  SC_RSPC,
        KC_LCTL,  LOPT_BRC,  KC_LCMD,  LT_SYM_SPC,  KC_SPC,  KC_ENT,  MO_NAV,   KC_ENT,   ROPT_BRC,  KC_COLN
    ),
    [LAYER_SYM] = LAYOUT(
        KC_GRV,   KC_F1,    KC_F2,    KC_F3,    KC_F4,    KC_F5,    KC_F6,    KC_F7,    KC_F8,    KC_F9,    KC_F10,   KC_F11,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_MINS,  KC_EQL,   KC_F12,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_LBRC,  KC_LCBR,  KC_RCBR,  KC_RBRC,  KC_BSLS,  KC_PIPE,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  _______,  _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  _______,
        _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______
    ),
    [LAYER_NAV] = LAYOUT(
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  KC_LEFT,  KC_DOWN,  KC_UP,    KC_RGHT,  XXXXXXX,  XXXXXXX,
        _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  _______,  _______,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  _______,
        _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______,  _______
    ),
    [LAYER_ADJ] = LAYOUT(
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        KC_CAPS,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,
        XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX,  XXXXXXX
    )
};
// clang-format on

// https://github.com/qmk/qmk_firmware/blob/master/docs/ref_functions.md
layer_state_t layer_state_set_user(layer_state_t state) {
    return update_tri_layer_state(state, LAYER_SYM, LAYER_NAV, LAYER_ADJ);
}

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_oled_driver.md
#ifdef OLED_ENABLE

static void render_logo(void) {
    // clang-format off
    static const char PROGMEM qmk_logo[] = {
        0x80,  0x81,  0x82,  0x83,  0x84,  0x85,  0x86,  0x87,  0x88,  0x89,  0x8A,  0x8B,  0x8C,  0x8D,  0x8E,  0x8F,  0x90,  0x91,  0x92,  0x93,  0x94,
        0xA0,  0xA1,  0xA2,  0xA3,  0xA4,  0xA5,  0xA6,  0xA7,  0xA8,  0xA9,  0xAA,  0xAB,  0xAC,  0xAD,  0xAE,  0xAF,  0xB0,  0xB1,  0xB2,  0xB3,  0xB4,
        0xC0,  0xC1,  0xC2,  0xC3,  0xC4,  0xC5,  0xC6,  0xC7,  0xC8,  0xC9,  0xCA,  0xCB,  0xCC,  0xCD,  0xCE,  0xCF,  0xD0,  0xD1,  0xD2,  0xD3,  0xD4,  0x00
    };
    // clang-format on

    oled_write_P(qmk_logo, false);
}

static void print_status_narrow(void) {
    oled_write_P(PSTR("\n\n"), false);

    // Print current layer
    oled_write_ln_P(PSTR("LAYER"), false);
    switch (get_highest_layer(layer_state)) {
        case LAYER_QWERTY:
            oled_write_ln_P(PSTR("BASE"), false);
            break;
        case LAYER_SYM:
            oled_write_ln_P(PSTR("SYM"), false);
            break;
        case LAYER_NAV:
            oled_write_ln_P(PSTR("NAV"), false);
            break;
        case LAYER_ADJ:
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
        render_logo();
    }
    return false;
}

#endif

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_encoders.md
#ifdef ENCODER_ENABLE

bool encoder_update_user(uint8_t index, bool clockwise) {
    if (index == 0) {
        if (clockwise) {
            tap_code(KC_VOLU);
        } else {
            tap_code(KC_VOLD);
        }
    } else if (index == 1) {
        if (clockwise) {
            tap_code(KC_PGDN);
        } else {
            tap_code(KC_PGUP);
        }
    }
    return true;
}

#endif
