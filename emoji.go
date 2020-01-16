package emoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
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
	// Work out where we are in relation to the caller
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	// Open the Emoji definition JSON and Unmarshal into map
	jsonFile, err := os.Open(path.Dir(filename) + "/data/emoji.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	byteValue, e := ioutil.ReadAll(jsonFile)
	if e != nil {
		panic(e)
	}

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
