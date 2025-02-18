package storage

import (
	"io"
	"net/url"
	"os"
	"path"
	"simpledouyin/config"
	"simpledouyin/logging"

	"github.com/sirupsen/logrus"
)

type FSStorage struct {
}

func (f FSStorage) Upload(fileName string, content io.Reader) (output *PutObjectOutput, err error) {
	methodFields := logrus.Fields{
		"function":  "FSStorage.Upload",
		"file_name": fileName,
	}
	logger := logging.Logger.WithFields(methodFields)
	logger.Debug("Process start")

	all, err := io.ReadAll(content)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Debug("failed reading content")
		return nil, err
	}
	filePath := path.Join(config.EnvConfig.LOCAL_FS_LOCATION, fileName)
	dir := path.Dir(filePath)
	err = os.MkdirAll(dir, os.FileMode(0755))
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Debug("failed writing creating directory before writing file")
		return nil, err
	}
	err = os.WriteFile(filePath, all, os.FileMode(0755))
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err,
		}).Debug("failed writing content to file")
		return nil, err
	}
	return &PutObjectOutput{}, nil
}

func (f FSStorage) GetLink(fileName string) (string, error) {
	return url.JoinPath(config.EnvConfig.LOCAL_FS_BASEURL, fileName)
}
