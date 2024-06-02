# https://github.com/qmk/qmk_firmware/blob/master/docs/config_options.md

# Optimize
# https://github.com/qmk/qmk_firmware/blob/master/docs/squeezing_avr.md
LTO_ENABLE = yes

CONSOLE_ENABLE = no
COMMAND_ENABLE = no
EXTRAKEY_ENABLE = yes

# Feature
# https://github.com/qmk/qmk_firmware/blob/master/docs/features/bootmagic.md
MAGIC_ENABLE = no
BOOTMAGIC_ENABLE = yes

# https://github.com/qmk/qmk_firmware/blob/master/docs/features/tap_dance.md
TAP_DANCE_ENABLE = yes

# https://github.com/qmk/qmk_firmware/blob/master/docs/features/tri_layer.md
TRI_LAYER_ENABLE = yes

# https://github.com/qmk/qmk_firmware/blob/master/docs/features/tri_layer.md
SWAP_HANDS_ENABLE = no

# Custom C
SRC += haunt98.c
SRC += tap_dance.c
SRC += oled.c
