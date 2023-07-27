package handler



func Preparation(placesMap map[string][]string, types string) (string, string) {
	var (
		respkey string
		resp    string
		nol     string
	)

	keyMapping := map[string][]string{
		"VIP":      {"01", "10"},
		"economy":  {"02", "03", "04", "05", "06", "07"},
		"business": {"08", "09"},
	}

	keys, ok := keyMapping[types]
	if !ok {
		return "", ""
	}

	for _, key := range keys {
		if value, ok := placesMap[key]; ok {
			for _, val := range value {
				for j := 0; j < 3; j++ {
					if string(val[2]) == "0" {
						nol = "0"
					}
					if string(val[j]) != "0" {
						resp = resp + string(val[j])
					}
				}
				respkey = key
				resp = resp + nol + "-" + resp + nol
				break
			}
			return resp, respkey
		}
	}

	return "", ""
}




func Preparation2(placesMap map[string][]string) (string, string) {
	var (
		respkey string
		resp    string
		nol     string
	)

	keys := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}

	for _, key := range keys {
		if value, ok := placesMap[key]; ok {
			for _, val := range value {
				for j := 0; j < 3; j++ {
					if string(val[2]) == "0" {
						nol = "0"
					}
					if string(val[j]) != "0" {
						resp = resp + string(val[j])
					}
				}
				respkey = key
				resp = resp + nol + "-" + resp + nol
				return resp, respkey
			}
		}
	}

	return "", ""
}
