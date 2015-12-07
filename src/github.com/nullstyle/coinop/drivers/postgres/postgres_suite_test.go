package postgres_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"strings"
	"testing"
)

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postgres Suite")
}

var db *sqlx.DB

const initDB = `
	DROP TABLE IF EXISTS webhooks;
	CREATE TABLE webhooks (
		id 								 	bigserial,
		url 							 	varchar(2048) NOT NULL,
		destination_filter 	varchar(255) 	NOT NULL,
		memo_type_filter		varchar(255),
		memo_filter					varchar(255),
		created_at 					timestamp without time zone
	);
`

const clearDB = `DELETE FROM webhooks`

var _ = BeforeSuite(func() {
	url := os.Getenv("POSTGRES_URL")
	if url == "" {
		url = "postgres://localhost/coinop_test?sslmode=disable"
	}
	var err error
	db, err = sqlx.Connect("postgres", url)
	Expect(err).NotTo(HaveOccurred())

	execAll(initDB)
})

var _ = AfterSuite(func() {
	db.Close()
})

var _ = BeforeEach(func() {
	execAll(clearDB)
})

func execAll(query string) {
	qs := strings.Split(query, ";")
	for _, q := range qs {
		q = strings.TrimSpace(q)
		if q == "" {
			continue
		}
		db.MustExec(q)
	}
}
