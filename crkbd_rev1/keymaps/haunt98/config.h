// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

// Experiment
#ifdef TAP_KEYCODE_DELAY
#    undef TAP_KEYCODE_DELAY
#endif
#define TAP_KEYCODE_DELAY 0

// https://github.com/qmk/qmk_firmware/issues/9962
#define NO_USB_STARTUP_CHECK

// https://github.com/qmk/qmk_firmware/issues/21686
#define SPLIT_WATCHDOG_ENABLE
