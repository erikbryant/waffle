package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/waffle/solver"
	"github.com/erikbryant/web"
	"io"
	"log"
	"math"
	"strings"
	"unicode"
)

var (
	DailyFiles  = []string{"daily1.txt", "daily2.txt"}
	DeluxeFiles = []string{"deluxe1.txt", "deluxe2.txt"}
	BaseURL     = "https://wafflegame.net/"
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
	// The JSON fails to unmarshal if non-printable unicode characters are present
	filtered := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, string(contents))

	var jsonObject map[string]interface{}

	err := json.Unmarshal([]byte(filtered), &jsonObject)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal json %s", err)
	}

	return jsonObject, nil
}

// calcSideLen returns the side length of the waffled square encoded by n
func calcSideLen(n int) int {
	// n is the difference of two squares (the whole board minus
	// the waffle holes):
	// n==4     4-0  2^2-0^2
	// n==8     9-1  3^2-1^2
	// n==21   25-4  5^2-2^2
	// n==40   49-9  7^2-3^2
	// n==65  81-16  9^2-4^2

	root := 0
	for {
		if root >= n {
			// n is not the difference of two squares
			break
		}
		candidate := n + root*root
		if algebra.IsSquare(candidate) {
			return int(math.Sqrt(float64(candidate)))
		}
		root++
	}

	panic(fmt.Errorf("could not find a length for: %d", n))
}

// insertSpaces returns a string with spaces added to represent the waffle holes
func insertSpaces(s string) string {
	out := ""

	sideLen := calcSideLen(len(s))

	i := 0
	for row := 0; row < sideLen; row++ {
		for col := 0; col < sideLen; col++ {
			if row%2 != 0 && col%2 != 0 {
				out += " "
				continue
			}
			out += s[i : i+1]
			i++
		}
	}

	return out
}

// signPuzzle downloads the given puzzle and returns its letters/colors signature
func signPuzzle(url string) (int, string) {
	msg, err := download(url)
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

	return number, waffle
}

// generateSignature returns letters/colors code suitable for pasting into the regress tests
func generateSignature(number int, waffle string) string {
	return fmt.Sprintf(`		{"%s", %d},`, waffle, number)
}

func main() {
	fmt.Printf("Welcome to decoder!\n\n")

	for _, file := range DailyFiles {
		number, waffle := signPuzzle(BaseURL + file)
		fmt.Println(generateSignature(number, waffle))
	}

	for _, file := range DeluxeFiles {
		number, waffle := signPuzzle(BaseURL + file)
		fmt.Println(generateSignature(number, waffle))
	}
}
