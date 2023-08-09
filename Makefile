.PHONY: all format clean go draw qmk_upstream qmk_init qmk_clean dztech_dz60rgb_wkl

all:
	$(MAKE) format
	$(MAKE) draw
	$(MAKE) qmk_upstream

format:
	clang-format -i dztech_dz60rgb_wkl/keymaps/haunt98/*.c dztech_dz60rgb_wkl/keymaps/haunt98/*.h & \
		clang-format-mp-15 -i dztech_dz60rgb_wkl/keymaps/haunt98/*.c dztech_dz60rgb_wkl/keymaps/haunt98/*.h

clean:
	rm -rf dztech_dz60rgb_wkl_v2_1_haunt98.bin

# From changeloguru
go:
	go install github.com/haunt98/go-test-color@latest
	go-test-color -race -failfast ./...
	golangci-lint run ./...
	go install github.com/haunt98/gofimports/cmd/gofimports@latest
	go install mvdan.cc/gofumpt@latest
	gofimports -w --company github.com/make-go-great,github.com/haunt98 .
	gofumpt -w -extra .

draw:
	$(MAKE) go
	go run ./cmd/qmkasciigen/*.go -qmk-keyboard dztech/dz60rgb_wkl/v2_1 -qmk-keymap-file dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json -out dztech_dz60rgb_wkl/asciiart/haunt98.txt
	# https://github.com/caksoylar/keymap-drawer
	pipx install keymap-drawer
	keymap parse -q dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json > dztech_dz60rgb_wkl/caksoylar_keymap_drawer/keymap.yaml
	keymap draw dztech_dz60rgb_wkl/caksoylar_keymap_drawer/keymap.yaml > dztech_dz60rgb_wkl/caksoylar_keymap_drawer/keymap.svg

qmk_upstream:
	curl https://raw.githubusercontent.com/qmk/qmk_firmware/master/.clang-format --output .clang-format

qmk_init:
	qmk setup -H ~/qmk_firmware
	qmk git-submodule
	qmk doctor

qmk_clean:
	qmk clean
	rm -rf ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/haunt98

dztech_dz60rgb_wkl:
	$(MAKE) qmk_init
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
