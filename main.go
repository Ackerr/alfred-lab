package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func readLines(filePath string) (lines []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buffer := bufio.NewReader(file)
	for {
		value, _, err := buffer.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(value))
	}
	return lines, err
}

func run() {
	value := strings.Trim(wf.Args()[0], " ")
	// set gitlab base url
	if len(wf.Args()) > 1 {
		if err := wf.Config.Set("GITLAB_BASE_URL", value, false).Do(); err != nil {
			wf.Warn(err.Error(), "")
		}
		wf.SendFeedback()
		return
	}

	baseURL := wf.Config.GetString("GITLAB_BASE_URL", "")
	if baseURL == "" {
		wf.Warn("Not found GITLAB BASE URL", "Try use laburl to set it")
		wf.SendFeedback()
		return
	}

	home, _ := os.UserHomeDir()
	filePath := filepath.Join(home, ".config", "lab", ".projects")

	lines, err := readLines(filePath)
	if err != nil {
		wf.Warn(err.Error(), "")
		wf.SendFeedback()
		return
	}
	for _, line := range lines {
		url := strings.Join([]string{baseURL, "/", line, "/-/merge_requests"}, "")
		wf.NewItem(line).Valid(true).Arg(url).UID(line)
	}

	if value != "" {
		wf.Filter(value)
	}

	wf.WarnEmpty("No matching project found", "Try a different query?")
	wf.SendFeedback()
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
