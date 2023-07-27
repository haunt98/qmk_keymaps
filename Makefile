.PHONY: all format clean go dztech_dz60rgb_wkl

all:
	qmk setup -H ~/qmk_firmware
	qmk git-submodule
	qmk doctor
	$(MAKE) format
	$(MAKE) go
	$(MAKE) dztech_dz60rgb_wkl

format:
	clang-format -i dztech_dz60rgb_wkl/keymaps/haunt98/*.c dztech_dz60rgb_wkl/keymaps/haunt98/*.h

clean:
	qmk clean
	rm -rf ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/haunt98
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
	# Draw
	curl https://raw.githubusercontent.com/qmk/qmk_firmware/master/keyboards/dztech/dz60rgb_wkl/info.json --output dztech_dz60rgb_wkl/info.json
	QMK_INFO=dztech_dz60rgb_wkl/info.json QMK_KEYMAP=dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json OUT=dztech_dz60rgb_wkl/asciiart/haunt98.txt POST_PROCESS_TABLE=true go run ./cmd/qmkasciigen/*.go
