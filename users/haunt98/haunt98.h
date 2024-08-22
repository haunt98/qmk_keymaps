// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_userspace.md

#include QMK_KEYBOARD_H
#include "oled.h"

// https://github.com/qmk/qmk_firmware/blob/master/docs/mod_tap.md
#define CTL_ESC CTL_T(KC_ESC)

// Use only 4 layers
enum layers {
    LAYER_QWERTY,
    LAYER_LOWER,
    LAYER_UPPER,
    LAYER_ADJUST,
};
