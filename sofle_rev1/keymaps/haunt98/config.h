// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

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

// https://github.com/qmk/qmk_firmware/blob/master/docs/features/encoders.md
#ifdef ENCODER_RESOLUTION
#    undef ENCODER_RESOLUTION
#endif
#define ENCODER_RESOLUTION 4

// Experiment
#ifdef TAP_KEYCODE_DELAY
#    undef TAP_KEYCODE_DELAY
#endif
#define TAP_KEYCODE_DELAY 0

// https://github.com/qmk/qmk_firmware/issues/9962
#define NO_USB_STARTUP_CHECK

// https://github.com/qmk/qmk_firmware/issues/21686
#define SPLIT_WATCHDOG_ENABLE
