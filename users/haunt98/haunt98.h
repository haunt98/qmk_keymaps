// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_userspace.md

#include QMK_KEYBOARD_H
#include "tap_dance.h"
#include "oled.h"

// https://github.com/qmk/qmk_firmware/blob/master/docs/mod_tap.md
#define CTL_ESC CTL_T(KC_ESC)
#define LOPT_BRC LOPT_T(KC_LBRC)
#define ROPT_BRC ROPT_T(KC_RBRC)
