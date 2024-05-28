# CHANGELOG

## v0.0.22 (2024-05-28)

### Added

- feat(qmkasciigen): support new qmk keyboard.json (2024-05-28)

- feat(qmk): use encoder for page up/down (2024-04-23)

- feat(qmk): enable SPLIT_WATCHDOG_ENABLE for sofle_rev1 (2024-04-20)

- feat(qmk): remove caps in encoder sofle (2024-04-19)

- feat(qmk): correct led width (2024-04-19)

- feat(qmk): use tri layers (2024-04-19)

- feat(qmk): experiemnt with NO_USB_STARTUP_CHECK (2024-04-16)

- feat(qmk): remove nkro (2024-04-16)

- feat(qmk): re-adds KC_CAPS for sofle_rev1 (2024-04-04)

- feat: put think65 to rest (2024-03-11)

### Others

- chore: use fd to auto find and exec (2024-05-12)

- chore: remove pipx (2024-05-11)

- chore: clearly explain how to flash (2024-05-08)

- chore: sync golangci-lint config (2024-05-08)

- chore(deps): bump golangci/golangci-lint-action from 5 to 6 (#17) (2024-05-07)

- chore: update golangci-lint (2024-05-04)

- chore(deps): bump golangci/golangci-lint-action from 4 to 5 (#16) (2024-04-25)

- chore: add encoder docs (2024-04-23)

- chore(draw): I give up dark mode for now (2024-04-19)

- chore(draw): try to fix dark mode (2/?) (2024-04-19)

- chore(draw): try to fix dark mode (1/?) (2024-04-19)

- config: disable dark mode (2024-04-19)

- chore(draw): fix dark mode (2024-04-19)

- chore(draw): improve icon (1/?) (2024-04-19)

- Revert "chore(draw): remove dark mode (for now)" (2024-04-19)

- chore(draw): remove dark mode (for now) (2024-04-19)

- chore(draw): fix dark_mode config value (2024-04-19)

- chore: more docs (2024-04-19)

- chore: column -t (2024-04-19)

- chore: fix typo (2024-04-19)

- chore: fix typo (2024-04-19)

- chore: re-draw using keymap-drawer (2024-04-15)

- chore: re-enable dztech/dz60rgb_wkl/v2_1 build (2024-04-07)

- config: try to fix git pull freeze (2024-04-07)

- chore: Add The SpaceFN concept to README.md (2024-04-05)

- chore: fix preview (2024-04-01)

- chore: clean up make (2024-04-01)

- chore: re-draw with latest keymap-drawer (2024-03-24)

- chore: change Bootmagic Lite to Bootmagic (2024-03-11)

- Merge branch 'main' of https://github.com/haunt98/qmk_keymaps (2024-03-04)

- chore: revert qmk to 0.23.9 (2024-03-04)

- chore: typo (2024-03-01)

- chore(readme): remove verbose (2024-03-01)

- chore(changelog): generate v0.0.21 (2024-03-01)

## v0.0.21 (2024-03-01)

### Added

- feat(qmk): move mouse wheel to layer adj (2024-02-27)

- feat(qmk): enable mouse key for sofle_rev1 (2024-01-08)

- feat(qmk): add mouse wheel in layer 2 (2024-01-05)

- feat(qmk): add mouse wheel (2024-01-05)

- feat(go): split draw func (2024-01-01)

### Fixed

- fix(qmk): try to fix :wq when using mod (2024-01-08)

### Others

- chore: remove make parallel (2024-03-01)

- chore: qmk clean -a (2024-03-01)

- chore: more build target (2024-03-01)

- chore(readme): update flash notes (2024-03-01)

- chore: rerun make (2024-02-29)

- Revert "chore(workflows): see qmk doctor" (2024-02-28)

- chore(workflows): see qmk doctor (2024-02-28)

- chore: disable dz60rgb_wkl for now until qmk nix fixed it (2024-02-28)

- Merge branch 'main' of https://github.com/haunt98/qmk_keymaps (2024-02-27)

- chore: disable go install (2024-02-10)

- chore(qmkasciigen): fix lint (2024-02-10)

- chore: go offline for bun, pipx (2024-01-08)

- chore: add TODO (2024-01-01)

- chore(deps): bump actions/upload-artifact from 3 to 4 (2023-12-15)

- chore(deps): bump actions/setup-go from 4 to 5 (2023-12-07)

- chore: prettier latest (2023-12-01)

- chore(changelog): generate v0.0.20 (2023-11-29)

## v0.0.20 (2023-11-29)

### Added

- feat(qmk): change debounce algo (2023-11-16)

- feat(qmk): experiment (2023-11-16)

- feat(qmk): use sym_defer_pk because noise-resistant (2023-11-14)

- feat(qmk): revision sofle_rev1 (2023-11-13)

- feat(qmk): sofle_rev1 tapping term 120-200 (2023-11-09)

- feat(qmk): move tapping term to out of users (2023-11-09)

- feat(qmk): another revision sofle_rev1 (2023-11-06)

- feat(qmk): another revision for sofle_rev1 (2023-11-02)

- feat(qmk): another revision sofle_rev1 (2023-11-02)

- feat(qmk): another revision sofle_rev1 (2023-11-02)

- feat(qmk): revision (2023-11-02)

- feat(qmk): use cyberpunk logo (2023-10-30)

- feat(qmk): move oled to users (2023-10-30)

- feat(qmk): revision sofle_rev1 (2023-10-30)

- feat(qmk): revision sofle_rev1 (2023-10-26)

- feat(qmk): try to fix debounce sofle_rev1 (2023-10-23)

- feat(qmk): revision sofle_rev1 (2023-10-23)

- feat(qmk): revision sofle_rev1 (2023-10-23)

- feat(qmk): revision sofle_rev1 (2023-10-23)

- feat(qmk): revision sofle_rev1 (2023-10-21)

- feat(draw): upgrade caksoylar_keymap_drawer v0.12.0 (2023-10-21)

- feat(qmk): revise layout (2023-10-20)

- feat(qmk): use _ in left hand sofle_rev1 (2023-10-19)

- feat(qmk): use : for vim (2023-10-19)

- feat(qmk): use define to reduce dupe code (2023-10-19)

- feat(qmk): add ENCODER_RESOLUTION (2023-10-19)

- feat(qmk): add debounce sofle_rev1 (2023-10-18)

- feat(qmk): revise sofle (2023-10-18)

- feat(qmk): update keymap sofle rev1 (2023-10-18)

- feat(qmk): add oled for sofle rev1 (2023-10-18)

### Others

- chore(deps): bump github.com/spf13/cast from 1.5.1 to 1.6.0 (2023-11-29)

- chore: more explain (2023-10-31)

- chore: clean up (2023-10-26)

- chore: sync qmk first (2023-10-23)

- chore: more qmk (2023-10-20)

- chore(draw): re sync (2023-10-19)

- chore(draw): sofle rev1 (2023-10-18)

- chore(draw): re draw sofle rev1 (2023-10-18)

- chore: bring back qmk (2023-10-18)

- chore: bring back qmk to Makefile (2023-10-18)

- chore: replace deno fmt -> qmk format-json (2023-10-18)

- chore(changelog): generate v0.0.19 (2023-10-14)

## v0.0.19 (2023-10-14)

### Added

- feat: support sofle (#9) (2023-10-14)

- feat(qmk): debounce on keyboard (2023-10-06)

### Others

- chore: move tools to my personal blogs (2023-10-06)

- chore(changelog): generate v0.0.18 (2023-09-29)

## v0.0.18 (2023-09-29)

### Added

- feat: support think65 (#8) (2023-09-29)

### Others

- chore: clena up readme (2023-09-29)

- chore: use clang-format from nix (2023-09-29)

- chore: add more badges (2023-09-23)

- chore: qmk gh action run daily (2023-09-20)

- chore: gh action run every hour (2023-09-20)

- chore: update README (2023-09-19)

- chore(changelog): generate v0.0.17 (2023-09-16)

## v0.0.17 (2023-09-16)

### Added

- feat(qmk): use qmk userspace (#7) (2023-09-16)

- feat(qmk): force nkro (2023-09-13)

### Others

- chore: more qmk read (2023-09-13)

- chore(qmk): add comment (2023-09-11)

- qmk: rgb matrix efffect (2023-09-11)

- qmk: update rgb matrix effect (2023-09-11)

- chore: add build develop qmk (#6) (2023-09-11)

- chore: update copyright for C file (2023-09-10)

- chore(changelog): generate v0.0.16 (2023-09-09)

## v0.0.16 (2023-09-09)

### Added

- feat: use CP437 character (#5) (2023-09-09)

- feat: remove thera (2023-09-07)

- feat: support thera (1/?) (2023-09-06)

- feat(qmk): add kc_del (2023-08-17)

- feat(qmk): add RGB_MATRIX_TIMEOUT (2023-08-15)

- feat(qmk): add USB_SUSPEND_WAKEUP_DELAY 200 (2023-08-11)

- feat(qmkasciigen): support caps (2023-08-10)

### Others

- chore: update TODO (2023-09-09)

- qmk: enable rgb matrix (2023-09-09)

- draw: re-draw (2023-09-07)

- qmk: right alt -> mo(2) (2023-09-07)

- chore: remove build locally (2023-09-06)

- qmk: optimize thera more (2023-09-06)

- draw: generate thera ascii art (2023-09-06)

- chore: remove qmk local (2023-09-06)

- chore: fix gh action (2023-09-06)

- chore: bump thera using qmk cli (2023-09-06)

- chore(deps): bump actions/checkout from 3 to 4 (2023-09-05)

- chore(draw): update style (2023-08-29)

- chore: upgrade draw (2023-08-29)

- chore: move hardware out of picture (2023-08-23)

- chore: more design (2023-08-20)

- chore: limit pro micro (2023-08-20)

- chore: re-draw (2023-08-17)

- chore: more interesting (2023-08-16)

- chore: more interesting (2023-08-13)

- chore: fix typo (2023-08-13)

- chore: add stock github repo (2023-08-11)

- chore: update license to match MIT (2023-08-11)

- chore: add karabiner (2023-08-11)

- chore: fix typo (2023-08-10)

- chore: fix typo (2023-08-10)

- chore: re-order posts (2023-08-10)

- chore(changelog): generate v0.0.15 (2023-08-10)

## v0.0.15 (2023-08-10)

### Added

- feat(qmkasciigen): re add support mod tap (2023-08-10)

- feat(qmkasciigen): re-order draw transform (2023-08-10)

- feat(qmkasciigen): add flag allow layers (2023-08-10)

- feat(qmkasciigen): add print-out flag (2023-08-10)

- feat(qmkasciigen): support add binding tiny (2023-08-10)

- feat(qmkasciigen): split raw binding and transform (2023-08-10)

- feat(qmk): change to sym_eager_pk (2023-08-10)

- feat(qmk): use layer name in keymaps (2023-08-09)

- Revert "feat(qmk): enable auto shift" (2023-08-09)

- feat(qmk): enable auto shift (2023-08-09)

- feat(qmk): use shift for space cadet not alt (2023-08-09)

### Fixed

- fix(qmk): remove redundant , in layout fn (2023-08-09)

### Others

- chore: better demo (2023-08-10)

- chore: add draw demo (2023-08-10)

- chore: sync latest json (2023-08-10)

- chore: more TODO (2023-08-10)

- chore(draw): docs (2023-08-10)

- chore: re draw (2023-08-10)

- chore: reformat (2023-08-10)

- chore(draw): big changes (2023-08-10)

- chore: format draw (2023-08-10)

- chore: format after draw (2023-08-10)

- chore: format yaml (2023-08-10)

- chore(draw): bare bone config (2023-08-10)

- chore(draw): map safe boot when (2023-08-10)

- chore: format keymap.json (2023-08-09)

- chore(qmk): reorder keymap by hand (2023-08-09)

- chore: add qmk upstream to sync with qmk master (2023-08-09)

- chore: sync clang format with upstream qmk (2023-08-09)

- chore(qmk): ignore clang format keymaps (2023-08-09)

- chore: don't space cadet (2023-08-09)

- chore: re draw (2023-08-09)

- chore: sync keymap json (2023-08-09)

- chore(changelog): generate v0.0.14 (2023-08-08)

## v0.0.14 (2023-08-08)

### Added

- feat(qmk): naming layers (2023-08-07)

### Fixed

- fix(qmk): missing ; in define enum (2023-08-07)

### Others

- chore: add bloznelis/kbt (2023-08-08)

- chore: goody for gh action (2023-08-08)

- chore: add qmk lint in gh action (2023-08-08)

- chore(changelog): generate v0.0.13 (2023-08-07)

## v0.0.13 (2023-08-07)

### Added

- feat: add tap dance (wip 4/?) (2023-08-07)

- Revert "feat: add tap dance (wip 3/?)" (2023-08-07)

- feat: add tap dance (wip 3/?) (2023-08-07)

- feat: add tap dance (wip 2/?) (2023-08-07)

- feat: add tap dance (wip 1/?) (2023-08-07)

### Others

- chore: fix typo (2023-08-07)

- chore: redraw (2023-08-07)

- chore(changelog): generate v0.0.12 (2023-08-07)

## v0.0.12 (2023-08-07)

### Added

- feat(qmkasciigen): support space cadet (2023-08-07)

- feat: increase tapping term (2023-08-07)

- feat(qmk): try to re-enable space cadet (2023-08-07)

- feat(qmk): add quick tap term and retro tapping (2023-08-07)

- Revert "feat(qmk): re-enable space cadet" (2023-08-07)

- feat(qmk): re-enable space cadet (2023-08-07)

- feat(qmk): prefer opt to alt, cmd to gui (2023-08-07)

- feat(qmk): remove retro tapping (2023-08-07)

- feat: move vol, brightness to layer 2 (2023-07-30)

- feat(qmkasciigen): local first, remote later (2023-07-30)

- feat(qmkasciigen): get qmk info from longest qmk keyboard name first (2023-07-30)

- feat(qmkasciigen): add flag to print layout, layer (2023-07-30)

### Others

- chore: fix typo (2023-08-06)

- chore: edit keymap (2023-08-06)

- chore: re-add caksoylar/keymap-drawer (2023-08-06)

- chore: add qmk info (2023-08-06)

- chore: split intetresting (2023-08-03)

- chore: more interesting (2023-08-03)

- chore: more interesting (2023-07-31)

- chore: remove default qmk compare from Makefile (2023-07-30)

- chore: add maintain (2023-07-30)

- chore: fix failed gh action (2023-07-30)

- chore: convert keymaps json using gh action (2023-07-30)

- chore: re gen ascii (2023-07-30)

- chore: re-gen asciiart (2023-07-30)

- chore(changelog): generate v0.0.11 (2023-07-30)

## v0.0.11 (2023-07-30)

### Added

- feat(qmk): split 4 layer (2023-07-30)

- feat(qmk): uncomment via config (2023-07-30)

- feat(qmk): remove music because its default is already no (2023-07-30)

### Others

- chore: re gen asciiart (2023-07-30)

- chore: more interesting (2023-07-30)

- chore(changelog): generate v0.0.10 (2023-07-29)

## v0.0.10 (2023-07-29)

### Added

- feat(qmk): adjust macOS brigthness (2023-07-29)

### Others

- chore(asciiart): redraw (2023-07-29)

- chore(qmkasciigen): consistent naming (2023-07-28)

- chore(changelog): generate v0.0.9 (2023-07-28)

## v0.0.9 (2023-07-28)

### Fixed

- fix: prevent colliding draw map (2023-07-28)

### Others

- chore(changelog): generate v0.0.8 (2023-07-28)

## v0.0.8 (2023-07-28)

### Added

- feat: revert back to backspace for now (2023-07-28)

- feat(qmkasciigen): split from to keycodes (2023-07-28)

### Others

- chore(changelog): generate v0.0.7 (2023-07-28)

## v0.0.7 (2023-07-28)

### Added

- feat(qmkasciigen): hide _______, XXXXXXX (2023-07-28)

### Others

- chore(changelog): generate v0.0.6 (2023-07-28)

## v0.0.6 (2023-07-28)

### Added

- feat(qmkasciigen): get remote keymap (2023-07-28)

- feat(qmkasciigen): get info json from qmk repo (2023-07-28)

- feat(qmkasciigen): use flag instead of os.ENV (2023-07-28)

- feat: set RGB_MATRIX_LED_FLUSH_LIMIT (2023-07-28)

### Fixed

- Revert "fix: ghosting rgb led (wip)" (2023-07-28)

- fix: ghosting rgb led (wip) (2023-07-28)

### Others

- chore(qmkasciigen): update docs (2023-07-28)

- chore: default make only format and compile (2023-07-28)

- chore: update readme (2023-07-28)

- chore: set debounce sym_eager_pr (2023-07-28)

- chore(changelog): generate v0.0.5 (2023-07-28)

## v0.0.5 (2023-07-28)

### Added

- feat: limit rgb matrix brightness (2023-07-28)

### Others

- chore: fix gh action (2023-07-28)

- chore: only run qmk gh action on c, h, mk files changed (2023-07-28)

- chore: remove firmware size (2023-07-28)

- chore(changelog): generate v0.0.4 (2023-07-28)

## v0.0.4 (2023-07-28)

### Others

- chore: guide qmk gh action workflow (2023-07-28)

- chore: use gh action badge (2023-07-28)

- chore: change artifact name (2023-07-28)

- chore: fix qmk gh action (2023-07-28)

- chore: fix qmk gh action (2023-07-28)

- chore: fix git checkout gh action (2023-07-28)

- chore: fix qmk github action (2023-07-28)

- chore: add qmk github action (2023-07-28)

- chore: add github action (2023-07-28)

- chore: Docker todo (2023-07-28)

- chore(changelog): generate v0.0.3 (2023-07-25)

## v0.0.3 (2023-07-25)

### Added

- feat(draw): remove post process table (2023-07-25)

- feat: remove shift enter, only enter (2023-07-24)

### Others

- chore(qmkasciigen): cleanup (2023-07-25)

- chore: simple license in code (2023-07-24)

- chore: add LICENSE (2023-07-24)

- chore(changelog): generate v0.0.2 (2023-07-24)

## v0.0.2 (2023-07-24)

### Added

- feat: change activate layer 2 using fn ` (2023-07-24)

- feat: move TG(2) from left alt to right command (2023-07-24)

- feat(draw): make post process table optional (2023-07-24)

- feat: remove tap dance (can not make it work) (2023-07-24)

- feat: better draw strategy (wip) (2023-07-23)

### Fixed

- fix(draw): init table with empty space to fix think65 (2023-07-24)

### Others

- chore: re-organize Makefile (2023-07-24)

- chore(draw): fix lint (2023-07-24)

- chore(draw): explain draw strategy (2023-07-23)

- chore(changelog): generate v0.0.1 (2023-07-23)

## v0.0.1 (2023-07-23)

### Added

- feat(draw): make output file is optional (2023-07-23)

- feat: use ACTION_TAP_DANCE_FN for safeboot (2023-07-23)

- feat: add tapdance for safe boot (2023-07-23)

- feat: increase debounce time (2023-07-23)

- feat: use shift enter (wip) (2023-07-23)

- feat: add layout (2023-07-23)

- feat: draw layer (2023-07-23)

- feat: draw almost keycodes (2023-07-23)

- feat: change keymap again (2023-07-23)

- feat: center key when draw (2023-07-23)

- feat: draw more symbol (2023-07-23)

- feat: gen ascii art with Go (2023-07-23)

- feat: remove usb delay config (wip) (2023-07-23)

- feat: add USB_SUSPEND_WAKEUP_DELAY (2023-07-22)

- feat: use KC_GRV in layer 1 (2023-07-22)

- feat: config mod term (2023-07-22)

- feat: disable VIA, NKRO (2023-07-22)

- feat: use mod tap for left control and esc (2023-07-22)

- feat: lower debounce to 5 (2023-07-21)

- feat: debounce 10ms (2023-07-21)

- feat: add qmk lint (2023-07-21)

- feat: disable more rgb (2023-07-21)

- feat: disable mod-tap (2023-07-21)

- feat: re-enable audio key (2023-07-21)

- feat: remove one shot (2023-07-21)

- feat: disable magic keycodes (2023-07-21)

- feat: disable space cadet (2023-07-21)

- feat: only have one mo(2) (2023-07-21)

- feat: remove del, add boot aka reset (2023-07-21)

- feat: new keymap (2023-07-21)

- feat: remove layer 3 (2023-07-21)

- feat: disable unused features (2023-07-21)

### Others

- chore: merge draw to keyboard (2023-07-23)

- chore: rename to qmkasciigen (2023-07-23)

- chore: draw + better (2023-07-23)

- chore: more interesting (2023-07-23)

- chore: support TD_ when draw (2023-07-23)

- chore: add pragma once in config.h (2023-07-23)

- chore: draw rgb (2023-07-23)

- chore: better draw this time (2023-07-23)

- chore: remove draw (2023-07-23)

- chore: qmk doctor (2023-07-22)

- chore: add more interesting (2023-07-22)

- dztech_dz60rgb_wkl: use hud, sai, vai directly (2023-07-22)

- chore: more interesting (2023-07-22)

- chore: more keeb in the future (2023-07-21)

- chore: compile faster (2023-07-21)

- chore: update docs (2023-07-21)

- chore: add comment (2023-07-21)

- chore: improve docs (2023-07-21)

- chore: add via config qmk (2023-07-21)

- chore: fix typo (2023-07-21)

- chore: add stock firmware (2023-07-21)

- chore: how to flash (2023-07-21)

- chore: add qmk clean (2023-07-21)

- chore fix typo (2023-07-21)

- chore: generate qmk json (2023-07-21)

- chore: add TODO (2023-07-21)

- chore: add Makefile (2023-07-21)

- chore: re-organize (2023-07-21)

- chore: more docs (2023-07-21)

- chore: more docs (2023-07-21)

- chore: add README (2023-07-21)

- chore init with barebones (2023-07-21)
