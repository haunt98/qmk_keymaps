.PHONY: all format cp compile

all:
	$(MAKE) format
	$(MAKE) cp
	$(MAKE) compile

format:
	clang-format -i keymaps/haunt98/*.c keymaps/haunt98/*.h

cp:
	rm -rf ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/haunt98
	cp -rf keymaps/haunt98 ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/

compile:
	qmk compile -kb dztech/dz60rgb_wkl/v2_1 -km haunt98
	mv ~/qmk_firmware/dztech_dz60rgb_wkl_v2_1_haunt98.bin .

clean:
	rm -rf dztech_dz60rgb_wkl_v2_1_haunt98.bin
