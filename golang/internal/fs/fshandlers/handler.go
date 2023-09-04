package fshandlers

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

const rootPath = "./storage"

type fsHandler struct {
	BasePath   string
	PublicPath string
}

func clean(path string) string {
	path = strings.TrimSuffix(path, "/")
	return strings.TrimPrefix(path, "/")
}

func concat(paths ...string) string {
	var cleaned []string
	for _, path := range paths {
		cleaned = append(cleaned, clean(path))
	}

	return strings.Join(cleaned, "/")
}

func (h fsHandler) GetBasePath() string {
	return clean(h.BasePath)
}

func (h fsHandler) GetFilePath(fileName string) string {
	return concat(h.GetBasePath(), fileName)
}

func (h fsHandler) ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(h.GetFilePath(fileName))
}

func (h fsHandler) WriteFile(fileName string, buffer []byte) error {
	err := os.Mkdir(h.GetBasePath(), os.ModePerm)
	if err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			return err
		}
	}

	err = os.WriteFile(h.GetFilePath(fileName), buffer, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (h fsHandler) GetPublicUrl(fileName string) string {
	pubUrl := viper.GetString("APPLICATION_URL")

	return concat(pubUrl, h.PublicPath, fileName)
}
