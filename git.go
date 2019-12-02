package main

import (
    "fmt"
    "os"
    "os/exec"
    "io"
    "io/ioutil"
    "strings"
    "log"
    "path/filepath"
)

func ExitIfNotGitDirectory() {
    // run git remote command
    cmd := exec.Command("git", "remote")

    stderr, err := cmd.StderrPipe()
    if err != nil {
        log.Fatal(err)
    }
    defer stderr.Close()

    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }

    result, err := ioutil.ReadAll(stderr)
    if err != nil {
        log.Fatal(err)
    }

    if err := cmd.Wait(); err != nil {
        fmt.Print(string(result))
        log.Fatal(err)
    }
}

func CommitMessage(message []byte, all bool) (string, error) {
    // save the commit message to temp file
    file, err := ioutil.TempFile("", "COMMIT_MESSAGE_")
    if err != nil {
        return "", err
    }
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
        return string(result), err
    }

    if err := cmd.Wait(); err != nil {
        return string(result), err
    }

    return string(result), nil
}

func CopyFile(dstName, srcName string) (written int64, err error) {
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

func Install() (string, error) {
    srcName, _ := exec.LookPath(os.Args[0])
    log.Printf("self path: %s", srcName)
    dstDir, err := execPath()
    if err != nil {
        return "", err
    }

    dstName := filepath.Join(dstDir, "git-cz")
    if _, err := CopyFile(dstName, srcName); err != nil {
        return dstName, err
    }

    return dstName, nil
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

