package main

import (
	"strings"
)

const (
	scaleX = 8
	scaleY = 3
)

// https://github.com/qmk/qmk_firmware/blob/master/docs/keycodes.md
var mapSpecialKey = map[string]string{
	"KC_TRNS": " ",
	"KC_NO":   " ",
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
	"KC_LALT": "ALT",
	"KC_RALT": "ALT",
	"KC_LGUI": "CMD",
	"KC_RGUI": "CMD",
	"KC_LSFT": "SHIFT",
	"KC_RSFT": "SHIFT",
	"RGB_TOG": "RGBTO",
	"RGB_MOD": "RGBMO",
	"RGB_HUI": "RGBH",
	"RGB_SAI": "RGBS",
	"RGB_VAI": "RGBV",
	// Custom
	"CTL_T(KC_ESC)": "CTRL ESC",
}

func Draw(
	layouts map[string]map[string][]QMKKeyDictionary,
	keymap QMKKeymap,
) string {
	result := ""

	for layout := range layouts {
		if keymap.Layout != layout {
			continue
		}

		keys, ok := layouts[layout]["layout"]
		if !ok {
			continue
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

		for _, layer := range keymap.Layers {
			// Preprocess table
			table := make([][]string, newMaxY)
			for i := 0; i < newMaxY; i++ {
				table[i] = make([]string, newMaxX)
			}

			// Fill layout
			count := 0
			for _, key := range keys {
				keyStr := layer[count]

				// Process keyStr
				if newKeyStr, ok := mapSpecialKey[keyStr]; ok {
					keyStr = newKeyStr
				} else if strings.HasPrefix(keyStr, "KC_") {
					keyStr = strings.TrimPrefix(keyStr, "KC_")
				} else if strings.HasPrefix(keyStr, "QK_") {
					keyStr = strings.TrimPrefix(keyStr, "QK_")
				}

				// Padding to center key
				padding := (key.NewW - len(keyStr)) / 2
				if padding <= 0 {
					padding = 1
				}

				for i := key.NewY; i < key.NewY+key.NewH; i++ {
					for j := key.NewX; j < key.NewX+key.NewW; j++ {
						if i == key.NewY {
							if j == key.NewX {
								table[i][j] = "+"
							} else {
								table[i][j] = "-"
							}
						} else if i == key.NewY+key.NewH/2 {
							// Write key in the middle
							if j == key.NewX {
								table[i][j] = "|"
							} else if j > key.NewX+padding && j < key.NewX+len(keyStr)+padding+1 && j <= key.NewX+key.NewW-padding {
								table[i][j] = string(keyStr[j-key.NewX-padding-1])
							} else {
								table[i][j] = " "
							}
						} else {
							table[i][j] = " "
						}
					}
				}

				count++
			}

			// Postprecess table
			newTable := make([][]string, 0, newMaxY+1)

			// Remove empty row
			for i := 0; i < len(table); i++ {
				isEmptyRow := true
				for j := 0; j < len(table[i]); j++ {
					if table[i][j] != " " {
						isEmptyRow = false
						break
					}
				}

				if isEmptyRow {
					continue
				}

				// Padding
				paddingRight := "|"
				if table[i][len(table[i])-1] == "-" {
					paddingRight = "+"
				}

				newTable = append(newTable, append(table[i], paddingRight))
			}

			// Padding
			paddingRow := make([]string, 0, newMaxX)
			for j := 0; j < newMaxX; j++ {
				paddingBottom := "-"
				if newTable[len(newTable)-1][j] == "|" {
					paddingBottom = "+"
				}
				paddingRow = append(paddingRow, paddingBottom)
			}
			newTable = append(newTable, append(paddingRow, "+"))

			str := ""
			for i := range newTable {
				for j := range newTable[i] {
					str += newTable[i][j]
				}
				str += "\n"
			}

			result += str + "\n"
		}

	}

	return result
}
