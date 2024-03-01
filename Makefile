.PHONY: all go draw_qmkasciigen format_draw_caksoylar_keymap_drawer draw_caksoylar_keymap_drawer draw qmk_upstream qmk_format qmk_compile qmk_clean qmk_c2json qmk_flash_sofle_rev1

all:
	$(MAKE) qmk_all
	$(MAKE) draw

# From changeloguru
go:
	# go install github.com/haunt98/gofimports/cmd/gofimports@latest
	# go install mvdan.cc/gofumpt@latest
	gofimports -w --company github.com/make-go-great,github.com/haunt98 .
	gofumpt -w -extra .
	# go install github.com/haunt98/go-test-color@latest
	go-test-color -race -failfast ./...
	golangci-lint run ./...

draw_qmkasciigen:
	$(MAKE) go
	go run ./cmd/qmkasciigen \
		-qmk-keyboard dztech/dz60rgb_wkl/v2_1 \
		-qmk-keymap-file dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json \
		-out dztech_dz60rgb_wkl/asciiart/haunt98.txt
	go run ./cmd/qmkasciigen \
		-qmk-keyboard gray_studio/think65/solder \
		-qmk-keymap-file gray_studio_think65_solder/keymaps_json/haunt98/keymap.json \
		-out gray_studio_think65_solder/asciiart/haunt98.txt
	go run ./cmd/qmkasciigen \
		-qmk-keyboard sofle/rev1 \
		-qmk-keymap-file sofle_rev1/keymaps_json/haunt98/keymap.json \
		-out sofle_rev1/asciiart/haunt98.txt

draw_qmkasciigen_demo:
	go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard matthewdias/m3n3van
	go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard matthewdias/minim
	go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard moondrop/dash75
	go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard ymdk/id75 -qmk-keymap via

format_draw_caksoylar_keymap_drawer:
	# bun upgrade
	bunx prettier@latest --write \
		dztech_dz60rgb_wkl/caksoylar_keymap_drawer/*.yaml \
		sofle_rev1/caksoylar_keymap_drawer/*.yaml

draw_caksoylar_keymap_drawer:
	# https://github.com/caksoylar/keymap-drawer
	# pipx install keymap-drawer
	# pipx upgrade keymap-drawer
	keymap -c caksoylar_keymap_drawer/config.yaml \
		parse -q dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json > \
		dztech_dz60rgb_wkl/caksoylar_keymap_drawer/keymap.yaml
	keymap -c caksoylar_keymap_drawer/config.yaml \
		draw dztech_dz60rgb_wkl/caksoylar_keymap_drawer/keymap.yaml > \
		dztech_dz60rgb_wkl/caksoylar_keymap_drawer/keymap.svg
	curl https://raw.githubusercontent.com/caksoylar/keymap-drawer/main/resources/qmk_layouts/sofle_rotated.json --output sofle_rev1/caksoylar_keymap_drawer/sofle_rotated.json
	keymap -c caksoylar_keymap_drawer/config.yaml \
		parse -q sofle_rev1/keymaps_json/haunt98/keymap.json > \
		sofle_rev1/caksoylar_keymap_drawer/keymap.yaml
	keymap -c caksoylar_keymap_drawer/config.yaml \
		draw sofle_rev1/caksoylar_keymap_drawer/keymap.yaml -j sofle_rev1/caksoylar_keymap_drawer/sofle_rotated.json > \
		sofle_rev1/caksoylar_keymap_drawer/keymap.svg

draw:
	$(MAKE) draw_qmkasciigen
	$(MAKE) draw_caksoylar_keymap_drawer
	$(MAKE) format_draw_caksoylar_keymap_drawer

qmk_all:
	$(MAKE) qmk_upstream
	$(MAKE) qmk_format
	$(MAKE) qmk_compile
	$(MAKE) qmk_c2json
	$(MAKE) qmk_clean

qmk_upstream:
	git -C ~/qmk_firmware pull
	qmk git-submodule
	qmk doctor
	curl https://raw.githubusercontent.com/qmk/qmk_firmware/master/.clang-format --output .clang-format

qmk_format:
	clang-format -i dztech_dz60rgb_wkl/keymaps/haunt98/*.c \
		dztech_dz60rgb_wkl/keymaps/haunt98/*.h \
		gray_studio_think65_solder/keymaps/haunt98/*.c \
		gray_studio_think65_solder/keymaps/haunt98/*.h \
		sofle_rev1/keymaps/haunt98/*.c \
		sofle_rev1/keymaps/haunt98/*.h \
		users/haunt98/*.c \
		users/haunt98/*.h

qmk_cp:
	$(MAKE) qmk_clean
	cp -rf users/haunt98 ~/qmk_firmware/users/
	cp -rf dztech_dz60rgb_wkl/keymaps/haunt98 ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/
	cp -rf gray_studio_think65_solder/keymaps/haunt98 ~/qmk_firmware/keyboards/gray_studio/think65/solder/keymaps/
	cp -rf sofle_rev1/keymaps/haunt98 ~/qmk_firmware/keyboards/sofle/keymaps/

qmk_compile:
	$(MAKE) qmk_cp
	# qmk compile -j 8 -kb dztech/dz60rgb_wkl/v2_1 -km haunt98
	# mv ~/qmk_firmware/dztech_dz60rgb_wkl_v2_1_haunt98.bin .
	qmk compile -j 8 -kb gray_studio/think65/solder -km haunt98
	mv ~/qmk_firmware/gray_studio_think65_solder_haunt98.hex .
	qmk compile -j 8 -kb sofle/rev1 -km haunt98

qmk_clean:
	rm -rf ~/qmk_firmware/users/haunt98
	rm -rf ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/haunt98
	rm -rf ~/qmk_firmware/keyboards/gray_studio/think65/solder/keymaps/haunt98
	rm -rf ~/qmk_firmware/keyboards/sofle/keymaps/haunt98
	qmk clean
	rm -rf *.bin *.hex

qmk_c2json:
	rm -rf dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json
	qmk c2json -kb dztech/dz60rgb_wkl/v2_1 -km haunt98 -o dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json --no-cpp \
		dztech_dz60rgb_wkl/keymaps/haunt98/keymap.c
	qmk format-json -i dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json
	rm -rf gray_studio_think65_solder/keymaps_json/haunt98/keymap.json
	qmk c2json -kb gray_studio/think65/solder -km haunt98 -o gray_studio_think65_solder/keymaps_json/haunt98/keymap.json --no-cpp \
		gray_studio_think65_solder/keymaps/haunt98/keymap.c
	qmk format-json -i gray_studio_think65_solder/keymaps_json/haunt98/keymap.json
	rm -rf sofle_rev1/keymaps_json/haunt98/keymap.json
	qmk c2json -kb sofle/rev1 -km haunt98 -o sofle_rev1/keymaps_json/haunt98/keymap.json --no-cpp \
		sofle_rev1/keymaps/haunt98/keymap.c
	qmk format-json -i sofle_rev1/keymaps_json/haunt98/keymap.json

qmk_flash_dztech_dz60rgb_wkl:
	$(MAKE) qmk_cp
	qmk flash -c -j 8 -kb dztech/dz60rgb_wkl/v2_1 -km haunt98
	$(MAKE) qmk_clean

qmk_flash_gray_studio_think65_solder:
	$(MAKE) qmk_cp
	qmk flash -c -j 8 -kb gray_studio/think65/solder -km haunt98
	$(MAKE) qmk_clean

qmk_flash_sofle_rev1:
	$(MAKE) qmk_cp
	qmk flash -c -j 8 -kb sofle/rev1 -km haunt98
	$(MAKE) qmk_clean
