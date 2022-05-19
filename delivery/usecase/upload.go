package usecase

import (
	"errors"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"

	"event/config"
)

func Upload(file *multipart.FileHeader) (string, error) {

	if file.Size >= 1000*1000 {
		return "", errors.New("file terlalu besar maks 1 mb")
	}
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return "", err
	}
	fileByte, _ := ioutil.ReadAll(src)
	fileType := http.DetectContentType(fileByte)
	if fileType != "image/png" {
		return "", errors.New("required PNG")
	}
	fileName := strconv.Itoa(int(time.Now().Unix())) + ".png"
	image := config.DoUpload(*file, fileName)
	return image, nil
}
func DeleteImg(img string) error {
	file := strings.Replace(img, "https://belajar-be.s3.ap-southeast-1.amazonaws.com/", "", 1)

	_, err := config.DeleteItem(&file)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
