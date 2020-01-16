package emoji

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tmdvs/Go-Emoji-Utils/utils"
)

// Emoji - Struct representing Emoji
type Emoji struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	Descriptor string `json:"descriptor"`
}

// Emojis - Map of Emoji Runes as Hex keys to their description
var Emojis map[string]Emoji

// Unmarshal the emoji JSON into the Emojis map
func init() {
	byteValue := []byte(`{
	"1F004": {
		"key": "1F004",
		"value": "🀄",
		"descriptor": "Mahjong Red Dragon"
	},
	"1F0CF": {
		"key": "1F0CF",
		"value": "🃏",
		"descriptor": "Joker"
	},
	"1F170-FE0F": {
		"key": "1F170-FE0F",
		"value": "🅰️",
		"descriptor": "A Button (Blood Type)"
	},
	"1F171-FE0F": {
		"key": "1F171-FE0F",
		"value": "🅱️",
		"descriptor": "B Button (Blood Type)"
	},
	"1F17E-FE0F": {
		"key": "1F17E-FE0F",
		"value": "🅾️",
		"descriptor": "O Button (Blood Type)"
	},
	"1F17F-FE0F": {
		"key": "1F17F-FE0F",
		"value": "🅿️",
		"descriptor": "P Button"
	},
	"1F18E": {
		"key": "1F18E",
		"value": "🆎",
		"descriptor": "AB Button (Blood Type)"
	},
	"1F191": {
		"key": "1F191",
		"value": "🆑",
		"descriptor": "CL Button"
	},
	"1F192": {
		"key": "1F192",
		"value": "🆒",
		"descriptor": "Cool Button"
	},
	"1F193": {
		"key": "1F193",
		"value": "🆓",
		"descriptor": "Free Button"
	},
	"1F194": {
		"key": "1F194",
		"value": "🆔",
		"descriptor": "ID Button"
	},
	"1F195": {
		"key": "1F195",
		"value": "🆕",
		"descriptor": "New Button"
	},
	"1F196": {
		"key": "1F196",
		"value": "🆖",
		"descriptor": "NG Button"
	},
	"1F197": {
		"key": "1F197",
		"value": "🆗",
		"descriptor": "OK Button"
	},
	"1F198": {
		"key": "1F198",
		"value": "🆘",
		"descriptor": "SOS Button"
	},
	"1F199": {
		"key": "1F199",
		"value": "🆙",
		"descriptor": "Up! Button"
	},
	"1F19A": {
		"key": "1F19A",
		"value": "🆚",
		"descriptor": "Vs Button"
	},
	"1F1E6-1F1E8": {
		"key": "1F1E6-1F1E8",
		"value": "🇦🇨",
		"descriptor": "Flag: Ascension Island"
	},
	"1F1E6-1F1E9": {
		"key": "1F1E6-1F1E9",
		"value": "🇦🇩",
		"descriptor": "Flag: Andorra"
	},
	"1F1E6-1F1EA": {
		"key": "1F1E6-1F1EA",
		"value": "🇦🇪",
		"descriptor": "Flag: United Arab Emirates"
	},
	"1F1E6-1F1EB": {
		"key": "1F1E6-1F1EB",
		"value": "🇦🇫",
		"descriptor": "Flag: Afghanistan"
	},
	"1F1E6-1F1EC": {
		"key": "1F1E6-1F1EC",
		"value": "🇦🇬",
		"descriptor": "Flag: Antigua \u0026 Barbuda"
	},
	"1F1E6-1F1EE": {
		"key": "1F1E6-1F1EE",
		"value": "🇦🇮",
		"descriptor": "Flag: Anguilla"
	},
	"1F1E6-1F1F1": {
		"key": "1F1E6-1F1F1",
		"value": "🇦🇱",
		"descriptor": "Flag: Albania"
	},
	"1F1E6-1F1F2": {
		"key": "1F1E6-1F1F2",
		"value": "🇦🇲",
		"descriptor": "Flag: Armenia"
	},
	"1F1E6-1F1F4": {
		"key": "1F1E6-1F1F4",
		"value": "🇦🇴",
		"descriptor": "Flag: Angola"
	},
	"1F1E6-1F1F6": {
		"key": "1F1E6-1F1F6",
		"value": "🇦🇶",
		"descriptor": "Flag: Antarctica"
	},
	"1F1E6-1F1F7": {
		"key": "1F1E6-1F1F7",
		"value": "🇦🇷",
		"descriptor": "Flag: Argentina"
	},
	"1F1E6-1F1F8": {
		"key": "1F1E6-1F1F8",
		"value": "🇦🇸",
		"descriptor": "Flag: American Samoa"
	},
	"1F1E6-1F1F9": {
		"key": "1F1E6-1F1F9",
		"value": "🇦🇹",
		"descriptor": "Flag: Austria"
	},
	"1F1E6-1F1FA": {
		"key": "1F1E6-1F1FA",
		"value": "🇦🇺",
		"descriptor": "Flag: Australia"
	},
	"1F1E6-1F1FC": {
		"key": "1F1E6-1F1FC",
		"value": "🇦🇼",
		"descriptor": "Flag: Aruba"
	},
	"1F1E6-1F1FD": {
		"key": "1F1E6-1F1FD",
		"value": "🇦🇽",
		"descriptor": "Flag: Åland Islands"
	},
	"1F1E6-1F1FF": {
		"key": "1F1E6-1F1FF",
		"value": "🇦🇿",
		"descriptor": "Flag: Azerbaijan"
	},
	"1F1E7-1F1E6": {
		"key": "1F1E7-1F1E6",
		"value": "🇧🇦",
		"descriptor": "Flag: Bosnia \u0026 Herzegovina"
	},
	"1F1E7-1F1E7": {
		"key": "1F1E7-1F1E7",
		"value": "🇧🇧",
		"descriptor": "Flag: Barbados"
	},
	"1F1E7-1F1E9": {
		"key": "1F1E7-1F1E9",
		"value": "🇧🇩",
		"descriptor": "Flag: Bangladesh"
	},
	"1F1E7-1F1EA": {
		"key": "1F1E7-1F1EA",
		"value": "🇧🇪",
		"descriptor": "Flag: Belgium"
	},
	"1F1E7-1F1EB": {
		"key": "1F1E7-1F1EB",
		"value": "🇧🇫",
		"descriptor": "Flag: Burkina Faso"
	},
	"1F1E7-1F1EC": {
		"key": "1F1E7-1F1EC",
		"value": "🇧🇬",
		"descriptor": "Flag: Bulgaria"
	},
	"1F1E7-1F1ED": {
		"key": "1F1E7-1F1ED",
		"value": "🇧🇭",
		"descriptor": "Flag: Bahrain"
	},
	"1F1E7-1F1EE": {
		"key": "1F1E7-1F1EE",
		"value": "🇧🇮",
		"descriptor": "Flag: Burundi"
	},
	"1F1E7-1F1EF": {
		"key": "1F1E7-1F1EF",
		"value": "🇧🇯",
		"descriptor": "Flag: Benin"
	},
	"1F1E7-1F1F1": {
		"key": "1F1E7-1F1F1",
		"value": "🇧🇱",
		"descriptor": "Flag: St. Barthélemy"
	},
	"1F1E7-1F1F2": {
		"key": "1F1E7-1F1F2",
		"value": "🇧🇲",
		"descriptor": "Flag: Bermuda"
	},
	"1F1E7-1F1F3": {
		"key": "1F1E7-1F1F3",
		"value": "🇧🇳",
		"descriptor": "Flag: Brunei"
	},
	"1F1E7-1F1F4": {
		"key": "1F1E7-1F1F4",
		"value": "🇧🇴",
		"descriptor": "Flag: Bolivia"
	},
	"1F1E7-1F1F6": {
		"key": "1F1E7-1F1F6",
		"value": "🇧🇶",
		"descriptor": "Flag: Caribbean Netherlands"
	},
	"1F1E7-1F1F7": {
		"key": "1F1E7-1F1F7",
		"value": "🇧🇷",
		"descriptor": "Flag: Brazil"
	},
	"1F1E7-1F1F8": {
		"key": "1F1E7-1F1F8",
		"value": "🇧🇸",
		"descriptor": "Flag: Bahamas"
	},
	"1F1E7-1F1F9": {
		"key": "1F1E7-1F1F9",
		"value": "🇧🇹",
		"descriptor": "Flag: Bhutan"
	},
	"1F1E7-1F1FB": {
		"key": "1F1E7-1F1FB",
		"value": "🇧🇻",
		"descriptor": "Flag: Bouvet Island"
	},
	"1F1E7-1F1FC": {
		"key": "1F1E7-1F1FC",
		"value": "🇧🇼",
		"descriptor": "Flag: Botswana"
	},
	"1F1E7-1F1FE": {
		"key": "1F1E7-1F1FE",
		"value": "🇧🇾",
		"descriptor": "Flag: Belarus"
	},
	"1F1E7-1F1FF": {
		"key": "1F1E7-1F1FF",
		"value": "🇧🇿",
		"descriptor": "Flag: Belize"
	},
	"1F1E8-1F1E6": {
		"key": "1F1E8-1F1E6",
		"value": "🇨🇦",
		"descriptor": "Flag: Canada"
	},
	"1F1E8-1F1E8": {
		"key": "1F1E8-1F1E8",
		"value": "🇨🇨",
		"descriptor": "Flag: Cocos (Keeling) Islands"
	},
	"1F1E8-1F1E9": {
		"key": "1F1E8-1F1E9",
		"value": "🇨🇩",
		"descriptor": "Flag: Congo - Kinshasa"
	},
	"1F1E8-1F1EB": {
		"key": "1F1E8-1F1EB",
		"value": "🇨🇫",
		"descriptor": "Flag: Central African Republic"
	},
	"1F1E8-1F1EC": {
		"key": "1F1E8-1F1EC",
		"value": "🇨🇬",
		"descriptor": "Flag: Congo - Brazzaville"
	},
	"1F1E8-1F1ED": {
		"key": "1F1E8-1F1ED",
		"value": "🇨🇭",
		"descriptor": "Flag: Switzerland"
	},
	"1F1E8-1F1EE": {
		"key": "1F1E8-1F1EE",
		"value": "🇨🇮",
		"descriptor": "Flag: Côte d’Ivoire"
	},
	"1F1E8-1F1F0": {
		"key": "1F1E8-1F1F0",
		"value": "🇨🇰",
		"descriptor": "Flag: Cook Islands"
	},
	"1F1E8-1F1F1": {
		"key": "1F1E8-1F1F1",
		"value": "🇨🇱",
		"descriptor": "Flag: Chile"
	},
	"1F1E8-1F1F2": {
		"key": "1F1E8-1F1F2",
		"value": "🇨🇲",
		"descriptor": "Flag: Cameroon"
	},
	"1F1E8-1F1F3": {
		"key": "1F1E8-1F1F3",
		"value": "🇨🇳",
		"descriptor": "Flag: China"
	},
	"1F1E8-1F1F4": {
		"key": "1F1E8-1F1F4",
		"value": "🇨🇴",
		"descriptor": "Flag: Colombia"
	},
	"1F1E8-1F1F5": {
		"key": "1F1E8-1F1F5",
		"value": "🇨🇵",
		"descriptor": "Flag: Clipperton Island"
	},
	"1F1E8-1F1F7": {
		"key": "1F1E8-1F1F7",
		"value": "🇨🇷",
		"descriptor": "Flag: Costa Rica"
	},
	"1F1E8-1F1FA": {
		"key": "1F1E8-1F1FA",
		"value": "🇨🇺",
		"descriptor": "Flag: Cuba"
	},
	"1F1E8-1F1FB": {
		"key": "1F1E8-1F1FB",
		"value": "🇨🇻",
		"descriptor": "Flag: Cape Verde"
	},
	"1F1E8-1F1FC": {
		"key": "1F1E8-1F1FC",
		"value": "🇨🇼",
		"descriptor": "Flag: Curaçao"
	},
	"1F1E8-1F1FD": {
		"key": "1F1E8-1F1FD",
		"value": "🇨🇽",
		"descriptor": "Flag: Christmas Island"
	},
	"1F1E8-1F1FE": {
		"key": "1F1E8-1F1FE",
		"value": "🇨🇾",
		"descriptor": "Flag: Cyprus"
	},
	"1F1E8-1F1FF": {
		"key": "1F1E8-1F1FF",
		"value": "🇨🇿",
		"descriptor": "Flag: Czechia"
	},
	"1F1E9-1F1EA": {
		"key": "1F1E9-1F1EA",
		"value": "🇩🇪",
		"descriptor": "Flag: Germany"
	},
	"1F1E9-1F1EC": {
		"key": "1F1E9-1F1EC",
		"value": "🇩🇬",
		"descriptor": "Flag: Diego Garcia"
	},
	"1F1E9-1F1EF": {
		"key": "1F1E9-1F1EF",
		"value": "🇩🇯",
		"descriptor": "Flag: Djibouti"
	},
	"1F1E9-1F1F0": {
		"key": "1F1E9-1F1F0",
		"value": "🇩🇰",
		"descriptor": "Flag: Denmark"
	},
	"1F1E9-1F1F2": {
		"key": "1F1E9-1F1F2",
		"value": "🇩🇲",
		"descriptor": "Flag: Dominica"
	},
	"1F1E9-1F1F4": {
		"key": "1F1E9-1F1F4",
		"value": "🇩🇴",
		"descriptor": "Flag: Dominican Republic"
	},
	"1F1E9-1F1FF": {
		"key": "1F1E9-1F1FF",
		"value": "🇩🇿",
		"descriptor": "Flag: Algeria"
	},
	"1F1EA-1F1E6": {
		"key": "1F1EA-1F1E6",
		"value": "🇪🇦",
		"descriptor": "Flag: Ceuta \u0026 Melilla"
	},
	"1F1EA-1F1E8": {
		"key": "1F1EA-1F1E8",
		"value": "🇪🇨",
		"descriptor": "Flag: Ecuador"
	},
	"1F1EA-1F1EA": {
		"key": "1F1EA-1F1EA",
		"value": "🇪🇪",
		"descriptor": "Flag: Estonia"
	},
	"1F1EA-1F1EC": {
		"key": "1F1EA-1F1EC",
		"value": "🇪🇬",
		"descriptor": "Flag: Egypt"
	},
	"1F1EA-1F1ED": {
		"key": "1F1EA-1F1ED",
		"value": "🇪🇭",
		"descriptor": "Flag: Western Sahara"
	},
	"1F1EA-1F1F7": {
		"key": "1F1EA-1F1F7",
		"value": "🇪🇷",
		"descriptor": "Flag: Eritrea"
	},
	"1F1EA-1F1F8": {
		"key": "1F1EA-1F1F8",
		"value": "🇪🇸",
		"descriptor": "Flag: Spain"
	},
	"1F1EA-1F1F9": {
		"key": "1F1EA-1F1F9",
		"value": "🇪🇹",
		"descriptor": "Flag: Ethiopia"
	},
	"1F1EA-1F1FA": {
		"key": "1F1EA-1F1FA",
		"value": "🇪🇺",
		"descriptor": "Flag: European Union"
	},
	"1F1EB-1F1EE": {
		"key": "1F1EB-1F1EE",
		"value": "🇫🇮",
		"descriptor": "Flag: Finland"
	},
	"1F1EB-1F1EF": {
		"key": "1F1EB-1F1EF",
		"value": "🇫🇯",
		"descriptor": "Flag: Fiji"
	},
	"1F1EB-1F1F0": {
		"key": "1F1EB-1F1F0",
		"value": "🇫🇰",
		"descriptor": "Flag: Falkland Islands"
	},
	"1F1EB-1F1F2": {
		"key": "1F1EB-1F1F2",
		"value": "🇫🇲",
		"descriptor": "Flag: Micronesia"
	},
	"1F1EB-1F1F4": {
		"key": "1F1EB-1F1F4",
		"value": "🇫🇴",
		"descriptor": "Flag: Faroe Islands"
	},
	"1F1EB-1F1F7": {
		"key": "1F1EB-1F1F7",
		"value": "🇫🇷",
		"descriptor": "Flag: France"
	},
	"1F1EC-1F1E6": {
		"key": "1F1EC-1F1E6",
		"value": "🇬🇦",
		"descriptor": "Flag: Gabon"
	},
	"1F1EC-1F1E7": {
		"key": "1F1EC-1F1E7",
		"value": "🇬🇧",
		"descriptor": "Flag: United Kingdom"
	},
	"1F1EC-1F1E9": {
		"key": "1F1EC-1F1E9",
		"value": "🇬🇩",
		"descriptor": "Flag: Grenada"
	},
	"1F1EC-1F1EA": {
		"key": "1F1EC-1F1EA",
		"value": "🇬🇪",
		"descriptor": "Flag: Georgia"
	},
	"1F1EC-1F1EB": {
		"key": "1F1EC-1F1EB",
		"value": "🇬🇫",
		"descriptor": "Flag: French Guiana"
	},
	"1F1EC-1F1EC": {
		"key": "1F1EC-1F1EC",
		"value": "🇬🇬",
		"descriptor": "Flag: Guernsey"
	},
	"1F1EC-1F1ED": {
		"key": "1F1EC-1F1ED",
		"value": "🇬🇭",
		"descriptor": "Flag: Ghana"
	},
	"1F1EC-1F1EE": {
		"key": "1F1EC-1F1EE",
		"value": "🇬🇮",
		"descriptor": "Flag: Gibraltar"
	},
	"1F1EC-1F1F1": {
		"key": "1F1EC-1F1F1",
		"value": "🇬🇱",
		"descriptor": "Flag: Greenland"
	},
	"1F1EC-1F1F2": {
		"key": "1F1EC-1F1F2",
		"value": "🇬🇲",
		"descriptor": "Flag: Gambia"
	},
	"1F1EC-1F1F3": {
		"key": "1F1EC-1F1F3",
		"value": "🇬🇳",
		"descriptor": "Flag: Guinea"
	},
	"1F1EC-1F1F5": {
		"key": "1F1EC-1F1F5",
		"value": "🇬🇵",
		"descriptor": "Flag: Guadeloupe"
	},
	"1F1EC-1F1F6": {
		"key": "1F1EC-1F1F6",
		"value": "🇬🇶",
		"descriptor": "Flag: Equatorial Guinea"
	},
	"1F1EC-1F1F7": {
		"key": "1F1EC-1F1F7",
		"value": "🇬🇷",
		"descriptor": "Flag: Greece"
	},
	"1F1EC-1F1F8": {
		"key": "1F1EC-1F1F8",
		"value": "🇬🇸",
		"descriptor": "Flag: South Georgia \u0026 South Sandwich Islands"
	},
	"1F1EC-1F1F9": {
		"key": "1F1EC-1F1F9",
		"value": "🇬🇹",
		"descriptor": "Flag: Guatemala"
	},
	"1F1EC-1F1FA": {
		"key": "1F1EC-1F1FA",
		"value": "🇬🇺",
		"descriptor": "Flag: Guam"
	},
	"1F1EC-1F1FC": {
		"key": "1F1EC-1F1FC",
		"value": "🇬🇼",
		"descriptor": "Flag: Guinea-Bissau"
	},
	"1F1EC-1F1FE": {
		"key": "1F1EC-1F1FE",
		"value": "🇬🇾",
		"descriptor": "Flag: Guyana"
	},
	"1F1ED-1F1F0": {
		"key": "1F1ED-1F1F0",
		"value": "🇭🇰",
		"descriptor": "Flag: Hong Kong SAR China"
	},
	"1F1ED-1F1F2": {
		"key": "1F1ED-1F1F2",
		"value": "🇭🇲",
		"descriptor": "Flag: Heard \u0026 McDonald Islands"
	},
	"1F1ED-1F1F3": {
		"key": "1F1ED-1F1F3",
		"value": "🇭🇳",
		"descriptor": "Flag: Honduras"
	},
	"1F1ED-1F1F7": {
		"key": "1F1ED-1F1F7",
		"value": "🇭🇷",
		"descriptor": "Flag: Croatia"
	},
	"1F1ED-1F1F9": {
		"key": "1F1ED-1F1F9",
		"value": "🇭🇹",
		"descriptor": "Flag: Haiti"
	},
	"1F1ED-1F1FA": {
		"key": "1F1ED-1F1FA",
		"value": "🇭🇺",
		"descriptor": "Flag: Hungary"
	},
	"1F1EE-1F1E8": {
		"key": "1F1EE-1F1E8",
		"value": "🇮🇨",
		"descriptor": "Flag: Canary Islands"
	},
	"1F1EE-1F1E9": {
		"key": "1F1EE-1F1E9",
		"value": "🇮🇩",
		"descriptor": "Flag: Indonesia"
	},
	"1F1EE-1F1EA": {
		"key": "1F1EE-1F1EA",
		"value": "🇮🇪",
		"descriptor": "Flag: Ireland"
	},
	"1F1EE-1F1F1": {
		"key": "1F1EE-1F1F1",
		"value": "🇮🇱",
		"descriptor": "Flag: Israel"
	},
	"1F1EE-1F1F2": {
		"key": "1F1EE-1F1F2",
		"value": "🇮🇲",
		"descriptor": "Flag: Isle of Man"
	},
	"1F1EE-1F1F3": {
		"key": "1F1EE-1F1F3",
		"value": "🇮🇳",
		"descriptor": "Flag: India"
	},
	"1F1EE-1F1F4": {
		"key": "1F1EE-1F1F4",
		"value": "🇮🇴",
		"descriptor": "Flag: British Indian Ocean Territory"
	},
	"1F1EE-1F1F6": {
		"key": "1F1EE-1F1F6",
		"value": "🇮🇶",
		"descriptor": "Flag: Iraq"
	},
	"1F1EE-1F1F7": {
		"key": "1F1EE-1F1F7",
		"value": "🇮🇷",
		"descriptor": "Flag: Iran"
	},
	"1F1EE-1F1F8": {
		"key": "1F1EE-1F1F8",
		"value": "🇮🇸",
		"descriptor": "Flag: Iceland"
	},
	"1F1EE-1F1F9": {
		"key": "1F1EE-1F1F9",
		"value": "🇮🇹",
		"descriptor": "Flag: Italy"
	},
	"1F1EF-1F1EA": {
		"key": "1F1EF-1F1EA",
		"value": "🇯🇪",
		"descriptor": "Flag: Jersey"
	},
	"1F1EF-1F1F2": {
		"key": "1F1EF-1F1F2",
		"value": "🇯🇲",
		"descriptor": "Flag: Jamaica"
	},
	"1F1EF-1F1F4": {
		"key": "1F1EF-1F1F4",
		"value": "🇯🇴",
		"descriptor": "Flag: Jordan"
	},
	"1F1EF-1F1F5": {
		"key": "1F1EF-1F1F5",
		"value": "🇯🇵",
		"descriptor": "Flag: Japan"
	},
	"1F1F0-1F1EA": {
		"key": "1F1F0-1F1EA",
		"value": "🇰🇪",
		"descriptor": "Flag: Kenya"
	},
	"1F1F0-1F1EC": {
		"key": "1F1F0-1F1EC",
		"value": "🇰🇬",
		"descriptor": "Flag: Kyrgyzstan"
	},
	"1F1F0-1F1ED": {
		"key": "1F1F0-1F1ED",
		"value": "🇰🇭",
		"descriptor": "Flag: Cambodia"
	},
	"1F1F0-1F1EE": {
		"key": "1F1F0-1F1EE",
		"value": "🇰🇮",
		"descriptor": "Flag: Kiribati"
	},
	"1F1F0-1F1F2": {
		"key": "1F1F0-1F1F2",
		"value": "🇰🇲",
		"descriptor": "Flag: Comoros"
	},
	"1F1F0-1F1F3": {
		"key": "1F1F0-1F1F3",
		"value": "🇰🇳",
		"descriptor": "Flag: St. Kitts \u0026 Nevis"
	},
	"1F1F0-1F1F5": {
		"key": "1F1F0-1F1F5",
		"value": "🇰🇵",
		"descriptor": "Flag: North Korea"
	},
	"1F1F0-1F1F7": {
		"key": "1F1F0-1F1F7",
		"value": "🇰🇷",
		"descriptor": "Flag: South Korea"
	},
	"1F1F0-1F1FC": {
		"key": "1F1F0-1F1FC",
		"value": "🇰🇼",
		"descriptor": "Flag: Kuwait"
	},
	"1F1F0-1F1FE": {
		"key": "1F1F0-1F1FE",
		"value": "🇰🇾",
		"descriptor": "Flag: Cayman Islands"
	},
	"1F1F0-1F1FF": {
		"key": "1F1F0-1F1FF",
		"value": "🇰🇿",
		"descriptor": "Flag: Kazakhstan"
	},
	"1F1F1-1F1E6": {
		"key": "1F1F1-1F1E6",
		"value": "🇱🇦",
		"descriptor": "Flag: Laos"
	},
	"1F1F1-1F1E7": {
		"key": "1F1F1-1F1E7",
		"value": "🇱🇧",
		"descriptor": "Flag: Lebanon"
	},
	"1F1F1-1F1E8": {
		"key": "1F1F1-1F1E8",
		"value": "🇱🇨",
		"descriptor": "Flag: St. Lucia"
	},
	"1F1F1-1F1EE": {
		"key": "1F1F1-1F1EE",
		"value": "🇱🇮",
		"descriptor": "Flag: Liechtenstein"
	},
	"1F1F1-1F1F0": {
		"key": "1F1F1-1F1F0",
		"value": "🇱🇰",
		"descriptor": "Flag: Sri Lanka"
	},
	"1F1F1-1F1F7": {
		"key": "1F1F1-1F1F7",
		"value": "🇱🇷",
		"descriptor": "Flag: Liberia"
	},
	"1F1F1-1F1F8": {
		"key": "1F1F1-1F1F8",
		"value": "🇱🇸",
		"descriptor": "Flag: Lesotho"
	},
	"1F1F1-1F1F9": {
		"key": "1F1F1-1F1F9",
		"value": "🇱🇹",
		"descriptor": "Flag: Lithuania"
	},
	"1F1F1-1F1FA": {
		"key": "1F1F1-1F1FA",
		"value": "🇱🇺",
		"descriptor": "Flag: Luxembourg"
	},
	"1F1F1-1F1FB": {
		"key": "1F1F1-1F1FB",
		"value": "🇱🇻",
		"descriptor": "Flag: Latvia"
	},
	"1F1F1-1F1FE": {
		"key": "1F1F1-1F1FE",
		"value": "🇱🇾",
		"descriptor": "Flag: Libya"
	},
	"1F1F2-1F1E6": {
		"key": "1F1F2-1F1E6",
		"value": "🇲🇦",
		"descriptor": "Flag: Morocco"
	},
	"1F1F2-1F1E8": {
		"key": "1F1F2-1F1E8",
		"value": "🇲🇨",
		"descriptor": "Flag: Monaco"
	},
	"1F1F2-1F1E9": {
		"key": "1F1F2-1F1E9",
		"value": "🇲🇩",
		"descriptor": "Flag: Moldova"
	},
	"1F1F2-1F1EA": {
		"key": "1F1F2-1F1EA",
		"value": "🇲🇪",
		"descriptor": "Flag: Montenegro"
	},
	"1F1F2-1F1EB": {
		"key": "1F1F2-1F1EB",
		"value": "🇲🇫",
		"descriptor": "Flag: St. Martin"
	},
	"1F1F2-1F1EC": {
		"key": "1F1F2-1F1EC",
		"value": "🇲🇬",
		"descriptor": "Flag: Madagascar"
	},
	"1F1F2-1F1ED": {
		"key": "1F1F2-1F1ED",
		"value": "🇲🇭",
		"descriptor": "Flag: Marshall Islands"
	},
	"1F1F2-1F1F0": {
		"key": "1F1F2-1F1F0",
		"value": "🇲🇰",
		"descriptor": "Flag: North Macedonia"
	},
	"1F1F2-1F1F1": {
		"key": "1F1F2-1F1F1",
		"value": "🇲🇱",
		"descriptor": "Flag: Mali"
	},
	"1F1F2-1F1F2": {
		"key": "1F1F2-1F1F2",
		"value": "🇲🇲",
		"descriptor": "Flag: Myanmar (Burma)"
	},
	"1F1F2-1F1F3": {
		"key": "1F1F2-1F1F3",
		"value": "🇲🇳",
		"descriptor": "Flag: Mongolia"
	},
	"1F1F2-1F1F4": {
		"key": "1F1F2-1F1F4",
		"value": "🇲🇴",
		"descriptor": "Flag: Macao Sar China"
	},
	"1F1F2-1F1F5": {
		"key": "1F1F2-1F1F5",
		"value": "🇲🇵",
		"descriptor": "Flag: Northern Mariana Islands"
	},
	"1F1F2-1F1F6": {
		"key": "1F1F2-1F1F6",
		"value": "🇲🇶",
		"descriptor": "Flag: Martinique"
	},
	"1F1F2-1F1F7": {
		"key": "1F1F2-1F1F7",
		"value": "🇲🇷",
		"descriptor": "Flag: Mauritania"
	},
	"1F1F2-1F1F8": {
		"key": "1F1F2-1F1F8",
		"value": "🇲🇸",
		"descriptor": "Flag: Montserrat"
	},
	"1F1F2-1F1F9": {
		"key": "1F1F2-1F1F9",
		"value": "🇲🇹",
		"descriptor": "Flag: Malta"
	},
	"1F1F2-1F1FA": {
		"key": "1F1F2-1F1FA",
		"value": "🇲🇺",
		"descriptor": "Flag: Mauritius"
	},
	"1F1F2-1F1FB": {
		"key": "1F1F2-1F1FB",
		"value": "🇲🇻",
		"descriptor": "Flag: Maldives"
	},
	"1F1F2-1F1FC": {
		"key": "1F1F2-1F1FC",
		"value": "🇲🇼",
		"descriptor": "Flag: Malawi"
	},
	"1F1F2-1F1FD": {
		"key": "1F1F2-1F1FD",
		"value": "🇲🇽",
		"descriptor": "Flag: Mexico"
	},
	"1F1F2-1F1FE": {
		"key": "1F1F2-1F1FE",
		"value": "🇲🇾",
		"descriptor": "Flag: Malaysia"
	},
	"1F1F2-1F1FF": {
		"key": "1F1F2-1F1FF",
		"value": "🇲🇿",
		"descriptor": "Flag: Mozambique"
	},
	"1F1F3-1F1E6": {
		"key": "1F1F3-1F1E6",
		"value": "🇳🇦",
		"descriptor": "Flag: Namibia"
	},
	"1F1F3-1F1E8": {
		"key": "1F1F3-1F1E8",
		"value": "🇳🇨",
		"descriptor": "Flag: New Caledonia"
	},
	"1F1F3-1F1EA": {
		"key": "1F1F3-1F1EA",
		"value": "🇳🇪",
		"descriptor": "Flag: Niger"
	},
	"1F1F3-1F1EB": {
		"key": "1F1F3-1F1EB",
		"value": "🇳🇫",
		"descriptor": "Flag: Norfolk Island"
	},
	"1F1F3-1F1EC": {
		"key": "1F1F3-1F1EC",
		"value": "🇳🇬",
		"descriptor": "Flag: Nigeria"
	},
	"1F1F3-1F1EE": {
		"key": "1F1F3-1F1EE",
		"value": "🇳🇮",
		"descriptor": "Flag: Nicaragua"
	},
	"1F1F3-1F1F1": {
		"key": "1F1F3-1F1F1",
		"value": "🇳🇱",
		"descriptor": "Flag: Netherlands"
	},
	"1F1F3-1F1F4": {
		"key": "1F1F3-1F1F4",
		"value": "🇳🇴",
		"descriptor": "Flag: Norway"
	},
	"1F1F3-1F1F5": {
		"key": "1F1F3-1F1F5",
		"value": "🇳🇵",
		"descriptor": "Flag: Nepal"
	},
	"1F1F3-1F1F7": {
		"key": "1F1F3-1F1F7",
		"value": "🇳🇷",
		"descriptor": "Flag: Nauru"
	},
	"1F1F3-1F1FA": {
		"key": "1F1F3-1F1FA",
		"value": "🇳🇺",
		"descriptor": "Flag: Niue"
	},
	"1F1F3-1F1FF": {
		"key": "1F1F3-1F1FF",
		"value": "🇳🇿",
		"descriptor": "Flag: New Zealand"
	},
	"1F1F4-1F1F2": {
		"key": "1F1F4-1F1F2",
		"value": "🇴🇲",
		"descriptor": "Flag: Oman"
	},
	"1F1F5-1F1E6": {
		"key": "1F1F5-1F1E6",
		"value": "🇵🇦",
		"descriptor": "Flag: Panama"
	},
	"1F1F5-1F1EA": {
		"key": "1F1F5-1F1EA",
		"value": "🇵🇪",
		"descriptor": "Flag: Peru"
	},
	"1F1F5-1F1EB": {
		"key": "1F1F5-1F1EB",
		"value": "🇵🇫",
		"descriptor": "Flag: French Polynesia"
	},
	"1F1F5-1F1EC": {
		"key": "1F1F5-1F1EC",
		"value": "🇵🇬",
		"descriptor": "Flag: Papua New Guinea"
	},
	"1F1F5-1F1ED": {
		"key": "1F1F5-1F1ED",
		"value": "🇵🇭",
		"descriptor": "Flag: Philippines"
	},
	"1F1F5-1F1F0": {
		"key": "1F1F5-1F1F0",
		"value": "🇵🇰",
		"descriptor": "Flag: Pakistan"
	},
	"1F1F5-1F1F1": {
		"key": "1F1F5-1F1F1",
		"value": "🇵🇱",
		"descriptor": "Flag: Poland"
	},
	"1F1F5-1F1F2": {
		"key": "1F1F5-1F1F2",
		"value": "🇵🇲",
		"descriptor": "Flag: St. Pierre \u0026 Miquelon"
	},
	"1F1F5-1F1F3": {
		"key": "1F1F5-1F1F3",
		"value": "🇵🇳",
		"descriptor": "Flag: Pitcairn Islands"
	},
	"1F1F5-1F1F7": {
		"key": "1F1F5-1F1F7",
		"value": "🇵🇷",
		"descriptor": "Flag: Puerto Rico"
	},
	"1F1F5-1F1F8": {
		"key": "1F1F5-1F1F8",
		"value": "🇵🇸",
		"descriptor": "Flag: Palestinian Territories"
	},
	"1F1F5-1F1F9": {
		"key": "1F1F5-1F1F9",
		"value": "🇵🇹",
		"descriptor": "Flag: Portugal"
	},
	"1F1F5-1F1FC": {
		"key": "1F1F5-1F1FC",
		"value": "🇵🇼",
		"descriptor": "Flag: Palau"
	},
	"1F1F5-1F1FE": {
		"key": "1F1F5-1F1FE",
		"value": "🇵🇾",
		"descriptor": "Flag: Paraguay"
	},
	"1F1F6-1F1E6": {
		"key": "1F1F6-1F1E6",
		"value": "🇶🇦",
		"descriptor": "Flag: Qatar"
	},
	"1F1F7-1F1EA": {
		"key": "1F1F7-1F1EA",
		"value": "🇷🇪",
		"descriptor": "Flag: Réunion"
	},
	"1F1F7-1F1F4": {
		"key": "1F1F7-1F1F4",
		"value": "🇷🇴",
		"descriptor": "Flag: Romania"
	},
	"1F1F7-1F1F8": {
		"key": "1F1F7-1F1F8",
		"value": "🇷🇸",
		"descriptor": "Flag: Serbia"
	},
	"1F1F7-1F1FA": {
		"key": "1F1F7-1F1FA",
		"value": "🇷🇺",
		"descriptor": "Flag: Russia"
	},
	"1F1F7-1F1FC": {
		"key": "1F1F7-1F1FC",
		"value": "🇷🇼",
		"descriptor": "Flag: Rwanda"
	},
	"1F1F8-1F1E6": {
		"key": "1F1F8-1F1E6",
		"value": "🇸🇦",
		"descriptor": "Flag: Saudi Arabia"
	},
	"1F1F8-1F1E7": {
		"key": "1F1F8-1F1E7",
		"value": "🇸🇧",
		"descriptor": "Flag: Solomon Islands"
	},
	"1F1F8-1F1E8": {
		"key": "1F1F8-1F1E8",
		"value": "🇸🇨",
		"descriptor": "Flag: Seychelles"
	},
	"1F1F8-1F1E9": {
		"key": "1F1F8-1F1E9",
		"value": "🇸🇩",
		"descriptor": "Flag: Sudan"
	},
	"1F1F8-1F1EA": {
		"key": "1F1F8-1F1EA",
		"value": "🇸🇪",
		"descriptor": "Flag: Sweden"
	},
	"1F1F8-1F1EC": {
		"key": "1F1F8-1F1EC",
		"value": "🇸🇬",
		"descriptor": "Flag: Singapore"
	},
	"1F1F8-1F1ED": {
		"key": "1F1F8-1F1ED",
		"value": "🇸🇭",
		"descriptor": "Flag: St. Helena"
	},
	"1F1F8-1F1EE": {
		"key": "1F1F8-1F1EE",
		"value": "🇸🇮",
		"descriptor": "Flag: Slovenia"
	},
	"1F1F8-1F1EF": {
		"key": "1F1F8-1F1EF",
		"value": "🇸🇯",
		"descriptor": "Flag: Svalbard \u0026 Jan Mayen"
	},
	"1F1F8-1F1F0": {
		"key": "1F1F8-1F1F0",
		"value": "🇸🇰",
		"descriptor": "Flag: Slovakia"
	},
	"1F1F8-1F1F1": {
		"key": "1F1F8-1F1F1",
		"value": "🇸🇱",
		"descriptor": "Flag: Sierra Leone"
	},
	"1F1F8-1F1F2": {
		"key": "1F1F8-1F1F2",
		"value": "🇸🇲",
		"descriptor": "Flag: San Marino"
	},
	"1F1F8-1F1F3": {
		"key": "1F1F8-1F1F3",
		"value": "🇸🇳",
		"descriptor": "Flag: Senegal"
	},
	"1F1F8-1F1F4": {
		"key": "1F1F8-1F1F4",
		"value": "🇸🇴",
		"descriptor": "Flag: Somalia"
	},
	"1F1F8-1F1F7": {
		"key": "1F1F8-1F1F7",
		"value": "🇸🇷",
		"descriptor": "Flag: Suriname"
	},
	"1F1F8-1F1F8": {
		"key": "1F1F8-1F1F8",
		"value": "🇸🇸",
		"descriptor": "Flag: South Sudan"
	},
	"1F1F8-1F1F9": {
		"key": "1F1F8-1F1F9",
		"value": "🇸🇹",
		"descriptor": "Flag: São Tomé \u0026 Príncipe"
	},
	"1F1F8-1F1FB": {
		"key": "1F1F8-1F1FB",
		"value": "🇸🇻",
		"descriptor": "Flag: El Salvador"
	},
	"1F1F8-1F1FD": {
		"key": "1F1F8-1F1FD",
		"value": "🇸🇽",
		"descriptor": "Flag: Sint Maarten"
	},
	"1F1F8-1F1FE": {
		"key": "1F1F8-1F1FE",
		"value": "🇸🇾",
		"descriptor": "Flag: Syria"
	},
	"1F1F8-1F1FF": {
		"key": "1F1F8-1F1FF",
		"value": "🇸🇿",
		"descriptor": "Flag: Eswatini"
	},
	"1F1F9-1F1E6": {
		"key": "1F1F9-1F1E6",
		"value": "🇹🇦",
		"descriptor": "Flag: Tristan Da Cunha"
	},
	"1F1F9-1F1E8": {
		"key": "1F1F9-1F1E8",
		"value": "🇹🇨",
		"descriptor": "Flag: Turks \u0026 Caicos Islands"
	},
	"1F1F9-1F1E9": {
		"key": "1F1F9-1F1E9",
		"value": "🇹🇩",
		"descriptor": "Flag: Chad"
	},
	"1F1F9-1F1EB": {
		"key": "1F1F9-1F1EB",
		"value": "🇹🇫",
		"descriptor": "Flag: French Southern Territories"
	},
	"1F1F9-1F1EC": {
		"key": "1F1F9-1F1EC",
		"value": "🇹🇬",
		"descriptor": "Flag: Togo"
	},
	"1F1F9-1F1ED": {
		"key": "1F1F9-1F1ED",
		"value": "🇹🇭",
		"descriptor": "Flag: Thailand"
	},
	"1F1F9-1F1EF": {
		"key": "1F1F9-1F1EF",
		"value": "🇹🇯",
		"descriptor": "Flag: Tajikistan"
	},
	"1F1F9-1F1F0": {
		"key": "1F1F9-1F1F0",
		"value": "🇹🇰",
		"descriptor": "Flag: Tokelau"
	},
	"1F1F9-1F1F1": {
		"key": "1F1F9-1F1F1",
		"value": "🇹🇱",
		"descriptor": "Flag: Timor-Leste"
	},
	"1F1F9-1F1F2": {
		"key": "1F1F9-1F1F2",
		"value": "🇹🇲",
		"descriptor": "Flag: Turkmenistan"
	},
	"1F1F9-1F1F3": {
		"key": "1F1F9-1F1F3",
		"value": "🇹🇳",
		"descriptor": "Flag: Tunisia"
	},
	"1F1F9-1F1F4": {
		"key": "1F1F9-1F1F4",
		"value": "🇹🇴",
		"descriptor": "Flag: Tonga"
	},
	"1F1F9-1F1F7": {
		"key": "1F1F9-1F1F7",
		"value": "🇹🇷",
		"descriptor": "Flag: Turkey"
	},
	"1F1F9-1F1F9": {
		"key": "1F1F9-1F1F9",
		"value": "🇹🇹",
		"descriptor": "Flag: Trinidad \u0026 Tobago"
	},
	"1F1F9-1F1FB": {
		"key": "1F1F9-1F1FB",
		"value": "🇹🇻",
		"descriptor": "Flag: Tuvalu"
	},
	"1F1F9-1F1FC": {
		"key": "1F1F9-1F1FC",
		"value": "🇹🇼",
		"descriptor": "Flag: Taiwan"
	},
	"1F1F9-1F1FF": {
		"key": "1F1F9-1F1FF",
		"value": "🇹🇿",
		"descriptor": "Flag: Tanzania"
	},
	"1F1FA-1F1E6": {
		"key": "1F1FA-1F1E6",
		"value": "🇺🇦",
		"descriptor": "Flag: Ukraine"
	},
	"1F1FA-1F1EC": {
		"key": "1F1FA-1F1EC",
		"value": "🇺🇬",
		"descriptor": "Flag: Uganda"
	},
	"1F1FA-1F1F2": {
		"key": "1F1FA-1F1F2",
		"value": "🇺🇲",
		"descriptor": "Flag: U.S. Outlying Islands"
	},
	"1F1FA-1F1F3": {
		"key": "1F1FA-1F1F3",
		"value": "🇺🇳",
		"descriptor": "Flag: United Nations"
	},
	"1F1FA-1F1F8": {
		"key": "1F1FA-1F1F8",
		"value": "🇺🇸",
		"descriptor": "Flag: United States"
	},
	"1F1FA-1F1FE": {
		"key": "1F1FA-1F1FE",
		"value": "🇺🇾",
		"descriptor": "Flag: Uruguay"
	},
	"1F1FA-1F1FF": {
		"key": "1F1FA-1F1FF",
		"value": "🇺🇿",
		"descriptor": "Flag: Uzbekistan"
	},
	"1F1FB-1F1E6": {
		"key": "1F1FB-1F1E6",
		"value": "🇻🇦",
		"descriptor": "Flag: Vatican City"
	},
	"1F1FB-1F1E8": {
		"key": "1F1FB-1F1E8",
		"value": "🇻🇨",
		"descriptor": "Flag: St. Vincent \u0026 Grenadines"
	},
	"1F1FB-1F1EA": {
		"key": "1F1FB-1F1EA",
		"value": "🇻🇪",
		"descriptor": "Flag: Venezuela"
	},
	"1F1FB-1F1EC": {
		"key": "1F1FB-1F1EC",
		"value": "🇻🇬",
		"descriptor": "Flag: British Virgin Islands"
	},
	"1F1FB-1F1EE": {
		"key": "1F1FB-1F1EE",
		"value": "🇻🇮",
		"descriptor": "Flag: U.S. Virgin Islands"
	},
	"1F1FB-1F1F3": {
		"key": "1F1FB-1F1F3",
		"value": "🇻🇳",
		"descriptor": "Flag: Vietnam"
	},
	"1F1FB-1F1FA": {
		"key": "1F1FB-1F1FA",
		"value": "🇻🇺",
		"descriptor": "Flag: Vanuatu"
	},
	"1F1FC-1F1EB": {
		"key": "1F1FC-1F1EB",
		"value": "🇼🇫",
		"descriptor": "Flag: Wallis \u0026 Futuna"
	},
	"1F1FC-1F1F8": {
		"key": "1F1FC-1F1F8",
		"value": "🇼🇸",
		"descriptor": "Flag: Samoa"
	},
	"1F1FD-1F1F0": {
		"key": "1F1FD-1F1F0",
		"value": "🇽🇰",
		"descriptor": "Flag: Kosovo"
	},
	"1F1FE-1F1EA": {
		"key": "1F1FE-1F1EA",
		"value": "🇾🇪",
		"descriptor": "Flag: Yemen"
	},
	"1F1FE-1F1F9": {
		"key": "1F1FE-1F1F9",
		"value": "🇾🇹",
		"descriptor": "Flag: Mayotte"
	},
	"1F1FF-1F1E6": {
		"key": "1F1FF-1F1E6",
		"value": "🇿🇦",
		"descriptor": "Flag: South Africa"
	},
	"1F1FF-1F1F2": {
		"key": "1F1FF-1F1F2",
		"value": "🇿🇲",
		"descriptor": "Flag: Zambia"
	},
	"1F1FF-1F1FC": {
		"key": "1F1FF-1F1FC",
		"value": "🇿🇼",
		"descriptor": "Flag: Zimbabwe"
	},
	"1F201": {
		"key": "1F201",
		"value": "🈁",
		"descriptor": "Japanese “Here” Button"
	},
	"1F202-FE0F": {
		"key": "1F202-FE0F",
		"value": "🈂️",
		"descriptor": "Japanese “Service Charge” Button"
	},
	"1F21A": {
		"key": "1F21A",
		"value": "🈚",
		"descriptor": "Japanese “Free of Charge” Button"
	},
	"1F22F": {
		"key": "1F22F",
		"value": "🈯",
		"descriptor": "Japanese “Reserved” Button"
	},
	"1F232": {
		"key": "1F232",
		"value": "🈲",
		"descriptor": "Japanese “Prohibited” Button"
	},
	"1F233": {
		"key": "1F233",
		"value": "🈳",
		"descriptor": "Japanese “Vacancy” Button"
	},
	"1F234": {
		"key": "1F234",
		"value": "🈴",
		"descriptor": "Japanese “Passing Grade” Button"
	},
	"1F235": {
		"key": "1F235",
		"value": "🈵",
		"descriptor": "Japanese “No Vacancy” Button"
	},
	"1F236": {
		"key": "1F236",
		"value": "🈶",
		"descriptor": "Japanese “Not Free of Charge” Button"
	},
	"1F237-FE0F": {
		"key": "1F237-FE0F",
		"value": "🈷️",
		"descriptor": "Japanese “Monthly Amount” Button"
	},
	"1F238": {
		"key": "1F238",
		"value": "🈸",
		"descriptor": "Japanese “Application” Button"
	},
	"1F239": {
		"key": "1F239",
		"value": "🈹",
		"descriptor": "Japanese “Discount” Button"
	},
	"1F23A": {
		"key": "1F23A",
		"value": "🈺",
		"descriptor": "Japanese “Open for Business” Button"
	},
	"1F250": {
		"key": "1F250",
		"value": "🉐",
		"descriptor": "Japanese “Bargain” Button"
	},
	"1F251": {
		"key": "1F251",
		"value": "🉑",
		"descriptor": "Japanese “Acceptable” Button"
	},
	"1F300": {
		"key": "1F300",
		"value": "🌀",
		"descriptor": "Cyclone"
	},
	"1F301": {
		"key": "1F301",
		"value": "🌁",
		"descriptor": "Foggy"
	},
	"1F302": {
		"key": "1F302",
		"value": "🌂",
		"descriptor": "Closed Umbrella"
	},
	"1F303": {
		"key": "1F303",
		"value": "🌃",
		"descriptor": "Night With Stars"
	},
	"1F304": {
		"key": "1F304",
		"value": "🌄",
		"descriptor": "Sunrise Over Mountains"
	},
	"1F305": {
		"key": "1F305",
		"value": "🌅",
		"descriptor": "Sunrise"
	},
	"1F306": {
		"key": "1F306",
		"value": "🌆",
		"descriptor": "Cityscape at Dusk"
	},
	"1F307": {
		"key": "1F307",
		"value": "🌇",
		"descriptor": "Sunset"
	},
	"1F308": {
		"key": "1F308",
		"value": "🌈",
		"descriptor": "Rainbow"
	},
	"1F309": {
		"key": "1F309",
		"value": "🌉",
		"descriptor": "Bridge at Night"
	},
	"1F30A": {
		"key": "1F30A",
		"value": "🌊",
		"descriptor": "Water Wave"
	},
	"1F30B": {
		"key": "1F30B",
		"value": "🌋",
		"descriptor": "Volcano"
	},
	"1F30C": {
		"key": "1F30C",
		"value": "🌌",
		"descriptor": "Milky Way"
	},
	"1F30D": {
		"key": "1F30D",
		"value": "🌍",
		"descriptor": "Globe Showing Europe-Africa"
	},
	"1F30E": {
		"key": "1F30E",
		"value": "🌎",
		"descriptor": "Globe Showing Americas"
	},
	"1F30F": {
		"key": "1F30F",
		"value": "🌏",
		"descriptor": "Globe Showing Asia-Australia"
	},
	"1F310": {
		"key": "1F310",
		"value": "🌐",
		"descriptor": "Globe With Meridians"
	},
	"1F311": {
		"key": "1F311",
		"value": "🌑",
		"descriptor": "New Moon"
	},
	"1F312": {
		"key": "1F312",
		"value": "🌒",
		"descriptor": "Waxing Crescent Moon"
	},
	"1F313": {
		"key": "1F313",
		"value": "🌓",
		"descriptor": "First Quarter Moon"
	},
	"1F314": {
		"key": "1F314",
		"value": "🌔",
		"descriptor": "Waxing Gibbous Moon"
	},
	"1F315": {
		"key": "1F315",
		"value": "🌕",
		"descriptor": "Full Moon"
	},
	"1F316": {
		"key": "1F316",
		"value": "🌖",
		"descriptor": "Waning Gibbous Moon"
	},
	"1F317": {
		"key": "1F317",
		"value": "🌗",
		"descriptor": "Last Quarter Moon"
	},
	"1F318": {
		"key": "1F318",
		"value": "🌘",
		"descriptor": "Waning Crescent Moon"
	},
	"1F319": {
		"key": "1F319",
		"value": "🌙",
		"descriptor": "Crescent Moon"
	},
	"1F31A": {
		"key": "1F31A",
		"value": "🌚",
		"descriptor": "New Moon Face"
	},
	"1F31B": {
		"key": "1F31B",
		"value": "🌛",
		"descriptor": "First Quarter Moon Face"
	},
	"1F31C": {
		"key": "1F31C",
		"value": "🌜",
		"descriptor": "Last Quarter Moon Face"
	},
	"1F31D": {
		"key": "1F31D",
		"value": "🌝",
		"descriptor": "Full Moon Face"
	},
	"1F31E": {
		"key": "1F31E",
		"value": "🌞",
		"descriptor": "Sun With Face"
	},
	"1F31F": {
		"key": "1F31F",
		"value": "🌟",
		"descriptor": "Glowing Star"
	},
	"1F320": {
		"key": "1F320",
		"value": "🌠",
		"descriptor": "Shooting Star"
	},
	"1F321-FE0F": {
		"key": "1F321-FE0F",
		"value": "🌡️",
		"descriptor": "Thermometer"
	},
	"1F324-FE0F": {
		"key": "1F324-FE0F",
		"value": "🌤️",
		"descriptor": "Sun Behind Small Cloud"
	},
	"1F325-FE0F": {
		"key": "1F325-FE0F",
		"value": "🌥️",
		"descriptor": "Sun Behind Large Cloud"
	},
	"1F326-FE0F": {
		"key": "1F326-FE0F",
		"value": "🌦️",
		"descriptor": "Sun Behind Rain Cloud"
	},
	"1F327-FE0F": {
		"key": "1F327-FE0F",
		"value": "🌧️",
		"descriptor": "Cloud With Rain"
	},
	"1F328-FE0F": {
		"key": "1F328-FE0F",
		"value": "🌨️",
		"descriptor": "Cloud With Snow"
	},
	"1F329-FE0F": {
		"key": "1F329-FE0F",
		"value": "🌩️",
		"descriptor": "Cloud With Lightning"
	},
	"1F32A-FE0F": {
		"key": "1F32A-FE0F",
		"value": "🌪️",
		"descriptor": "Tornado"
	},
	"1F32B-FE0F": {
		"key": "1F32B-FE0F",
		"value": "🌫️",
		"descriptor": "Fog"
	},
	"1F32C-FE0F": {
		"key": "1F32C-FE0F",
		"value": "🌬️",
		"descriptor": "Wind Face"
	},
	"1F32D": {
		"key": "1F32D",
		"value": "🌭",
		"descriptor": "Hot Dog"
	},
	"1F32E": {
		"key": "1F32E",
		"value": "🌮",
		"descriptor": "Taco"
	},
	"1F32F": {
		"key": "1F32F",
		"value": "🌯",
		"descriptor": "Burrito"
	},
	"1F330": {
		"key": "1F330",
		"value": "🌰",
		"descriptor": "Chestnut"
	},
	"1F331": {
		"key": "1F331",
		"value": "🌱",
		"descriptor": "Seedling"
	},
	"1F332": {
		"key": "1F332",
		"value": "🌲",
		"descriptor": "Evergreen Tree"
	},
	"1F333": {
		"key": "1F333",
		"value": "🌳",
		"descriptor": "Deciduous Tree"
	},
	"1F334": {
		"key": "1F334",
		"value": "🌴",
		"descriptor": "Palm Tree"
	},
	"1F335": {
		"key": "1F335",
		"value": "🌵",
		"descriptor": "Cactus"
	},
	"1F336-FE0F": {
		"key": "1F336-FE0F",
		"value": "🌶️",
		"descriptor": "Hot Pepper"
	},
	"1F337": {
		"key": "1F337",
		"value": "🌷",
		"descriptor": "Tulip"
	},
	"1F338": {
		"key": "1F338",
		"value": "🌸",
		"descriptor": "Cherry Blossom"
	},
	"1F339": {
		"key": "1F339",
		"value": "🌹",
		"descriptor": "Rose"
	},
	"1F33A": {
		"key": "1F33A",
		"value": "🌺",
		"descriptor": "Hibiscus"
	},
	"1F33B": {
		"key": "1F33B",
		"value": "🌻",
		"descriptor": "Sunflower"
	},
	"1F33C": {
		"key": "1F33C",
		"value": "🌼",
		"descriptor": "Blossom"
	},
	"1F33D": {
		"key": "1F33D",
		"value": "🌽",
		"descriptor": "Ear of Corn"
	},
	"1F33E": {
		"key": "1F33E",
		"value": "🌾",
		"descriptor": "Sheaf of Rice"
	},
	"1F33F": {
		"key": "1F33F",
		"value": "🌿",
		"descriptor": "Herb"
	},
	"1F340": {
		"key": "1F340",
		"value": "🍀",
		"descriptor": "Four Leaf Clover"
	},
	"1F341": {
		"key": "1F341",
		"value": "🍁",
		"descriptor": "Maple Leaf"
	},
	"1F342": {
		"key": "1F342",
		"value": "🍂",
		"descriptor": "Fallen Leaf"
	},
	"1F343": {
		"key": "1F343",
		"value": "🍃",
		"descriptor": "Leaf Fluttering in Wind"
	},
	"1F344": {
		"key": "1F344",
		"value": "🍄",
		"descriptor": "Mushroom"
	},
	"1F345": {
		"key": "1F345",
		"value": "🍅",
		"descriptor": "Tomato"
	},
	"1F346": {
		"key": "1F346",
		"value": "🍆",
		"descriptor": "Eggplant"
	},
	"1F347": {
		"key": "1F347",
		"value": "🍇",
		"descriptor": "Grapes"
	},
	"1F348": {
		"key": "1F348",
		"value": "🍈",
		"descriptor": "Melon"
	},
	"1F349": {
		"key": "1F349",
		"value": "🍉",
		"descriptor": "Watermelon"
	},
	"1F34A": {
		"key": "1F34A",
		"value": "🍊",
		"descriptor": "Tangerine"
	},
	"1F34B": {
		"key": "1F34B",
		"value": "🍋",
		"descriptor": "Lemon"
	},
	"1F34C": {
		"key": "1F34C",
		"value": "🍌",
		"descriptor": "Banana"
	},
	"1F34D": {
		"key": "1F34D",
		"value": "🍍",
		"descriptor": "Pineapple"
	},
	"1F34E": {
		"key": "1F34E",
		"value": "🍎",
		"descriptor": "Red Apple"
	},
	"1F34F": {
		"key": "1F34F",
		"value": "🍏",
		"descriptor": "Green Apple"
	},
	"1F350": {
		"key": "1F350",
		"value": "🍐",
		"descriptor": "Pear"
	},
	"1F351": {
		"key": "1F351",
		"value": "🍑",
		"descriptor": "Peach"
	},
	"1F352": {
		"key": "1F352",
		"value": "🍒",
		"descriptor": "Cherries"
	},
	"1F353": {
		"key": "1F353",
		"value": "🍓",
		"descriptor": "Strawberry"
	},
	"1F354": {
		"key": "1F354",
		"value": "🍔",
		"descriptor": "Hamburger"
	},
	"1F355": {
		"key": "1F355",
		"value": "🍕",
		"descriptor": "Pizza"
	},
	"1F356": {
		"key": "1F356",
		"value": "🍖",
		"descriptor": "Meat on Bone"
	},
	"1F357": {
		"key": "1F357",
		"value": "🍗",
		"descriptor": "Poultry Leg"
	},
	"1F358": {
		"key": "1F358",
		"value": "🍘",
		"descriptor": "Rice Cracker"
	},
	"1F359": {
		"key": "1F359",
		"value": "🍙",
		"descriptor": "Rice Ball"
	},
	"1F35A": {
		"key": "1F35A",
		"value": "🍚",
		"descriptor": "Cooked Rice"
	},
	"1F35B": {
		"key": "1F35B",
		"value": "🍛",
		"descriptor": "Curry Rice"
	},
	"1F35C": {
		"key": "1F35C",
		"value": "🍜",
		"descriptor": "Steaming Bowl"
	},
	"1F35D": {
		"key": "1F35D",
		"value": "🍝",
		"descriptor": "Spaghetti"
	},
	"1F35E": {
		"key": "1F35E",
		"value": "🍞",
		"descriptor": "Bread"
	},
	"1F35F": {
		"key": "1F35F",
		"value": "🍟",
		"descriptor": "French Fries"
	},
	"1F360": {
		"key": "1F360",
		"value": "🍠",
		"descriptor": "Roasted Sweet Potato"
	},
	"1F361": {
		"key": "1F361",
		"value": "🍡",
		"descriptor": "Dango"
	},
	"1F362": {
		"key": "1F362",
		"value": "🍢",
		"descriptor": "Oden"
	},
	"1F363": {
		"key": "1F363",
		"value": "🍣",
		"descriptor": "Sushi"
	},
	"1F364": {
		"key": "1F364",
		"value": "🍤",
		"descriptor": "Fried Shrimp"
	},
	"1F365": {
		"key": "1F365",
		"value": "🍥",
		"descriptor": "Fish Cake With Swirl"
	},
	"1F366": {
		"key": "1F366",
		"value": "🍦",
		"descriptor": "Soft Ice Cream"
	},
	"1F367": {
		"key": "1F367",
		"value": "🍧",
		"descriptor": "Shaved Ice"
	},
	"1F368": {
		"key": "1F368",
		"value": "🍨",
		"descriptor": "Ice Cream"
	},
	"1F369": {
		"key": "1F369",
		"value": "🍩",
		"descriptor": "Doughnut"
	},
	"1F36A": {
		"key": "1F36A",
		"value": "🍪",
		"descriptor": "Cookie"
	},
	"1F36B": {
		"key": "1F36B",
		"value": "🍫",
		"descriptor": "Chocolate Bar"
	},
	"1F36C": {
		"key": "1F36C",
		"value": "🍬",
		"descriptor": "Candy"
	},
	"1F36D": {
		"key": "1F36D",
		"value": "🍭",
		"descriptor": "Lollipop"
	},
	"1F36E": {
		"key": "1F36E",
		"value": "🍮",
		"descriptor": "Custard"
	},
	"1F36F": {
		"key": "1F36F",
		"value": "🍯",
		"descriptor": "Honey Pot"
	},
	"1F370": {
		"key": "1F370",
		"value": "🍰",
		"descriptor": "Shortcake"
	},
	"1F371": {
		"key": "1F371",
		"value": "🍱",
		"descriptor": "Bento Box"
	},
	"1F372": {
		"key": "1F372",
		"value": "🍲",
		"descriptor": "Pot of Food"
	},
	"1F373": {
		"key": "1F373",
		"value": "🍳",
		"descriptor": "Cooking"
	},
	"1F374": {
		"key": "1F374",
		"value": "🍴",
		"descriptor": "Fork and Knife"
	},
	"1F375": {
		"key": "1F375",
		"value": "🍵",
		"descriptor": "Teacup Without Handle"
	},
	"1F376": {
		"key": "1F376",
		"value": "🍶",
		"descriptor": "Sake"
	},
	"1F377": {
		"key": "1F377",
		"value": "🍷",
		"descriptor": "Wine Glass"
	},
	"1F378": {
		"key": "1F378",
		"value": "🍸",
		"descriptor": "Cocktail Glass"
	},
	"1F379": {
		"key": "1F379",
		"value": "🍹",
		"descriptor": "Tropical Drink"
	},
	"1F37A": {
		"key": "1F37A",
		"value": "🍺",
		"descriptor": "Beer Mug"
	},
	"1F37B": {
		"key": "1F37B",
		"value": "🍻",
		"descriptor": "Clinking Beer Mugs"
	},
	"1F37C": {
		"key": "1F37C",
		"value": "🍼",
		"descriptor": "Baby Bottle"
	},
	"1F37D-FE0F": {
		"key": "1F37D-FE0F",
		"value": "🍽️",
		"descriptor": "Fork and Knife With Plate"
	},
	"1F37E": {
		"key": "1F37E",
		"value": "🍾",
		"descriptor": "Bottle With Popping Cork"
	},
	"1F37F": {
		"key": "1F37F",
		"value": "🍿",
		"descriptor": "Popcorn"
	},
	"1F380": {
		"key": "1F380",
		"value": "🎀",
		"descriptor": "Ribbon"
	},
	"1F381": {
		"key": "1F381",
		"value": "🎁",
		"descriptor": "Wrapped Gift"
	},
	"1F382": {
		"key": "1F382",
		"value": "🎂",
		"descriptor": "Birthday Cake"
	},
	"1F383": {
		"key": "1F383",
		"value": "🎃",
		"descriptor": "Jack-O-Lantern"
	},
	"1F384": {
		"key": "1F384",
		"value": "🎄",
		"descriptor": "Christmas Tree"
	},
	"1F385": {
		"key": "1F385",
		"value": "🎅",
		"descriptor": "Santa Claus"
	},
	"1F385-1F3FB": {
		"key": "1F385-1F3FB",
		"value": "🎅🏻",
		"descriptor": "Santa Claus: Light Skin Tone"
	},
	"1F385-1F3FC": {
		"key": "1F385-1F3FC",
		"value": "🎅🏼",
		"descriptor": "Santa Claus: Medium-Light Skin Tone"
	},
	"1F385-1F3FD": {
		"key": "1F385-1F3FD",
		"value": "🎅🏽",
		"descriptor": "Santa Claus: Medium Skin Tone"
	},
	"1F385-1F3FE": {
		"key": "1F385-1F3FE",
		"value": "🎅🏾",
		"descriptor": "Santa Claus: Medium-Dark Skin Tone"
	},
	"1F385-1F3FF": {
		"key": "1F385-1F3FF",
		"value": "🎅🏿",
		"descriptor": "Santa Claus: Dark Skin Tone"
	},
	"1F386": {
		"key": "1F386",
		"value": "🎆",
		"descriptor": "Fireworks"
	},
	"1F387": {
		"key": "1F387",
		"value": "🎇",
		"descriptor": "Sparkler"
	},
	"1F388": {
		"key": "1F388",
		"value": "🎈",
		"descriptor": "Balloon"
	},
	"1F389": {
		"key": "1F389",
		"value": "🎉",
		"descriptor": "Party Popper"
	},
	"1F38A": {
		"key": "1F38A",
		"value": "🎊",
		"descriptor": "Confetti Ball"
	},
	"1F38B": {
		"key": "1F38B",
		"value": "🎋",
		"descriptor": "Tanabata Tree"
	},
	"1F38C": {
		"key": "1F38C",
		"value": "🎌",
		"descriptor": "Crossed Flags"
	},
	"1F38D": {
		"key": "1F38D",
		"value": "🎍",
		"descriptor": "Pine Decoration"
	},
	"1F38E": {
		"key": "1F38E",
		"value": "🎎",
		"descriptor": "Japanese Dolls"
	},
	"1F38F": {
		"key": "1F38F",
		"value": "🎏",
		"descriptor": "Carp Streamer"
	},
	"1F390": {
		"key": "1F390",
		"value": "🎐",
		"descriptor": "Wind Chime"
	},
	"1F391": {
		"key": "1F391",
		"value": "🎑",
		"descriptor": "Moon Viewing Ceremony"
	},
	"1F392": {
		"key": "1F392",
		"value": "🎒",
		"descriptor": "Backpack"
	},
	"1F393": {
		"key": "1F393",
		"value": "🎓",
		"descriptor": "Graduation Cap"
	},
	"1F396-FE0F": {
		"key": "1F396-FE0F",
		"value": "🎖️",
		"descriptor": "Military Medal"
	},
	"1F397-FE0F": {
		"key": "1F397-FE0F",
		"value": "🎗️",
		"descriptor": "Reminder Ribbon"
	},
	"1F399-FE0F": {
		"key": "1F399-FE0F",
		"value": "🎙️",
		"descriptor": "Studio Microphone"
	},
	"1F39A-FE0F": {
		"key": "1F39A-FE0F",
		"value": "🎚️",
		"descriptor": "Level Slider"
	},
	"1F39B-FE0F": {
		"key": "1F39B-FE0F",
		"value": "🎛️",
		"descriptor": "Control Knobs"
	},
	"1F39E-FE0F": {
		"key": "1F39E-FE0F",
		"value": "🎞️",
		"descriptor": "Film Frames"
	},
	"1F39F-FE0F": {
		"key": "1F39F-FE0F",
		"value": "🎟️",
		"descriptor": "Admission Tickets"
	},
	"1F3A0": {
		"key": "1F3A0",
		"value": "🎠",
		"descriptor": "Carousel Horse"
	},
	"1F3A1": {
		"key": "1F3A1",
		"value": "🎡",
		"descriptor": "Ferris Wheel"
	},
	"1F3A2": {
		"key": "1F3A2",
		"value": "🎢",
		"descriptor": "Roller Coaster"
	},
	"1F3A3": {
		"key": "1F3A3",
		"value": "🎣",
		"descriptor": "Fishing Pole"
	},
	"1F3A4": {
		"key": "1F3A4",
		"value": "🎤",
		"descriptor": "Microphone"
	},
	"1F3A5": {
		"key": "1F3A5",
		"value": "🎥",
		"descriptor": "Movie Camera"
	},
	"1F3A6": {
		"key": "1F3A6",
		"value": "🎦",
		"descriptor": "Cinema"
	},
	"1F3A7": {
		"key": "1F3A7",
		"value": "🎧",
		"descriptor": "Headphone"
	},
	"1F3A8": {
		"key": "1F3A8",
		"value": "🎨",
		"descriptor": "Artist Palette"
	},
	"1F3A9": {
		"key": "1F3A9",
		"value": "🎩",
		"descriptor": "Top Hat"
	},
	"1F3AA": {
		"key": "1F3AA",
		"value": "🎪",
		"descriptor": "Circus Tent"
	},
	"1F3AB": {
		"key": "1F3AB",
		"value": "🎫",
		"descriptor": "Ticket"
	},
	"1F3AC": {
		"key": "1F3AC",
		"value": "🎬",
		"descriptor": "Clapper Board"
	},
	"1F3AD": {
		"key": "1F3AD",
		"value": "🎭",
		"descriptor": "Performing Arts"
	},
	"1F3AE": {
		"key": "1F3AE",
		"value": "🎮",
		"descriptor": "Video Game"
	},
	"1F3AF": {
		"key": "1F3AF",
		"value": "🎯",
		"descriptor": "Direct Hit"
	},
	"1F3B0": {
		"key": "1F3B0",
		"value": "🎰",
		"descriptor": "Slot Machine"
	},
	"1F3B1": {
		"key": "1F3B1",
		"value": "🎱",
		"descriptor": "Pool 8 Ball"
	},
	"1F3B2": {
		"key": "1F3B2",
		"value": "🎲",
		"descriptor": "Game Die"
	},
	"1F3B3": {
		"key": "1F3B3",
		"value": "🎳",
		"descriptor": "Bowling"
	},
	"1F3B4": {
		"key": "1F3B4",
		"value": "🎴",
		"descriptor": "Flower Playing Cards"
	},
	"1F3B5": {
		"key": "1F3B5",
		"value": "🎵",
		"descriptor": "Musical Note"
	},
	"1F3B6": {
		"key": "1F3B6",
		"value": "🎶",
		"descriptor": "Musical Notes"
	},
	"1F3B7": {
		"key": "1F3B7",
		"value": "🎷",
		"descriptor": "Saxophone"
	},
	"1F3B8": {
		"key": "1F3B8",
		"value": "🎸",
		"descriptor": "Guitar"
	},
	"1F3B9": {
		"key": "1F3B9",
		"value": "🎹",
		"descriptor": "Musical Keyboard"
	},
	"1F3BA": {
		"key": "1F3BA",
		"value": "🎺",
		"descriptor": "Trumpet"
	},
	"1F3BB": {
		"key": "1F3BB",
		"value": "🎻",
		"descriptor": "Violin"
	},
	"1F3BC": {
		"key": "1F3BC",
		"value": "🎼",
		"descriptor": "Musical Score"
	},
	"1F3BD": {
		"key": "1F3BD",
		"value": "🎽",
		"descriptor": "Running Shirt"
	},
	"1F3BE": {
		"key": "1F3BE",
		"value": "🎾",
		"descriptor": "Tennis"
	},
	"1F3BF": {
		"key": "1F3BF",
		"value": "🎿",
		"descriptor": "Skis"
	},
	"1F3C0": {
		"key": "1F3C0",
		"value": "🏀",
		"descriptor": "Basketball"
	},
	"1F3C1": {
		"key": "1F3C1",
		"value": "🏁",
		"descriptor": "Chequered Flag"
	},
	"1F3C2": {
		"key": "1F3C2",
		"value": "🏂",
		"descriptor": "Snowboarder"
	},
	"1F3C3": {
		"key": "1F3C3",
		"value": "🏃",
		"descriptor": "Person Running"
	},
	"1F3C3-1F3FB": {
		"key": "1F3C3-1F3FB",
		"value": "🏃🏻",
		"descriptor": "Person Running: Light Skin Tone"
	},
	"1F3C3-1F3FB-200D-2640-FE0F": {
		"key": "1F3C3-1F3FB-200D-2640-FE0F",
		"value": "🏃🏻‍♀️",
		"descriptor": "Woman Running: Light Skin Tone"
	},
	"1F3C3-1F3FB-200D-2642-FE0F": {
		"key": "1F3C3-1F3FB-200D-2642-FE0F",
		"value": "🏃🏻‍♂️",
		"descriptor": "Man Running: Light Skin Tone"
	},
	"1F3C3-1F3FC": {
		"key": "1F3C3-1F3FC",
		"value": "🏃🏼",
		"descriptor": "Person Running: Medium-Light Skin Tone"
	},
	"1F3C3-1F3FC-200D-2640-FE0F": {
		"key": "1F3C3-1F3FC-200D-2640-FE0F",
		"value": "🏃🏼‍♀️",
		"descriptor": "Woman Running: Medium-Light Skin Tone"
	},
	"1F3C3-1F3FC-200D-2642-FE0F": {
		"key": "1F3C3-1F3FC-200D-2642-FE0F",
		"value": "🏃🏼‍♂️",
		"descriptor": "Man Running: Medium-Light Skin Tone"
	},
	"1F3C3-1F3FD": {
		"key": "1F3C3-1F3FD",
		"value": "🏃🏽",
		"descriptor": "Person Running: Medium Skin Tone"
	},
	"1F3C3-1F3FD-200D-2640-FE0F": {
		"key": "1F3C3-1F3FD-200D-2640-FE0F",
		"value": "🏃🏽‍♀️",
		"descriptor": "Woman Running: Medium Skin Tone"
	},
	"1F3C3-1F3FD-200D-2642-FE0F": {
		"key": "1F3C3-1F3FD-200D-2642-FE0F",
		"value": "🏃🏽‍♂️",
		"descriptor": "Man Running: Medium Skin Tone"
	},
	"1F3C3-1F3FE": {
		"key": "1F3C3-1F3FE",
		"value": "🏃🏾",
		"descriptor": "Person Running: Medium-Dark Skin Tone"
	},
	"1F3C3-1F3FE-200D-2640-FE0F": {
		"key": "1F3C3-1F3FE-200D-2640-FE0F",
		"value": "🏃🏾‍♀️",
		"descriptor": "Woman Running: Medium-Dark Skin Tone"
	},
	"1F3C3-1F3FE-200D-2642-FE0F": {
		"key": "1F3C3-1F3FE-200D-2642-FE0F",
		"value": "🏃🏾‍♂️",
		"descriptor": "Man Running: Medium-Dark Skin Tone"
	},
	"1F3C3-1F3FF": {
		"key": "1F3C3-1F3FF",
		"value": "🏃🏿",
		"descriptor": "Person Running: Dark Skin Tone"
	},
	"1F3C3-1F3FF-200D-2640-FE0F": {
		"key": "1F3C3-1F3FF-200D-2640-FE0F",
		"value": "🏃🏿‍♀️",
		"descriptor": "Woman Running: Dark Skin Tone"
	},
	"1F3C3-1F3FF-200D-2642-FE0F": {
		"key": "1F3C3-1F3FF-200D-2642-FE0F",
		"value": "🏃🏿‍♂️",
		"descriptor": "Man Running: Dark Skin Tone"
	},
	"1F3C3-200D-2640-FE0F": {
		"key": "1F3C3-200D-2640-FE0F",
		"value": "🏃‍♀️",
		"descriptor": "Woman Running"
	},
	"1F3C3-200D-2642-FE0F": {
		"key": "1F3C3-200D-2642-FE0F",
		"value": "🏃‍♂️",
		"descriptor": "Man Running"
	},
	"1F3C4": {
		"key": "1F3C4",
		"value": "🏄",
		"descriptor": "Person Surfing"
	},
	"1F3C4-1F3FB": {
		"key": "1F3C4-1F3FB",
		"value": "🏄🏻",
		"descriptor": "Person Surfing: Light Skin Tone"
	},
	"1F3C4-1F3FB-200D-2640-FE0F": {
		"key": "1F3C4-1F3FB-200D-2640-FE0F",
		"value": "🏄🏻‍♀️",
		"descriptor": "Woman Surfing: Light Skin Tone"
	},
	"1F3C4-1F3FB-200D-2642-FE0F": {
		"key": "1F3C4-1F3FB-200D-2642-FE0F",
		"value": "🏄🏻‍♂️",
		"descriptor": "Man Surfing: Light Skin Tone"
	},
	"1F3C4-1F3FC": {
		"key": "1F3C4-1F3FC",
		"value": "🏄🏼",
		"descriptor": "Person Surfing: Medium-Light Skin Tone"
	},
	"1F3C4-1F3FC-200D-2640-FE0F": {
		"key": "1F3C4-1F3FC-200D-2640-FE0F",
		"value": "🏄🏼‍♀️",
		"descriptor": "Woman Surfing: Medium-Light Skin Tone"
	},
	"1F3C4-1F3FC-200D-2642-FE0F": {
		"key": "1F3C4-1F3FC-200D-2642-FE0F",
		"value": "🏄🏼‍♂️",
		"descriptor": "Man Surfing: Medium-Light Skin Tone"
	},
	"1F3C4-1F3FD": {
		"key": "1F3C4-1F3FD",
		"value": "🏄🏽",
		"descriptor": "Person Surfing: Medium Skin Tone"
	},
	"1F3C4-1F3FD-200D-2640-FE0F": {
		"key": "1F3C4-1F3FD-200D-2640-FE0F",
		"value": "🏄🏽‍♀️",
		"descriptor": "Woman Surfing: Medium Skin Tone"
	},
	"1F3C4-1F3FD-200D-2642-FE0F": {
		"key": "1F3C4-1F3FD-200D-2642-FE0F",
		"value": "🏄🏽‍♂️",
		"descriptor": "Man Surfing: Medium Skin Tone"
	},
	"1F3C4-1F3FE": {
		"key": "1F3C4-1F3FE",
		"value": "🏄🏾",
		"descriptor": "Person Surfing: Medium-Dark Skin Tone"
	},
	"1F3C4-1F3FE-200D-2640-FE0F": {
		"key": "1F3C4-1F3FE-200D-2640-FE0F",
		"value": "🏄🏾‍♀️",
		"descriptor": "Woman Surfing: Medium-Dark Skin Tone"
	},
	"1F3C4-1F3FE-200D-2642-FE0F": {
		"key": "1F3C4-1F3FE-200D-2642-FE0F",
		"value": "🏄🏾‍♂️",
		"descriptor": "Man Surfing: Medium-Dark Skin Tone"
	},
	"1F3C4-1F3FF": {
		"key": "1F3C4-1F3FF",
		"value": "🏄🏿",
		"descriptor": "Person Surfing: Dark Skin Tone"
	},
	"1F3C4-1F3FF-200D-2640-FE0F": {
		"key": "1F3C4-1F3FF-200D-2640-FE0F",
		"value": "🏄🏿‍♀️",
		"descriptor": "Woman Surfing: Dark Skin Tone"
	},
	"1F3C4-1F3FF-200D-2642-FE0F": {
		"key": "1F3C4-1F3FF-200D-2642-FE0F",
		"value": "🏄🏿‍♂️",
		"descriptor": "Man Surfing: Dark Skin Tone"
	},
	"1F3C4-200D-2640-FE0F": {
		"key": "1F3C4-200D-2640-FE0F",
		"value": "🏄‍♀️",
		"descriptor": "Woman Surfing"
	},
	"1F3C4-200D-2642-FE0F": {
		"key": "1F3C4-200D-2642-FE0F",
		"value": "🏄‍♂️",
		"descriptor": "Man Surfing"
	},
	"1F3C5": {
		"key": "1F3C5",
		"value": "🏅",
		"descriptor": "Sports Medal"
	},
	"1F3C6": {
		"key": "1F3C6",
		"value": "🏆",
		"descriptor": "Trophy"
	},
	"1F3C7": {
		"key": "1F3C7",
		"value": "🏇",
		"descriptor": "Horse Racing"
	},
	"1F3C7-1F3FB": {
		"key": "1F3C7-1F3FB",
		"value": "🏇🏻",
		"descriptor": "Horse Racing: Light Skin Tone"
	},
	"1F3C7-1F3FC": {
		"key": "1F3C7-1F3FC",
		"value": "🏇🏼",
		"descriptor": "Horse Racing: Medium-Light Skin Tone"
	},
	"1F3C7-1F3FD": {
		"key": "1F3C7-1F3FD",
		"value": "🏇🏽",
		"descriptor": "Horse Racing: Medium Skin Tone"
	},
	"1F3C7-1F3FE": {
		"key": "1F3C7-1F3FE",
		"value": "🏇🏾",
		"descriptor": "Horse Racing: Medium-Dark Skin Tone"
	},
	"1F3C7-1F3FF": {
		"key": "1F3C7-1F3FF",
		"value": "🏇🏿",
		"descriptor": "Horse Racing: Dark Skin Tone"
	},
	"1F3C8": {
		"key": "1F3C8",
		"value": "🏈",
		"descriptor": "American Football"
	},
	"1F3C9": {
		"key": "1F3C9",
		"value": "🏉",
		"descriptor": "Rugby Football"
	},
	"1F3CA": {
		"key": "1F3CA",
		"value": "🏊",
		"descriptor": "Person Swimming"
	},
	"1F3CA-1F3FB": {
		"key": "1F3CA-1F3FB",
		"value": "🏊🏻",
		"descriptor": "Person Swimming: Light Skin Tone"
	},
	"1F3CA-1F3FB-200D-2640-FE0F": {
		"key": "1F3CA-1F3FB-200D-2640-FE0F",
		"value": "🏊🏻‍♀️",
		"descriptor": "Woman Swimming: Light Skin Tone"
	},
	"1F3CA-1F3FB-200D-2642-FE0F": {
		"key": "1F3CA-1F3FB-200D-2642-FE0F",
		"value": "🏊🏻‍♂️",
		"descriptor": "Man Swimming: Light Skin Tone"
	},
	"1F3CA-1F3FC": {
		"key": "1F3CA-1F3FC",
		"value": "🏊🏼",
		"descriptor": "Person Swimming: Medium-Light Skin Tone"
	},
	"1F3CA-1F3FC-200D-2640-FE0F": {
		"key": "1F3CA-1F3FC-200D-2640-FE0F",
		"value": "🏊🏼‍♀️",
		"descriptor": "Woman Swimming: Medium-Light Skin Tone"
	},
	"1F3CA-1F3FC-200D-2642-FE0F": {
		"key": "1F3CA-1F3FC-200D-2642-FE0F",
		"value": "🏊🏼‍♂️",
		"descriptor": "Man Swimming: Medium-Light Skin Tone"
	},
	"1F3CA-1F3FD": {
		"key": "1F3CA-1F3FD",
		"value": "🏊🏽",
		"descriptor": "Person Swimming: Medium Skin Tone"
	},
	"1F3CA-1F3FD-200D-2640-FE0F": {
		"key": "1F3CA-1F3FD-200D-2640-FE0F",
		"value": "🏊🏽‍♀️",
		"descriptor": "Woman Swimming: Medium Skin Tone"
	},
	"1F3CA-1F3FD-200D-2642-FE0F": {
		"key": "1F3CA-1F3FD-200D-2642-FE0F",
		"value": "🏊🏽‍♂️",
		"descriptor": "Man Swimming: Medium Skin Tone"
	},
	"1F3CA-1F3FE": {
		"key": "1F3CA-1F3FE",
		"value": "🏊🏾",
		"descriptor": "Person Swimming: Medium-Dark Skin Tone"
	},
	"1F3CA-1F3FE-200D-2640-FE0F": {
		"key": "1F3CA-1F3FE-200D-2640-FE0F",
		"value": "🏊🏾‍♀️",
		"descriptor": "Woman Swimming: Medium-Dark Skin Tone"
	},
	"1F3CA-1F3FE-200D-2642-FE0F": {
		"key": "1F3CA-1F3FE-200D-2642-FE0F",
		"value": "🏊🏾‍♂️",
		"descriptor": "Man Swimming: Medium-Dark Skin Tone"
	},
	"1F3CA-1F3FF": {
		"key": "1F3CA-1F3FF",
		"value": "🏊🏿",
		"descriptor": "Person Swimming: Dark Skin Tone"
	},
	"1F3CA-1F3FF-200D-2640-FE0F": {
		"key": "1F3CA-1F3FF-200D-2640-FE0F",
		"value": "🏊🏿‍♀️",
		"descriptor": "Woman Swimming: Dark Skin Tone"
	},
	"1F3CA-1F3FF-200D-2642-FE0F": {
		"key": "1F3CA-1F3FF-200D-2642-FE0F",
		"value": "🏊🏿‍♂️",
		"descriptor": "Man Swimming: Dark Skin Tone"
	},
	"1F3CA-200D-2640-FE0F": {
		"key": "1F3CA-200D-2640-FE0F",
		"value": "🏊‍♀️",
		"descriptor": "Woman Swimming"
	},
	"1F3CA-200D-2642-FE0F": {
		"key": "1F3CA-200D-2642-FE0F",
		"value": "🏊‍♂️",
		"descriptor": "Man Swimming"
	},
	"1F3CB-1F3FB": {
		"key": "1F3CB-1F3FB",
		"value": "🏋🏻",
		"descriptor": "Person Lifting Weights: Light Skin Tone"
	},
	"1F3CB-1F3FB-200D-2640-FE0F": {
		"key": "1F3CB-1F3FB-200D-2640-FE0F",
		"value": "🏋🏻‍♀️",
		"descriptor": "Woman Lifting Weights: Light Skin Tone"
	},
	"1F3CB-1F3FB-200D-2642-FE0F": {
		"key": "1F3CB-1F3FB-200D-2642-FE0F",
		"value": "🏋🏻‍♂️",
		"descriptor": "Man Lifting Weights: Light Skin Tone"
	},
	"1F3CB-1F3FC": {
		"key": "1F3CB-1F3FC",
		"value": "🏋🏼",
		"descriptor": "Person Lifting Weights: Medium-Light Skin Tone"
	},
	"1F3CB-1F3FC-200D-2640-FE0F": {
		"key": "1F3CB-1F3FC-200D-2640-FE0F",
		"value": "🏋🏼‍♀️",
		"descriptor": "Woman Lifting Weights: Medium-Light Skin Tone"
	},
	"1F3CB-1F3FC-200D-2642-FE0F": {
		"key": "1F3CB-1F3FC-200D-2642-FE0F",
		"value": "🏋🏼‍♂️",
		"descriptor": "Man Lifting Weights: Medium-Light Skin Tone"
	},
	"1F3CB-1F3FD": {
		"key": "1F3CB-1F3FD",
		"value": "🏋🏽",
		"descriptor": "Person Lifting Weights: Medium Skin Tone"
	},
	"1F3CB-1F3FD-200D-2640-FE0F": {
		"key": "1F3CB-1F3FD-200D-2640-FE0F",
		"value": "🏋🏽‍♀️",
		"descriptor": "Woman Lifting Weights: Medium Skin Tone"
	},
	"1F3CB-1F3FD-200D-2642-FE0F": {
		"key": "1F3CB-1F3FD-200D-2642-FE0F",
		"value": "🏋🏽‍♂️",
		"descriptor": "Man Lifting Weights: Medium Skin Tone"
	},
	"1F3CB-1F3FE": {
		"key": "1F3CB-1F3FE",
		"value": "🏋🏾",
		"descriptor": "Person Lifting Weights: Medium-Dark Skin Tone"
	},
	"1F3CB-1F3FE-200D-2640-FE0F": {
		"key": "1F3CB-1F3FE-200D-2640-FE0F",
		"value": "🏋🏾‍♀️",
		"descriptor": "Woman Lifting Weights: Medium-Dark Skin Tone"
	},
	"1F3CB-1F3FE-200D-2642-FE0F": {
		"key": "1F3CB-1F3FE-200D-2642-FE0F",
		"value": "🏋🏾‍♂️",
		"descriptor": "Man Lifting Weights: Medium-Dark Skin Tone"
	},
	"1F3CB-1F3FF": {
		"key": "1F3CB-1F3FF",
		"value": "🏋🏿",
		"descriptor": "Person Lifting Weights: Dark Skin Tone"
	},
	"1F3CB-1F3FF-200D-2640-FE0F": {
		"key": "1F3CB-1F3FF-200D-2640-FE0F",
		"value": "🏋🏿‍♀️",
		"descriptor": "Woman Lifting Weights: Dark Skin Tone"
	},
	"1F3CB-1F3FF-200D-2642-FE0F": {
		"key": "1F3CB-1F3FF-200D-2642-FE0F",
		"value": "🏋🏿‍♂️",
		"descriptor": "Man Lifting Weights: Dark Skin Tone"
	},
	"1F3CB-FE0F": {
		"key": "1F3CB-FE0F",
		"value": "🏋️",
		"descriptor": "Person Lifting Weights"
	},
	"1F3CB-FE0F-200D-2640-FE0F": {
		"key": "1F3CB-FE0F-200D-2640-FE0F",
		"value": "🏋️‍♀️",
		"descriptor": "Woman Lifting Weights"
	},
	"1F3CB-FE0F-200D-2642-FE0F": {
		"key": "1F3CB-FE0F-200D-2642-FE0F",
		"value": "🏋️‍♂️",
		"descriptor": "Man Lifting Weights"
	},
	"1F3CC-1F3FB": {
		"key": "1F3CC-1F3FB",
		"value": "🏌🏻",
		"descriptor": "Person Golfing: Light Skin Tone"
	},
	"1F3CC-1F3FB-200D-2640-FE0F": {
		"key": "1F3CC-1F3FB-200D-2640-FE0F",
		"value": "🏌🏻‍♀️",
		"descriptor": "Woman Golfing: Light Skin Tone"
	},
	"1F3CC-1F3FB-200D-2642-FE0F": {
		"key": "1F3CC-1F3FB-200D-2642-FE0F",
		"value": "🏌🏻‍♂️",
		"descriptor": "Man Golfing: Light Skin Tone"
	},
	"1F3CC-1F3FC": {
		"key": "1F3CC-1F3FC",
		"value": "🏌🏼",
		"descriptor": "Person Golfing: Medium-Light Skin Tone"
	},
	"1F3CC-1F3FC-200D-2640-FE0F": {
		"key": "1F3CC-1F3FC-200D-2640-FE0F",
		"value": "🏌🏼‍♀️",
		"descriptor": "Woman Golfing: Medium-Light Skin Tone"
	},
	"1F3CC-1F3FC-200D-2642-FE0F": {
		"key": "1F3CC-1F3FC-200D-2642-FE0F",
		"value": "🏌🏼‍♂️",
		"descriptor": "Man Golfing: Medium-Light Skin Tone"
	},
	"1F3CC-1F3FD": {
		"key": "1F3CC-1F3FD",
		"value": "🏌🏽",
		"descriptor": "Person Golfing: Medium Skin Tone"
	},
	"1F3CC-1F3FD-200D-2640-FE0F": {
		"key": "1F3CC-1F3FD-200D-2640-FE0F",
		"value": "🏌🏽‍♀️",
		"descriptor": "Woman Golfing: Medium Skin Tone"
	},
	"1F3CC-1F3FD-200D-2642-FE0F": {
		"key": "1F3CC-1F3FD-200D-2642-FE0F",
		"value": "🏌🏽‍♂️",
		"descriptor": "Man Golfing: Medium Skin Tone"
	},
	"1F3CC-1F3FE": {
		"key": "1F3CC-1F3FE",
		"value": "🏌🏾",
		"descriptor": "Person Golfing: Medium-Dark Skin Tone"
	},
	"1F3CC-1F3FE-200D-2640-FE0F": {
		"key": "1F3CC-1F3FE-200D-2640-FE0F",
		"value": "🏌🏾‍♀️",
		"descriptor": "Woman Golfing: Medium-Dark Skin Tone"
	},
	"1F3CC-1F3FE-200D-2642-FE0F": {
		"key": "1F3CC-1F3FE-200D-2642-FE0F",
		"value": "🏌🏾‍♂️",
		"descriptor": "Man Golfing: Medium-Dark Skin Tone"
	},
	"1F3CC-1F3FF": {
		"key": "1F3CC-1F3FF",
		"value": "🏌🏿",
		"descriptor": "Person Golfing: Dark Skin Tone"
	},
	"1F3CC-1F3FF-200D-2640-FE0F": {
		"key": "1F3CC-1F3FF-200D-2640-FE0F",
		"value": "🏌🏿‍♀️",
		"descriptor": "Woman Golfing: Dark Skin Tone"
	},
	"1F3CC-1F3FF-200D-2642-FE0F": {
		"key": "1F3CC-1F3FF-200D-2642-FE0F",
		"value": "🏌🏿‍♂️",
		"descriptor": "Man Golfing: Dark Skin Tone"
	},
	"1F3CC-FE0F": {
		"key": "1F3CC-FE0F",
		"value": "🏌️",
		"descriptor": "Person Golfing"
	},
	"1F3CC-FE0F-200D-2640-FE0F": {
		"key": "1F3CC-FE0F-200D-2640-FE0F",
		"value": "🏌️‍♀️",
		"descriptor": "Woman Golfing"
	},
	"1F3CC-FE0F-200D-2642-FE0F": {
		"key": "1F3CC-FE0F-200D-2642-FE0F",
		"value": "🏌️‍♂️",
		"descriptor": "Man Golfing"
	},
	"1F3CD-FE0F": {
		"key": "1F3CD-FE0F",
		"value": "🏍️",
		"descriptor": "Motorcycle"
	},
	"1F3CE-FE0F": {
		"key": "1F3CE-FE0F",
		"value": "🏎️",
		"descriptor": "Racing Car"
	},
	"1F3CF": {
		"key": "1F3CF",
		"value": "🏏",
		"descriptor": "Cricket Game"
	},
	"1F3D0": {
		"key": "1F3D0",
		"value": "🏐",
		"descriptor": "Volleyball"
	},
	"1F3D1": {
		"key": "1F3D1",
		"value": "🏑",
		"descriptor": "Field Hockey"
	},
	"1F3D2": {
		"key": "1F3D2",
		"value": "🏒",
		"descriptor": "Ice Hockey"
	},
	"1F3D3": {
		"key": "1F3D3",
		"value": "🏓",
		"descriptor": "Ping Pong"
	},
	"1F3D4-FE0F": {
		"key": "1F3D4-FE0F",
		"value": "🏔️",
		"descriptor": "Snow-Capped Mountain"
	},
	"1F3D5-FE0F": {
		"key": "1F3D5-FE0F",
		"value": "🏕️",
		"descriptor": "Camping"
	},
	"1F3D6-FE0F": {
		"key": "1F3D6-FE0F",
		"value": "🏖️",
		"descriptor": "Beach With Umbrella"
	},
	"1F3D7-FE0F": {
		"key": "1F3D7-FE0F",
		"value": "🏗️",
		"descriptor": "Building Construction"
	},
	"1F3D8-FE0F": {
		"key": "1F3D8-FE0F",
		"value": "🏘️",
		"descriptor": "Houses"
	},
	"1F3D9-FE0F": {
		"key": "1F3D9-FE0F",
		"value": "🏙️",
		"descriptor": "Cityscape"
	},
	"1F3DA-FE0F": {
		"key": "1F3DA-FE0F",
		"value": "🏚️",
		"descriptor": "Derelict House"
	},
	"1F3DB-FE0F": {
		"key": "1F3DB-FE0F",
		"value": "🏛️",
		"descriptor": "Classical Building"
	},
	"1F3DC-FE0F": {
		"key": "1F3DC-FE0F",
		"value": "🏜️",
		"descriptor": "Desert"
	},
	"1F3DD-FE0F": {
		"key": "1F3DD-FE0F",
		"value": "🏝️",
		"descriptor": "Desert Island"
	},
	"1F3DE-FE0F": {
		"key": "1F3DE-FE0F",
		"value": "🏞️",
		"descriptor": "National Park"
	},
	"1F3DF-FE0F": {
		"key": "1F3DF-FE0F",
		"value": "🏟️",
		"descriptor": "Stadium"
	},
	"1F3E0": {
		"key": "1F3E0",
		"value": "🏠",
		"descriptor": "House"
	},
	"1F3E1": {
		"key": "1F3E1",
		"value": "🏡",
		"descriptor": "House With Garden"
	},
	"1F3E2": {
		"key": "1F3E2",
		"value": "🏢",
		"descriptor": "Office Building"
	},
	"1F3E3": {
		"key": "1F3E3",
		"value": "🏣",
		"descriptor": "Japanese Post Office"
	},
	"1F3E4": {
		"key": "1F3E4",
		"value": "🏤",
		"descriptor": "Post Office"
	},
	"1F3E5": {
		"key": "1F3E5",
		"value": "🏥",
		"descriptor": "Hospital"
	},
	"1F3E6": {
		"key": "1F3E6",
		"value": "🏦",
		"descriptor": "Bank"
	},
	"1F3E7": {
		"key": "1F3E7",
		"value": "🏧",
		"descriptor": "ATM Sign"
	},
	"1F3E8": {
		"key": "1F3E8",
		"value": "🏨",
		"descriptor": "Hotel"
	},
	"1F3E9": {
		"key": "1F3E9",
		"value": "🏩",
		"descriptor": "Love Hotel"
	},
	"1F3EA": {
		"key": "1F3EA",
		"value": "🏪",
		"descriptor": "Convenience Store"
	},
	"1F3EB": {
		"key": "1F3EB",
		"value": "🏫",
		"descriptor": "School"
	},
	"1F3EC": {
		"key": "1F3EC",
		"value": "🏬",
		"descriptor": "Department Store"
	},
	"1F3ED": {
		"key": "1F3ED",
		"value": "🏭",
		"descriptor": "Factory"
	},
	"1F3EE": {
		"key": "1F3EE",
		"value": "🏮",
		"descriptor": "Red Paper Lantern"
	},
	"1F3EF": {
		"key": "1F3EF",
		"value": "🏯",
		"descriptor": "Japanese Castle"
	},
	"1F3F0": {
		"key": "1F3F0",
		"value": "🏰",
		"descriptor": "Castle"
	},
	"1F3F3-FE0F": {
		"key": "1F3F3-FE0F",
		"value": "🏳️",
		"descriptor": "White Flag"
	},
	"1F3F3-FE0F-200D-1F308": {
		"key": "1F3F3-FE0F-200D-1F308",
		"value": "🏳️‍🌈",
		"descriptor": "Rainbow Flag"
	},
	"1F3F4": {
		"key": "1F3F4",
		"value": "🏴",
		"descriptor": "Black Flag"
	},
	"1F3F4-200D-2620-FE0F": {
		"key": "1F3F4-200D-2620-FE0F",
		"value": "🏴‍☠️",
		"descriptor": "Pirate Flag"
	},
	"1F3F4-E0067-E0062-E0065-E006E-E0067-E007F": {
		"key": "1F3F4-E0067-E0062-E0065-E006E-E0067-E007F",
		"value": "🏴󠁧󠁢󠁥󠁮󠁧󠁿",
		"descriptor": "Flag: England"
	},
	"1F3F4-E0067-E0062-E0073-E0063-E0074-E007F": {
		"key": "1F3F4-E0067-E0062-E0073-E0063-E0074-E007F",
		"value": "🏴󠁧󠁢󠁳󠁣󠁴󠁿",
		"descriptor": "Flag: Scotland"
	},
	"1F3F4-E0067-E0062-E0077-E006C-E0073-E007F": {
		"key": "1F3F4-E0067-E0062-E0077-E006C-E0073-E007F",
		"value": "🏴󠁧󠁢󠁷󠁬󠁳󠁿",
		"descriptor": "Flag: Wales"
	},
	"1F3F5-FE0F": {
		"key": "1F3F5-FE0F",
		"value": "🏵️",
		"descriptor": "Rosette"
	},
	"1F3F7-FE0F": {
		"key": "1F3F7-FE0F",
		"value": "🏷️",
		"descriptor": "Label"
	},
	"1F3F8": {
		"key": "1F3F8",
		"value": "🏸",
		"descriptor": "Badminton"
	},
	"1F3F9": {
		"key": "1F3F9",
		"value": "🏹",
		"descriptor": "Bow and Arrow"
	},
	"1F3FA": {
		"key": "1F3FA",
		"value": "🏺",
		"descriptor": "Amphora"
	},
	"1F3FB": {
		"key": "1F3FB",
		"value": "🏻",
		"descriptor": "Light Skin Tone"
	},
	"1F3FC": {
		"key": "1F3FC",
		"value": "🏼",
		"descriptor": "Medium-Light Skin Tone"
	},
	"1F3FD": {
		"key": "1F3FD",
		"value": "🏽",
		"descriptor": "Medium Skin Tone"
	},
	"1F3FE": {
		"key": "1F3FE",
		"value": "🏾",
		"descriptor": "Medium-Dark Skin Tone"
	},
	"1F3FF": {
		"key": "1F3FF",
		"value": "🏿",
		"descriptor": "Dark Skin Tone"
	},
	"1F400": {
		"key": "1F400",
		"value": "🐀",
		"descriptor": "Rat"
	},
	"1F401": {
		"key": "1F401",
		"value": "🐁",
		"descriptor": "Mouse"
	},
	"1F402": {
		"key": "1F402",
		"value": "🐂",
		"descriptor": "Ox"
	},
	"1F403": {
		"key": "1F403",
		"value": "🐃",
		"descriptor": "Water Buffalo"
	},
	"1F404": {
		"key": "1F404",
		"value": "🐄",
		"descriptor": "Cow"
	},
	"1F405": {
		"key": "1F405",
		"value": "🐅",
		"descriptor": "Tiger"
	},
	"1F406": {
		"key": "1F406",
		"value": "🐆",
		"descriptor": "Leopard"
	},
	"1F407": {
		"key": "1F407",
		"value": "🐇",
		"descriptor": "Rabbit"
	},
	"1F408": {
		"key": "1F408",
		"value": "🐈",
		"descriptor": "Cat"
	},
	"1F409": {
		"key": "1F409",
		"value": "🐉",
		"descriptor": "Dragon"
	},
	"1F40A": {
		"key": "1F40A",
		"value": "🐊",
		"descriptor": "Crocodile"
	},
	"1F40B": {
		"key": "1F40B",
		"value": "🐋",
		"descriptor": "Whale"
	},
	"1F40C": {
		"key": "1F40C",
		"value": "🐌",
		"descriptor": "Snail"
	},
	"1F40D": {
		"key": "1F40D",
		"value": "🐍",
		"descriptor": "Snake"
	},
	"1F40E": {
		"key": "1F40E",
		"value": "🐎",
		"descriptor": "Horse"
	},
	"1F40F": {
		"key": "1F40F",
		"value": "🐏",
		"descriptor": "Ram"
	},
	"1F410": {
		"key": "1F410",
		"value": "🐐",
		"descriptor": "Goat"
	},
	"1F411": {
		"key": "1F411",
		"value": "🐑",
		"descriptor": "Ewe"
	},
	"1F412": {
		"key": "1F412",
		"value": "🐒",
		"descriptor": "Monkey"
	},
	"1F413": {
		"key": "1F413",
		"value": "🐓",
		"descriptor": "Rooster"
	},
	"1F414": {
		"key": "1F414",
		"value": "🐔",
		"descriptor": "Chicken"
	},
	"1F415": {
		"key": "1F415",
		"value": "🐕",
		"descriptor": "Dog"
	},
	"1F415-200D-1F9BA": {
		"key": "1F415-200D-1F9BA",
		"value": "🐕‍🦺",
		"descriptor": "Service Dog"
	},
	"1F416": {
		"key": "1F416",
		"value": "🐖",
		"descriptor": "Pig"
	},
	"1F417": {
		"key": "1F417",
		"value": "🐗",
		"descriptor": "Boar"
	},
	"1F418": {
		"key": "1F418",
		"value": "🐘",
		"descriptor": "Elephant"
	},
	"1F419": {
		"key": "1F419",
		"value": "🐙",
		"descriptor": "Octopus"
	},
	"1F41A": {
		"key": "1F41A",
		"value": "🐚",
		"descriptor": "Spiral Shell"
	},
	"1F41B": {
		"key": "1F41B",
		"value": "🐛",
		"descriptor": "Bug"
	},
	"1F41C": {
		"key": "1F41C",
		"value": "🐜",
		"descriptor": "Ant"
	},
	"1F41D": {
		"key": "1F41D",
		"value": "🐝",
		"descriptor": "Honeybee"
	},
	"1F41E": {
		"key": "1F41E",
		"value": "🐞",
		"descriptor": "Lady Beetle"
	},
	"1F41F": {
		"key": "1F41F",
		"value": "🐟",
		"descriptor": "Fish"
	},
	"1F420": {
		"key": "1F420",
		"value": "🐠",
		"descriptor": "Tropical Fish"
	},
	"1F421": {
		"key": "1F421",
		"value": "🐡",
		"descriptor": "Blowfish"
	},
	"1F422": {
		"key": "1F422",
		"value": "🐢",
		"descriptor": "Turtle"
	},
	"1F423": {
		"key": "1F423",
		"value": "🐣",
		"descriptor": "Hatching Chick"
	},
	"1F424": {
		"key": "1F424",
		"value": "🐤",
		"descriptor": "Baby Chick"
	},
	"1F425": {
		"key": "1F425",
		"value": "🐥",
		"descriptor": "Front-Facing Baby Chick"
	},
	"1F426": {
		"key": "1F426",
		"value": "🐦",
		"descriptor": "Bird"
	},
	"1F427": {
		"key": "1F427",
		"value": "🐧",
		"descriptor": "Penguin"
	},
	"1F428": {
		"key": "1F428",
		"value": "🐨",
		"descriptor": "Koala"
	},
	"1F429": {
		"key": "1F429",
		"value": "🐩",
		"descriptor": "Poodle"
	},
	"1F42A": {
		"key": "1F42A",
		"value": "🐪",
		"descriptor": "Camel"
	},
	"1F42B": {
		"key": "1F42B",
		"value": "🐫",
		"descriptor": "Two-Hump Camel"
	},
	"1F42C": {
		"key": "1F42C",
		"value": "🐬",
		"descriptor": "Dolphin"
	},
	"1F42D": {
		"key": "1F42D",
		"value": "🐭",
		"descriptor": "Mouse Face"
	},
	"1F42E": {
		"key": "1F42E",
		"value": "🐮",
		"descriptor": "Cow Face"
	},
	"1F42F": {
		"key": "1F42F",
		"value": "🐯",
		"descriptor": "Tiger Face"
	},
	"1F430": {
		"key": "1F430",
		"value": "🐰",
		"descriptor": "Rabbit Face"
	},
	"1F431": {
		"key": "1F431",
		"value": "🐱",
		"descriptor": "Cat Face"
	},
	"1F432": {
		"key": "1F432",
		"value": "🐲",
		"descriptor": "Dragon Face"
	},
	"1F433": {
		"key": "1F433",
		"value": "🐳",
		"descriptor": "Spouting Whale"
	},
	"1F434": {
		"key": "1F434",
		"value": "🐴",
		"descriptor": "Horse Face"
	},
	"1F435": {
		"key": "1F435",
		"value": "🐵",
		"descriptor": "Monkey Face"
	},
	"1F436": {
		"key": "1F436",
		"value": "🐶",
		"descriptor": "Dog Face"
	},
	"1F437": {
		"key": "1F437",
		"value": "🐷",
		"descriptor": "Pig Face"
	},
	"1F438": {
		"key": "1F438",
		"value": "🐸",
		"descriptor": "Frog"
	},
	"1F439": {
		"key": "1F439",
		"value": "🐹",
		"descriptor": "Hamster"
	},
	"1F43A": {
		"key": "1F43A",
		"value": "🐺",
		"descriptor": "Wolf"
	},
	"1F43B": {
		"key": "1F43B",
		"value": "🐻",
		"descriptor": "Bear"
	},
	"1F43C": {
		"key": "1F43C",
		"value": "🐼",
		"descriptor": "Panda"
	},
	"1F43D": {
		"key": "1F43D",
		"value": "🐽",
		"descriptor": "Pig Nose"
	},
	"1F43E": {
		"key": "1F43E",
		"value": "🐾",
		"descriptor": "Paw Prints"
	},
	"1F43F-FE0F": {
		"key": "1F43F-FE0F",
		"value": "🐿️",
		"descriptor": "Chipmunk"
	},
	"1F440": {
		"key": "1F440",
		"value": "👀",
		"descriptor": "Eyes"
	},
	"1F441-FE0F": {
		"key": "1F441-FE0F",
		"value": "👁️",
		"descriptor": "Eye"
	},
	"1F441-FE0F-200D-1F5E8-FE0F": {
		"key": "1F441-FE0F-200D-1F5E8-FE0F",
		"value": "👁️‍🗨️",
		"descriptor": "Eye in Speech Bubble"
	},
	"1F442": {
		"key": "1F442",
		"value": "👂",
		"descriptor": "Ear"
	},
	"1F442-1F3FB": {
		"key": "1F442-1F3FB",
		"value": "👂🏻",
		"descriptor": "Ear: Light Skin Tone"
	},
	"1F442-1F3FC": {
		"key": "1F442-1F3FC",
		"value": "👂🏼",
		"descriptor": "Ear: Medium-Light Skin Tone"
	},
	"1F442-1F3FD": {
		"key": "1F442-1F3FD",
		"value": "👂🏽",
		"descriptor": "Ear: Medium Skin Tone"
	},
	"1F442-1F3FE": {
		"key": "1F442-1F3FE",
		"value": "👂🏾",
		"descriptor": "Ear: Medium-Dark Skin Tone"
	},
	"1F442-1F3FF": {
		"key": "1F442-1F3FF",
		"value": "👂🏿",
		"descriptor": "Ear: Dark Skin Tone"
	},
	"1F443": {
		"key": "1F443",
		"value": "👃",
		"descriptor": "Nose"
	},
	"1F443-1F3FB": {
		"key": "1F443-1F3FB",
		"value": "👃🏻",
		"descriptor": "Nose: Light Skin Tone"
	},
	"1F443-1F3FC": {
		"key": "1F443-1F3FC",
		"value": "👃🏼",
		"descriptor": "Nose: Medium-Light Skin Tone"
	},
	"1F443-1F3FD": {
		"key": "1F443-1F3FD",
		"value": "👃🏽",
		"descriptor": "Nose: Medium Skin Tone"
	},
	"1F443-1F3FE": {
		"key": "1F443-1F3FE",
		"value": "👃🏾",
		"descriptor": "Nose: Medium-Dark Skin Tone"
	},
	"1F443-1F3FF": {
		"key": "1F443-1F3FF",
		"value": "👃🏿",
		"descriptor": "Nose: Dark Skin Tone"
	},
	"1F444": {
		"key": "1F444",
		"value": "👄",
		"descriptor": "Mouth"
	},
	"1F445": {
		"key": "1F445",
		"value": "👅",
		"descriptor": "Tongue"
	},
	"1F446": {
		"key": "1F446",
		"value": "👆",
		"descriptor": "Backhand Index Pointing Up"
	},
	"1F446-1F3FB": {
		"key": "1F446-1F3FB",
		"value": "👆🏻",
		"descriptor": "Backhand Index Pointing Up: Light Skin Tone"
	},
	"1F446-1F3FC": {
		"key": "1F446-1F3FC",
		"value": "👆🏼",
		"descriptor": "Backhand Index Pointing Up: Medium-Light Skin Tone"
	},
	"1F446-1F3FD": {
		"key": "1F446-1F3FD",
		"value": "👆🏽",
		"descriptor": "Backhand Index Pointing Up: Medium Skin Tone"
	},
	"1F446-1F3FE": {
		"key": "1F446-1F3FE",
		"value": "👆🏾",
		"descriptor": "Backhand Index Pointing Up: Medium-Dark Skin Tone"
	},
	"1F446-1F3FF": {
		"key": "1F446-1F3FF",
		"value": "👆🏿",
		"descriptor": "Backhand Index Pointing Up: Dark Skin Tone"
	},
	"1F447": {
		"key": "1F447",
		"value": "👇",
		"descriptor": "Backhand Index Pointing Down"
	},
	"1F447-1F3FB": {
		"key": "1F447-1F3FB",
		"value": "👇🏻",
		"descriptor": "Backhand Index Pointing Down: Light Skin Tone"
	},
	"1F447-1F3FC": {
		"key": "1F447-1F3FC",
		"value": "👇🏼",
		"descriptor": "Backhand Index Pointing Down: Medium-Light Skin Tone"
	},
	"1F447-1F3FD": {
		"key": "1F447-1F3FD",
		"value": "👇🏽",
		"descriptor": "Backhand Index Pointing Down: Medium Skin Tone"
	},
	"1F447-1F3FE": {
		"key": "1F447-1F3FE",
		"value": "👇🏾",
		"descriptor": "Backhand Index Pointing Down: Medium-Dark Skin Tone"
	},
	"1F447-1F3FF": {
		"key": "1F447-1F3FF",
		"value": "👇🏿",
		"descriptor": "Backhand Index Pointing Down: Dark Skin Tone"
	},
	"1F448": {
		"key": "1F448",
		"value": "👈",
		"descriptor": "Backhand Index Pointing Left"
	},
	"1F448-1F3FB": {
		"key": "1F448-1F3FB",
		"value": "👈🏻",
		"descriptor": "Backhand Index Pointing Left: Light Skin Tone"
	},
	"1F448-1F3FC": {
		"key": "1F448-1F3FC",
		"value": "👈🏼",
		"descriptor": "Backhand Index Pointing Left: Medium-Light Skin Tone"
	},
	"1F448-1F3FD": {
		"key": "1F448-1F3FD",
		"value": "👈🏽",
		"descriptor": "Backhand Index Pointing Left: Medium Skin Tone"
	},
	"1F448-1F3FE": {
		"key": "1F448-1F3FE",
		"value": "👈🏾",
		"descriptor": "Backhand Index Pointing Left: Medium-Dark Skin Tone"
	},
	"1F448-1F3FF": {
		"key": "1F448-1F3FF",
		"value": "👈🏿",
		"descriptor": "Backhand Index Pointing Left: Dark Skin Tone"
	},
	"1F449": {
		"key": "1F449",
		"value": "👉",
		"descriptor": "Backhand Index Pointing Right"
	},
	"1F449-1F3FB": {
		"key": "1F449-1F3FB",
		"value": "👉🏻",
		"descriptor": "Backhand Index Pointing Right: Light Skin Tone"
	},
	"1F449-1F3FC": {
		"key": "1F449-1F3FC",
		"value": "👉🏼",
		"descriptor": "Backhand Index Pointing Right: Medium-Light Skin Tone"
	},
	"1F449-1F3FD": {
		"key": "1F449-1F3FD",
		"value": "👉🏽",
		"descriptor": "Backhand Index Pointing Right: Medium Skin Tone"
	},
	"1F449-1F3FE": {
		"key": "1F449-1F3FE",
		"value": "👉🏾",
		"descriptor": "Backhand Index Pointing Right: Medium-Dark Skin Tone"
	},
	"1F449-1F3FF": {
		"key": "1F449-1F3FF",
		"value": "👉🏿",
		"descriptor": "Backhand Index Pointing Right: Dark Skin Tone"
	},
	"1F44A": {
		"key": "1F44A",
		"value": "👊",
		"descriptor": "Oncoming Fist"
	},
	"1F44A-1F3FB": {
		"key": "1F44A-1F3FB",
		"value": "👊🏻",
		"descriptor": "Oncoming Fist: Light Skin Tone"
	},
	"1F44A-1F3FC": {
		"key": "1F44A-1F3FC",
		"value": "👊🏼",
		"descriptor": "Oncoming Fist: Medium-Light Skin Tone"
	},
	"1F44A-1F3FD": {
		"key": "1F44A-1F3FD",
		"value": "👊🏽",
		"descriptor": "Oncoming Fist: Medium Skin Tone"
	},
	"1F44A-1F3FE": {
		"key": "1F44A-1F3FE",
		"value": "👊🏾",
		"descriptor": "Oncoming Fist: Medium-Dark Skin Tone"
	},
	"1F44A-1F3FF": {
		"key": "1F44A-1F3FF",
		"value": "👊🏿",
		"descriptor": "Oncoming Fist: Dark Skin Tone"
	},
	"1F44B": {
		"key": "1F44B",
		"value": "👋",
		"descriptor": "Waving Hand"
	},
	"1F44B-1F3FB": {
		"key": "1F44B-1F3FB",
		"value": "👋🏻",
		"descriptor": "Waving Hand: Light Skin Tone"
	},
	"1F44B-1F3FC": {
		"key": "1F44B-1F3FC",
		"value": "👋🏼",
		"descriptor": "Waving Hand: Medium-Light Skin Tone"
	},
	"1F44B-1F3FD": {
		"key": "1F44B-1F3FD",
		"value": "👋🏽",
		"descriptor": "Waving Hand: Medium Skin Tone"
	},
	"1F44B-1F3FE": {
		"key": "1F44B-1F3FE",
		"value": "👋🏾",
		"descriptor": "Waving Hand: Medium-Dark Skin Tone"
	},
	"1F44B-1F3FF": {
		"key": "1F44B-1F3FF",
		"value": "👋🏿",
		"descriptor": "Waving Hand: Dark Skin Tone"
	},
	"1F44C": {
		"key": "1F44C",
		"value": "👌",
		"descriptor": "OK Hand"
	},
	"1F44C-1F3FB": {
		"key": "1F44C-1F3FB",
		"value": "👌🏻",
		"descriptor": "OK Hand: Light Skin Tone"
	},
	"1F44C-1F3FC": {
		"key": "1F44C-1F3FC",
		"value": "👌🏼",
		"descriptor": "OK Hand: Medium-Light Skin Tone"
	},
	"1F44C-1F3FD": {
		"key": "1F44C-1F3FD",
		"value": "👌🏽",
		"descriptor": "OK Hand: Medium Skin Tone"
	},
	"1F44C-1F3FE": {
		"key": "1F44C-1F3FE",
		"value": "👌🏾",
		"descriptor": "OK Hand: Medium-Dark Skin Tone"
	},
	"1F44C-1F3FF": {
		"key": "1F44C-1F3FF",
		"value": "👌🏿",
		"descriptor": "OK Hand: Dark Skin Tone"
	},
	"1F44D": {
		"key": "1F44D",
		"value": "👍",
		"descriptor": "Thumbs Up"
	},
	"1F44D-1F3FB": {
		"key": "1F44D-1F3FB",
		"value": "👍🏻",
		"descriptor": "Thumbs Up: Light Skin Tone"
	},
	"1F44D-1F3FC": {
		"key": "1F44D-1F3FC",
		"value": "👍🏼",
		"descriptor": "Thumbs Up: Medium-Light Skin Tone"
	},
	"1F44D-1F3FD": {
		"key": "1F44D-1F3FD",
		"value": "👍🏽",
		"descriptor": "Thumbs Up: Medium Skin Tone"
	},
	"1F44D-1F3FE": {
		"key": "1F44D-1F3FE",
		"value": "👍🏾",
		"descriptor": "Thumbs Up: Medium-Dark Skin Tone"
	},
	"1F44D-1F3FF": {
		"key": "1F44D-1F3FF",
		"value": "👍🏿",
		"descriptor": "Thumbs Up: Dark Skin Tone"
	},
	"1F44E": {
		"key": "1F44E",
		"value": "👎",
		"descriptor": "Thumbs Down"
	},
	"1F44E-1F3FB": {
		"key": "1F44E-1F3FB",
		"value": "👎🏻",
		"descriptor": "Thumbs Down: Light Skin Tone"
	},
	"1F44E-1F3FC": {
		"key": "1F44E-1F3FC",
		"value": "👎🏼",
		"descriptor": "Thumbs Down: Medium-Light Skin Tone"
	},
	"1F44E-1F3FD": {
		"key": "1F44E-1F3FD",
		"value": "👎🏽",
		"descriptor": "Thumbs Down: Medium Skin Tone"
	},
	"1F44E-1F3FE": {
		"key": "1F44E-1F3FE",
		"value": "👎🏾",
		"descriptor": "Thumbs Down: Medium-Dark Skin Tone"
	},
	"1F44E-1F3FF": {
		"key": "1F44E-1F3FF",
		"value": "👎🏿",
		"descriptor": "Thumbs Down: Dark Skin Tone"
	},
	"1F44F": {
		"key": "1F44F",
		"value": "👏",
		"descriptor": "Clapping Hands"
	},
	"1F44F-1F3FB": {
		"key": "1F44F-1F3FB",
		"value": "👏🏻",
		"descriptor": "Clapping Hands: Light Skin Tone"
	},
	"1F44F-1F3FC": {
		"key": "1F44F-1F3FC",
		"value": "👏🏼",
		"descriptor": "Clapping Hands: Medium-Light Skin Tone"
	},
	"1F44F-1F3FD": {
		"key": "1F44F-1F3FD",
		"value": "👏🏽",
		"descriptor": "Clapping Hands: Medium Skin Tone"
	},
	"1F44F-1F3FE": {
		"key": "1F44F-1F3FE",
		"value": "👏🏾",
		"descriptor": "Clapping Hands: Medium-Dark Skin Tone"
	},
	"1F44F-1F3FF": {
		"key": "1F44F-1F3FF",
		"value": "👏🏿",
		"descriptor": "Clapping Hands: Dark Skin Tone"
	},
	"1F450": {
		"key": "1F450",
		"value": "👐",
		"descriptor": "Open Hands"
	},
	"1F450-1F3FB": {
		"key": "1F450-1F3FB",
		"value": "👐🏻",
		"descriptor": "Open Hands: Light Skin Tone"
	},
	"1F450-1F3FC": {
		"key": "1F450-1F3FC",
		"value": "👐🏼",
		"descriptor": "Open Hands: Medium-Light Skin Tone"
	},
	"1F450-1F3FD": {
		"key": "1F450-1F3FD",
		"value": "👐🏽",
		"descriptor": "Open Hands: Medium Skin Tone"
	},
	"1F450-1F3FE": {
		"key": "1F450-1F3FE",
		"value": "👐🏾",
		"descriptor": "Open Hands: Medium-Dark Skin Tone"
	},
	"1F450-1F3FF": {
		"key": "1F450-1F3FF",
		"value": "👐🏿",
		"descriptor": "Open Hands: Dark Skin Tone"
	},
	"1F451": {
		"key": "1F451",
		"value": "👑",
		"descriptor": "Crown"
	},
	"1F452": {
		"key": "1F452",
		"value": "👒",
		"descriptor": "Woman’s Hat"
	},
	"1F453": {
		"key": "1F453",
		"value": "👓",
		"descriptor": "Glasses"
	},
	"1F454": {
		"key": "1F454",
		"value": "👔",
		"descriptor": "Necktie"
	},
	"1F455": {
		"key": "1F455",
		"value": "👕",
		"descriptor": "T-Shirt"
	},
	"1F456": {
		"key": "1F456",
		"value": "👖",
		"descriptor": "Jeans"
	},
	"1F457": {
		"key": "1F457",
		"value": "👗",
		"descriptor": "Dress"
	},
	"1F458": {
		"key": "1F458",
		"value": "👘",
		"descriptor": "Kimono"
	},
	"1F459": {
		"key": "1F459",
		"value": "👙",
		"descriptor": "Bikini"
	},
	"1F45A": {
		"key": "1F45A",
		"value": "👚",
		"descriptor": "Woman’s Clothes"
	},
	"1F45B": {
		"key": "1F45B",
		"value": "👛",
		"descriptor": "Purse"
	},
	"1F45C": {
		"key": "1F45C",
		"value": "👜",
		"descriptor": "Handbag"
	},
	"1F45D": {
		"key": "1F45D",
		"value": "👝",
		"descriptor": "Clutch Bag"
	},
	"1F45E": {
		"key": "1F45E",
		"value": "👞",
		"descriptor": "Man’s Shoe"
	},
	"1F45F": {
		"key": "1F45F",
		"value": "👟",
		"descriptor": "Running Shoe"
	},
	"1F460": {
		"key": "1F460",
		"value": "👠",
		"descriptor": "High-Heeled Shoe"
	},
	"1F461": {
		"key": "1F461",
		"value": "👡",
		"descriptor": "Woman’s Sandal"
	},
	"1F462": {
		"key": "1F462",
		"value": "👢",
		"descriptor": "Woman’s Boot"
	},
	"1F463": {
		"key": "1F463",
		"value": "👣",
		"descriptor": "Footprints"
	},
	"1F464": {
		"key": "1F464",
		"value": "👤",
		"descriptor": "Bust in Silhouette"
	},
	"1F465": {
		"key": "1F465",
		"value": "👥",
		"descriptor": "Busts in Silhouette"
	},
	"1F466": {
		"key": "1F466",
		"value": "👦",
		"descriptor": "Boy"
	},
	"1F466-1F3FB": {
		"key": "1F466-1F3FB",
		"value": "👦🏻",
		"descriptor": "Boy: Light Skin Tone"
	},
	"1F466-1F3FC": {
		"key": "1F466-1F3FC",
		"value": "👦🏼",
		"descriptor": "Boy: Medium-Light Skin Tone"
	},
	"1F466-1F3FD": {
		"key": "1F466-1F3FD",
		"value": "👦🏽",
		"descriptor": "Boy: Medium Skin Tone"
	},
	"1F466-1F3FE": {
		"key": "1F466-1F3FE",
		"value": "👦🏾",
		"descriptor": "Boy: Medium-Dark Skin Tone"
	},
	"1F466-1F3FF": {
		"key": "1F466-1F3FF",
		"value": "👦🏿",
		"descriptor": "Boy: Dark Skin Tone"
	},
	"1F467": {
		"key": "1F467",
		"value": "👧",
		"descriptor": "Girl"
	},
	"1F467-1F3FB": {
		"key": "1F467-1F3FB",
		"value": "👧🏻",
		"descriptor": "Girl: Light Skin Tone"
	},
	"1F467-1F3FC": {
		"key": "1F467-1F3FC",
		"value": "👧🏼",
		"descriptor": "Girl: Medium-Light Skin Tone"
	},
	"1F467-1F3FD": {
		"key": "1F467-1F3FD",
		"value": "👧🏽",
		"descriptor": "Girl: Medium Skin Tone"
	},
	"1F467-1F3FE": {
		"key": "1F467-1F3FE",
		"value": "👧🏾",
		"descriptor": "Girl: Medium-Dark Skin Tone"
	},
	"1F467-1F3FF": {
		"key": "1F467-1F3FF",
		"value": "👧🏿",
		"descriptor": "Girl: Dark Skin Tone"
	},
	"1F468": {
		"key": "1F468",
		"value": "👨",
		"descriptor": "Man"
	},
	"1F468-1F3FB": {
		"key": "1F468-1F3FB",
		"value": "👨🏻",
		"descriptor": "Man: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F33E": {
		"key": "1F468-1F3FB-200D-1F33E",
		"value": "👨🏻‍🌾",
		"descriptor": "Man Farmer: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F373": {
		"key": "1F468-1F3FB-200D-1F373",
		"value": "👨🏻‍🍳",
		"descriptor": "Man Cook: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F393": {
		"key": "1F468-1F3FB-200D-1F393",
		"value": "👨🏻‍🎓",
		"descriptor": "Man Student: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F3A4": {
		"key": "1F468-1F3FB-200D-1F3A4",
		"value": "👨🏻‍🎤",
		"descriptor": "Man Singer: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F3A8": {
		"key": "1F468-1F3FB-200D-1F3A8",
		"value": "👨🏻‍🎨",
		"descriptor": "Man Artist: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F3EB": {
		"key": "1F468-1F3FB-200D-1F3EB",
		"value": "👨🏻‍🏫",
		"descriptor": "Man Teacher: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F3ED": {
		"key": "1F468-1F3FB-200D-1F3ED",
		"value": "👨🏻‍🏭",
		"descriptor": "Man Factory Worker: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F4BB": {
		"key": "1F468-1F3FB-200D-1F4BB",
		"value": "👨🏻‍💻",
		"descriptor": "Man Technologist: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F4BC": {
		"key": "1F468-1F3FB-200D-1F4BC",
		"value": "👨🏻‍💼",
		"descriptor": "Man Office Worker: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F527": {
		"key": "1F468-1F3FB-200D-1F527",
		"value": "👨🏻‍🔧",
		"descriptor": "Man Mechanic: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F52C": {
		"key": "1F468-1F3FB-200D-1F52C",
		"value": "👨🏻‍🔬",
		"descriptor": "Man Scientist: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F680": {
		"key": "1F468-1F3FB-200D-1F680",
		"value": "👨🏻‍🚀",
		"descriptor": "Man Astronaut: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F692": {
		"key": "1F468-1F3FB-200D-1F692",
		"value": "👨🏻‍🚒",
		"descriptor": "Man Firefighter: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FC": {
		"key": "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FC",
		"value": "👨🏻‍🤝‍👨🏼",
		"descriptor": "Men Holding Hands: Light Skin Tone, Medium-Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FD": {
		"key": "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FD",
		"value": "👨🏻‍🤝‍👨🏽",
		"descriptor": "Men Holding Hands: Light Skin Tone, Medium Skin Tone"
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FE": {
		"key": "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FE",
		"value": "👨🏻‍🤝‍👨🏾",
		"descriptor": "Men Holding Hands: Light Skin Tone, Medium-Dark Skin Tone"
	},
	"1F468-1F3FB-200D-1F91D-200D-1F468-1F3FF": {
		"key": "1F468-1F3FB-200D-1F91D-200D-1F468-1F3FF",
		"value": "👨🏻‍🤝‍👨🏿",
		"descriptor": "Men Holding Hands: Light Skin Tone, Dark Skin Tone"
	},
	"1F468-1F3FB-200D-1F9AF": {
		"key": "1F468-1F3FB-200D-1F9AF",
		"value": "👨🏻‍🦯",
		"descriptor": "Man With Probing Cane: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F9B0": {
		"key": "1F468-1F3FB-200D-1F9B0",
		"value": "👨🏻‍🦰",
		"descriptor": "Man: Light Skin Tone, Red Hair"
	},
	"1F468-1F3FB-200D-1F9B1": {
		"key": "1F468-1F3FB-200D-1F9B1",
		"value": "👨🏻‍🦱",
		"descriptor": "Man: Light Skin Tone, Curly Hair"
	},
	"1F468-1F3FB-200D-1F9B2": {
		"key": "1F468-1F3FB-200D-1F9B2",
		"value": "👨🏻‍🦲",
		"descriptor": "Man: Light Skin Tone, Bald"
	},
	"1F468-1F3FB-200D-1F9B3": {
		"key": "1F468-1F3FB-200D-1F9B3",
		"value": "👨🏻‍🦳",
		"descriptor": "Man: Light Skin Tone, White Hair"
	},
	"1F468-1F3FB-200D-1F9BC": {
		"key": "1F468-1F3FB-200D-1F9BC",
		"value": "👨🏻‍🦼",
		"descriptor": "Man in Motorized Wheelchair: Light Skin Tone"
	},
	"1F468-1F3FB-200D-1F9BD": {
		"key": "1F468-1F3FB-200D-1F9BD",
		"value": "👨🏻‍🦽",
		"descriptor": "Man in Manual Wheelchair: Light Skin Tone"
	},
	"1F468-1F3FB-200D-2695-FE0F": {
		"key": "1F468-1F3FB-200D-2695-FE0F",
		"value": "👨🏻‍⚕️",
		"descriptor": "Man Health Worker: Light Skin Tone"
	},
	"1F468-1F3FB-200D-2696-FE0F": {
		"key": "1F468-1F3FB-200D-2696-FE0F",
		"value": "👨🏻‍⚖️",
		"descriptor": "Man Judge: Light Skin Tone"
	},
	"1F468-1F3FB-200D-2708-FE0F": {
		"key": "1F468-1F3FB-200D-2708-FE0F",
		"value": "👨🏻‍✈️",
		"descriptor": "Man Pilot: Light Skin Tone"
	},
	"1F468-1F3FC": {
		"key": "1F468-1F3FC",
		"value": "👨🏼",
		"descriptor": "Man: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F33E": {
		"key": "1F468-1F3FC-200D-1F33E",
		"value": "👨🏼‍🌾",
		"descriptor": "Man Farmer: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F373": {
		"key": "1F468-1F3FC-200D-1F373",
		"value": "👨🏼‍🍳",
		"descriptor": "Man Cook: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F393": {
		"key": "1F468-1F3FC-200D-1F393",
		"value": "👨🏼‍🎓",
		"descriptor": "Man Student: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F3A4": {
		"key": "1F468-1F3FC-200D-1F3A4",
		"value": "👨🏼‍🎤",
		"descriptor": "Man Singer: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F3A8": {
		"key": "1F468-1F3FC-200D-1F3A8",
		"value": "👨🏼‍🎨",
		"descriptor": "Man Artist: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F3EB": {
		"key": "1F468-1F3FC-200D-1F3EB",
		"value": "👨🏼‍🏫",
		"descriptor": "Man Teacher: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F3ED": {
		"key": "1F468-1F3FC-200D-1F3ED",
		"value": "👨🏼‍🏭",
		"descriptor": "Man Factory Worker: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F4BB": {
		"key": "1F468-1F3FC-200D-1F4BB",
		"value": "👨🏼‍💻",
		"descriptor": "Man Technologist: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F4BC": {
		"key": "1F468-1F3FC-200D-1F4BC",
		"value": "👨🏼‍💼",
		"descriptor": "Man Office Worker: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F527": {
		"key": "1F468-1F3FC-200D-1F527",
		"value": "👨🏼‍🔧",
		"descriptor": "Man Mechanic: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F52C": {
		"key": "1F468-1F3FC-200D-1F52C",
		"value": "👨🏼‍🔬",
		"descriptor": "Man Scientist: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F680": {
		"key": "1F468-1F3FC-200D-1F680",
		"value": "👨🏼‍🚀",
		"descriptor": "Man Astronaut: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F692": {
		"key": "1F468-1F3FC-200D-1F692",
		"value": "👨🏼‍🚒",
		"descriptor": "Man Firefighter: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FB": {
		"key": "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FB",
		"value": "👨🏼‍🤝‍👨🏻",
		"descriptor": "Men Holding Hands: Medium-Light Skin Tone, Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FD": {
		"key": "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FD",
		"value": "👨🏼‍🤝‍👨🏽",
		"descriptor": "Men Holding Hands: Medium-Light Skin Tone, Medium Skin Tone"
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FE": {
		"key": "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FE",
		"value": "👨🏼‍🤝‍👨🏾",
		"descriptor": "Men Holding Hands: Medium-Light Skin Tone, Medium-Dark Skin Tone"
	},
	"1F468-1F3FC-200D-1F91D-200D-1F468-1F3FF": {
		"key": "1F468-1F3FC-200D-1F91D-200D-1F468-1F3FF",
		"value": "👨🏼‍🤝‍👨🏿",
		"descriptor": "Men Holding Hands: Medium-Light Skin Tone, Dark Skin Tone"
	},
	"1F468-1F3FC-200D-1F9AF": {
		"key": "1F468-1F3FC-200D-1F9AF",
		"value": "👨🏼‍🦯",
		"descriptor": "Man With Probing Cane: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F9B0": {
		"key": "1F468-1F3FC-200D-1F9B0",
		"value": "👨🏼‍🦰",
		"descriptor": "Man: Medium-Light Skin Tone, Red Hair"
	},
	"1F468-1F3FC-200D-1F9B1": {
		"key": "1F468-1F3FC-200D-1F9B1",
		"value": "👨🏼‍🦱",
		"descriptor": "Man: Medium-Light Skin Tone, Curly Hair"
	},
	"1F468-1F3FC-200D-1F9B2": {
		"key": "1F468-1F3FC-200D-1F9B2",
		"value": "👨🏼‍🦲",
		"descriptor": "Man: Medium-Light Skin Tone, Bald"
	},
	"1F468-1F3FC-200D-1F9B3": {
		"key": "1F468-1F3FC-200D-1F9B3",
		"value": "👨🏼‍🦳",
		"descriptor": "Man: Medium-Light Skin Tone, White Hair"
	},
	"1F468-1F3FC-200D-1F9BC": {
		"key": "1F468-1F3FC-200D-1F9BC",
		"value": "👨🏼‍🦼",
		"descriptor": "Man in Motorized Wheelchair: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-1F9BD": {
		"key": "1F468-1F3FC-200D-1F9BD",
		"value": "👨🏼‍🦽",
		"descriptor": "Man in Manual Wheelchair: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-2695-FE0F": {
		"key": "1F468-1F3FC-200D-2695-FE0F",
		"value": "👨🏼‍⚕️",
		"descriptor": "Man Health Worker: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-2696-FE0F": {
		"key": "1F468-1F3FC-200D-2696-FE0F",
		"value": "👨🏼‍⚖️",
		"descriptor": "Man Judge: Medium-Light Skin Tone"
	},
	"1F468-1F3FC-200D-2708-FE0F": {
		"key": "1F468-1F3FC-200D-2708-FE0F",
		"value": "👨🏼‍✈️",
		"descriptor": "Man Pilot: Medium-Light Skin Tone"
	},
	"1F468-1F3FD": {
		"key": "1F468-1F3FD",
		"value": "👨🏽",
		"descriptor": "Man: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F33E": {
		"key": "1F468-1F3FD-200D-1F33E",
		"value": "👨🏽‍🌾",
		"descriptor": "Man Farmer: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F373": {
		"key": "1F468-1F3FD-200D-1F373",
		"value": "👨🏽‍🍳",
		"descriptor": "Man Cook: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F393": {
		"key": "1F468-1F3FD-200D-1F393",
		"value": "👨🏽‍🎓",
		"descriptor": "Man Student: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F3A4": {
		"key": "1F468-1F3FD-200D-1F3A4",
		"value": "👨🏽‍🎤",
		"descriptor": "Man Singer: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F3A8": {
		"key": "1F468-1F3FD-200D-1F3A8",
		"value": "👨🏽‍🎨",
		"descriptor": "Man Artist: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F3EB": {
		"key": "1F468-1F3FD-200D-1F3EB",
		"value": "👨🏽‍🏫",
		"descriptor": "Man Teacher: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F3ED": {
		"key": "1F468-1F3FD-200D-1F3ED",
		"value": "👨🏽‍🏭",
		"descriptor": "Man Factory Worker: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F4BB": {
		"key": "1F468-1F3FD-200D-1F4BB",
		"value": "👨🏽‍💻",
		"descriptor": "Man Technologist: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F4BC": {
		"key": "1F468-1F3FD-200D-1F4BC",
		"value": "👨🏽‍💼",
		"descriptor": "Man Office Worker: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F527": {
		"key": "1F468-1F3FD-200D-1F527",
		"value": "👨🏽‍🔧",
		"descriptor": "Man Mechanic: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F52C": {
		"key": "1F468-1F3FD-200D-1F52C",
		"value": "👨🏽‍🔬",
		"descriptor": "Man Scientist: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F680": {
		"key": "1F468-1F3FD-200D-1F680",
		"value": "👨🏽‍🚀",
		"descriptor": "Man Astronaut: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F692": {
		"key": "1F468-1F3FD-200D-1F692",
		"value": "👨🏽‍🚒",
		"descriptor": "Man Firefighter: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FB": {
		"key": "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FB",
		"value": "👨🏽‍🤝‍👨🏻",
		"descriptor": "Men Holding Hands: Medium Skin Tone, Light Skin Tone"
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FC": {
		"key": "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FC",
		"value": "👨🏽‍🤝‍👨🏼",
		"descriptor": "Men Holding Hands: Medium Skin Tone, Medium-Light Skin Tone"
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FE": {
		"key": "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FE",
		"value": "👨🏽‍🤝‍👨🏾",
		"descriptor": "Men Holding Hands: Medium Skin Tone, Medium-Dark Skin Tone"
	},
	"1F468-1F3FD-200D-1F91D-200D-1F468-1F3FF": {
		"key": "1F468-1F3FD-200D-1F91D-200D-1F468-1F3FF",
		"value": "👨🏽‍🤝‍👨🏿",
		"descriptor": "Men Holding Hands: Medium Skin Tone, Dark Skin Tone"
	},
	"1F468-1F3FD-200D-1F9AF": {
		"key": "1F468-1F3FD-200D-1F9AF",
		"value": "👨🏽‍🦯",
		"descriptor": "Man With Probing Cane: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F9B0": {
		"key": "1F468-1F3FD-200D-1F9B0",
		"value": "👨🏽‍🦰",
		"descriptor": "Man: Medium Skin Tone, Red Hair"
	},
	"1F468-1F3FD-200D-1F9B1": {
		"key": "1F468-1F3FD-200D-1F9B1",
		"value": "👨🏽‍🦱",
		"descriptor": "Man: Medium Skin Tone, Curly Hair"
	},
	"1F468-1F3FD-200D-1F9B2": {
		"key": "1F468-1F3FD-200D-1F9B2",
		"value": "👨🏽‍🦲",
		"descriptor": "Man: Medium Skin Tone, Bald"
	},
	"1F468-1F3FD-200D-1F9B3": {
		"key": "1F468-1F3FD-200D-1F9B3",
		"value": "👨🏽‍🦳",
		"descriptor": "Man: Medium Skin Tone, White Hair"
	},
	"1F468-1F3FD-200D-1F9BC": {
		"key": "1F468-1F3FD-200D-1F9BC",
		"value": "👨🏽‍🦼",
		"descriptor": "Man in Motorized Wheelchair: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-1F9BD": {
		"key": "1F468-1F3FD-200D-1F9BD",
		"value": "👨🏽‍🦽",
		"descriptor": "Man in Manual Wheelchair: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-2695-FE0F": {
		"key": "1F468-1F3FD-200D-2695-FE0F",
		"value": "👨🏽‍⚕️",
		"descriptor": "Man Health Worker: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-2696-FE0F": {
		"key": "1F468-1F3FD-200D-2696-FE0F",
		"value": "👨🏽‍⚖️",
		"descriptor": "Man Judge: Medium Skin Tone"
	},
	"1F468-1F3FD-200D-2708-FE0F": {
		"key": "1F468-1F3FD-200D-2708-FE0F",
		"value": "👨🏽‍✈️",
		"descriptor": "Man Pilot: Medium Skin Tone"
	},
	"1F468-1F3FE": {
		"key": "1F468-1F3FE",
		"value": "👨🏾",
		"descriptor": "Man: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F33E": {
		"key": "1F468-1F3FE-200D-1F33E",
		"value": "👨🏾‍🌾",
		"descriptor": "Man Farmer: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F373": {
		"key": "1F468-1F3FE-200D-1F373",
		"value": "👨🏾‍🍳",
		"descriptor": "Man Cook: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F393": {
		"key": "1F468-1F3FE-200D-1F393",
		"value": "👨🏾‍🎓",
		"descriptor": "Man Student: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F3A4": {
		"key": "1F468-1F3FE-200D-1F3A4",
		"value": "👨🏾‍🎤",
		"descriptor": "Man Singer: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F3A8": {
		"key": "1F468-1F3FE-200D-1F3A8",
		"value": "👨🏾‍🎨",
		"descriptor": "Man Artist: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F3EB": {
		"key": "1F468-1F3FE-200D-1F3EB",
		"value": "👨🏾‍🏫",
		"descriptor": "Man Teacher: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F3ED": {
		"key": "1F468-1F3FE-200D-1F3ED",
		"value": "👨🏾‍🏭",
		"descriptor": "Man Factory Worker: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F4BB": {
		"key": "1F468-1F3FE-200D-1F4BB",
		"value": "👨🏾‍💻",
		"descriptor": "Man Technologist: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F4BC": {
		"key": "1F468-1F3FE-200D-1F4BC",
		"value": "👨🏾‍💼",
		"descriptor": "Man Office Worker: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F527": {
		"key": "1F468-1F3FE-200D-1F527",
		"value": "👨🏾‍🔧",
		"descriptor": "Man Mechanic: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F52C": {
		"key": "1F468-1F3FE-200D-1F52C",
		"value": "👨🏾‍🔬",
		"descriptor": "Man Scientist: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F680": {
		"key": "1F468-1F3FE-200D-1F680",
		"value": "👨🏾‍🚀",
		"descriptor": "Man Astronaut: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F692": {
		"key": "1F468-1F3FE-200D-1F692",
		"value": "👨🏾‍🚒",
		"descriptor": "Man Firefighter: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FB": {
		"key": "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FB",
		"value": "👨🏾‍🤝‍👨🏻",
		"descriptor": "Men Holding Hands: Medium-Dark Skin Tone, Light Skin Tone"
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FC": {
		"key": "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FC",
		"value": "👨🏾‍🤝‍👨🏼",
		"descriptor": "Men Holding Hands: Medium-Dark Skin Tone, Medium-Light Skin Tone"
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FD": {
		"key": "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FD",
		"value": "👨🏾‍🤝‍👨🏽",
		"descriptor": "Men Holding Hands: Medium-Dark Skin Tone, Medium Skin Tone"
	},
	"1F468-1F3FE-200D-1F91D-200D-1F468-1F3FF": {
		"key": "1F468-1F3FE-200D-1F91D-200D-1F468-1F3FF",
		"value": "👨🏾‍🤝‍👨🏿",
		"descriptor": "Men Holding Hands: Medium-Dark Skin Tone, Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F9AF": {
		"key": "1F468-1F3FE-200D-1F9AF",
		"value": "👨🏾‍🦯",
		"descriptor": "Man With Probing Cane: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F9B0": {
		"key": "1F468-1F3FE-200D-1F9B0",
		"value": "👨🏾‍🦰",
		"descriptor": "Man: Medium-Dark Skin Tone, Red Hair"
	},
	"1F468-1F3FE-200D-1F9B1": {
		"key": "1F468-1F3FE-200D-1F9B1",
		"value": "👨🏾‍🦱",
		"descriptor": "Man: Medium-Dark Skin Tone, Curly Hair"
	},
	"1F468-1F3FE-200D-1F9B2": {
		"key": "1F468-1F3FE-200D-1F9B2",
		"value": "👨🏾‍🦲",
		"descriptor": "Man: Medium-Dark Skin Tone, Bald"
	},
	"1F468-1F3FE-200D-1F9B3": {
		"key": "1F468-1F3FE-200D-1F9B3",
		"value": "👨🏾‍🦳",
		"descriptor": "Man: Medium-Dark Skin Tone, White Hair"
	},
	"1F468-1F3FE-200D-1F9BC": {
		"key": "1F468-1F3FE-200D-1F9BC",
		"value": "👨🏾‍🦼",
		"descriptor": "Man in Motorized Wheelchair: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-1F9BD": {
		"key": "1F468-1F3FE-200D-1F9BD",
		"value": "👨🏾‍🦽",
		"descriptor": "Man in Manual Wheelchair: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-2695-FE0F": {
		"key": "1F468-1F3FE-200D-2695-FE0F",
		"value": "👨🏾‍⚕️",
		"descriptor": "Man Health Worker: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-2696-FE0F": {
		"key": "1F468-1F3FE-200D-2696-FE0F",
		"value": "👨🏾‍⚖️",
		"descriptor": "Man Judge: Medium-Dark Skin Tone"
	},
	"1F468-1F3FE-200D-2708-FE0F": {
		"key": "1F468-1F3FE-200D-2708-FE0F",
		"value": "👨🏾‍✈️",
		"descriptor": "Man Pilot: Medium-Dark Skin Tone"
	},
	"1F468-1F3FF": {
		"key": "1F468-1F3FF",
		"value": "👨🏿",
		"descriptor": "Man: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F33E": {
		"key": "1F468-1F3FF-200D-1F33E",
		"value": "👨🏿‍🌾",
		"descriptor": "Man Farmer: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F373": {
		"key": "1F468-1F3FF-200D-1F373",
		"value": "👨🏿‍🍳",
		"descriptor": "Man Cook: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F393": {
		"key": "1F468-1F3FF-200D-1F393",
		"value": "👨🏿‍🎓",
		"descriptor": "Man Student: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F3A4": {
		"key": "1F468-1F3FF-200D-1F3A4",
		"value": "👨🏿‍🎤",
		"descriptor": "Man Singer: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F3A8": {
		"key": "1F468-1F3FF-200D-1F3A8",
		"value": "👨🏿‍🎨",
		"descriptor": "Man Artist: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F3EB": {
		"key": "1F468-1F3FF-200D-1F3EB",
		"value": "👨🏿‍🏫",
		"descriptor": "Man Teacher: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F3ED": {
		"key": "1F468-1F3FF-200D-1F3ED",
		"value": "👨🏿‍🏭",
		"descriptor": "Man Factory Worker: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F4BB": {
		"key": "1F468-1F3FF-200D-1F4BB",
		"value": "👨🏿‍💻",
		"descriptor": "Man Technologist: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F4BC": {
		"key": "1F468-1F3FF-200D-1F4BC",
		"value": "👨🏿‍💼",
		"descriptor": "Man Office Worker: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F527": {
		"key": "1F468-1F3FF-200D-1F527",
		"value": "👨🏿‍🔧",
		"descriptor": "Man Mechanic: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F52C": {
		"key": "1F468-1F3FF-200D-1F52C",
		"value": "👨🏿‍🔬",
		"descriptor": "Man Scientist: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F680": {
		"key": "1F468-1F3FF-200D-1F680",
		"value": "👨🏿‍🚀",
		"descriptor": "Man Astronaut: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F692": {
		"key": "1F468-1F3FF-200D-1F692",
		"value": "👨🏿‍🚒",
		"descriptor": "Man Firefighter: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FB": {
		"key": "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FB",
		"value": "👨🏿‍🤝‍👨🏻",
		"descriptor": "Men Holding Hands: Dark Skin Tone, Light Skin Tone"
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FC": {
		"key": "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FC",
		"value": "👨🏿‍🤝‍👨🏼",
		"descriptor": "Men Holding Hands: Dark Skin Tone, Medium-Light Skin Tone"
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FD": {
		"key": "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FD",
		"value": "👨🏿‍🤝‍👨🏽",
		"descriptor": "Men Holding Hands: Dark Skin Tone, Medium Skin Tone"
	},
	"1F468-1F3FF-200D-1F91D-200D-1F468-1F3FE": {
		"key": "1F468-1F3FF-200D-1F91D-200D-1F468-1F3FE",
		"value": "👨🏿‍🤝‍👨🏾",
		"descriptor": "Men Holding Hands: Dark Skin Tone, Medium-Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F9AF": {
		"key": "1F468-1F3FF-200D-1F9AF",
		"value": "👨🏿‍🦯",
		"descriptor": "Man With Probing Cane: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F9B0": {
		"key": "1F468-1F3FF-200D-1F9B0",
		"value": "👨🏿‍🦰",
		"descriptor": "Man: Dark Skin Tone, Red Hair"
	},
	"1F468-1F3FF-200D-1F9B1": {
		"key": "1F468-1F3FF-200D-1F9B1",
		"value": "👨🏿‍🦱",
		"descriptor": "Man: Dark Skin Tone, Curly Hair"
	},
	"1F468-1F3FF-200D-1F9B2": {
		"key": "1F468-1F3FF-200D-1F9B2",
		"value": "👨🏿‍🦲",
		"descriptor": "Man: Dark Skin Tone, Bald"
	},
	"1F468-1F3FF-200D-1F9B3": {
		"key": "1F468-1F3FF-200D-1F9B3",
		"value": "👨🏿‍🦳",
		"descriptor": "Man: Dark Skin Tone, White Hair"
	},
	"1F468-1F3FF-200D-1F9BC": {
		"key": "1F468-1F3FF-200D-1F9BC",
		"value": "👨🏿‍🦼",
		"descriptor": "Man in Motorized Wheelchair: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-1F9BD": {
		"key": "1F468-1F3FF-200D-1F9BD",
		"value": "👨🏿‍🦽",
		"descriptor": "Man in Manual Wheelchair: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-2695-FE0F": {
		"key": "1F468-1F3FF-200D-2695-FE0F",
		"value": "👨🏿‍⚕️",
		"descriptor": "Man Health Worker: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-2696-FE0F": {
		"key": "1F468-1F3FF-200D-2696-FE0F",
		"value": "👨🏿‍⚖️",
		"descriptor": "Man Judge: Dark Skin Tone"
	},
	"1F468-1F3FF-200D-2708-FE0F": {
		"key": "1F468-1F3FF-200D-2708-FE0F",
		"value": "👨🏿‍✈️",
		"descriptor": "Man Pilot: Dark Skin Tone"
	},
	"1F468-200D-1F33E": {
		"key": "1F468-200D-1F33E",
		"value": "👨‍🌾",
		"descriptor": "Man Farmer"
	},
	"1F468-200D-1F373": {
		"key": "1F468-200D-1F373",
		"value": "👨‍🍳",
		"descriptor": "Man Cook"
	},
	"1F468-200D-1F393": {
		"key": "1F468-200D-1F393",
		"value": "👨‍🎓",
		"descriptor": "Man Student"
	},
	"1F468-200D-1F3A4": {
		"key": "1F468-200D-1F3A4",
		"value": "👨‍🎤",
		"descriptor": "Man Singer"
	},
	"1F468-200D-1F3A8": {
		"key": "1F468-200D-1F3A8",
		"value": "👨‍🎨",
		"descriptor": "Man Artist"
	},
	"1F468-200D-1F3EB": {
		"key": "1F468-200D-1F3EB",
		"value": "👨‍🏫",
		"descriptor": "Man Teacher"
	},
	"1F468-200D-1F3ED": {
		"key": "1F468-200D-1F3ED",
		"value": "👨‍🏭",
		"descriptor": "Man Factory Worker"
	},
	"1F468-200D-1F466": {
		"key": "1F468-200D-1F466",
		"value": "👨‍👦",
		"descriptor": "Family: Man, Boy"
	},
	"1F468-200D-1F466-200D-1F466": {
		"key": "1F468-200D-1F466-200D-1F466",
		"value": "👨‍👦‍👦",
		"descriptor": "Family: Man, Boy, Boy"
	},
	"1F468-200D-1F467": {
		"key": "1F468-200D-1F467",
		"value": "👨‍👧",
		"descriptor": "Family: Man, Girl"
	},
	"1F468-200D-1F467-200D-1F466": {
		"key": "1F468-200D-1F467-200D-1F466",
		"value": "👨‍👧‍👦",
		"descriptor": "Family: Man, Girl, Boy"
	},
	"1F468-200D-1F467-200D-1F467": {
		"key": "1F468-200D-1F467-200D-1F467",
		"value": "👨‍👧‍👧",
		"descriptor": "Family: Man, Girl, Girl"
	},
	"1F468-200D-1F468-200D-1F466": {
		"key": "1F468-200D-1F468-200D-1F466",
		"value": "👨‍👨‍👦",
		"descriptor": "Family: Man, Man, Boy"
	},
	"1F468-200D-1F468-200D-1F466-200D-1F466": {
		"key": "1F468-200D-1F468-200D-1F466-200D-1F466",
		"value": "👨‍👨‍👦‍👦",
		"descriptor": "Family: Man, Man, Boy, Boy"
	},
	"1F468-200D-1F468-200D-1F467": {
		"key": "1F468-200D-1F468-200D-1F467",
		"value": "👨‍👨‍👧",
		"descriptor": "Family: Man, Man, Girl"
	},
	"1F468-200D-1F468-200D-1F467-200D-1F466": {
		"key": "1F468-200D-1F468-200D-1F467-200D-1F466",
		"value": "👨‍👨‍👧‍👦",
		"descriptor": "Family: Man, Man, Girl, Boy"
	},
	"1F468-200D-1F468-200D-1F467-200D-1F467": {
		"key": "1F468-200D-1F468-200D-1F467-200D-1F467",
		"value": "👨‍👨‍👧‍👧",
		"descriptor": "Family: Man, Man, Girl, Girl"
	},
	"1F468-200D-1F469-200D-1F466": {
		"key": "1F468-200D-1F469-200D-1F466",
		"value": "👨‍👩‍👦",
		"descriptor": "Family: Man, Woman, Boy"
	},
	"1F468-200D-1F469-200D-1F466-200D-1F466": {
		"key": "1F468-200D-1F469-200D-1F466-200D-1F466",
		"value": "👨‍👩‍👦‍👦",
		"descriptor": "Family: Man, Woman, Boy, Boy"
	},
	"1F468-200D-1F469-200D-1F467": {
		"key": "1F468-200D-1F469-200D-1F467",
		"value": "👨‍👩‍👧",
		"descriptor": "Family: Man, Woman, Girl"
	},
	"1F468-200D-1F469-200D-1F467-200D-1F466": {
		"key": "1F468-200D-1F469-200D-1F467-200D-1F466",
		"value": "👨‍👩‍👧‍👦",
		"descriptor": "Family: Man, Woman, Girl, Boy"
	},
	"1F468-200D-1F469-200D-1F467-200D-1F467": {
		"key": "1F468-200D-1F469-200D-1F467-200D-1F467",
		"value": "👨‍👩‍👧‍👧",
		"descriptor": "Family: Man, Woman, Girl, Girl"
	},
	"1F468-200D-1F4BB": {
		"key": "1F468-200D-1F4BB",
		"value": "👨‍💻",
		"descriptor": "Man Technologist"
	},
	"1F468-200D-1F4BC": {
		"key": "1F468-200D-1F4BC",
		"value": "👨‍💼",
		"descriptor": "Man Office Worker"
	},
	"1F468-200D-1F527": {
		"key": "1F468-200D-1F527",
		"value": "👨‍🔧",
		"descriptor": "Man Mechanic"
	},
	"1F468-200D-1F52C": {
		"key": "1F468-200D-1F52C",
		"value": "👨‍🔬",
		"descriptor": "Man Scientist"
	},
	"1F468-200D-1F680": {
		"key": "1F468-200D-1F680",
		"value": "👨‍🚀",
		"descriptor": "Man Astronaut"
	},
	"1F468-200D-1F692": {
		"key": "1F468-200D-1F692",
		"value": "👨‍🚒",
		"descriptor": "Man Firefighter"
	},
	"1F468-200D-1F9AF": {
		"key": "1F468-200D-1F9AF",
		"value": "👨‍🦯",
		"descriptor": "Man With Probing Cane"
	},
	"1F468-200D-1F9B0": {
		"key": "1F468-200D-1F9B0",
		"value": "👨‍🦰",
		"descriptor": "Man: Red Hair"
	},
	"1F468-200D-1F9B1": {
		"key": "1F468-200D-1F9B1",
		"value": "👨‍🦱",
		"descriptor": "Man: Curly Hair"
	},
	"1F468-200D-1F9B2": {
		"key": "1F468-200D-1F9B2",
		"value": "👨‍🦲",
		"descriptor": "Man: Bald"
	},
	"1F468-200D-1F9B3": {
		"key": "1F468-200D-1F9B3",
		"value": "👨‍🦳",
		"descriptor": "Man: White Hair"
	},
	"1F468-200D-1F9BC": {
		"key": "1F468-200D-1F9BC",
		"value": "👨‍🦼",
		"descriptor": "Man in Motorized Wheelchair"
	},
	"1F468-200D-1F9BD": {
		"key": "1F468-200D-1F9BD",
		"value": "👨‍🦽",
		"descriptor": "Man in Manual Wheelchair"
	},
	"1F468-200D-2695-FE0F": {
		"key": "1F468-200D-2695-FE0F",
		"value": "👨‍⚕️",
		"descriptor": "Man Health Worker"
	},
	"1F468-200D-2696-FE0F": {
		"key": "1F468-200D-2696-FE0F",
		"value": "👨‍⚖️",
		"descriptor": "Man Judge"
	},
	"1F468-200D-2708-FE0F": {
		"key": "1F468-200D-2708-FE0F",
		"value": "👨‍✈️",
		"descriptor": "Man Pilot"
	},
	"1F468-200D-2764-FE0F-200D-1F468": {
		"key": "1F468-200D-2764-FE0F-200D-1F468",
		"value": "👨‍❤️‍👨",
		"descriptor": "Couple With Heart: Man, Man"
	},
	"1F468-200D-2764-FE0F-200D-1F48B-200D-1F468": {
		"key": "1F468-200D-2764-FE0F-200D-1F48B-200D-1F468",
		"value": "👨‍❤️‍💋‍👨",
		"descriptor": "Kiss: Man, Man"
	},
	"1F469": {
		"key": "1F469",
		"value": "👩",
		"descriptor": "Woman"
	},
	"1F469-1F3FB": {
		"key": "1F469-1F3FB",
		"value": "👩🏻",
		"descriptor": "Woman: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F33E": {
		"key": "1F469-1F3FB-200D-1F33E",
		"value": "👩🏻‍🌾",
		"descriptor": "Woman Farmer: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F373": {
		"key": "1F469-1F3FB-200D-1F373",
		"value": "👩🏻‍🍳",
		"descriptor": "Woman Cook: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F393": {
		"key": "1F469-1F3FB-200D-1F393",
		"value": "👩🏻‍🎓",
		"descriptor": "Woman Student: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F3A4": {
		"key": "1F469-1F3FB-200D-1F3A4",
		"value": "👩🏻‍🎤",
		"descriptor": "Woman Singer: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F3A8": {
		"key": "1F469-1F3FB-200D-1F3A8",
		"value": "👩🏻‍🎨",
		"descriptor": "Woman Artist: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F3EB": {
		"key": "1F469-1F3FB-200D-1F3EB",
		"value": "👩🏻‍🏫",
		"descriptor": "Woman Teacher: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F3ED": {
		"key": "1F469-1F3FB-200D-1F3ED",
		"value": "👩🏻‍🏭",
		"descriptor": "Woman Factory Worker: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F4BB": {
		"key": "1F469-1F3FB-200D-1F4BB",
		"value": "👩🏻‍💻",
		"descriptor": "Woman Technologist: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F4BC": {
		"key": "1F469-1F3FB-200D-1F4BC",
		"value": "👩🏻‍💼",
		"descriptor": "Woman Office Worker: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F527": {
		"key": "1F469-1F3FB-200D-1F527",
		"value": "👩🏻‍🔧",
		"descriptor": "Woman Mechanic: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F52C": {
		"key": "1F469-1F3FB-200D-1F52C",
		"value": "👩🏻‍🔬",
		"descriptor": "Woman Scientist: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F680": {
		"key": "1F469-1F3FB-200D-1F680",
		"value": "👩🏻‍🚀",
		"descriptor": "Woman Astronaut: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F692": {
		"key": "1F469-1F3FB-200D-1F692",
		"value": "👩🏻‍🚒",
		"descriptor": "Woman Firefighter: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FC": {
		"key": "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FC",
		"value": "👩🏻‍🤝‍👨🏼",
		"descriptor": "Woman and Man Holding Hands: Light Skin Tone, Medium-Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FD": {
		"key": "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FD",
		"value": "👩🏻‍🤝‍👨🏽",
		"descriptor": "Woman and Man Holding Hands: Light Skin Tone, Medium Skin Tone"
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FE": {
		"key": "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FE",
		"value": "👩🏻‍🤝‍👨🏾",
		"descriptor": "Woman and Man Holding Hands: Light Skin Tone, Medium-Dark Skin Tone"
	},
	"1F469-1F3FB-200D-1F91D-200D-1F468-1F3FF": {
		"key": "1F469-1F3FB-200D-1F91D-200D-1F468-1F3FF",
		"value": "👩🏻‍🤝‍👨🏿",
		"descriptor": "Woman and Man Holding Hands: Light Skin Tone, Dark Skin Tone"
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FC": {
		"key": "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FC",
		"value": "👩🏻‍🤝‍👩🏼",
		"descriptor": "Women Holding Hands: Light Skin Tone, Medium-Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FD": {
		"key": "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FD",
		"value": "👩🏻‍🤝‍👩🏽",
		"descriptor": "Women Holding Hands: Light Skin Tone, Medium Skin Tone"
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FE": {
		"key": "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FE",
		"value": "👩🏻‍🤝‍👩🏾",
		"descriptor": "Women Holding Hands: Light Skin Tone, Medium-Dark Skin Tone"
	},
	"1F469-1F3FB-200D-1F91D-200D-1F469-1F3FF": {
		"key": "1F469-1F3FB-200D-1F91D-200D-1F469-1F3FF",
		"value": "👩🏻‍🤝‍👩🏿",
		"descriptor": "Women Holding Hands: Light Skin Tone, Dark Skin Tone"
	},
	"1F469-1F3FB-200D-1F9AF": {
		"key": "1F469-1F3FB-200D-1F9AF",
		"value": "👩🏻‍🦯",
		"descriptor": "Woman With Probing Cane: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F9B0": {
		"key": "1F469-1F3FB-200D-1F9B0",
		"value": "👩🏻‍🦰",
		"descriptor": "Woman: Light Skin Tone, Red Hair"
	},
	"1F469-1F3FB-200D-1F9B1": {
		"key": "1F469-1F3FB-200D-1F9B1",
		"value": "👩🏻‍🦱",
		"descriptor": "Woman: Light Skin Tone, Curly Hair"
	},
	"1F469-1F3FB-200D-1F9B2": {
		"key": "1F469-1F3FB-200D-1F9B2",
		"value": "👩🏻‍🦲",
		"descriptor": "Woman: Light Skin Tone, Bald"
	},
	"1F469-1F3FB-200D-1F9B3": {
		"key": "1F469-1F3FB-200D-1F9B3",
		"value": "👩🏻‍🦳",
		"descriptor": "Woman: Light Skin Tone, White Hair"
	},
	"1F469-1F3FB-200D-1F9BC": {
		"key": "1F469-1F3FB-200D-1F9BC",
		"value": "👩🏻‍🦼",
		"descriptor": "Woman in Motorized Wheelchair: Light Skin Tone"
	},
	"1F469-1F3FB-200D-1F9BD": {
		"key": "1F469-1F3FB-200D-1F9BD",
		"value": "👩🏻‍🦽",
		"descriptor": "Woman in Manual Wheelchair: Light Skin Tone"
	},
	"1F469-1F3FB-200D-2695-FE0F": {
		"key": "1F469-1F3FB-200D-2695-FE0F",
		"value": "👩🏻‍⚕️",
		"descriptor": "Woman Health Worker: Light Skin Tone"
	},
	"1F469-1F3FB-200D-2696-FE0F": {
		"key": "1F469-1F3FB-200D-2696-FE0F",
		"value": "👩🏻‍⚖️",
		"descriptor": "Woman Judge: Light Skin Tone"
	},
	"1F469-1F3FB-200D-2708-FE0F": {
		"key": "1F469-1F3FB-200D-2708-FE0F",
		"value": "👩🏻‍✈️",
		"descriptor": "Woman Pilot: Light Skin Tone"
	},
	"1F469-1F3FC": {
		"key": "1F469-1F3FC",
		"value": "👩🏼",
		"descriptor": "Woman: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F33E": {
		"key": "1F469-1F3FC-200D-1F33E",
		"value": "👩🏼‍🌾",
		"descriptor": "Woman Farmer: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F373": {
		"key": "1F469-1F3FC-200D-1F373",
		"value": "👩🏼‍🍳",
		"descriptor": "Woman Cook: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F393": {
		"key": "1F469-1F3FC-200D-1F393",
		"value": "👩🏼‍🎓",
		"descriptor": "Woman Student: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F3A4": {
		"key": "1F469-1F3FC-200D-1F3A4",
		"value": "👩🏼‍🎤",
		"descriptor": "Woman Singer: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F3A8": {
		"key": "1F469-1F3FC-200D-1F3A8",
		"value": "👩🏼‍🎨",
		"descriptor": "Woman Artist: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F3EB": {
		"key": "1F469-1F3FC-200D-1F3EB",
		"value": "👩🏼‍🏫",
		"descriptor": "Woman Teacher: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F3ED": {
		"key": "1F469-1F3FC-200D-1F3ED",
		"value": "👩🏼‍🏭",
		"descriptor": "Woman Factory Worker: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F4BB": {
		"key": "1F469-1F3FC-200D-1F4BB",
		"value": "👩🏼‍💻",
		"descriptor": "Woman Technologist: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F4BC": {
		"key": "1F469-1F3FC-200D-1F4BC",
		"value": "👩🏼‍💼",
		"descriptor": "Woman Office Worker: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F527": {
		"key": "1F469-1F3FC-200D-1F527",
		"value": "👩🏼‍🔧",
		"descriptor": "Woman Mechanic: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F52C": {
		"key": "1F469-1F3FC-200D-1F52C",
		"value": "👩🏼‍🔬",
		"descriptor": "Woman Scientist: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F680": {
		"key": "1F469-1F3FC-200D-1F680",
		"value": "👩🏼‍🚀",
		"descriptor": "Woman Astronaut: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F692": {
		"key": "1F469-1F3FC-200D-1F692",
		"value": "👩🏼‍🚒",
		"descriptor": "Woman Firefighter: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FB": {
		"key": "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FB",
		"value": "👩🏼‍🤝‍👨🏻",
		"descriptor": "Woman and Man Holding Hands: Medium-Light Skin Tone, Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FD": {
		"key": "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FD",
		"value": "👩🏼‍🤝‍👨🏽",
		"descriptor": "Woman and Man Holding Hands: Medium-Light Skin Tone, Medium Skin Tone"
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FE": {
		"key": "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FE",
		"value": "👩🏼‍🤝‍👨🏾",
		"descriptor": "Woman and Man Holding Hands: Medium-Light Skin Tone, Medium-Dark Skin Tone"
	},
	"1F469-1F3FC-200D-1F91D-200D-1F468-1F3FF": {
		"key": "1F469-1F3FC-200D-1F91D-200D-1F468-1F3FF",
		"value": "👩🏼‍🤝‍👨🏿",
		"descriptor": "Woman and Man Holding Hands: Medium-Light Skin Tone, Dark Skin Tone"
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FB": {
		"key": "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FB",
		"value": "👩🏼‍🤝‍👩🏻",
		"descriptor": "Women Holding Hands: Medium-Light Skin Tone, Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FD": {
		"key": "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FD",
		"value": "👩🏼‍🤝‍👩🏽",
		"descriptor": "Women Holding Hands: Medium-Light Skin Tone, Medium Skin Tone"
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FE": {
		"key": "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FE",
		"value": "👩🏼‍🤝‍👩🏾",
		"descriptor": "Women Holding Hands: Medium-Light Skin Tone, Medium-Dark Skin Tone"
	},
	"1F469-1F3FC-200D-1F91D-200D-1F469-1F3FF": {
		"key": "1F469-1F3FC-200D-1F91D-200D-1F469-1F3FF",
		"value": "👩🏼‍🤝‍👩🏿",
		"descriptor": "Women Holding Hands: Medium-Light Skin Tone, Dark Skin Tone"
	},
	"1F469-1F3FC-200D-1F9AF": {
		"key": "1F469-1F3FC-200D-1F9AF",
		"value": "👩🏼‍🦯",
		"descriptor": "Woman With Probing Cane: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F9B0": {
		"key": "1F469-1F3FC-200D-1F9B0",
		"value": "👩🏼‍🦰",
		"descriptor": "Woman: Medium-Light Skin Tone, Red Hair"
	},
	"1F469-1F3FC-200D-1F9B1": {
		"key": "1F469-1F3FC-200D-1F9B1",
		"value": "👩🏼‍🦱",
		"descriptor": "Woman: Medium-Light Skin Tone, Curly Hair"
	},
	"1F469-1F3FC-200D-1F9B2": {
		"key": "1F469-1F3FC-200D-1F9B2",
		"value": "👩🏼‍🦲",
		"descriptor": "Woman: Medium-Light Skin Tone, Bald"
	},
	"1F469-1F3FC-200D-1F9B3": {
		"key": "1F469-1F3FC-200D-1F9B3",
		"value": "👩🏼‍🦳",
		"descriptor": "Woman: Medium-Light Skin Tone, White Hair"
	},
	"1F469-1F3FC-200D-1F9BC": {
		"key": "1F469-1F3FC-200D-1F9BC",
		"value": "👩🏼‍🦼",
		"descriptor": "Woman in Motorized Wheelchair: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-1F9BD": {
		"key": "1F469-1F3FC-200D-1F9BD",
		"value": "👩🏼‍🦽",
		"descriptor": "Woman in Manual Wheelchair: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-2695-FE0F": {
		"key": "1F469-1F3FC-200D-2695-FE0F",
		"value": "👩🏼‍⚕️",
		"descriptor": "Woman Health Worker: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-2696-FE0F": {
		"key": "1F469-1F3FC-200D-2696-FE0F",
		"value": "👩🏼‍⚖️",
		"descriptor": "Woman Judge: Medium-Light Skin Tone"
	},
	"1F469-1F3FC-200D-2708-FE0F": {
		"key": "1F469-1F3FC-200D-2708-FE0F",
		"value": "👩🏼‍✈️",
		"descriptor": "Woman Pilot: Medium-Light Skin Tone"
	},
	"1F469-1F3FD": {
		"key": "1F469-1F3FD",
		"value": "👩🏽",
		"descriptor": "Woman: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F33E": {
		"key": "1F469-1F3FD-200D-1F33E",
		"value": "👩🏽‍🌾",
		"descriptor": "Woman Farmer: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F373": {
		"key": "1F469-1F3FD-200D-1F373",
		"value": "👩🏽‍🍳",
		"descriptor": "Woman Cook: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F393": {
		"key": "1F469-1F3FD-200D-1F393",
		"value": "👩🏽‍🎓",
		"descriptor": "Woman Student: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F3A4": {
		"key": "1F469-1F3FD-200D-1F3A4",
		"value": "👩🏽‍🎤",
		"descriptor": "Woman Singer: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F3A8": {
		"key": "1F469-1F3FD-200D-1F3A8",
		"value": "👩🏽‍🎨",
		"descriptor": "Woman Artist: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F3EB": {
		"key": "1F469-1F3FD-200D-1F3EB",
		"value": "👩🏽‍🏫",
		"descriptor": "Woman Teacher: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F3ED": {
		"key": "1F469-1F3FD-200D-1F3ED",
		"value": "👩🏽‍🏭",
		"descriptor": "Woman Factory Worker: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F4BB": {
		"key": "1F469-1F3FD-200D-1F4BB",
		"value": "👩🏽‍💻",
		"descriptor": "Woman Technologist: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F4BC": {
		"key": "1F469-1F3FD-200D-1F4BC",
		"value": "👩🏽‍💼",
		"descriptor": "Woman Office Worker: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F527": {
		"key": "1F469-1F3FD-200D-1F527",
		"value": "👩🏽‍🔧",
		"descriptor": "Woman Mechanic: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F52C": {
		"key": "1F469-1F3FD-200D-1F52C",
		"value": "👩🏽‍🔬",
		"descriptor": "Woman Scientist: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F680": {
		"key": "1F469-1F3FD-200D-1F680",
		"value": "👩🏽‍🚀",
		"descriptor": "Woman Astronaut: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F692": {
		"key": "1F469-1F3FD-200D-1F692",
		"value": "👩🏽‍🚒",
		"descriptor": "Woman Firefighter: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FB": {
		"key": "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FB",
		"value": "👩🏽‍🤝‍👨🏻",
		"descriptor": "Woman and Man Holding Hands: Medium Skin Tone, Light Skin Tone"
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FC": {
		"key": "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FC",
		"value": "👩🏽‍🤝‍👨🏼",
		"descriptor": "Woman and Man Holding Hands: Medium Skin Tone, Medium-Light Skin Tone"
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FE": {
		"key": "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FE",
		"value": "👩🏽‍🤝‍👨🏾",
		"descriptor": "Woman and Man Holding Hands: Medium Skin Tone, Medium-Dark Skin Tone"
	},
	"1F469-1F3FD-200D-1F91D-200D-1F468-1F3FF": {
		"key": "1F469-1F3FD-200D-1F91D-200D-1F468-1F3FF",
		"value": "👩🏽‍🤝‍👨🏿",
		"descriptor": "Woman and Man Holding Hands: Medium Skin Tone, Dark Skin Tone"
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FB": {
		"key": "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FB",
		"value": "👩🏽‍🤝‍👩🏻",
		"descriptor": "Women Holding Hands: Medium Skin Tone, Light Skin Tone"
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FC": {
		"key": "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FC",
		"value": "👩🏽‍🤝‍👩🏼",
		"descriptor": "Women Holding Hands: Medium Skin Tone, Medium-Light Skin Tone"
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FE": {
		"key": "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FE",
		"value": "👩🏽‍🤝‍👩🏾",
		"descriptor": "Women Holding Hands: Medium Skin Tone, Medium-Dark Skin Tone"
	},
	"1F469-1F3FD-200D-1F91D-200D-1F469-1F3FF": {
		"key": "1F469-1F3FD-200D-1F91D-200D-1F469-1F3FF",
		"value": "👩🏽‍🤝‍👩🏿",
		"descriptor": "Women Holding Hands: Medium Skin Tone, Dark Skin Tone"
	},
	"1F469-1F3FD-200D-1F9AF": {
		"key": "1F469-1F3FD-200D-1F9AF",
		"value": "👩🏽‍🦯",
		"descriptor": "Woman With Probing Cane: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F9B0": {
		"key": "1F469-1F3FD-200D-1F9B0",
		"value": "👩🏽‍🦰",
		"descriptor": "Woman: Medium Skin Tone, Red Hair"
	},
	"1F469-1F3FD-200D-1F9B1": {
		"key": "1F469-1F3FD-200D-1F9B1",
		"value": "👩🏽‍🦱",
		"descriptor": "Woman: Medium Skin Tone, Curly Hair"
	},
	"1F469-1F3FD-200D-1F9B2": {
		"key": "1F469-1F3FD-200D-1F9B2",
		"value": "👩🏽‍🦲",
		"descriptor": "Woman: Medium Skin Tone, Bald"
	},
	"1F469-1F3FD-200D-1F9B3": {
		"key": "1F469-1F3FD-200D-1F9B3",
		"value": "👩🏽‍🦳",
		"descriptor": "Woman: Medium Skin Tone, White Hair"
	},
	"1F469-1F3FD-200D-1F9BC": {
		"key": "1F469-1F3FD-200D-1F9BC",
		"value": "👩🏽‍🦼",
		"descriptor": "Woman in Motorized Wheelchair: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-1F9BD": {
		"key": "1F469-1F3FD-200D-1F9BD",
		"value": "👩🏽‍🦽",
		"descriptor": "Woman in Manual Wheelchair: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-2695-FE0F": {
		"key": "1F469-1F3FD-200D-2695-FE0F",
		"value": "👩🏽‍⚕️",
		"descriptor": "Woman Health Worker: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-2696-FE0F": {
		"key": "1F469-1F3FD-200D-2696-FE0F",
		"value": "👩🏽‍⚖️",
		"descriptor": "Woman Judge: Medium Skin Tone"
	},
	"1F469-1F3FD-200D-2708-FE0F": {
		"key": "1F469-1F3FD-200D-2708-FE0F",
		"value": "👩🏽‍✈️",
		"descriptor": "Woman Pilot: Medium Skin Tone"
	},
	"1F469-1F3FE": {
		"key": "1F469-1F3FE",
		"value": "👩🏾",
		"descriptor": "Woman: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F33E": {
		"key": "1F469-1F3FE-200D-1F33E",
		"value": "👩🏾‍🌾",
		"descriptor": "Woman Farmer: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F373": {
		"key": "1F469-1F3FE-200D-1F373",
		"value": "👩🏾‍🍳",
		"descriptor": "Woman Cook: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F393": {
		"key": "1F469-1F3FE-200D-1F393",
		"value": "👩🏾‍🎓",
		"descriptor": "Woman Student: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F3A4": {
		"key": "1F469-1F3FE-200D-1F3A4",
		"value": "👩🏾‍🎤",
		"descriptor": "Woman Singer: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F3A8": {
		"key": "1F469-1F3FE-200D-1F3A8",
		"value": "👩🏾‍🎨",
		"descriptor": "Woman Artist: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F3EB": {
		"key": "1F469-1F3FE-200D-1F3EB",
		"value": "👩🏾‍🏫",
		"descriptor": "Woman Teacher: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F3ED": {
		"key": "1F469-1F3FE-200D-1F3ED",
		"value": "👩🏾‍🏭",
		"descriptor": "Woman Factory Worker: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F4BB": {
		"key": "1F469-1F3FE-200D-1F4BB",
		"value": "👩🏾‍💻",
		"descriptor": "Woman Technologist: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F4BC": {
		"key": "1F469-1F3FE-200D-1F4BC",
		"value": "👩🏾‍💼",
		"descriptor": "Woman Office Worker: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F527": {
		"key": "1F469-1F3FE-200D-1F527",
		"value": "👩🏾‍🔧",
		"descriptor": "Woman Mechanic: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F52C": {
		"key": "1F469-1F3FE-200D-1F52C",
		"value": "👩🏾‍🔬",
		"descriptor": "Woman Scientist: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F680": {
		"key": "1F469-1F3FE-200D-1F680",
		"value": "👩🏾‍🚀",
		"descriptor": "Woman Astronaut: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F692": {
		"key": "1F469-1F3FE-200D-1F692",
		"value": "👩🏾‍🚒",
		"descriptor": "Woman Firefighter: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FB": {
		"key": "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FB",
		"value": "👩🏾‍🤝‍👨🏻",
		"descriptor": "Woman and Man Holding Hands: Medium-Dark Skin Tone, Light Skin Tone"
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FC": {
		"key": "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FC",
		"value": "👩🏾‍🤝‍👨🏼",
		"descriptor": "Woman and Man Holding Hands: Medium-Dark Skin Tone, Medium-Light Skin Tone"
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FD": {
		"key": "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FD",
		"value": "👩🏾‍🤝‍👨🏽",
		"descriptor": "Woman and Man Holding Hands: Medium-Dark Skin Tone, Medium Skin Tone"
	},
	"1F469-1F3FE-200D-1F91D-200D-1F468-1F3FF": {
		"key": "1F469-1F3FE-200D-1F91D-200D-1F468-1F3FF",
		"value": "👩🏾‍🤝‍👨🏿",
		"descriptor": "Woman and Man Holding Hands: Medium-Dark Skin Tone, Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FB": {
		"key": "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FB",
		"value": "👩🏾‍🤝‍👩🏻",
		"descriptor": "Women Holding Hands: Medium-Dark Skin Tone, Light Skin Tone"
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FC": {
		"key": "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FC",
		"value": "👩🏾‍🤝‍👩🏼",
		"descriptor": "Women Holding Hands: Medium-Dark Skin Tone, Medium-Light Skin Tone"
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FD": {
		"key": "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FD",
		"value": "👩🏾‍🤝‍👩🏽",
		"descriptor": "Women Holding Hands: Medium-Dark Skin Tone, Medium Skin Tone"
	},
	"1F469-1F3FE-200D-1F91D-200D-1F469-1F3FF": {
		"key": "1F469-1F3FE-200D-1F91D-200D-1F469-1F3FF",
		"value": "👩🏾‍🤝‍👩🏿",
		"descriptor": "Women Holding Hands: Medium-Dark Skin Tone, Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F9AF": {
		"key": "1F469-1F3FE-200D-1F9AF",
		"value": "👩🏾‍🦯",
		"descriptor": "Woman With Probing Cane: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F9B0": {
		"key": "1F469-1F3FE-200D-1F9B0",
		"value": "👩🏾‍🦰",
		"descriptor": "Woman: Medium-Dark Skin Tone, Red Hair"
	},
	"1F469-1F3FE-200D-1F9B1": {
		"key": "1F469-1F3FE-200D-1F9B1",
		"value": "👩🏾‍🦱",
		"descriptor": "Woman: Medium-Dark Skin Tone, Curly Hair"
	},
	"1F469-1F3FE-200D-1F9B2": {
		"key": "1F469-1F3FE-200D-1F9B2",
		"value": "👩🏾‍🦲",
		"descriptor": "Woman: Medium-Dark Skin Tone, Bald"
	},
	"1F469-1F3FE-200D-1F9B3": {
		"key": "1F469-1F3FE-200D-1F9B3",
		"value": "👩🏾‍🦳",
		"descriptor": "Woman: Medium-Dark Skin Tone, White Hair"
	},
	"1F469-1F3FE-200D-1F9BC": {
		"key": "1F469-1F3FE-200D-1F9BC",
		"value": "👩🏾‍🦼",
		"descriptor": "Woman in Motorized Wheelchair: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-1F9BD": {
		"key": "1F469-1F3FE-200D-1F9BD",
		"value": "👩🏾‍🦽",
		"descriptor": "Woman in Manual Wheelchair: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-2695-FE0F": {
		"key": "1F469-1F3FE-200D-2695-FE0F",
		"value": "👩🏾‍⚕️",
		"descriptor": "Woman Health Worker: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-2696-FE0F": {
		"key": "1F469-1F3FE-200D-2696-FE0F",
		"value": "👩🏾‍⚖️",
		"descriptor": "Woman Judge: Medium-Dark Skin Tone"
	},
	"1F469-1F3FE-200D-2708-FE0F": {
		"key": "1F469-1F3FE-200D-2708-FE0F",
		"value": "👩🏾‍✈️",
		"descriptor": "Woman Pilot: Medium-Dark Skin Tone"
	},
	"1F469-1F3FF": {
		"key": "1F469-1F3FF",
		"value": "👩🏿",
		"descriptor": "Woman: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F33E": {
		"key": "1F469-1F3FF-200D-1F33E",
		"value": "👩🏿‍🌾",
		"descriptor": "Woman Farmer: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F373": {
		"key": "1F469-1F3FF-200D-1F373",
		"value": "👩🏿‍🍳",
		"descriptor": "Woman Cook: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F393": {
		"key": "1F469-1F3FF-200D-1F393",
		"value": "👩🏿‍🎓",
		"descriptor": "Woman Student: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F3A4": {
		"key": "1F469-1F3FF-200D-1F3A4",
		"value": "👩🏿‍🎤",
		"descriptor": "Woman Singer: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F3A8": {
		"key": "1F469-1F3FF-200D-1F3A8",
		"value": "👩🏿‍🎨",
		"descriptor": "Woman Artist: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F3EB": {
		"key": "1F469-1F3FF-200D-1F3EB",
		"value": "👩🏿‍🏫",
		"descriptor": "Woman Teacher: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F3ED": {
		"key": "1F469-1F3FF-200D-1F3ED",
		"value": "👩🏿‍🏭",
		"descriptor": "Woman Factory Worker: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F4BB": {
		"key": "1F469-1F3FF-200D-1F4BB",
		"value": "👩🏿‍💻",
		"descriptor": "Woman Technologist: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F4BC": {
		"key": "1F469-1F3FF-200D-1F4BC",
		"value": "👩🏿‍💼",
		"descriptor": "Woman Office Worker: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F527": {
		"key": "1F469-1F3FF-200D-1F527",
		"value": "👩🏿‍🔧",
		"descriptor": "Woman Mechanic: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F52C": {
		"key": "1F469-1F3FF-200D-1F52C",
		"value": "👩🏿‍🔬",
		"descriptor": "Woman Scientist: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F680": {
		"key": "1F469-1F3FF-200D-1F680",
		"value": "👩🏿‍🚀",
		"descriptor": "Woman Astronaut: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F692": {
		"key": "1F469-1F3FF-200D-1F692",
		"value": "👩🏿‍🚒",
		"descriptor": "Woman Firefighter: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FB": {
		"key": "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FB",
		"value": "👩🏿‍🤝‍👨🏻",
		"descriptor": "Woman and Man Holding Hands: Dark Skin Tone, Light Skin Tone"
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FC": {
		"key": "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FC",
		"value": "👩🏿‍🤝‍👨🏼",
		"descriptor": "Woman and Man Holding Hands: Dark Skin Tone, Medium-Light Skin Tone"
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FD": {
		"key": "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FD",
		"value": "👩🏿‍🤝‍👨🏽",
		"descriptor": "Woman and Man Holding Hands: Dark Skin Tone, Medium Skin Tone"
	},
	"1F469-1F3FF-200D-1F91D-200D-1F468-1F3FE": {
		"key": "1F469-1F3FF-200D-1F91D-200D-1F468-1F3FE",
		"value": "👩🏿‍🤝‍👨🏾",
		"descriptor": "Woman and Man Holding Hands: Dark Skin Tone, Medium-Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FB": {
		"key": "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FB",
		"value": "👩🏿‍🤝‍👩🏻",
		"descriptor": "Women Holding Hands: Dark Skin Tone, Light Skin Tone"
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FC": {
		"key": "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FC",
		"value": "👩🏿‍🤝‍👩🏼",
		"descriptor": "Women Holding Hands: Dark Skin Tone, Medium-Light Skin Tone"
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FD": {
		"key": "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FD",
		"value": "👩🏿‍🤝‍👩🏽",
		"descriptor": "Women Holding Hands: Dark Skin Tone, Medium Skin Tone"
	},
	"1F469-1F3FF-200D-1F91D-200D-1F469-1F3FE": {
		"key": "1F469-1F3FF-200D-1F91D-200D-1F469-1F3FE",
		"value": "👩🏿‍🤝‍👩🏾",
		"descriptor": "Women Holding Hands: Dark Skin Tone, Medium-Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F9AF": {
		"key": "1F469-1F3FF-200D-1F9AF",
		"value": "👩🏿‍🦯",
		"descriptor": "Woman With Probing Cane: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F9B0": {
		"key": "1F469-1F3FF-200D-1F9B0",
		"value": "👩🏿‍🦰",
		"descriptor": "Woman: Dark Skin Tone, Red Hair"
	},
	"1F469-1F3FF-200D-1F9B1": {
		"key": "1F469-1F3FF-200D-1F9B1",
		"value": "👩🏿‍🦱",
		"descriptor": "Woman: Dark Skin Tone, Curly Hair"
	},
	"1F469-1F3FF-200D-1F9B2": {
		"key": "1F469-1F3FF-200D-1F9B2",
		"value": "👩🏿‍🦲",
		"descriptor": "Woman: Dark Skin Tone, Bald"
	},
	"1F469-1F3FF-200D-1F9B3": {
		"key": "1F469-1F3FF-200D-1F9B3",
		"value": "👩🏿‍🦳",
		"descriptor": "Woman: Dark Skin Tone, White Hair"
	},
	"1F469-1F3FF-200D-1F9BC": {
		"key": "1F469-1F3FF-200D-1F9BC",
		"value": "👩🏿‍🦼",
		"descriptor": "Woman in Motorized Wheelchair: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-1F9BD": {
		"key": "1F469-1F3FF-200D-1F9BD",
		"value": "👩🏿‍🦽",
		"descriptor": "Woman in Manual Wheelchair: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-2695-FE0F": {
		"key": "1F469-1F3FF-200D-2695-FE0F",
		"value": "👩🏿‍⚕️",
		"descriptor": "Woman Health Worker: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-2696-FE0F": {
		"key": "1F469-1F3FF-200D-2696-FE0F",
		"value": "👩🏿‍⚖️",
		"descriptor": "Woman Judge: Dark Skin Tone"
	},
	"1F469-1F3FF-200D-2708-FE0F": {
		"key": "1F469-1F3FF-200D-2708-FE0F",
		"value": "👩🏿‍✈️",
		"descriptor": "Woman Pilot: Dark Skin Tone"
	},
	"1F469-200D-1F33E": {
		"key": "1F469-200D-1F33E",
		"value": "👩‍🌾",
		"descriptor": "Woman Farmer"
	},
	"1F469-200D-1F373": {
		"key": "1F469-200D-1F373",
		"value": "👩‍🍳",
		"descriptor": "Woman Cook"
	},
	"1F469-200D-1F393": {
		"key": "1F469-200D-1F393",
		"value": "👩‍🎓",
		"descriptor": "Woman Student"
	},
	"1F469-200D-1F3A4": {
		"key": "1F469-200D-1F3A4",
		"value": "👩‍🎤",
		"descriptor": "Woman Singer"
	},
	"1F469-200D-1F3A8": {
		"key": "1F469-200D-1F3A8",
		"value": "👩‍🎨",
		"descriptor": "Woman Artist"
	},
	"1F469-200D-1F3EB": {
		"key": "1F469-200D-1F3EB",
		"value": "👩‍🏫",
		"descriptor": "Woman Teacher"
	},
	"1F469-200D-1F3ED": {
		"key": "1F469-200D-1F3ED",
		"value": "👩‍🏭",
		"descriptor": "Woman Factory Worker"
	},
	"1F469-200D-1F466": {
		"key": "1F469-200D-1F466",
		"value": "👩‍👦",
		"descriptor": "Family: Woman, Boy"
	},
	"1F469-200D-1F466-200D-1F466": {
		"key": "1F469-200D-1F466-200D-1F466",
		"value": "👩‍👦‍👦",
		"descriptor": "Family: Woman, Boy, Boy"
	},
	"1F469-200D-1F467": {
		"key": "1F469-200D-1F467",
		"value": "👩‍👧",
		"descriptor": "Family: Woman, Girl"
	},
	"1F469-200D-1F467-200D-1F466": {
		"key": "1F469-200D-1F467-200D-1F466",
		"value": "👩‍👧‍👦",
		"descriptor": "Family: Woman, Girl, Boy"
	},
	"1F469-200D-1F467-200D-1F467": {
		"key": "1F469-200D-1F467-200D-1F467",
		"value": "👩‍👧‍👧",
		"descriptor": "Family: Woman, Girl, Girl"
	},
	"1F469-200D-1F469-200D-1F466": {
		"key": "1F469-200D-1F469-200D-1F466",
		"value": "👩‍👩‍👦",
		"descriptor": "Family: Woman, Woman, Boy"
	},
	"1F469-200D-1F469-200D-1F466-200D-1F466": {
		"key": "1F469-200D-1F469-200D-1F466-200D-1F466",
		"value": "👩‍👩‍👦‍👦",
		"descriptor": "Family: Woman, Woman, Boy, Boy"
	},
	"1F469-200D-1F469-200D-1F467": {
		"key": "1F469-200D-1F469-200D-1F467",
		"value": "👩‍👩‍👧",
		"descriptor": "Family: Woman, Woman, Girl"
	},
	"1F469-200D-1F469-200D-1F467-200D-1F466": {
		"key": "1F469-200D-1F469-200D-1F467-200D-1F466",
		"value": "👩‍👩‍👧‍👦",
		"descriptor": "Family: Woman, Woman, Girl, Boy"
	},
	"1F469-200D-1F469-200D-1F467-200D-1F467": {
		"key": "1F469-200D-1F469-200D-1F467-200D-1F467",
		"value": "👩‍👩‍👧‍👧",
		"descriptor": "Family: Woman, Woman, Girl, Girl"
	},
	"1F469-200D-1F4BB": {
		"key": "1F469-200D-1F4BB",
		"value": "👩‍💻",
		"descriptor": "Woman Technologist"
	},
	"1F469-200D-1F4BC": {
		"key": "1F469-200D-1F4BC",
		"value": "👩‍💼",
		"descriptor": "Woman Office Worker"
	},
	"1F469-200D-1F527": {
		"key": "1F469-200D-1F527",
		"value": "👩‍🔧",
		"descriptor": "Woman Mechanic"
	},
	"1F469-200D-1F52C": {
		"key": "1F469-200D-1F52C",
		"value": "👩‍🔬",
		"descriptor": "Woman Scientist"
	},
	"1F469-200D-1F680": {
		"key": "1F469-200D-1F680",
		"value": "👩‍🚀",
		"descriptor": "Woman Astronaut"
	},
	"1F469-200D-1F692": {
		"key": "1F469-200D-1F692",
		"value": "👩‍🚒",
		"descriptor": "Woman Firefighter"
	},
	"1F469-200D-1F9AF": {
		"key": "1F469-200D-1F9AF",
		"value": "👩‍🦯",
		"descriptor": "Woman With Probing Cane"
	},
	"1F469-200D-1F9B0": {
		"key": "1F469-200D-1F9B0",
		"value": "👩‍🦰",
		"descriptor": "Woman: Red Hair"
	},
	"1F469-200D-1F9B1": {
		"key": "1F469-200D-1F9B1",
		"value": "👩‍🦱",
		"descriptor": "Woman: Curly Hair"
	},
	"1F469-200D-1F9B2": {
		"key": "1F469-200D-1F9B2",
		"value": "👩‍🦲",
		"descriptor": "Woman: Bald"
	},
	"1F469-200D-1F9B3": {
		"key": "1F469-200D-1F9B3",
		"value": "👩‍🦳",
		"descriptor": "Woman: White Hair"
	},
	"1F469-200D-1F9BC": {
		"key": "1F469-200D-1F9BC",
		"value": "👩‍🦼",
		"descriptor": "Woman in Motorized Wheelchair"
	},
	"1F469-200D-1F9BD": {
		"key": "1F469-200D-1F9BD",
		"value": "👩‍🦽",
		"descriptor": "Woman in Manual Wheelchair"
	},
	"1F469-200D-2695-FE0F": {
		"key": "1F469-200D-2695-FE0F",
		"value": "👩‍⚕️",
		"descriptor": "Woman Health Worker"
	},
	"1F469-200D-2696-FE0F": {
		"key": "1F469-200D-2696-FE0F",
		"value": "👩‍⚖️",
		"descriptor": "Woman Judge"
	},
	"1F469-200D-2708-FE0F": {
		"key": "1F469-200D-2708-FE0F",
		"value": "👩‍✈️",
		"descriptor": "Woman Pilot"
	},
	"1F469-200D-2764-FE0F-200D-1F468": {
		"key": "1F469-200D-2764-FE0F-200D-1F468",
		"value": "👩‍❤️‍👨",
		"descriptor": "Couple With Heart: Woman, Man"
	},
	"1F469-200D-2764-FE0F-200D-1F469": {
		"key": "1F469-200D-2764-FE0F-200D-1F469",
		"value": "👩‍❤️‍👩",
		"descriptor": "Couple With Heart: Woman, Woman"
	},
	"1F469-200D-2764-FE0F-200D-1F48B-200D-1F468": {
		"key": "1F469-200D-2764-FE0F-200D-1F48B-200D-1F468",
		"value": "👩‍❤️‍💋‍👨",
		"descriptor": "Kiss: Woman, Man"
	},
	"1F469-200D-2764-FE0F-200D-1F48B-200D-1F469": {
		"key": "1F469-200D-2764-FE0F-200D-1F48B-200D-1F469",
		"value": "👩‍❤️‍💋‍👩",
		"descriptor": "Kiss: Woman, Woman"
	},
	"1F46A": {
		"key": "1F46A",
		"value": "👪",
		"descriptor": "Family"
	},
	"1F46B": {
		"key": "1F46B",
		"value": "👫",
		"descriptor": "Woman and Man Holding Hands"
	},
	"1F46B-1F3FB": {
		"key": "1F46B-1F3FB",
		"value": "👫🏻",
		"descriptor": "Woman and Man Holding Hands: Light Skin Tone"
	},
	"1F46B-1F3FC": {
		"key": "1F46B-1F3FC",
		"value": "👫🏼",
		"descriptor": "Woman and Man Holding Hands: Medium-Light Skin Tone"
	},
	"1F46B-1F3FD": {
		"key": "1F46B-1F3FD",
		"value": "👫🏽",
		"descriptor": "Woman and Man Holding Hands: Medium Skin Tone"
	},
	"1F46B-1F3FE": {
		"key": "1F46B-1F3FE",
		"value": "👫🏾",
		"descriptor": "Woman and Man Holding Hands: Medium-Dark Skin Tone"
	},
	"1F46B-1F3FF": {
		"key": "1F46B-1F3FF",
		"value": "👫🏿",
		"descriptor": "Woman and Man Holding Hands: Dark Skin Tone"
	},
	"1F46C": {
		"key": "1F46C",
		"value": "👬",
		"descriptor": "Men Holding Hands"
	},
	"1F46C-1F3FB": {
		"key": "1F46C-1F3FB",
		"value": "👬🏻",
		"descriptor": "Men Holding Hands: Light Skin Tone"
	},
	"1F46C-1F3FC": {
		"key": "1F46C-1F3FC",
		"value": "👬🏼",
		"descriptor": "Men Holding Hands: Medium-Light Skin Tone"
	},
	"1F46C-1F3FD": {
		"key": "1F46C-1F3FD",
		"value": "👬🏽",
		"descriptor": "Men Holding Hands: Medium Skin Tone"
	},
	"1F46C-1F3FE": {
		"key": "1F46C-1F3FE",
		"value": "👬🏾",
		"descriptor": "Men Holding Hands: Medium-Dark Skin Tone"
	},
	"1F46C-1F3FF": {
		"key": "1F46C-1F3FF",
		"value": "👬🏿",
		"descriptor": "Men Holding Hands: Dark Skin Tone"
	},
	"1F46D": {
		"key": "1F46D",
		"value": "👭",
		"descriptor": "Women Holding Hands"
	},
	"1F46D-1F3FB": {
		"key": "1F46D-1F3FB",
		"value": "👭🏻",
		"descriptor": "Women Holding Hands: Light Skin Tone"
	},
	"1F46D-1F3FC": {
		"key": "1F46D-1F3FC",
		"value": "👭🏼",
		"descriptor": "Women Holding Hands: Medium-Light Skin Tone"
	},
	"1F46D-1F3FD": {
		"key": "1F46D-1F3FD",
		"value": "👭🏽",
		"descriptor": "Women Holding Hands: Medium Skin Tone"
	},
	"1F46D-1F3FE": {
		"key": "1F46D-1F3FE",
		"value": "👭🏾",
		"descriptor": "Women Holding Hands: Medium-Dark Skin Tone"
	},
	"1F46D-1F3FF": {
		"key": "1F46D-1F3FF",
		"value": "👭🏿",
		"descriptor": "Women Holding Hands: Dark Skin Tone"
	},
	"1F46E": {
		"key": "1F46E",
		"value": "👮",
		"descriptor": "Police Officer"
	},
	"1F46E-1F3FB": {
		"key": "1F46E-1F3FB",
		"value": "👮🏻",
		"descriptor": "Police Officer: Light Skin Tone"
	},
	"1F46E-1F3FB-200D-2640-FE0F": {
		"key": "1F46E-1F3FB-200D-2640-FE0F",
		"value": "👮🏻‍♀️",
		"descriptor": "Woman Police Officer: Light Skin Tone"
	},
	"1F46E-1F3FB-200D-2642-FE0F": {
		"key": "1F46E-1F3FB-200D-2642-FE0F",
		"value": "👮🏻‍♂️",
		"descriptor": "Man Police Officer: Light Skin Tone"
	},
	"1F46E-1F3FC": {
		"key": "1F46E-1F3FC",
		"value": "👮🏼",
		"descriptor": "Police Officer: Medium-Light Skin Tone"
	},
	"1F46E-1F3FC-200D-2640-FE0F": {
		"key": "1F46E-1F3FC-200D-2640-FE0F",
		"value": "👮🏼‍♀️",
		"descriptor": "Woman Police Officer: Medium-Light Skin Tone"
	},
	"1F46E-1F3FC-200D-2642-FE0F": {
		"key": "1F46E-1F3FC-200D-2642-FE0F",
		"value": "👮🏼‍♂️",
		"descriptor": "Man Police Officer: Medium-Light Skin Tone"
	},
	"1F46E-1F3FD": {
		"key": "1F46E-1F3FD",
		"value": "👮🏽",
		"descriptor": "Police Officer: Medium Skin Tone"
	},
	"1F46E-1F3FD-200D-2640-FE0F": {
		"key": "1F46E-1F3FD-200D-2640-FE0F",
		"value": "👮🏽‍♀️",
		"descriptor": "Woman Police Officer: Medium Skin Tone"
	},
	"1F46E-1F3FD-200D-2642-FE0F": {
		"key": "1F46E-1F3FD-200D-2642-FE0F",
		"value": "👮🏽‍♂️",
		"descriptor": "Man Police Officer: Medium Skin Tone"
	},
	"1F46E-1F3FE": {
		"key": "1F46E-1F3FE",
		"value": "👮🏾",
		"descriptor": "Police Officer: Medium-Dark Skin Tone"
	},
	"1F46E-1F3FE-200D-2640-FE0F": {
		"key": "1F46E-1F3FE-200D-2640-FE0F",
		"value": "👮🏾‍♀️",
		"descriptor": "Woman Police Officer: Medium-Dark Skin Tone"
	},
	"1F46E-1F3FE-200D-2642-FE0F": {
		"key": "1F46E-1F3FE-200D-2642-FE0F",
		"value": "👮🏾‍♂️",
		"descriptor": "Man Police Officer: Medium-Dark Skin Tone"
	},
	"1F46E-1F3FF": {
		"key": "1F46E-1F3FF",
		"value": "👮🏿",
		"descriptor": "Police Officer: Dark Skin Tone"
	},
	"1F46E-1F3FF-200D-2640-FE0F": {
		"key": "1F46E-1F3FF-200D-2640-FE0F",
		"value": "👮🏿‍♀️",
		"descriptor": "Woman Police Officer: Dark Skin Tone"
	},
	"1F46E-1F3FF-200D-2642-FE0F": {
		"key": "1F46E-1F3FF-200D-2642-FE0F",
		"value": "👮🏿‍♂️",
		"descriptor": "Man Police Officer: Dark Skin Tone"
	},
	"1F46E-200D-2640-FE0F": {
		"key": "1F46E-200D-2640-FE0F",
		"value": "👮‍♀️",
		"descriptor": "Woman Police Officer"
	},
	"1F46E-200D-2642-FE0F": {
		"key": "1F46E-200D-2642-FE0F",
		"value": "👮‍♂️",
		"descriptor": "Man Police Officer"
	},
	"1F46F": {
		"key": "1F46F",
		"value": "👯",
		"descriptor": "People With Bunny Ears"
	},
	"1F46F-200D-2640-FE0F": {
		"key": "1F46F-200D-2640-FE0F",
		"value": "👯‍♀️",
		"descriptor": "Women With Bunny Ears"
	},
	"1F46F-200D-2642-FE0F": {
		"key": "1F46F-200D-2642-FE0F",
		"value": "👯‍♂️",
		"descriptor": "Men With Bunny Ears"
	},
	"1F470": {
		"key": "1F470",
		"value": "👰",
		"descriptor": "Bride With Veil"
	},
	"1F470-1F3FB": {
		"key": "1F470-1F3FB",
		"value": "👰🏻",
		"descriptor": "Bride With Veil: Light Skin Tone"
	},
	"1F470-1F3FC": {
		"key": "1F470-1F3FC",
		"value": "👰🏼",
		"descriptor": "Bride With Veil: Medium-Light Skin Tone"
	},
	"1F470-1F3FD": {
		"key": "1F470-1F3FD",
		"value": "👰🏽",
		"descriptor": "Bride With Veil: Medium Skin Tone"
	},
	"1F470-1F3FE": {
		"key": "1F470-1F3FE",
		"value": "👰🏾",
		"descriptor": "Bride With Veil: Medium-Dark Skin Tone"
	},
	"1F470-1F3FF": {
		"key": "1F470-1F3FF",
		"value": "👰🏿",
		"descriptor": "Bride With Veil: Dark Skin Tone"
	},
	"1F471": {
		"key": "1F471",
		"value": "👱",
		"descriptor": "Person: Blond Hair"
	},
	"1F471-1F3FB": {
		"key": "1F471-1F3FB",
		"value": "👱🏻",
		"descriptor": "Person: Light Skin Tone, Blond Hair"
	},
	"1F471-1F3FB-200D-2640-FE0F": {
		"key": "1F471-1F3FB-200D-2640-FE0F",
		"value": "👱🏻‍♀️",
		"descriptor": "Woman: Light Skin Tone, Blond Hair"
	},
	"1F471-1F3FB-200D-2642-FE0F": {
		"key": "1F471-1F3FB-200D-2642-FE0F",
		"value": "👱🏻‍♂️",
		"descriptor": "Man: Light Skin Tone, Blond Hair"
	},
	"1F471-1F3FC": {
		"key": "1F471-1F3FC",
		"value": "👱🏼",
		"descriptor": "Person: Medium-Light Skin Tone, Blond Hair"
	},
	"1F471-1F3FC-200D-2640-FE0F": {
		"key": "1F471-1F3FC-200D-2640-FE0F",
		"value": "👱🏼‍♀️",
		"descriptor": "Woman: Medium-Light Skin Tone, Blond Hair"
	},
	"1F471-1F3FC-200D-2642-FE0F": {
		"key": "1F471-1F3FC-200D-2642-FE0F",
		"value": "👱🏼‍♂️",
		"descriptor": "Man: Medium-Light Skin Tone, Blond Hair"
	},
	"1F471-1F3FD": {
		"key": "1F471-1F3FD",
		"value": "👱🏽",
		"descriptor": "Person: Medium Skin Tone, Blond Hair"
	},
	"1F471-1F3FD-200D-2640-FE0F": {
		"key": "1F471-1F3FD-200D-2640-FE0F",
		"value": "👱🏽‍♀️",
		"descriptor": "Woman: Medium Skin Tone, Blond Hair"
	},
	"1F471-1F3FD-200D-2642-FE0F": {
		"key": "1F471-1F3FD-200D-2642-FE0F",
		"value": "👱🏽‍♂️",
		"descriptor": "Man: Medium Skin Tone, Blond Hair"
	},
	"1F471-1F3FE": {
		"key": "1F471-1F3FE",
		"value": "👱🏾",
		"descriptor": "Person: Medium-Dark Skin Tone, Blond Hair"
	},
	"1F471-1F3FE-200D-2640-FE0F": {
		"key": "1F471-1F3FE-200D-2640-FE0F",
		"value": "👱🏾‍♀️",
		"descriptor": "Woman: Medium-Dark Skin Tone, Blond Hair"
	},
	"1F471-1F3FE-200D-2642-FE0F": {
		"key": "1F471-1F3FE-200D-2642-FE0F",
		"value": "👱🏾‍♂️",
		"descriptor": "Man: Medium-Dark Skin Tone, Blond Hair"
	},
	"1F471-1F3FF": {
		"key": "1F471-1F3FF",
		"value": "👱🏿",
		"descriptor": "Person: Dark Skin Tone, Blond Hair"
	},
	"1F471-1F3FF-200D-2640-FE0F": {
		"key": "1F471-1F3FF-200D-2640-FE0F",
		"value": "👱🏿‍♀️",
		"descriptor": "Woman: Dark Skin Tone, Blond Hair"
	},
	"1F471-1F3FF-200D-2642-FE0F": {
		"key": "1F471-1F3FF-200D-2642-FE0F",
		"value": "👱🏿‍♂️",
		"descriptor": "Man: Dark Skin Tone, Blond Hair"
	},
	"1F471-200D-2640-FE0F": {
		"key": "1F471-200D-2640-FE0F",
		"value": "👱‍♀️",
		"descriptor": "Woman: Blond Hair"
	},
	"1F471-200D-2642-FE0F": {
		"key": "1F471-200D-2642-FE0F",
		"value": "👱‍♂️",
		"descriptor": "Man: Blond Hair"
	},
	"1F472": {
		"key": "1F472",
		"value": "👲",
		"descriptor": "Man With Skullcap"
	},
	"1F472-1F3FB": {
		"key": "1F472-1F3FB",
		"value": "👲🏻",
		"descriptor": "Man With Skullcap: Light Skin Tone"
	},
	"1F472-1F3FC": {
		"key": "1F472-1F3FC",
		"value": "👲🏼",
		"descriptor": "Man With Skullcap: Medium-Light Skin Tone"
	},
	"1F472-1F3FD": {
		"key": "1F472-1F3FD",
		"value": "👲🏽",
		"descriptor": "Man With Skullcap: Medium Skin Tone"
	},
	"1F472-1F3FE": {
		"key": "1F472-1F3FE",
		"value": "👲🏾",
		"descriptor": "Man With Skullcap: Medium-Dark Skin Tone"
	},
	"1F472-1F3FF": {
		"key": "1F472-1F3FF",
		"value": "👲🏿",
		"descriptor": "Man With Skullcap: Dark Skin Tone"
	},
	"1F473": {
		"key": "1F473",
		"value": "👳",
		"descriptor": "Person Wearing Turban"
	},
	"1F473-1F3FB": {
		"key": "1F473-1F3FB",
		"value": "👳🏻",
		"descriptor": "Person Wearing Turban: Light Skin Tone"
	},
	"1F473-1F3FB-200D-2640-FE0F": {
		"key": "1F473-1F3FB-200D-2640-FE0F",
		"value": "👳🏻‍♀️",
		"descriptor": "Woman Wearing Turban: Light Skin Tone"
	},
	"1F473-1F3FB-200D-2642-FE0F": {
		"key": "1F473-1F3FB-200D-2642-FE0F",
		"value": "👳🏻‍♂️",
		"descriptor": "Man Wearing Turban: Light Skin Tone"
	},
	"1F473-1F3FC": {
		"key": "1F473-1F3FC",
		"value": "👳🏼",
		"descriptor": "Person Wearing Turban: Medium-Light Skin Tone"
	},
	"1F473-1F3FC-200D-2640-FE0F": {
		"key": "1F473-1F3FC-200D-2640-FE0F",
		"value": "👳🏼‍♀️",
		"descriptor": "Woman Wearing Turban: Medium-Light Skin Tone"
	},
	"1F473-1F3FC-200D-2642-FE0F": {
		"key": "1F473-1F3FC-200D-2642-FE0F",
		"value": "👳🏼‍♂️",
		"descriptor": "Man Wearing Turban: Medium-Light Skin Tone"
	},
	"1F473-1F3FD": {
		"key": "1F473-1F3FD",
		"value": "👳🏽",
		"descriptor": "Person Wearing Turban: Medium Skin Tone"
	},
	"1F473-1F3FD-200D-2640-FE0F": {
		"key": "1F473-1F3FD-200D-2640-FE0F",
		"value": "👳🏽‍♀️",
		"descriptor": "Woman Wearing Turban: Medium Skin Tone"
	},
	"1F473-1F3FD-200D-2642-FE0F": {
		"key": "1F473-1F3FD-200D-2642-FE0F",
		"value": "👳🏽‍♂️",
		"descriptor": "Man Wearing Turban: Medium Skin Tone"
	},
	"1F473-1F3FE": {
		"key": "1F473-1F3FE",
		"value": "👳🏾",
		"descriptor": "Person Wearing Turban: Medium-Dark Skin Tone"
	},
	"1F473-1F3FE-200D-2640-FE0F": {
		"key": "1F473-1F3FE-200D-2640-FE0F",
		"value": "👳🏾‍♀️",
		"descriptor": "Woman Wearing Turban: Medium-Dark Skin Tone"
	},
	"1F473-1F3FE-200D-2642-FE0F": {
		"key": "1F473-1F3FE-200D-2642-FE0F",
		"value": "👳🏾‍♂️",
		"descriptor": "Man Wearing Turban: Medium-Dark Skin Tone"
	},
	"1F473-1F3FF": {
		"key": "1F473-1F3FF",
		"value": "👳🏿",
		"descriptor": "Person Wearing Turban: Dark Skin Tone"
	},
	"1F473-1F3FF-200D-2640-FE0F": {
		"key": "1F473-1F3FF-200D-2640-FE0F",
		"value": "👳🏿‍♀️",
		"descriptor": "Woman Wearing Turban: Dark Skin Tone"
	},
	"1F473-1F3FF-200D-2642-FE0F": {
		"key": "1F473-1F3FF-200D-2642-FE0F",
		"value": "👳🏿‍♂️",
		"descriptor": "Man Wearing Turban: Dark Skin Tone"
	},
	"1F473-200D-2640-FE0F": {
		"key": "1F473-200D-2640-FE0F",
		"value": "👳‍♀️",
		"descriptor": "Woman Wearing Turban"
	},
	"1F473-200D-2642-FE0F": {
		"key": "1F473-200D-2642-FE0F",
		"value": "👳‍♂️",
		"descriptor": "Man Wearing Turban"
	},
	"1F474": {
		"key": "1F474",
		"value": "👴",
		"descriptor": "Old Man"
	},
	"1F474-1F3FB": {
		"key": "1F474-1F3FB",
		"value": "👴🏻",
		"descriptor": "Old Man: Light Skin Tone"
	},
	"1F474-1F3FC": {
		"key": "1F474-1F3FC",
		"value": "👴🏼",
		"descriptor": "Old Man: Medium-Light Skin Tone"
	},
	"1F474-1F3FD": {
		"key": "1F474-1F3FD",
		"value": "👴🏽",
		"descriptor": "Old Man: Medium Skin Tone"
	},
	"1F474-1F3FE": {
		"key": "1F474-1F3FE",
		"value": "👴🏾",
		"descriptor": "Old Man: Medium-Dark Skin Tone"
	},
	"1F474-1F3FF": {
		"key": "1F474-1F3FF",
		"value": "👴🏿",
		"descriptor": "Old Man: Dark Skin Tone"
	},
	"1F475": {
		"key": "1F475",
		"value": "👵",
		"descriptor": "Old Woman"
	},
	"1F475-1F3FB": {
		"key": "1F475-1F3FB",
		"value": "👵🏻",
		"descriptor": "Old Woman: Light Skin Tone"
	},
	"1F475-1F3FC": {
		"key": "1F475-1F3FC",
		"value": "👵🏼",
		"descriptor": "Old Woman: Medium-Light Skin Tone"
	},
	"1F475-1F3FD": {
		"key": "1F475-1F3FD",
		"value": "👵🏽",
		"descriptor": "Old Woman: Medium Skin Tone"
	},
	"1F475-1F3FE": {
		"key": "1F475-1F3FE",
		"value": "👵🏾",
		"descriptor": "Old Woman: Medium-Dark Skin Tone"
	},
	"1F475-1F3FF": {
		"key": "1F475-1F3FF",
		"value": "👵🏿",
		"descriptor": "Old Woman: Dark Skin Tone"
	},
	"1F476": {
		"key": "1F476",
		"value": "👶",
		"descriptor": "Baby"
	},
	"1F476-1F3FB": {
		"key": "1F476-1F3FB",
		"value": "👶🏻",
		"descriptor": "Baby: Light Skin Tone"
	},
	"1F476-1F3FC": {
		"key": "1F476-1F3FC",
		"value": "👶🏼",
		"descriptor": "Baby: Medium-Light Skin Tone"
	},
	"1F476-1F3FD": {
		"key": "1F476-1F3FD",
		"value": "👶🏽",
		"descriptor": "Baby: Medium Skin Tone"
	},
	"1F476-1F3FE": {
		"key": "1F476-1F3FE",
		"value": "👶🏾",
		"descriptor": "Baby: Medium-Dark Skin Tone"
	},
	"1F476-1F3FF": {
		"key": "1F476-1F3FF",
		"value": "👶🏿",
		"descriptor": "Baby: Dark Skin Tone"
	},
	"1F477": {
		"key": "1F477",
		"value": "👷",
		"descriptor": "Construction Worker"
	},
	"1F477-1F3FB": {
		"key": "1F477-1F3FB",
		"value": "👷🏻",
		"descriptor": "Construction Worker: Light Skin Tone"
	},
	"1F477-1F3FB-200D-2640-FE0F": {
		"key": "1F477-1F3FB-200D-2640-FE0F",
		"value": "👷🏻‍♀️",
		"descriptor": "Woman Construction Worker: Light Skin Tone"
	},
	"1F477-1F3FB-200D-2642-FE0F": {
		"key": "1F477-1F3FB-200D-2642-FE0F",
		"value": "👷🏻‍♂️",
		"descriptor": "Man Construction Worker: Light Skin Tone"
	},
	"1F477-1F3FC": {
		"key": "1F477-1F3FC",
		"value": "👷🏼",
		"descriptor": "Construction Worker: Medium-Light Skin Tone"
	},
	"1F477-1F3FC-200D-2640-FE0F": {
		"key": "1F477-1F3FC-200D-2640-FE0F",
		"value": "👷🏼‍♀️",
		"descriptor": "Woman Construction Worker: Medium-Light Skin Tone"
	},
	"1F477-1F3FC-200D-2642-FE0F": {
		"key": "1F477-1F3FC-200D-2642-FE0F",
		"value": "👷🏼‍♂️",
		"descriptor": "Man Construction Worker: Medium-Light Skin Tone"
	},
	"1F477-1F3FD": {
		"key": "1F477-1F3FD",
		"value": "👷🏽",
		"descriptor": "Construction Worker: Medium Skin Tone"
	},
	"1F477-1F3FD-200D-2640-FE0F": {
		"key": "1F477-1F3FD-200D-2640-FE0F",
		"value": "👷🏽‍♀️",
		"descriptor": "Woman Construction Worker: Medium Skin Tone"
	},
	"1F477-1F3FD-200D-2642-FE0F": {
		"key": "1F477-1F3FD-200D-2642-FE0F",
		"value": "👷🏽‍♂️",
		"descriptor": "Man Construction Worker: Medium Skin Tone"
	},
	"1F477-1F3FE": {
		"key": "1F477-1F3FE",
		"value": "👷🏾",
		"descriptor": "Construction Worker: Medium-Dark Skin Tone"
	},
	"1F477-1F3FE-200D-2640-FE0F": {
		"key": "1F477-1F3FE-200D-2640-FE0F",
		"value": "👷🏾‍♀️",
		"descriptor": "Woman Construction Worker: Medium-Dark Skin Tone"
	},
	"1F477-1F3FE-200D-2642-FE0F": {
		"key": "1F477-1F3FE-200D-2642-FE0F",
		"value": "👷🏾‍♂️",
		"descriptor": "Man Construction Worker: Medium-Dark Skin Tone"
	},
	"1F477-1F3FF": {
		"key": "1F477-1F3FF",
		"value": "👷🏿",
		"descriptor": "Construction Worker: Dark Skin Tone"
	},
	"1F477-1F3FF-200D-2640-FE0F": {
		"key": "1F477-1F3FF-200D-2640-FE0F",
		"value": "👷🏿‍♀️",
		"descriptor": "Woman Construction Worker: Dark Skin Tone"
	},
	"1F477-1F3FF-200D-2642-FE0F": {
		"key": "1F477-1F3FF-200D-2642-FE0F",
		"value": "👷🏿‍♂️",
		"descriptor": "Man Construction Worker: Dark Skin Tone"
	},
	"1F477-200D-2640-FE0F": {
		"key": "1F477-200D-2640-FE0F",
		"value": "👷‍♀️",
		"descriptor": "Woman Construction Worker"
	},
	"1F477-200D-2642-FE0F": {
		"key": "1F477-200D-2642-FE0F",
		"value": "👷‍♂️",
		"descriptor": "Man Construction Worker"
	},
	"1F478": {
		"key": "1F478",
		"value": "👸",
		"descriptor": "Princess"
	},
	"1F478-1F3FB": {
		"key": "1F478-1F3FB",
		"value": "👸🏻",
		"descriptor": "Princess: Light Skin Tone"
	},
	"1F478-1F3FC": {
		"key": "1F478-1F3FC",
		"value": "👸🏼",
		"descriptor": "Princess: Medium-Light Skin Tone"
	},
	"1F478-1F3FD": {
		"key": "1F478-1F3FD",
		"value": "👸🏽",
		"descriptor": "Princess: Medium Skin Tone"
	},
	"1F478-1F3FE": {
		"key": "1F478-1F3FE",
		"value": "👸🏾",
		"descriptor": "Princess: Medium-Dark Skin Tone"
	},
	"1F478-1F3FF": {
		"key": "1F478-1F3FF",
		"value": "👸🏿",
		"descriptor": "Princess: Dark Skin Tone"
	},
	"1F479": {
		"key": "1F479",
		"value": "👹",
		"descriptor": "Ogre"
	},
	"1F47A": {
		"key": "1F47A",
		"value": "👺",
		"descriptor": "Goblin"
	},
	"1F47B": {
		"key": "1F47B",
		"value": "👻",
		"descriptor": "Ghost"
	},
	"1F47C": {
		"key": "1F47C",
		"value": "👼",
		"descriptor": "Baby Angel"
	},
	"1F47C-1F3FB": {
		"key": "1F47C-1F3FB",
		"value": "👼🏻",
		"descriptor": "Baby Angel: Light Skin Tone"
	},
	"1F47C-1F3FC": {
		"key": "1F47C-1F3FC",
		"value": "👼🏼",
		"descriptor": "Baby Angel: Medium-Light Skin Tone"
	},
	"1F47C-1F3FD": {
		"key": "1F47C-1F3FD",
		"value": "👼🏽",
		"descriptor": "Baby Angel: Medium Skin Tone"
	},
	"1F47C-1F3FE": {
		"key": "1F47C-1F3FE",
		"value": "👼🏾",
		"descriptor": "Baby Angel: Medium-Dark Skin Tone"
	},
	"1F47C-1F3FF": {
		"key": "1F47C-1F3FF",
		"value": "👼🏿",
		"descriptor": "Baby Angel: Dark Skin Tone"
	},
	"1F47D": {
		"key": "1F47D",
		"value": "👽",
		"descriptor": "Alien"
	},
	"1F47E": {
		"key": "1F47E",
		"value": "👾",
		"descriptor": "Alien Monster"
	},
	"1F47F": {
		"key": "1F47F",
		"value": "👿",
		"descriptor": "Angry Face With Horns"
	},
	"1F480": {
		"key": "1F480",
		"value": "💀",
		"descriptor": "Skull"
	},
	"1F481": {
		"key": "1F481",
		"value": "💁",
		"descriptor": "Person Tipping Hand"
	},
	"1F481-1F3FB": {
		"key": "1F481-1F3FB",
		"value": "💁🏻",
		"descriptor": "Person Tipping Hand: Light Skin Tone"
	},
	"1F481-1F3FB-200D-2640-FE0F": {
		"key": "1F481-1F3FB-200D-2640-FE0F",
		"value": "💁🏻‍♀️",
		"descriptor": "Woman Tipping Hand: Light Skin Tone"
	},
	"1F481-1F3FB-200D-2642-FE0F": {
		"key": "1F481-1F3FB-200D-2642-FE0F",
		"value": "💁🏻‍♂️",
		"descriptor": "Man Tipping Hand: Light Skin Tone"
	},
	"1F481-1F3FC": {
		"key": "1F481-1F3FC",
		"value": "💁🏼",
		"descriptor": "Person Tipping Hand: Medium-Light Skin Tone"
	},
	"1F481-1F3FC-200D-2640-FE0F": {
		"key": "1F481-1F3FC-200D-2640-FE0F",
		"value": "💁🏼‍♀️",
		"descriptor": "Woman Tipping Hand: Medium-Light Skin Tone"
	},
	"1F481-1F3FC-200D-2642-FE0F": {
		"key": "1F481-1F3FC-200D-2642-FE0F",
		"value": "💁🏼‍♂️",
		"descriptor": "Man Tipping Hand: Medium-Light Skin Tone"
	},
	"1F481-1F3FD": {
		"key": "1F481-1F3FD",
		"value": "💁🏽",
		"descriptor": "Person Tipping Hand: Medium Skin Tone"
	},
	"1F481-1F3FD-200D-2640-FE0F": {
		"key": "1F481-1F3FD-200D-2640-FE0F",
		"value": "💁🏽‍♀️",
		"descriptor": "Woman Tipping Hand: Medium Skin Tone"
	},
	"1F481-1F3FD-200D-2642-FE0F": {
		"key": "1F481-1F3FD-200D-2642-FE0F",
		"value": "💁🏽‍♂️",
		"descriptor": "Man Tipping Hand: Medium Skin Tone"
	},
	"1F481-1F3FE": {
		"key": "1F481-1F3FE",
		"value": "💁🏾",
		"descriptor": "Person Tipping Hand: Medium-Dark Skin Tone"
	},
	"1F481-1F3FE-200D-2640-FE0F": {
		"key": "1F481-1F3FE-200D-2640-FE0F",
		"value": "💁🏾‍♀️",
		"descriptor": "Woman Tipping Hand: Medium-Dark Skin Tone"
	},
	"1F481-1F3FE-200D-2642-FE0F": {
		"key": "1F481-1F3FE-200D-2642-FE0F",
		"value": "💁🏾‍♂️",
		"descriptor": "Man Tipping Hand: Medium-Dark Skin Tone"
	},
	"1F481-1F3FF": {
		"key": "1F481-1F3FF",
		"value": "💁🏿",
		"descriptor": "Person Tipping Hand: Dark Skin Tone"
	},
	"1F481-1F3FF-200D-2640-FE0F": {
		"key": "1F481-1F3FF-200D-2640-FE0F",
		"value": "💁🏿‍♀️",
		"descriptor": "Woman Tipping Hand: Dark Skin Tone"
	},
	"1F481-1F3FF-200D-2642-FE0F": {
		"key": "1F481-1F3FF-200D-2642-FE0F",
		"value": "💁🏿‍♂️",
		"descriptor": "Man Tipping Hand: Dark Skin Tone"
	},
	"1F481-200D-2640-FE0F": {
		"key": "1F481-200D-2640-FE0F",
		"value": "💁‍♀️",
		"descriptor": "Woman Tipping Hand"
	},
	"1F481-200D-2642-FE0F": {
		"key": "1F481-200D-2642-FE0F",
		"value": "💁‍♂️",
		"descriptor": "Man Tipping Hand"
	},
	"1F482": {
		"key": "1F482",
		"value": "💂",
		"descriptor": "Guard"
	},
	"1F482-1F3FB": {
		"key": "1F482-1F3FB",
		"value": "💂🏻",
		"descriptor": "Guard: Light Skin Tone"
	},
	"1F482-1F3FB-200D-2640-FE0F": {
		"key": "1F482-1F3FB-200D-2640-FE0F",
		"value": "💂🏻‍♀️",
		"descriptor": "Woman Guard: Light Skin Tone"
	},
	"1F482-1F3FB-200D-2642-FE0F": {
		"key": "1F482-1F3FB-200D-2642-FE0F",
		"value": "💂🏻‍♂️",
		"descriptor": "Man Guard: Light Skin Tone"
	},
	"1F482-1F3FC": {
		"key": "1F482-1F3FC",
		"value": "💂🏼",
		"descriptor": "Guard: Medium-Light Skin Tone"
	},
	"1F482-1F3FC-200D-2640-FE0F": {
		"key": "1F482-1F3FC-200D-2640-FE0F",
		"value": "💂🏼‍♀️",
		"descriptor": "Woman Guard: Medium-Light Skin Tone"
	},
	"1F482-1F3FC-200D-2642-FE0F": {
		"key": "1F482-1F3FC-200D-2642-FE0F",
		"value": "💂🏼‍♂️",
		"descriptor": "Man Guard: Medium-Light Skin Tone"
	},
	"1F482-1F3FD": {
		"key": "1F482-1F3FD",
		"value": "💂🏽",
		"descriptor": "Guard: Medium Skin Tone"
	},
	"1F482-1F3FD-200D-2640-FE0F": {
		"key": "1F482-1F3FD-200D-2640-FE0F",
		"value": "💂🏽‍♀️",
		"descriptor": "Woman Guard: Medium Skin Tone"
	},
	"1F482-1F3FD-200D-2642-FE0F": {
		"key": "1F482-1F3FD-200D-2642-FE0F",
		"value": "💂🏽‍♂️",
		"descriptor": "Man Guard: Medium Skin Tone"
	},
	"1F482-1F3FE": {
		"key": "1F482-1F3FE",
		"value": "💂🏾",
		"descriptor": "Guard: Medium-Dark Skin Tone"
	},
	"1F482-1F3FE-200D-2640-FE0F": {
		"key": "1F482-1F3FE-200D-2640-FE0F",
		"value": "💂🏾‍♀️",
		"descriptor": "Woman Guard: Medium-Dark Skin Tone"
	},
	"1F482-1F3FE-200D-2642-FE0F": {
		"key": "1F482-1F3FE-200D-2642-FE0F",
		"value": "💂🏾‍♂️",
		"descriptor": "Man Guard: Medium-Dark Skin Tone"
	},
	"1F482-1F3FF": {
		"key": "1F482-1F3FF",
		"value": "💂🏿",
		"descriptor": "Guard: Dark Skin Tone"
	},
	"1F482-1F3FF-200D-2640-FE0F": {
		"key": "1F482-1F3FF-200D-2640-FE0F",
		"value": "💂🏿‍♀️",
		"descriptor": "Woman Guard: Dark Skin Tone"
	},
	"1F482-1F3FF-200D-2642-FE0F": {
		"key": "1F482-1F3FF-200D-2642-FE0F",
		"value": "💂🏿‍♂️",
		"descriptor": "Man Guard: Dark Skin Tone"
	},
	"1F482-200D-2640-FE0F": {
		"key": "1F482-200D-2640-FE0F",
		"value": "💂‍♀️",
		"descriptor": "Woman Guard"
	},
	"1F482-200D-2642-FE0F": {
		"key": "1F482-200D-2642-FE0F",
		"value": "💂‍♂️",
		"descriptor": "Man Guard"
	},
	"1F483": {
		"key": "1F483",
		"value": "💃",
		"descriptor": "Woman Dancing"
	},
	"1F483-1F3FB": {
		"key": "1F483-1F3FB",
		"value": "💃🏻",
		"descriptor": "Woman Dancing: Light Skin Tone"
	},
	"1F483-1F3FC": {
		"key": "1F483-1F3FC",
		"value": "💃🏼",
		"descriptor": "Woman Dancing: Medium-Light Skin Tone"
	},
	"1F483-1F3FD": {
		"key": "1F483-1F3FD",
		"value": "💃🏽",
		"descriptor": "Woman Dancing: Medium Skin Tone"
	},
	"1F483-1F3FE": {
		"key": "1F483-1F3FE",
		"value": "💃🏾",
		"descriptor": "Woman Dancing: Medium-Dark Skin Tone"
	},
	"1F483-1F3FF": {
		"key": "1F483-1F3FF",
		"value": "💃🏿",
		"descriptor": "Woman Dancing: Dark Skin Tone"
	},
	"1F484": {
		"key": "1F484",
		"value": "💄",
		"descriptor": "Lipstick"
	},
	"1F485": {
		"key": "1F485",
		"value": "💅",
		"descriptor": "Nail Polish"
	},
	"1F485-1F3FB": {
		"key": "1F485-1F3FB",
		"value": "💅🏻",
		"descriptor": "Nail Polish: Light Skin Tone"
	},
	"1F485-1F3FC": {
		"key": "1F485-1F3FC",
		"value": "💅🏼",
		"descriptor": "Nail Polish: Medium-Light Skin Tone"
	},
	"1F485-1F3FD": {
		"key": "1F485-1F3FD",
		"value": "💅🏽",
		"descriptor": "Nail Polish: Medium Skin Tone"
	},
	"1F485-1F3FE": {
		"key": "1F485-1F3FE",
		"value": "💅🏾",
		"descriptor": "Nail Polish: Medium-Dark Skin Tone"
	},
	"1F485-1F3FF": {
		"key": "1F485-1F3FF",
		"value": "💅🏿",
		"descriptor": "Nail Polish: Dark Skin Tone"
	},
	"1F486": {
		"key": "1F486",
		"value": "💆",
		"descriptor": "Person Getting Massage"
	},
	"1F486-1F3FB": {
		"key": "1F486-1F3FB",
		"value": "💆🏻",
		"descriptor": "Person Getting Massage: Light Skin Tone"
	},
	"1F486-1F3FB-200D-2640-FE0F": {
		"key": "1F486-1F3FB-200D-2640-FE0F",
		"value": "💆🏻‍♀️",
		"descriptor": "Woman Getting Massage: Light Skin Tone"
	},
	"1F486-1F3FB-200D-2642-FE0F": {
		"key": "1F486-1F3FB-200D-2642-FE0F",
		"value": "💆🏻‍♂️",
		"descriptor": "Man Getting Massage: Light Skin Tone"
	},
	"1F486-1F3FC": {
		"key": "1F486-1F3FC",
		"value": "💆🏼",
		"descriptor": "Person Getting Massage: Medium-Light Skin Tone"
	},
	"1F486-1F3FC-200D-2640-FE0F": {
		"key": "1F486-1F3FC-200D-2640-FE0F",
		"value": "💆🏼‍♀️",
		"descriptor": "Woman Getting Massage: Medium-Light Skin Tone"
	},
	"1F486-1F3FC-200D-2642-FE0F": {
		"key": "1F486-1F3FC-200D-2642-FE0F",
		"value": "💆🏼‍♂️",
		"descriptor": "Man Getting Massage: Medium-Light Skin Tone"
	},
	"1F486-1F3FD": {
		"key": "1F486-1F3FD",
		"value": "💆🏽",
		"descriptor": "Person Getting Massage: Medium Skin Tone"
	},
	"1F486-1F3FD-200D-2640-FE0F": {
		"key": "1F486-1F3FD-200D-2640-FE0F",
		"value": "💆🏽‍♀️",
		"descriptor": "Woman Getting Massage: Medium Skin Tone"
	},
	"1F486-1F3FD-200D-2642-FE0F": {
		"key": "1F486-1F3FD-200D-2642-FE0F",
		"value": "💆🏽‍♂️",
		"descriptor": "Man Getting Massage: Medium Skin Tone"
	},
	"1F486-1F3FE": {
		"key": "1F486-1F3FE",
		"value": "💆🏾",
		"descriptor": "Person Getting Massage: Medium-Dark Skin Tone"
	},
	"1F486-1F3FE-200D-2640-FE0F": {
		"key": "1F486-1F3FE-200D-2640-FE0F",
		"value": "💆🏾‍♀️",
		"descriptor": "Woman Getting Massage: Medium-Dark Skin Tone"
	},
	"1F486-1F3FE-200D-2642-FE0F": {
		"key": "1F486-1F3FE-200D-2642-FE0F",
		"value": "💆🏾‍♂️",
		"descriptor": "Man Getting Massage: Medium-Dark Skin Tone"
	},
	"1F486-1F3FF": {
		"key": "1F486-1F3FF",
		"value": "💆🏿",
		"descriptor": "Person Getting Massage: Dark Skin Tone"
	},
	"1F486-1F3FF-200D-2640-FE0F": {
		"key": "1F486-1F3FF-200D-2640-FE0F",
		"value": "💆🏿‍♀️",
		"descriptor": "Woman Getting Massage: Dark Skin Tone"
	},
	"1F486-1F3FF-200D-2642-FE0F": {
		"key": "1F486-1F3FF-200D-2642-FE0F",
		"value": "💆🏿‍♂️",
		"descriptor": "Man Getting Massage: Dark Skin Tone"
	},
	"1F486-200D-2640-FE0F": {
		"key": "1F486-200D-2640-FE0F",
		"value": "💆‍♀️",
		"descriptor": "Woman Getting Massage"
	},
	"1F486-200D-2642-FE0F": {
		"key": "1F486-200D-2642-FE0F",
		"value": "💆‍♂️",
		"descriptor": "Man Getting Massage"
	},
	"1F487": {
		"key": "1F487",
		"value": "💇",
		"descriptor": "Person Getting Haircut"
	},
	"1F487-1F3FB": {
		"key": "1F487-1F3FB",
		"value": "💇🏻",
		"descriptor": "Person Getting Haircut: Light Skin Tone"
	},
	"1F487-1F3FB-200D-2640-FE0F": {
		"key": "1F487-1F3FB-200D-2640-FE0F",
		"value": "💇🏻‍♀️",
		"descriptor": "Woman Getting Haircut: Light Skin Tone"
	},
	"1F487-1F3FB-200D-2642-FE0F": {
		"key": "1F487-1F3FB-200D-2642-FE0F",
		"value": "💇🏻‍♂️",
		"descriptor": "Man Getting Haircut: Light Skin Tone"
	},
	"1F487-1F3FC": {
		"key": "1F487-1F3FC",
		"value": "💇🏼",
		"descriptor": "Person Getting Haircut: Medium-Light Skin Tone"
	},
	"1F487-1F3FC-200D-2640-FE0F": {
		"key": "1F487-1F3FC-200D-2640-FE0F",
		"value": "💇🏼‍♀️",
		"descriptor": "Woman Getting Haircut: Medium-Light Skin Tone"
	},
	"1F487-1F3FC-200D-2642-FE0F": {
		"key": "1F487-1F3FC-200D-2642-FE0F",
		"value": "💇🏼‍♂️",
		"descriptor": "Man Getting Haircut: Medium-Light Skin Tone"
	},
	"1F487-1F3FD": {
		"key": "1F487-1F3FD",
		"value": "💇🏽",
		"descriptor": "Person Getting Haircut: Medium Skin Tone"
	},
	"1F487-1F3FD-200D-2640-FE0F": {
		"key": "1F487-1F3FD-200D-2640-FE0F",
		"value": "💇🏽‍♀️",
		"descriptor": "Woman Getting Haircut: Medium Skin Tone"
	},
	"1F487-1F3FD-200D-2642-FE0F": {
		"key": "1F487-1F3FD-200D-2642-FE0F",
		"value": "💇🏽‍♂️",
		"descriptor": "Man Getting Haircut: Medium Skin Tone"
	},
	"1F487-1F3FE": {
		"key": "1F487-1F3FE",
		"value": "💇🏾",
		"descriptor": "Person Getting Haircut: Medium-Dark Skin Tone"
	},
	"1F487-1F3FE-200D-2640-FE0F": {
		"key": "1F487-1F3FE-200D-2640-FE0F",
		"value": "💇🏾‍♀️",
		"descriptor": "Woman Getting Haircut: Medium-Dark Skin Tone"
	},
	"1F487-1F3FE-200D-2642-FE0F": {
		"key": "1F487-1F3FE-200D-2642-FE0F",
		"value": "💇🏾‍♂️",
		"descriptor": "Man Getting Haircut: Medium-Dark Skin Tone"
	},
	"1F487-1F3FF": {
		"key": "1F487-1F3FF",
		"value": "💇🏿",
		"descriptor": "Person Getting Haircut: Dark Skin Tone"
	},
	"1F487-1F3FF-200D-2640-FE0F": {
		"key": "1F487-1F3FF-200D-2640-FE0F",
		"value": "💇🏿‍♀️",
		"descriptor": "Woman Getting Haircut: Dark Skin Tone"
	},
	"1F487-1F3FF-200D-2642-FE0F": {
		"key": "1F487-1F3FF-200D-2642-FE0F",
		"value": "💇🏿‍♂️",
		"descriptor": "Man Getting Haircut: Dark Skin Tone"
	},
	"1F487-200D-2640-FE0F": {
		"key": "1F487-200D-2640-FE0F",
		"value": "💇‍♀️",
		"descriptor": "Woman Getting Haircut"
	},
	"1F487-200D-2642-FE0F": {
		"key": "1F487-200D-2642-FE0F",
		"value": "💇‍♂️",
		"descriptor": "Man Getting Haircut"
	},
	"1F488": {
		"key": "1F488",
		"value": "💈",
		"descriptor": "Barber Pole"
	},
	"1F489": {
		"key": "1F489",
		"value": "💉",
		"descriptor": "Syringe"
	},
	"1F48A": {
		"key": "1F48A",
		"value": "💊",
		"descriptor": "Pill"
	},
	"1F48B": {
		"key": "1F48B",
		"value": "💋",
		"descriptor": "Kiss Mark"
	},
	"1F48C": {
		"key": "1F48C",
		"value": "💌",
		"descriptor": "Love Letter"
	},
	"1F48D": {
		"key": "1F48D",
		"value": "💍",
		"descriptor": "Ring"
	},
	"1F48E": {
		"key": "1F48E",
		"value": "💎",
		"descriptor": "Gem Stone"
	},
	"1F48F": {
		"key": "1F48F",
		"value": "💏",
		"descriptor": "Kiss"
	},
	"1F490": {
		"key": "1F490",
		"value": "💐",
		"descriptor": "Bouquet"
	},
	"1F491": {
		"key": "1F491",
		"value": "💑",
		"descriptor": "Couple With Heart"
	},
	"1F492": {
		"key": "1F492",
		"value": "💒",
		"descriptor": "Wedding"
	},
	"1F493": {
		"key": "1F493",
		"value": "💓",
		"descriptor": "Beating Heart"
	},
	"1F494": {
		"key": "1F494",
		"value": "💔",
		"descriptor": "Broken Heart"
	},
	"1F495": {
		"key": "1F495",
		"value": "💕",
		"descriptor": "Two Hearts"
	},
	"1F496": {
		"key": "1F496",
		"value": "💖",
		"descriptor": "Sparkling Heart"
	},
	"1F497": {
		"key": "1F497",
		"value": "💗",
		"descriptor": "Growing Heart"
	},
	"1F498": {
		"key": "1F498",
		"value": "💘",
		"descriptor": "Heart With Arrow"
	},
	"1F499": {
		"key": "1F499",
		"value": "💙",
		"descriptor": "Blue Heart"
	},
	"1F49A": {
		"key": "1F49A",
		"value": "💚",
		"descriptor": "Green Heart"
	},
	"1F49B": {
		"key": "1F49B",
		"value": "💛",
		"descriptor": "Yellow Heart"
	},
	"1F49C": {
		"key": "1F49C",
		"value": "💜",
		"descriptor": "Purple Heart"
	},
	"1F49D": {
		"key": "1F49D",
		"value": "💝",
		"descriptor": "Heart With Ribbon"
	},
	"1F49E": {
		"key": "1F49E",
		"value": "💞",
		"descriptor": "Revolving Hearts"
	},
	"1F49F": {
		"key": "1F49F",
		"value": "💟",
		"descriptor": "Heart Decoration"
	},
	"1F4A0": {
		"key": "1F4A0",
		"value": "💠",
		"descriptor": "Diamond With a Dot"
	},
	"1F4A1": {
		"key": "1F4A1",
		"value": "💡",
		"descriptor": "Light Bulb"
	},
	"1F4A2": {
		"key": "1F4A2",
		"value": "💢",
		"descriptor": "Anger Symbol"
	},
	"1F4A3": {
		"key": "1F4A3",
		"value": "💣",
		"descriptor": "Bomb"
	},
	"1F4A4": {
		"key": "1F4A4",
		"value": "💤",
		"descriptor": "Zzz"
	},
	"1F4A5": {
		"key": "1F4A5",
		"value": "💥",
		"descriptor": "Collision"
	},
	"1F4A6": {
		"key": "1F4A6",
		"value": "💦",
		"descriptor": "Sweat Droplets"
	},
	"1F4A7": {
		"key": "1F4A7",
		"value": "💧",
		"descriptor": "Droplet"
	},
	"1F4A8": {
		"key": "1F4A8",
		"value": "💨",
		"descriptor": "Dashing Away"
	},
	"1F4A9": {
		"key": "1F4A9",
		"value": "💩",
		"descriptor": "Pile of Poo"
	},
	"1F4AA": {
		"key": "1F4AA",
		"value": "💪",
		"descriptor": "Flexed Biceps"
	},
	"1F4AA-1F3FB": {
		"key": "1F4AA-1F3FB",
		"value": "💪🏻",
		"descriptor": "Flexed Biceps: Light Skin Tone"
	},
	"1F4AA-1F3FC": {
		"key": "1F4AA-1F3FC",
		"value": "💪🏼",
		"descriptor": "Flexed Biceps: Medium-Light Skin Tone"
	},
	"1F4AA-1F3FD": {
		"key": "1F4AA-1F3FD",
		"value": "💪🏽",
		"descriptor": "Flexed Biceps: Medium Skin Tone"
	},
	"1F4AA-1F3FE": {
		"key": "1F4AA-1F3FE",
		"value": "💪🏾",
		"descriptor": "Flexed Biceps: Medium-Dark Skin Tone"
	},
	"1F4AA-1F3FF": {
		"key": "1F4AA-1F3FF",
		"value": "💪🏿",
		"descriptor": "Flexed Biceps: Dark Skin Tone"
	},
	"1F4AB": {
		"key": "1F4AB",
		"value": "💫",
		"descriptor": "Dizzy"
	},
	"1F4AC": {
		"key": "1F4AC",
		"value": "💬",
		"descriptor": "Speech Balloon"
	},
	"1F4AD": {
		"key": "1F4AD",
		"value": "💭",
		"descriptor": "Thought Balloon"
	},
	"1F4AE": {
		"key": "1F4AE",
		"value": "💮",
		"descriptor": "White Flower"
	},
	"1F4AF": {
		"key": "1F4AF",
		"value": "💯",
		"descriptor": "Hundred Points"
	},
	"1F4B0": {
		"key": "1F4B0",
		"value": "💰",
		"descriptor": "Money Bag"
	},
	"1F4B1": {
		"key": "1F4B1",
		"value": "💱",
		"descriptor": "Currency Exchange"
	},
	"1F4B2": {
		"key": "1F4B2",
		"value": "💲",
		"descriptor": "Heavy Dollar Sign"
	},
	"1F4B3": {
		"key": "1F4B3",
		"value": "💳",
		"descriptor": "Credit Card"
	},
	"1F4B4": {
		"key": "1F4B4",
		"value": "💴",
		"descriptor": "Yen Banknote"
	},
	"1F4B5": {
		"key": "1F4B5",
		"value": "💵",
		"descriptor": "Dollar Banknote"
	},
	"1F4B6": {
		"key": "1F4B6",
		"value": "💶",
		"descriptor": "Euro Banknote"
	},
	"1F4B7": {
		"key": "1F4B7",
		"value": "💷",
		"descriptor": "Pound Banknote"
	},
	"1F4B8": {
		"key": "1F4B8",
		"value": "💸",
		"descriptor": "Money With Wings"
	},
	"1F4B9": {
		"key": "1F4B9",
		"value": "💹",
		"descriptor": "Chart Increasing With Yen"
	},
	"1F4BA": {
		"key": "1F4BA",
		"value": "💺",
		"descriptor": "Seat"
	},
	"1F4BB": {
		"key": "1F4BB",
		"value": "💻",
		"descriptor": "Laptop"
	},
	"1F4BC": {
		"key": "1F4BC",
		"value": "💼",
		"descriptor": "Briefcase"
	},
	"1F4BD": {
		"key": "1F4BD",
		"value": "💽",
		"descriptor": "Computer Disk"
	},
	"1F4BE": {
		"key": "1F4BE",
		"value": "💾",
		"descriptor": "Floppy Disk"
	},
	"1F4BF": {
		"key": "1F4BF",
		"value": "💿",
		"descriptor": "Optical Disk"
	},
	"1F4C0": {
		"key": "1F4C0",
		"value": "📀",
		"descriptor": "DVD"
	},
	"1F4C1": {
		"key": "1F4C1",
		"value": "📁",
		"descriptor": "File Folder"
	},
	"1F4C2": {
		"key": "1F4C2",
		"value": "📂",
		"descriptor": "Open File Folder"
	},
	"1F4C3": {
		"key": "1F4C3",
		"value": "📃",
		"descriptor": "Page With Curl"
	},
	"1F4C4": {
		"key": "1F4C4",
		"value": "📄",
		"descriptor": "Page Facing Up"
	},
	"1F4C5": {
		"key": "1F4C5",
		"value": "📅",
		"descriptor": "Calendar"
	},
	"1F4C6": {
		"key": "1F4C6",
		"value": "📆",
		"descriptor": "Tear-Off Calendar"
	},
	"1F4C7": {
		"key": "1F4C7",
		"value": "📇",
		"descriptor": "Card Index"
	},
	"1F4C8": {
		"key": "1F4C8",
		"value": "📈",
		"descriptor": "Chart Increasing"
	},
	"1F4C9": {
		"key": "1F4C9",
		"value": "📉",
		"descriptor": "Chart Decreasing"
	},
	"1F4CA": {
		"key": "1F4CA",
		"value": "📊",
		"descriptor": "Bar Chart"
	},
	"1F4CB": {
		"key": "1F4CB",
		"value": "📋",
		"descriptor": "Clipboard"
	},
	"1F4CC": {
		"key": "1F4CC",
		"value": "📌",
		"descriptor": "Pushpin"
	},
	"1F4CD": {
		"key": "1F4CD",
		"value": "📍",
		"descriptor": "Round Pushpin"
	},
	"1F4CE": {
		"key": "1F4CE",
		"value": "📎",
		"descriptor": "Paperclip"
	},
	"1F4CF": {
		"key": "1F4CF",
		"value": "📏",
		"descriptor": "Straight Ruler"
	},
	"1F4D0": {
		"key": "1F4D0",
		"value": "📐",
		"descriptor": "Triangular Ruler"
	},
	"1F4D1": {
		"key": "1F4D1",
		"value": "📑",
		"descriptor": "Bookmark Tabs"
	},
	"1F4D2": {
		"key": "1F4D2",
		"value": "📒",
		"descriptor": "Ledger"
	},
	"1F4D3": {
		"key": "1F4D3",
		"value": "📓",
		"descriptor": "Notebook"
	},
	"1F4D4": {
		"key": "1F4D4",
		"value": "📔",
		"descriptor": "Notebook With Decorative Cover"
	},
	"1F4D5": {
		"key": "1F4D5",
		"value": "📕",
		"descriptor": "Closed Book"
	},
	"1F4D6": {
		"key": "1F4D6",
		"value": "📖",
		"descriptor": "Open Book"
	},
	"1F4D7": {
		"key": "1F4D7",
		"value": "📗",
		"descriptor": "Green Book"
	},
	"1F4D8": {
		"key": "1F4D8",
		"value": "📘",
		"descriptor": "Blue Book"
	},
	"1F4D9": {
		"key": "1F4D9",
		"value": "📙",
		"descriptor": "Orange Book"
	},
	"1F4DA": {
		"key": "1F4DA",
		"value": "📚",
		"descriptor": "Books"
	},
	"1F4DB": {
		"key": "1F4DB",
		"value": "📛",
		"descriptor": "Name Badge"
	},
	"1F4DC": {
		"key": "1F4DC",
		"value": "📜",
		"descriptor": "Scroll"
	},
	"1F4DD": {
		"key": "1F4DD",
		"value": "📝",
		"descriptor": "Memo"
	},
	"1F4DE": {
		"key": "1F4DE",
		"value": "📞",
		"descriptor": "Telephone Receiver"
	},
	"1F4DF": {
		"key": "1F4DF",
		"value": "📟",
		"descriptor": "Pager"
	},
	"1F4E0": {
		"key": "1F4E0",
		"value": "📠",
		"descriptor": "Fax Machine"
	},
	"1F4E1": {
		"key": "1F4E1",
		"value": "📡",
		"descriptor": "Satellite Antenna"
	},
	"1F4E2": {
		"key": "1F4E2",
		"value": "📢",
		"descriptor": "Loudspeaker"
	},
	"1F4E3": {
		"key": "1F4E3",
		"value": "📣",
		"descriptor": "Megaphone"
	},
	"1F4E4": {
		"key": "1F4E4",
		"value": "📤",
		"descriptor": "Outbox Tray"
	},
	"1F4E5": {
		"key": "1F4E5",
		"value": "📥",
		"descriptor": "Inbox Tray"
	},
	"1F4E6": {
		"key": "1F4E6",
		"value": "📦",
		"descriptor": "Package"
	},
	"1F4E7": {
		"key": "1F4E7",
		"value": "📧",
		"descriptor": "E-Mail"
	},
	"1F4E8": {
		"key": "1F4E8",
		"value": "📨",
		"descriptor": "Incoming Envelope"
	},
	"1F4E9": {
		"key": "1F4E9",
		"value": "📩",
		"descriptor": "Envelope With Arrow"
	},
	"1F4EA": {
		"key": "1F4EA",
		"value": "📪",
		"descriptor": "Closed Mailbox With Lowered Flag"
	},
	"1F4EB": {
		"key": "1F4EB",
		"value": "📫",
		"descriptor": "Closed Mailbox With Raised Flag"
	},
	"1F4EC": {
		"key": "1F4EC",
		"value": "📬",
		"descriptor": "Open Mailbox With Raised Flag"
	},
	"1F4ED": {
		"key": "1F4ED",
		"value": "📭",
		"descriptor": "Open Mailbox With Lowered Flag"
	},
	"1F4EE": {
		"key": "1F4EE",
		"value": "📮",
		"descriptor": "Postbox"
	},
	"1F4EF": {
		"key": "1F4EF",
		"value": "📯",
		"descriptor": "Postal Horn"
	},
	"1F4F0": {
		"key": "1F4F0",
		"value": "📰",
		"descriptor": "Newspaper"
	},
	"1F4F1": {
		"key": "1F4F1",
		"value": "📱",
		"descriptor": "Mobile Phone"
	},
	"1F4F2": {
		"key": "1F4F2",
		"value": "📲",
		"descriptor": "Mobile Phone With Arrow"
	},
	"1F4F3": {
		"key": "1F4F3",
		"value": "📳",
		"descriptor": "Vibration Mode"
	},
	"1F4F4": {
		"key": "1F4F4",
		"value": "📴",
		"descriptor": "Mobile Phone Off"
	},
	"1F4F5": {
		"key": "1F4F5",
		"value": "📵",
		"descriptor": "No Mobile Phones"
	},
	"1F4F6": {
		"key": "1F4F6",
		"value": "📶",
		"descriptor": "Antenna Bars"
	},
	"1F4F7": {
		"key": "1F4F7",
		"value": "📷",
		"descriptor": "Camera"
	},
	"1F4F8": {
		"key": "1F4F8",
		"value": "📸",
		"descriptor": "Camera With Flash"
	},
	"1F4F9": {
		"key": "1F4F9",
		"value": "📹",
		"descriptor": "Video Camera"
	},
	"1F4FA": {
		"key": "1F4FA",
		"value": "📺",
		"descriptor": "Television"
	},
	"1F4FB": {
		"key": "1F4FB",
		"value": "📻",
		"descriptor": "Radio"
	},
	"1F4FC": {
		"key": "1F4FC",
		"value": "📼",
		"descriptor": "Videocassette"
	},
	"1F4FD-FE0F": {
		"key": "1F4FD-FE0F",
		"value": "📽️",
		"descriptor": "Film Projector"
	},
	"1F4FF": {
		"key": "1F4FF",
		"value": "📿",
		"descriptor": "Prayer Beads"
	},
	"1F500": {
		"key": "1F500",
		"value": "🔀",
		"descriptor": "Shuffle Tracks Button"
	},
	"1F501": {
		"key": "1F501",
		"value": "🔁",
		"descriptor": "Repeat Button"
	},
	"1F502": {
		"key": "1F502",
		"value": "🔂",
		"descriptor": "Repeat Single Button"
	},
	"1F503": {
		"key": "1F503",
		"value": "🔃",
		"descriptor": "Clockwise Vertical Arrows"
	},
	"1F504": {
		"key": "1F504",
		"value": "🔄",
		"descriptor": "Counterclockwise Arrows Button"
	},
	"1F505": {
		"key": "1F505",
		"value": "🔅",
		"descriptor": "Dim Button"
	},
	"1F506": {
		"key": "1F506",
		"value": "🔆",
		"descriptor": "Bright Button"
	},
	"1F507": {
		"key": "1F507",
		"value": "🔇",
		"descriptor": "Muted Speaker"
	},
	"1F508": {
		"key": "1F508",
		"value": "🔈",
		"descriptor": "Speaker Low Volume"
	},
	"1F509": {
		"key": "1F509",
		"value": "🔉",
		"descriptor": "Speaker Medium Volume"
	},
	"1F50A": {
		"key": "1F50A",
		"value": "🔊",
		"descriptor": "Speaker High Volume"
	},
	"1F50B": {
		"key": "1F50B",
		"value": "🔋",
		"descriptor": "Battery"
	},
	"1F50C": {
		"key": "1F50C",
		"value": "🔌",
		"descriptor": "Electric Plug"
	},
	"1F50D": {
		"key": "1F50D",
		"value": "🔍",
		"descriptor": "Magnifying Glass Tilted Left"
	},
	"1F50E": {
		"key": "1F50E",
		"value": "🔎",
		"descriptor": "Magnifying Glass Tilted Right"
	},
	"1F50F": {
		"key": "1F50F",
		"value": "🔏",
		"descriptor": "Locked With Pen"
	},
	"1F510": {
		"key": "1F510",
		"value": "🔐",
		"descriptor": "Locked With Key"
	},
	"1F511": {
		"key": "1F511",
		"value": "🔑",
		"descriptor": "Key"
	},
	"1F512": {
		"key": "1F512",
		"value": "🔒",
		"descriptor": "Locked"
	},
	"1F513": {
		"key": "1F513",
		"value": "🔓",
		"descriptor": "Unlocked"
	},
	"1F514": {
		"key": "1F514",
		"value": "🔔",
		"descriptor": "Bell"
	},
	"1F515": {
		"key": "1F515",
		"value": "🔕",
		"descriptor": "Bell With Slash"
	},
	"1F516": {
		"key": "1F516",
		"value": "🔖",
		"descriptor": "Bookmark"
	},
	"1F517": {
		"key": "1F517",
		"value": "🔗",
		"descriptor": "Link"
	},
	"1F518": {
		"key": "1F518",
		"value": "🔘",
		"descriptor": "Radio Button"
	},
	"1F519": {
		"key": "1F519",
		"value": "🔙",
		"descriptor": "Back Arrow"
	},
	"1F51A": {
		"key": "1F51A",
		"value": "🔚",
		"descriptor": "End Arrow"
	},
	"1F51B": {
		"key": "1F51B",
		"value": "🔛",
		"descriptor": "On! Arrow"
	},
	"1F51C": {
		"key": "1F51C",
		"value": "🔜",
		"descriptor": "Soon Arrow"
	},
	"1F51D": {
		"key": "1F51D",
		"value": "🔝",
		"descriptor": "Top Arrow"
	},
	"1F51E": {
		"key": "1F51E",
		"value": "🔞",
		"descriptor": "No One Under Eighteen"
	},
	"1F51F": {
		"key": "1F51F",
		"value": "🔟",
		"descriptor": "Keycap: 10"
	},
	"1F520": {
		"key": "1F520",
		"value": "🔠",
		"descriptor": "Input Latin Uppercase"
	},
	"1F521": {
		"key": "1F521",
		"value": "🔡",
		"descriptor": "Input Latin Lowercase"
	},
	"1F522": {
		"key": "1F522",
		"value": "🔢",
		"descriptor": "Input Numbers"
	},
	"1F523": {
		"key": "1F523",
		"value": "🔣",
		"descriptor": "Input Symbols"
	},
	"1F524": {
		"key": "1F524",
		"value": "🔤",
		"descriptor": "Input Latin Letters"
	},
	"1F525": {
		"key": "1F525",
		"value": "🔥",
		"descriptor": "Fire"
	},
	"1F526": {
		"key": "1F526",
		"value": "🔦",
		"descriptor": "Flashlight"
	},
	"1F527": {
		"key": "1F527",
		"value": "🔧",
		"descriptor": "Wrench"
	},
	"1F528": {
		"key": "1F528",
		"value": "🔨",
		"descriptor": "Hammer"
	},
	"1F529": {
		"key": "1F529",
		"value": "🔩",
		"descriptor": "Nut and Bolt"
	},
	"1F52A": {
		"key": "1F52A",
		"value": "🔪",
		"descriptor": "Kitchen Knife"
	},
	"1F52B": {
		"key": "1F52B",
		"value": "🔫",
		"descriptor": "Pistol"
	},
	"1F52C": {
		"key": "1F52C",
		"value": "🔬",
		"descriptor": "Microscope"
	},
	"1F52D": {
		"key": "1F52D",
		"value": "🔭",
		"descriptor": "Telescope"
	},
	"1F52E": {
		"key": "1F52E",
		"value": "🔮",
		"descriptor": "Crystal Ball"
	},
	"1F52F": {
		"key": "1F52F",
		"value": "🔯",
		"descriptor": "Dotted Six-Pointed Star"
	},
	"1F530": {
		"key": "1F530",
		"value": "🔰",
		"descriptor": "Japanese Symbol for Beginner"
	},
	"1F531": {
		"key": "1F531",
		"value": "🔱",
		"descriptor": "Trident Emblem"
	},
	"1F532": {
		"key": "1F532",
		"value": "🔲",
		"descriptor": "Black Square Button"
	},
	"1F533": {
		"key": "1F533",
		"value": "🔳",
		"descriptor": "White Square Button"
	},
	"1F534": {
		"key": "1F534",
		"value": "🔴",
		"descriptor": "Red Circle"
	},
	"1F535": {
		"key": "1F535",
		"value": "🔵",
		"descriptor": "Blue Circle"
	},
	"1F536": {
		"key": "1F536",
		"value": "🔶",
		"descriptor": "Large Orange Diamond"
	},
	"1F537": {
		"key": "1F537",
		"value": "🔷",
		"descriptor": "Large Blue Diamond"
	},
	"1F538": {
		"key": "1F538",
		"value": "🔸",
		"descriptor": "Small Orange Diamond"
	},
	"1F539": {
		"key": "1F539",
		"value": "🔹",
		"descriptor": "Small Blue Diamond"
	},
	"1F53A": {
		"key": "1F53A",
		"value": "🔺",
		"descriptor": "Red Triangle Pointed Up"
	},
	"1F53B": {
		"key": "1F53B",
		"value": "🔻",
		"descriptor": "Red Triangle Pointed Down"
	},
	"1F53C": {
		"key": "1F53C",
		"value": "🔼",
		"descriptor": "Upwards Button"
	},
	"1F53D": {
		"key": "1F53D",
		"value": "🔽",
		"descriptor": "Downwards Button"
	},
	"1F549-FE0F": {
		"key": "1F549-FE0F",
		"value": "🕉️",
		"descriptor": "Om"
	},
	"1F54A-FE0F": {
		"key": "1F54A-FE0F",
		"value": "🕊️",
		"descriptor": "Dove"
	},
	"1F54B": {
		"key": "1F54B",
		"value": "🕋",
		"descriptor": "Kaaba"
	},
	"1F54C": {
		"key": "1F54C",
		"value": "🕌",
		"descriptor": "Mosque"
	},
	"1F54D": {
		"key": "1F54D",
		"value": "🕍",
		"descriptor": "Synagogue"
	},
	"1F54E": {
		"key": "1F54E",
		"value": "🕎",
		"descriptor": "Menorah"
	},
	"1F550": {
		"key": "1F550",
		"value": "🕐",
		"descriptor": "One O’Clock"
	},
	"1F551": {
		"key": "1F551",
		"value": "🕑",
		"descriptor": "Two O’Clock"
	},
	"1F552": {
		"key": "1F552",
		"value": "🕒",
		"descriptor": "Three O’Clock"
	},
	"1F553": {
		"key": "1F553",
		"value": "🕓",
		"descriptor": "Four O’Clock"
	},
	"1F554": {
		"key": "1F554",
		"value": "🕔",
		"descriptor": "Five O’Clock"
	},
	"1F555": {
		"key": "1F555",
		"value": "🕕",
		"descriptor": "Six O’Clock"
	},
	"1F556": {
		"key": "1F556",
		"value": "🕖",
		"descriptor": "Seven O’Clock"
	},
	"1F557": {
		"key": "1F557",
		"value": "🕗",
		"descriptor": "Eight O’Clock"
	},
	"1F558": {
		"key": "1F558",
		"value": "🕘",
		"descriptor": "Nine O’Clock"
	},
	"1F559": {
		"key": "1F559",
		"value": "🕙",
		"descriptor": "Ten O’Clock"
	},
	"1F55A": {
		"key": "1F55A",
		"value": "🕚",
		"descriptor": "Eleven O’Clock"
	},
	"1F55B": {
		"key": "1F55B",
		"value": "🕛",
		"descriptor": "Twelve O’Clock"
	},
	"1F55C": {
		"key": "1F55C",
		"value": "🕜",
		"descriptor": "One-Thirty"
	},
	"1F55D": {
		"key": "1F55D",
		"value": "🕝",
		"descriptor": "Two-Thirty"
	},
	"1F55E": {
		"key": "1F55E",
		"value": "🕞",
		"descriptor": "Three-Thirty"
	},
	"1F55F": {
		"key": "1F55F",
		"value": "🕟",
		"descriptor": "Four-Thirty"
	},
	"1F560": {
		"key": "1F560",
		"value": "🕠",
		"descriptor": "Five-Thirty"
	},
	"1F561": {
		"key": "1F561",
		"value": "🕡",
		"descriptor": "Six-Thirty"
	},
	"1F562": {
		"key": "1F562",
		"value": "🕢",
		"descriptor": "Seven-Thirty"
	},
	"1F563": {
		"key": "1F563",
		"value": "🕣",
		"descriptor": "Eight-Thirty"
	},
	"1F564": {
		"key": "1F564",
		"value": "🕤",
		"descriptor": "Nine-Thirty"
	},
	"1F565": {
		"key": "1F565",
		"value": "🕥",
		"descriptor": "Ten-Thirty"
	},
	"1F566": {
		"key": "1F566",
		"value": "🕦",
		"descriptor": "Eleven-Thirty"
	},
	"1F567": {
		"key": "1F567",
		"value": "🕧",
		"descriptor": "Twelve-Thirty"
	},
	"1F56F-FE0F": {
		"key": "1F56F-FE0F",
		"value": "🕯️",
		"descriptor": "Candle"
	},
	"1F570-FE0F": {
		"key": "1F570-FE0F",
		"value": "🕰️",
		"descriptor": "Mantelpiece Clock"
	},
	"1F573-FE0F": {
		"key": "1F573-FE0F",
		"value": "🕳️",
		"descriptor": "Hole"
	},
	"1F574-1F3FB": {
		"key": "1F574-1F3FB",
		"value": "🕴🏻",
		"descriptor": "Man in Suit Levitating: Light Skin Tone"
	},
	"1F574-1F3FC": {
		"key": "1F574-1F3FC",
		"value": "🕴🏼",
		"descriptor": "Man in Suit Levitating: Medium-Light Skin Tone"
	},
	"1F574-1F3FD": {
		"key": "1F574-1F3FD",
		"value": "🕴🏽",
		"descriptor": "Man in Suit Levitating: Medium Skin Tone"
	},
	"1F574-1F3FE": {
		"key": "1F574-1F3FE",
		"value": "🕴🏾",
		"descriptor": "Man in Suit Levitating: Medium-Dark Skin Tone"
	},
	"1F574-1F3FF": {
		"key": "1F574-1F3FF",
		"value": "🕴🏿",
		"descriptor": "Man in Suit Levitating: Dark Skin Tone"
	},
	"1F574-FE0F": {
		"key": "1F574-FE0F",
		"value": "🕴️",
		"descriptor": "Man in Suit Levitating"
	},
	"1F575-1F3FB": {
		"key": "1F575-1F3FB",
		"value": "🕵🏻",
		"descriptor": "Detective: Light Skin Tone"
	},
	"1F575-1F3FB-200D-2640-FE0F": {
		"key": "1F575-1F3FB-200D-2640-FE0F",
		"value": "🕵🏻‍♀️",
		"descriptor": "Woman Detective: Light Skin Tone"
	},
	"1F575-1F3FB-200D-2642-FE0F": {
		"key": "1F575-1F3FB-200D-2642-FE0F",
		"value": "🕵🏻‍♂️",
		"descriptor": "Man Detective: Light Skin Tone"
	},
	"1F575-1F3FC": {
		"key": "1F575-1F3FC",
		"value": "🕵🏼",
		"descriptor": "Detective: Medium-Light Skin Tone"
	},
	"1F575-1F3FC-200D-2640-FE0F": {
		"key": "1F575-1F3FC-200D-2640-FE0F",
		"value": "🕵🏼‍♀️",
		"descriptor": "Woman Detective: Medium-Light Skin Tone"
	},
	"1F575-1F3FC-200D-2642-FE0F": {
		"key": "1F575-1F3FC-200D-2642-FE0F",
		"value": "🕵🏼‍♂️",
		"descriptor": "Man Detective: Medium-Light Skin Tone"
	},
	"1F575-1F3FD": {
		"key": "1F575-1F3FD",
		"value": "🕵🏽",
		"descriptor": "Detective: Medium Skin Tone"
	},
	"1F575-1F3FD-200D-2640-FE0F": {
		"key": "1F575-1F3FD-200D-2640-FE0F",
		"value": "🕵🏽‍♀️",
		"descriptor": "Woman Detective: Medium Skin Tone"
	},
	"1F575-1F3FD-200D-2642-FE0F": {
		"key": "1F575-1F3FD-200D-2642-FE0F",
		"value": "🕵🏽‍♂️",
		"descriptor": "Man Detective: Medium Skin Tone"
	},
	"1F575-1F3FE": {
		"key": "1F575-1F3FE",
		"value": "🕵🏾",
		"descriptor": "Detective: Medium-Dark Skin Tone"
	},
	"1F575-1F3FE-200D-2640-FE0F": {
		"key": "1F575-1F3FE-200D-2640-FE0F",
		"value": "🕵🏾‍♀️",
		"descriptor": "Woman Detective: Medium-Dark Skin Tone"
	},
	"1F575-1F3FE-200D-2642-FE0F": {
		"key": "1F575-1F3FE-200D-2642-FE0F",
		"value": "🕵🏾‍♂️",
		"descriptor": "Man Detective: Medium-Dark Skin Tone"
	},
	"1F575-1F3FF": {
		"key": "1F575-1F3FF",
		"value": "🕵🏿",
		"descriptor": "Detective: Dark Skin Tone"
	},
	"1F575-1F3FF-200D-2640-FE0F": {
		"key": "1F575-1F3FF-200D-2640-FE0F",
		"value": "🕵🏿‍♀️",
		"descriptor": "Woman Detective: Dark Skin Tone"
	},
	"1F575-1F3FF-200D-2642-FE0F": {
		"key": "1F575-1F3FF-200D-2642-FE0F",
		"value": "🕵🏿‍♂️",
		"descriptor": "Man Detective: Dark Skin Tone"
	},
	"1F575-FE0F": {
		"key": "1F575-FE0F",
		"value": "🕵️",
		"descriptor": "Detective"
	},
	"1F575-FE0F-200D-2640-FE0F": {
		"key": "1F575-FE0F-200D-2640-FE0F",
		"value": "🕵️‍♀️",
		"descriptor": "Woman Detective"
	},
	"1F575-FE0F-200D-2642-FE0F": {
		"key": "1F575-FE0F-200D-2642-FE0F",
		"value": "🕵️‍♂️",
		"descriptor": "Man Detective"
	},
	"1F576-FE0F": {
		"key": "1F576-FE0F",
		"value": "🕶️",
		"descriptor": "Sunglasses"
	},
	"1F577-FE0F": {
		"key": "1F577-FE0F",
		"value": "🕷️",
		"descriptor": "Spider"
	},
	"1F578-FE0F": {
		"key": "1F578-FE0F",
		"value": "🕸️",
		"descriptor": "Spider Web"
	},
	"1F579-FE0F": {
		"key": "1F579-FE0F",
		"value": "🕹️",
		"descriptor": "Joystick"
	},
	"1F57A": {
		"key": "1F57A",
		"value": "🕺",
		"descriptor": "Man Dancing"
	},
	"1F57A-1F3FB": {
		"key": "1F57A-1F3FB",
		"value": "🕺🏻",
		"descriptor": "Man Dancing: Light Skin Tone"
	},
	"1F57A-1F3FC": {
		"key": "1F57A-1F3FC",
		"value": "🕺🏼",
		"descriptor": "Man Dancing: Medium-Light Skin Tone"
	},
	"1F57A-1F3FD": {
		"key": "1F57A-1F3FD",
		"value": "🕺🏽",
		"descriptor": "Man Dancing: Medium Skin Tone"
	},
	"1F57A-1F3FE": {
		"key": "1F57A-1F3FE",
		"value": "🕺🏾",
		"descriptor": "Man Dancing: Medium-Dark Skin Tone"
	},
	"1F57A-1F3FF": {
		"key": "1F57A-1F3FF",
		"value": "🕺🏿",
		"descriptor": "Man Dancing: Dark Skin Tone"
	},
	"1F587-FE0F": {
		"key": "1F587-FE0F",
		"value": "🖇️",
		"descriptor": "Linked Paperclips"
	},
	"1F58A-FE0F": {
		"key": "1F58A-FE0F",
		"value": "🖊️",
		"descriptor": "Pen"
	},
	"1F58B-FE0F": {
		"key": "1F58B-FE0F",
		"value": "🖋️",
		"descriptor": "Fountain Pen"
	},
	"1F58C-FE0F": {
		"key": "1F58C-FE0F",
		"value": "🖌️",
		"descriptor": "Paintbrush"
	},
	"1F58D-FE0F": {
		"key": "1F58D-FE0F",
		"value": "🖍️",
		"descriptor": "Crayon"
	},
	"1F590-1F3FB": {
		"key": "1F590-1F3FB",
		"value": "🖐🏻",
		"descriptor": "Hand With Fingers Splayed: Light Skin Tone"
	},
	"1F590-1F3FC": {
		"key": "1F590-1F3FC",
		"value": "🖐🏼",
		"descriptor": "Hand With Fingers Splayed: Medium-Light Skin Tone"
	},
	"1F590-1F3FD": {
		"key": "1F590-1F3FD",
		"value": "🖐🏽",
		"descriptor": "Hand With Fingers Splayed: Medium Skin Tone"
	},
	"1F590-1F3FE": {
		"key": "1F590-1F3FE",
		"value": "🖐🏾",
		"descriptor": "Hand With Fingers Splayed: Medium-Dark Skin Tone"
	},
	"1F590-1F3FF": {
		"key": "1F590-1F3FF",
		"value": "🖐🏿",
		"descriptor": "Hand With Fingers Splayed: Dark Skin Tone"
	},
	"1F590-FE0F": {
		"key": "1F590-FE0F",
		"value": "🖐️",
		"descriptor": "Hand With Fingers Splayed"
	},
	"1F595": {
		"key": "1F595",
		"value": "🖕",
		"descriptor": "Middle Finger"
	},
	"1F595-1F3FB": {
		"key": "1F595-1F3FB",
		"value": "🖕🏻",
		"descriptor": "Middle Finger: Light Skin Tone"
	},
	"1F595-1F3FC": {
		"key": "1F595-1F3FC",
		"value": "🖕🏼",
		"descriptor": "Middle Finger: Medium-Light Skin Tone"
	},
	"1F595-1F3FD": {
		"key": "1F595-1F3FD",
		"value": "🖕🏽",
		"descriptor": "Middle Finger: Medium Skin Tone"
	},
	"1F595-1F3FE": {
		"key": "1F595-1F3FE",
		"value": "🖕🏾",
		"descriptor": "Middle Finger: Medium-Dark Skin Tone"
	},
	"1F595-1F3FF": {
		"key": "1F595-1F3FF",
		"value": "🖕🏿",
		"descriptor": "Middle Finger: Dark Skin Tone"
	},
	"1F596": {
		"key": "1F596",
		"value": "🖖",
		"descriptor": "Vulcan Salute"
	},
	"1F596-1F3FB": {
		"key": "1F596-1F3FB",
		"value": "🖖🏻",
		"descriptor": "Vulcan Salute: Light Skin Tone"
	},
	"1F596-1F3FC": {
		"key": "1F596-1F3FC",
		"value": "🖖🏼",
		"descriptor": "Vulcan Salute: Medium-Light Skin Tone"
	},
	"1F596-1F3FD": {
		"key": "1F596-1F3FD",
		"value": "🖖🏽",
		"descriptor": "Vulcan Salute: Medium Skin Tone"
	},
	"1F596-1F3FE": {
		"key": "1F596-1F3FE",
		"value": "🖖🏾",
		"descriptor": "Vulcan Salute: Medium-Dark Skin Tone"
	},
	"1F596-1F3FF": {
		"key": "1F596-1F3FF",
		"value": "🖖🏿",
		"descriptor": "Vulcan Salute: Dark Skin Tone"
	},
	"1F5A4": {
		"key": "1F5A4",
		"value": "🖤",
		"descriptor": "Black Heart"
	},
	"1F5A5-FE0F": {
		"key": "1F5A5-FE0F",
		"value": "🖥️",
		"descriptor": "Desktop Computer"
	},
	"1F5A8-FE0F": {
		"key": "1F5A8-FE0F",
		"value": "🖨️",
		"descriptor": "Printer"
	},
	"1F5B1-FE0F": {
		"key": "1F5B1-FE0F",
		"value": "🖱️",
		"descriptor": "Computer Mouse"
	},
	"1F5B2-FE0F": {
		"key": "1F5B2-FE0F",
		"value": "🖲️",
		"descriptor": "Trackball"
	},
	"1F5BC-FE0F": {
		"key": "1F5BC-FE0F",
		"value": "🖼️",
		"descriptor": "Framed Picture"
	},
	"1F5C2-FE0F": {
		"key": "1F5C2-FE0F",
		"value": "🗂️",
		"descriptor": "Card Index Dividers"
	},
	"1F5C3-FE0F": {
		"key": "1F5C3-FE0F",
		"value": "🗃️",
		"descriptor": "Card File Box"
	},
	"1F5C4-FE0F": {
		"key": "1F5C4-FE0F",
		"value": "🗄️",
		"descriptor": "File Cabinet"
	},
	"1F5D1-FE0F": {
		"key": "1F5D1-FE0F",
		"value": "🗑️",
		"descriptor": "Wastebasket"
	},
	"1F5D2-FE0F": {
		"key": "1F5D2-FE0F",
		"value": "🗒️",
		"descriptor": "Spiral Notepad"
	},
	"1F5D3-FE0F": {
		"key": "1F5D3-FE0F",
		"value": "🗓️",
		"descriptor": "Spiral Calendar"
	},
	"1F5DC-FE0F": {
		"key": "1F5DC-FE0F",
		"value": "🗜️",
		"descriptor": "Clamp"
	},
	"1F5DD-FE0F": {
		"key": "1F5DD-FE0F",
		"value": "🗝️",
		"descriptor": "Old Key"
	},
	"1F5DE-FE0F": {
		"key": "1F5DE-FE0F",
		"value": "🗞️",
		"descriptor": "Rolled-Up Newspaper"
	},
	"1F5E1-FE0F": {
		"key": "1F5E1-FE0F",
		"value": "🗡️",
		"descriptor": "Dagger"
	},
	"1F5E3-FE0F": {
		"key": "1F5E3-FE0F",
		"value": "🗣️",
		"descriptor": "Speaking Head"
	},
	"1F5E8-FE0F": {
		"key": "1F5E8-FE0F",
		"value": "🗨️",
		"descriptor": "Left Speech Bubble"
	},
	"1F5EF-FE0F": {
		"key": "1F5EF-FE0F",
		"value": "🗯️",
		"descriptor": "Right Anger Bubble"
	},
	"1F5F3-FE0F": {
		"key": "1F5F3-FE0F",
		"value": "🗳️",
		"descriptor": "Ballot Box With Ballot"
	},
	"1F5FA-FE0F": {
		"key": "1F5FA-FE0F",
		"value": "🗺️",
		"descriptor": "World Map"
	},
	"1F5FB": {
		"key": "1F5FB",
		"value": "🗻",
		"descriptor": "Mount Fuji"
	},
	"1F5FC": {
		"key": "1F5FC",
		"value": "🗼",
		"descriptor": "Tokyo Tower"
	},
	"1F5FD": {
		"key": "1F5FD",
		"value": "🗽",
		"descriptor": "Statue of Liberty"
	},
	"1F5FE": {
		"key": "1F5FE",
		"value": "🗾",
		"descriptor": "Map of Japan"
	},
	"1F5FF": {
		"key": "1F5FF",
		"value": "🗿",
		"descriptor": "Moai"
	},
	"1F600": {
		"key": "1F600",
		"value": "😀",
		"descriptor": "Grinning Face"
	},
	"1F601": {
		"key": "1F601",
		"value": "😁",
		"descriptor": "Beaming Face With Smiling Eyes"
	},
	"1F602": {
		"key": "1F602",
		"value": "😂",
		"descriptor": "Face With Tears of Joy"
	},
	"1F603": {
		"key": "1F603",
		"value": "😃",
		"descriptor": "Grinning Face With Big Eyes"
	},
	"1F604": {
		"key": "1F604",
		"value": "😄",
		"descriptor": "Grinning Face With Smiling Eyes"
	},
	"1F605": {
		"key": "1F605",
		"value": "😅",
		"descriptor": "Grinning Face With Sweat"
	},
	"1F606": {
		"key": "1F606",
		"value": "😆",
		"descriptor": "Grinning Squinting Face"
	},
	"1F607": {
		"key": "1F607",
		"value": "😇",
		"descriptor": "Smiling Face With Halo"
	},
	"1F608": {
		"key": "1F608",
		"value": "😈",
		"descriptor": "Smiling Face With Horns"
	},
	"1F609": {
		"key": "1F609",
		"value": "😉",
		"descriptor": "Winking Face"
	},
	"1F60A": {
		"key": "1F60A",
		"value": "😊",
		"descriptor": "Smiling Face With Smiling Eyes"
	},
	"1F60B": {
		"key": "1F60B",
		"value": "😋",
		"descriptor": "Face Savoring Food"
	},
	"1F60C": {
		"key": "1F60C",
		"value": "😌",
		"descriptor": "Relieved Face"
	},
	"1F60D": {
		"key": "1F60D",
		"value": "😍",
		"descriptor": "Smiling Face With Heart-Eyes"
	},
	"1F60E": {
		"key": "1F60E",
		"value": "😎",
		"descriptor": "Smiling Face With Sunglasses"
	},
	"1F60F": {
		"key": "1F60F",
		"value": "😏",
		"descriptor": "Smirking Face"
	},
	"1F610": {
		"key": "1F610",
		"value": "😐",
		"descriptor": "Neutral Face"
	},
	"1F611": {
		"key": "1F611",
		"value": "😑",
		"descriptor": "Expressionless Face"
	},
	"1F612": {
		"key": "1F612",
		"value": "😒",
		"descriptor": "Unamused Face"
	},
	"1F613": {
		"key": "1F613",
		"value": "😓",
		"descriptor": "Downcast Face With Sweat"
	},
	"1F614": {
		"key": "1F614",
		"value": "😔",
		"descriptor": "Pensive Face"
	},
	"1F615": {
		"key": "1F615",
		"value": "😕",
		"descriptor": "Confused Face"
	},
	"1F616": {
		"key": "1F616",
		"value": "😖",
		"descriptor": "Confounded Face"
	},
	"1F617": {
		"key": "1F617",
		"value": "😗",
		"descriptor": "Kissing Face"
	},
	"1F618": {
		"key": "1F618",
		"value": "😘",
		"descriptor": "Face Blowing a Kiss"
	},
	"1F619": {
		"key": "1F619",
		"value": "😙",
		"descriptor": "Kissing Face With Smiling Eyes"
	},
	"1F61A": {
		"key": "1F61A",
		"value": "😚",
		"descriptor": "Kissing Face With Closed Eyes"
	},
	"1F61B": {
		"key": "1F61B",
		"value": "😛",
		"descriptor": "Face With Tongue"
	},
	"1F61C": {
		"key": "1F61C",
		"value": "😜",
		"descriptor": "Winking Face With Tongue"
	},
	"1F61D": {
		"key": "1F61D",
		"value": "😝",
		"descriptor": "Squinting Face With Tongue"
	},
	"1F61E": {
		"key": "1F61E",
		"value": "😞",
		"descriptor": "Disappointed Face"
	},
	"1F61F": {
		"key": "1F61F",
		"value": "😟",
		"descriptor": "Worried Face"
	},
	"1F620": {
		"key": "1F620",
		"value": "😠",
		"descriptor": "Angry Face"
	},
	"1F621": {
		"key": "1F621",
		"value": "😡",
		"descriptor": "Pouting Face"
	},
	"1F622": {
		"key": "1F622",
		"value": "😢",
		"descriptor": "Crying Face"
	},
	"1F623": {
		"key": "1F623",
		"value": "😣",
		"descriptor": "Persevering Face"
	},
	"1F624": {
		"key": "1F624",
		"value": "😤",
		"descriptor": "Face With Steam From Nose"
	},
	"1F625": {
		"key": "1F625",
		"value": "😥",
		"descriptor": "Sad but Relieved Face"
	},
	"1F626": {
		"key": "1F626",
		"value": "😦",
		"descriptor": "Frowning Face With Open Mouth"
	},
	"1F627": {
		"key": "1F627",
		"value": "😧",
		"descriptor": "Anguished Face"
	},
	"1F628": {
		"key": "1F628",
		"value": "😨",
		"descriptor": "Fearful Face"
	},
	"1F629": {
		"key": "1F629",
		"value": "😩",
		"descriptor": "Weary Face"
	},
	"1F62A": {
		"key": "1F62A",
		"value": "😪",
		"descriptor": "Sleepy Face"
	},
	"1F62B": {
		"key": "1F62B",
		"value": "😫",
		"descriptor": "Tired Face"
	},
	"1F62C": {
		"key": "1F62C",
		"value": "😬",
		"descriptor": "Grimacing Face"
	},
	"1F62D": {
		"key": "1F62D",
		"value": "😭",
		"descriptor": "Loudly Crying Face"
	},
	"1F62E": {
		"key": "1F62E",
		"value": "😮",
		"descriptor": "Face With Open Mouth"
	},
	"1F62F": {
		"key": "1F62F",
		"value": "😯",
		"descriptor": "Hushed Face"
	},
	"1F630": {
		"key": "1F630",
		"value": "😰",
		"descriptor": "Anxious Face With Sweat"
	},
	"1F631": {
		"key": "1F631",
		"value": "😱",
		"descriptor": "Face Screaming in Fear"
	},
	"1F632": {
		"key": "1F632",
		"value": "😲",
		"descriptor": "Astonished Face"
	},
	"1F633": {
		"key": "1F633",
		"value": "😳",
		"descriptor": "Flushed Face"
	},
	"1F634": {
		"key": "1F634",
		"value": "😴",
		"descriptor": "Sleeping Face"
	},
	"1F635": {
		"key": "1F635",
		"value": "😵",
		"descriptor": "Dizzy Face"
	},
	"1F636": {
		"key": "1F636",
		"value": "😶",
		"descriptor": "Face Without Mouth"
	},
	"1F637": {
		"key": "1F637",
		"value": "😷",
		"descriptor": "Face With Medical Mask"
	},
	"1F638": {
		"key": "1F638",
		"value": "😸",
		"descriptor": "Grinning Cat With Smiling Eyes"
	},
	"1F639": {
		"key": "1F639",
		"value": "😹",
		"descriptor": "Cat With Tears of Joy"
	},
	"1F63A": {
		"key": "1F63A",
		"value": "😺",
		"descriptor": "Grinning Cat"
	},
	"1F63B": {
		"key": "1F63B",
		"value": "😻",
		"descriptor": "Smiling Cat With Heart-Eyes"
	},
	"1F63C": {
		"key": "1F63C",
		"value": "😼",
		"descriptor": "Cat With Wry Smile"
	},
	"1F63D": {
		"key": "1F63D",
		"value": "😽",
		"descriptor": "Kissing Cat"
	},
	"1F63E": {
		"key": "1F63E",
		"value": "😾",
		"descriptor": "Pouting Cat"
	},
	"1F63F": {
		"key": "1F63F",
		"value": "😿",
		"descriptor": "Crying Cat"
	},
	"1F640": {
		"key": "1F640",
		"value": "🙀",
		"descriptor": "Weary Cat"
	},
	"1F641": {
		"key": "1F641",
		"value": "🙁",
		"descriptor": "Slightly Frowning Face"
	},
	"1F642": {
		"key": "1F642",
		"value": "🙂",
		"descriptor": "Slightly Smiling Face"
	},
	"1F643": {
		"key": "1F643",
		"value": "🙃",
		"descriptor": "Upside-Down Face"
	},
	"1F644": {
		"key": "1F644",
		"value": "🙄",
		"descriptor": "Face With Rolling Eyes"
	},
	"1F645": {
		"key": "1F645",
		"value": "🙅",
		"descriptor": "Person Gesturing No"
	},
	"1F645-1F3FB": {
		"key": "1F645-1F3FB",
		"value": "🙅🏻",
		"descriptor": "Person Gesturing No: Light Skin Tone"
	},
	"1F645-1F3FB-200D-2640-FE0F": {
		"key": "1F645-1F3FB-200D-2640-FE0F",
		"value": "🙅🏻‍♀️",
		"descriptor": "Woman Gesturing No: Light Skin Tone"
	},
	"1F645-1F3FB-200D-2642-FE0F": {
		"key": "1F645-1F3FB-200D-2642-FE0F",
		"value": "🙅🏻‍♂️",
		"descriptor": "Man Gesturing No: Light Skin Tone"
	},
	"1F645-1F3FC": {
		"key": "1F645-1F3FC",
		"value": "🙅🏼",
		"descriptor": "Person Gesturing No: Medium-Light Skin Tone"
	},
	"1F645-1F3FC-200D-2640-FE0F": {
		"key": "1F645-1F3FC-200D-2640-FE0F",
		"value": "🙅🏼‍♀️",
		"descriptor": "Woman Gesturing No: Medium-Light Skin Tone"
	},
	"1F645-1F3FC-200D-2642-FE0F": {
		"key": "1F645-1F3FC-200D-2642-FE0F",
		"value": "🙅🏼‍♂️",
		"descriptor": "Man Gesturing No: Medium-Light Skin Tone"
	},
	"1F645-1F3FD": {
		"key": "1F645-1F3FD",
		"value": "🙅🏽",
		"descriptor": "Person Gesturing No: Medium Skin Tone"
	},
	"1F645-1F3FD-200D-2640-FE0F": {
		"key": "1F645-1F3FD-200D-2640-FE0F",
		"value": "🙅🏽‍♀️",
		"descriptor": "Woman Gesturing No: Medium Skin Tone"
	},
	"1F645-1F3FD-200D-2642-FE0F": {
		"key": "1F645-1F3FD-200D-2642-FE0F",
		"value": "🙅🏽‍♂️",
		"descriptor": "Man Gesturing No: Medium Skin Tone"
	},
	"1F645-1F3FE": {
		"key": "1F645-1F3FE",
		"value": "🙅🏾",
		"descriptor": "Person Gesturing No: Medium-Dark Skin Tone"
	},
	"1F645-1F3FE-200D-2640-FE0F": {
		"key": "1F645-1F3FE-200D-2640-FE0F",
		"value": "🙅🏾‍♀️",
		"descriptor": "Woman Gesturing No: Medium-Dark Skin Tone"
	},
	"1F645-1F3FE-200D-2642-FE0F": {
		"key": "1F645-1F3FE-200D-2642-FE0F",
		"value": "🙅🏾‍♂️",
		"descriptor": "Man Gesturing No: Medium-Dark Skin Tone"
	},
	"1F645-1F3FF": {
		"key": "1F645-1F3FF",
		"value": "🙅🏿",
		"descriptor": "Person Gesturing No: Dark Skin Tone"
	},
	"1F645-1F3FF-200D-2640-FE0F": {
		"key": "1F645-1F3FF-200D-2640-FE0F",
		"value": "🙅🏿‍♀️",
		"descriptor": "Woman Gesturing No: Dark Skin Tone"
	},
	"1F645-1F3FF-200D-2642-FE0F": {
		"key": "1F645-1F3FF-200D-2642-FE0F",
		"value": "🙅🏿‍♂️",
		"descriptor": "Man Gesturing No: Dark Skin Tone"
	},
	"1F645-200D-2640-FE0F": {
		"key": "1F645-200D-2640-FE0F",
		"value": "🙅‍♀️",
		"descriptor": "Woman Gesturing No"
	},
	"1F645-200D-2642-FE0F": {
		"key": "1F645-200D-2642-FE0F",
		"value": "🙅‍♂️",
		"descriptor": "Man Gesturing No"
	},
	"1F646": {
		"key": "1F646",
		"value": "🙆",
		"descriptor": "Person Gesturing OK"
	},
	"1F646-1F3FB": {
		"key": "1F646-1F3FB",
		"value": "🙆🏻",
		"descriptor": "Person Gesturing OK: Light Skin Tone"
	},
	"1F646-1F3FB-200D-2640-FE0F": {
		"key": "1F646-1F3FB-200D-2640-FE0F",
		"value": "🙆🏻‍♀️",
		"descriptor": "Woman Gesturing OK: Light Skin Tone"
	},
	"1F646-1F3FB-200D-2642-FE0F": {
		"key": "1F646-1F3FB-200D-2642-FE0F",
		"value": "🙆🏻‍♂️",
		"descriptor": "Man Gesturing OK: Light Skin Tone"
	},
	"1F646-1F3FC": {
		"key": "1F646-1F3FC",
		"value": "🙆🏼",
		"descriptor": "Person Gesturing OK: Medium-Light Skin Tone"
	},
	"1F646-1F3FC-200D-2640-FE0F": {
		"key": "1F646-1F3FC-200D-2640-FE0F",
		"value": "🙆🏼‍♀️",
		"descriptor": "Woman Gesturing OK: Medium-Light Skin Tone"
	},
	"1F646-1F3FC-200D-2642-FE0F": {
		"key": "1F646-1F3FC-200D-2642-FE0F",
		"value": "🙆🏼‍♂️",
		"descriptor": "Man Gesturing OK: Medium-Light Skin Tone"
	},
	"1F646-1F3FD": {
		"key": "1F646-1F3FD",
		"value": "🙆🏽",
		"descriptor": "Person Gesturing OK: Medium Skin Tone"
	},
	"1F646-1F3FD-200D-2640-FE0F": {
		"key": "1F646-1F3FD-200D-2640-FE0F",
		"value": "🙆🏽‍♀️",
		"descriptor": "Woman Gesturing OK: Medium Skin Tone"
	},
	"1F646-1F3FD-200D-2642-FE0F": {
		"key": "1F646-1F3FD-200D-2642-FE0F",
		"value": "🙆🏽‍♂️",
		"descriptor": "Man Gesturing OK: Medium Skin Tone"
	},
	"1F646-1F3FE": {
		"key": "1F646-1F3FE",
		"value": "🙆🏾",
		"descriptor": "Person Gesturing OK: Medium-Dark Skin Tone"
	},
	"1F646-1F3FE-200D-2640-FE0F": {
		"key": "1F646-1F3FE-200D-2640-FE0F",
		"value": "🙆🏾‍♀️",
		"descriptor": "Woman Gesturing OK: Medium-Dark Skin Tone"
	},
	"1F646-1F3FE-200D-2642-FE0F": {
		"key": "1F646-1F3FE-200D-2642-FE0F",
		"value": "🙆🏾‍♂️",
		"descriptor": "Man Gesturing OK: Medium-Dark Skin Tone"
	},
	"1F646-1F3FF": {
		"key": "1F646-1F3FF",
		"value": "🙆🏿",
		"descriptor": "Person Gesturing OK: Dark Skin Tone"
	},
	"1F646-1F3FF-200D-2640-FE0F": {
		"key": "1F646-1F3FF-200D-2640-FE0F",
		"value": "🙆🏿‍♀️",
		"descriptor": "Woman Gesturing OK: Dark Skin Tone"
	},
	"1F646-1F3FF-200D-2642-FE0F": {
		"key": "1F646-1F3FF-200D-2642-FE0F",
		"value": "🙆🏿‍♂️",
		"descriptor": "Man Gesturing OK: Dark Skin Tone"
	},
	"1F646-200D-2640-FE0F": {
		"key": "1F646-200D-2640-FE0F",
		"value": "🙆‍♀️",
		"descriptor": "Woman Gesturing OK"
	},
	"1F646-200D-2642-FE0F": {
		"key": "1F646-200D-2642-FE0F",
		"value": "🙆‍♂️",
		"descriptor": "Man Gesturing OK"
	},
	"1F647": {
		"key": "1F647",
		"value": "🙇",
		"descriptor": "Person Bowing"
	},
	"1F647-1F3FB": {
		"key": "1F647-1F3FB",
		"value": "🙇🏻",
		"descriptor": "Person Bowing: Light Skin Tone"
	},
	"1F647-1F3FB-200D-2640-FE0F": {
		"key": "1F647-1F3FB-200D-2640-FE0F",
		"value": "🙇🏻‍♀️",
		"descriptor": "Woman Bowing: Light Skin Tone"
	},
	"1F647-1F3FB-200D-2642-FE0F": {
		"key": "1F647-1F3FB-200D-2642-FE0F",
		"value": "🙇🏻‍♂️",
		"descriptor": "Man Bowing: Light Skin Tone"
	},
	"1F647-1F3FC": {
		"key": "1F647-1F3FC",
		"value": "🙇🏼",
		"descriptor": "Person Bowing: Medium-Light Skin Tone"
	},
	"1F647-1F3FC-200D-2640-FE0F": {
		"key": "1F647-1F3FC-200D-2640-FE0F",
		"value": "🙇🏼‍♀️",
		"descriptor": "Woman Bowing: Medium-Light Skin Tone"
	},
	"1F647-1F3FC-200D-2642-FE0F": {
		"key": "1F647-1F3FC-200D-2642-FE0F",
		"value": "🙇🏼‍♂️",
		"descriptor": "Man Bowing: Medium-Light Skin Tone"
	},
	"1F647-1F3FD": {
		"key": "1F647-1F3FD",
		"value": "🙇🏽",
		"descriptor": "Person Bowing: Medium Skin Tone"
	},
	"1F647-1F3FD-200D-2640-FE0F": {
		"key": "1F647-1F3FD-200D-2640-FE0F",
		"value": "🙇🏽‍♀️",
		"descriptor": "Woman Bowing: Medium Skin Tone"
	},
	"1F647-1F3FD-200D-2642-FE0F": {
		"key": "1F647-1F3FD-200D-2642-FE0F",
		"value": "🙇🏽‍♂️",
		"descriptor": "Man Bowing: Medium Skin Tone"
	},
	"1F647-1F3FE": {
		"key": "1F647-1F3FE",
		"value": "🙇🏾",
		"descriptor": "Person Bowing: Medium-Dark Skin Tone"
	},
	"1F647-1F3FE-200D-2640-FE0F": {
		"key": "1F647-1F3FE-200D-2640-FE0F",
		"value": "🙇🏾‍♀️",
		"descriptor": "Woman Bowing: Medium-Dark Skin Tone"
	},
	"1F647-1F3FE-200D-2642-FE0F": {
		"key": "1F647-1F3FE-200D-2642-FE0F",
		"value": "🙇🏾‍♂️",
		"descriptor": "Man Bowing: Medium-Dark Skin Tone"
	},
	"1F647-1F3FF": {
		"key": "1F647-1F3FF",
		"value": "🙇🏿",
		"descriptor": "Person Bowing: Dark Skin Tone"
	},
	"1F647-1F3FF-200D-2640-FE0F": {
		"key": "1F647-1F3FF-200D-2640-FE0F",
		"value": "🙇🏿‍♀️",
		"descriptor": "Woman Bowing: Dark Skin Tone"
	},
	"1F647-1F3FF-200D-2642-FE0F": {
		"key": "1F647-1F3FF-200D-2642-FE0F",
		"value": "🙇🏿‍♂️",
		"descriptor": "Man Bowing: Dark Skin Tone"
	},
	"1F647-200D-2640-FE0F": {
		"key": "1F647-200D-2640-FE0F",
		"value": "🙇‍♀️",
		"descriptor": "Woman Bowing"
	},
	"1F647-200D-2642-FE0F": {
		"key": "1F647-200D-2642-FE0F",
		"value": "🙇‍♂️",
		"descriptor": "Man Bowing"
	},
	"1F648": {
		"key": "1F648",
		"value": "🙈",
		"descriptor": "See-No-Evil Monkey"
	},
	"1F649": {
		"key": "1F649",
		"value": "🙉",
		"descriptor": "Hear-No-Evil Monkey"
	},
	"1F64A": {
		"key": "1F64A",
		"value": "🙊",
		"descriptor": "Speak-No-Evil Monkey"
	},
	"1F64B": {
		"key": "1F64B",
		"value": "🙋",
		"descriptor": "Person Raising Hand"
	},
	"1F64B-1F3FB": {
		"key": "1F64B-1F3FB",
		"value": "🙋🏻",
		"descriptor": "Person Raising Hand: Light Skin Tone"
	},
	"1F64B-1F3FB-200D-2640-FE0F": {
		"key": "1F64B-1F3FB-200D-2640-FE0F",
		"value": "🙋🏻‍♀️",
		"descriptor": "Woman Raising Hand: Light Skin Tone"
	},
	"1F64B-1F3FB-200D-2642-FE0F": {
		"key": "1F64B-1F3FB-200D-2642-FE0F",
		"value": "🙋🏻‍♂️",
		"descriptor": "Man Raising Hand: Light Skin Tone"
	},
	"1F64B-1F3FC": {
		"key": "1F64B-1F3FC",
		"value": "🙋🏼",
		"descriptor": "Person Raising Hand: Medium-Light Skin Tone"
	},
	"1F64B-1F3FC-200D-2640-FE0F": {
		"key": "1F64B-1F3FC-200D-2640-FE0F",
		"value": "🙋🏼‍♀️",
		"descriptor": "Woman Raising Hand: Medium-Light Skin Tone"
	},
	"1F64B-1F3FC-200D-2642-FE0F": {
		"key": "1F64B-1F3FC-200D-2642-FE0F",
		"value": "🙋🏼‍♂️",
		"descriptor": "Man Raising Hand: Medium-Light Skin Tone"
	},
	"1F64B-1F3FD": {
		"key": "1F64B-1F3FD",
		"value": "🙋🏽",
		"descriptor": "Person Raising Hand: Medium Skin Tone"
	},
	"1F64B-1F3FD-200D-2640-FE0F": {
		"key": "1F64B-1F3FD-200D-2640-FE0F",
		"value": "🙋🏽‍♀️",
		"descriptor": "Woman Raising Hand: Medium Skin Tone"
	},
	"1F64B-1F3FD-200D-2642-FE0F": {
		"key": "1F64B-1F3FD-200D-2642-FE0F",
		"value": "🙋🏽‍♂️",
		"descriptor": "Man Raising Hand: Medium Skin Tone"
	},
	"1F64B-1F3FE": {
		"key": "1F64B-1F3FE",
		"value": "🙋🏾",
		"descriptor": "Person Raising Hand: Medium-Dark Skin Tone"
	},
	"1F64B-1F3FE-200D-2640-FE0F": {
		"key": "1F64B-1F3FE-200D-2640-FE0F",
		"value": "🙋🏾‍♀️",
		"descriptor": "Woman Raising Hand: Medium-Dark Skin Tone"
	},
	"1F64B-1F3FE-200D-2642-FE0F": {
		"key": "1F64B-1F3FE-200D-2642-FE0F",
		"value": "🙋🏾‍♂️",
		"descriptor": "Man Raising Hand: Medium-Dark Skin Tone"
	},
	"1F64B-1F3FF": {
		"key": "1F64B-1F3FF",
		"value": "🙋🏿",
		"descriptor": "Person Raising Hand: Dark Skin Tone"
	},
	"1F64B-1F3FF-200D-2640-FE0F": {
		"key": "1F64B-1F3FF-200D-2640-FE0F",
		"value": "🙋🏿‍♀️",
		"descriptor": "Woman Raising Hand: Dark Skin Tone"
	},
	"1F64B-1F3FF-200D-2642-FE0F": {
		"key": "1F64B-1F3FF-200D-2642-FE0F",
		"value": "🙋🏿‍♂️",
		"descriptor": "Man Raising Hand: Dark Skin Tone"
	},
	"1F64B-200D-2640-FE0F": {
		"key": "1F64B-200D-2640-FE0F",
		"value": "🙋‍♀️",
		"descriptor": "Woman Raising Hand"
	},
	"1F64B-200D-2642-FE0F": {
		"key": "1F64B-200D-2642-FE0F",
		"value": "🙋‍♂️",
		"descriptor": "Man Raising Hand"
	},
	"1F64C": {
		"key": "1F64C",
		"value": "🙌",
		"descriptor": "Raising Hands"
	},
	"1F64C-1F3FB": {
		"key": "1F64C-1F3FB",
		"value": "🙌🏻",
		"descriptor": "Raising Hands: Light Skin Tone"
	},
	"1F64C-1F3FC": {
		"key": "1F64C-1F3FC",
		"value": "🙌🏼",
		"descriptor": "Raising Hands: Medium-Light Skin Tone"
	},
	"1F64C-1F3FD": {
		"key": "1F64C-1F3FD",
		"value": "🙌🏽",
		"descriptor": "Raising Hands: Medium Skin Tone"
	},
	"1F64C-1F3FE": {
		"key": "1F64C-1F3FE",
		"value": "🙌🏾",
		"descriptor": "Raising Hands: Medium-Dark Skin Tone"
	},
	"1F64C-1F3FF": {
		"key": "1F64C-1F3FF",
		"value": "🙌🏿",
		"descriptor": "Raising Hands: Dark Skin Tone"
	},
	"1F64D": {
		"key": "1F64D",
		"value": "🙍",
		"descriptor": "Person Frowning"
	},
	"1F64D-1F3FB": {
		"key": "1F64D-1F3FB",
		"value": "🙍🏻",
		"descriptor": "Person Frowning: Light Skin Tone"
	},
	"1F64D-1F3FB-200D-2640-FE0F": {
		"key": "1F64D-1F3FB-200D-2640-FE0F",
		"value": "🙍🏻‍♀️",
		"descriptor": "Woman Frowning: Light Skin Tone"
	},
	"1F64D-1F3FB-200D-2642-FE0F": {
		"key": "1F64D-1F3FB-200D-2642-FE0F",
		"value": "🙍🏻‍♂️",
		"descriptor": "Man Frowning: Light Skin Tone"
	},
	"1F64D-1F3FC": {
		"key": "1F64D-1F3FC",
		"value": "🙍🏼",
		"descriptor": "Person Frowning: Medium-Light Skin Tone"
	},
	"1F64D-1F3FC-200D-2640-FE0F": {
		"key": "1F64D-1F3FC-200D-2640-FE0F",
		"value": "🙍🏼‍♀️",
		"descriptor": "Woman Frowning: Medium-Light Skin Tone"
	},
	"1F64D-1F3FC-200D-2642-FE0F": {
		"key": "1F64D-1F3FC-200D-2642-FE0F",
		"value": "🙍🏼‍♂️",
		"descriptor": "Man Frowning: Medium-Light Skin Tone"
	},
	"1F64D-1F3FD": {
		"key": "1F64D-1F3FD",
		"value": "🙍🏽",
		"descriptor": "Person Frowning: Medium Skin Tone"
	},
	"1F64D-1F3FD-200D-2640-FE0F": {
		"key": "1F64D-1F3FD-200D-2640-FE0F",
		"value": "🙍🏽‍♀️",
		"descriptor": "Woman Frowning: Medium Skin Tone"
	},
	"1F64D-1F3FD-200D-2642-FE0F": {
		"key": "1F64D-1F3FD-200D-2642-FE0F",
		"value": "🙍🏽‍♂️",
		"descriptor": "Man Frowning: Medium Skin Tone"
	},
	"1F64D-1F3FE": {
		"key": "1F64D-1F3FE",
		"value": "🙍🏾",
		"descriptor": "Person Frowning: Medium-Dark Skin Tone"
	},
	"1F64D-1F3FE-200D-2640-FE0F": {
		"key": "1F64D-1F3FE-200D-2640-FE0F",
		"value": "🙍🏾‍♀️",
		"descriptor": "Woman Frowning: Medium-Dark Skin Tone"
	},
	"1F64D-1F3FE-200D-2642-FE0F": {
		"key": "1F64D-1F3FE-200D-2642-FE0F",
		"value": "🙍🏾‍♂️",
		"descriptor": "Man Frowning: Medium-Dark Skin Tone"
	},
	"1F64D-1F3FF": {
		"key": "1F64D-1F3FF",
		"value": "🙍🏿",
		"descriptor": "Person Frowning: Dark Skin Tone"
	},
	"1F64D-1F3FF-200D-2640-FE0F": {
		"key": "1F64D-1F3FF-200D-2640-FE0F",
		"value": "🙍🏿‍♀️",
		"descriptor": "Woman Frowning: Dark Skin Tone"
	},
	"1F64D-1F3FF-200D-2642-FE0F": {
		"key": "1F64D-1F3FF-200D-2642-FE0F",
		"value": "🙍🏿‍♂️",
		"descriptor": "Man Frowning: Dark Skin Tone"
	},
	"1F64D-200D-2640-FE0F": {
		"key": "1F64D-200D-2640-FE0F",
		"value": "🙍‍♀️",
		"descriptor": "Woman Frowning"
	},
	"1F64D-200D-2642-FE0F": {
		"key": "1F64D-200D-2642-FE0F",
		"value": "🙍‍♂️",
		"descriptor": "Man Frowning"
	},
	"1F64E": {
		"key": "1F64E",
		"value": "🙎",
		"descriptor": "Person Pouting"
	},
	"1F64E-1F3FB": {
		"key": "1F64E-1F3FB",
		"value": "🙎🏻",
		"descriptor": "Person Pouting: Light Skin Tone"
	},
	"1F64E-1F3FB-200D-2640-FE0F": {
		"key": "1F64E-1F3FB-200D-2640-FE0F",
		"value": "🙎🏻‍♀️",
		"descriptor": "Woman Pouting: Light Skin Tone"
	},
	"1F64E-1F3FB-200D-2642-FE0F": {
		"key": "1F64E-1F3FB-200D-2642-FE0F",
		"value": "🙎🏻‍♂️",
		"descriptor": "Man Pouting: Light Skin Tone"
	},
	"1F64E-1F3FC": {
		"key": "1F64E-1F3FC",
		"value": "🙎🏼",
		"descriptor": "Person Pouting: Medium-Light Skin Tone"
	},
	"1F64E-1F3FC-200D-2640-FE0F": {
		"key": "1F64E-1F3FC-200D-2640-FE0F",
		"value": "🙎🏼‍♀️",
		"descriptor": "Woman Pouting: Medium-Light Skin Tone"
	},
	"1F64E-1F3FC-200D-2642-FE0F": {
		"key": "1F64E-1F3FC-200D-2642-FE0F",
		"value": "🙎🏼‍♂️",
		"descriptor": "Man Pouting: Medium-Light Skin Tone"
	},
	"1F64E-1F3FD": {
		"key": "1F64E-1F3FD",
		"value": "🙎🏽",
		"descriptor": "Person Pouting: Medium Skin Tone"
	},
	"1F64E-1F3FD-200D-2640-FE0F": {
		"key": "1F64E-1F3FD-200D-2640-FE0F",
		"value": "🙎🏽‍♀️",
		"descriptor": "Woman Pouting: Medium Skin Tone"
	},
	"1F64E-1F3FD-200D-2642-FE0F": {
		"key": "1F64E-1F3FD-200D-2642-FE0F",
		"value": "🙎🏽‍♂️",
		"descriptor": "Man Pouting: Medium Skin Tone"
	},
	"1F64E-1F3FE": {
		"key": "1F64E-1F3FE",
		"value": "🙎🏾",
		"descriptor": "Person Pouting: Medium-Dark Skin Tone"
	},
	"1F64E-1F3FE-200D-2640-FE0F": {
		"key": "1F64E-1F3FE-200D-2640-FE0F",
		"value": "🙎🏾‍♀️",
		"descriptor": "Woman Pouting: Medium-Dark Skin Tone"
	},
	"1F64E-1F3FE-200D-2642-FE0F": {
		"key": "1F64E-1F3FE-200D-2642-FE0F",
		"value": "🙎🏾‍♂️",
		"descriptor": "Man Pouting: Medium-Dark Skin Tone"
	},
	"1F64E-1F3FF": {
		"key": "1F64E-1F3FF",
		"value": "🙎🏿",
		"descriptor": "Person Pouting: Dark Skin Tone"
	},
	"1F64E-1F3FF-200D-2640-FE0F": {
		"key": "1F64E-1F3FF-200D-2640-FE0F",
		"value": "🙎🏿‍♀️",
		"descriptor": "Woman Pouting: Dark Skin Tone"
	},
	"1F64E-1F3FF-200D-2642-FE0F": {
		"key": "1F64E-1F3FF-200D-2642-FE0F",
		"value": "🙎🏿‍♂️",
		"descriptor": "Man Pouting: Dark Skin Tone"
	},
	"1F64E-200D-2640-FE0F": {
		"key": "1F64E-200D-2640-FE0F",
		"value": "🙎‍♀️",
		"descriptor": "Woman Pouting"
	},
	"1F64E-200D-2642-FE0F": {
		"key": "1F64E-200D-2642-FE0F",
		"value": "🙎‍♂️",
		"descriptor": "Man Pouting"
	},
	"1F64F": {
		"key": "1F64F",
		"value": "🙏",
		"descriptor": "Folded Hands"
	},
	"1F64F-1F3FB": {
		"key": "1F64F-1F3FB",
		"value": "🙏🏻",
		"descriptor": "Folded Hands: Light Skin Tone"
	},
	"1F64F-1F3FC": {
		"key": "1F64F-1F3FC",
		"value": "🙏🏼",
		"descriptor": "Folded Hands: Medium-Light Skin Tone"
	},
	"1F64F-1F3FD": {
		"key": "1F64F-1F3FD",
		"value": "🙏🏽",
		"descriptor": "Folded Hands: Medium Skin Tone"
	},
	"1F64F-1F3FE": {
		"key": "1F64F-1F3FE",
		"value": "🙏🏾",
		"descriptor": "Folded Hands: Medium-Dark Skin Tone"
	},
	"1F64F-1F3FF": {
		"key": "1F64F-1F3FF",
		"value": "🙏🏿",
		"descriptor": "Folded Hands: Dark Skin Tone"
	},
	"1F680": {
		"key": "1F680",
		"value": "🚀",
		"descriptor": "Rocket"
	},
	"1F681": {
		"key": "1F681",
		"value": "🚁",
		"descriptor": "Helicopter"
	},
	"1F682": {
		"key": "1F682",
		"value": "🚂",
		"descriptor": "Locomotive"
	},
	"1F683": {
		"key": "1F683",
		"value": "🚃",
		"descriptor": "Railway Car"
	},
	"1F684": {
		"key": "1F684",
		"value": "🚄",
		"descriptor": "High-Speed Train"
	},
	"1F685": {
		"key": "1F685",
		"value": "🚅",
		"descriptor": "Bullet Train"
	},
	"1F686": {
		"key": "1F686",
		"value": "🚆",
		"descriptor": "Train"
	},
	"1F687": {
		"key": "1F687",
		"value": "🚇",
		"descriptor": "Metro"
	},
	"1F688": {
		"key": "1F688",
		"value": "🚈",
		"descriptor": "Light Rail"
	},
	"1F689": {
		"key": "1F689",
		"value": "🚉",
		"descriptor": "Station"
	},
	"1F68A": {
		"key": "1F68A",
		"value": "🚊",
		"descriptor": "Tram"
	},
	"1F68B": {
		"key": "1F68B",
		"value": "🚋",
		"descriptor": "Tram Car"
	},
	"1F68C": {
		"key": "1F68C",
		"value": "🚌",
		"descriptor": "Bus"
	},
	"1F68D": {
		"key": "1F68D",
		"value": "🚍",
		"descriptor": "Oncoming Bus"
	},
	"1F68E": {
		"key": "1F68E",
		"value": "🚎",
		"descriptor": "Trolleybus"
	},
	"1F68F": {
		"key": "1F68F",
		"value": "🚏",
		"descriptor": "Bus Stop"
	},
	"1F690": {
		"key": "1F690",
		"value": "🚐",
		"descriptor": "Minibus"
	},
	"1F691": {
		"key": "1F691",
		"value": "🚑",
		"descriptor": "Ambulance"
	},
	"1F692": {
		"key": "1F692",
		"value": "🚒",
		"descriptor": "Fire Engine"
	},
	"1F693": {
		"key": "1F693",
		"value": "🚓",
		"descriptor": "Police Car"
	},
	"1F694": {
		"key": "1F694",
		"value": "🚔",
		"descriptor": "Oncoming Police Car"
	},
	"1F695": {
		"key": "1F695",
		"value": "🚕",
		"descriptor": "Taxi"
	},
	"1F696": {
		"key": "1F696",
		"value": "🚖",
		"descriptor": "Oncoming Taxi"
	},
	"1F697": {
		"key": "1F697",
		"value": "🚗",
		"descriptor": "Automobile"
	},
	"1F698": {
		"key": "1F698",
		"value": "🚘",
		"descriptor": "Oncoming Automobile"
	},
	"1F699": {
		"key": "1F699",
		"value": "🚙",
		"descriptor": "Sport Utility Vehicle"
	},
	"1F69A": {
		"key": "1F69A",
		"value": "🚚",
		"descriptor": "Delivery Truck"
	},
	"1F69B": {
		"key": "1F69B",
		"value": "🚛",
		"descriptor": "Articulated Lorry"
	},
	"1F69C": {
		"key": "1F69C",
		"value": "🚜",
		"descriptor": "Tractor"
	},
	"1F69D": {
		"key": "1F69D",
		"value": "🚝",
		"descriptor": "Monorail"
	},
	"1F69E": {
		"key": "1F69E",
		"value": "🚞",
		"descriptor": "Mountain Railway"
	},
	"1F69F": {
		"key": "1F69F",
		"value": "🚟",
		"descriptor": "Suspension Railway"
	},
	"1F6A0": {
		"key": "1F6A0",
		"value": "🚠",
		"descriptor": "Mountain Cableway"
	},
	"1F6A1": {
		"key": "1F6A1",
		"value": "🚡",
		"descriptor": "Aerial Tramway"
	},
	"1F6A2": {
		"key": "1F6A2",
		"value": "🚢",
		"descriptor": "Ship"
	},
	"1F6A3": {
		"key": "1F6A3",
		"value": "🚣",
		"descriptor": "Person Rowing Boat"
	},
	"1F6A3-1F3FB": {
		"key": "1F6A3-1F3FB",
		"value": "🚣🏻",
		"descriptor": "Person Rowing Boat: Light Skin Tone"
	},
	"1F6A3-1F3FB-200D-2640-FE0F": {
		"key": "1F6A3-1F3FB-200D-2640-FE0F",
		"value": "🚣🏻‍♀️",
		"descriptor": "Woman Rowing Boat: Light Skin Tone"
	},
	"1F6A3-1F3FB-200D-2642-FE0F": {
		"key": "1F6A3-1F3FB-200D-2642-FE0F",
		"value": "🚣🏻‍♂️",
		"descriptor": "Man Rowing Boat: Light Skin Tone"
	},
	"1F6A3-1F3FC": {
		"key": "1F6A3-1F3FC",
		"value": "🚣🏼",
		"descriptor": "Person Rowing Boat: Medium-Light Skin Tone"
	},
	"1F6A3-1F3FC-200D-2640-FE0F": {
		"key": "1F6A3-1F3FC-200D-2640-FE0F",
		"value": "🚣🏼‍♀️",
		"descriptor": "Woman Rowing Boat: Medium-Light Skin Tone"
	},
	"1F6A3-1F3FC-200D-2642-FE0F": {
		"key": "1F6A3-1F3FC-200D-2642-FE0F",
		"value": "🚣🏼‍♂️",
		"descriptor": "Man Rowing Boat: Medium-Light Skin Tone"
	},
	"1F6A3-1F3FD": {
		"key": "1F6A3-1F3FD",
		"value": "🚣🏽",
		"descriptor": "Person Rowing Boat: Medium Skin Tone"
	},
	"1F6A3-1F3FD-200D-2640-FE0F": {
		"key": "1F6A3-1F3FD-200D-2640-FE0F",
		"value": "🚣🏽‍♀️",
		"descriptor": "Woman Rowing Boat: Medium Skin Tone"
	},
	"1F6A3-1F3FD-200D-2642-FE0F": {
		"key": "1F6A3-1F3FD-200D-2642-FE0F",
		"value": "🚣🏽‍♂️",
		"descriptor": "Man Rowing Boat: Medium Skin Tone"
	},
	"1F6A3-1F3FE": {
		"key": "1F6A3-1F3FE",
		"value": "🚣🏾",
		"descriptor": "Person Rowing Boat: Medium-Dark Skin Tone"
	},
	"1F6A3-1F3FE-200D-2640-FE0F": {
		"key": "1F6A3-1F3FE-200D-2640-FE0F",
		"value": "🚣🏾‍♀️",
		"descriptor": "Woman Rowing Boat: Medium-Dark Skin Tone"
	},
	"1F6A3-1F3FE-200D-2642-FE0F": {
		"key": "1F6A3-1F3FE-200D-2642-FE0F",
		"value": "🚣🏾‍♂️",
		"descriptor": "Man Rowing Boat: Medium-Dark Skin Tone"
	},
	"1F6A3-1F3FF": {
		"key": "1F6A3-1F3FF",
		"value": "🚣🏿",
		"descriptor": "Person Rowing Boat: Dark Skin Tone"
	},
	"1F6A3-1F3FF-200D-2640-FE0F": {
		"key": "1F6A3-1F3FF-200D-2640-FE0F",
		"value": "🚣🏿‍♀️",
		"descriptor": "Woman Rowing Boat: Dark Skin Tone"
	},
	"1F6A3-1F3FF-200D-2642-FE0F": {
		"key": "1F6A3-1F3FF-200D-2642-FE0F",
		"value": "🚣🏿‍♂️",
		"descriptor": "Man Rowing Boat: Dark Skin Tone"
	},
	"1F6A3-200D-2640-FE0F": {
		"key": "1F6A3-200D-2640-FE0F",
		"value": "🚣‍♀️",
		"descriptor": "Woman Rowing Boat"
	},
	"1F6A3-200D-2642-FE0F": {
		"key": "1F6A3-200D-2642-FE0F",
		"value": "🚣‍♂️",
		"descriptor": "Man Rowing Boat"
	},
	"1F6A4": {
		"key": "1F6A4",
		"value": "🚤",
		"descriptor": "Speedboat"
	},
	"1F6A5": {
		"key": "1F6A5",
		"value": "🚥",
		"descriptor": "Horizontal Traffic Light"
	},
	"1F6A6": {
		"key": "1F6A6",
		"value": "🚦",
		"descriptor": "Vertical Traffic Light"
	},
	"1F6A7": {
		"key": "1F6A7",
		"value": "🚧",
		"descriptor": "Construction"
	},
	"1F6A8": {
		"key": "1F6A8",
		"value": "🚨",
		"descriptor": "Police Car Light"
	},
	"1F6A9": {
		"key": "1F6A9",
		"value": "🚩",
		"descriptor": "Triangular Flag"
	},
	"1F6AA": {
		"key": "1F6AA",
		"value": "🚪",
		"descriptor": "Door"
	},
	"1F6AB": {
		"key": "1F6AB",
		"value": "🚫",
		"descriptor": "Prohibited"
	},
	"1F6AC": {
		"key": "1F6AC",
		"value": "🚬",
		"descriptor": "Cigarette"
	},
	"1F6AD": {
		"key": "1F6AD",
		"value": "🚭",
		"descriptor": "No Smoking"
	},
	"1F6AE": {
		"key": "1F6AE",
		"value": "🚮",
		"descriptor": "Litter in Bin Sign"
	},
	"1F6AF": {
		"key": "1F6AF",
		"value": "🚯",
		"descriptor": "No Littering"
	},
	"1F6B0": {
		"key": "1F6B0",
		"value": "🚰",
		"descriptor": "Potable Water"
	},
	"1F6B1": {
		"key": "1F6B1",
		"value": "🚱",
		"descriptor": "Non-Potable Water"
	},
	"1F6B2": {
		"key": "1F6B2",
		"value": "🚲",
		"descriptor": "Bicycle"
	},
	"1F6B3": {
		"key": "1F6B3",
		"value": "🚳",
		"descriptor": "No Bicycles"
	},
	"1F6B4": {
		"key": "1F6B4",
		"value": "🚴",
		"descriptor": "Person Biking"
	},
	"1F6B4-1F3FB": {
		"key": "1F6B4-1F3FB",
		"value": "🚴🏻",
		"descriptor": "Person Biking: Light Skin Tone"
	},
	"1F6B4-1F3FB-200D-2640-FE0F": {
		"key": "1F6B4-1F3FB-200D-2640-FE0F",
		"value": "🚴🏻‍♀️",
		"descriptor": "Woman Biking: Light Skin Tone"
	},
	"1F6B4-1F3FB-200D-2642-FE0F": {
		"key": "1F6B4-1F3FB-200D-2642-FE0F",
		"value": "🚴🏻‍♂️",
		"descriptor": "Man Biking: Light Skin Tone"
	},
	"1F6B4-1F3FC": {
		"key": "1F6B4-1F3FC",
		"value": "🚴🏼",
		"descriptor": "Person Biking: Medium-Light Skin Tone"
	},
	"1F6B4-1F3FC-200D-2640-FE0F": {
		"key": "1F6B4-1F3FC-200D-2640-FE0F",
		"value": "🚴🏼‍♀️",
		"descriptor": "Woman Biking: Medium-Light Skin Tone"
	},
	"1F6B4-1F3FC-200D-2642-FE0F": {
		"key": "1F6B4-1F3FC-200D-2642-FE0F",
		"value": "🚴🏼‍♂️",
		"descriptor": "Man Biking: Medium-Light Skin Tone"
	},
	"1F6B4-1F3FD": {
		"key": "1F6B4-1F3FD",
		"value": "🚴🏽",
		"descriptor": "Person Biking: Medium Skin Tone"
	},
	"1F6B4-1F3FD-200D-2640-FE0F": {
		"key": "1F6B4-1F3FD-200D-2640-FE0F",
		"value": "🚴🏽‍♀️",
		"descriptor": "Woman Biking: Medium Skin Tone"
	},
	"1F6B4-1F3FD-200D-2642-FE0F": {
		"key": "1F6B4-1F3FD-200D-2642-FE0F",
		"value": "🚴🏽‍♂️",
		"descriptor": "Man Biking: Medium Skin Tone"
	},
	"1F6B4-1F3FE": {
		"key": "1F6B4-1F3FE",
		"value": "🚴🏾",
		"descriptor": "Person Biking: Medium-Dark Skin Tone"
	},
	"1F6B4-1F3FE-200D-2640-FE0F": {
		"key": "1F6B4-1F3FE-200D-2640-FE0F",
		"value": "🚴🏾‍♀️",
		"descriptor": "Woman Biking: Medium-Dark Skin Tone"
	},
	"1F6B4-1F3FE-200D-2642-FE0F": {
		"key": "1F6B4-1F3FE-200D-2642-FE0F",
		"value": "🚴🏾‍♂️",
		"descriptor": "Man Biking: Medium-Dark Skin Tone"
	},
	"1F6B4-1F3FF": {
		"key": "1F6B4-1F3FF",
		"value": "🚴🏿",
		"descriptor": "Person Biking: Dark Skin Tone"
	},
	"1F6B4-1F3FF-200D-2640-FE0F": {
		"key": "1F6B4-1F3FF-200D-2640-FE0F",
		"value": "🚴🏿‍♀️",
		"descriptor": "Woman Biking: Dark Skin Tone"
	},
	"1F6B4-1F3FF-200D-2642-FE0F": {
		"key": "1F6B4-1F3FF-200D-2642-FE0F",
		"value": "🚴🏿‍♂️",
		"descriptor": "Man Biking: Dark Skin Tone"
	},
	"1F6B4-200D-2640-FE0F": {
		"key": "1F6B4-200D-2640-FE0F",
		"value": "🚴‍♀️",
		"descriptor": "Woman Biking"
	},
	"1F6B4-200D-2642-FE0F": {
		"key": "1F6B4-200D-2642-FE0F",
		"value": "🚴‍♂️",
		"descriptor": "Man Biking"
	},
	"1F6B5": {
		"key": "1F6B5",
		"value": "🚵",
		"descriptor": "Person Mountain Biking"
	},
	"1F6B5-1F3FB": {
		"key": "1F6B5-1F3FB",
		"value": "🚵🏻",
		"descriptor": "Person Mountain Biking: Light Skin Tone"
	},
	"1F6B5-1F3FB-200D-2640-FE0F": {
		"key": "1F6B5-1F3FB-200D-2640-FE0F",
		"value": "🚵🏻‍♀️",
		"descriptor": "Woman Mountain Biking: Light Skin Tone"
	},
	"1F6B5-1F3FB-200D-2642-FE0F": {
		"key": "1F6B5-1F3FB-200D-2642-FE0F",
		"value": "🚵🏻‍♂️",
		"descriptor": "Man Mountain Biking: Light Skin Tone"
	},
	"1F6B5-1F3FC": {
		"key": "1F6B5-1F3FC",
		"value": "🚵🏼",
		"descriptor": "Person Mountain Biking: Medium-Light Skin Tone"
	},
	"1F6B5-1F3FC-200D-2640-FE0F": {
		"key": "1F6B5-1F3FC-200D-2640-FE0F",
		"value": "🚵🏼‍♀️",
		"descriptor": "Woman Mountain Biking: Medium-Light Skin Tone"
	},
	"1F6B5-1F3FC-200D-2642-FE0F": {
		"key": "1F6B5-1F3FC-200D-2642-FE0F",
		"value": "🚵🏼‍♂️",
		"descriptor": "Man Mountain Biking: Medium-Light Skin Tone"
	},
	"1F6B5-1F3FD": {
		"key": "1F6B5-1F3FD",
		"value": "🚵🏽",
		"descriptor": "Person Mountain Biking: Medium Skin Tone"
	},
	"1F6B5-1F3FD-200D-2640-FE0F": {
		"key": "1F6B5-1F3FD-200D-2640-FE0F",
		"value": "🚵🏽‍♀️",
		"descriptor": "Woman Mountain Biking: Medium Skin Tone"
	},
	"1F6B5-1F3FD-200D-2642-FE0F": {
		"key": "1F6B5-1F3FD-200D-2642-FE0F",
		"value": "🚵🏽‍♂️",
		"descriptor": "Man Mountain Biking: Medium Skin Tone"
	},
	"1F6B5-1F3FE": {
		"key": "1F6B5-1F3FE",
		"value": "🚵🏾",
		"descriptor": "Person Mountain Biking: Medium-Dark Skin Tone"
	},
	"1F6B5-1F3FE-200D-2640-FE0F": {
		"key": "1F6B5-1F3FE-200D-2640-FE0F",
		"value": "🚵🏾‍♀️",
		"descriptor": "Woman Mountain Biking: Medium-Dark Skin Tone"
	},
	"1F6B5-1F3FE-200D-2642-FE0F": {
		"key": "1F6B5-1F3FE-200D-2642-FE0F",
		"value": "🚵🏾‍♂️",
		"descriptor": "Man Mountain Biking: Medium-Dark Skin Tone"
	},
	"1F6B5-1F3FF": {
		"key": "1F6B5-1F3FF",
		"value": "🚵🏿",
		"descriptor": "Person Mountain Biking: Dark Skin Tone"
	},
	"1F6B5-1F3FF-200D-2640-FE0F": {
		"key": "1F6B5-1F3FF-200D-2640-FE0F",
		"value": "🚵🏿‍♀️",
		"descriptor": "Woman Mountain Biking: Dark Skin Tone"
	},
	"1F6B5-1F3FF-200D-2642-FE0F": {
		"key": "1F6B5-1F3FF-200D-2642-FE0F",
		"value": "🚵🏿‍♂️",
		"descriptor": "Man Mountain Biking: Dark Skin Tone"
	},
	"1F6B5-200D-2640-FE0F": {
		"key": "1F6B5-200D-2640-FE0F",
		"value": "🚵‍♀️",
		"descriptor": "Woman Mountain Biking"
	},
	"1F6B5-200D-2642-FE0F": {
		"key": "1F6B5-200D-2642-FE0F",
		"value": "🚵‍♂️",
		"descriptor": "Man Mountain Biking"
	},
	"1F6B6": {
		"key": "1F6B6",
		"value": "🚶",
		"descriptor": "Person Walking"
	},
	"1F6B6-1F3FB": {
		"key": "1F6B6-1F3FB",
		"value": "🚶🏻",
		"descriptor": "Person Walking: Light Skin Tone"
	},
	"1F6B6-1F3FB-200D-2640-FE0F": {
		"key": "1F6B6-1F3FB-200D-2640-FE0F",
		"value": "🚶🏻‍♀️",
		"descriptor": "Woman Walking: Light Skin Tone"
	},
	"1F6B6-1F3FB-200D-2642-FE0F": {
		"key": "1F6B6-1F3FB-200D-2642-FE0F",
		"value": "🚶🏻‍♂️",
		"descriptor": "Man Walking: Light Skin Tone"
	},
	"1F6B6-1F3FC": {
		"key": "1F6B6-1F3FC",
		"value": "🚶🏼",
		"descriptor": "Person Walking: Medium-Light Skin Tone"
	},
	"1F6B6-1F3FC-200D-2640-FE0F": {
		"key": "1F6B6-1F3FC-200D-2640-FE0F",
		"value": "🚶🏼‍♀️",
		"descriptor": "Woman Walking: Medium-Light Skin Tone"
	},
	"1F6B6-1F3FC-200D-2642-FE0F": {
		"key": "1F6B6-1F3FC-200D-2642-FE0F",
		"value": "🚶🏼‍♂️",
		"descriptor": "Man Walking: Medium-Light Skin Tone"
	},
	"1F6B6-1F3FD": {
		"key": "1F6B6-1F3FD",
		"value": "🚶🏽",
		"descriptor": "Person Walking: Medium Skin Tone"
	},
	"1F6B6-1F3FD-200D-2640-FE0F": {
		"key": "1F6B6-1F3FD-200D-2640-FE0F",
		"value": "🚶🏽‍♀️",
		"descriptor": "Woman Walking: Medium Skin Tone"
	},
	"1F6B6-1F3FD-200D-2642-FE0F": {
		"key": "1F6B6-1F3FD-200D-2642-FE0F",
		"value": "🚶🏽‍♂️",
		"descriptor": "Man Walking: Medium Skin Tone"
	},
	"1F6B6-1F3FE": {
		"key": "1F6B6-1F3FE",
		"value": "🚶🏾",
		"descriptor": "Person Walking: Medium-Dark Skin Tone"
	},
	"1F6B6-1F3FE-200D-2640-FE0F": {
		"key": "1F6B6-1F3FE-200D-2640-FE0F",
		"value": "🚶🏾‍♀️",
		"descriptor": "Woman Walking: Medium-Dark Skin Tone"
	},
	"1F6B6-1F3FE-200D-2642-FE0F": {
		"key": "1F6B6-1F3FE-200D-2642-FE0F",
		"value": "🚶🏾‍♂️",
		"descriptor": "Man Walking: Medium-Dark Skin Tone"
	},
	"1F6B6-1F3FF": {
		"key": "1F6B6-1F3FF",
		"value": "🚶🏿",
		"descriptor": "Person Walking: Dark Skin Tone"
	},
	"1F6B6-1F3FF-200D-2640-FE0F": {
		"key": "1F6B6-1F3FF-200D-2640-FE0F",
		"value": "🚶🏿‍♀️",
		"descriptor": "Woman Walking: Dark Skin Tone"
	},
	"1F6B6-1F3FF-200D-2642-FE0F": {
		"key": "1F6B6-1F3FF-200D-2642-FE0F",
		"value": "🚶🏿‍♂️",
		"descriptor": "Man Walking: Dark Skin Tone"
	},
	"1F6B6-200D-2640-FE0F": {
		"key": "1F6B6-200D-2640-FE0F",
		"value": "🚶‍♀️",
		"descriptor": "Woman Walking"
	},
	"1F6B6-200D-2642-FE0F": {
		"key": "1F6B6-200D-2642-FE0F",
		"value": "🚶‍♂️",
		"descriptor": "Man Walking"
	},
	"1F6B7": {
		"key": "1F6B7",
		"value": "🚷",
		"descriptor": "No Pedestrians"
	},
	"1F6B8": {
		"key": "1F6B8",
		"value": "🚸",
		"descriptor": "Children Crossing"
	},
	"1F6B9": {
		"key": "1F6B9",
		"value": "🚹",
		"descriptor": "Men’s Room"
	},
	"1F6BA": {
		"key": "1F6BA",
		"value": "🚺",
		"descriptor": "Women’s Room"
	},
	"1F6BB": {
		"key": "1F6BB",
		"value": "🚻",
		"descriptor": "Restroom"
	},
	"1F6BC": {
		"key": "1F6BC",
		"value": "🚼",
		"descriptor": "Baby Symbol"
	},
	"1F6BD": {
		"key": "1F6BD",
		"value": "🚽",
		"descriptor": "Toilet"
	},
	"1F6BE": {
		"key": "1F6BE",
		"value": "🚾",
		"descriptor": "Water Closet"
	},
	"1F6BF": {
		"key": "1F6BF",
		"value": "🚿",
		"descriptor": "Shower"
	},
	"1F6C0": {
		"key": "1F6C0",
		"value": "🛀",
		"descriptor": "Person Taking Bath"
	},
	"1F6C0-1F3FB": {
		"key": "1F6C0-1F3FB",
		"value": "🛀🏻",
		"descriptor": "Person Taking Bath: Light Skin Tone"
	},
	"1F6C0-1F3FC": {
		"key": "1F6C0-1F3FC",
		"value": "🛀🏼",
		"descriptor": "Person Taking Bath: Medium-Light Skin Tone"
	},
	"1F6C0-1F3FD": {
		"key": "1F6C0-1F3FD",
		"value": "🛀🏽",
		"descriptor": "Person Taking Bath: Medium Skin Tone"
	},
	"1F6C0-1F3FE": {
		"key": "1F6C0-1F3FE",
		"value": "🛀🏾",
		"descriptor": "Person Taking Bath: Medium-Dark Skin Tone"
	},
	"1F6C0-1F3FF": {
		"key": "1F6C0-1F3FF",
		"value": "🛀🏿",
		"descriptor": "Person Taking Bath: Dark Skin Tone"
	},
	"1F6C1": {
		"key": "1F6C1",
		"value": "🛁",
		"descriptor": "Bathtub"
	},
	"1F6C2": {
		"key": "1F6C2",
		"value": "🛂",
		"descriptor": "Passport Control"
	},
	"1F6C3": {
		"key": "1F6C3",
		"value": "🛃",
		"descriptor": "Customs"
	},
	"1F6C4": {
		"key": "1F6C4",
		"value": "🛄",
		"descriptor": "Baggage Claim"
	},
	"1F6C5": {
		"key": "1F6C5",
		"value": "🛅",
		"descriptor": "Left Luggage"
	},
	"1F6CB-FE0F": {
		"key": "1F6CB-FE0F",
		"value": "🛋️",
		"descriptor": "Couch and Lamp"
	},
	"1F6CC": {
		"key": "1F6CC",
		"value": "🛌",
		"descriptor": "Person in Bed"
	},
	"1F6CD-FE0F": {
		"key": "1F6CD-FE0F",
		"value": "🛍️",
		"descriptor": "Shopping Bags"
	},
	"1F6CE-FE0F": {
		"key": "1F6CE-FE0F",
		"value": "🛎️",
		"descriptor": "Bellhop Bell"
	},
	"1F6CF-FE0F": {
		"key": "1F6CF-FE0F",
		"value": "🛏️",
		"descriptor": "Bed"
	},
	"1F6D0": {
		"key": "1F6D0",
		"value": "🛐",
		"descriptor": "Place of Worship"
	},
	"1F6D1": {
		"key": "1F6D1",
		"value": "🛑",
		"descriptor": "Stop Sign"
	},
	"1F6D2": {
		"key": "1F6D2",
		"value": "🛒",
		"descriptor": "Shopping Cart"
	},
	"1F6D5": {
		"key": "1F6D5",
		"value": "🛕",
		"descriptor": "Hindu Temple"
	},
	"1F6E0-FE0F": {
		"key": "1F6E0-FE0F",
		"value": "🛠️",
		"descriptor": "Hammer and Wrench"
	},
	"1F6E1-FE0F": {
		"key": "1F6E1-FE0F",
		"value": "🛡️",
		"descriptor": "Shield"
	},
	"1F6E2-FE0F": {
		"key": "1F6E2-FE0F",
		"value": "🛢️",
		"descriptor": "Oil Drum"
	},
	"1F6E3-FE0F": {
		"key": "1F6E3-FE0F",
		"value": "🛣️",
		"descriptor": "Motorway"
	},
	"1F6E4-FE0F": {
		"key": "1F6E4-FE0F",
		"value": "🛤️",
		"descriptor": "Railway Track"
	},
	"1F6E5-FE0F": {
		"key": "1F6E5-FE0F",
		"value": "🛥️",
		"descriptor": "Motor Boat"
	},
	"1F6E9-FE0F": {
		"key": "1F6E9-FE0F",
		"value": "🛩️",
		"descriptor": "Small Airplane"
	},
	"1F6EB": {
		"key": "1F6EB",
		"value": "🛫",
		"descriptor": "Airplane Departure"
	},
	"1F6EC": {
		"key": "1F6EC",
		"value": "🛬",
		"descriptor": "Airplane Arrival"
	},
	"1F6F0-FE0F": {
		"key": "1F6F0-FE0F",
		"value": "🛰️",
		"descriptor": "Satellite"
	},
	"1F6F3-FE0F": {
		"key": "1F6F3-FE0F",
		"value": "🛳️",
		"descriptor": "Passenger Ship"
	},
	"1F6F4": {
		"key": "1F6F4",
		"value": "🛴",
		"descriptor": "Kick Scooter"
	},
	"1F6F5": {
		"key": "1F6F5",
		"value": "🛵",
		"descriptor": "Motor Scooter"
	},
	"1F6F6": {
		"key": "1F6F6",
		"value": "🛶",
		"descriptor": "Canoe"
	},
	"1F6F7": {
		"key": "1F6F7",
		"value": "🛷",
		"descriptor": "Sled"
	},
	"1F6F8": {
		"key": "1F6F8",
		"value": "🛸",
		"descriptor": "Flying Saucer"
	},
	"1F6F9": {
		"key": "1F6F9",
		"value": "🛹",
		"descriptor": "Skateboard"
	},
	"1F6FA": {
		"key": "1F6FA",
		"value": "🛺",
		"descriptor": "Auto Rickshaw"
	},
	"1F7E0": {
		"key": "1F7E0",
		"value": "🟠",
		"descriptor": "Orange Circle"
	},
	"1F7E1": {
		"key": "1F7E1",
		"value": "🟡",
		"descriptor": "Yellow Circle"
	},
	"1F7E2": {
		"key": "1F7E2",
		"value": "🟢",
		"descriptor": "Green Circle"
	},
	"1F7E3": {
		"key": "1F7E3",
		"value": "🟣",
		"descriptor": "Purple Circle"
	},
	"1F7E4": {
		"key": "1F7E4",
		"value": "🟤",
		"descriptor": "Brown Circle"
	},
	"1F7E5": {
		"key": "1F7E5",
		"value": "🟥",
		"descriptor": "Red Square"
	},
	"1F7E6": {
		"key": "1F7E6",
		"value": "🟦",
		"descriptor": "Blue Square"
	},
	"1F7E7": {
		"key": "1F7E7",
		"value": "🟧",
		"descriptor": "Orange Square"
	},
	"1F7E8": {
		"key": "1F7E8",
		"value": "🟨",
		"descriptor": "Yellow Square"
	},
	"1F7E9": {
		"key": "1F7E9",
		"value": "🟩",
		"descriptor": "Green Square"
	},
	"1F7EA": {
		"key": "1F7EA",
		"value": "🟪",
		"descriptor": "Purple Square"
	},
	"1F7EB": {
		"key": "1F7EB",
		"value": "🟫",
		"descriptor": "Brown Square"
	},
	"1F90D": {
		"key": "1F90D",
		"value": "🤍",
		"descriptor": "White Heart"
	},
	"1F90E": {
		"key": "1F90E",
		"value": "🤎",
		"descriptor": "Brown Heart"
	},
	"1F90F": {
		"key": "1F90F",
		"value": "🤏",
		"descriptor": "Pinching Hand"
	},
	"1F90F-1F3FB": {
		"key": "1F90F-1F3FB",
		"value": "🤏🏻",
		"descriptor": "Pinching Hand: Light Skin Tone"
	},
	"1F90F-1F3FC": {
		"key": "1F90F-1F3FC",
		"value": "🤏🏼",
		"descriptor": "Pinching Hand: Medium-Light Skin Tone"
	},
	"1F90F-1F3FD": {
		"key": "1F90F-1F3FD",
		"value": "🤏🏽",
		"descriptor": "Pinching Hand: Medium Skin Tone"
	},
	"1F90F-1F3FE": {
		"key": "1F90F-1F3FE",
		"value": "🤏🏾",
		"descriptor": "Pinching Hand: Medium-Dark Skin Tone"
	},
	"1F90F-1F3FF": {
		"key": "1F90F-1F3FF",
		"value": "🤏🏿",
		"descriptor": "Pinching Hand: Dark Skin Tone"
	},
	"1F910": {
		"key": "1F910",
		"value": "🤐",
		"descriptor": "Zipper-Mouth Face"
	},
	"1F911": {
		"key": "1F911",
		"value": "🤑",
		"descriptor": "Money-Mouth Face"
	},
	"1F912": {
		"key": "1F912",
		"value": "🤒",
		"descriptor": "Face With Thermometer"
	},
	"1F913": {
		"key": "1F913",
		"value": "🤓",
		"descriptor": "Nerd Face"
	},
	"1F914": {
		"key": "1F914",
		"value": "🤔",
		"descriptor": "Thinking Face"
	},
	"1F915": {
		"key": "1F915",
		"value": "🤕",
		"descriptor": "Face With Head-Bandage"
	},
	"1F916": {
		"key": "1F916",
		"value": "🤖",
		"descriptor": "Robot"
	},
	"1F917": {
		"key": "1F917",
		"value": "🤗",
		"descriptor": "Hugging Face"
	},
	"1F918": {
		"key": "1F918",
		"value": "🤘",
		"descriptor": "Sign of the Horns"
	},
	"1F918-1F3FB": {
		"key": "1F918-1F3FB",
		"value": "🤘🏻",
		"descriptor": "Sign of the Horns: Light Skin Tone"
	},
	"1F918-1F3FC": {
		"key": "1F918-1F3FC",
		"value": "🤘🏼",
		"descriptor": "Sign of the Horns: Medium-Light Skin Tone"
	},
	"1F918-1F3FD": {
		"key": "1F918-1F3FD",
		"value": "🤘🏽",
		"descriptor": "Sign of the Horns: Medium Skin Tone"
	},
	"1F918-1F3FE": {
		"key": "1F918-1F3FE",
		"value": "🤘🏾",
		"descriptor": "Sign of the Horns: Medium-Dark Skin Tone"
	},
	"1F918-1F3FF": {
		"key": "1F918-1F3FF",
		"value": "🤘🏿",
		"descriptor": "Sign of the Horns: Dark Skin Tone"
	},
	"1F919": {
		"key": "1F919",
		"value": "🤙",
		"descriptor": "Call Me Hand"
	},
	"1F919-1F3FB": {
		"key": "1F919-1F3FB",
		"value": "🤙🏻",
		"descriptor": "Call Me Hand: Light Skin Tone"
	},
	"1F919-1F3FC": {
		"key": "1F919-1F3FC",
		"value": "🤙🏼",
		"descriptor": "Call Me Hand: Medium-Light Skin Tone"
	},
	"1F919-1F3FD": {
		"key": "1F919-1F3FD",
		"value": "🤙🏽",
		"descriptor": "Call Me Hand: Medium Skin Tone"
	},
	"1F919-1F3FE": {
		"key": "1F919-1F3FE",
		"value": "🤙🏾",
		"descriptor": "Call Me Hand: Medium-Dark Skin Tone"
	},
	"1F919-1F3FF": {
		"key": "1F919-1F3FF",
		"value": "🤙🏿",
		"descriptor": "Call Me Hand: Dark Skin Tone"
	},
	"1F91A": {
		"key": "1F91A",
		"value": "🤚",
		"descriptor": "Raised Back of Hand"
	},
	"1F91A-1F3FB": {
		"key": "1F91A-1F3FB",
		"value": "🤚🏻",
		"descriptor": "Raised Back of Hand: Light Skin Tone"
	},
	"1F91A-1F3FC": {
		"key": "1F91A-1F3FC",
		"value": "🤚🏼",
		"descriptor": "Raised Back of Hand: Medium-Light Skin Tone"
	},
	"1F91A-1F3FD": {
		"key": "1F91A-1F3FD",
		"value": "🤚🏽",
		"descriptor": "Raised Back of Hand: Medium Skin Tone"
	},
	"1F91A-1F3FE": {
		"key": "1F91A-1F3FE",
		"value": "🤚🏾",
		"descriptor": "Raised Back of Hand: Medium-Dark Skin Tone"
	},
	"1F91A-1F3FF": {
		"key": "1F91A-1F3FF",
		"value": "🤚🏿",
		"descriptor": "Raised Back of Hand: Dark Skin Tone"
	},
	"1F91B": {
		"key": "1F91B",
		"value": "🤛",
		"descriptor": "Left-Facing Fist"
	},
	"1F91B-1F3FB": {
		"key": "1F91B-1F3FB",
		"value": "🤛🏻",
		"descriptor": "Left-Facing Fist: Light Skin Tone"
	},
	"1F91B-1F3FC": {
		"key": "1F91B-1F3FC",
		"value": "🤛🏼",
		"descriptor": "Left-Facing Fist: Medium-Light Skin Tone"
	},
	"1F91B-1F3FD": {
		"key": "1F91B-1F3FD",
		"value": "🤛🏽",
		"descriptor": "Left-Facing Fist: Medium Skin Tone"
	},
	"1F91B-1F3FE": {
		"key": "1F91B-1F3FE",
		"value": "🤛🏾",
		"descriptor": "Left-Facing Fist: Medium-Dark Skin Tone"
	},
	"1F91B-1F3FF": {
		"key": "1F91B-1F3FF",
		"value": "🤛🏿",
		"descriptor": "Left-Facing Fist: Dark Skin Tone"
	},
	"1F91C": {
		"key": "1F91C",
		"value": "🤜",
		"descriptor": "Right-Facing Fist"
	},
	"1F91C-1F3FB": {
		"key": "1F91C-1F3FB",
		"value": "🤜🏻",
		"descriptor": "Right-Facing Fist: Light Skin Tone"
	},
	"1F91C-1F3FC": {
		"key": "1F91C-1F3FC",
		"value": "🤜🏼",
		"descriptor": "Right-Facing Fist: Medium-Light Skin Tone"
	},
	"1F91C-1F3FD": {
		"key": "1F91C-1F3FD",
		"value": "🤜🏽",
		"descriptor": "Right-Facing Fist: Medium Skin Tone"
	},
	"1F91C-1F3FE": {
		"key": "1F91C-1F3FE",
		"value": "🤜🏾",
		"descriptor": "Right-Facing Fist: Medium-Dark Skin Tone"
	},
	"1F91C-1F3FF": {
		"key": "1F91C-1F3FF",
		"value": "🤜🏿",
		"descriptor": "Right-Facing Fist: Dark Skin Tone"
	},
	"1F91D": {
		"key": "1F91D",
		"value": "🤝",
		"descriptor": "Handshake"
	},
	"1F91E": {
		"key": "1F91E",
		"value": "🤞",
		"descriptor": "Crossed Fingers"
	},
	"1F91E-1F3FB": {
		"key": "1F91E-1F3FB",
		"value": "🤞🏻",
		"descriptor": "Crossed Fingers: Light Skin Tone"
	},
	"1F91E-1F3FC": {
		"key": "1F91E-1F3FC",
		"value": "🤞🏼",
		"descriptor": "Crossed Fingers: Medium-Light Skin Tone"
	},
	"1F91E-1F3FD": {
		"key": "1F91E-1F3FD",
		"value": "🤞🏽",
		"descriptor": "Crossed Fingers: Medium Skin Tone"
	},
	"1F91E-1F3FE": {
		"key": "1F91E-1F3FE",
		"value": "🤞🏾",
		"descriptor": "Crossed Fingers: Medium-Dark Skin Tone"
	},
	"1F91E-1F3FF": {
		"key": "1F91E-1F3FF",
		"value": "🤞🏿",
		"descriptor": "Crossed Fingers: Dark Skin Tone"
	},
	"1F91F": {
		"key": "1F91F",
		"value": "🤟",
		"descriptor": "Love-You Gesture"
	},
	"1F91F-1F3FB": {
		"key": "1F91F-1F3FB",
		"value": "🤟🏻",
		"descriptor": "Love-You Gesture: Light Skin Tone"
	},
	"1F91F-1F3FC": {
		"key": "1F91F-1F3FC",
		"value": "🤟🏼",
		"descriptor": "Love-You Gesture: Medium-Light Skin Tone"
	},
	"1F91F-1F3FD": {
		"key": "1F91F-1F3FD",
		"value": "🤟🏽",
		"descriptor": "Love-You Gesture: Medium Skin Tone"
	},
	"1F91F-1F3FE": {
		"key": "1F91F-1F3FE",
		"value": "🤟🏾",
		"descriptor": "Love-You Gesture: Medium-Dark Skin Tone"
	},
	"1F91F-1F3FF": {
		"key": "1F91F-1F3FF",
		"value": "🤟🏿",
		"descriptor": "Love-You Gesture: Dark Skin Tone"
	},
	"1F920": {
		"key": "1F920",
		"value": "🤠",
		"descriptor": "Cowboy Hat Face"
	},
	"1F921": {
		"key": "1F921",
		"value": "🤡",
		"descriptor": "Clown Face"
	},
	"1F922": {
		"key": "1F922",
		"value": "🤢",
		"descriptor": "Nauseated Face"
	},
	"1F923": {
		"key": "1F923",
		"value": "🤣",
		"descriptor": "Rolling on the Floor Laughing"
	},
	"1F924": {
		"key": "1F924",
		"value": "🤤",
		"descriptor": "Drooling Face"
	},
	"1F925": {
		"key": "1F925",
		"value": "🤥",
		"descriptor": "Lying Face"
	},
	"1F926": {
		"key": "1F926",
		"value": "🤦",
		"descriptor": "Person Facepalming"
	},
	"1F926-1F3FB": {
		"key": "1F926-1F3FB",
		"value": "🤦🏻",
		"descriptor": "Person Facepalming: Light Skin Tone"
	},
	"1F926-1F3FB-200D-2640-FE0F": {
		"key": "1F926-1F3FB-200D-2640-FE0F",
		"value": "🤦🏻‍♀️",
		"descriptor": "Woman Facepalming: Light Skin Tone"
	},
	"1F926-1F3FB-200D-2642-FE0F": {
		"key": "1F926-1F3FB-200D-2642-FE0F",
		"value": "🤦🏻‍♂️",
		"descriptor": "Man Facepalming: Light Skin Tone"
	},
	"1F926-1F3FC": {
		"key": "1F926-1F3FC",
		"value": "🤦🏼",
		"descriptor": "Person Facepalming: Medium-Light Skin Tone"
	},
	"1F926-1F3FC-200D-2640-FE0F": {
		"key": "1F926-1F3FC-200D-2640-FE0F",
		"value": "🤦🏼‍♀️",
		"descriptor": "Woman Facepalming: Medium-Light Skin Tone"
	},
	"1F926-1F3FC-200D-2642-FE0F": {
		"key": "1F926-1F3FC-200D-2642-FE0F",
		"value": "🤦🏼‍♂️",
		"descriptor": "Man Facepalming: Medium-Light Skin Tone"
	},
	"1F926-1F3FD": {
		"key": "1F926-1F3FD",
		"value": "🤦🏽",
		"descriptor": "Person Facepalming: Medium Skin Tone"
	},
	"1F926-1F3FD-200D-2640-FE0F": {
		"key": "1F926-1F3FD-200D-2640-FE0F",
		"value": "🤦🏽‍♀️",
		"descriptor": "Woman Facepalming: Medium Skin Tone"
	},
	"1F926-1F3FD-200D-2642-FE0F": {
		"key": "1F926-1F3FD-200D-2642-FE0F",
		"value": "🤦🏽‍♂️",
		"descriptor": "Man Facepalming: Medium Skin Tone"
	},
	"1F926-1F3FE": {
		"key": "1F926-1F3FE",
		"value": "🤦🏾",
		"descriptor": "Person Facepalming: Medium-Dark Skin Tone"
	},
	"1F926-1F3FE-200D-2640-FE0F": {
		"key": "1F926-1F3FE-200D-2640-FE0F",
		"value": "🤦🏾‍♀️",
		"descriptor": "Woman Facepalming: Medium-Dark Skin Tone"
	},
	"1F926-1F3FE-200D-2642-FE0F": {
		"key": "1F926-1F3FE-200D-2642-FE0F",
		"value": "🤦🏾‍♂️",
		"descriptor": "Man Facepalming: Medium-Dark Skin Tone"
	},
	"1F926-1F3FF": {
		"key": "1F926-1F3FF",
		"value": "🤦🏿",
		"descriptor": "Person Facepalming: Dark Skin Tone"
	},
	"1F926-1F3FF-200D-2640-FE0F": {
		"key": "1F926-1F3FF-200D-2640-FE0F",
		"value": "🤦🏿‍♀️",
		"descriptor": "Woman Facepalming: Dark Skin Tone"
	},
	"1F926-1F3FF-200D-2642-FE0F": {
		"key": "1F926-1F3FF-200D-2642-FE0F",
		"value": "🤦🏿‍♂️",
		"descriptor": "Man Facepalming: Dark Skin Tone"
	},
	"1F926-200D-2640-FE0F": {
		"key": "1F926-200D-2640-FE0F",
		"value": "🤦‍♀️",
		"descriptor": "Woman Facepalming"
	},
	"1F926-200D-2642-FE0F": {
		"key": "1F926-200D-2642-FE0F",
		"value": "🤦‍♂️",
		"descriptor": "Man Facepalming"
	},
	"1F927": {
		"key": "1F927",
		"value": "🤧",
		"descriptor": "Sneezing Face"
	},
	"1F928": {
		"key": "1F928",
		"value": "🤨",
		"descriptor": "Face With Raised Eyebrow"
	},
	"1F929": {
		"key": "1F929",
		"value": "🤩",
		"descriptor": "Star-Struck"
	},
	"1F92A": {
		"key": "1F92A",
		"value": "🤪",
		"descriptor": "Zany Face"
	},
	"1F92B": {
		"key": "1F92B",
		"value": "🤫",
		"descriptor": "Shushing Face"
	},
	"1F92C": {
		"key": "1F92C",
		"value": "🤬",
		"descriptor": "Face With Symbols on Mouth"
	},
	"1F92D": {
		"key": "1F92D",
		"value": "🤭",
		"descriptor": "Face With Hand Over Mouth"
	},
	"1F92E": {
		"key": "1F92E",
		"value": "🤮",
		"descriptor": "Face Vomiting"
	},
	"1F92F": {
		"key": "1F92F",
		"value": "🤯",
		"descriptor": "Exploding Head"
	},
	"1F930": {
		"key": "1F930",
		"value": "🤰",
		"descriptor": "Pregnant Woman"
	},
	"1F930-1F3FB": {
		"key": "1F930-1F3FB",
		"value": "🤰🏻",
		"descriptor": "Pregnant Woman: Light Skin Tone"
	},
	"1F930-1F3FC": {
		"key": "1F930-1F3FC",
		"value": "🤰🏼",
		"descriptor": "Pregnant Woman: Medium-Light Skin Tone"
	},
	"1F930-1F3FD": {
		"key": "1F930-1F3FD",
		"value": "🤰🏽",
		"descriptor": "Pregnant Woman: Medium Skin Tone"
	},
	"1F930-1F3FE": {
		"key": "1F930-1F3FE",
		"value": "🤰🏾",
		"descriptor": "Pregnant Woman: Medium-Dark Skin Tone"
	},
	"1F930-1F3FF": {
		"key": "1F930-1F3FF",
		"value": "🤰🏿",
		"descriptor": "Pregnant Woman: Dark Skin Tone"
	},
	"1F931": {
		"key": "1F931",
		"value": "🤱",
		"descriptor": "Breast-Feeding"
	},
	"1F931-1F3FB": {
		"key": "1F931-1F3FB",
		"value": "🤱🏻",
		"descriptor": "Breast-Feeding: Light Skin Tone"
	},
	"1F931-1F3FC": {
		"key": "1F931-1F3FC",
		"value": "🤱🏼",
		"descriptor": "Breast-Feeding: Medium-Light Skin Tone"
	},
	"1F931-1F3FD": {
		"key": "1F931-1F3FD",
		"value": "🤱🏽",
		"descriptor": "Breast-Feeding: Medium Skin Tone"
	},
	"1F931-1F3FE": {
		"key": "1F931-1F3FE",
		"value": "🤱🏾",
		"descriptor": "Breast-Feeding: Medium-Dark Skin Tone"
	},
	"1F931-1F3FF": {
		"key": "1F931-1F3FF",
		"value": "🤱🏿",
		"descriptor": "Breast-Feeding: Dark Skin Tone"
	},
	"1F932": {
		"key": "1F932",
		"value": "🤲",
		"descriptor": "Palms Up Together"
	},
	"1F932-1F3FB": {
		"key": "1F932-1F3FB",
		"value": "🤲🏻",
		"descriptor": "Palms Up Together: Light Skin Tone"
	},
	"1F932-1F3FC": {
		"key": "1F932-1F3FC",
		"value": "🤲🏼",
		"descriptor": "Palms Up Together: Medium-Light Skin Tone"
	},
	"1F932-1F3FD": {
		"key": "1F932-1F3FD",
		"value": "🤲🏽",
		"descriptor": "Palms Up Together: Medium Skin Tone"
	},
	"1F932-1F3FE": {
		"key": "1F932-1F3FE",
		"value": "🤲🏾",
		"descriptor": "Palms Up Together: Medium-Dark Skin Tone"
	},
	"1F932-1F3FF": {
		"key": "1F932-1F3FF",
		"value": "🤲🏿",
		"descriptor": "Palms Up Together: Dark Skin Tone"
	},
	"1F933": {
		"key": "1F933",
		"value": "🤳",
		"descriptor": "Selfie"
	},
	"1F933-1F3FB": {
		"key": "1F933-1F3FB",
		"value": "🤳🏻",
		"descriptor": "Selfie: Light Skin Tone"
	},
	"1F933-1F3FC": {
		"key": "1F933-1F3FC",
		"value": "🤳🏼",
		"descriptor": "Selfie: Medium-Light Skin Tone"
	},
	"1F933-1F3FD": {
		"key": "1F933-1F3FD",
		"value": "🤳🏽",
		"descriptor": "Selfie: Medium Skin Tone"
	},
	"1F933-1F3FE": {
		"key": "1F933-1F3FE",
		"value": "🤳🏾",
		"descriptor": "Selfie: Medium-Dark Skin Tone"
	},
	"1F933-1F3FF": {
		"key": "1F933-1F3FF",
		"value": "🤳🏿",
		"descriptor": "Selfie: Dark Skin Tone"
	},
	"1F934": {
		"key": "1F934",
		"value": "🤴",
		"descriptor": "Prince"
	},
	"1F934-1F3FB": {
		"key": "1F934-1F3FB",
		"value": "🤴🏻",
		"descriptor": "Prince: Light Skin Tone"
	},
	"1F934-1F3FC": {
		"key": "1F934-1F3FC",
		"value": "🤴🏼",
		"descriptor": "Prince: Medium-Light Skin Tone"
	},
	"1F934-1F3FD": {
		"key": "1F934-1F3FD",
		"value": "🤴🏽",
		"descriptor": "Prince: Medium Skin Tone"
	},
	"1F934-1F3FE": {
		"key": "1F934-1F3FE",
		"value": "🤴🏾",
		"descriptor": "Prince: Medium-Dark Skin Tone"
	},
	"1F934-1F3FF": {
		"key": "1F934-1F3FF",
		"value": "🤴🏿",
		"descriptor": "Prince: Dark Skin Tone"
	},
	"1F935": {
		"key": "1F935",
		"value": "🤵",
		"descriptor": "Man in Tuxedo"
	},
	"1F935-1F3FB": {
		"key": "1F935-1F3FB",
		"value": "🤵🏻",
		"descriptor": "Man in Tuxedo: Light Skin Tone"
	},
	"1F935-1F3FC": {
		"key": "1F935-1F3FC",
		"value": "🤵🏼",
		"descriptor": "Man in Tuxedo: Medium-Light Skin Tone"
	},
	"1F935-1F3FD": {
		"key": "1F935-1F3FD",
		"value": "🤵🏽",
		"descriptor": "Man in Tuxedo: Medium Skin Tone"
	},
	"1F935-1F3FE": {
		"key": "1F935-1F3FE",
		"value": "🤵🏾",
		"descriptor": "Man in Tuxedo: Medium-Dark Skin Tone"
	},
	"1F935-1F3FF": {
		"key": "1F935-1F3FF",
		"value": "🤵🏿",
		"descriptor": "Man in Tuxedo: Dark Skin Tone"
	},
	"1F936": {
		"key": "1F936",
		"value": "🤶",
		"descriptor": "Mrs. Claus"
	},
	"1F936-1F3FB": {
		"key": "1F936-1F3FB",
		"value": "🤶🏻",
		"descriptor": "Mrs. Claus: Light Skin Tone"
	},
	"1F936-1F3FC": {
		"key": "1F936-1F3FC",
		"value": "🤶🏼",
		"descriptor": "Mrs. Claus: Medium-Light Skin Tone"
	},
	"1F936-1F3FD": {
		"key": "1F936-1F3FD",
		"value": "🤶🏽",
		"descriptor": "Mrs. Claus: Medium Skin Tone"
	},
	"1F936-1F3FE": {
		"key": "1F936-1F3FE",
		"value": "🤶🏾",
		"descriptor": "Mrs. Claus: Medium-Dark Skin Tone"
	},
	"1F936-1F3FF": {
		"key": "1F936-1F3FF",
		"value": "🤶🏿",
		"descriptor": "Mrs. Claus: Dark Skin Tone"
	},
	"1F937": {
		"key": "1F937",
		"value": "🤷",
		"descriptor": "Person Shrugging"
	},
	"1F937-1F3FB": {
		"key": "1F937-1F3FB",
		"value": "🤷🏻",
		"descriptor": "Person Shrugging: Light Skin Tone"
	},
	"1F937-1F3FB-200D-2640-FE0F": {
		"key": "1F937-1F3FB-200D-2640-FE0F",
		"value": "🤷🏻‍♀️",
		"descriptor": "Woman Shrugging: Light Skin Tone"
	},
	"1F937-1F3FB-200D-2642-FE0F": {
		"key": "1F937-1F3FB-200D-2642-FE0F",
		"value": "🤷🏻‍♂️",
		"descriptor": "Man Shrugging: Light Skin Tone"
	},
	"1F937-1F3FC": {
		"key": "1F937-1F3FC",
		"value": "🤷🏼",
		"descriptor": "Person Shrugging: Medium-Light Skin Tone"
	},
	"1F937-1F3FC-200D-2640-FE0F": {
		"key": "1F937-1F3FC-200D-2640-FE0F",
		"value": "🤷🏼‍♀️",
		"descriptor": "Woman Shrugging: Medium-Light Skin Tone"
	},
	"1F937-1F3FC-200D-2642-FE0F": {
		"key": "1F937-1F3FC-200D-2642-FE0F",
		"value": "🤷🏼‍♂️",
		"descriptor": "Man Shrugging: Medium-Light Skin Tone"
	},
	"1F937-1F3FD": {
		"key": "1F937-1F3FD",
		"value": "🤷🏽",
		"descriptor": "Person Shrugging: Medium Skin Tone"
	},
	"1F937-1F3FD-200D-2640-FE0F": {
		"key": "1F937-1F3FD-200D-2640-FE0F",
		"value": "🤷🏽‍♀️",
		"descriptor": "Woman Shrugging: Medium Skin Tone"
	},
	"1F937-1F3FD-200D-2642-FE0F": {
		"key": "1F937-1F3FD-200D-2642-FE0F",
		"value": "🤷🏽‍♂️",
		"descriptor": "Man Shrugging: Medium Skin Tone"
	},
	"1F937-1F3FE": {
		"key": "1F937-1F3FE",
		"value": "🤷🏾",
		"descriptor": "Person Shrugging: Medium-Dark Skin Tone"
	},
	"1F937-1F3FE-200D-2640-FE0F": {
		"key": "1F937-1F3FE-200D-2640-FE0F",
		"value": "🤷🏾‍♀️",
		"descriptor": "Woman Shrugging: Medium-Dark Skin Tone"
	},
	"1F937-1F3FE-200D-2642-FE0F": {
		"key": "1F937-1F3FE-200D-2642-FE0F",
		"value": "🤷🏾‍♂️",
		"descriptor": "Man Shrugging: Medium-Dark Skin Tone"
	},
	"1F937-1F3FF": {
		"key": "1F937-1F3FF",
		"value": "🤷🏿",
		"descriptor": "Person Shrugging: Dark Skin Tone"
	},
	"1F937-1F3FF-200D-2640-FE0F": {
		"key": "1F937-1F3FF-200D-2640-FE0F",
		"value": "🤷🏿‍♀️",
		"descriptor": "Woman Shrugging: Dark Skin Tone"
	},
	"1F937-1F3FF-200D-2642-FE0F": {
		"key": "1F937-1F3FF-200D-2642-FE0F",
		"value": "🤷🏿‍♂️",
		"descriptor": "Man Shrugging: Dark Skin Tone"
	},
	"1F937-200D-2640-FE0F": {
		"key": "1F937-200D-2640-FE0F",
		"value": "🤷‍♀️",
		"descriptor": "Woman Shrugging"
	},
	"1F937-200D-2642-FE0F": {
		"key": "1F937-200D-2642-FE0F",
		"value": "🤷‍♂️",
		"descriptor": "Man Shrugging"
	},
	"1F938": {
		"key": "1F938",
		"value": "🤸",
		"descriptor": "Person Cartwheeling"
	},
	"1F938-1F3FB": {
		"key": "1F938-1F3FB",
		"value": "🤸🏻",
		"descriptor": "Person Cartwheeling: Light Skin Tone"
	},
	"1F938-1F3FB-200D-2640-FE0F": {
		"key": "1F938-1F3FB-200D-2640-FE0F",
		"value": "🤸🏻‍♀️",
		"descriptor": "Woman Cartwheeling: Light Skin Tone"
	},
	"1F938-1F3FB-200D-2642-FE0F": {
		"key": "1F938-1F3FB-200D-2642-FE0F",
		"value": "🤸🏻‍♂️",
		"descriptor": "Man Cartwheeling: Light Skin Tone"
	},
	"1F938-1F3FC": {
		"key": "1F938-1F3FC",
		"value": "🤸🏼",
		"descriptor": "Person Cartwheeling: Medium-Light Skin Tone"
	},
	"1F938-1F3FC-200D-2640-FE0F": {
		"key": "1F938-1F3FC-200D-2640-FE0F",
		"value": "🤸🏼‍♀️",
		"descriptor": "Woman Cartwheeling: Medium-Light Skin Tone"
	},
	"1F938-1F3FC-200D-2642-FE0F": {
		"key": "1F938-1F3FC-200D-2642-FE0F",
		"value": "🤸🏼‍♂️",
		"descriptor": "Man Cartwheeling: Medium-Light Skin Tone"
	},
	"1F938-1F3FD": {
		"key": "1F938-1F3FD",
		"value": "🤸🏽",
		"descriptor": "Person Cartwheeling: Medium Skin Tone"
	},
	"1F938-1F3FD-200D-2640-FE0F": {
		"key": "1F938-1F3FD-200D-2640-FE0F",
		"value": "🤸🏽‍♀️",
		"descriptor": "Woman Cartwheeling: Medium Skin Tone"
	},
	"1F938-1F3FD-200D-2642-FE0F": {
		"key": "1F938-1F3FD-200D-2642-FE0F",
		"value": "🤸🏽‍♂️",
		"descriptor": "Man Cartwheeling: Medium Skin Tone"
	},
	"1F938-1F3FE": {
		"key": "1F938-1F3FE",
		"value": "🤸🏾",
		"descriptor": "Person Cartwheeling: Medium-Dark Skin Tone"
	},
	"1F938-1F3FE-200D-2640-FE0F": {
		"key": "1F938-1F3FE-200D-2640-FE0F",
		"value": "🤸🏾‍♀️",
		"descriptor": "Woman Cartwheeling: Medium-Dark Skin Tone"
	},
	"1F938-1F3FE-200D-2642-FE0F": {
		"key": "1F938-1F3FE-200D-2642-FE0F",
		"value": "🤸🏾‍♂️",
		"descriptor": "Man Cartwheeling: Medium-Dark Skin Tone"
	},
	"1F938-1F3FF": {
		"key": "1F938-1F3FF",
		"value": "🤸🏿",
		"descriptor": "Person Cartwheeling: Dark Skin Tone"
	},
	"1F938-1F3FF-200D-2640-FE0F": {
		"key": "1F938-1F3FF-200D-2640-FE0F",
		"value": "🤸🏿‍♀️",
		"descriptor": "Woman Cartwheeling: Dark Skin Tone"
	},
	"1F938-1F3FF-200D-2642-FE0F": {
		"key": "1F938-1F3FF-200D-2642-FE0F",
		"value": "🤸🏿‍♂️",
		"descriptor": "Man Cartwheeling: Dark Skin Tone"
	},
	"1F938-200D-2640-FE0F": {
		"key": "1F938-200D-2640-FE0F",
		"value": "🤸‍♀️",
		"descriptor": "Woman Cartwheeling"
	},
	"1F938-200D-2642-FE0F": {
		"key": "1F938-200D-2642-FE0F",
		"value": "🤸‍♂️",
		"descriptor": "Man Cartwheeling"
	},
	"1F939": {
		"key": "1F939",
		"value": "🤹",
		"descriptor": "Person Juggling"
	},
	"1F939-1F3FB": {
		"key": "1F939-1F3FB",
		"value": "🤹🏻",
		"descriptor": "Person Juggling: Light Skin Tone"
	},
	"1F939-1F3FB-200D-2640-FE0F": {
		"key": "1F939-1F3FB-200D-2640-FE0F",
		"value": "🤹🏻‍♀️",
		"descriptor": "Woman Juggling: Light Skin Tone"
	},
	"1F939-1F3FB-200D-2642-FE0F": {
		"key": "1F939-1F3FB-200D-2642-FE0F",
		"value": "🤹🏻‍♂️",
		"descriptor": "Man Juggling: Light Skin Tone"
	},
	"1F939-1F3FC": {
		"key": "1F939-1F3FC",
		"value": "🤹🏼",
		"descriptor": "Person Juggling: Medium-Light Skin Tone"
	},
	"1F939-1F3FC-200D-2640-FE0F": {
		"key": "1F939-1F3FC-200D-2640-FE0F",
		"value": "🤹🏼‍♀️",
		"descriptor": "Woman Juggling: Medium-Light Skin Tone"
	},
	"1F939-1F3FC-200D-2642-FE0F": {
		"key": "1F939-1F3FC-200D-2642-FE0F",
		"value": "🤹🏼‍♂️",
		"descriptor": "Man Juggling: Medium-Light Skin Tone"
	},
	"1F939-1F3FD": {
		"key": "1F939-1F3FD",
		"value": "🤹🏽",
		"descriptor": "Person Juggling: Medium Skin Tone"
	},
	"1F939-1F3FD-200D-2640-FE0F": {
		"key": "1F939-1F3FD-200D-2640-FE0F",
		"value": "🤹🏽‍♀️",
		"descriptor": "Woman Juggling: Medium Skin Tone"
	},
	"1F939-1F3FD-200D-2642-FE0F": {
		"key": "1F939-1F3FD-200D-2642-FE0F",
		"value": "🤹🏽‍♂️",
		"descriptor": "Man Juggling: Medium Skin Tone"
	},
	"1F939-1F3FE": {
		"key": "1F939-1F3FE",
		"value": "🤹🏾",
		"descriptor": "Person Juggling: Medium-Dark Skin Tone"
	},
	"1F939-1F3FE-200D-2640-FE0F": {
		"key": "1F939-1F3FE-200D-2640-FE0F",
		"value": "🤹🏾‍♀️",
		"descriptor": "Woman Juggling: Medium-Dark Skin Tone"
	},
	"1F939-1F3FE-200D-2642-FE0F": {
		"key": "1F939-1F3FE-200D-2642-FE0F",
		"value": "🤹🏾‍♂️",
		"descriptor": "Man Juggling: Medium-Dark Skin Tone"
	},
	"1F939-1F3FF": {
		"key": "1F939-1F3FF",
		"value": "🤹🏿",
		"descriptor": "Person Juggling: Dark Skin Tone"
	},
	"1F939-1F3FF-200D-2640-FE0F": {
		"key": "1F939-1F3FF-200D-2640-FE0F",
		"value": "🤹🏿‍♀️",
		"descriptor": "Woman Juggling: Dark Skin Tone"
	},
	"1F939-1F3FF-200D-2642-FE0F": {
		"key": "1F939-1F3FF-200D-2642-FE0F",
		"value": "🤹🏿‍♂️",
		"descriptor": "Man Juggling: Dark Skin Tone"
	},
	"1F939-200D-2640-FE0F": {
		"key": "1F939-200D-2640-FE0F",
		"value": "🤹‍♀️",
		"descriptor": "Woman Juggling"
	},
	"1F939-200D-2642-FE0F": {
		"key": "1F939-200D-2642-FE0F",
		"value": "🤹‍♂️",
		"descriptor": "Man Juggling"
	},
	"1F93A": {
		"key": "1F93A",
		"value": "🤺",
		"descriptor": "Person Fencing"
	},
	"1F93C": {
		"key": "1F93C",
		"value": "🤼",
		"descriptor": "People Wrestling"
	},
	"1F93C-200D-2640-FE0F": {
		"key": "1F93C-200D-2640-FE0F",
		"value": "🤼‍♀️",
		"descriptor": "Women Wrestling"
	},
	"1F93C-200D-2642-FE0F": {
		"key": "1F93C-200D-2642-FE0F",
		"value": "🤼‍♂️",
		"descriptor": "Men Wrestling"
	},
	"1F93D": {
		"key": "1F93D",
		"value": "🤽",
		"descriptor": "Person Playing Water Polo"
	},
	"1F93D-1F3FB": {
		"key": "1F93D-1F3FB",
		"value": "🤽🏻",
		"descriptor": "Person Playing Water Polo: Light Skin Tone"
	},
	"1F93D-1F3FB-200D-2640-FE0F": {
		"key": "1F93D-1F3FB-200D-2640-FE0F",
		"value": "🤽🏻‍♀️",
		"descriptor": "Woman Playing Water Polo: Light Skin Tone"
	},
	"1F93D-1F3FB-200D-2642-FE0F": {
		"key": "1F93D-1F3FB-200D-2642-FE0F",
		"value": "🤽🏻‍♂️",
		"descriptor": "Man Playing Water Polo: Light Skin Tone"
	},
	"1F93D-1F3FC": {
		"key": "1F93D-1F3FC",
		"value": "🤽🏼",
		"descriptor": "Person Playing Water Polo: Medium-Light Skin Tone"
	},
	"1F93D-1F3FC-200D-2640-FE0F": {
		"key": "1F93D-1F3FC-200D-2640-FE0F",
		"value": "🤽🏼‍♀️",
		"descriptor": "Woman Playing Water Polo: Medium-Light Skin Tone"
	},
	"1F93D-1F3FC-200D-2642-FE0F": {
		"key": "1F93D-1F3FC-200D-2642-FE0F",
		"value": "🤽🏼‍♂️",
		"descriptor": "Man Playing Water Polo: Medium-Light Skin Tone"
	},
	"1F93D-1F3FD": {
		"key": "1F93D-1F3FD",
		"value": "🤽🏽",
		"descriptor": "Person Playing Water Polo: Medium Skin Tone"
	},
	"1F93D-1F3FD-200D-2640-FE0F": {
		"key": "1F93D-1F3FD-200D-2640-FE0F",
		"value": "🤽🏽‍♀️",
		"descriptor": "Woman Playing Water Polo: Medium Skin Tone"
	},
	"1F93D-1F3FD-200D-2642-FE0F": {
		"key": "1F93D-1F3FD-200D-2642-FE0F",
		"value": "🤽🏽‍♂️",
		"descriptor": "Man Playing Water Polo: Medium Skin Tone"
	},
	"1F93D-1F3FE": {
		"key": "1F93D-1F3FE",
		"value": "🤽🏾",
		"descriptor": "Person Playing Water Polo: Medium-Dark Skin Tone"
	},
	"1F93D-1F3FE-200D-2640-FE0F": {
		"key": "1F93D-1F3FE-200D-2640-FE0F",
		"value": "🤽🏾‍♀️",
		"descriptor": "Woman Playing Water Polo: Medium-Dark Skin Tone"
	},
	"1F93D-1F3FE-200D-2642-FE0F": {
		"key": "1F93D-1F3FE-200D-2642-FE0F",
		"value": "🤽🏾‍♂️",
		"descriptor": "Man Playing Water Polo: Medium-Dark Skin Tone"
	},
	"1F93D-1F3FF": {
		"key": "1F93D-1F3FF",
		"value": "🤽🏿",
		"descriptor": "Person Playing Water Polo: Dark Skin Tone"
	},
	"1F93D-1F3FF-200D-2640-FE0F": {
		"key": "1F93D-1F3FF-200D-2640-FE0F",
		"value": "🤽🏿‍♀️",
		"descriptor": "Woman Playing Water Polo: Dark Skin Tone"
	},
	"1F93D-1F3FF-200D-2642-FE0F": {
		"key": "1F93D-1F3FF-200D-2642-FE0F",
		"value": "🤽🏿‍♂️",
		"descriptor": "Man Playing Water Polo: Dark Skin Tone"
	},
	"1F93D-200D-2640-FE0F": {
		"key": "1F93D-200D-2640-FE0F",
		"value": "🤽‍♀️",
		"descriptor": "Woman Playing Water Polo"
	},
	"1F93D-200D-2642-FE0F": {
		"key": "1F93D-200D-2642-FE0F",
		"value": "🤽‍♂️",
		"descriptor": "Man Playing Water Polo"
	},
	"1F93E": {
		"key": "1F93E",
		"value": "🤾",
		"descriptor": "Person Playing Handball"
	},
	"1F93E-1F3FB": {
		"key": "1F93E-1F3FB",
		"value": "🤾🏻",
		"descriptor": "Person Playing Handball: Light Skin Tone"
	},
	"1F93E-1F3FB-200D-2640-FE0F": {
		"key": "1F93E-1F3FB-200D-2640-FE0F",
		"value": "🤾🏻‍♀️",
		"descriptor": "Woman Playing Handball: Light Skin Tone"
	},
	"1F93E-1F3FB-200D-2642-FE0F": {
		"key": "1F93E-1F3FB-200D-2642-FE0F",
		"value": "🤾🏻‍♂️",
		"descriptor": "Man Playing Handball: Light Skin Tone"
	},
	"1F93E-1F3FC": {
		"key": "1F93E-1F3FC",
		"value": "🤾🏼",
		"descriptor": "Person Playing Handball: Medium-Light Skin Tone"
	},
	"1F93E-1F3FC-200D-2640-FE0F": {
		"key": "1F93E-1F3FC-200D-2640-FE0F",
		"value": "🤾🏼‍♀️",
		"descriptor": "Woman Playing Handball: Medium-Light Skin Tone"
	},
	"1F93E-1F3FC-200D-2642-FE0F": {
		"key": "1F93E-1F3FC-200D-2642-FE0F",
		"value": "🤾🏼‍♂️",
		"descriptor": "Man Playing Handball: Medium-Light Skin Tone"
	},
	"1F93E-1F3FD": {
		"key": "1F93E-1F3FD",
		"value": "🤾🏽",
		"descriptor": "Person Playing Handball: Medium Skin Tone"
	},
	"1F93E-1F3FD-200D-2640-FE0F": {
		"key": "1F93E-1F3FD-200D-2640-FE0F",
		"value": "🤾🏽‍♀️",
		"descriptor": "Woman Playing Handball: Medium Skin Tone"
	},
	"1F93E-1F3FD-200D-2642-FE0F": {
		"key": "1F93E-1F3FD-200D-2642-FE0F",
		"value": "🤾🏽‍♂️",
		"descriptor": "Man Playing Handball: Medium Skin Tone"
	},
	"1F93E-1F3FE": {
		"key": "1F93E-1F3FE",
		"value": "🤾🏾",
		"descriptor": "Person Playing Handball: Medium-Dark Skin Tone"
	},
	"1F93E-1F3FE-200D-2640-FE0F": {
		"key": "1F93E-1F3FE-200D-2640-FE0F",
		"value": "🤾🏾‍♀️",
		"descriptor": "Woman Playing Handball: Medium-Dark Skin Tone"
	},
	"1F93E-1F3FE-200D-2642-FE0F": {
		"key": "1F93E-1F3FE-200D-2642-FE0F",
		"value": "🤾🏾‍♂️",
		"descriptor": "Man Playing Handball: Medium-Dark Skin Tone"
	},
	"1F93E-1F3FF": {
		"key": "1F93E-1F3FF",
		"value": "🤾🏿",
		"descriptor": "Person Playing Handball: Dark Skin Tone"
	},
	"1F93E-1F3FF-200D-2640-FE0F": {
		"key": "1F93E-1F3FF-200D-2640-FE0F",
		"value": "🤾🏿‍♀️",
		"descriptor": "Woman Playing Handball: Dark Skin Tone"
	},
	"1F93E-1F3FF-200D-2642-FE0F": {
		"key": "1F93E-1F3FF-200D-2642-FE0F",
		"value": "🤾🏿‍♂️",
		"descriptor": "Man Playing Handball: Dark Skin Tone"
	},
	"1F93E-200D-2640-FE0F": {
		"key": "1F93E-200D-2640-FE0F",
		"value": "🤾‍♀️",
		"descriptor": "Woman Playing Handball"
	},
	"1F93E-200D-2642-FE0F": {
		"key": "1F93E-200D-2642-FE0F",
		"value": "🤾‍♂️",
		"descriptor": "Man Playing Handball"
	},
	"1F93F": {
		"key": "1F93F",
		"value": "🤿",
		"descriptor": "Diving Mask"
	},
	"1F940": {
		"key": "1F940",
		"value": "🥀",
		"descriptor": "Wilted Flower"
	},
	"1F941": {
		"key": "1F941",
		"value": "🥁",
		"descriptor": "Drum"
	},
	"1F942": {
		"key": "1F942",
		"value": "🥂",
		"descriptor": "Clinking Glasses"
	},
	"1F943": {
		"key": "1F943",
		"value": "🥃",
		"descriptor": "Tumbler Glass"
	},
	"1F944": {
		"key": "1F944",
		"value": "🥄",
		"descriptor": "Spoon"
	},
	"1F945": {
		"key": "1F945",
		"value": "🥅",
		"descriptor": "Goal Net"
	},
	"1F947": {
		"key": "1F947",
		"value": "🥇",
		"descriptor": "1st Place Medal"
	},
	"1F948": {
		"key": "1F948",
		"value": "🥈",
		"descriptor": "2nd Place Medal"
	},
	"1F949": {
		"key": "1F949",
		"value": "🥉",
		"descriptor": "3rd Place Medal"
	},
	"1F94A": {
		"key": "1F94A",
		"value": "🥊",
		"descriptor": "Boxing Glove"
	},
	"1F94B": {
		"key": "1F94B",
		"value": "🥋",
		"descriptor": "Martial Arts Uniform"
	},
	"1F94C": {
		"key": "1F94C",
		"value": "🥌",
		"descriptor": "Curling Stone"
	},
	"1F94D": {
		"key": "1F94D",
		"value": "🥍",
		"descriptor": "Lacrosse"
	},
	"1F94E": {
		"key": "1F94E",
		"value": "🥎",
		"descriptor": "Softball"
	},
	"1F94F": {
		"key": "1F94F",
		"value": "🥏",
		"descriptor": "Flying Disc"
	},
	"1F950": {
		"key": "1F950",
		"value": "🥐",
		"descriptor": "Croissant"
	},
	"1F951": {
		"key": "1F951",
		"value": "🥑",
		"descriptor": "Avocado"
	},
	"1F952": {
		"key": "1F952",
		"value": "🥒",
		"descriptor": "Cucumber"
	},
	"1F953": {
		"key": "1F953",
		"value": "🥓",
		"descriptor": "Bacon"
	},
	"1F954": {
		"key": "1F954",
		"value": "🥔",
		"descriptor": "Potato"
	},
	"1F955": {
		"key": "1F955",
		"value": "🥕",
		"descriptor": "Carrot"
	},
	"1F956": {
		"key": "1F956",
		"value": "🥖",
		"descriptor": "Baguette Bread"
	},
	"1F957": {
		"key": "1F957",
		"value": "🥗",
		"descriptor": "Green Salad"
	},
	"1F958": {
		"key": "1F958",
		"value": "🥘",
		"descriptor": "Shallow Pan of Food"
	},
	"1F959": {
		"key": "1F959",
		"value": "🥙",
		"descriptor": "Stuffed Flatbread"
	},
	"1F95A": {
		"key": "1F95A",
		"value": "🥚",
		"descriptor": "Egg"
	},
	"1F95B": {
		"key": "1F95B",
		"value": "🥛",
		"descriptor": "Glass of Milk"
	},
	"1F95C": {
		"key": "1F95C",
		"value": "🥜",
		"descriptor": "Peanuts"
	},
	"1F95D": {
		"key": "1F95D",
		"value": "🥝",
		"descriptor": "Kiwi Fruit"
	},
	"1F95E": {
		"key": "1F95E",
		"value": "🥞",
		"descriptor": "Pancakes"
	},
	"1F95F": {
		"key": "1F95F",
		"value": "🥟",
		"descriptor": "Dumpling"
	},
	"1F960": {
		"key": "1F960",
		"value": "🥠",
		"descriptor": "Fortune Cookie"
	},
	"1F961": {
		"key": "1F961",
		"value": "🥡",
		"descriptor": "Takeout Box"
	},
	"1F962": {
		"key": "1F962",
		"value": "🥢",
		"descriptor": "Chopsticks"
	},
	"1F963": {
		"key": "1F963",
		"value": "🥣",
		"descriptor": "Bowl With Spoon"
	},
	"1F964": {
		"key": "1F964",
		"value": "🥤",
		"descriptor": "Cup With Straw"
	},
	"1F965": {
		"key": "1F965",
		"value": "🥥",
		"descriptor": "Coconut"
	},
	"1F966": {
		"key": "1F966",
		"value": "🥦",
		"descriptor": "Broccoli"
	},
	"1F967": {
		"key": "1F967",
		"value": "🥧",
		"descriptor": "Pie"
	},
	"1F968": {
		"key": "1F968",
		"value": "🥨",
		"descriptor": "Pretzel"
	},
	"1F969": {
		"key": "1F969",
		"value": "🥩",
		"descriptor": "Cut of Meat"
	},
	"1F96A": {
		"key": "1F96A",
		"value": "🥪",
		"descriptor": "Sandwich"
	},
	"1F96B": {
		"key": "1F96B",
		"value": "🥫",
		"descriptor": "Canned Food"
	},
	"1F96C": {
		"key": "1F96C",
		"value": "🥬",
		"descriptor": "Leafy Green"
	},
	"1F96D": {
		"key": "1F96D",
		"value": "🥭",
		"descriptor": "Mango"
	},
	"1F96E": {
		"key": "1F96E",
		"value": "🥮",
		"descriptor": "Moon Cake"
	},
	"1F96F": {
		"key": "1F96F",
		"value": "🥯",
		"descriptor": "Bagel"
	},
	"1F970": {
		"key": "1F970",
		"value": "🥰",
		"descriptor": "Smiling Face With Hearts"
	},
	"1F971": {
		"key": "1F971",
		"value": "🥱",
		"descriptor": "Yawning Face"
	},
	"1F973": {
		"key": "1F973",
		"value": "🥳",
		"descriptor": "Partying Face"
	},
	"1F974": {
		"key": "1F974",
		"value": "🥴",
		"descriptor": "Woozy Face"
	},
	"1F975": {
		"key": "1F975",
		"value": "🥵",
		"descriptor": "Hot Face"
	},
	"1F976": {
		"key": "1F976",
		"value": "🥶",
		"descriptor": "Cold Face"
	},
	"1F97A": {
		"key": "1F97A",
		"value": "🥺",
		"descriptor": "Pleading Face"
	},
	"1F97B": {
		"key": "1F97B",
		"value": "🥻",
		"descriptor": "Sari"
	},
	"1F97C": {
		"key": "1F97C",
		"value": "🥼",
		"descriptor": "Lab Coat"
	},
	"1F97D": {
		"key": "1F97D",
		"value": "🥽",
		"descriptor": "Goggles"
	},
	"1F97E": {
		"key": "1F97E",
		"value": "🥾",
		"descriptor": "Hiking Boot"
	},
	"1F97F": {
		"key": "1F97F",
		"value": "🥿",
		"descriptor": "Flat Shoe"
	},
	"1F980": {
		"key": "1F980",
		"value": "🦀",
		"descriptor": "Crab"
	},
	"1F981": {
		"key": "1F981",
		"value": "🦁",
		"descriptor": "Lion"
	},
	"1F982": {
		"key": "1F982",
		"value": "🦂",
		"descriptor": "Scorpion"
	},
	"1F983": {
		"key": "1F983",
		"value": "🦃",
		"descriptor": "Turkey"
	},
	"1F984": {
		"key": "1F984",
		"value": "🦄",
		"descriptor": "Unicorn"
	},
	"1F985": {
		"key": "1F985",
		"value": "🦅",
		"descriptor": "Eagle"
	},
	"1F986": {
		"key": "1F986",
		"value": "🦆",
		"descriptor": "Duck"
	},
	"1F987": {
		"key": "1F987",
		"value": "🦇",
		"descriptor": "Bat"
	},
	"1F988": {
		"key": "1F988",
		"value": "🦈",
		"descriptor": "Shark"
	},
	"1F989": {
		"key": "1F989",
		"value": "🦉",
		"descriptor": "Owl"
	},
	"1F98A": {
		"key": "1F98A",
		"value": "🦊",
		"descriptor": "Fox"
	},
	"1F98B": {
		"key": "1F98B",
		"value": "🦋",
		"descriptor": "Butterfly"
	},
	"1F98C": {
		"key": "1F98C",
		"value": "🦌",
		"descriptor": "Deer"
	},
	"1F98D": {
		"key": "1F98D",
		"value": "🦍",
		"descriptor": "Gorilla"
	},
	"1F98E": {
		"key": "1F98E",
		"value": "🦎",
		"descriptor": "Lizard"
	},
	"1F98F": {
		"key": "1F98F",
		"value": "🦏",
		"descriptor": "Rhinoceros"
	},
	"1F990": {
		"key": "1F990",
		"value": "🦐",
		"descriptor": "Shrimp"
	},
	"1F991": {
		"key": "1F991",
		"value": "🦑",
		"descriptor": "Squid"
	},
	"1F992": {
		"key": "1F992",
		"value": "🦒",
		"descriptor": "Giraffe"
	},
	"1F993": {
		"key": "1F993",
		"value": "🦓",
		"descriptor": "Zebra"
	},
	"1F994": {
		"key": "1F994",
		"value": "🦔",
		"descriptor": "Hedgehog"
	},
	"1F995": {
		"key": "1F995",
		"value": "🦕",
		"descriptor": "Sauropod"
	},
	"1F996": {
		"key": "1F996",
		"value": "🦖",
		"descriptor": "T-Rex"
	},
	"1F997": {
		"key": "1F997",
		"value": "🦗",
		"descriptor": "Cricket"
	},
	"1F998": {
		"key": "1F998",
		"value": "🦘",
		"descriptor": "Kangaroo"
	},
	"1F999": {
		"key": "1F999",
		"value": "🦙",
		"descriptor": "Llama"
	},
	"1F99A": {
		"key": "1F99A",
		"value": "🦚",
		"descriptor": "Peacock"
	},
	"1F99B": {
		"key": "1F99B",
		"value": "🦛",
		"descriptor": "Hippopotamus"
	},
	"1F99C": {
		"key": "1F99C",
		"value": "🦜",
		"descriptor": "Parrot"
	},
	"1F99D": {
		"key": "1F99D",
		"value": "🦝",
		"descriptor": "Raccoon"
	},
	"1F99E": {
		"key": "1F99E",
		"value": "🦞",
		"descriptor": "Lobster"
	},
	"1F99F": {
		"key": "1F99F",
		"value": "🦟",
		"descriptor": "Mosquito"
	},
	"1F9A0": {
		"key": "1F9A0",
		"value": "🦠",
		"descriptor": "Microbe"
	},
	"1F9A1": {
		"key": "1F9A1",
		"value": "🦡",
		"descriptor": "Badger"
	},
	"1F9A2": {
		"key": "1F9A2",
		"value": "🦢",
		"descriptor": "Swan"
	},
	"1F9A5": {
		"key": "1F9A5",
		"value": "🦥",
		"descriptor": "Sloth"
	},
	"1F9A6": {
		"key": "1F9A6",
		"value": "🦦",
		"descriptor": "Otter"
	},
	"1F9A7": {
		"key": "1F9A7",
		"value": "🦧",
		"descriptor": "Orangutan"
	},
	"1F9A8": {
		"key": "1F9A8",
		"value": "🦨",
		"descriptor": "Skunk"
	},
	"1F9A9": {
		"key": "1F9A9",
		"value": "🦩",
		"descriptor": "Flamingo"
	},
	"1F9AA": {
		"key": "1F9AA",
		"value": "🦪",
		"descriptor": "Oyster"
	},
	"1F9AE": {
		"key": "1F9AE",
		"value": "🦮",
		"descriptor": "Guide Dog"
	},
	"1F9AF": {
		"key": "1F9AF",
		"value": "🦯",
		"descriptor": "Probing Cane"
	},
	"1F9B0": {
		"key": "1F9B0",
		"value": "🦰",
		"descriptor": "Red Hair"
	},
	"1F9B1": {
		"key": "1F9B1",
		"value": "🦱",
		"descriptor": "Curly Hair"
	},
	"1F9B2": {
		"key": "1F9B2",
		"value": "🦲",
		"descriptor": "Bald"
	},
	"1F9B3": {
		"key": "1F9B3",
		"value": "🦳",
		"descriptor": "White Hair"
	},
	"1F9B4": {
		"key": "1F9B4",
		"value": "🦴",
		"descriptor": "Bone"
	},
	"1F9B5": {
		"key": "1F9B5",
		"value": "🦵",
		"descriptor": "Leg"
	},
	"1F9B5-1F3FB": {
		"key": "1F9B5-1F3FB",
		"value": "🦵🏻",
		"descriptor": "Leg: Light Skin Tone"
	},
	"1F9B5-1F3FC": {
		"key": "1F9B5-1F3FC",
		"value": "🦵🏼",
		"descriptor": "Leg: Medium-Light Skin Tone"
	},
	"1F9B5-1F3FD": {
		"key": "1F9B5-1F3FD",
		"value": "🦵🏽",
		"descriptor": "Leg: Medium Skin Tone"
	},
	"1F9B5-1F3FE": {
		"key": "1F9B5-1F3FE",
		"value": "🦵🏾",
		"descriptor": "Leg: Medium-Dark Skin Tone"
	},
	"1F9B5-1F3FF": {
		"key": "1F9B5-1F3FF",
		"value": "🦵🏿",
		"descriptor": "Leg: Dark Skin Tone"
	},
	"1F9B6": {
		"key": "1F9B6",
		"value": "🦶",
		"descriptor": "Foot"
	},
	"1F9B6-1F3FB": {
		"key": "1F9B6-1F3FB",
		"value": "🦶🏻",
		"descriptor": "Foot: Light Skin Tone"
	},
	"1F9B6-1F3FC": {
		"key": "1F9B6-1F3FC",
		"value": "🦶🏼",
		"descriptor": "Foot: Medium-Light Skin Tone"
	},
	"1F9B6-1F3FD": {
		"key": "1F9B6-1F3FD",
		"value": "🦶🏽",
		"descriptor": "Foot: Medium Skin Tone"
	},
	"1F9B6-1F3FE": {
		"key": "1F9B6-1F3FE",
		"value": "🦶🏾",
		"descriptor": "Foot: Medium-Dark Skin Tone"
	},
	"1F9B6-1F3FF": {
		"key": "1F9B6-1F3FF",
		"value": "🦶🏿",
		"descriptor": "Foot: Dark Skin Tone"
	},
	"1F9B7": {
		"key": "1F9B7",
		"value": "🦷",
		"descriptor": "Tooth"
	},
	"1F9B8": {
		"key": "1F9B8",
		"value": "🦸",
		"descriptor": "Superhero"
	},
	"1F9B8-1F3FB": {
		"key": "1F9B8-1F3FB",
		"value": "🦸🏻",
		"descriptor": "Superhero: Light Skin Tone"
	},
	"1F9B8-1F3FB-200D-2640-FE0F": {
		"key": "1F9B8-1F3FB-200D-2640-FE0F",
		"value": "🦸🏻‍♀️",
		"descriptor": "Woman Superhero: Light Skin Tone"
	},
	"1F9B8-1F3FB-200D-2642-FE0F": {
		"key": "1F9B8-1F3FB-200D-2642-FE0F",
		"value": "🦸🏻‍♂️",
		"descriptor": "Man Superhero: Light Skin Tone"
	},
	"1F9B8-1F3FC": {
		"key": "1F9B8-1F3FC",
		"value": "🦸🏼",
		"descriptor": "Superhero: Medium-Light Skin Tone"
	},
	"1F9B8-1F3FC-200D-2640-FE0F": {
		"key": "1F9B8-1F3FC-200D-2640-FE0F",
		"value": "🦸🏼‍♀️",
		"descriptor": "Woman Superhero: Medium-Light Skin Tone"
	},
	"1F9B8-1F3FC-200D-2642-FE0F": {
		"key": "1F9B8-1F3FC-200D-2642-FE0F",
		"value": "🦸🏼‍♂️",
		"descriptor": "Man Superhero: Medium-Light Skin Tone"
	},
	"1F9B8-1F3FD": {
		"key": "1F9B8-1F3FD",
		"value": "🦸🏽",
		"descriptor": "Superhero: Medium Skin Tone"
	},
	"1F9B8-1F3FD-200D-2640-FE0F": {
		"key": "1F9B8-1F3FD-200D-2640-FE0F",
		"value": "🦸🏽‍♀️",
		"descriptor": "Woman Superhero: Medium Skin Tone"
	},
	"1F9B8-1F3FD-200D-2642-FE0F": {
		"key": "1F9B8-1F3FD-200D-2642-FE0F",
		"value": "🦸🏽‍♂️",
		"descriptor": "Man Superhero: Medium Skin Tone"
	},
	"1F9B8-1F3FE": {
		"key": "1F9B8-1F3FE",
		"value": "🦸🏾",
		"descriptor": "Superhero: Medium-Dark Skin Tone"
	},
	"1F9B8-1F3FE-200D-2640-FE0F": {
		"key": "1F9B8-1F3FE-200D-2640-FE0F",
		"value": "🦸🏾‍♀️",
		"descriptor": "Woman Superhero: Medium-Dark Skin Tone"
	},
	"1F9B8-1F3FE-200D-2642-FE0F": {
		"key": "1F9B8-1F3FE-200D-2642-FE0F",
		"value": "🦸🏾‍♂️",
		"descriptor": "Man Superhero: Medium-Dark Skin Tone"
	},
	"1F9B8-1F3FF": {
		"key": "1F9B8-1F3FF",
		"value": "🦸🏿",
		"descriptor": "Superhero: Dark Skin Tone"
	},
	"1F9B8-1F3FF-200D-2640-FE0F": {
		"key": "1F9B8-1F3FF-200D-2640-FE0F",
		"value": "🦸🏿‍♀️",
		"descriptor": "Woman Superhero: Dark Skin Tone"
	},
	"1F9B8-1F3FF-200D-2642-FE0F": {
		"key": "1F9B8-1F3FF-200D-2642-FE0F",
		"value": "🦸🏿‍♂️",
		"descriptor": "Man Superhero: Dark Skin Tone"
	},
	"1F9B8-200D-2640-FE0F": {
		"key": "1F9B8-200D-2640-FE0F",
		"value": "🦸‍♀️",
		"descriptor": "Woman Superhero"
	},
	"1F9B8-200D-2642-FE0F": {
		"key": "1F9B8-200D-2642-FE0F",
		"value": "🦸‍♂️",
		"descriptor": "Man Superhero"
	},
	"1F9B9": {
		"key": "1F9B9",
		"value": "🦹",
		"descriptor": "Supervillain"
	},
	"1F9B9-1F3FB": {
		"key": "1F9B9-1F3FB",
		"value": "🦹🏻",
		"descriptor": "Supervillain: Light Skin Tone"
	},
	"1F9B9-1F3FB-200D-2640-FE0F": {
		"key": "1F9B9-1F3FB-200D-2640-FE0F",
		"value": "🦹🏻‍♀️",
		"descriptor": "Woman Supervillain: Light Skin Tone"
	},
	"1F9B9-1F3FB-200D-2642-FE0F": {
		"key": "1F9B9-1F3FB-200D-2642-FE0F",
		"value": "🦹🏻‍♂️",
		"descriptor": "Man Supervillain: Light Skin Tone"
	},
	"1F9B9-1F3FC": {
		"key": "1F9B9-1F3FC",
		"value": "🦹🏼",
		"descriptor": "Supervillain: Medium-Light Skin Tone"
	},
	"1F9B9-1F3FC-200D-2640-FE0F": {
		"key": "1F9B9-1F3FC-200D-2640-FE0F",
		"value": "🦹🏼‍♀️",
		"descriptor": "Woman Supervillain: Medium-Light Skin Tone"
	},
	"1F9B9-1F3FC-200D-2642-FE0F": {
		"key": "1F9B9-1F3FC-200D-2642-FE0F",
		"value": "🦹🏼‍♂️",
		"descriptor": "Man Supervillain: Medium-Light Skin Tone"
	},
	"1F9B9-1F3FD": {
		"key": "1F9B9-1F3FD",
		"value": "🦹🏽",
		"descriptor": "Supervillain: Medium Skin Tone"
	},
	"1F9B9-1F3FD-200D-2640-FE0F": {
		"key": "1F9B9-1F3FD-200D-2640-FE0F",
		"value": "🦹🏽‍♀️",
		"descriptor": "Woman Supervillain: Medium Skin Tone"
	},
	"1F9B9-1F3FD-200D-2642-FE0F": {
		"key": "1F9B9-1F3FD-200D-2642-FE0F",
		"value": "🦹🏽‍♂️",
		"descriptor": "Man Supervillain: Medium Skin Tone"
	},
	"1F9B9-1F3FE": {
		"key": "1F9B9-1F3FE",
		"value": "🦹🏾",
		"descriptor": "Supervillain: Medium-Dark Skin Tone"
	},
	"1F9B9-1F3FE-200D-2640-FE0F": {
		"key": "1F9B9-1F3FE-200D-2640-FE0F",
		"value": "🦹🏾‍♀️",
		"descriptor": "Woman Supervillain: Medium-Dark Skin Tone"
	},
	"1F9B9-1F3FE-200D-2642-FE0F": {
		"key": "1F9B9-1F3FE-200D-2642-FE0F",
		"value": "🦹🏾‍♂️",
		"descriptor": "Man Supervillain: Medium-Dark Skin Tone"
	},
	"1F9B9-1F3FF": {
		"key": "1F9B9-1F3FF",
		"value": "🦹🏿",
		"descriptor": "Supervillain: Dark Skin Tone"
	},
	"1F9B9-1F3FF-200D-2640-FE0F": {
		"key": "1F9B9-1F3FF-200D-2640-FE0F",
		"value": "🦹🏿‍♀️",
		"descriptor": "Woman Supervillain: Dark Skin Tone"
	},
	"1F9B9-1F3FF-200D-2642-FE0F": {
		"key": "1F9B9-1F3FF-200D-2642-FE0F",
		"value": "🦹🏿‍♂️",
		"descriptor": "Man Supervillain: Dark Skin Tone"
	},
	"1F9B9-200D-2640-FE0F": {
		"key": "1F9B9-200D-2640-FE0F",
		"value": "🦹‍♀️",
		"descriptor": "Woman Supervillain"
	},
	"1F9B9-200D-2642-FE0F": {
		"key": "1F9B9-200D-2642-FE0F",
		"value": "🦹‍♂️",
		"descriptor": "Man Supervillain"
	},
	"1F9BA": {
		"key": "1F9BA",
		"value": "🦺",
		"descriptor": "Safety Vest"
	},
	"1F9BB": {
		"key": "1F9BB",
		"value": "🦻",
		"descriptor": "Ear With Hearing Aid"
	},
	"1F9BB-1F3FB": {
		"key": "1F9BB-1F3FB",
		"value": "🦻🏻",
		"descriptor": "Ear With Hearing Aid: Light Skin Tone"
	},
	"1F9BB-1F3FC": {
		"key": "1F9BB-1F3FC",
		"value": "🦻🏼",
		"descriptor": "Ear With Hearing Aid: Medium-Light Skin Tone"
	},
	"1F9BB-1F3FD": {
		"key": "1F9BB-1F3FD",
		"value": "🦻🏽",
		"descriptor": "Ear With Hearing Aid: Medium Skin Tone"
	},
	"1F9BB-1F3FE": {
		"key": "1F9BB-1F3FE",
		"value": "🦻🏾",
		"descriptor": "Ear With Hearing Aid: Medium-Dark Skin Tone"
	},
	"1F9BB-1F3FF": {
		"key": "1F9BB-1F3FF",
		"value": "🦻🏿",
		"descriptor": "Ear With Hearing Aid: Dark Skin Tone"
	},
	"1F9BC": {
		"key": "1F9BC",
		"value": "🦼",
		"descriptor": "Motorized Wheelchair"
	},
	"1F9BD": {
		"key": "1F9BD",
		"value": "🦽",
		"descriptor": "Manual Wheelchair"
	},
	"1F9BE": {
		"key": "1F9BE",
		"value": "🦾",
		"descriptor": "Mechanical Arm"
	},
	"1F9BF": {
		"key": "1F9BF",
		"value": "🦿",
		"descriptor": "Mechanical Leg"
	},
	"1F9C0": {
		"key": "1F9C0",
		"value": "🧀",
		"descriptor": "Cheese Wedge"
	},
	"1F9C1": {
		"key": "1F9C1",
		"value": "🧁",
		"descriptor": "Cupcake"
	},
	"1F9C2": {
		"key": "1F9C2",
		"value": "🧂",
		"descriptor": "Salt"
	},
	"1F9C3": {
		"key": "1F9C3",
		"value": "🧃",
		"descriptor": "Beverage Box"
	},
	"1F9C4": {
		"key": "1F9C4",
		"value": "🧄",
		"descriptor": "Garlic"
	},
	"1F9C5": {
		"key": "1F9C5",
		"value": "🧅",
		"descriptor": "Onion"
	},
	"1F9C6": {
		"key": "1F9C6",
		"value": "🧆",
		"descriptor": "Falafel"
	},
	"1F9C7": {
		"key": "1F9C7",
		"value": "🧇",
		"descriptor": "Waffle"
	},
	"1F9C8": {
		"key": "1F9C8",
		"value": "🧈",
		"descriptor": "Butter"
	},
	"1F9C9": {
		"key": "1F9C9",
		"value": "🧉",
		"descriptor": "Mate"
	},
	"1F9CA": {
		"key": "1F9CA",
		"value": "🧊",
		"descriptor": "Ice"
	},
	"1F9CD": {
		"key": "1F9CD",
		"value": "🧍",
		"descriptor": "Person Standing"
	},
	"1F9CD-1F3FB": {
		"key": "1F9CD-1F3FB",
		"value": "🧍🏻",
		"descriptor": "Person Standing: Light Skin Tone"
	},
	"1F9CD-1F3FB-200D-2640-FE0F": {
		"key": "1F9CD-1F3FB-200D-2640-FE0F",
		"value": "🧍🏻‍♀️",
		"descriptor": "Woman Standing: Light Skin Tone"
	},
	"1F9CD-1F3FB-200D-2642-FE0F": {
		"key": "1F9CD-1F3FB-200D-2642-FE0F",
		"value": "🧍🏻‍♂️",
		"descriptor": "Man Standing: Light Skin Tone"
	},
	"1F9CD-1F3FC": {
		"key": "1F9CD-1F3FC",
		"value": "🧍🏼",
		"descriptor": "Person Standing: Medium-Light Skin Tone"
	},
	"1F9CD-1F3FC-200D-2640-FE0F": {
		"key": "1F9CD-1F3FC-200D-2640-FE0F",
		"value": "🧍🏼‍♀️",
		"descriptor": "Woman Standing: Medium-Light Skin Tone"
	},
	"1F9CD-1F3FC-200D-2642-FE0F": {
		"key": "1F9CD-1F3FC-200D-2642-FE0F",
		"value": "🧍🏼‍♂️",
		"descriptor": "Man Standing: Medium-Light Skin Tone"
	},
	"1F9CD-1F3FD": {
		"key": "1F9CD-1F3FD",
		"value": "🧍🏽",
		"descriptor": "Person Standing: Medium Skin Tone"
	},
	"1F9CD-1F3FD-200D-2640-FE0F": {
		"key": "1F9CD-1F3FD-200D-2640-FE0F",
		"value": "🧍🏽‍♀️",
		"descriptor": "Woman Standing: Medium Skin Tone"
	},
	"1F9CD-1F3FD-200D-2642-FE0F": {
		"key": "1F9CD-1F3FD-200D-2642-FE0F",
		"value": "🧍🏽‍♂️",
		"descriptor": "Man Standing: Medium Skin Tone"
	},
	"1F9CD-1F3FE": {
		"key": "1F9CD-1F3FE",
		"value": "🧍🏾",
		"descriptor": "Person Standing: Medium-Dark Skin Tone"
	},
	"1F9CD-1F3FE-200D-2640-FE0F": {
		"key": "1F9CD-1F3FE-200D-2640-FE0F",
		"value": "🧍🏾‍♀️",
		"descriptor": "Woman Standing: Medium-Dark Skin Tone"
	},
	"1F9CD-1F3FE-200D-2642-FE0F": {
		"key": "1F9CD-1F3FE-200D-2642-FE0F",
		"value": "🧍🏾‍♂️",
		"descriptor": "Man Standing: Medium-Dark Skin Tone"
	},
	"1F9CD-1F3FF": {
		"key": "1F9CD-1F3FF",
		"value": "🧍🏿",
		"descriptor": "Person Standing: Dark Skin Tone"
	},
	"1F9CD-1F3FF-200D-2640-FE0F": {
		"key": "1F9CD-1F3FF-200D-2640-FE0F",
		"value": "🧍🏿‍♀️",
		"descriptor": "Woman Standing: Dark Skin Tone"
	},
	"1F9CD-1F3FF-200D-2642-FE0F": {
		"key": "1F9CD-1F3FF-200D-2642-FE0F",
		"value": "🧍🏿‍♂️",
		"descriptor": "Man Standing: Dark Skin Tone"
	},
	"1F9CD-200D-2640-FE0F": {
		"key": "1F9CD-200D-2640-FE0F",
		"value": "🧍‍♀️",
		"descriptor": "Woman Standing"
	},
	"1F9CD-200D-2642-FE0F": {
		"key": "1F9CD-200D-2642-FE0F",
		"value": "🧍‍♂️",
		"descriptor": "Man Standing"
	},
	"1F9CE": {
		"key": "1F9CE",
		"value": "🧎",
		"descriptor": "Person Kneeling"
	},
	"1F9CE-1F3FB": {
		"key": "1F9CE-1F3FB",
		"value": "🧎🏻",
		"descriptor": "Person Kneeling: Light Skin Tone"
	},
	"1F9CE-1F3FB-200D-2640-FE0F": {
		"key": "1F9CE-1F3FB-200D-2640-FE0F",
		"value": "🧎🏻‍♀️",
		"descriptor": "Woman Kneeling: Light Skin Tone"
	},
	"1F9CE-1F3FB-200D-2642-FE0F": {
		"key": "1F9CE-1F3FB-200D-2642-FE0F",
		"value": "🧎🏻‍♂️",
		"descriptor": "Man Kneeling: Light Skin Tone"
	},
	"1F9CE-1F3FC": {
		"key": "1F9CE-1F3FC",
		"value": "🧎🏼",
		"descriptor": "Person Kneeling: Medium-Light Skin Tone"
	},
	"1F9CE-1F3FC-200D-2640-FE0F": {
		"key": "1F9CE-1F3FC-200D-2640-FE0F",
		"value": "🧎🏼‍♀️",
		"descriptor": "Woman Kneeling: Medium-Light Skin Tone"
	},
	"1F9CE-1F3FC-200D-2642-FE0F": {
		"key": "1F9CE-1F3FC-200D-2642-FE0F",
		"value": "🧎🏼‍♂️",
		"descriptor": "Man Kneeling: Medium-Light Skin Tone"
	},
	"1F9CE-1F3FD": {
		"key": "1F9CE-1F3FD",
		"value": "🧎🏽",
		"descriptor": "Person Kneeling: Medium Skin Tone"
	},
	"1F9CE-1F3FD-200D-2640-FE0F": {
		"key": "1F9CE-1F3FD-200D-2640-FE0F",
		"value": "🧎🏽‍♀️",
		"descriptor": "Woman Kneeling: Medium Skin Tone"
	},
	"1F9CE-1F3FD-200D-2642-FE0F": {
		"key": "1F9CE-1F3FD-200D-2642-FE0F",
		"value": "🧎🏽‍♂️",
		"descriptor": "Man Kneeling: Medium Skin Tone"
	},
	"1F9CE-1F3FE": {
		"key": "1F9CE-1F3FE",
		"value": "🧎🏾",
		"descriptor": "Person Kneeling: Medium-Dark Skin Tone"
	},
	"1F9CE-1F3FE-200D-2640-FE0F": {
		"key": "1F9CE-1F3FE-200D-2640-FE0F",
		"value": "🧎🏾‍♀️",
		"descriptor": "Woman Kneeling: Medium-Dark Skin Tone"
	},
	"1F9CE-1F3FE-200D-2642-FE0F": {
		"key": "1F9CE-1F3FE-200D-2642-FE0F",
		"value": "🧎🏾‍♂️",
		"descriptor": "Man Kneeling: Medium-Dark Skin Tone"
	},
	"1F9CE-1F3FF": {
		"key": "1F9CE-1F3FF",
		"value": "🧎🏿",
		"descriptor": "Person Kneeling: Dark Skin Tone"
	},
	"1F9CE-1F3FF-200D-2640-FE0F": {
		"key": "1F9CE-1F3FF-200D-2640-FE0F",
		"value": "🧎🏿‍♀️",
		"descriptor": "Woman Kneeling: Dark Skin Tone"
	},
	"1F9CE-1F3FF-200D-2642-FE0F": {
		"key": "1F9CE-1F3FF-200D-2642-FE0F",
		"value": "🧎🏿‍♂️",
		"descriptor": "Man Kneeling: Dark Skin Tone"
	},
	"1F9CE-200D-2640-FE0F": {
		"key": "1F9CE-200D-2640-FE0F",
		"value": "🧎‍♀️",
		"descriptor": "Woman Kneeling"
	},
	"1F9CE-200D-2642-FE0F": {
		"key": "1F9CE-200D-2642-FE0F",
		"value": "🧎‍♂️",
		"descriptor": "Man Kneeling"
	},
	"1F9CF": {
		"key": "1F9CF",
		"value": "🧏",
		"descriptor": "Deaf Person"
	},
	"1F9CF-1F3FB": {
		"key": "1F9CF-1F3FB",
		"value": "🧏🏻",
		"descriptor": "Deaf Person: Light Skin Tone"
	},
	"1F9CF-1F3FB-200D-2640-FE0F": {
		"key": "1F9CF-1F3FB-200D-2640-FE0F",
		"value": "🧏🏻‍♀️",
		"descriptor": "Deaf Woman: Light Skin Tone"
	},
	"1F9CF-1F3FB-200D-2642-FE0F": {
		"key": "1F9CF-1F3FB-200D-2642-FE0F",
		"value": "🧏🏻‍♂️",
		"descriptor": "Deaf Man: Light Skin Tone"
	},
	"1F9CF-1F3FC": {
		"key": "1F9CF-1F3FC",
		"value": "🧏🏼",
		"descriptor": "Deaf Person: Medium-Light Skin Tone"
	},
	"1F9CF-1F3FC-200D-2640-FE0F": {
		"key": "1F9CF-1F3FC-200D-2640-FE0F",
		"value": "🧏🏼‍♀️",
		"descriptor": "Deaf Woman: Medium-Light Skin Tone"
	},
	"1F9CF-1F3FC-200D-2642-FE0F": {
		"key": "1F9CF-1F3FC-200D-2642-FE0F",
		"value": "🧏🏼‍♂️",
		"descriptor": "Deaf Man: Medium-Light Skin Tone"
	},
	"1F9CF-1F3FD": {
		"key": "1F9CF-1F3FD",
		"value": "🧏🏽",
		"descriptor": "Deaf Person: Medium Skin Tone"
	},
	"1F9CF-1F3FD-200D-2640-FE0F": {
		"key": "1F9CF-1F3FD-200D-2640-FE0F",
		"value": "🧏🏽‍♀️",
		"descriptor": "Deaf Woman: Medium Skin Tone"
	},
	"1F9CF-1F3FD-200D-2642-FE0F": {
		"key": "1F9CF-1F3FD-200D-2642-FE0F",
		"value": "🧏🏽‍♂️",
		"descriptor": "Deaf Man: Medium Skin Tone"
	},
	"1F9CF-1F3FE": {
		"key": "1F9CF-1F3FE",
		"value": "🧏🏾",
		"descriptor": "Deaf Person: Medium-Dark Skin Tone"
	},
	"1F9CF-1F3FE-200D-2640-FE0F": {
		"key": "1F9CF-1F3FE-200D-2640-FE0F",
		"value": "🧏🏾‍♀️",
		"descriptor": "Deaf Woman: Medium-Dark Skin Tone"
	},
	"1F9CF-1F3FE-200D-2642-FE0F": {
		"key": "1F9CF-1F3FE-200D-2642-FE0F",
		"value": "🧏🏾‍♂️",
		"descriptor": "Deaf Man: Medium-Dark Skin Tone"
	},
	"1F9CF-1F3FF": {
		"key": "1F9CF-1F3FF",
		"value": "🧏🏿",
		"descriptor": "Deaf Person: Dark Skin Tone"
	},
	"1F9CF-1F3FF-200D-2640-FE0F": {
		"key": "1F9CF-1F3FF-200D-2640-FE0F",
		"value": "🧏🏿‍♀️",
		"descriptor": "Deaf Woman: Dark Skin Tone"
	},
	"1F9CF-1F3FF-200D-2642-FE0F": {
		"key": "1F9CF-1F3FF-200D-2642-FE0F",
		"value": "🧏🏿‍♂️",
		"descriptor": "Deaf Man: Dark Skin Tone"
	},
	"1F9CF-200D-2640-FE0F": {
		"key": "1F9CF-200D-2640-FE0F",
		"value": "🧏‍♀️",
		"descriptor": "Deaf Woman"
	},
	"1F9CF-200D-2642-FE0F": {
		"key": "1F9CF-200D-2642-FE0F",
		"value": "🧏‍♂️",
		"descriptor": "Deaf Man"
	},
	"1F9D0": {
		"key": "1F9D0",
		"value": "🧐",
		"descriptor": "Face With Monocle"
	},
	"1F9D1": {
		"key": "1F9D1",
		"value": "🧑",
		"descriptor": "Person"
	},
	"1F9D1-1F3FB": {
		"key": "1F9D1-1F3FB",
		"value": "🧑🏻",
		"descriptor": "Person: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F33E": {
		"key": "1F9D1-1F3FB-200D-1F33E",
		"value": "🧑🏻‍🌾",
		"descriptor": "Farmer: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F373": {
		"key": "1F9D1-1F3FB-200D-1F373",
		"value": "🧑🏻‍🍳",
		"descriptor": "Cook: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F393": {
		"key": "1F9D1-1F3FB-200D-1F393",
		"value": "🧑🏻‍🎓",
		"descriptor": "Student: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F3A4": {
		"key": "1F9D1-1F3FB-200D-1F3A4",
		"value": "🧑🏻‍🎤",
		"descriptor": "Singer: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F3A8": {
		"key": "1F9D1-1F3FB-200D-1F3A8",
		"value": "🧑🏻‍🎨",
		"descriptor": "Artist: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F3EB": {
		"key": "1F9D1-1F3FB-200D-1F3EB",
		"value": "🧑🏻‍🏫",
		"descriptor": "Teacher: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F3ED": {
		"key": "1F9D1-1F3FB-200D-1F3ED",
		"value": "🧑🏻‍🏭",
		"descriptor": "Factory Worker: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F4BB": {
		"key": "1F9D1-1F3FB-200D-1F4BB",
		"value": "🧑🏻‍💻",
		"descriptor": "Technologist: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F4BC": {
		"key": "1F9D1-1F3FB-200D-1F4BC",
		"value": "🧑🏻‍💼",
		"descriptor": "Office Worker: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F527": {
		"key": "1F9D1-1F3FB-200D-1F527",
		"value": "🧑🏻‍🔧",
		"descriptor": "Mechanic: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F52C": {
		"key": "1F9D1-1F3FB-200D-1F52C",
		"value": "🧑🏻‍🔬",
		"descriptor": "Scientist: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F680": {
		"key": "1F9D1-1F3FB-200D-1F680",
		"value": "🧑🏻‍🚀",
		"descriptor": "Astronaut: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F692": {
		"key": "1F9D1-1F3FB-200D-1F692",
		"value": "🧑🏻‍🚒",
		"descriptor": "Firefighter: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FB": {
		"key": "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FB",
		"value": "🧑🏻‍🤝‍🧑🏻",
		"descriptor": "People Holding Hands: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FC": {
		"key": "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FC",
		"value": "🧑🏻‍🤝‍🧑🏼",
		"descriptor": "People Holding Hands: Light Skin Tone, Medium-Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FD": {
		"key": "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FD",
		"value": "🧑🏻‍🤝‍🧑🏽",
		"descriptor": "People Holding Hands: Light Skin Tone, Medium Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FE": {
		"key": "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FE",
		"value": "🧑🏻‍🤝‍🧑🏾",
		"descriptor": "People Holding Hands: Light Skin Tone, Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FF": {
		"key": "1F9D1-1F3FB-200D-1F91D-200D-1F9D1-1F3FF",
		"value": "🧑🏻‍🤝‍🧑🏿",
		"descriptor": "People Holding Hands: Light Skin Tone, Dark Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F9AF": {
		"key": "1F9D1-1F3FB-200D-1F9AF",
		"value": "🧑🏻‍🦯",
		"descriptor": "Person With Probing Cane: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F9B0": {
		"key": "1F9D1-1F3FB-200D-1F9B0",
		"value": "🧑🏻‍🦰",
		"descriptor": "Person: Light Skin Tone, Red Hair"
	},
	"1F9D1-1F3FB-200D-1F9B1": {
		"key": "1F9D1-1F3FB-200D-1F9B1",
		"value": "🧑🏻‍🦱",
		"descriptor": "Person: Light Skin Tone, Curly Hair"
	},
	"1F9D1-1F3FB-200D-1F9B2": {
		"key": "1F9D1-1F3FB-200D-1F9B2",
		"value": "🧑🏻‍🦲",
		"descriptor": "Person: Light Skin Tone, Bald"
	},
	"1F9D1-1F3FB-200D-1F9B3": {
		"key": "1F9D1-1F3FB-200D-1F9B3",
		"value": "🧑🏻‍🦳",
		"descriptor": "Person: Light Skin Tone, White Hair"
	},
	"1F9D1-1F3FB-200D-1F9BC": {
		"key": "1F9D1-1F3FB-200D-1F9BC",
		"value": "🧑🏻‍🦼",
		"descriptor": "Person in Motorized Wheelchair: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-1F9BD": {
		"key": "1F9D1-1F3FB-200D-1F9BD",
		"value": "🧑🏻‍🦽",
		"descriptor": "Person in Manual Wheelchair: Light Skin Ton"
	},
	"1F9D1-1F3FB-200D-2695-FE0F": {
		"key": "1F9D1-1F3FB-200D-2695-FE0F",
		"value": "🧑🏻‍⚕️",
		"descriptor": "Health Worker: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-2696-FE0F": {
		"key": "1F9D1-1F3FB-200D-2696-FE0F",
		"value": "🧑🏻‍⚖️",
		"descriptor": "Judge: Light Skin Tone"
	},
	"1F9D1-1F3FB-200D-2708-FE0F": {
		"key": "1F9D1-1F3FB-200D-2708-FE0F",
		"value": "🧑🏻‍✈️",
		"descriptor": "Pilot: Light Skin Tone"
	},
	"1F9D1-1F3FC": {
		"key": "1F9D1-1F3FC",
		"value": "🧑🏼",
		"descriptor": "Person: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F33E": {
		"key": "1F9D1-1F3FC-200D-1F33E",
		"value": "🧑🏼‍🌾",
		"descriptor": "Farmer: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F373": {
		"key": "1F9D1-1F3FC-200D-1F373",
		"value": "🧑🏼‍🍳",
		"descriptor": "Cook: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F393": {
		"key": "1F9D1-1F3FC-200D-1F393",
		"value": "🧑🏼‍🎓",
		"descriptor": "Student: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F3A4": {
		"key": "1F9D1-1F3FC-200D-1F3A4",
		"value": "🧑🏼‍🎤",
		"descriptor": "Singer: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F3A8": {
		"key": "1F9D1-1F3FC-200D-1F3A8",
		"value": "🧑🏼‍🎨",
		"descriptor": "Artist: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F3EB": {
		"key": "1F9D1-1F3FC-200D-1F3EB",
		"value": "🧑🏼‍🏫",
		"descriptor": "Teacher: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F3ED": {
		"key": "1F9D1-1F3FC-200D-1F3ED",
		"value": "🧑🏼‍🏭",
		"descriptor": "Factory Worker: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F4BB": {
		"key": "1F9D1-1F3FC-200D-1F4BB",
		"value": "🧑🏼‍💻",
		"descriptor": "Technologist: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F4BC": {
		"key": "1F9D1-1F3FC-200D-1F4BC",
		"value": "🧑🏼‍💼",
		"descriptor": "Office Worker: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F527": {
		"key": "1F9D1-1F3FC-200D-1F527",
		"value": "🧑🏼‍🔧",
		"descriptor": "Mechanic: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F52C": {
		"key": "1F9D1-1F3FC-200D-1F52C",
		"value": "🧑🏼‍🔬",
		"descriptor": "Scientist: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F680": {
		"key": "1F9D1-1F3FC-200D-1F680",
		"value": "🧑🏼‍🚀",
		"descriptor": "Astronaut: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F692": {
		"key": "1F9D1-1F3FC-200D-1F692",
		"value": "🧑🏼‍🚒",
		"descriptor": "Firefighter: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FB": {
		"key": "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FB",
		"value": "🧑🏼‍🤝‍🧑🏻",
		"descriptor": "People Holding Hands: Medium-Light Skin Tone, Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FC": {
		"key": "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FC",
		"value": "🧑🏼‍🤝‍🧑🏼",
		"descriptor": "People Holding Hands: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FD": {
		"key": "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FD",
		"value": "🧑🏼‍🤝‍🧑🏽",
		"descriptor": "People Holding Hands: Medium-Light Skin Tone, Medium Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FE": {
		"key": "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FE",
		"value": "🧑🏼‍🤝‍🧑🏾",
		"descriptor": "People Holding Hands: Medium-Light Skin Tone, Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FF": {
		"key": "1F9D1-1F3FC-200D-1F91D-200D-1F9D1-1F3FF",
		"value": "🧑🏼‍🤝‍🧑🏿",
		"descriptor": "People Holding Hands: Medium-Light Skin Tone, Dark Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F9AF": {
		"key": "1F9D1-1F3FC-200D-1F9AF",
		"value": "🧑🏼‍🦯",
		"descriptor": "Person With Probing Cane: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F9B0": {
		"key": "1F9D1-1F3FC-200D-1F9B0",
		"value": "🧑🏼‍🦰",
		"descriptor": "Person: Medium-Light Skin Tone, Red Hair"
	},
	"1F9D1-1F3FC-200D-1F9B1": {
		"key": "1F9D1-1F3FC-200D-1F9B1",
		"value": "🧑🏼‍🦱",
		"descriptor": "Person: Medium-Light Skin Tone, Curly Hair"
	},
	"1F9D1-1F3FC-200D-1F9B2": {
		"key": "1F9D1-1F3FC-200D-1F9B2",
		"value": "🧑🏼‍🦲",
		"descriptor": "Person: Medium-Light Skin Tone, Bald"
	},
	"1F9D1-1F3FC-200D-1F9B3": {
		"key": "1F9D1-1F3FC-200D-1F9B3",
		"value": "🧑🏼‍🦳",
		"descriptor": "Person: Medium-Light Skin Tone, White Hair"
	},
	"1F9D1-1F3FC-200D-1F9BC": {
		"key": "1F9D1-1F3FC-200D-1F9BC",
		"value": "🧑🏼‍🦼",
		"descriptor": "Person in Motorized Wheelchair: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-1F9BD": {
		"key": "1F9D1-1F3FC-200D-1F9BD",
		"value": "🧑🏼‍🦽",
		"descriptor": "Person in Manual Wheelchair: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-2695-FE0F": {
		"key": "1F9D1-1F3FC-200D-2695-FE0F",
		"value": "🧑🏼‍⚕️",
		"descriptor": "Health Worker: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-2696-FE0F": {
		"key": "1F9D1-1F3FC-200D-2696-FE0F",
		"value": "🧑🏼‍⚖️",
		"descriptor": "Judge: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FC-200D-2708-FE0F": {
		"key": "1F9D1-1F3FC-200D-2708-FE0F",
		"value": "🧑🏼‍✈️",
		"descriptor": "Pilot: Medium-Light Skin Tone"
	},
	"1F9D1-1F3FD": {
		"key": "1F9D1-1F3FD",
		"value": "🧑🏽",
		"descriptor": "Person: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F33E": {
		"key": "1F9D1-1F3FD-200D-1F33E",
		"value": "🧑🏽‍🌾",
		"descriptor": "Farmer: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F373": {
		"key": "1F9D1-1F3FD-200D-1F373",
		"value": "🧑🏽‍🍳",
		"descriptor": "Cook: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F393": {
		"key": "1F9D1-1F3FD-200D-1F393",
		"value": "🧑🏽‍🎓",
		"descriptor": "Student: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F3A4": {
		"key": "1F9D1-1F3FD-200D-1F3A4",
		"value": "🧑🏽‍🎤",
		"descriptor": "Singer: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F3A8": {
		"key": "1F9D1-1F3FD-200D-1F3A8",
		"value": "🧑🏽‍🎨",
		"descriptor": "Artist: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F3EB": {
		"key": "1F9D1-1F3FD-200D-1F3EB",
		"value": "🧑🏽‍🏫",
		"descriptor": "Teacher: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F3ED": {
		"key": "1F9D1-1F3FD-200D-1F3ED",
		"value": "🧑🏽‍🏭",
		"descriptor": "Factory Worker: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F4BB": {
		"key": "1F9D1-1F3FD-200D-1F4BB",
		"value": "🧑🏽‍💻",
		"descriptor": "Technologist: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F4BC": {
		"key": "1F9D1-1F3FD-200D-1F4BC",
		"value": "🧑🏽‍💼",
		"descriptor": "Office Worker: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F527": {
		"key": "1F9D1-1F3FD-200D-1F527",
		"value": "🧑🏽‍🔧",
		"descriptor": "Mechanic: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F52C": {
		"key": "1F9D1-1F3FD-200D-1F52C",
		"value": "🧑🏽‍🔬",
		"descriptor": "Scientist: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F680": {
		"key": "1F9D1-1F3FD-200D-1F680",
		"value": "🧑🏽‍🚀",
		"descriptor": "Astronaut: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F692": {
		"key": "1F9D1-1F3FD-200D-1F692",
		"value": "🧑🏽‍🚒",
		"descriptor": "Firefighter: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FB": {
		"key": "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FB",
		"value": "🧑🏽‍🤝‍🧑🏻",
		"descriptor": "People Holding Hands: Medium Skin Tone, Light Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FC": {
		"key": "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FC",
		"value": "🧑🏽‍🤝‍🧑🏼",
		"descriptor": "People Holding Hands: Medium Skin Tone, Medium-Light Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FD": {
		"key": "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FD",
		"value": "🧑🏽‍🤝‍🧑🏽",
		"descriptor": "People Holding Hands: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FE": {
		"key": "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FE",
		"value": "🧑🏽‍🤝‍🧑🏾",
		"descriptor": "People Holding Hands: Medium Skin Tone, Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FF": {
		"key": "1F9D1-1F3FD-200D-1F91D-200D-1F9D1-1F3FF",
		"value": "🧑🏽‍🤝‍🧑🏿",
		"descriptor": "People Holding Hands: Medium Skin Tone, Dark Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F9AF": {
		"key": "1F9D1-1F3FD-200D-1F9AF",
		"value": "🧑🏽‍🦯",
		"descriptor": "Person With Probing Cane: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F9B0": {
		"key": "1F9D1-1F3FD-200D-1F9B0",
		"value": "🧑🏽‍🦰",
		"descriptor": "Person: Medium Skin Tone, Red Hair"
	},
	"1F9D1-1F3FD-200D-1F9B1": {
		"key": "1F9D1-1F3FD-200D-1F9B1",
		"value": "🧑🏽‍🦱",
		"descriptor": "Person: Medium Skin Tone, Curly Hair"
	},
	"1F9D1-1F3FD-200D-1F9B2": {
		"key": "1F9D1-1F3FD-200D-1F9B2",
		"value": "🧑🏽‍🦲",
		"descriptor": "Person: Medium Skin Tone, Bald"
	},
	"1F9D1-1F3FD-200D-1F9B3": {
		"key": "1F9D1-1F3FD-200D-1F9B3",
		"value": "🧑🏽‍🦳",
		"descriptor": "Person: Medium Skin Tone, White Hair"
	},
	"1F9D1-1F3FD-200D-1F9BC": {
		"key": "1F9D1-1F3FD-200D-1F9BC",
		"value": "🧑🏽‍🦼",
		"descriptor": "Person in Motorized Wheelchair: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-1F9BD": {
		"key": "1F9D1-1F3FD-200D-1F9BD",
		"value": "🧑🏽‍🦽",
		"descriptor": "Person in Manual Wheelchair: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-2695-FE0F": {
		"key": "1F9D1-1F3FD-200D-2695-FE0F",
		"value": "🧑🏽‍⚕️",
		"descriptor": "Health Worker: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-2696-FE0F": {
		"key": "1F9D1-1F3FD-200D-2696-FE0F",
		"value": "🧑🏽‍⚖️",
		"descriptor": "Judge: Medium Skin Tone"
	},
	"1F9D1-1F3FD-200D-2708-FE0F": {
		"key": "1F9D1-1F3FD-200D-2708-FE0F",
		"value": "🧑🏽‍✈️",
		"descriptor": "Pilot: Medium Skin Tone"
	},
	"1F9D1-1F3FE": {
		"key": "1F9D1-1F3FE",
		"value": "🧑🏾",
		"descriptor": "Person: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F33E": {
		"key": "1F9D1-1F3FE-200D-1F33E",
		"value": "🧑🏾‍🌾",
		"descriptor": "Farmer: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F373": {
		"key": "1F9D1-1F3FE-200D-1F373",
		"value": "🧑🏾‍🍳",
		"descriptor": "Cook: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F393": {
		"key": "1F9D1-1F3FE-200D-1F393",
		"value": "🧑🏾‍🎓",
		"descriptor": "Student: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F3A4": {
		"key": "1F9D1-1F3FE-200D-1F3A4",
		"value": "🧑🏾‍🎤",
		"descriptor": "Singer: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F3A8": {
		"key": "1F9D1-1F3FE-200D-1F3A8",
		"value": "🧑🏾‍🎨",
		"descriptor": "Artist: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F3EB": {
		"key": "1F9D1-1F3FE-200D-1F3EB",
		"value": "🧑🏾‍🏫",
		"descriptor": "Teacher: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F3ED": {
		"key": "1F9D1-1F3FE-200D-1F3ED",
		"value": "🧑🏾‍🏭",
		"descriptor": "Factory Worker: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F4BB": {
		"key": "1F9D1-1F3FE-200D-1F4BB",
		"value": "🧑🏾‍💻",
		"descriptor": "Technologist: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F4BC": {
		"key": "1F9D1-1F3FE-200D-1F4BC",
		"value": "🧑🏾‍💼",
		"descriptor": "Office Worker: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F527": {
		"key": "1F9D1-1F3FE-200D-1F527",
		"value": "🧑🏾‍🔧",
		"descriptor": "Mechanic: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F52C": {
		"key": "1F9D1-1F3FE-200D-1F52C",
		"value": "🧑🏾‍🔬",
		"descriptor": "Scientist: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F680": {
		"key": "1F9D1-1F3FE-200D-1F680",
		"value": "🧑🏾‍🚀",
		"descriptor": "Astronaut: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F692": {
		"key": "1F9D1-1F3FE-200D-1F692",
		"value": "🧑🏾‍🚒",
		"descriptor": "Firefighter: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FB": {
		"key": "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FB",
		"value": "🧑🏾‍🤝‍🧑🏻",
		"descriptor": "People Holding Hands: Medium-Dark Skin Tone, Light Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FC": {
		"key": "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FC",
		"value": "🧑🏾‍🤝‍🧑🏼",
		"descriptor": "People Holding Hands: Medium-Dark Skin Tone, Medium-Light Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FD": {
		"key": "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FD",
		"value": "🧑🏾‍🤝‍🧑🏽",
		"descriptor": "People Holding Hands: Medium-Dark Skin Tone, Medium Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FE": {
		"key": "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FE",
		"value": "🧑🏾‍🤝‍🧑🏾",
		"descriptor": "People Holding Hands: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FF": {
		"key": "1F9D1-1F3FE-200D-1F91D-200D-1F9D1-1F3FF",
		"value": "🧑🏾‍🤝‍🧑🏿",
		"descriptor": "People Holding Hands: Medium-Dark Skin Tone, Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F9AF": {
		"key": "1F9D1-1F3FE-200D-1F9AF",
		"value": "🧑🏾‍🦯",
		"descriptor": "Person With Probing Cane: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F9B0": {
		"key": "1F9D1-1F3FE-200D-1F9B0",
		"value": "🧑🏾‍🦰",
		"descriptor": "Person: Medium-Dark Skin Tone, Red Hair"
	},
	"1F9D1-1F3FE-200D-1F9B1": {
		"key": "1F9D1-1F3FE-200D-1F9B1",
		"value": "🧑🏾‍🦱",
		"descriptor": "Person: Medium-Dark Skin Tone, Curly Hair"
	},
	"1F9D1-1F3FE-200D-1F9B2": {
		"key": "1F9D1-1F3FE-200D-1F9B2",
		"value": "🧑🏾‍🦲",
		"descriptor": "Person: Medium-Dark Skin Tone, Bald"
	},
	"1F9D1-1F3FE-200D-1F9B3": {
		"key": "1F9D1-1F3FE-200D-1F9B3",
		"value": "🧑🏾‍🦳",
		"descriptor": "Person: Medium-Dark Skin Tone, White Hair"
	},
	"1F9D1-1F3FE-200D-1F9BC": {
		"key": "1F9D1-1F3FE-200D-1F9BC",
		"value": "🧑🏾‍🦼",
		"descriptor": "Person in Motorized Wheelchair: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-1F9BD": {
		"key": "1F9D1-1F3FE-200D-1F9BD",
		"value": "🧑🏾‍🦽",
		"descriptor": "Person in Manual Wheelchair: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-2695-FE0F": {
		"key": "1F9D1-1F3FE-200D-2695-FE0F",
		"value": "🧑🏾‍⚕️",
		"descriptor": "Health Worker: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-2696-FE0F": {
		"key": "1F9D1-1F3FE-200D-2696-FE0F",
		"value": "🧑🏾‍⚖️",
		"descriptor": "Judge: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FE-200D-2708-FE0F": {
		"key": "1F9D1-1F3FE-200D-2708-FE0F",
		"value": "🧑🏾‍✈️",
		"descriptor": "Pilot: Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FF": {
		"key": "1F9D1-1F3FF",
		"value": "🧑🏿",
		"descriptor": "Person: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F33E": {
		"key": "1F9D1-1F3FF-200D-1F33E",
		"value": "🧑🏿‍🌾",
		"descriptor": "Farmer: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F373": {
		"key": "1F9D1-1F3FF-200D-1F373",
		"value": "🧑🏿‍🍳",
		"descriptor": "Cook: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F393": {
		"key": "1F9D1-1F3FF-200D-1F393",
		"value": "🧑🏿‍🎓",
		"descriptor": "Student: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F3A4": {
		"key": "1F9D1-1F3FF-200D-1F3A4",
		"value": "🧑🏿‍🎤",
		"descriptor": "Singer: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F3A8": {
		"key": "1F9D1-1F3FF-200D-1F3A8",
		"value": "🧑🏿‍🎨",
		"descriptor": "Artist: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F3EB": {
		"key": "1F9D1-1F3FF-200D-1F3EB",
		"value": "🧑🏿‍🏫",
		"descriptor": "Teacher: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F3ED": {
		"key": "1F9D1-1F3FF-200D-1F3ED",
		"value": "🧑🏿‍🏭",
		"descriptor": "Factory Worker: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F4BB": {
		"key": "1F9D1-1F3FF-200D-1F4BB",
		"value": "🧑🏿‍💻",
		"descriptor": "Technologist: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F4BC": {
		"key": "1F9D1-1F3FF-200D-1F4BC",
		"value": "🧑🏿‍💼",
		"descriptor": "Office Worker: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F527": {
		"key": "1F9D1-1F3FF-200D-1F527",
		"value": "🧑🏿‍🔧",
		"descriptor": "Mechanic: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F52C": {
		"key": "1F9D1-1F3FF-200D-1F52C",
		"value": "🧑🏿‍🔬",
		"descriptor": "Scientist: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F680": {
		"key": "1F9D1-1F3FF-200D-1F680",
		"value": "🧑🏿‍🚀",
		"descriptor": "Astronaut: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F692": {
		"key": "1F9D1-1F3FF-200D-1F692",
		"value": "🧑🏿‍🚒",
		"descriptor": "Firefighter: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FB": {
		"key": "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FB",
		"value": "🧑🏿‍🤝‍🧑🏻",
		"descriptor": "People Holding Hands: Dark Skin Tone, Light Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FC": {
		"key": "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FC",
		"value": "🧑🏿‍🤝‍🧑🏼",
		"descriptor": "People Holding Hands: Dark Skin Tone, Medium-Light Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FD": {
		"key": "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FD",
		"value": "🧑🏿‍🤝‍🧑🏽",
		"descriptor": "People Holding Hands: Dark Skin Tone, Medium Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FE": {
		"key": "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FE",
		"value": "🧑🏿‍🤝‍🧑🏾",
		"descriptor": "People Holding Hands: Dark Skin Tone, Medium-Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FF": {
		"key": "1F9D1-1F3FF-200D-1F91D-200D-1F9D1-1F3FF",
		"value": "🧑🏿‍🤝‍🧑🏿",
		"descriptor": "People Holding Hands: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F9AF": {
		"key": "1F9D1-1F3FF-200D-1F9AF",
		"value": "🧑🏿‍🦯",
		"descriptor": "Person With Probing Cane: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F9B0": {
		"key": "1F9D1-1F3FF-200D-1F9B0",
		"value": "🧑🏿‍🦰",
		"descriptor": "Person: Dark Skin Tone, Red Hair"
	},
	"1F9D1-1F3FF-200D-1F9B1": {
		"key": "1F9D1-1F3FF-200D-1F9B1",
		"value": "🧑🏿‍🦱",
		"descriptor": "Person: Dark Skin Tone, Curly Hair"
	},
	"1F9D1-1F3FF-200D-1F9B2": {
		"key": "1F9D1-1F3FF-200D-1F9B2",
		"value": "🧑🏿‍🦲",
		"descriptor": "Person: Dark Skin Tone, Bald"
	},
	"1F9D1-1F3FF-200D-1F9B3": {
		"key": "1F9D1-1F3FF-200D-1F9B3",
		"value": "🧑🏿‍🦳",
		"descriptor": "Person: Dark Skin Tone, White Hair"
	},
	"1F9D1-1F3FF-200D-1F9BC": {
		"key": "1F9D1-1F3FF-200D-1F9BC",
		"value": "🧑🏿‍🦼",
		"descriptor": "Person in Motorized Wheelchair: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-1F9BD": {
		"key": "1F9D1-1F3FF-200D-1F9BD",
		"value": "🧑🏿‍🦽",
		"descriptor": "Person in Manual Wheelchair: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-2695-FE0F": {
		"key": "1F9D1-1F3FF-200D-2695-FE0F",
		"value": "🧑🏿‍⚕️",
		"descriptor": "Health Worker: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-2696-FE0F": {
		"key": "1F9D1-1F3FF-200D-2696-FE0F",
		"value": "🧑🏿‍⚖️",
		"descriptor": "Judge: Dark Skin Tone"
	},
	"1F9D1-1F3FF-200D-2708-FE0F": {
		"key": "1F9D1-1F3FF-200D-2708-FE0F",
		"value": "🧑🏿‍✈️",
		"descriptor": "Pilot: Dark Skin Tone"
	},
	"1F9D1-200D-1F33E": {
		"key": "1F9D1-200D-1F33E",
		"value": "🧑‍🌾",
		"descriptor": "Farmer"
	},
	"1F9D1-200D-1F373": {
		"key": "1F9D1-200D-1F373",
		"value": "🧑‍🍳",
		"descriptor": "Cook"
	},
	"1F9D1-200D-1F393": {
		"key": "1F9D1-200D-1F393",
		"value": "🧑‍🎓",
		"descriptor": "Student"
	},
	"1F9D1-200D-1F3A4": {
		"key": "1F9D1-200D-1F3A4",
		"value": "🧑‍🎤",
		"descriptor": "Singer"
	},
	"1F9D1-200D-1F3A8": {
		"key": "1F9D1-200D-1F3A8",
		"value": "🧑‍🎨",
		"descriptor": "Artist"
	},
	"1F9D1-200D-1F3EB": {
		"key": "1F9D1-200D-1F3EB",
		"value": "🧑‍🏫",
		"descriptor": "Teacher"
	},
	"1F9D1-200D-1F3ED": {
		"key": "1F9D1-200D-1F3ED",
		"value": "🧑‍🏭",
		"descriptor": "Factory Worker"
	},
	"1F9D1-200D-1F4BB": {
		"key": "1F9D1-200D-1F4BB",
		"value": "🧑‍💻",
		"descriptor": "Technologist"
	},
	"1F9D1-200D-1F4BC": {
		"key": "1F9D1-200D-1F4BC",
		"value": "🧑‍💼",
		"descriptor": "Office Worker"
	},
	"1F9D1-200D-1F527": {
		"key": "1F9D1-200D-1F527",
		"value": "🧑‍🔧",
		"descriptor": "Mechanic"
	},
	"1F9D1-200D-1F52C": {
		"key": "1F9D1-200D-1F52C",
		"value": "🧑‍🔬",
		"descriptor": "Scientist"
	},
	"1F9D1-200D-1F680": {
		"key": "1F9D1-200D-1F680",
		"value": "🧑‍🚀",
		"descriptor": "Astronaut"
	},
	"1F9D1-200D-1F692": {
		"key": "1F9D1-200D-1F692",
		"value": "🧑‍🚒",
		"descriptor": "Firefighter"
	},
	"1F9D1-200D-1F91D-200D-1F9D1": {
		"key": "1F9D1-200D-1F91D-200D-1F9D1",
		"value": "🧑‍🤝‍🧑",
		"descriptor": "People Holding Hands"
	},
	"1F9D1-200D-1F9AF": {
		"key": "1F9D1-200D-1F9AF",
		"value": "🧑‍🦯",
		"descriptor": "Person With Probing Cane"
	},
	"1F9D1-200D-1F9B0": {
		"key": "1F9D1-200D-1F9B0",
		"value": "🧑‍🦰",
		"descriptor": "Person: Red Hair"
	},
	"1F9D1-200D-1F9B1": {
		"key": "1F9D1-200D-1F9B1",
		"value": "🧑‍🦱",
		"descriptor": "Person: Curly Hair"
	},
	"1F9D1-200D-1F9B2": {
		"key": "1F9D1-200D-1F9B2",
		"value": "🧑‍🦲",
		"descriptor": "Person: Bald"
	},
	"1F9D1-200D-1F9B3": {
		"key": "1F9D1-200D-1F9B3",
		"value": "🧑‍🦳",
		"descriptor": "Person: White Hair"
	},
	"1F9D1-200D-1F9BC": {
		"key": "1F9D1-200D-1F9BC",
		"value": "🧑‍🦼",
		"descriptor": "Person in Motorized Wheelchair"
	},
	"1F9D1-200D-1F9BD": {
		"key": "1F9D1-200D-1F9BD",
		"value": "🧑‍🦽",
		"descriptor": "Person in Manual Wheelchair"
	},
	"1F9D1-200D-2695-FE0F": {
		"key": "1F9D1-200D-2695-FE0F",
		"value": "🧑‍⚕️",
		"descriptor": "Health Worker"
	},
	"1F9D1-200D-2696-FE0F": {
		"key": "1F9D1-200D-2696-FE0F",
		"value": "🧑‍⚖️",
		"descriptor": "Judge"
	},
	"1F9D1-200D-2708-FE0F": {
		"key": "1F9D1-200D-2708-FE0F",
		"value": "🧑‍✈️",
		"descriptor": "Pilot"
	},
	"1F9D2": {
		"key": "1F9D2",
		"value": "🧒",
		"descriptor": "Child"
	},
	"1F9D2-1F3FB": {
		"key": "1F9D2-1F3FB",
		"value": "🧒🏻",
		"descriptor": "Child: Light Skin Tone"
	},
	"1F9D2-1F3FC": {
		"key": "1F9D2-1F3FC",
		"value": "🧒🏼",
		"descriptor": "Child: Medium-Light Skin Tone"
	},
	"1F9D2-1F3FD": {
		"key": "1F9D2-1F3FD",
		"value": "🧒🏽",
		"descriptor": "Child: Medium Skin Tone"
	},
	"1F9D2-1F3FE": {
		"key": "1F9D2-1F3FE",
		"value": "🧒🏾",
		"descriptor": "Child: Medium-Dark Skin Tone"
	},
	"1F9D2-1F3FF": {
		"key": "1F9D2-1F3FF",
		"value": "🧒🏿",
		"descriptor": "Child: Dark Skin Tone"
	},
	"1F9D3": {
		"key": "1F9D3",
		"value": "🧓",
		"descriptor": "Older Person"
	},
	"1F9D3-1F3FB": {
		"key": "1F9D3-1F3FB",
		"value": "🧓🏻",
		"descriptor": "Older Person: Light Skin Tone"
	},
	"1F9D3-1F3FC": {
		"key": "1F9D3-1F3FC",
		"value": "🧓🏼",
		"descriptor": "Older Person: Medium-Light Skin Tone"
	},
	"1F9D3-1F3FD": {
		"key": "1F9D3-1F3FD",
		"value": "🧓🏽",
		"descriptor": "Older Person: Medium Skin Tone"
	},
	"1F9D3-1F3FE": {
		"key": "1F9D3-1F3FE",
		"value": "🧓🏾",
		"descriptor": "Older Person: Medium-Dark Skin Tone"
	},
	"1F9D3-1F3FF": {
		"key": "1F9D3-1F3FF",
		"value": "🧓🏿",
		"descriptor": "Older Person: Dark Skin Tone"
	},
	"1F9D4": {
		"key": "1F9D4",
		"value": "🧔",
		"descriptor": "Man: Beard"
	},
	"1F9D4-1F3FB": {
		"key": "1F9D4-1F3FB",
		"value": "🧔🏻",
		"descriptor": "Man: Light Skin Tone, Beard"
	},
	"1F9D4-1F3FC": {
		"key": "1F9D4-1F3FC",
		"value": "🧔🏼",
		"descriptor": "Man: Medium-Light Skin Tone, Beard"
	},
	"1F9D4-1F3FD": {
		"key": "1F9D4-1F3FD",
		"value": "🧔🏽",
		"descriptor": "Man: Medium Skin Tone, Beard"
	},
	"1F9D4-1F3FE": {
		"key": "1F9D4-1F3FE",
		"value": "🧔🏾",
		"descriptor": "Man: Medium-Dark Skin Tone, Beard"
	},
	"1F9D4-1F3FF": {
		"key": "1F9D4-1F3FF",
		"value": "🧔🏿",
		"descriptor": "Man: Dark Skin Tone, Beard"
	},
	"1F9D5": {
		"key": "1F9D5",
		"value": "🧕",
		"descriptor": "Woman With Headscarf"
	},
	"1F9D5-1F3FB": {
		"key": "1F9D5-1F3FB",
		"value": "🧕🏻",
		"descriptor": "Woman With Headscarf: Light Skin Tone"
	},
	"1F9D5-1F3FC": {
		"key": "1F9D5-1F3FC",
		"value": "🧕🏼",
		"descriptor": "Woman With Headscarf: Medium-Light Skin Tone"
	},
	"1F9D5-1F3FD": {
		"key": "1F9D5-1F3FD",
		"value": "🧕🏽",
		"descriptor": "Woman With Headscarf: Medium Skin Tone"
	},
	"1F9D5-1F3FE": {
		"key": "1F9D5-1F3FE",
		"value": "🧕🏾",
		"descriptor": "Woman With Headscarf: Medium-Dark Skin Tone"
	},
	"1F9D5-1F3FF": {
		"key": "1F9D5-1F3FF",
		"value": "🧕🏿",
		"descriptor": "Woman With Headscarf: Dark Skin Tone"
	},
	"1F9D6": {
		"key": "1F9D6",
		"value": "🧖",
		"descriptor": "Person in Steamy Room"
	},
	"1F9D6-1F3FB": {
		"key": "1F9D6-1F3FB",
		"value": "🧖🏻",
		"descriptor": "Person in Steamy Room: Light Skin Tone"
	},
	"1F9D6-1F3FB-200D-2640-FE0F": {
		"key": "1F9D6-1F3FB-200D-2640-FE0F",
		"value": "🧖🏻‍♀️",
		"descriptor": "Woman in Steamy Room: Light Skin Tone"
	},
	"1F9D6-1F3FB-200D-2642-FE0F": {
		"key": "1F9D6-1F3FB-200D-2642-FE0F",
		"value": "🧖🏻‍♂️",
		"descriptor": "Man in Steamy Room: Light Skin Tone"
	},
	"1F9D6-1F3FC": {
		"key": "1F9D6-1F3FC",
		"value": "🧖🏼",
		"descriptor": "Person in Steamy Room: Medium-Light Skin Tone"
	},
	"1F9D6-1F3FC-200D-2640-FE0F": {
		"key": "1F9D6-1F3FC-200D-2640-FE0F",
		"value": "🧖🏼‍♀️",
		"descriptor": "Woman in Steamy Room: Medium-Light Skin Tone"
	},
	"1F9D6-1F3FC-200D-2642-FE0F": {
		"key": "1F9D6-1F3FC-200D-2642-FE0F",
		"value": "🧖🏼‍♂️",
		"descriptor": "Man in Steamy Room: Medium-Light Skin Tone"
	},
	"1F9D6-1F3FD": {
		"key": "1F9D6-1F3FD",
		"value": "🧖🏽",
		"descriptor": "Person in Steamy Room: Medium Skin Tone"
	},
	"1F9D6-1F3FD-200D-2640-FE0F": {
		"key": "1F9D6-1F3FD-200D-2640-FE0F",
		"value": "🧖🏽‍♀️",
		"descriptor": "Woman in Steamy Room: Medium Skin Tone"
	},
	"1F9D6-1F3FD-200D-2642-FE0F": {
		"key": "1F9D6-1F3FD-200D-2642-FE0F",
		"value": "🧖🏽‍♂️",
		"descriptor": "Man in Steamy Room: Medium Skin Tone"
	},
	"1F9D6-1F3FE": {
		"key": "1F9D6-1F3FE",
		"value": "🧖🏾",
		"descriptor": "Person in Steamy Room: Medium-Dark Skin Tone"
	},
	"1F9D6-1F3FE-200D-2640-FE0F": {
		"key": "1F9D6-1F3FE-200D-2640-FE0F",
		"value": "🧖🏾‍♀️",
		"descriptor": "Woman in Steamy Room: Medium-Dark Skin Tone"
	},
	"1F9D6-1F3FE-200D-2642-FE0F": {
		"key": "1F9D6-1F3FE-200D-2642-FE0F",
		"value": "🧖🏾‍♂️",
		"descriptor": "Man in Steamy Room: Medium-Dark Skin Tone"
	},
	"1F9D6-1F3FF": {
		"key": "1F9D6-1F3FF",
		"value": "🧖🏿",
		"descriptor": "Person in Steamy Room: Dark Skin Tone"
	},
	"1F9D6-1F3FF-200D-2640-FE0F": {
		"key": "1F9D6-1F3FF-200D-2640-FE0F",
		"value": "🧖🏿‍♀️",
		"descriptor": "Woman in Steamy Room: Dark Skin Tone"
	},
	"1F9D6-1F3FF-200D-2642-FE0F": {
		"key": "1F9D6-1F3FF-200D-2642-FE0F",
		"value": "🧖🏿‍♂️",
		"descriptor": "Man in Steamy Room: Dark Skin Tone"
	},
	"1F9D6-200D-2640-FE0F": {
		"key": "1F9D6-200D-2640-FE0F",
		"value": "🧖‍♀️",
		"descriptor": "Woman in Steamy Room"
	},
	"1F9D6-200D-2642-FE0F": {
		"key": "1F9D6-200D-2642-FE0F",
		"value": "🧖‍♂️",
		"descriptor": "Man in Steamy Room"
	},
	"1F9D7": {
		"key": "1F9D7",
		"value": "🧗",
		"descriptor": "Person Climbing"
	},
	"1F9D7-1F3FB": {
		"key": "1F9D7-1F3FB",
		"value": "🧗🏻",
		"descriptor": "Person Climbing: Light Skin Tone"
	},
	"1F9D7-1F3FB-200D-2640-FE0F": {
		"key": "1F9D7-1F3FB-200D-2640-FE0F",
		"value": "🧗🏻‍♀️",
		"descriptor": "Woman Climbing: Light Skin Tone"
	},
	"1F9D7-1F3FB-200D-2642-FE0F": {
		"key": "1F9D7-1F3FB-200D-2642-FE0F",
		"value": "🧗🏻‍♂️",
		"descriptor": "Man Climbing: Light Skin Tone"
	},
	"1F9D7-1F3FC": {
		"key": "1F9D7-1F3FC",
		"value": "🧗🏼",
		"descriptor": "Person Climbing: Medium-Light Skin Tone"
	},
	"1F9D7-1F3FC-200D-2640-FE0F": {
		"key": "1F9D7-1F3FC-200D-2640-FE0F",
		"value": "🧗🏼‍♀️",
		"descriptor": "Woman Climbing: Medium-Light Skin Tone"
	},
	"1F9D7-1F3FC-200D-2642-FE0F": {
		"key": "1F9D7-1F3FC-200D-2642-FE0F",
		"value": "🧗🏼‍♂️",
		"descriptor": "Man Climbing: Medium-Light Skin Tone"
	},
	"1F9D7-1F3FD": {
		"key": "1F9D7-1F3FD",
		"value": "🧗🏽",
		"descriptor": "Person Climbing: Medium Skin Tone"
	},
	"1F9D7-1F3FD-200D-2640-FE0F": {
		"key": "1F9D7-1F3FD-200D-2640-FE0F",
		"value": "🧗🏽‍♀️",
		"descriptor": "Woman Climbing: Medium Skin Tone"
	},
	"1F9D7-1F3FD-200D-2642-FE0F": {
		"key": "1F9D7-1F3FD-200D-2642-FE0F",
		"value": "🧗🏽‍♂️",
		"descriptor": "Man Climbing: Medium Skin Tone"
	},
	"1F9D7-1F3FE": {
		"key": "1F9D7-1F3FE",
		"value": "🧗🏾",
		"descriptor": "Person Climbing: Medium-Dark Skin Tone"
	},
	"1F9D7-1F3FE-200D-2640-FE0F": {
		"key": "1F9D7-1F3FE-200D-2640-FE0F",
		"value": "🧗🏾‍♀️",
		"descriptor": "Woman Climbing: Medium-Dark Skin Tone"
	},
	"1F9D7-1F3FE-200D-2642-FE0F": {
		"key": "1F9D7-1F3FE-200D-2642-FE0F",
		"value": "🧗🏾‍♂️",
		"descriptor": "Man Climbing: Medium-Dark Skin Tone"
	},
	"1F9D7-1F3FF": {
		"key": "1F9D7-1F3FF",
		"value": "🧗🏿",
		"descriptor": "Person Climbing: Dark Skin Tone"
	},
	"1F9D7-1F3FF-200D-2640-FE0F": {
		"key": "1F9D7-1F3FF-200D-2640-FE0F",
		"value": "🧗🏿‍♀️",
		"descriptor": "Woman Climbing: Dark Skin Tone"
	},
	"1F9D7-1F3FF-200D-2642-FE0F": {
		"key": "1F9D7-1F3FF-200D-2642-FE0F",
		"value": "🧗🏿‍♂️",
		"descriptor": "Man Climbing: Dark Skin Tone"
	},
	"1F9D7-200D-2640-FE0F": {
		"key": "1F9D7-200D-2640-FE0F",
		"value": "🧗‍♀️",
		"descriptor": "Woman Climbing"
	},
	"1F9D7-200D-2642-FE0F": {
		"key": "1F9D7-200D-2642-FE0F",
		"value": "🧗‍♂️",
		"descriptor": "Man Climbing"
	},
	"1F9D8": {
		"key": "1F9D8",
		"value": "🧘",
		"descriptor": "Person in Lotus Position"
	},
	"1F9D8-1F3FB": {
		"key": "1F9D8-1F3FB",
		"value": "🧘🏻",
		"descriptor": "Person in Lotus Position: Light Skin Tone"
	},
	"1F9D8-1F3FB-200D-2640-FE0F": {
		"key": "1F9D8-1F3FB-200D-2640-FE0F",
		"value": "🧘🏻‍♀️",
		"descriptor": "Woman in Lotus Position: Light Skin Tone"
	},
	"1F9D8-1F3FB-200D-2642-FE0F": {
		"key": "1F9D8-1F3FB-200D-2642-FE0F",
		"value": "🧘🏻‍♂️",
		"descriptor": "Man in Lotus Position: Light Skin Tone"
	},
	"1F9D8-1F3FC": {
		"key": "1F9D8-1F3FC",
		"value": "🧘🏼",
		"descriptor": "Person in Lotus Position: Medium-Light Skin Tone"
	},
	"1F9D8-1F3FC-200D-2640-FE0F": {
		"key": "1F9D8-1F3FC-200D-2640-FE0F",
		"value": "🧘🏼‍♀️",
		"descriptor": "Woman in Lotus Position: Medium-Light Skin Tone"
	},
	"1F9D8-1F3FC-200D-2642-FE0F": {
		"key": "1F9D8-1F3FC-200D-2642-FE0F",
		"value": "🧘🏼‍♂️",
		"descriptor": "Man in Lotus Position: Medium-Light Skin Tone"
	},
	"1F9D8-1F3FD": {
		"key": "1F9D8-1F3FD",
		"value": "🧘🏽",
		"descriptor": "Person in Lotus Position: Medium Skin Tone"
	},
	"1F9D8-1F3FD-200D-2640-FE0F": {
		"key": "1F9D8-1F3FD-200D-2640-FE0F",
		"value": "🧘🏽‍♀️",
		"descriptor": "Woman in Lotus Position: Medium Skin Tone"
	},
	"1F9D8-1F3FD-200D-2642-FE0F": {
		"key": "1F9D8-1F3FD-200D-2642-FE0F",
		"value": "🧘🏽‍♂️",
		"descriptor": "Man in Lotus Position: Medium Skin Tone"
	},
	"1F9D8-1F3FE": {
		"key": "1F9D8-1F3FE",
		"value": "🧘🏾",
		"descriptor": "Person in Lotus Position: Medium-Dark Skin Tone"
	},
	"1F9D8-1F3FE-200D-2640-FE0F": {
		"key": "1F9D8-1F3FE-200D-2640-FE0F",
		"value": "🧘🏾‍♀️",
		"descriptor": "Woman in Lotus Position: Medium-Dark Skin Tone"
	},
	"1F9D8-1F3FE-200D-2642-FE0F": {
		"key": "1F9D8-1F3FE-200D-2642-FE0F",
		"value": "🧘🏾‍♂️",
		"descriptor": "Man in Lotus Position: Medium-Dark Skin Tone"
	},
	"1F9D8-1F3FF": {
		"key": "1F9D8-1F3FF",
		"value": "🧘🏿",
		"descriptor": "Person in Lotus Position: Dark Skin Tone"
	},
	"1F9D8-1F3FF-200D-2640-FE0F": {
		"key": "1F9D8-1F3FF-200D-2640-FE0F",
		"value": "🧘🏿‍♀️",
		"descriptor": "Woman in Lotus Position: Dark Skin Tone"
	},
	"1F9D8-1F3FF-200D-2642-FE0F": {
		"key": "1F9D8-1F3FF-200D-2642-FE0F",
		"value": "🧘🏿‍♂️",
		"descriptor": "Man in Lotus Position: Dark Skin Tone"
	},
	"1F9D8-200D-2640-FE0F": {
		"key": "1F9D8-200D-2640-FE0F",
		"value": "🧘‍♀️",
		"descriptor": "Woman in Lotus Position"
	},
	"1F9D8-200D-2642-FE0F": {
		"key": "1F9D8-200D-2642-FE0F",
		"value": "🧘‍♂️",
		"descriptor": "Man in Lotus Position"
	},
	"1F9D9": {
		"key": "1F9D9",
		"value": "🧙",
		"descriptor": "Mage"
	},
	"1F9D9-1F3FB": {
		"key": "1F9D9-1F3FB",
		"value": "🧙🏻",
		"descriptor": "Mage: Light Skin Tone"
	},
	"1F9D9-1F3FB-200D-2640-FE0F": {
		"key": "1F9D9-1F3FB-200D-2640-FE0F",
		"value": "🧙🏻‍♀️",
		"descriptor": "Woman Mage: Light Skin Tone"
	},
	"1F9D9-1F3FB-200D-2642-FE0F": {
		"key": "1F9D9-1F3FB-200D-2642-FE0F",
		"value": "🧙🏻‍♂️",
		"descriptor": "Man Mage: Light Skin Tone"
	},
	"1F9D9-1F3FC": {
		"key": "1F9D9-1F3FC",
		"value": "🧙🏼",
		"descriptor": "Mage: Medium-Light Skin Tone"
	},
	"1F9D9-1F3FC-200D-2640-FE0F": {
		"key": "1F9D9-1F3FC-200D-2640-FE0F",
		"value": "🧙🏼‍♀️",
		"descriptor": "Woman Mage: Medium-Light Skin Tone"
	},
	"1F9D9-1F3FC-200D-2642-FE0F": {
		"key": "1F9D9-1F3FC-200D-2642-FE0F",
		"value": "🧙🏼‍♂️",
		"descriptor": "Man Mage: Medium-Light Skin Tone"
	},
	"1F9D9-1F3FD": {
		"key": "1F9D9-1F3FD",
		"value": "🧙🏽",
		"descriptor": "Mage: Medium Skin Tone"
	},
	"1F9D9-1F3FD-200D-2640-FE0F": {
		"key": "1F9D9-1F3FD-200D-2640-FE0F",
		"value": "🧙🏽‍♀️",
		"descriptor": "Woman Mage: Medium Skin Tone"
	},
	"1F9D9-1F3FD-200D-2642-FE0F": {
		"key": "1F9D9-1F3FD-200D-2642-FE0F",
		"value": "🧙🏽‍♂️",
		"descriptor": "Man Mage: Medium Skin Tone"
	},
	"1F9D9-1F3FE": {
		"key": "1F9D9-1F3FE",
		"value": "🧙🏾",
		"descriptor": "Mage: Medium-Dark Skin Tone"
	},
	"1F9D9-1F3FE-200D-2640-FE0F": {
		"key": "1F9D9-1F3FE-200D-2640-FE0F",
		"value": "🧙🏾‍♀️",
		"descriptor": "Woman Mage: Medium-Dark Skin Tone"
	},
	"1F9D9-1F3FE-200D-2642-FE0F": {
		"key": "1F9D9-1F3FE-200D-2642-FE0F",
		"value": "🧙🏾‍♂️",
		"descriptor": "Man Mage: Medium-Dark Skin Tone"
	},
	"1F9D9-1F3FF": {
		"key": "1F9D9-1F3FF",
		"value": "🧙🏿",
		"descriptor": "Mage: Dark Skin Tone"
	},
	"1F9D9-1F3FF-200D-2640-FE0F": {
		"key": "1F9D9-1F3FF-200D-2640-FE0F",
		"value": "🧙🏿‍♀️",
		"descriptor": "Woman Mage: Dark Skin Tone"
	},
	"1F9D9-1F3FF-200D-2642-FE0F": {
		"key": "1F9D9-1F3FF-200D-2642-FE0F",
		"value": "🧙🏿‍♂️",
		"descriptor": "Man Mage: Dark Skin Tone"
	},
	"1F9D9-200D-2640-FE0F": {
		"key": "1F9D9-200D-2640-FE0F",
		"value": "🧙‍♀️",
		"descriptor": "Woman Mage"
	},
	"1F9D9-200D-2642-FE0F": {
		"key": "1F9D9-200D-2642-FE0F",
		"value": "🧙‍♂️",
		"descriptor": "Man Mage"
	},
	"1F9DA": {
		"key": "1F9DA",
		"value": "🧚",
		"descriptor": "Fairy"
	},
	"1F9DA-1F3FB": {
		"key": "1F9DA-1F3FB",
		"value": "🧚🏻",
		"descriptor": "Fairy: Light Skin Tone"
	},
	"1F9DA-1F3FB-200D-2640-FE0F": {
		"key": "1F9DA-1F3FB-200D-2640-FE0F",
		"value": "🧚🏻‍♀️",
		"descriptor": "Woman Fairy: Light Skin Tone"
	},
	"1F9DA-1F3FB-200D-2642-FE0F": {
		"key": "1F9DA-1F3FB-200D-2642-FE0F",
		"value": "🧚🏻‍♂️",
		"descriptor": "Man Fairy: Light Skin Tone"
	},
	"1F9DA-1F3FC": {
		"key": "1F9DA-1F3FC",
		"value": "🧚🏼",
		"descriptor": "Fairy: Medium-Light Skin Tone"
	},
	"1F9DA-1F3FC-200D-2640-FE0F": {
		"key": "1F9DA-1F3FC-200D-2640-FE0F",
		"value": "🧚🏼‍♀️",
		"descriptor": "Woman Fairy: Medium-Light Skin Tone"
	},
	"1F9DA-1F3FC-200D-2642-FE0F": {
		"key": "1F9DA-1F3FC-200D-2642-FE0F",
		"value": "🧚🏼‍♂️",
		"descriptor": "Man Fairy: Medium-Light Skin Tone"
	},
	"1F9DA-1F3FD": {
		"key": "1F9DA-1F3FD",
		"value": "🧚🏽",
		"descriptor": "Fairy: Medium Skin Tone"
	},
	"1F9DA-1F3FD-200D-2640-FE0F": {
		"key": "1F9DA-1F3FD-200D-2640-FE0F",
		"value": "🧚🏽‍♀️",
		"descriptor": "Woman Fairy: Medium Skin Tone"
	},
	"1F9DA-1F3FD-200D-2642-FE0F": {
		"key": "1F9DA-1F3FD-200D-2642-FE0F",
		"value": "🧚🏽‍♂️",
		"descriptor": "Man Fairy: Medium Skin Tone"
	},
	"1F9DA-1F3FE": {
		"key": "1F9DA-1F3FE",
		"value": "🧚🏾",
		"descriptor": "Fairy: Medium-Dark Skin Tone"
	},
	"1F9DA-1F3FE-200D-2640-FE0F": {
		"key": "1F9DA-1F3FE-200D-2640-FE0F",
		"value": "🧚🏾‍♀️",
		"descriptor": "Woman Fairy: Medium-Dark Skin Tone"
	},
	"1F9DA-1F3FE-200D-2642-FE0F": {
		"key": "1F9DA-1F3FE-200D-2642-FE0F",
		"value": "🧚🏾‍♂️",
		"descriptor": "Man Fairy: Medium-Dark Skin Tone"
	},
	"1F9DA-1F3FF": {
		"key": "1F9DA-1F3FF",
		"value": "🧚🏿",
		"descriptor": "Fairy: Dark Skin Tone"
	},
	"1F9DA-1F3FF-200D-2640-FE0F": {
		"key": "1F9DA-1F3FF-200D-2640-FE0F",
		"value": "🧚🏿‍♀️",
		"descriptor": "Woman Fairy: Dark Skin Tone"
	},
	"1F9DA-1F3FF-200D-2642-FE0F": {
		"key": "1F9DA-1F3FF-200D-2642-FE0F",
		"value": "🧚🏿‍♂️",
		"descriptor": "Man Fairy: Dark Skin Tone"
	},
	"1F9DA-200D-2640-FE0F": {
		"key": "1F9DA-200D-2640-FE0F",
		"value": "🧚‍♀️",
		"descriptor": "Woman Fairy"
	},
	"1F9DA-200D-2642-FE0F": {
		"key": "1F9DA-200D-2642-FE0F",
		"value": "🧚‍♂️",
		"descriptor": "Man Fairy"
	},
	"1F9DB": {
		"key": "1F9DB",
		"value": "🧛",
		"descriptor": "Vampire"
	},
	"1F9DB-1F3FB": {
		"key": "1F9DB-1F3FB",
		"value": "🧛🏻",
		"descriptor": "Vampire: Light Skin Tone"
	},
	"1F9DB-1F3FB-200D-2640-FE0F": {
		"key": "1F9DB-1F3FB-200D-2640-FE0F",
		"value": "🧛🏻‍♀️",
		"descriptor": "Woman Vampire: Light Skin Tone"
	},
	"1F9DB-1F3FB-200D-2642-FE0F": {
		"key": "1F9DB-1F3FB-200D-2642-FE0F",
		"value": "🧛🏻‍♂️",
		"descriptor": "Man Vampire: Light Skin Tone"
	},
	"1F9DB-1F3FC": {
		"key": "1F9DB-1F3FC",
		"value": "🧛🏼",
		"descriptor": "Vampire: Medium-Light Skin Tone"
	},
	"1F9DB-1F3FC-200D-2640-FE0F": {
		"key": "1F9DB-1F3FC-200D-2640-FE0F",
		"value": "🧛🏼‍♀️",
		"descriptor": "Woman Vampire: Medium-Light Skin Tone"
	},
	"1F9DB-1F3FC-200D-2642-FE0F": {
		"key": "1F9DB-1F3FC-200D-2642-FE0F",
		"value": "🧛🏼‍♂️",
		"descriptor": "Man Vampire: Medium-Light Skin Tone"
	},
	"1F9DB-1F3FD": {
		"key": "1F9DB-1F3FD",
		"value": "🧛🏽",
		"descriptor": "Vampire: Medium Skin Tone"
	},
	"1F9DB-1F3FD-200D-2640-FE0F": {
		"key": "1F9DB-1F3FD-200D-2640-FE0F",
		"value": "🧛🏽‍♀️",
		"descriptor": "Woman Vampire: Medium Skin Tone"
	},
	"1F9DB-1F3FD-200D-2642-FE0F": {
		"key": "1F9DB-1F3FD-200D-2642-FE0F",
		"value": "🧛🏽‍♂️",
		"descriptor": "Man Vampire: Medium Skin Tone"
	},
	"1F9DB-1F3FE": {
		"key": "1F9DB-1F3FE",
		"value": "🧛🏾",
		"descriptor": "Vampire: Medium-Dark Skin Tone"
	},
	"1F9DB-1F3FE-200D-2640-FE0F": {
		"key": "1F9DB-1F3FE-200D-2640-FE0F",
		"value": "🧛🏾‍♀️",
		"descriptor": "Woman Vampire: Medium-Dark Skin Tone"
	},
	"1F9DB-1F3FE-200D-2642-FE0F": {
		"key": "1F9DB-1F3FE-200D-2642-FE0F",
		"value": "🧛🏾‍♂️",
		"descriptor": "Man Vampire: Medium-Dark Skin Tone"
	},
	"1F9DB-1F3FF": {
		"key": "1F9DB-1F3FF",
		"value": "🧛🏿",
		"descriptor": "Vampire: Dark Skin Tone"
	},
	"1F9DB-1F3FF-200D-2640-FE0F": {
		"key": "1F9DB-1F3FF-200D-2640-FE0F",
		"value": "🧛🏿‍♀️",
		"descriptor": "Woman Vampire: Dark Skin Tone"
	},
	"1F9DB-1F3FF-200D-2642-FE0F": {
		"key": "1F9DB-1F3FF-200D-2642-FE0F",
		"value": "🧛🏿‍♂️",
		"descriptor": "Man Vampire: Dark Skin Tone"
	},
	"1F9DB-200D-2640-FE0F": {
		"key": "1F9DB-200D-2640-FE0F",
		"value": "🧛‍♀️",
		"descriptor": "Woman Vampire"
	},
	"1F9DB-200D-2642-FE0F": {
		"key": "1F9DB-200D-2642-FE0F",
		"value": "🧛‍♂️",
		"descriptor": "Man Vampire"
	},
	"1F9DC": {
		"key": "1F9DC",
		"value": "🧜",
		"descriptor": "Merperson"
	},
	"1F9DC-1F3FB": {
		"key": "1F9DC-1F3FB",
		"value": "🧜🏻",
		"descriptor": "Merperson: Light Skin Tone"
	},
	"1F9DC-1F3FB-200D-2640-FE0F": {
		"key": "1F9DC-1F3FB-200D-2640-FE0F",
		"value": "🧜🏻‍♀️",
		"descriptor": "Mermaid: Light Skin Tone"
	},
	"1F9DC-1F3FB-200D-2642-FE0F": {
		"key": "1F9DC-1F3FB-200D-2642-FE0F",
		"value": "🧜🏻‍♂️",
		"descriptor": "Merman: Light Skin Tone"
	},
	"1F9DC-1F3FC": {
		"key": "1F9DC-1F3FC",
		"value": "🧜🏼",
		"descriptor": "Merperson: Medium-Light Skin Tone"
	},
	"1F9DC-1F3FC-200D-2640-FE0F": {
		"key": "1F9DC-1F3FC-200D-2640-FE0F",
		"value": "🧜🏼‍♀️",
		"descriptor": "Mermaid: Medium-Light Skin Tone"
	},
	"1F9DC-1F3FC-200D-2642-FE0F": {
		"key": "1F9DC-1F3FC-200D-2642-FE0F",
		"value": "🧜🏼‍♂️",
		"descriptor": "Merman: Medium-Light Skin Tone"
	},
	"1F9DC-1F3FD": {
		"key": "1F9DC-1F3FD",
		"value": "🧜🏽",
		"descriptor": "Merperson: Medium Skin Tone"
	},
	"1F9DC-1F3FD-200D-2640-FE0F": {
		"key": "1F9DC-1F3FD-200D-2640-FE0F",
		"value": "🧜🏽‍♀️",
		"descriptor": "Mermaid: Medium Skin Tone"
	},
	"1F9DC-1F3FD-200D-2642-FE0F": {
		"key": "1F9DC-1F3FD-200D-2642-FE0F",
		"value": "🧜🏽‍♂️",
		"descriptor": "Merman: Medium Skin Tone"
	},
	"1F9DC-1F3FE": {
		"key": "1F9DC-1F3FE",
		"value": "🧜🏾",
		"descriptor": "Merperson: Medium-Dark Skin Tone"
	},
	"1F9DC-1F3FE-200D-2640-FE0F": {
		"key": "1F9DC-1F3FE-200D-2640-FE0F",
		"value": "🧜🏾‍♀️",
		"descriptor": "Mermaid: Medium-Dark Skin Tone"
	},
	"1F9DC-1F3FE-200D-2642-FE0F": {
		"key": "1F9DC-1F3FE-200D-2642-FE0F",
		"value": "🧜🏾‍♂️",
		"descriptor": "Merman: Medium-Dark Skin Tone"
	},
	"1F9DC-1F3FF": {
		"key": "1F9DC-1F3FF",
		"value": "🧜🏿",
		"descriptor": "Merperson: Dark Skin Tone"
	},
	"1F9DC-1F3FF-200D-2640-FE0F": {
		"key": "1F9DC-1F3FF-200D-2640-FE0F",
		"value": "🧜🏿‍♀️",
		"descriptor": "Mermaid: Dark Skin Tone"
	},
	"1F9DC-1F3FF-200D-2642-FE0F": {
		"key": "1F9DC-1F3FF-200D-2642-FE0F",
		"value": "🧜🏿‍♂️",
		"descriptor": "Merman: Dark Skin Tone"
	},
	"1F9DC-200D-2640-FE0F": {
		"key": "1F9DC-200D-2640-FE0F",
		"value": "🧜‍♀️",
		"descriptor": "Mermaid"
	},
	"1F9DC-200D-2642-FE0F": {
		"key": "1F9DC-200D-2642-FE0F",
		"value": "🧜‍♂️",
		"descriptor": "Merman"
	},
	"1F9DD": {
		"key": "1F9DD",
		"value": "🧝",
		"descriptor": "Elf"
	},
	"1F9DD-1F3FB": {
		"key": "1F9DD-1F3FB",
		"value": "🧝🏻",
		"descriptor": "Elf: Light Skin Tone"
	},
	"1F9DD-1F3FB-200D-2640-FE0F": {
		"key": "1F9DD-1F3FB-200D-2640-FE0F",
		"value": "🧝🏻‍♀️",
		"descriptor": "Woman Elf: Light Skin Tone"
	},
	"1F9DD-1F3FB-200D-2642-FE0F": {
		"key": "1F9DD-1F3FB-200D-2642-FE0F",
		"value": "🧝🏻‍♂️",
		"descriptor": "Man Elf: Light Skin Tone"
	},
	"1F9DD-1F3FC": {
		"key": "1F9DD-1F3FC",
		"value": "🧝🏼",
		"descriptor": "Elf: Medium-Light Skin Tone"
	},
	"1F9DD-1F3FC-200D-2640-FE0F": {
		"key": "1F9DD-1F3FC-200D-2640-FE0F",
		"value": "🧝🏼‍♀️",
		"descriptor": "Woman Elf: Medium-Light Skin Tone"
	},
	"1F9DD-1F3FC-200D-2642-FE0F": {
		"key": "1F9DD-1F3FC-200D-2642-FE0F",
		"value": "🧝🏼‍♂️",
		"descriptor": "Man Elf: Medium-Light Skin Tone"
	},
	"1F9DD-1F3FD": {
		"key": "1F9DD-1F3FD",
		"value": "🧝🏽",
		"descriptor": "Elf: Medium Skin Tone"
	},
	"1F9DD-1F3FD-200D-2640-FE0F": {
		"key": "1F9DD-1F3FD-200D-2640-FE0F",
		"value": "🧝🏽‍♀️",
		"descriptor": "Woman Elf: Medium Skin Tone"
	},
	"1F9DD-1F3FD-200D-2642-FE0F": {
		"key": "1F9DD-1F3FD-200D-2642-FE0F",
		"value": "🧝🏽‍♂️",
		"descriptor": "Man Elf: Medium Skin Tone"
	},
	"1F9DD-1F3FE": {
		"key": "1F9DD-1F3FE",
		"value": "🧝🏾",
		"descriptor": "Elf: Medium-Dark Skin Tone"
	},
	"1F9DD-1F3FE-200D-2640-FE0F": {
		"key": "1F9DD-1F3FE-200D-2640-FE0F",
		"value": "🧝🏾‍♀️",
		"descriptor": "Woman Elf: Medium-Dark Skin Tone"
	},
	"1F9DD-1F3FE-200D-2642-FE0F": {
		"key": "1F9DD-1F3FE-200D-2642-FE0F",
		"value": "🧝🏾‍♂️",
		"descriptor": "Man Elf: Medium-Dark Skin Tone"
	},
	"1F9DD-1F3FF": {
		"key": "1F9DD-1F3FF",
		"value": "🧝🏿",
		"descriptor": "Elf: Dark Skin Tone"
	},
	"1F9DD-1F3FF-200D-2640-FE0F": {
		"key": "1F9DD-1F3FF-200D-2640-FE0F",
		"value": "🧝🏿‍♀️",
		"descriptor": "Woman Elf: Dark Skin Tone"
	},
	"1F9DD-1F3FF-200D-2642-FE0F": {
		"key": "1F9DD-1F3FF-200D-2642-FE0F",
		"value": "🧝🏿‍♂️",
		"descriptor": "Man Elf: Dark Skin Tone"
	},
	"1F9DD-200D-2640-FE0F": {
		"key": "1F9DD-200D-2640-FE0F",
		"value": "🧝‍♀️",
		"descriptor": "Woman Elf"
	},
	"1F9DD-200D-2642-FE0F": {
		"key": "1F9DD-200D-2642-FE0F",
		"value": "🧝‍♂️",
		"descriptor": "Man Elf"
	},
	"1F9DE": {
		"key": "1F9DE",
		"value": "🧞",
		"descriptor": "Genie"
	},
	"1F9DE-200D-2640-FE0F": {
		"key": "1F9DE-200D-2640-FE0F",
		"value": "🧞‍♀️",
		"descriptor": "Woman Genie"
	},
	"1F9DE-200D-2642-FE0F": {
		"key": "1F9DE-200D-2642-FE0F",
		"value": "🧞‍♂️",
		"descriptor": "Man Genie"
	},
	"1F9DF": {
		"key": "1F9DF",
		"value": "🧟",
		"descriptor": "Zombie"
	},
	"1F9DF-200D-2640-FE0F": {
		"key": "1F9DF-200D-2640-FE0F",
		"value": "🧟‍♀️",
		"descriptor": "Woman Zombie"
	},
	"1F9DF-200D-2642-FE0F": {
		"key": "1F9DF-200D-2642-FE0F",
		"value": "🧟‍♂️",
		"descriptor": "Man Zombie"
	},
	"1F9E0": {
		"key": "1F9E0",
		"value": "🧠",
		"descriptor": "Brain"
	},
	"1F9E1": {
		"key": "1F9E1",
		"value": "🧡",
		"descriptor": "Orange Heart"
	},
	"1F9E2": {
		"key": "1F9E2",
		"value": "🧢",
		"descriptor": "Billed Cap"
	},
	"1F9E3": {
		"key": "1F9E3",
		"value": "🧣",
		"descriptor": "Scarf"
	},
	"1F9E4": {
		"key": "1F9E4",
		"value": "🧤",
		"descriptor": "Gloves"
	},
	"1F9E5": {
		"key": "1F9E5",
		"value": "🧥",
		"descriptor": "Coat"
	},
	"1F9E6": {
		"key": "1F9E6",
		"value": "🧦",
		"descriptor": "Socks"
	},
	"1F9E7": {
		"key": "1F9E7",
		"value": "🧧",
		"descriptor": "Red Envelope"
	},
	"1F9E8": {
		"key": "1F9E8",
		"value": "🧨",
		"descriptor": "Firecracker"
	},
	"1F9E9": {
		"key": "1F9E9",
		"value": "🧩",
		"descriptor": "Puzzle Piece"
	},
	"1F9EA": {
		"key": "1F9EA",
		"value": "🧪",
		"descriptor": "Test Tube"
	},
	"1F9EB": {
		"key": "1F9EB",
		"value": "🧫",
		"descriptor": "Petri Dish"
	},
	"1F9EC": {
		"key": "1F9EC",
		"value": "🧬",
		"descriptor": "DNA"
	},
	"1F9ED": {
		"key": "1F9ED",
		"value": "🧭",
		"descriptor": "Compass"
	},
	"1F9EE": {
		"key": "1F9EE",
		"value": "🧮",
		"descriptor": "Abacus"
	},
	"1F9EF": {
		"key": "1F9EF",
		"value": "🧯",
		"descriptor": "Fire Extinguisher"
	},
	"1F9F0": {
		"key": "1F9F0",
		"value": "🧰",
		"descriptor": "Toolbox"
	},
	"1F9F1": {
		"key": "1F9F1",
		"value": "🧱",
		"descriptor": "Brick"
	},
	"1F9F2": {
		"key": "1F9F2",
		"value": "🧲",
		"descriptor": "Magnet"
	},
	"1F9F3": {
		"key": "1F9F3",
		"value": "🧳",
		"descriptor": "Luggage"
	},
	"1F9F4": {
		"key": "1F9F4",
		"value": "🧴",
		"descriptor": "Lotion Bottle"
	},
	"1F9F5": {
		"key": "1F9F5",
		"value": "🧵",
		"descriptor": "Thread"
	},
	"1F9F6": {
		"key": "1F9F6",
		"value": "🧶",
		"descriptor": "Yarn"
	},
	"1F9F7": {
		"key": "1F9F7",
		"value": "🧷",
		"descriptor": "Safety Pin"
	},
	"1F9F8": {
		"key": "1F9F8",
		"value": "🧸",
		"descriptor": "Teddy Bear"
	},
	"1F9F9": {
		"key": "1F9F9",
		"value": "🧹",
		"descriptor": "Broom"
	},
	"1F9FA": {
		"key": "1F9FA",
		"value": "🧺",
		"descriptor": "Basket"
	},
	"1F9FB": {
		"key": "1F9FB",
		"value": "🧻",
		"descriptor": "Roll of Paper"
	},
	"1F9FC": {
		"key": "1F9FC",
		"value": "🧼",
		"descriptor": "Soap"
	},
	"1F9FD": {
		"key": "1F9FD",
		"value": "🧽",
		"descriptor": "Sponge"
	},
	"1F9FE": {
		"key": "1F9FE",
		"value": "🧾",
		"descriptor": "Receipt"
	},
	"1F9FF": {
		"key": "1F9FF",
		"value": "🧿",
		"descriptor": "Nazar Amulet"
	},
	"1FA70": {
		"key": "1FA70",
		"value": "🩰",
		"descriptor": "Ballet Shoes"
	},
	"1FA71": {
		"key": "1FA71",
		"value": "🩱",
		"descriptor": "One-Piece Swimsuit"
	},
	"1FA72": {
		"key": "1FA72",
		"value": "🩲",
		"descriptor": "Briefs"
	},
	"1FA73": {
		"key": "1FA73",
		"value": "🩳",
		"descriptor": "Shorts"
	},
	"1FA78": {
		"key": "1FA78",
		"value": "🩸",
		"descriptor": "Drop of Blood"
	},
	"1FA79": {
		"key": "1FA79",
		"value": "🩹",
		"descriptor": "Adhesive Bandage"
	},
	"1FA7A": {
		"key": "1FA7A",
		"value": "🩺",
		"descriptor": "Stethoscope"
	},
	"1FA80": {
		"key": "1FA80",
		"value": "🪀",
		"descriptor": "Yo-Yo"
	},
	"1FA81": {
		"key": "1FA81",
		"value": "🪁",
		"descriptor": "Kite"
	},
	"1FA82": {
		"key": "1FA82",
		"value": "🪂",
		"descriptor": "Parachute"
	},
	"1FA90": {
		"key": "1FA90",
		"value": "🪐",
		"descriptor": "Ringed Planet"
	},
	"1FA91": {
		"key": "1FA91",
		"value": "🪑",
		"descriptor": "Chair"
	},
	"1FA92": {
		"key": "1FA92",
		"value": "🪒",
		"descriptor": "Razor"
	},
	"1FA93": {
		"key": "1FA93",
		"value": "🪓",
		"descriptor": "Axe"
	},
	"1FA94": {
		"key": "1FA94",
		"value": "🪔",
		"descriptor": "Diya Lamp"
	},
	"1FA95": {
		"key": "1FA95",
		"value": "🪕",
		"descriptor": "Banjo"
	},
	"203C-FE0F": {
		"key": "203C-FE0F",
		"value": "‼️",
		"descriptor": "Double Exclamation Mark"
	},
	"2049-FE0F": {
		"key": "2049-FE0F",
		"value": "⁉️",
		"descriptor": "Exclamation Question Mark"
	},
	"2122-FE0F": {
		"key": "2122-FE0F",
		"value": "™️",
		"descriptor": "Trade Mark"
	},
	"2139-FE0F": {
		"key": "2139-FE0F",
		"value": "ℹ️",
		"descriptor": "Information"
	},
	"2194-FE0F": {
		"key": "2194-FE0F",
		"value": "↔️",
		"descriptor": "Left-Right Arrow"
	},
	"2195-FE0F": {
		"key": "2195-FE0F",
		"value": "↕️",
		"descriptor": "Up-Down Arrow"
	},
	"2196-FE0F": {
		"key": "2196-FE0F",
		"value": "↖️",
		"descriptor": "Up-Left Arrow"
	},
	"2197-FE0F": {
		"key": "2197-FE0F",
		"value": "↗️",
		"descriptor": "Up-Right Arrow"
	},
	"2198-FE0F": {
		"key": "2198-FE0F",
		"value": "↘️",
		"descriptor": "Down-Right Arrow"
	},
	"2199-FE0F": {
		"key": "2199-FE0F",
		"value": "↙️",
		"descriptor": "Down-Left Arrow"
	},
	"21A9-FE0F": {
		"key": "21A9-FE0F",
		"value": "↩️",
		"descriptor": "Right Arrow Curving Left"
	},
	"21AA-FE0F": {
		"key": "21AA-FE0F",
		"value": "↪️",
		"descriptor": "Left Arrow Curving Right"
	},
	"23-FE0F-20E3": {
		"key": "23-FE0F-20E3",
		"value": "#️⃣",
		"descriptor": "Keycap Number Sign"
	},
	"231A": {
		"key": "231A",
		"value": "⌚",
		"descriptor": "Watch"
	},
	"231B": {
		"key": "231B",
		"value": "⌛",
		"descriptor": "Hourglass Done"
	},
	"2328-FE0F": {
		"key": "2328-FE0F",
		"value": "⌨️",
		"descriptor": "Keyboard"
	},
	"23CF-FE0F": {
		"key": "23CF-FE0F",
		"value": "⏏️",
		"descriptor": "Eject Button"
	},
	"23E9": {
		"key": "23E9",
		"value": "⏩",
		"descriptor": "Fast-Forward Button"
	},
	"23EA": {
		"key": "23EA",
		"value": "⏪",
		"descriptor": "Fast Reverse Button"
	},
	"23EB": {
		"key": "23EB",
		"value": "⏫",
		"descriptor": "Fast Up Button"
	},
	"23EC": {
		"key": "23EC",
		"value": "⏬",
		"descriptor": "Fast Down Button"
	},
	"23ED-FE0F": {
		"key": "23ED-FE0F",
		"value": "⏭️",
		"descriptor": "Next Track Button"
	},
	"23EE-FE0F": {
		"key": "23EE-FE0F",
		"value": "⏮️",
		"descriptor": "Last Track Button"
	},
	"23EF-FE0F": {
		"key": "23EF-FE0F",
		"value": "⏯️",
		"descriptor": "Play or Pause Button"
	},
	"23F0": {
		"key": "23F0",
		"value": "⏰",
		"descriptor": "Alarm Clock"
	},
	"23F1-FE0F": {
		"key": "23F1-FE0F",
		"value": "⏱️",
		"descriptor": "Stopwatch"
	},
	"23F2-FE0F": {
		"key": "23F2-FE0F",
		"value": "⏲️",
		"descriptor": "Timer Clock"
	},
	"23F3": {
		"key": "23F3",
		"value": "⏳",
		"descriptor": "Hourglass Not Done"
	},
	"23F8-FE0F": {
		"key": "23F8-FE0F",
		"value": "⏸️",
		"descriptor": "Pause Button"
	},
	"23F9-FE0F": {
		"key": "23F9-FE0F",
		"value": "⏹️",
		"descriptor": "Stop Button"
	},
	"23FA-FE0F": {
		"key": "23FA-FE0F",
		"value": "⏺️",
		"descriptor": "Record Button"
	},
	"24C2-FE0F": {
		"key": "24C2-FE0F",
		"value": "Ⓜ️",
		"descriptor": "Circled M"
	},
	"25AA-FE0F": {
		"key": "25AA-FE0F",
		"value": "▪️",
		"descriptor": "Black Small Square"
	},
	"25AB-FE0F": {
		"key": "25AB-FE0F",
		"value": "▫️",
		"descriptor": "White Small Square"
	},
	"25B6-FE0F": {
		"key": "25B6-FE0F",
		"value": "▶️",
		"descriptor": "Play Button"
	},
	"25C0-FE0F": {
		"key": "25C0-FE0F",
		"value": "◀️",
		"descriptor": "Reverse Button"
	},
	"25FB-FE0F": {
		"key": "25FB-FE0F",
		"value": "◻️",
		"descriptor": "White Medium Square"
	},
	"25FC-FE0F": {
		"key": "25FC-FE0F",
		"value": "◼️",
		"descriptor": "Black Medium Square"
	},
	"25FD": {
		"key": "25FD",
		"value": "◽",
		"descriptor": "White Medium-Small Square"
	},
	"25FE": {
		"key": "25FE",
		"value": "◾",
		"descriptor": "Black Medium-Small Square"
	},
	"2600-FE0F": {
		"key": "2600-FE0F",
		"value": "☀️",
		"descriptor": "Sun"
	},
	"2601-FE0F": {
		"key": "2601-FE0F",
		"value": "☁️",
		"descriptor": "Cloud"
	},
	"2602-FE0F": {
		"key": "2602-FE0F",
		"value": "☂️",
		"descriptor": "Umbrella"
	},
	"2603-FE0F": {
		"key": "2603-FE0F",
		"value": "☃️",
		"descriptor": "Snowman"
	},
	"2604-FE0F": {
		"key": "2604-FE0F",
		"value": "☄️",
		"descriptor": "Comet"
	},
	"260E-FE0F": {
		"key": "260E-FE0F",
		"value": "☎️",
		"descriptor": "Telephone"
	},
	"2611-FE0F": {
		"key": "2611-FE0F",
		"value": "☑️",
		"descriptor": "Check Box With Check"
	},
	"2614": {
		"key": "2614",
		"value": "☔",
		"descriptor": "Umbrella With Rain Drops"
	},
	"2615": {
		"key": "2615",
		"value": "☕",
		"descriptor": "Hot Beverage"
	},
	"2618-FE0F": {
		"key": "2618-FE0F",
		"value": "☘️",
		"descriptor": "Shamrock"
	},
	"261D-1F3FB": {
		"key": "261D-1F3FB",
		"value": "☝🏻",
		"descriptor": "Index Pointing Up: Light Skin Tone"
	},
	"261D-1F3FC": {
		"key": "261D-1F3FC",
		"value": "☝🏼",
		"descriptor": "Index Pointing Up: Medium-Light Skin Tone"
	},
	"261D-1F3FD": {
		"key": "261D-1F3FD",
		"value": "☝🏽",
		"descriptor": "Index Pointing Up: Medium Skin Tone"
	},
	"261D-1F3FE": {
		"key": "261D-1F3FE",
		"value": "☝🏾",
		"descriptor": "Index Pointing Up: Medium-Dark Skin Tone"
	},
	"261D-1F3FF": {
		"key": "261D-1F3FF",
		"value": "☝🏿",
		"descriptor": "Index Pointing Up: Dark Skin Tone"
	},
	"261D-FE0F": {
		"key": "261D-FE0F",
		"value": "☝️",
		"descriptor": "Index Pointing Up"
	},
	"2620-FE0F": {
		"key": "2620-FE0F",
		"value": "☠️",
		"descriptor": "Skull and Crossbones"
	},
	"2622-FE0F": {
		"key": "2622-FE0F",
		"value": "☢️",
		"descriptor": "Radioactive"
	},
	"2623-FE0F": {
		"key": "2623-FE0F",
		"value": "☣️",
		"descriptor": "Biohazard"
	},
	"2626-FE0F": {
		"key": "2626-FE0F",
		"value": "☦️",
		"descriptor": "Orthodox Cross"
	},
	"262A-FE0F": {
		"key": "262A-FE0F",
		"value": "☪️",
		"descriptor": "Star and Crescent"
	},
	"262E-FE0F": {
		"key": "262E-FE0F",
		"value": "☮️",
		"descriptor": "Peace Symbol"
	},
	"262F-FE0F": {
		"key": "262F-FE0F",
		"value": "☯️",
		"descriptor": "Yin Yang"
	},
	"2638-FE0F": {
		"key": "2638-FE0F",
		"value": "☸️",
		"descriptor": "Wheel of Dharma"
	},
	"2639-FE0F": {
		"key": "2639-FE0F",
		"value": "☹️",
		"descriptor": "Frowning Face"
	},
	"263A-FE0F": {
		"key": "263A-FE0F",
		"value": "☺️",
		"descriptor": "Smiling Face"
	},
	"2648": {
		"key": "2648",
		"value": "♈",
		"descriptor": "Aries"
	},
	"2649": {
		"key": "2649",
		"value": "♉",
		"descriptor": "Taurus"
	},
	"264A": {
		"key": "264A",
		"value": "♊",
		"descriptor": "Gemini"
	},
	"264B": {
		"key": "264B",
		"value": "♋",
		"descriptor": "Cancer"
	},
	"264C": {
		"key": "264C",
		"value": "♌",
		"descriptor": "Leo"
	},
	"264D": {
		"key": "264D",
		"value": "♍",
		"descriptor": "Virgo"
	},
	"264E": {
		"key": "264E",
		"value": "♎",
		"descriptor": "Libra"
	},
	"264F": {
		"key": "264F",
		"value": "♏",
		"descriptor": "Scorpio"
	},
	"2650": {
		"key": "2650",
		"value": "♐",
		"descriptor": "Sagittarius"
	},
	"2651": {
		"key": "2651",
		"value": "♑",
		"descriptor": "Capricorn"
	},
	"2652": {
		"key": "2652",
		"value": "♒",
		"descriptor": "Aquarius"
	},
	"2653": {
		"key": "2653",
		"value": "♓",
		"descriptor": "Pisces"
	},
	"265F-FE0F": {
		"key": "265F-FE0F",
		"value": "♟️",
		"descriptor": "Chess Pawn"
	},
	"2660-FE0F": {
		"key": "2660-FE0F",
		"value": "♠️",
		"descriptor": "Spade Suit"
	},
	"2663-FE0F": {
		"key": "2663-FE0F",
		"value": "♣️",
		"descriptor": "Club Suit"
	},
	"2665-FE0F": {
		"key": "2665-FE0F",
		"value": "♥️",
		"descriptor": "Heart Suit"
	},
	"2666-FE0F": {
		"key": "2666-FE0F",
		"value": "♦️",
		"descriptor": "Diamond Suit"
	},
	"2668-FE0F": {
		"key": "2668-FE0F",
		"value": "♨️",
		"descriptor": "Hot Springs"
	},
	"267B-FE0F": {
		"key": "267B-FE0F",
		"value": "♻️",
		"descriptor": "Recycling Symbol"
	},
	"267E-FE0F": {
		"key": "267E-FE0F",
		"value": "♾️",
		"descriptor": "Infinity"
	},
	"267F": {
		"key": "267F",
		"value": "♿",
		"descriptor": "Wheelchair Symbol"
	},
	"2692-FE0F": {
		"key": "2692-FE0F",
		"value": "⚒️",
		"descriptor": "Hammer and Pick"
	},
	"2693": {
		"key": "2693",
		"value": "⚓",
		"descriptor": "Anchor"
	},
	"2694-FE0F": {
		"key": "2694-FE0F",
		"value": "⚔️",
		"descriptor": "Crossed Swords"
	},
	"2696-FE0F": {
		"key": "2696-FE0F",
		"value": "⚖️",
		"descriptor": "Balance Scale"
	},
	"2697-FE0F": {
		"key": "2697-FE0F",
		"value": "⚗️",
		"descriptor": "Alembic"
	},
	"2699-FE0F": {
		"key": "2699-FE0F",
		"value": "⚙️",
		"descriptor": "Gear"
	},
	"269B-FE0F": {
		"key": "269B-FE0F",
		"value": "⚛️",
		"descriptor": "Atom Symbol"
	},
	"269C-FE0F": {
		"key": "269C-FE0F",
		"value": "⚜️",
		"descriptor": "Fleur-de-lis"
	},
	"26A0-FE0F": {
		"key": "26A0-FE0F",
		"value": "⚠️",
		"descriptor": "Warning"
	},
	"26A1": {
		"key": "26A1",
		"value": "⚡",
		"descriptor": "High Voltage"
	},
	"26AA": {
		"key": "26AA",
		"value": "⚪",
		"descriptor": "White Circle"
	},
	"26AB": {
		"key": "26AB",
		"value": "⚫",
		"descriptor": "Black Circle"
	},
	"26B0-FE0F": {
		"key": "26B0-FE0F",
		"value": "⚰️",
		"descriptor": "Coffin"
	},
	"26B1-FE0F": {
		"key": "26B1-FE0F",
		"value": "⚱️",
		"descriptor": "Funeral Urn"
	},
	"26BD": {
		"key": "26BD",
		"value": "⚽",
		"descriptor": "Soccer Ball"
	},
	"26BE": {
		"key": "26BE",
		"value": "⚾",
		"descriptor": "Baseball"
	},
	"26C4": {
		"key": "26C4",
		"value": "⛄",
		"descriptor": "Snowman Without Snow"
	},
	"26C5": {
		"key": "26C5",
		"value": "⛅",
		"descriptor": "Sun Behind Cloud"
	},
	"26C8-FE0F": {
		"key": "26C8-FE0F",
		"value": "⛈️",
		"descriptor": "Cloud With Lightning and Rain"
	},
	"26CE": {
		"key": "26CE",
		"value": "⛎",
		"descriptor": "Ophiuchus"
	},
	"26CF-FE0F": {
		"key": "26CF-FE0F",
		"value": "⛏️",
		"descriptor": "Pick"
	},
	"26D1-FE0F": {
		"key": "26D1-FE0F",
		"value": "⛑️",
		"descriptor": "Rescue Worker’s Helmet"
	},
	"26D3-FE0F": {
		"key": "26D3-FE0F",
		"value": "⛓️",
		"descriptor": "Chains"
	},
	"26D4": {
		"key": "26D4",
		"value": "⛔",
		"descriptor": "No Entry"
	},
	"26E9-FE0F": {
		"key": "26E9-FE0F",
		"value": "⛩️",
		"descriptor": "Shinto Shrine"
	},
	"26EA": {
		"key": "26EA",
		"value": "⛪",
		"descriptor": "Church"
	},
	"26F0-FE0F": {
		"key": "26F0-FE0F",
		"value": "⛰️",
		"descriptor": "Mountain"
	},
	"26F1-FE0F": {
		"key": "26F1-FE0F",
		"value": "⛱️",
		"descriptor": "Umbrella on Ground"
	},
	"26F2": {
		"key": "26F2",
		"value": "⛲",
		"descriptor": "Fountain"
	},
	"26F3": {
		"key": "26F3",
		"value": "⛳",
		"descriptor": "Flag in Hole"
	},
	"26F4-FE0F": {
		"key": "26F4-FE0F",
		"value": "⛴️",
		"descriptor": "Ferry"
	},
	"26F5": {
		"key": "26F5",
		"value": "⛵",
		"descriptor": "Sailboat"
	},
	"26F7-FE0F": {
		"key": "26F7-FE0F",
		"value": "⛷️",
		"descriptor": "Skier"
	},
	"26F8-FE0F": {
		"key": "26F8-FE0F",
		"value": "⛸️",
		"descriptor": "Ice Skate"
	},
	"26F9-1F3FB": {
		"key": "26F9-1F3FB",
		"value": "⛹🏻",
		"descriptor": "Person Bouncing Ball: Light Skin Tone"
	},
	"26F9-1F3FB-200D-2640-FE0F": {
		"key": "26F9-1F3FB-200D-2640-FE0F",
		"value": "⛹🏻‍♀️",
		"descriptor": "Woman Bouncing Ball: Light Skin Tone"
	},
	"26F9-1F3FB-200D-2642-FE0F": {
		"key": "26F9-1F3FB-200D-2642-FE0F",
		"value": "⛹🏻‍♂️",
		"descriptor": "Man Bouncing Ball: Light Skin Tone"
	},
	"26F9-1F3FC": {
		"key": "26F9-1F3FC",
		"value": "⛹🏼",
		"descriptor": "Person Bouncing Ball: Medium-Light Skin Tone"
	},
	"26F9-1F3FC-200D-2640-FE0F": {
		"key": "26F9-1F3FC-200D-2640-FE0F",
		"value": "⛹🏼‍♀️",
		"descriptor": "Woman Bouncing Ball: Medium-Light Skin Tone"
	},
	"26F9-1F3FC-200D-2642-FE0F": {
		"key": "26F9-1F3FC-200D-2642-FE0F",
		"value": "⛹🏼‍♂️",
		"descriptor": "Man Bouncing Ball: Medium-Light Skin Tone"
	},
	"26F9-1F3FD": {
		"key": "26F9-1F3FD",
		"value": "⛹🏽",
		"descriptor": "Person Bouncing Ball: Medium Skin Tone"
	},
	"26F9-1F3FD-200D-2640-FE0F": {
		"key": "26F9-1F3FD-200D-2640-FE0F",
		"value": "⛹🏽‍♀️",
		"descriptor": "Woman Bouncing Ball: Medium Skin Tone"
	},
	"26F9-1F3FD-200D-2642-FE0F": {
		"key": "26F9-1F3FD-200D-2642-FE0F",
		"value": "⛹🏽‍♂️",
		"descriptor": "Man Bouncing Ball: Medium Skin Tone"
	},
	"26F9-1F3FE": {
		"key": "26F9-1F3FE",
		"value": "⛹🏾",
		"descriptor": "Person Bouncing Ball: Medium-Dark Skin Tone"
	},
	"26F9-1F3FE-200D-2640-FE0F": {
		"key": "26F9-1F3FE-200D-2640-FE0F",
		"value": "⛹🏾‍♀️",
		"descriptor": "Woman Bouncing Ball: Medium-Dark Skin Tone"
	},
	"26F9-1F3FE-200D-2642-FE0F": {
		"key": "26F9-1F3FE-200D-2642-FE0F",
		"value": "⛹🏾‍♂️",
		"descriptor": "Man Bouncing Ball: Medium-Dark Skin Tone"
	},
	"26F9-1F3FF": {
		"key": "26F9-1F3FF",
		"value": "⛹🏿",
		"descriptor": "Person Bouncing Ball: Dark Skin Tone"
	},
	"26F9-1F3FF-200D-2640-FE0F": {
		"key": "26F9-1F3FF-200D-2640-FE0F",
		"value": "⛹🏿‍♀️",
		"descriptor": "Woman Bouncing Ball: Dark Skin Tone"
	},
	"26F9-1F3FF-200D-2642-FE0F": {
		"key": "26F9-1F3FF-200D-2642-FE0F",
		"value": "⛹🏿‍♂️",
		"descriptor": "Man Bouncing Ball: Dark Skin Tone"
	},
	"26F9-FE0F": {
		"key": "26F9-FE0F",
		"value": "⛹️",
		"descriptor": "Person Bouncing Ball"
	},
	"26F9-FE0F-200D-2640-FE0F": {
		"key": "26F9-FE0F-200D-2640-FE0F",
		"value": "⛹️‍♀️",
		"descriptor": "Woman Bouncing Ball"
	},
	"26F9-FE0F-200D-2642-FE0F": {
		"key": "26F9-FE0F-200D-2642-FE0F",
		"value": "⛹️‍♂️",
		"descriptor": "Man Bouncing Ball"
	},
	"26FA": {
		"key": "26FA",
		"value": "⛺",
		"descriptor": "Tent"
	},
	"26FD": {
		"key": "26FD",
		"value": "⛽",
		"descriptor": "Fuel Pump"
	},
	"2702-FE0F": {
		"key": "2702-FE0F",
		"value": "✂️",
		"descriptor": "Scissors"
	},
	"2705": {
		"key": "2705",
		"value": "✅",
		"descriptor": "Check Mark Button"
	},
	"2708-FE0F": {
		"key": "2708-FE0F",
		"value": "✈️",
		"descriptor": "Airplane"
	},
	"2709-FE0F": {
		"key": "2709-FE0F",
		"value": "✉️",
		"descriptor": "Envelope"
	},
	"270A": {
		"key": "270A",
		"value": "✊",
		"descriptor": "Raised Fist"
	},
	"270A-1F3FB": {
		"key": "270A-1F3FB",
		"value": "✊🏻",
		"descriptor": "Raised Fist: Light Skin Tone"
	},
	"270A-1F3FC": {
		"key": "270A-1F3FC",
		"value": "✊🏼",
		"descriptor": "Raised Fist: Medium-Light Skin Tone"
	},
	"270A-1F3FD": {
		"key": "270A-1F3FD",
		"value": "✊🏽",
		"descriptor": "Raised Fist: Medium Skin Tone"
	},
	"270A-1F3FE": {
		"key": "270A-1F3FE",
		"value": "✊🏾",
		"descriptor": "Raised Fist: Medium-Dark Skin Tone"
	},
	"270A-1F3FF": {
		"key": "270A-1F3FF",
		"value": "✊🏿",
		"descriptor": "Raised Fist: Dark Skin Tone"
	},
	"270B": {
		"key": "270B",
		"value": "✋",
		"descriptor": "Raised Hand"
	},
	"270B-1F3FB": {
		"key": "270B-1F3FB",
		"value": "✋🏻",
		"descriptor": "Raised Hand: Light Skin Tone"
	},
	"270B-1F3FC": {
		"key": "270B-1F3FC",
		"value": "✋🏼",
		"descriptor": "Raised Hand: Medium-Light Skin Tone"
	},
	"270B-1F3FD": {
		"key": "270B-1F3FD",
		"value": "✋🏽",
		"descriptor": "Raised Hand: Medium Skin Tone"
	},
	"270B-1F3FE": {
		"key": "270B-1F3FE",
		"value": "✋🏾",
		"descriptor": "Raised Hand: Medium-Dark Skin Tone"
	},
	"270B-1F3FF": {
		"key": "270B-1F3FF",
		"value": "✋🏿",
		"descriptor": "Raised Hand: Dark Skin Tone"
	},
	"270C-1F3FB": {
		"key": "270C-1F3FB",
		"value": "✌🏻",
		"descriptor": "Victory Hand: Light Skin Tone"
	},
	"270C-1F3FC": {
		"key": "270C-1F3FC",
		"value": "✌🏼",
		"descriptor": "Victory Hand: Medium-Light Skin Tone"
	},
	"270C-1F3FD": {
		"key": "270C-1F3FD",
		"value": "✌🏽",
		"descriptor": "Victory Hand: Medium Skin Tone"
	},
	"270C-1F3FE": {
		"key": "270C-1F3FE",
		"value": "✌🏾",
		"descriptor": "Victory Hand: Medium-Dark Skin Tone"
	},
	"270C-1F3FF": {
		"key": "270C-1F3FF",
		"value": "✌🏿",
		"descriptor": "Victory Hand: Dark Skin Tone"
	},
	"270C-FE0F": {
		"key": "270C-FE0F",
		"value": "✌️",
		"descriptor": "Victory Hand"
	},
	"270D-1F3FB": {
		"key": "270D-1F3FB",
		"value": "✍🏻",
		"descriptor": "Writing Hand: Light Skin Tone"
	},
	"270D-1F3FC": {
		"key": "270D-1F3FC",
		"value": "✍🏼",
		"descriptor": "Writing Hand: Medium-Light Skin Tone"
	},
	"270D-1F3FD": {
		"key": "270D-1F3FD",
		"value": "✍🏽",
		"descriptor": "Writing Hand: Medium Skin Tone"
	},
	"270D-1F3FE": {
		"key": "270D-1F3FE",
		"value": "✍🏾",
		"descriptor": "Writing Hand: Medium-Dark Skin Tone"
	},
	"270D-1F3FF": {
		"key": "270D-1F3FF",
		"value": "✍🏿",
		"descriptor": "Writing Hand: Dark Skin Tone"
	},
	"270D-FE0F": {
		"key": "270D-FE0F",
		"value": "✍️",
		"descriptor": "Writing Hand"
	},
	"270F-FE0F": {
		"key": "270F-FE0F",
		"value": "✏️",
		"descriptor": "Pencil"
	},
	"2712-FE0F": {
		"key": "2712-FE0F",
		"value": "✒️",
		"descriptor": "Black Nib"
	},
	"2714-FE0F": {
		"key": "2714-FE0F",
		"value": "✔️",
		"descriptor": "Check Mark"
	},
	"2716-FE0F": {
		"key": "2716-FE0F",
		"value": "✖️",
		"descriptor": "Multiplication Sign"
	},
	"271D-FE0F": {
		"key": "271D-FE0F",
		"value": "✝️",
		"descriptor": "Latin Cross"
	},
	"2721-FE0F": {
		"key": "2721-FE0F",
		"value": "✡️",
		"descriptor": "Star of David"
	},
	"2728": {
		"key": "2728",
		"value": "✨",
		"descriptor": "Sparkles"
	},
	"2733-FE0F": {
		"key": "2733-FE0F",
		"value": "✳️",
		"descriptor": "Eight-Spoked Asterisk"
	},
	"2734-FE0F": {
		"key": "2734-FE0F",
		"value": "✴️",
		"descriptor": "Eight-Pointed Star"
	},
	"2744-FE0F": {
		"key": "2744-FE0F",
		"value": "❄️",
		"descriptor": "Snowflake"
	},
	"2747-FE0F": {
		"key": "2747-FE0F",
		"value": "❇️",
		"descriptor": "Sparkle"
	},
	"274C": {
		"key": "274C",
		"value": "❌",
		"descriptor": "Cross Mark"
	},
	"274E": {
		"key": "274E",
		"value": "❎",
		"descriptor": "Cross Mark Button"
	},
	"2753": {
		"key": "2753",
		"value": "❓",
		"descriptor": "Question Mark"
	},
	"2754": {
		"key": "2754",
		"value": "❔",
		"descriptor": "White Question Mark"
	},
	"2755": {
		"key": "2755",
		"value": "❕",
		"descriptor": "White Exclamation Mark"
	},
	"2757": {
		"key": "2757",
		"value": "❗",
		"descriptor": "Exclamation Mark"
	},
	"2763-FE0F": {
		"key": "2763-FE0F",
		"value": "❣️",
		"descriptor": "Heart Exclamation"
	},
	"2764-FE0F": {
		"key": "2764-FE0F",
		"value": "❤️",
		"descriptor": "Red Heart"
	},
	"2795": {
		"key": "2795",
		"value": "➕",
		"descriptor": "Plus Sign"
	},
	"2796": {
		"key": "2796",
		"value": "➖",
		"descriptor": "Minus Sign"
	},
	"2797": {
		"key": "2797",
		"value": "➗",
		"descriptor": "Division Sign"
	},
	"27A1-FE0F": {
		"key": "27A1-FE0F",
		"value": "➡️",
		"descriptor": "Right Arrow"
	},
	"27B0": {
		"key": "27B0",
		"value": "➰",
		"descriptor": "Curly Loop"
	},
	"27BF": {
		"key": "27BF",
		"value": "➿",
		"descriptor": "Double Curly Loop"
	},
	"2934-FE0F": {
		"key": "2934-FE0F",
		"value": "⤴️",
		"descriptor": "Right Arrow Curving Up"
	},
	"2935-FE0F": {
		"key": "2935-FE0F",
		"value": "⤵️",
		"descriptor": "Right Arrow Curving Down"
	},
	"2A-FE0F-20E3": {
		"key": "2A-FE0F-20E3",
		"value": "*️⃣",
		"descriptor": "Keycap Asterisk"
	},
	"2B05-FE0F": {
		"key": "2B05-FE0F",
		"value": "⬅️",
		"descriptor": "Left Arrow"
	},
	"2B06-FE0F": {
		"key": "2B06-FE0F",
		"value": "⬆️",
		"descriptor": "Up Arrow"
	},
	"2B07-FE0F": {
		"key": "2B07-FE0F",
		"value": "⬇️",
		"descriptor": "Down Arrow"
	},
	"2B1B": {
		"key": "2B1B",
		"value": "⬛",
		"descriptor": "Black Large Square"
	},
	"2B1C": {
		"key": "2B1C",
		"value": "⬜",
		"descriptor": "White Large Square"
	},
	"2B50": {
		"key": "2B50",
		"value": "⭐",
		"descriptor": "Star"
	},
	"2B55": {
		"key": "2B55",
		"value": "⭕",
		"descriptor": "Hollow Red Circle"
	},
	"30-FE0F-20E3": {
		"key": "30-FE0F-20E3",
		"value": "0️⃣",
		"descriptor": "Keycap Digit Zero"
	},
	"3030-FE0F": {
		"key": "3030-FE0F",
		"value": "〰️",
		"descriptor": "Wavy Dash"
	},
	"303D-FE0F": {
		"key": "303D-FE0F",
		"value": "〽️",
		"descriptor": "Part Alternation Mark"
	},
	"31-FE0F-20E3": {
		"key": "31-FE0F-20E3",
		"value": "1️⃣",
		"descriptor": "Keycap Digit One"
	},
	"32-FE0F-20E3": {
		"key": "32-FE0F-20E3",
		"value": "2️⃣",
		"descriptor": "Keycap Digit Two"
	},
	"3297-FE0F": {
		"key": "3297-FE0F",
		"value": "㊗️",
		"descriptor": "Japanese “Congratulations” Button"
	},
	"3299-FE0F": {
		"key": "3299-FE0F",
		"value": "㊙️",
		"descriptor": "Japanese “Secret” Button"
	},
	"33-FE0F-20E3": {
		"key": "33-FE0F-20E3",
		"value": "3️⃣",
		"descriptor": "Keycap Digit Three"
	},
	"34-FE0F-20E3": {
		"key": "34-FE0F-20E3",
		"value": "4️⃣",
		"descriptor": "Keycap Digit Four"
	},
	"35-FE0F-20E3": {
		"key": "35-FE0F-20E3",
		"value": "5️⃣",
		"descriptor": "Keycap Digit Five"
	},
	"36-FE0F-20E3": {
		"key": "36-FE0F-20E3",
		"value": "6️⃣",
		"descriptor": "Keycap Digit Six"
	},
	"37-FE0F-20E3": {
		"key": "37-FE0F-20E3",
		"value": "7️⃣",
		"descriptor": "Keycap Digit Seven"
	},
	"38-FE0F-20E3": {
		"key": "38-FE0F-20E3",
		"value": "8️⃣",
		"descriptor": "Keycap Digit Eight"
	},
	"39-FE0F-20E3": {
		"key": "39-FE0F-20E3",
		"value": "9️⃣",
		"descriptor": "Keycap Digit Nine"
	},
	"A9-FE0F": {
		"key": "A9-FE0F",
		"value": "©️",
		"descriptor": "Copyright"
	},
	"AE-FE0F": {
		"key": "AE-FE0F",
		"value": "®️",
		"descriptor": "Registered"
	}
}`)
	json.Unmarshal(byteValue, &Emojis)
}

// LookupEmoji - Lookup a single emoji definition
func LookupEmoji(emojiString string) (emoji Emoji, err error) {

	hexKey := utils.StringToHexKey(emojiString)

	// If we have a definition for this string we'll return it,
	// else we'll return an error
	if e, ok := Emojis[hexKey]; ok {
		emoji = e
	} else {
		err = fmt.Errorf("No record for \"%s\" could be found", emojiString)
	}

	return emoji, err
}

// LookupEmojis - Lookup definitions for each emoji in the input
func LookupEmojis(emoji []string) (matches []interface{}) {
	for _, emoji := range emoji {
		if match, err := LookupEmoji(emoji); err == nil {
			matches = append(matches, match)
		} else {
			matches = append(matches, err)
		}
	}

	return
}

// RemoveAll - Remove all emoji
func RemoveAll(input string) string {

	// Find all the emojis in this string
	matches := FindAll(input)

	for _, item := range matches {
		emo := item.Match.(Emoji)
		rs := []rune(emo.Value)
		for _, r := range rs {
			input = strings.Replace(input, string([]rune{r}), "", -1)
		}
	}

	// Remove and trim and left over whitespace
	return strings.TrimSpace(strings.Join(strings.Fields(input), " "))
	//return input
}
