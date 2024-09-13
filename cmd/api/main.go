package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-chi/chi"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/auth"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/bucket"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/files"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/folders"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/queue"
	"github.com/wagnersilvafilho/aprendagolang/imersao/internal/users"
	"github.com/wagnersilvafilho/aprendagolang/imersao/pkg/database"
)

func main() {
	db, b, qc := getSessions()

	r := chi.NewRouter()

	// define endpoints
	r.Post("/auth", auth.HandleAuth(
		func(login, pass string) (auth.Authenticated, error) {
			return users.Authenticate(login, pass)
		},
	))

	files.SetRoutes(r, db, b, qc)
	folders.SetRoutes(r, db)
	users.SetRoutes(r, db)

	// start server
	http.ListenAndServe(os.Getenv(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))), r)
}

func getSessions() (*sql.DB, *bucket.Bucket, *queue.Queue) {
	// create new database connection
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	qcfg := queue.RabbitMQConfig{
		URL:       os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	// create new queue
	qc, err := queue.New(queue.RabbitMQ, qcfg)
	if err != nil {
		log.Fatal(err)
	}

	// bucket config
	bcfg := bucket.AwsConfig{
		Config: &aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_KEY"),
				os.Getenv("AWS_SECRET"),
				"",
			)},
		BucketDownload: "aprenda-golang-drive-raw",
		BucketUpload:   "aprenda-golang-drive-gzip",
	}

	//create new bucket session
	b, err := bucket.New(bucket.AwsProvider, bcfg)
	if err != nil {
		log.Fatal(err)
	}

	return db, b, qc
}
