.PHONY: all format clean draw dztech_dz60rgb_wkl

all:
	$(MAKE) format
	$(MAKE) dztech_dz60rgb_wkl

format:
	clang-format -i dztech_dz60rgb_wkl/keymaps/haunt98/*.c dztech_dz60rgb_wkl/keymaps/haunt98/*.h

clean:
	qmk clean
	rm -rf ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/haunt98
	rm -rf dztech_dz60rgb_wkl_v2_1_haunt98.bin

draw:
	pipx install keymap-drawer
	# keymap -c keymap-drawer-custom.yaml parse -q dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json > dztech_dz60rgb_wkl/keymap-drawer/haunt98/keymap.yaml
	keymap -c keymap-drawer-custom.yaml draw dztech_dz60rgb_wkl/keymap-drawer/haunt98/keymap.yaml > dztech_dz60rgb_wkl/keymap-drawer/haunt98/keymap.svg

dztech_dz60rgb_wkl:
	# Copy
	rm -rf ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/haunt98
	cp -rf dztech_dz60rgb_wkl/keymaps/haunt98 ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/
	# Lint
	qmk lint -kb dztech/dz60rgb_wkl/v2_1 -km haunt98
	# Compile
	qmk compile -j 8 -kb dztech/dz60rgb_wkl/v2_1 -km haunt98
	mv ~/qmk_firmware/dztech_dz60rgb_wkl_v2_1_haunt98.bin .
	# JSON
	rm -rf dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json
	qmk c2json -kb dztech/dz60rgb_wkl/v2_1 -km haunt98 -o dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json dztech_dz60rgb_wkl/keymaps/haunt98/keymap.c
