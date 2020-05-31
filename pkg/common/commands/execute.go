package commands

import (
	"bytes"
	"github.com/conplementAG/k8s-semantic-detective/pkg/common/logging"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// ExecuteCommand executes the command and logs the pipes the output to the logging system
func ExecuteCommand(command *exec.Cmd) string {
	out, _ := execute(command, true, true, true, false)
	return out
}

// ExecuteCommandWithSecretContents is similar to ExecuteCommandWithExcessiveOutput in behaviour, with the
// additional that even the command itself will not be logged / shown
func ExecuteCommandWithSecretContents(command *exec.Cmd) string {
	out, _ := execute(command, false, false, true, false)
	return out
}

func execute(command *exec.Cmd, pipeStdoutToLogs bool, showExecutedCommand bool, panicOnError bool, collectStdErrAsOutput bool) (string, error) {
	command.Stdin = os.Stdin

	commandStdoutPipe, _ := command.StdoutPipe()
	commandstderrPipe, _ := command.StderrPipe()

	var stdoutBuffer, stderrBuffer bytes.Buffer

	stdoutLogWritter := ioutil.Discard

	if pipeStdoutToLogs {
		stdoutLogWritter = newDebugLogWriter()
	}

	outputCollector := newOutputCollector()

	stdoutMultiwriter := io.MultiWriter(stdoutLogWritter, outputCollector, &stdoutBuffer)
	stderrMultiwriter := io.MultiWriter(newDebugLogWriter(), &stderrBuffer)

	if collectStdErrAsOutput {
		stderrMultiwriter = io.MultiWriter(newDebugLogWriter(), outputCollector, &stderrBuffer)
	}

	if showExecutedCommand {
		logging.Debugf("Executing: %s %s", command.Path, strings.Join(command.Args[1:], " "))
	}

	err := command.Start()

	if err != nil {
		logging.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var multiWritingSteps sync.WaitGroup
	multiWritingSteps.Add(2)

	go func() {
		io.Copy(stdoutMultiwriter, commandStdoutPipe)
		multiWritingSteps.Done()
	}()

	go func() {
		io.Copy(stderrMultiwriter, commandstderrPipe)
		multiWritingSteps.Done()
	}()

	commandError := command.Wait()
	multiWritingSteps.Wait()

	if panicOnError && commandError != nil {
		panic(commandError)
	}

	return outputCollector.output, commandError
}

type debugLogWriter struct{}

func newDebugLogWriter() *debugLogWriter {
	return &debugLogWriter{}
}

func (w *debugLogWriter) Write(p []byte) (int, error) {
	logging.Debug(string(p))

	return len(p), nil
}

type outputCollector struct {
	output string
}

func newOutputCollector() *outputCollector {
	return &outputCollector{}
}

func (w *outputCollector) Write(p []byte) (int, error) {
	w.output += string(p)

	return len(p), nil
}
