package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var Mode string
var zshHistoryFilePath string
var zshHistoryData []string

func init() {
	if Mode == "release" {
		// go build -ldflags "-X 'main.Mode=release'" -o rzh
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("사용자 홈 디렉토리를 찾을 수 없습니다: %v", err)
		}
		zshHistoryFilePath = filepath.Join(home, ".zsh_history")
	} else {
		zshHistoryFilePath = "test/.zsh_history"
	}
}

func main() {
	fmt.Println(zshHistoryFilePath)
	f, err := os.Open(zshHistoryFilePath)
	if err != nil {
		fmt.Println("error")
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	for input.Scan() {
		zshHistoryData = append(zshHistoryData, input.Text()) // 파일과 동일한 순서대로 저장됨
	}

	var metadataSepIdx = 0      // 메타데이터의 시작을 알리는 고정된 구분자
	var durationSepIdx = 12     // 시작 시간과-실행 시간을 구분하는 구분자
	var cmdSeparatorSepIdx = 14 // 모든 메타데이터와-실제 명령어를 구분하는 구분자

	for n, s := range zshHistoryData {
		lineNumber := n + 1
		timestamp := s[metadataSepIdx+2 : durationSepIdx]    // 명령어 실행 시작 시간
		duration := s[durationSepIdx+1 : cmdSeparatorSepIdx] // 명령어 지속 시간
		cmd := s[cmdSeparatorSepIdx+1:]                      // 명령어
		_, _, _, _ = lineNumber, duration, cmd, timestamp
	}

	// subcommand
	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ContinueOnError) // add 하위에서 사용할 수 있는 플래그 집합 선언

		// flag
		var enableBackup bool
		addCmd.BoolVar(&enableBackup, "backup", false, "")
		addCmd.BoolVar(&enableBackup, "b", false, "")

		var flagArgs []string
		var positionalArgs []string

		for _, arg := range os.Args[2:] {
			if strings.HasPrefix(arg, "-") {
				flagArgs = append(flagArgs, arg)
			} else {
				positionalArgs = append(positionalArgs, arg)
			}
		}

		if err := addCmd.Parse(flagArgs); err != nil {
			// flagArgs에 담긴 플래그들이 사전에 정의된 플래그가 맞는지 확인.
			// 만약 정의되지 않은 플래그가 있으면 에러를 반환.
			return
		}

		if len(positionalArgs) == 0 {
			fmt.Println("Error: 1개 이상의 인수 필요.")
			return
		}

		// main logic
		// ": 1754785989:0;foo"
		lastData := zshHistoryData[len(zshHistoryData)-1]
		lastTimestampStr := lastData[metadataSepIdx+2 : durationSepIdx]
		lastTimestamp, err := strconv.Atoi(lastTimestampStr)
		if err != nil {
			fmt.Println("error: string to int")
			return
		}

		var lines []string
		for i, a := range positionalArgs {
			a := fmt.Sprintf(": %d:0;%s", lastTimestamp+i+1, a) // 임의로 timestamp는 lastTimestamp + n으로 설정함.
			lines = append(lines, a)
		}
		content := strings.Join(lines, "\n")
		_ = content

		var input string
		var confirmed bool
		fmt.Print("Are you sure? (y/n): ")
		n, err := fmt.Scan(&input)
		_ = n
		if err != nil {
			fmt.Println("error")
		}
		if input == "y" || input == "Y" {
			confirmed = true
		}

		switch confirmed {
		case true:
			f, err := os.OpenFile(zshHistoryFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("error")
			}
			defer f.Close()
			if _, err := fmt.Fprintln(f, content); err != nil { // 파일에 쓰기(맨 뒤에 \n는 항상 포함해야함.)
				fmt.Println("error")
			}
			fmt.Println("추가됨.")
		case false:
			fmt.Println("추가되지 않음.")
		}
	case "remove":
		fmt.Println("remove")
	case "dedup":
		fmt.Println("dedup")
	case "backup":
		fmt.Println("backup")
	}
}
