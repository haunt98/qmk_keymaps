all: qmk draw

# From changeloguru
go:
    # go install github.com/haunt98/gofimports/cmd/gofimports@latest
    # go install mvdan.cc/gofumpt@latest
    gofimports -w --company github.com/make-go-great,github.com/haunt98 .
    gofumpt -w -extra .
    # go install github.com/haunt98/go-test-color@latest
    go-test-color -race -failfast ./...
    golangci-lint run --fix ./...
    modernize -fix -test ./...

draw_qmkasciigen: go
    go run ./cmd/qmkasciigen \
    	-qmk-keyboard dztech/dz60rgb_wkl/v2_1 \
    	-qmk-keymap-file dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json \
    	-out dztech_dz60rgb_wkl/asciiart/haunt98.txt
    go run ./cmd/qmkasciigen \
    	-qmk-keyboard hineybush/h60 \
    	-qmk-keymap-file hineybush_h60/keymaps_json/haunt98/keymap.json \
    	-out hineybush_h60/asciiart/haunt98.txt
    go run ./cmd/qmkasciigen \
    	-qmk-keyboard sofle/rev1 \
    	-qmk-keymap-file sofle_rev1/keymaps_json/haunt98/keymap.json \
    	-out sofle_rev1/asciiart/haunt98.txt
    go run ./cmd/qmkasciigen \
        -qmk-keyboard coban/pad9a \
        -qmk-keymap-file coban_pad9a/keymaps_json/haunt98/keymap.json \
        -out coban_pad9a/asciiart/haunt98.txt

draw_qmkasciigen_demo:
    go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard matthewdias/m3n3van
    go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard matthewdias/minim
    go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard moondrop/dash75
    go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard ymdk/id75 -qmk-keymap via

draw_keymap_drawer:
    # https://github.com/caksoylar/keymap-drawer
    keymap -c keymap_drawer/config.yaml \
    	parse -q dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json > \
    	dztech_dz60rgb_wkl/keymap_drawer/keymap.yaml
    keymap -c keymap_drawer/config.yaml \
    	draw dztech_dz60rgb_wkl/keymap_drawer/keymap.yaml > \
    	dztech_dz60rgb_wkl/keymap_drawer/keymap.svg
    keymap -c keymap_drawer/config.yaml \
    	parse -q hineybush_h60/keymaps_json/haunt98/keymap.json > \
    	hineybush_h60/keymap_drawer/keymap.yaml
    keymap -c keymap_drawer/config.yaml \
    	draw hineybush_h60/keymap_drawer/keymap.yaml > \
    	hineybush_h60/keymap_drawer/keymap.svg
    keymap -c keymap_drawer/config.yaml \
    	parse -q sofle_rev1/keymaps_json/haunt98/keymap.json > \
    	sofle_rev1/keymap_drawer/keymap.yaml
    keymap -c keymap_drawer/config.yaml \
    	draw sofle_rev1/keymap_drawer/keymap.yaml -j sofle_rev1/keymap_drawer/sofle_rotated.json > \
    	sofle_rev1/keymap_drawer/keymap.svg
    keymap -c keymap_drawer/config.yaml \
        parse -q coban_pad9a/keymaps_json/haunt98/keymap.json > \
        coban_pad9a/keymap_drawer/keymap.yaml
    keymap -c keymap_drawer/config.yaml \
        draw coban_pad9a/keymap_drawer/keymap.yaml > \
        coban_pad9a/keymap_drawer/keymap.svg

draw_keymap_drawer_upstream:
    wcurl --curl-options="--clobber --netrc" https://raw.githubusercontent.com/caksoylar/keymap-drawer/main/resources/extra_layouts/sofle_rotated.json --output sofle_rev1/keymap_drawer/sofle_rotated.json

draw_keymap_drawer_format:
    bunx prettier --write \
    	dztech_dz60rgb_wkl/keymap_drawer/*.yaml \
        hineybush_h60/keymap_drawer/*.yaml \
    	sofle_rev1/keymap_drawer/*.yaml \
        coban_pad9a/keymap_drawer/*.yaml

draw: draw_qmkasciigen draw_keymap_drawer draw_keymap_drawer_format

qmk: qmk_clean qmk_format qmk_compile qmk_c2json

qmk_upstream:
    qmk doctor
    git -C ~/qmk_firmware pull origin master
    qmk git-submodule
    qmk doctor
    curl https://raw.githubusercontent.com/qmk/qmk_firmware/master/.clang-format --output .clang-format

qmk_format:
    fd --extension c --extension h --exec clang-format -i

qmk_cp:
    cp -rf users/haunt98 ~/qmk_firmware/users/
    cp -rf dztech_dz60rgb_wkl/keymaps/haunt98 ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/
    cp -rf hineybush_h60/keymaps/haunt98 ~/qmk_firmware/keyboards/hineybush/h60/keymaps/
    cp -rf sofle_rev1/keymaps/haunt98 ~/qmk_firmware/keyboards/sofle/keymaps/
    cp -rf coban_pad9a/keymaps/haunt98 ~/qmk_firmware/keyboards/coban/pad9a/keymaps/

qmk_compile: qmk_cp
    # Ignore qmk lint as header errors
    qmk compile -c -kb dztech/dz60rgb_wkl/v2_1 -km haunt98
    mv ~/qmk_firmware/dztech_dz60rgb_wkl_v2_1_haunt98.bin .
    qmk compile -c -kb hineybush/h60 -km haunt98
    mv ~/qmk_firmware/hineybush_h60_haunt98.hex .
    qmk compile -c -kb sofle/rev1 -km haunt98
    qmk compile -c -kb coban/pad9a -km haunt98

qmk_clean:
    rm -rf ~/qmk_firmware/users/haunt98
    rm -rf ~/qmk_firmware/keyboards/dztech/dz60rgb_wkl/keymaps/haunt98
    rm -rf ~/qmk_firmware/keyboards/hineybush/h60/keymaps/haunt98
    rm -rf ~/qmk_firmware/keyboards/sofle/keymaps/haunt98
    rm -rf ~/qmk_firmware/keyboards/coban/pad9a/keymaps/haunt98
    qmk clean -a
    rm -rf *.bin *.hex

qmk_c2json:
    fd "keymap.json" --exec-batch rm
    qmk c2json -kb dztech/dz60rgb_wkl/v2_1 -km haunt98 -o dztech_dz60rgb_wkl/keymaps_json/haunt98/keymap.json --no-cpp \
    	dztech_dz60rgb_wkl/keymaps/haunt98/keymap.c
    qmk c2json -kb hineybush/h60 -km haunt98 -o hineybush_h60/keymaps_json/haunt98/keymap.json --no-cpp \
    	hineybush_h60/keymaps/haunt98/keymap.c
    qmk c2json -kb sofle/rev1 -km haunt98 -o sofle_rev1/keymaps_json/haunt98/keymap.json --no-cpp \
    	sofle_rev1/keymaps/haunt98/keymap.c
    qmk c2json -kb coban/pad9a -km haunt98 -o coban_pad9a/keymaps_json/haunt98/keymap.json --no-cpp \
    	coban_pad9a/keymaps/haunt98/keymap.c
    fd "keymap.json" --exec qmk format-json -i

qmk_flash_sofle_rev1: qmk_clean qmk_cp
    qmk flash -c -kb sofle/rev1 -km haunt98

qmk_flash_coban_pad9a: qmk_clean qmk_cp
    qmk flash -c -kb coban/pad9a -km haunt98
