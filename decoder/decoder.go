package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/erikbryant/waffle/solver"
	"github.com/erikbryant/web"
)

// download returns the response body from requesting the given URL
func download(url string) (string, error) {
	response, err := web.Request2(url, map[string]string{})
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		return "", fmt.Errorf("error requesting data: %d", response.StatusCode)
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(contents), nil
}

// decodeBase64 returns the plain text of the given base64 string
func decodeBase64(msg string) ([]byte, error) {
	plainText, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return nil, err
	}

	return plainText, err
}

// parseJson returns the JSON representation of the contents
func parseJson(contents []byte) (map[string]interface{}, error) {
	// The JSON fails to unmarshal if these special characters are present.
	// Filter them out.
	filtered := []byte{}
	for _, b := range contents {
		if b <= 0x1d {
			continue
		}
		filtered = append(filtered, b)
	}

	var jsonObject map[string]interface{}

	err := json.Unmarshal(filtered, &jsonObject)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal json %s", err)
	}

	return jsonObject, nil
}

// insertSpaces returns a string with spaces added to represent the waffle holes
func insertSpaces(s string) string {
	out := ""

	for i := range s {
		if i == 6 || i == 7 || i == 14 || i == 15 {
			out += " "
		}
		out += s[i : i+1]
	}

	return out
}

// generateSignature returns code suitable for pasting into the regress tests
func generateSignature(number int, waffle string) string {
	return fmt.Sprintf(`		{"%s", %d},`, waffle, number)
}

func main() {
	fmt.Printf("Welcome to decoder!\n\n")

	msg, err := download("https://wafflegame.net/daily1.txt")
	if err != nil {
		log.Fatal(err)
	}

	plainText, err := decodeBase64(msg)
	if err != nil {
		log.Fatal(err)
	}

	jsonMap, err := parseJson(plainText)
	if err != nil {
		log.Fatal(err)
	}

	number := int(jsonMap["number"].(float64))
	puzzle := jsonMap["puzzle"].(string)
	solution := jsonMap["solution"].(string)

	puzzle = strings.ToLower(puzzle)
	puzzle = insertSpaces(puzzle)

	solution = strings.ToLower(solution)
	solution = insertSpaces(solution)

	waffle := solver.ParseSolution(puzzle, solution)

	sig := generateSignature(number, waffle)

	fmt.Println(sig)
}
