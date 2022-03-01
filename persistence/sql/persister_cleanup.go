package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/ory/kratos/continuity"
	"github.com/ory/kratos/corp"
	"github.com/ory/kratos/selfservice/flow/login"
	"github.com/ory/kratos/selfservice/flow/recovery"
	"github.com/ory/kratos/selfservice/flow/registration"
	"github.com/ory/kratos/selfservice/flow/settings"
	"github.com/ory/kratos/selfservice/flow/verification"
	"github.com/ory/x/sqlcon"
)

type CleanupTableNameType int

const (
	CleanupCorpContextualizeTable CleanupTableNameType = iota
	CleanupContinuityContainerTable
	CleanupSettingsFlowTable
	CleanupLoginFlowTable
	CleanupRecoveryFlowTable
	CleanupRegistrationFlowTable
	CleanupVerificationFlowTable
)

func (typ CleanupTableNameType) TableName(ctx context.Context) string {
	var s string
	switch typ {
	case CleanupCorpContextualizeTable:
		s = corp.ContextualizeTableName(ctx, "sessions")
	case CleanupContinuityContainerTable:
		s = new(continuity.Container).TableName(ctx)
	case CleanupSettingsFlowTable:
		s = new(settings.Flow).TableName(ctx)
	case CleanupLoginFlowTable:
		s = new(login.Flow).TableName(ctx)
	case CleanupRecoveryFlowTable:
		s = new(recovery.Flow).TableName(ctx)
	case CleanupRegistrationFlowTable:
		s = new(registration.Flow).TableName(ctx)
	case CleanupVerificationFlowTable:
		s = new(verification.Flow).TableName(ctx)
	}
	return s
}

type CleanupParams struct {
	Tables    []CleanupTableNameType
	Batch     int
	Limit     int
	ExpiresAt time.Time
}

func (p *Persister) Cleanup(ctx context.Context, opt CleanupParams) error {
	const stmt = `
DELETE FROM "%s"
WHERE ctid IN (
    SELECT ctid
    FROM "%s"
	WHERE expires_at < ?
    ORDER BY expires_at
    LIMIT ?
)`
	if opt.Limit < opt.Batch {
		opt.Batch = opt.Limit
	}
	for _, typ := range opt.Tables {

		tableName := typ.TableName(ctx)
		p.r.Logger().Printf("Starting clean expired records for %q table", tableName)

		for ok := true; ok; ok = opt.Batch <= opt.Limit {
			opt.Limit -= opt.Batch
			// #nosec G201
			count, err := p.GetConnection(ctx).RawQuery(
				fmt.Sprintf(stmt, tableName, tableName),
				opt.ExpiresAt,
				opt.Batch,
			).ExecWithCount()
			if err != nil {
				return sqlcon.HandleError(err)
			}

			if count == 0 || opt.Limit <= 0 {
				break
			}
		}
		p.r.Logger().Printf("Done. Expired records in %q table cleaned", tableName)
	}
	return nil
}
