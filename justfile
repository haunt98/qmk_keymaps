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
    	-qmk-keyboard hineybush/h60 \
    	-qmk-keymap-file hineybush_h60/keymaps_json/haunt98/keymap.json \
    	-out hineybush_h60/asciiart/haunt98.txt
    go run ./cmd/qmkasciigen \
    	-qmk-keyboard sofle/rev1 \
    	-qmk-keymap-file sofle_rev1/keymaps_json/haunt98/keymap.json \
    	-out sofle_rev1/asciiart/haunt98.txt
    go run ./cmd/qmkasciigen \
    	-qmk-keyboard crkbd/rev1 \
    	-qmk-keymap-file crkbd_rev1/keymaps_json/haunt98/keymap.json \
    	-out crkbd_rev1/asciiart/haunt98.txt
    go run ./cmd/qmkasciigen \
        -qmk-keyboard coban/pad9a \
        -qmk-keymap-file coban_pad9a/keymaps_json/haunt98/keymap.json \
        -out coban_pad9a/asciiart/haunt98.txt

draw_qmkasciigen_demo:
    go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard matthewdias/m3n3van
    go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard matthewdias/minim
    go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard moondrop/dash75
    go run ./cmd/qmkasciigen -print-out -allow-layers 0 -qmk-keyboard ymdk/id75 -qmk-keymap via

draw_keymap_drawer_init:
    mkdir -p hineybush_h60/keymap_drawer
    mkdir -p sofle_rev1/keymap_drawer
    mkdir -p crkbd_rev1/keymap_drawer
    mkdir -p coban_pad9a/keymap_drawer

draw_keymap_drawer: draw_keymap_drawer_init
    # https://github.com/caksoylar/keymap-drawer
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
    	parse -q crkbd_rev1/keymaps_json/haunt98/keymap.json > \
    	crkbd_rev1/keymap_drawer/keymap.yaml
    keymap -c keymap_drawer/config.yaml \
    	draw crkbd_rev1/keymap_drawer/keymap.yaml -j crkbd_rev1/keymap_drawer/corne_rotated.json > \
    	crkbd_rev1/keymap_drawer/keymap.svg
    keymap -c keymap_drawer/config.yaml \
        parse -q coban_pad9a/keymaps_json/haunt98/keymap.json > \
        coban_pad9a/keymap_drawer/keymap.yaml
    keymap -c keymap_drawer/config.yaml \
        draw coban_pad9a/keymap_drawer/keymap.yaml > \
        coban_pad9a/keymap_drawer/keymap.svg

draw_keymap_drawer_upstream: draw_keymap_drawer_init
    wcurl --curl-options="--clobber --netrc" https://raw.githubusercontent.com/caksoylar/keymap-drawer/main/resources/extra_layouts/sofle_rotated.json --output sofle_rev1/keymap_drawer/sofle_rotated.json
    wcurl --curl-options="--clobber --netrc" https://raw.githubusercontent.com/caksoylar/keymap-drawer/main/resources/extra_layouts/corne_rotated.json --output crkbd_rev1/keymap_drawer/corne_rotated.json

draw_keymap_drawer_format:
    bunx prettier --write \
        hineybush_h60/keymap_drawer/*.yaml \
    	sofle_rev1/keymap_drawer/*.yaml \
    	crkbd_rev1/keymap_drawer/*.yaml \
        coban_pad9a/keymap_drawer/*.yaml

draw: draw_qmkasciigen draw_keymap_drawer draw_keymap_drawer_format

qmk: qmk_clean qmk_format qmk_compile qmk_c2json

qmk_upstream:
    # curl -fsSL https://install.qmk.fm | sh
    qmk doctor
    git -C ~/qmk_firmware pull origin master --tags
    qmk git-submodule
    qmk doctor
    curl https://raw.githubusercontent.com/qmk/qmk_firmware/master/.clang-format --output .clang-format

qmk_format:
    fd --extension c --extension h --exec clang-format -i

qmk_cp:
    cp -rf users/haunt98 ~/qmk_firmware/users/
    cp -rf hineybush_h60/keymaps/haunt98 ~/qmk_firmware/keyboards/hineybush/h60/keymaps/
    cp -rf sofle_rev1/keymaps/haunt98 ~/qmk_firmware/keyboards/sofle/keymaps/
    cp -rf crkbd_rev1/keymaps/haunt98 ~/qmk_firmware/keyboards/crkbd/keymaps/
    cp -rf coban_pad9a/keymaps/haunt98 ~/qmk_firmware/keyboards/coban/pad9a/keymaps/

qmk_compile: qmk_cp
    # Ignore qmk lint as header errors
    qmk compile -c -kb hineybush/h60 -km haunt98
    mv ~/qmk_firmware/hineybush_h60_haunt98.hex .
    qmk compile -c -kb sofle/rev1 -km haunt98
    qmk compile -c -kb crkbd/rev1 -km haunt98
    qmk compile -c -kb coban/pad9a -km haunt98

qmk_clean:
    rm -rf ~/qmk_firmware/users/haunt98
    rm -rf ~/qmk_firmware/keyboards/hineybush/h60/keymaps/haunt98
    rm -rf ~/qmk_firmware/keyboards/sofle/keymaps/haunt98
    rm -rf ~/qmk_firmware/keyboards/crkbd/keymaps/haunt98
    rm -rf ~/qmk_firmware/keyboards/coban/pad9a/keymaps/haunt98
    qmk clean -a
    rm -rf *.bin *.hex *.uf2

qmk_c2json:
    fd "keymap.json" --exec-batch rm
    qmk c2json -kb hineybush/h60 -km haunt98 -o hineybush_h60/keymaps_json/haunt98/keymap.json --no-cpp \
    	hineybush_h60/keymaps/haunt98/keymap.c
    qmk c2json -kb sofle/rev1 -km haunt98 -o sofle_rev1/keymaps_json/haunt98/keymap.json --no-cpp \
    	sofle_rev1/keymaps/haunt98/keymap.c
    qmk c2json -kb crkbd/rev1 -km haunt98 -o crkbd_rev1/keymaps_json/haunt98/keymap.json --no-cpp \
    	crkbd_rev1/keymaps/haunt98/keymap.c
    qmk c2json -kb coban/pad9a -km haunt98 -o coban_pad9a/keymaps_json/haunt98/keymap.json --no-cpp \
    	coban_pad9a/keymaps/haunt98/keymap.c
    fd "keymap.json" --exec qmk format-json -i

qmk_flash_sofle_rev1: qmk_clean qmk_cp
    qmk flash -c -kb sofle/rev1 -km haunt98

qmk_flash_sofle_rev1_default: qmk_clean
    qmk flash -c -kb sofle/rev1 -km default -e CONVERT_TO=rp2040_ce

qmk_flash_crkbd_rev1: qmk_clean qmk_cp
    qmk flash -c -kb crkbd/rev1 -km haunt98

qmk_flash_crkbd_rev1_default: qmk_clean
    qmk flash -c -kb crkbd/rev1 -km default -e CONVERT_TO=rp2040_ce

qmk_flash_coban_pad9a: qmk_clean qmk_cp
    qmk flash -c -kb coban/pad9a -km haunt98

qmk_flash_coban_pad9a_default: qmk_clean
    qmk flash -c -kb coban/pad9a -km default -e CONVERT_TO=rp2040_ce
