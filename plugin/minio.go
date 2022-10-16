package plugin

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func InitMinio() *minio.Client {
	endpoint := viper.GetString("minio.endpoint")
	accessKey := viper.GetString("minio.accessKey")
	secretKey := viper.GetString("minio.secretKey")
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up

	return minioClient
}
