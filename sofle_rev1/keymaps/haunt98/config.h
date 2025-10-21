// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

// https://github.com/qmk/qmk_firmware/blob/master/docs/features/encoders.md
#ifdef ENCODER_RESOLUTION
#    undef ENCODER_RESOLUTION
#endif
#define ENCODER_RESOLUTION 4

// https://github.com/qmk/qmk_firmware/issues/21686
#define SPLIT_WATCHDOG_ENABLE
