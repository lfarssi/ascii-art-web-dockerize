package ascii

import (
	"bufio"
	"net/http"
	"os"
	"strings"
)

func TraitmentData(w http.ResponseWriter, bnr string, arg string) string {
	banner := bnr
	fileName := "./banners/" + banner + ".txt"
	
	// Open the ASCII art file
	file, err := os.Open(fileName)
	if err != nil {
		http.Error(w, "Error opening the file", http.StatusInternalServerError)
		return ""
	}
	defer file.Close()

	var asciiArt []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asciiArt = append(asciiArt, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading the file", http.StatusInternalServerError)
		return ""
	}

	var result string
	lines := strings.Split(arg, "\n")
	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}
		// Iterate over each row of the ASCII art (0 to 7, for the 8 rows)
		for i := 1; i <= 8; i++ {
			for _, r := range line {
				// Ensure the character is within the valid ASCII range
				if r < 32 || r > 126 {
					http.Error(w, "Please enter a valid character between ASCII code 32 and 126", http.StatusBadRequest)
					return ""
				}
				index := 9*(int(r)-32) + i
				result += asciiArt[index]
			}
			result += "\n" // Add newline after finishing the current row of the line
		}
	}
	return result
}

func BannerExists(banner string) bool {
	// Check if the provided banner exists (this is a placeholder function)
	// For example, compare the banner string to a list of supported banners
	supportedBanners := []string{"standard", "shadow", "thinkertoy"}
	for _, b := range supportedBanners {
		if banner == b {
			return true
		}
	}
	return false
}
