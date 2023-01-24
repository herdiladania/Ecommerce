package helper

import (
	"e-commerce/config"
	"errors"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// const charset = "abcdefghijklmnopqrstuvwxyz" +
// 	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// var seededRand *rand.Rand = rand.New(
// 	rand.NewSource(time.Now().UnixNano()))

// func autoGenerate(length int, charset string) string {
// 	b := make([]byte, length)
// 	for i := range b {
// 		b[i] = charset[seededRand.Intn(len(charset))]
// 	}
// 	return string(b)
// }

// func String(length int) string {
// 	return autoGenerate(length, charset)
// }

// func UploadImage(c echo.Context) (string, error) {

// 	file, fileheader, err := c.Request().FormFile("file")
// 	if err != nil {
// 		fmt.Print("\n\nfailed get pah file. err = ", err)
// 		return "", err
// 	}

// 	randomStr := String(20)

// 	godotenv.Load("local.env")

// 	s3Config := &aws.Config{
// 		Region:      aws.String(os.Getenv("AWS_REGION")),
// 		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_IAM"), os.Getenv("SECRET_KEY_IAM"), ""),
// 	}
// 	s3Session := session.New(s3Config)

// 	uploader := s3manager.NewUploader(s3Session)

// 	input := &s3manager.UploadInput{
// 		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),
// 		Key:         aws.String("posting/" + randomStr + "-" + fileheader.Filename),
// 		Body:        file,
// 		ContentType: aws.String("image/jpg"),
// 	}
// 	res, err := uploader.UploadWithContext(context.Background(), input)
// 	fmt.Println("\n\nerror upload to s3. err = ", err)
// 	return res.Location, err
// }

var ObjectURL string = "https://ecommercegroup7.s3.ap-southeast-1.amazonaws.com/"

func UploadProfilePhotoS3(file multipart.FileHeader, email string) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("ecommercegroup7"),
		Key:    aws.String("files/user/" + email + "/profile-photo" + ext),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload profile photo")
	}
	path := ObjectURL + "files/user/" + email + "/profile-photo" + ext
	return path, nil
}

func UploadProductImageS3(file multipart.FileHeader, userID int) (string, error) {
	s3Session := config.S3Config()
	uploader := s3manager.NewUploader(s3Session)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// ext := filepath.Ext(file.Filename)

	cnv := strconv.Itoa(userID)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("ecommercegroup7"),
		Key:    aws.String("files/post/" + cnv + "/" + file.Filename),
		Body:   src,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", errors.New("problem with upload post photo")
	}
	path := ObjectURL + "files/post/" + cnv + "/" + file.Filename
	return path, nil
}
