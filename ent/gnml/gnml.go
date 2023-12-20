package gnml

import (
	"strings"
)

type TextPart struct {
	SecNum  int
	Content string
	Start   int
	End     int
	Length  int
}

func SplitText(text string, chunkSize, overlap int) []TextPart {
	var parts []TextPart
	var i, prevI, count int
	for {
		// Ensure that 'i' is always advancing to avoid an infinite loop
		if i != 0 && i <= prevI {
			i = prevI + 1
		}

		// Determine the end of the current chunk
		end := i + chunkSize
		if end >= len(text) {
			part := TextPart{
				SecNum:  count,
				Content: text[i:],
				Start:   i,
				End:     len(text),
				Length:  len(text) - i,
			}

			parts = append(parts, part)
			return parts
		}

		// Find the best place to split the chunk
		partEnd := findSplitPoint(text[i:end])
		actualEnd := i + partEnd

		part := TextPart{
			SecNum:  count,
			Content: text[i:actualEnd],
			Start:   i,
			End:     actualEnd,
			Length:  actualEnd - i,
		}
		parts = append(parts, part)

		prevI, i = i, actualEnd-overlap
		count++
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
