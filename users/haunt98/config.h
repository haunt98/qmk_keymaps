// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

// https://github.com/qmk/qmk_firmware/blob/master/docs/config_options.md

// Optimize
// https://github.com/qmk/qmk_firmware/blob/master/docs/squeezing_avr.md
#define NO_ACTION_ONESHOT

#define LAYER_STATE_8BIT

// Feature
// https://github.com/qmk/qmk_firmware/blob/master/docs/tap_hold.md
#undef TAPPING_TERM
#define TAPPING_TERM 200
#undef QUICK_TAP_TERM
#define QUICK_TAP_TERM 120
#define HOLD_ON_OTHER_KEY_PRESS
#define RETRO_TAPPING

// https://github.com/qmk/qmk_firmware/blob/master/docs/config_options.md
// https://github.com/qmk/qmk_firmware/issues/17316
#define USB_SUSPEND_WAKEUP_DELAY 200

#define FORCE_NKRO
