package nowpaymentsrepository

import (
	"context"
	"fmt"
	"log"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	nowpaymentsmodel "github.com/jSierraB3991/now_payment_api/domain/now_payments_model"
	"gorm.io/gorm"
)

func (repo *Repository) RunMigrations(schemas []string) error {

	for _, schema := range schemas {
		// Asegúrate de que el schema existe antes de migrar
		if err := ensureSchemaExists(repo.db, schema); err != nil {
			return fmt.Errorf("schema '%s' creation failed: %w", schema, err)
		}

		dbTenant, err := repo.db.Session(&gorm.Session{NewDB: true}).DB()
		if err != nil {
			log.Fatalf("could not create session for %s: %v", schema, err)
		}

		// establecer el search_path de forma segura
		_, err = dbTenant.Exec(`SET search_path TO ` + jsierralibs.QuoteIdentifier(schema)) // o con una query preparada
		if err != nil {
			log.Fatalf("could not set search_path for %s: %v", schema, err)

		}

		ctx := context.WithValue(context.Background(), jsierralibs.ContextTenantKey, schema)
		err = repo.runAutoMigrate(ctx)
		if err != nil {
			return err
		}

		log.Printf("✅ Migrated schema: %s", schema)
	}
	return nil
}

func ensureSchemaExists(db *gorm.DB, schema string) error {
	return db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", jsierralibs.QuoteIdentifier(schema))).Error
}

func (repo *Repository) runAutoMigrate(ctx context.Context) error {
	db, err := repo.GetDb(ctx)
	if err != nil {
		return err
	}
	return db.AutoMigrate(
		&nowpaymentsmodel.NowPaymentCreatePayment{},
		&nowpaymentsmodel.NowPaymentCreateInvoice{},
	)
}
