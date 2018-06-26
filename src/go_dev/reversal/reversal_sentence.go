// reserver_sentence
package reversal

import (
	"strings"
)

func ReserverSentence(sentence string) string {
	no_space_sentence := strings.TrimSpace(sentence)
	sentence_array := strings.Split(no_space_sentence, " ")
	sentence_array_len := len(sentence_array)
	new_sentence_array := []string{}
	for i, _ := range sentence_array {
		last_word := sentence_array[sentence_array_len-i-1] //因为索引从0开始，所以这里多减去1
		new_sentence_array = append(new_sentence_array, last_word)
	}
	var new_sentence string
	for i, v := range new_sentence_array {
		if i == 0 {
			new_sentence = new_sentence + v
		} else {
			new_sentence = new_sentence + " " + v
		}
	}
	return new_sentence
}
