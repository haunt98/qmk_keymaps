// Copyright 2023 haunt98 <hauvipapro@gmail.com>
// SPDX-License-Identifier: MIT

#include "tap_dance.h"

#ifdef TAP_DANCE_ENABLE

void tap_dance_safe_boot(tap_dance_state_t* state, void* user_data) {
    if (state->count >= 3) {
        reset_keyboard();
    }
}

tap_dance_action_t tap_dance_actions[] = {
    [TD_SAFE_BOOT] = ACTION_TAP_DANCE_FN(tap_dance_safe_boot),
};

#endif
