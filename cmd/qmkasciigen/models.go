package main

const (
	DefaultQMKW = 1
	DefaultQMKH = 1
)

// https://github.com/qmk/qmk_firmware/blob/master/docs/reference_info_json.md
type QMKInfo struct {
	Layouts map[string]map[string][]QMKKeyDictionary `json:"layouts"`
}

type QMKKeyDictionary struct {
	X    float32 `json:"x"`
	Y    float32 `json:"y"`
	W    float32 `json:"w"`
	H    float32 `json:"h"`
	NewX int     `json:"-"`
	NewY int     `json:"-"`
	NewW int     `json:"-"`
	NewH int     `json:"-"`
}

// https://config.qmk.fm/#/
type QMKKeymap struct {
	Layout string     `json:"layout"`
	Layers [][]string `json:"layers"`
}
