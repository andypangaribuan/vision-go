package util

import (
	"bufio"
	"github.com/andypangaribuan/vision-go/models"
	"github.com/andypangaribuan/vision-go/vis"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (*utilStruct) GetFileExtension(fileName string) string {
	idx := strings.LastIndex(fileName, ".")
	if idx < 0 {
		return ""
	}
	return fileName[idx+1:]
}


func (slf *utilStruct) ScanDir(includeDirectory, recursively bool, limit int, directoryPath string, extensions []string, condition func(model models.DirScan) (include bool)) (fns chan func()(*models.DirScan, error)) {
	fns = make(chan func()(*models.DirScan, error))
	totalModel := 0

	mapExtensions := make(map[string]interface{}, 0)
	for i:=0; i<len(extensions); i++ {
		extensions[i] = strings.ToLower(extensions[i])
		mapExtensions[extensions[i]] = nil
	}

	if len(directoryPath) > 0 {
		length := len(directoryPath)
		if directoryPath[length-1:] == vis.Const.PathSeparator {
			directoryPath = directoryPath[:length-1]
		}
	}


	canAddFileByExtension := func(fileName string) bool {
		if len(extensions) == 0 {
			return true
		}

		ext := slf.GetFileExtension(fileName)
		ext = strings.ToLower(ext)
		if _, ok := mapExtensions[ext]; ok {
			return true
		}
		return false
	}


	getModel := func(isDir bool, dirPath, fileName string) (model models.DirScan) {
		model = models.DirScan{
			IsDirectory: isDir,
			DirPath:     dirPath,
		}

		if !isDir {
			model.FileName = fileName
			model.FileExtension = slf.GetFileExtension(fileName)
			model.FilePath = dirPath + vis.Const.PathSeparator + fileName

			if idx := strings.LastIndex(fileName, "."); idx >= 0 {
				model.FileNameWithoutExtension = fileName[:idx]
			} else {
				model.FileNameWithoutExtension = fileName
			}

			if info, err := os.Stat(model.FilePath); err == nil {
				model.FileSize = info.Size()
			}
		}

		if idx := strings.LastIndex(dirPath, vis.Const.PathSeparator); idx >= 0 {
			model.DirName = dirPath[idx+1:]
		} else {
			model.DirName = dirPath
		}
		return
	}


	var doScan func(dirPath string) error
	doScan = func(dirPath string) error {
		infos, err := ioutil.ReadDir(dirPath)
		if err != nil {
			err = errors.WithStack(err)
			fns <- func() (*models.DirScan, error) { return nil, err }
			return err
		}

		sort.Slice(infos, func(i, j int) bool {
			if infos[i].IsDir() && !infos[j].IsDir() {
				return false
			}
			return infos[i].Name() < infos[j].Name()
		})

		for _, info := range infos {
			if limit != -1 && totalModel >= limit {
				break
			}

			if info.IsDir() {
				path := dirPath + vis.Const.PathSeparator + info.Name()
				if includeDirectory {
					model := getModel(true, path, "")
					canAdd := true
					if condition != nil {
						canAdd = condition(model)
					}
					if canAdd {
						fns <- func() (*models.DirScan, error) { return &model, nil }
						totalModel++
					}
				}
				if recursively {
					if err := doScan(path); err != nil {
						return err
					}
				}
			} else {
				canAdd := canAddFileByExtension(info.Name())
				model := getModel(false, dirPath, info.Name())
				if canAdd && condition != nil {
					canAdd = condition(model)
				}
				if canAdd {
					fns <- func() (*models.DirScan, error) { return &model, nil }
					totalModel++
				}
			}
		}
		return nil
	}

	go func() {
		_ = doScan(directoryPath)
		close(fns)
	}()
	return
}


func (slf *utilStruct) CountFileLines(filePath string) (count int, err error) {
	file, _err := os.Open(filePath)
	if _err != nil {
		return 0, errors.WithStack(_err)
	}
	defer file.Close()

	switch info, err := file.Stat(); {
	case err != nil:
		return 0, errors.WithStack(err)
	case info.Size() == 0:
		return 0, nil
	}

	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		count++
		if err != nil {
			return count, nil
		}
	}
}


func (slf *utilStruct) ScanFileLines(filePath string) (scan chan func()(isLast bool, index int, line string), err error) {
	canClose := true
	scan = make(chan func()(isLast bool, index int, line string))
	defer func() {
		if canClose {
			close(scan)
		}
	}()

	file, _err := os.Open(filePath)
	if _err != nil {
		return scan, errors.WithStack(_err)
	}
	defer func() {
		if canClose {
			file.Close()
		}
	}()

	switch info, err := file.Stat(); {
	case err != nil:
		return scan, errors.WithStack(err)
	case info.Size() == 0:
		return scan, nil
	}

	doScan := func() {
		reader := bufio.NewReader(file)
		index := -1
		for {
			index++
			idx := index //prevent updated value from var index
			line, err := reader.ReadString('\n')
			line = strings.Replace(line, "\n", "", 1)

			if err == nil {
				scan <- func() (bool, int, string) {
					return false, idx, line
				}
			} else {
				scan <- func() (bool, int, string) {
					return true, idx, line
				}
				break
			}
		}
	}

	canClose = false
	go func() {
		doScan()
		close(scan)
		file.Close()
	}()
	return
}
