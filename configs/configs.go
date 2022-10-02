package configs

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type Configs struct {
	Logger    *logrus.Logger
	DB        *sql.DB
	Validator *Validator
}

type Validator struct {
	Validate *validator.Validate
	Trans    *ut.Translator
}

func (c *Configs) Close() {
	c.DB.Close()
}

var instance *Configs
var lock *sync.Mutex = &sync.Mutex{}

func GetInstance() *Configs {
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = &Configs{
			Logger:    createLogger(),
			DB:        createDBConnection(),
			Validator: createValidator(),
		}
	}

	return instance

}

func createValidator() *Validator {
	id := id.New()
	uni := ut.New(id, id)

	trans, _ := uni.GetTranslator("id")

	validate := validator.New()
	id_translations.RegisterDefaultTranslations(validate, trans)

	return &Validator{
		Validate: validate,
		Trans:    &trans,
	}
}

func createLogger() *logrus.Logger {
	logger := logrus.New()
	log.SetOutput(logger.Writer())
	log.SetFlags(0)
	logger.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		ForceQuote:    true,
		FullTimestamp: true,
	}
	logger.Info("Setup logger complete")

	return logger
}

func createDBConnection() *sql.DB {
	// TODO: change
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		// os.Getenv("MYSQL_DB_USER"),
		// os.Getenv("MYSQL_DB_PASS"),
		// os.Getenv("MYSQL_DB_HOST"),
		// os.Getenv("MYSQL_DB_PORT"),
		// os.Getenv("MYSQL_DB_NAME"),

		"user",
		"password",
		"localhost",
		"3306",
		"cake_store",
	)

	retryCount := 0
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("MySQL not ready yet...")
			retryCount++
		} else {
			log.Println("Connected to MySQL...")
			return connection
		}

		if retryCount > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two second")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
