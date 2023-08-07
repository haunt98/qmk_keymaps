# https://github.com/qmk/qmk_firmware/blob/master/docs/config_options.md

# VIA
# https://www.caniusevia.com/docs/configuring_qmk
VIA_ENABLE = no

# Optimize
# https://github.com/qmk/qmk_firmware/blob/master/docs/squeezing_avr.md
LTO_ENABLE = yes

CONSOLE_ENABLE = no
COMMAND_ENABLE = no
MOUSEKEY_ENABLE = no
EXTRAKEY_ENABLE = yes

SPACE_CADET_ENABLE = no
MAGIC_ENABLE = no

# Feature
# https://github.com/qmk/qmk_firmware/blob/master/docs/feature_bootmagic.md
BOOTMAGIC_ENABLE = yes

# https://github.com/qmk/qmk_firmware/blob/master/docs/feature_debounce_type.md
DEBOUNCE_TYPE = sym_eager_pr
