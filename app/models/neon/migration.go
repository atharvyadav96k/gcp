package neon

import "errors"

// RunMigrations executes AutoMigrate for all provided models.
// This should be called once when the database is ready.
// Call it manually from your application before using CRUD operations.
//
// Example usage:
//   app := common.Init()
//   defer app.Close()
//
//   err := app.GetNeonService().RunMigrations(
//       &database_models.User{},
//       &database_models.Org{},
//       &database_models.Vehicle{},
//       &database_models.RFID{},
//       &database_models.Scanner{},
//       &database_models.RootUser{},
//       &database_models.WhiteListedEmail{},
//   )
//   if err != nil {
//       log.Fatalf("Migration failed: %v", err)
//   }
func (s *Service) RunMigrations(models ...interface{}) error {
	if s == nil || s.GetDB() == nil {
		return errors.New("Neon service is not initialized. Call InitDB() before running migrations.")
	}

	// Run AutoMigrate for all provided models
	if err := s.GetDB().AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}
