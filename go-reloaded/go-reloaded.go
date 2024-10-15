package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// Create the input.txt file programmatically
	createInputFile()

	// Open the input file
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Read the input file line by line
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	for scanner.Scan() {
		line := scanner.Text()
		processedLine := processText(line)
		writer.WriteString(processedLine + "\n")
	}

	// Flush the writer to ensure all data is written to the output file
	writer.Flush()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
}

// createInputFile creates the input.txt file with test cases programmatically
func createInputFile() {
	content := `it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
There is no greater agony than bearing a untold story inside you.
Punctuation tests are ... kinda boring ,what do you think ?
I am exactly how they describe me: ' awesome '
As Elton John said: ' I am the most well-known homosexual in the world '
This is amazing (up, 3)
make IT SMALL (low, 2) but CAPTURE the moment
capitalize these three words (cap, 3)`

	_ = os.WriteFile("input.txt", []byte(content), 0o644)
}

// processText processes the text as per specified transformations
func processText(text string) string {
	// Apply multi-word transformations like (up, N), (low, N), (cap, N) using the instances function
	text = instances(text)

	words := strings.Fields(text)
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Check if the word is followed by a specific transformation tag
		if i+1 < len(words) {
			nextWord := words[i+1]

			switch nextWord {
			case "(hex)":
				word = hexToDecimal(word)
				i++ // Skip the transformation tag
			case "(bin)":
				word = binToDecimal(word)
				i++ // Skip the transformation tag
			case "(up)":
				word = strings.ToUpper(word)
				i++ // Skip the transformation tag
			case "(low)":
				word = strings.ToLower(word)
				i++ // Skip the transformation tag
			case "(cap)":
				word = capitalize(word)
				i++ // Skip the transformation tag
			}
		}

		// Handle grammar correction (a -> an)
		if word == "a" && i+1 < len(words) && isVowel(words[i+1][0]) {
			word = "an"
		}

		result = append(result, word)
	}

	// Join the result words back together and fix punctuation
	joinedResult := strings.Join(result, " ")
	joinedResult = formatPunctuation(joinedResult)
	return fixQuotes(joinedResult)
}

// instances modifies N previous words based on (up, N), (low, N), (cap, N)
func instances(content string) string {
	// Regular expression to match instances like (up, 2), (low, 1), (cap, 3)
	re := regexp.MustCompile(`\((up|low|cap),\s?([0-9]+)\)`)
	// Keep processing while there are matches
	for {
		// Find the first match
		match := re.FindStringSubmatchIndex(content)
		if match == nil {
			// No more matches
			break
		}
		// Extract the directive (up, low, cap)
		direction := content[match[2]:match[3]]
		// Extract the number of words to modify
		number, _ := strconv.Atoi(content[match[4]:match[5]])
		// Get all the words before the match
		words := strings.Fields(content[:match[0]])
		// Ensure that the number of words to modify doesn't exceed available words
		if len(words) < number {
			number = len(words)
		}
		// Modify the previous N words based on the directive
		for i := len(words) - number; i < len(words); i++ {
			switch direction {
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "low":
				words[i] = strings.ToLower(words[i])
			case "cap":
				words[i] = capitalize(words[i]) // Title case
			}
		}
		// Rebuild the content by joining the modified words
		content = strings.Join(words, " ") + content[match[1]:] // Append the rest after the match
	}
	return content
}

// hexToDecimal converts a hexadecimal string to its decimal equivalent
func hexToDecimal(word string) string {
	hexValue, err := strconv.ParseInt(word, 16, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(hexValue, 10)
}

// binToDecimal converts a binary string to its decimal equivalent
func binToDecimal(word string) string {
	binValue, err := strconv.ParseInt(word, 2, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(binValue, 10)
}

// capitalize capitalizes the first letter of a word
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// isVowel checks if a character is a vowel or 'h'
func isVowel(c byte) bool {
	return strings.ContainsAny(string(unicode.ToLower(rune(c))), "aeiouh")
}

// formatPunctuation formats punctuation correctly as per the rules
func formatPunctuation(text string) string {
	// Remove extra spaces around punctuation
	text = strings.ReplaceAll(text, " ,", ",")
	text = strings.ReplaceAll(text, " .", ".")
	text = strings.ReplaceAll(text, " !", "!")
	text = strings.ReplaceAll(text, " ?", "?")
	text = strings.ReplaceAll(text, " :", ":")
	text = strings.ReplaceAll(text, " ;", ";")
	// Handle ellipses and ensure they stick to the left word only
	text = strings.ReplaceAll(text, " ...", "...")
	text = strings.ReplaceAll(text, "... ", "...")
	text = strings.ReplaceAll(text, "...", "...") // This keeps ellipses intact
	// Ensure no space before the ellipsis
	text = strings.ReplaceAll(text, " ...", "...")
	// Add a space after the ellipsis if not present
	text = strings.ReplaceAll(text, "... ", "... ")
	text = strings.ReplaceAll(text, "...", "... ")
	// Add a space after a comma if it's not present
	text = strings.ReplaceAll(text, ",", ", ")
	// Clean up any double spaces
	text = strings.ReplaceAll(text, "  ", " ")
	return text
}

// fixQuotes ensures quotes are properly formatted without extra spaces
func fixQuotes(text string) string {
	// Regular expression to match single quotes and remove spaces between the quotes and the inner content
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(text, "'$1'")
}
