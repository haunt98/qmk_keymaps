.PHONY: all format cp compile json clean

all:
	$(MAKE) format
	$(MAKE) compile

format:
	clang-format -i keymaps/haunt98/*.c keymaps/haunt98/*.h

cp:
	rm -rf ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/haunt98
	cp -rf keymaps/haunt98 ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/

compile:
	$(MAKE) cp
	qmk compile -kb dztech/dz60rgb_wkl/v2_1 -km haunt98
	mv ~/qmk_firmware/dztech_dz60rgb_wkl_v2_1_haunt98.bin .

json:
	$(MAKE) cp
	rm -rf keymaps_json/haunt98/keymap.json
	qmk c2json -kb dztech/dz60rgb_wkl/v2_1 -km haunt98 -o keymaps_json/haunt98/keymap.json keymaps/haunt98/keymap.c

clean:
	rm -rf dztech_dz60rgb_wkl_v2_1_haunt98.bin
