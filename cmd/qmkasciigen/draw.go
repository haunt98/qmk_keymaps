package main

import (
	"fmt"
	"strings"
)

const (
	scaleX = 8
	scaleY = 2
)

var mapRawBinding = []map[string]string{
	{
		// My custom
		"CTL_ESC": "ESC CTRL",
	},
}

// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
var mapTransform = []map[string]string{
	{
		"MO(0)": "L0",
		"MO(1)": "L1",
		"MO(2)": "L2",
		"MO(3)": "L3",
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
		"KC_CAPS": "CAPSLOCK",
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
		"UG_TOGG": "UGTO",
		"UG_NEXT": "UGNEXT",
		"UG_HUEU": "UGH",
		"UG_SATU": "UGS",
		"UG_VALU": "UGV",
		"UG_SPDU": "UGSPD",
		"RGB_TOG": "RGBTO",
		"RGB_MOD": "RGBMO",
		"RGB_HUI": "RGBH",
		"RGB_SAI": "RGBS",
		"RGB_VAI": "RGBV",
		"KC_LCBR": "{",
		"KC_RCBR": "}",
		"KC_UNDS": "_",
		"KC_PLUS": "+",
		"KC_PIPE": "|",
		"KC_COLN": ":",
		"KC_DQUO": `"`,
		"KC_LT":   "<",
		"KC_GT":   ">",
		// Advance
		// Mod-Tap
		"LCTL_T": "CTRL",
		// Space cadet
		"SC_LSPO": "( SHIFT",
		"SC_RSPC": ") SHIFT",
		// Quantum
		"EE_CLR": "EECLR",
	},
	{
		// Mod-Tap
		"CTL_T": "CTRL",
	},
	{
		// Prefix
		"KC_": "",
		"QK_": "",
	},
}

// Only use this as last resort
var mapTransformTiny = []map[string]string{
	{
		"BACKSPACE": "BACK",
		"CAPSLOCK":  "CAPS",
		"ESC CTRL":  "ESCTL",
		"( SHIFT":   "( SFT",
		") SHIFT":   "( SFT",
	},
}

type DrawConfig struct {
	AllowLayers map[int]struct{}
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

		layoutStr := drawLayout(
			keymap,
			layout,
			keys,
			cfg,
		)

		layoutsStr = append(layoutsStr, layoutStr)
	}

	return strings.Join(layoutsStr, "\n")
}

func drawLayout(
	keymap QMKKeymap,
	layout string,
	keys []QMKKeyDictionary,
	cfg DrawConfig,
) string {
	layoutStr := ""
	if cfg.PrintLayout {
		layoutStr += fmt.Sprintf("Layout %s\n", layout)
	}

	// Preprocess keys
	maxX, maxY := findMaxXY(keys)

	// Each keymap layout has many layers
	layersStr := make([]string, 0, len(keymap.Layers))
	for layerIndex, layer := range keymap.Layers {
		layerStr := drawLayer(
			keys,
			maxX,
			maxY,
			layerIndex,
			layer,
			cfg,
		)

		layersStr = append(layersStr, layerStr)
	}

	// Print
	layoutStr += strings.Join(layersStr, "\n")

	return layoutStr
}

func drawLayer(
	keys []QMKKeyDictionary,
	maxX int,
	maxY int,
	layerIndex int,
	layer []string,
	cfg DrawConfig,
) string {
	// Only check if cfg valid
	if len(cfg.AllowLayers) > 0 {
		if _, ok := cfg.AllowLayers[layerIndex]; !ok {
			// Skip if not in list
			return ""
		}
	}

	// Preprocess table with space
	table := make([][]string, maxY+1)
	for i := 0; i <= maxY; i++ {
		// Padding 1 in the right
		table[i] = make([]string, maxX+1)
		for j := 0; j <= maxX; j++ {
			table[i][j] = cp437Empty
		}
	}

	// Fill layout key by key
	// layer and keys is same len
	for i, key := range keys {
		keyStr, padding := getKeyStr(key, layer[i])

		drawKey(
			table,
			key,
			keyStr,
			padding,
		)
	}

	// Print
	layerStr := ""
	if cfg.PrintLayer {
		layerStr += fmt.Sprintf("Layer %d\n", layerIndex)
	}

	for i := range table {
		for j := range table[i] {
			layerStr += table[i][j]
		}
		layerStr += "\n"
	}

	return layerStr
}

