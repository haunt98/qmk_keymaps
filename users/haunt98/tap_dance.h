// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

#include QMK_KEYBOARD_H

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_tap_dance.md
#ifdef TAP_DANCE_ENABLE

enum custom_tap_dance {
    TD_SAFE_BOOT,
};

#endif
