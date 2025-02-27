// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_userspace.md

#include "oled.h"

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_debounce_type.md
#ifdef DEBOUNCE
#    undef DEBOUNCE
#endif
#define DEBOUNCE 5

// https://github.com/qmk/qmk_firmware/blob/master/docs/tap_hold.md
#ifdef TAPPING_TERM
#    undef TAPPING_TERM
#endif
#define TAPPING_TERM 200

#ifdef QUICK_TAP_TERM
#    undef QUICK_TAP_TERM
#endif
#define QUICK_TAP_TERM 120

// https://github.com/qmk/qmk_firmware/blob/master/docs/mod_tap.md
#define CTL_ESC CTL_T(KC_ESC)

// Use only 4 layers
enum layers {
    LAYER_QWERTY,
    LAYER_LOWER,
    LAYER_UPPER,
    LAYER_ADJUST,
};
