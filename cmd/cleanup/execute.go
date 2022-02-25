package cleanup

import (
	"context"
	"os"

	"github.com/ory/kratos/driver"
)

func Execute(ctx context.Context, r driver.Registry) {
	r.Logger().Println("Cleanup started")
	err := r.Persister().CleanupDatabase(ctx)
	if err != nil {
		r.Logger().WithError(err).Fatalf("Cleanup failed")
	}
	r.Logger().Println("Cleanup successful done")

	os.Exit(0)
}
