// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#pragma once

#include QMK_KEYBOARD_H

// https://github.com/qmk/qmk_firmware/blob/master/docs/feature_oled_driver.md
#ifdef OLED_ENABLE

void render_logo_qmk(void);
void render_logo_cyberpunk(void);

#endif
