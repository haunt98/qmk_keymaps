package main

import (
	"fmt"
	"strings"
)

const (
	scaleX = 8
	scaleY = 2
)

// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
var mapRawBinding = []map[string]string{
	{
		// My custom
		"CTL_T(KC_ESC)":    "ESC CTRL",
		"TD(TD_SAFE_BOOT)": "BOOT",
	},
	{
		// Basic
		// Prefer OPT > ALT
		// Prefer CMD > GUI, WIN
		"KC_TRNS": "",
		"_______": "",
		"KC_NO":   "",
		"XXXXXXX": "",
		"KC_ENT":  "ENTER",
		"KC_BSPC": "BACKSPACE",
		"KC_SPC":  "SPACE",
		"KC_MINS": "-",
		"KC_EQL":  "=",
		"KC_LBRC": "[",
		"KC_RBRC": "]",
		"KC_BSLS": "\\",
		"KC_SCLN": ";",
		"KC_QUOT": "'",
		"KC_GRV":  "`",
		"KC_COMM": ",",
		"KC_DOT":  ".",
		"KC_SLSH": "/",
		"KC_LCTL": "CTRL",
		"KC_RCTL": "CTRL",
		"KC_LALT": "OPT",
		"KC_RALT": "OPT",
		"KC_LOPT": "OPT",
		"KC_ROPT": "OPT",
		"KC_LGUI": "CMD",
		"KC_RGUI": "CMD",
		"KC_LCMD": "CMD",
		"KC_RCMD": "CMD",
		"KC_LSFT": "SHIFT",
		"KC_RSFT": "SHIFT",
		"RGB_TOG": "RGBTO",
		"RGB_MOD": "RGBMO",
		"RGB_HUI": "RGBH",
		"RGB_SAI": "RGBS",
		"RGB_VAI": "RGBV",
		// Advance
		// Space cadet
		"SC_LSPO": "( SHIFT",
		"SC_RSPC": ") SHIFT",
		// Quantum
		"EE_CLR": "EECLR",
	},
}

var mapTransform = []map[string]string{
	{
		// Layer enum
		"LAYER_QWERTY":   "0",
		"LAYER_COMMON":   "1",
		"LAYER_RARELY":   "2",
		"LAYER_RECOVERY": "3",
	},
	{
		// Prefix
		"KC_": "",
		"QK_": "",
	},
	{
		"MO(0)": "L0",
		"MO(1)": "L1",
		"MO(2)": "L2",
		"MO(3)": "L3",
	},
}

// Only use this as last resort
var mapBindingTiny = []map[string]string{
	{
		"BACKSPACE": "BACK",
	},
}

type DrawConfig struct {
	PrintLayout bool
	PrintLayer  bool
}

func Draw(
	layouts map[string]map[string][]QMKKeyDictionary,
	keymap QMKKeymap,
	cfg DrawConfig,
) string {
	layoutsStr := make([]string, 0, len(layouts))

	// Each keyboard has many layouts
	for layout := range layouts {
		if keymap.Layout != layout {
			continue
		}

		keys, ok := layouts[layout]["layout"]
		if !ok {
			continue
		}

		layoutStr := ""
		if cfg.PrintLayout {
			layoutStr += fmt.Sprintf("Layout %s\n", layout)
		}

		// Preprocess keys
		// Y aka row -> X aka col
		var newMaxX, newMaxY int
		for i := range keys {
			if int(keys[i].W) == 0 {
				keys[i].W = 1
			}

			if int(keys[i].H) == 0 {
				keys[i].H = 1
			}

			// Because 0.25
			keys[i].NewX = int(keys[i].X * scaleX)
			keys[i].NewY = int(keys[i].Y * scaleY)
			keys[i].NewW = int(keys[i].W * scaleX)
			keys[i].NewH = int(keys[i].H * scaleY)

			if keys[i].NewX+keys[i].NewW > newMaxX {
				newMaxX = keys[i].NewX + keys[i].NewW
			}

			if keys[i].NewY+keys[i].NewH > newMaxY {
				newMaxY = keys[i].NewY + keys[i].NewH
			}
		}

		if newMaxX == 0 || newMaxY == 0 {
			continue
		}

		// Each keymap has many layers
		layersStr := make([]string, 0, len(keymap.Layers))
		for iLayer, layer := range keymap.Layers {
			// PreProcess table with space
			table := make([][]string, newMaxY+1)
			for i := 0; i <= newMaxY; i++ {
				// Padding 1 in the right
				table[i] = make([]string, newMaxX+1)
				for j := 0; j <= newMaxX; j++ {
					table[i][j] = " "
				}
			}

			// Fill layout
			count := 0
			for _, key := range keys {
				keyStr := layer[count]

				// Convert keyStr
				isRaw := false
				for _, m := range mapRawBinding {
					if _, ok := m[keyStr]; ok {
						isRaw = true
						keyStr = m[keyStr]
						break
					}
				}

				if !isRaw {
					for _, m := range mapTransform {
						for from, to := range m {
							keyStr = strings.ReplaceAll(keyStr, from, to)
						}
					}
				}

				// Padding to center key
				// Why / 2, why - 1 ?
				// Base on my feeling of course
				padding := (key.NewW-len(keyStr))/2 - 1
				if padding <= 0 {
					padding = 1
				}

				if len(keyStr)+2*padding > key.NewW {
					// Make it smaller
					for _, m := range mapBindingTiny {
						if _, ok := m[keyStr]; ok {
							keyStr = m[keyStr]
							break
						}
					}

					// Re calc padding
					padding = (key.NewW-len(keyStr))/2 - 1
					if padding <= 0 {
						padding = 1
					}
				}

				// Draw strategy
				// Draw 4 + in corner
				// Draw - on top and bottom
				// Draw | on left and right
				for i := key.NewY; i <= key.NewY+key.NewH; i++ {
					for j := key.NewX; j <= key.NewX+key.NewW; j++ {
						if i == key.NewY || i == key.NewY+key.NewH {
							if j == key.NewX || j == key.NewX+key.NewW {
								// Draw corner
								table[i][j] = "+"
							} else if table[i][j] != "+" {
								// Draw top/bottom
								table[i][j] = "-"
							}
						} else if i == key.NewY+key.NewH/2 {
							// Write key in the middle
							if j == key.NewX || j == key.NewX+key.NewW {
								// Draw left/right
								table[i][j] = "|"
							} else if len(keyStr) > 0 && j > key.NewX+padding && j < key.NewX+len(keyStr)+padding+1 && j <= key.NewX+key.NewW-padding {
								// Only handle ASCII keyStr
								table[i][j] = string(keyStr[j-key.NewX-padding-1])
							}
						} else {
							if j == key.NewX || j == key.NewX+key.NewW {
								table[i][j] = "|"
							}
						}
					}
				}

				count++
			}

			// Print
			layerStr := ""
			if cfg.PrintLayer {
				layerStr += fmt.Sprintf("Layer %d\n", iLayer)
			}

			for i := range table {
				for j := range table[i] {
					layerStr += table[i][j]
				}
				layerStr += "\n"
			}
			layersStr = append(layersStr, layerStr)
		}

		// Print
		layoutStr += strings.Join(layersStr, "\n")
		layoutsStr = append(layoutsStr, layoutStr)
	}

	return strings.Join(layoutsStr, "\n")
}
