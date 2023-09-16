# https://github.com/qmk/qmk_firmware/blob/master/docs/config_options.md

# Optimize
# https://github.com/qmk/qmk_firmware/blob/master/docs/squeezing_avr.md
LTO_ENABLE = yes

CONSOLE_ENABLE = no
COMMAND_ENABLE = no
MOUSEKEY_ENABLE = no
EXTRAKEY_ENABLE = yes

# Feature
# https://github.com/qmk/qmk_firmware/blob/master/docs/feature_bootmagic.md
MAGIC_ENABLE = no
BOOTMAGIC_ENABLE = yes

# https://github.com/qmk/qmk_firmware/blob/master/docs/feature_tap_dance.md
TAP_DANCE_ENABLE = yes

# Feature
# https://github.com/qmk/qmk_firmware/blob/master/docs/feature_debounce_type.md
DEBOUNCE_TYPE = sym_eager_pk

NKRO_ENABLE = yes

# Custom C
SRC += haunt98.c
SRC += tap_dance.c
