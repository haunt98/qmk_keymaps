# https://github.com/qmk/qmk_firmware/blob/master/docs/config_options.md

# Optimize
# https://github.com/qmk/qmk_firmware/blob/master/docs/squeezing_avr.md
LTO_ENABLE = yes
CONSOLE_ENABLE = no
COMMAND_ENABLE = no
EXTRAKEY_ENABLE = yes

# https://github.com/qmk/qmk_firmware/blob/master/docs/features/mouse_keys.md
MOUSEKEY_ENABLE = no

# https://github.com/qmk/qmk_firmware/blob/master/docs/features/rgb_matrix.md
RGB_MATRIX_ENABLE = no

# https://github.com/qmk/qmk_firmware/blob/master/docs/features/rgblight.md
RGBLIGHT_ENABLE = no

# Feature
# https://github.com/qmk/qmk_firmware/blob/master/docs/features/bootmagic.md
MAGIC_ENABLE = no
BOOTMAGIC_ENABLE = yes

# https://github.com/qmk/qmk_firmware/blob/master/docs/feature_debounce_type.md
DEBOUNCE_TYPE = sym_defer_pr

# https://github.com/qmk/qmk_firmware/blob/master/docs/features/tri_layer.md
TRI_LAYER_ENABLE = yes

# https://github.com/qmk/qmk_firmware/blob/master/docs/features/tri_layer.md
SWAP_HANDS_ENABLE = no

# Custom C
SRC += haunt98.c
SRC += oled.c

# https://github.com/qmk/qmk_firmware/issues/21137
INTROSPECTION_KEYMAP_C = combos.c
