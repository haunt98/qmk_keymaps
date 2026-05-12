// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_userspace.md

#include "oled.h"

// https://github.com/qmk/qmk_firmware/blob/master/docs/mod_tap.md
#define CTL_ESC CTL_T(KC_ESC)
#define OPT_KCS OPT_T(KC_S)

// Use only 4 layers
enum layers {
    LAYER_QWERTY,
    LAYER_LOWER,
    LAYER_UPPER,
    LAYER_ADJUST,
};
