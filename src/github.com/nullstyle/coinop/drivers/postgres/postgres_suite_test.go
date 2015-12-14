package postgres_test

import (
	"os"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	. "github.com/nullstyle/coinop/drivers/postgres"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Postgres Suite")
}

var db *sqlx.DB
var badDB *sqlx.DB

const clearDB = `
	DELETE FROM coinop.webhooks;
	DELETE FROM coinop.deliveries;
	DELETE FROM coinop.kv;
`

var _ = BeforeSuite(func() {
	connectDB()
	var err error
	badDB, err = sqlx.Open("postgres", "postgres://localhost/SHOULD_NOT_EXIST?sslmode=disable")
	Expect(err).NotTo(HaveOccurred())
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

func connectDB() {
	url := os.Getenv("POSTGRES_URL")
	if url == "" {
		url = "postgres://localhost/coinop_test?sslmode=disable"
	}
	var err error
	db, err = sqlx.Connect("postgres", url)
	Expect(err).NotTo(HaveOccurred())

	d := &Driver{db}
	err = d.RebuildSchema()
	Expect(err).NotTo(HaveOccurred())
}
