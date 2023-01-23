package helper

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func autoGenerate(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return autoGenerate(length, charset)
}

func UploadImage(c echo.Context) (string, error) {

	file, fileheader, err := c.Request().FormFile("file")
	if err != nil {
		fmt.Print("\n\nfailed get pah file. err = ", err)
		return "", err
	}

	randomStr := String(20)

	godotenv.Load("local.env")

	s3Config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_IAM"), os.Getenv("SECRET_KEY_IAM"), ""),
	}
	s3Session := session.New(s3Config)

	uploader := s3manager.NewUploader(s3Session)

	input := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:         aws.String("posting/" + randomStr + "-" + fileheader.Filename),
		Body:        file,
		ContentType: aws.String("image/jpg"),
	}
	res, err := uploader.UploadWithContext(context.Background(), input)
	fmt.Println("\n\nerror upload to s3. err = ", err)
	return res.Location, err
}
