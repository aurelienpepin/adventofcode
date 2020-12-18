package main

import (
	"fmt"
	"github.com/aurelienpepin/adventofcode/2020/day18"
	"time"
)

const (
	ADVENT_OF_CODE_URL = "https://adventofcode.com/2020/day/"
	ADVENT_OF_CODE_INPUT = "/input"
	INPUT_FOLDER = "2020/inputs/"
)

func main() {
	start := time.Now()
	fmt.Println(day18.Part2())
	fmt.Printf("Time elapsed: %s\n", time.Since(start))
}

// If `force` is false, the input file won't be downloaded if a local file
// with the same name already exists
/*func DownloadDailyInput(force bool) error {
	today := strconv.Itoa(time.Now().Day())
	fileUrl := fmt.Sprint(ADVENT_OF_CODE_URL, today, ADVENT_OF_CODE_INPUT)
	filePath := fmt.Sprint(INPUT_FOLDER, "day", today)

	// Get input file
	response, err := http.Get(fileUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Check the existing file
	if _, err := os.Open(filePath); !force && err == nil {
		panic("file already exists")
	}

	// Write input file locally
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}*/