package service

import (
	"errors"
	logUtils "github.com/aaronchen2k/deeptest/internal/pkg/lib/log"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/dir"
	"github.com/snowlyg/helper/str"
	"go.uber.org/zap"
)

var (
	ErrEmpty = errors.New("请上传正确的文件")
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

// UploadFile 上传文件
func (s *FileService) UploadFile(ctx iris.Context, fh *multipart.FileHeader) (iris.Map, error) {
	filename, err := GetFileName(fh.Filename)
	if err != nil {

		return nil, err
	}
	path := filepath.Join(dir.GetCurrentAbPath(), "static", "upload", "images")
	err = dir.InsureDir(path)
	if err != nil {
		logUtils.Errorf("文件上传失败", zap.String("dir.InsureDir", err.Error()))
		return nil, err
	}
	_, err = ctx.SaveFormFile(fh, filepath.Join(path, filename))
	if err != nil {
		logUtils.Errorf("文件上传失败", zap.String("ctx.SaveFormFile", "保存文件到本地"))
		return nil, err
	}

	qiniuUrl := ""
	path = s.GetPath(filename)
	// if libs.ConfiserverConsts.Qiniu.Enable {
	// 	var key string
	// 	var hash string
	// 	key, hash, err = libs.Upload(filepath.Join(libs.CWD(), path), filename)
	// 	if err != nil {
	// 		logUtils.Errorf("文件上传失败", zap.String("ctx.SaveFormFile", "图片上传七牛云失败"))
	// 		ctx.JSON(consts.Response{Code: consts.SystemErr.Code, Data: nil, Msg: err.Error()})
	// 		return
	// 	}

	// 	logUtils.Debugf("文件上传失败", zap.String("key", key), zap.String("hash", hash))

	// 	if key != "" {
	// 		qiniuUrl = fmt.Sprintf("%s/%s", libs.Config.Qiniu.Host, key)
	// 	}
	// }

	return iris.Map{"local": path, "qiniu": qiniuUrl}, nil
}

// GetFileName 获取文件名称
func GetFileName(name string) (string, error) {
	fns := strings.Split(strings.TrimLeft(name, "./"), ".")
	if len(fns) != 2 {
		logUtils.Errorf("文件上传失败", zap.String("trings.Split", name))
		return "", ErrEmpty
	}
	ext := fns[1]
	md5, err := dir.MD5(name)
	if err != nil {
		logUtils.Errorf("文件上传失败", zap.String("dir.MD5", err.Error()))
		return "", err
	}
	return str.Join(md5, ".", ext), nil
}

// GetPath 获取文件路径
func (s *FileService) GetPath(filename string) string {
	return filepath.Join("upload", "images", filename)
}
