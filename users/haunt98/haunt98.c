// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include QMK_KEYBOARD_H
#include "haunt98.h"

bool get_permissive_hold(uint16_t keycode, keyrecord_t *record) {
    switch (keycode) {
        case OPT_KCS:
            return true;
        default:
            return false;
    }
}

bool get_hold_on_other_key_press(uint16_t keycode, keyrecord_t *record) {
    switch (keycode) {
        case CTL_ESC:
            return true;
        case SC_LSPO:
            return true;
        case SC_RSPC:
            return true;
        default:
            return false;
    }
}

bool get_chordal_hold(uint16_t tap_hold_keycode, keyrecord_t *tap_hold_record, uint16_t other_keycode, keyrecord_t *other_record) {
    switch (tap_hold_keycode) {
        case CTL_ESC:
            return true;
        case SC_LSPO:
            return true;
        case SC_RSPC:
            return true;
    }

    return get_chordal_hold_default(tap_hold_record, other_record);
}
