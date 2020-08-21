package git

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func InstallSubCmd(srcFilePath, subCmdName string) (string, error) {
	dstDir, err := execPath()
	if err != nil {
		return "", err
	}

	subCmdFileName := "git-" + subCmdName

	dstFilePath := filepath.Join(dstDir, subCmdFileName)
	if _, err := copyFile(dstFilePath, srcFilePath); err != nil {
		return dstFilePath, err
	}

	return dstFilePath, nil
}

func IsCurrentDirectoryGitRepo() (bool, error) {
	// run git remote command
	cmd := exec.Command("git", "remote")

	if stderr, err := cmd.StderrPipe(); err != nil {
		return false, err
	} else {
		defer stderr.Close()

		if err := cmd.Start(); err != nil {
			return false, err
		}

		if result, err := ioutil.ReadAll(stderr); err != nil {
			return false, err
		} else {

			if err := cmd.Wait(); err != nil {
				return false, fmt.Errorf("%s", string(result))
			}

			return true, nil
		}
	}
}

func CommitMessage(message []byte, all bool) (string, error) {
	// save the commit message to temp file
	if file, err := ioutil.TempFile("", "COMMIT_MESSAGE_"); err != nil {
		return "", err
	} else {
		defer os.Remove(file.Name())

		if _, err := file.Write(message); err != nil {
			return "", err
		}

		// run git commit command
		cmd := exec.Command("git", "commit", "-F")
		cmd.Args = append(cmd.Args, file.Name())
		if all {
			cmd.Args = append(cmd.Args, "-a")
		}

		if stdout, err := cmd.StdoutPipe(); err != nil {
			return "", err
		} else {
			defer stdout.Close()

			if err := cmd.Start(); err != nil {
				return "", err
			}

			if result, err := ioutil.ReadAll(stdout); err != nil {
				return string(result), err
			} else {
				if err := cmd.Wait(); err != nil {
					return string(result), err
				} else {
					return string(result), nil
				}
			}
		}
	}
}

func copyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func execPath() (string, error) {
	cmd := exec.Command("git", "--exec-path")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	defer stdout.Close()

	if err := cmd.Start(); err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	if err := cmd.Wait(); err != nil {
		return "", err
	}

	return strings.TrimSpace(string(result)), nil
}
