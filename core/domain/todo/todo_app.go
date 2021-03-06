package todo

import (
	"github.com/phodal/coca/core/domain/gitt"
	"github.com/phodal/coca/core/domain/todo/astitodo"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type TodoApp struct {
}

func NewTodoApp() TodoApp {
	return *&TodoApp{

	}
}

type TodoDetail struct {
	Date     string
	FileName string
	Author   string
	Line     string
	Assignee string
	Message  []string
}

func (a TodoApp) AnalysisPath(path string) []TodoDetail {
	var todoList []TodoDetail = nil
	todos, err := astitodo.Extract(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, todo := range todos {
		lineOutput := runGitGetLog(todo.Line, todo.Filename)

		todoDetail := &TodoDetail{
			Date:     "",
			FileName: todo.Filename,
			Author:   "",
			Line:     strconv.Itoa(todo.Line),
			Assignee: todo.Assignee,
			Message:  todo.Message,
		}
		commitMessages := gitt.BuildMessageByInput(lineOutput)

		if len(commitMessages) > 0 {
			commit := commitMessages[0]
			todoDetail.Date = commit.Date
			todoDetail.Author = commit.Author
		}
		todoList = append(todoList, *todoDetail)
	}

	return todoList
}

func runGitGetLog(line int, fileName string) string {
	// git log -1 -L2:README.md --pretty="format:[%h] %aN %ad %s" --date=short  --numstat
	historyArgs := []string{"log", "-1", "-L" + strconv.Itoa(line) + ":" + fileName, "--pretty=format:[%h] %aN %ad %s", "--date=short", "--numstat", "--summary"}
	cmd := exec.Command("git", historyArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	split := strings.Split(string(out), "\n")
	output := split[0] + "\n "
	return output
}
