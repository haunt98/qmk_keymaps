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
#ifdef TAPPING_TERM
#    undef TAPPING_TERM
#endif
#define TAPPING_TERM 200

#define PERMISSIVE_HOLD

// https://github.com/qmk/qmk_firmware/issues/9962
#define NO_USB_STARTUP_CHECK

// https://github.com/qmk/qmk_firmware/blob/master/docs/config_options.md
// https://github.com/qmk/qmk_firmware/issues/17316
#define USB_SUSPEND_WAKEUP_DELAY 200