func drawKey(
	table [][]string,
	key QMKKeyDictionary,
	keyStr string,
	padding int,
) {
	// Draw strategy
	// Draw 4 + in corner
	// Draw - on top and bottom
	// Draw | on left and right
	for i := key.NewY; i <= key.NewY+key.NewH; i++ {
		for j := key.NewX; j <= key.NewX+key.NewW; j++ {
			switch i {
			case key.NewY, key.NewY + key.NewH:
				if j == key.NewX || j == key.NewX+key.NewW {
					// Draw corner
					var temp string
					if i == key.NewY {
						if j == key.NewX {
							// Corner top left
							temp = cp437BottomVerticalRight
						} else {
							// Corner top right
							temp = cp437BottomVerticalLeft
						}
					} else {
						if j == key.NewX {
							// Corner bottom left
							temp = cp437TopVerticalRight
						} else {
							// Corner top right
							temp = cp437TopVerticalLeft
						}
					}

					// Need to combine with current
					table[i][j] = cp437Plus(table[i][j], temp)
				} else {
					// Draw horizontal
					table[i][j] = cp437Plus(table[i][j], cp437Horizontal)
				}
			case key.NewY + key.NewH/2:
				// Write key in the middle
				if j == key.NewX || j == key.NewX+key.NewW {
					// Draw vertical most left/right
					table[i][j] = cp437Plus(table[i][j], cp437Vertical)
				} else if keyStr != "" &&
					j > key.NewX+padding && j < key.NewX+len(keyStr)+padding+1 && j <= key.NewX+key.NewW-padding {
					// Only handle ASCII keyStr
					table[i][j] = string(keyStr[j-key.NewX-padding-1])
				}
			default:
				// Draw vertical most left/right
				if j == key.NewX || j == key.NewX+key.NewW {
					table[i][j] = cp437Plus(table[i][j], cp437Vertical)
				}
			}
		}
	}
}

// Y aka row -> X aka col
// Try to find the furthest corner to draw rectangle which covers layout
func findMaxXY(keys []QMKKeyDictionary) (int, int) {
	var maxX, maxY int
	for i := range keys {
		if int(keys[i].W) == 0 {
			keys[i].W = 1
		}

		if int(keys[i].H) == 0 {
			keys[i].H = 1
		}

		// TODO: Move scale to config
		// Because 0.25
		keys[i].NewX = int(keys[i].X * scaleX)
		keys[i].NewY = int(keys[i].Y * scaleY)
		keys[i].NewW = int(keys[i].W * scaleX)
		keys[i].NewH = int(keys[i].H * scaleY)

		if keys[i].NewX+keys[i].NewW > maxX {
			maxX = keys[i].NewX + keys[i].NewW
		}

		if keys[i].NewY+keys[i].NewH > maxY {
			maxY = keys[i].NewY + keys[i].NewH
		}
	}

	return maxX, maxY
}

// Return key str and padding
func getKeyStr(key QMKKeyDictionary, rawKeyStr string) (string, int) {
	keyStr := rawKeyStr

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

		// Handle layer
		if strings.HasPrefix(keyStr, "MO(") && strings.HasSuffix(keyStr, ")") {
			keyStr = strings.TrimPrefix(keyStr, "MO(")
			keyStr = strings.TrimSuffix(keyStr, ")")
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
		for _, m := range mapTransformTiny {
			for from, to := range m {
				keyStr = strings.ReplaceAll(keyStr, from, to)
			}
		}

		// Re calc padding
		padding = (key.NewW-len(keyStr))/2 - 1
		if padding <= 0 {
			padding = 1
		}
	}

	return keyStr, padding
}
