package config

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/labstack/gommon/log"
)

func InitS3() *session.Session {
	ACCESS_KEY := GetEnv("ACCESS_KEY", "AKIAWEUURQ3TXBVRIDUG")
	SECRET_KEY := GetEnv("SECRET_KEY", "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")
	AWS_S3_REGION := GetEnv("AWS_S3_REGION", "ap-southeast-1")

	conn, err := session.NewSession(
		&aws.Config{
			Region: aws.String(AWS_S3_REGION),
			Credentials: credentials.NewStaticCredentials(
				ACCESS_KEY, SECRET_KEY, "",
			),
		},
	)
	if err != nil {
		log.Error("S3 Config error:", err)
	}
	return conn
}

func DoUpload(file multipart.FileHeader, fileName string) string {
	AWS_S3_BUCKET := GetEnv("AWS_S3_BUCKET", "belajar-be")
	Conn := InitS3()
	manager := s3manager.NewUploader(Conn)
	src, err := file.Open()
	if err != nil {
		log.Info(err)
	}
	defer src.Close()
	buffer := make([]byte, file.Size)
	src.Read(buffer)
	body, _ := file.Open()

	res, err := manager.Upload(
		&s3manager.UploadInput{
			Bucket: aws.String(AWS_S3_BUCKET),
			// ACL:         aws.String("public-read"),
			ContentType: aws.String(http.DetectContentType(buffer)),
			Key:         aws.String(fileName),
			Body:        body,
		},
	)
	fmt.Println(file.Filename)
	if err != nil {
		log.Info(res)
		log.Error("Upload error : ", err)
	}

	return res.Location
}
func DeleteItem(item *string) (string, error) {
	AWS_S3_BUCKET := GetEnv("AWS_S3_BUCKET", "belajar-be")
	Conn := InitS3()
	svc := s3.New(Conn)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    item,
	})
	if err != nil {
		return "", err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    item,
	})
	if err != nil {
		return "", err
	}

	return "sukses", nil
}
