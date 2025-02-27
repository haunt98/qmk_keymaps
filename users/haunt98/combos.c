// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include QMK_KEYBOARD_H
#include "haunt98.h"

#ifdef COMBO_ENABLE

// https://github.com/qmk/qmk_firmware/blob/master/docs/features/combo.md
const uint16_t PROGMEM combo_F12[] = {KC_F1, KC_F2, COMBO_END};

combo_t key_combos[] = {
    COMBO(combo_F12, KC_F12),
};

#endif
