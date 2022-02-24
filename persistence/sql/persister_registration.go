package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/ory/kratos/corp"

	"github.com/ory/x/sqlcon"

	"github.com/ory/kratos/selfservice/flow/registration"
)

func (p *Persister) CreateRegistrationFlow(ctx context.Context, r *registration.Flow) error {
	r.NID = corp.ContextualizeNID(ctx, p.nid)
	r.EnsureInternalContext()
	return p.GetConnection(ctx).Create(r)
}

func (p *Persister) UpdateRegistrationFlow(ctx context.Context, r *registration.Flow) error {
	r.EnsureInternalContext()
	cp := *r
	cp.NID = corp.ContextualizeNID(ctx, p.nid)
	return p.update(ctx, cp)
}

func (p *Persister) GetRegistrationFlow(ctx context.Context, id uuid.UUID) (*registration.Flow, error) {
	var r registration.Flow
	if err := p.GetConnection(ctx).Where("id = ? AND nid = ?",
		id, corp.ContextualizeNID(ctx, p.nid)).First(&r); err != nil {
		return nil, sqlcon.HandleError(err)
	}

	return &r, nil
}

func (p *Persister) DeleteExpiredRegistrationFlows(ctx context.Context, expiresAt time.Time, limit, batch int) error {
	for ok := true; ok; ok = batch <= limit {
		limit -= batch
		// #nosec G201
		count, err := p.GetConnection(ctx).RawQuery(fmt.Sprintf(
			"DELETE FROM `%s` WHERE `expires_at` <= ? LIMIT ?",
			new(registration.Flow).TableName(ctx),
		),
			expiresAt,
			batch,
		).ExecWithCount()
		if err != nil {
			return sqlcon.HandleError(err)
		}

		if count == 0 {
			break
		}
	}
	return nil
}
