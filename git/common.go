package git

import (
	"github.com/chmodke/gitmm/log"
	"github.com/chmodke/gitmm/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Command(localRepo string, args []string) bool {
	builder := &util.CmdBuilder{}
	builder.Add("git")
	for _, arg := range args {
		builder.Add(arg)
	}
	out, ret := util.GetOut(util.Execute(localRepo, builder.Build()))
	log.Consoleln(out)
	return ret
}

func GetWorkDir(workDir string) (string, error) {
	var localDir string
	if filepath.IsAbs(workDir) {
		localDir = workDir
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		localDir = filepath.Join(pwd, workDir)
	}

	if PathExists(localDir) {
		return localDir, nil
	}

	err := os.MkdirAll(localDir, 0750)
	if err != nil {
		return "", err
	}
	return localDir, nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func FindGit(dirPth string) (files []string, err error) {
	files = make([]string, 0, 10)

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		if isGit(filepath.Join(dirPth, fi.Name())) {
			files = append(files, fi.Name())
		}
	}
	return files, nil
}

func isGit(dirPth string) bool {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return false
	}
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		if strings.EqualFold(".git", fi.Name()) {
			return true
		}
	}
	return false
}
