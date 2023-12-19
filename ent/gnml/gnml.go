package gnml

import (
	"strings"
)

func SplitText(text string, chunkSize, overlap int) []string {
	var chunks []string
	var i, prevI int
	for {
		// Ensure that 'i' is always advancing to avoid an infinite loop
		if i != 0 && i <= prevI {
			i = prevI + 1
		}

		// Determine the end of the current chunk
		end := i + chunkSize
		if end >= len(text) {
			chunks = append(chunks, text[i:])
			return chunks
		}

		// Find the best place to split the chunk
		chunkEnd := findSplitPoint(text[i:end])
		actualEnd := i + chunkEnd

		chunk := text[i:actualEnd]
		chunks = append(chunks, chunk)

		prevI, i = i, actualEnd-overlap
	}
}

// findSplitPoint finds the best split point in a chunk of text.
func findSplitPoint(chunk string) int {
	threshold := len(chunk) / 2
	if threshold < 10 {
		return len(chunk)
	}
	delimiters := []string{"\r\r\n\r\r\n", "\r\n\r\n", "\n\n", ". ", " ", "\n"}
	for _, delimiter := range delimiters {
		if idx := strings.LastIndex(chunk, delimiter); idx != -1 {
			if len(chunk)-idx < threshold {
				return idx + len(delimiter)
			}
		}
	}
	return len(chunk)
}
