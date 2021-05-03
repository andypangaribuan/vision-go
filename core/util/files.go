package util

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
	"vision/models"
	"vision/vis"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/03
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func (*VS) GetFileExtension(fileName string) string {
	idx := strings.LastIndex(fileName, ".")
	if idx < 0 {
		return ""
	}
	return fileName[idx+1:]
}


func (slf *VS) ScanFiles(recursively bool, limit int, dirPath string, extensions []string, condition func(file models.FileScan) (include bool)) (files []models.FileScan, err error) {
	files = make([]models.FileScan, 0)

	mapExtensions := make(map[string]interface{}, 0)
	for i:=0; i<len(extensions); i++ {
		extensions[i] = strings.ToLower(extensions[i])
		mapExtensions[extensions[i]] = nil
	}

	if len(dirPath) > 0 {
		length := len(dirPath)
		if dirPath[length-1:] == vis.Const.PathSeparator {
			dirPath = dirPath[:length-1]
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


	getFileModel := func(dirPath, fileName string) models.FileScan {
		model := models.FileScan{
			FileName:  fileName,
			Extension: slf.GetFileExtension(fileName),
			FilePath:  dirPath + vis.Const.PathSeparator + fileName,
			DirPath:   dirPath,
		}

		if idx := strings.LastIndex(fileName, "."); idx >= 0 {
			model.FileNameWithoutExtension = fileName[:idx]
		} else {
			model.FileNameWithoutExtension = fileName
		}

		if idx := strings.LastIndex(dirPath, "/"); idx >= 0 {
			model.DirName = dirPath[idx+1:]
		} else {
			model.DirName = dirPath
		}

		return model
	}


	var readDir func(dirPath string) error
	readDir = func(dirPath string) error {
		fileInfos, err := ioutil.ReadDir(dirPath)
		if err != nil {
			return errors.WithStack(err)
		}

		dirs := make([]string, 0)
		for _, file := range fileInfos {
			if len(files) >= limit && limit != -1 {
				break
			}

			if file.IsDir() {
				path := dirPath + vis.Const.PathSeparator + file.Name()
				dirs = append(dirs, path)
				continue
			}

			if file.Size() > 0 {
				canAddFile := canAddFileByExtension(file.Name())
				fileScan := getFileModel(dirPath, file.Name())
				if canAddFile && condition != nil {
					canAddFile = condition(fileScan)
				}
				if canAddFile {
					files = append(files, fileScan)
				}
			}
		}

		if !recursively {
			return nil
		}

		if len(files) < limit || limit == -1 {
			for _, dir := range dirs {
				err := readDir(dir)
				if err != nil {
					return err
				}

				if len(files) >= limit && limit != -1 {
					break
				}
			}
		}

		return nil
	}

	err = readDir(dirPath)
	return
}
