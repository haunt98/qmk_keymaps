// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

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
