package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestStorage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Storage Suite")
}

var _ = Describe("Storage tests", func() {

	const (
		insertQueryShort = "INSERT INTO \"table\""
	)

	Context("when db get", func() {
		It("returns db", func() {
			db, _, err := sqlmock.New()
			Expect(err).ShouldNot(HaveOccurred())
			defer db.Close()

			storage := postgresStorage{db: db}

			Expect(storage.Get()).To(Equal(db))
		})
	})

	Describe("Transactions", func() {

		Context("when transaction is successful", func() {
			It("changes are committed", func() {
				db, sqlMock, err := sqlmock.New()
				Expect(err).ShouldNot(HaveOccurred())

				defer db.Close()

				storage := postgresStorage{db: db}

				sqlMock.ExpectBegin()

				var expectedLastInsertID int64 = 1
				var expectedLastRowsAffected int64 = 1
				expectedResult := sqlmock.NewResult(expectedLastInsertID, expectedLastRowsAffected)
				sqlMock.ExpectExec(insertQueryShort).WithArgs("arg1", "arg2", "arg3").WillReturnResult(expectedResult)

				sqlMock.ExpectCommit()

				txErr := storage.Transaction(context.TODO(), func(context context.Context, tx *sql.Tx) error {
					actualResult, err := tx.Exec(insertQueryShort, "arg1", "arg2", "arg3")
					actualLastInsertID, _ := actualResult.LastInsertId()
					actualRowsAffected, _ := actualResult.RowsAffected()

					Expect(actualLastInsertID != expectedLastInsertID || actualRowsAffected != expectedLastRowsAffected).To(BeFalse())
					return err
				})

				Expect(txErr).ShouldNot(HaveOccurred())
				err = sqlMock.ExpectationsWereMet()
				Expect(err).ShouldNot(HaveOccurred())
			})

		})

		Context("when transaction begin fails", func() {
			It("transaction was not started", func() {
				db, sqlMock, err := sqlmock.New()
				Expect(err).ShouldNot(HaveOccurred())
				defer db.Close()

				sqlMock.ExpectBegin()

				storage := postgresStorage{db: db}
				txErr := storage.Transaction(context.TODO(), func(context context.Context, tx *sql.Tx) error {
					return nil
				})

				Expect(txErr).Should(HaveOccurred())
				err = sqlMock.ExpectationsWereMet()
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Describe("Rollback fails", func() {
			Context("when transaction successfully rollbacks", func() {
				It("should return operation error", func() {
					db, sqlMock, err := sqlmock.New()
					Expect(err).ShouldNot(HaveOccurred())

					defer db.Close()

					sqlMock.ExpectBegin()
					sqlMock.ExpectRollback()

					storage := postgresStorage{db: db}
					expectedErr := fmt.Errorf("unexpected error")
					txErr := storage.Transaction(context.TODO(), func(context context.Context, tx *sql.Tx) error {
						return expectedErr
					})

					Expect(txErr).To(Equal(expectedErr))

					err = sqlMock.ExpectationsWereMet()
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			Context("when transaction rollback fails", func() {
				It("should return rollback error", func() {
					db, sqlMock, err := sqlmock.New()
					Expect(err).ShouldNot(HaveOccurred())

					defer db.Close()

					sqlMock.ExpectBegin()

					storage := postgresStorage{db: db}
					expectedErr := fmt.Errorf("unexpected error")
					txErr := storage.Transaction(context.TODO(), func(context context.Context, tx *sql.Tx) error {
						return expectedErr
					})

					Expect(txErr).Should(HaveOccurred())
					Expect(txErr).NotTo(Equal(expectedErr))

					err = sqlMock.ExpectationsWereMet()
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Context("when transaction commit fails", func() {
			It("should return commit error", func() {
				db, sqlMock, err := sqlmock.New()
				Expect(err).ShouldNot(HaveOccurred())

				defer db.Close()

				sqlMock.ExpectBegin()

				storage := postgresStorage{db: db}
				txErr := storage.Transaction(context.TODO(), func(context context.Context, tx *sql.Tx) error {
					return nil
				})

				Expect(txErr).Should(HaveOccurred())

				err = sqlMock.ExpectationsWereMet()
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
	})

	Context("when db close", func() {
		It("should close db", func() {
			db, _, err := sqlmock.New()
			Expect(err).ShouldNot(HaveOccurred())

			storage := postgresStorage{db: db}
			storage.Close()

			err = db.Ping()
			Expect(err).Should(HaveOccurred())
		})
	})
})
