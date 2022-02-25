package sql

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/ory/kratos/corp"

	"github.com/gofrs/uuid"

	"github.com/ory/x/sqlcon"

	"github.com/ory/kratos/continuity"
)

var _ continuity.Persister = new(Persister)

func (p *Persister) SaveContinuitySession(ctx context.Context, c *continuity.Container) error {
	c.NID = corp.ContextualizeNID(ctx, p.nid)
	return sqlcon.HandleError(p.GetConnection(ctx).Create(c))
}

func (p *Persister) GetContinuitySession(ctx context.Context, id uuid.UUID) (*continuity.Container, error) {
	var c continuity.Container
	if err := p.GetConnection(ctx).Where("id = ? AND nid = ?", id, corp.ContextualizeNID(ctx, p.nid)).First(&c); err != nil {
		return nil, sqlcon.HandleError(err)
	}
	return &c, nil
}

func (p *Persister) DeleteContinuitySession(ctx context.Context, id uuid.UUID) error {
	if count, err := p.GetConnection(ctx).RawQuery(
		// #nosec
		fmt.Sprintf("DELETE FROM %s WHERE id=? AND nid=?",
			new(continuity.Container).TableName(ctx)), id, corp.ContextualizeNID(ctx, p.nid)).ExecWithCount(); err != nil {
		return sqlcon.HandleError(err)
	} else if count == 0 {
		return errors.WithStack(sqlcon.ErrNoRows)
	}
	return nil
}

func (p *Persister) DeleteExpiredContinuitySessions(ctx context.Context, expiresAt time.Time, limit, batch int) error {
	for ok := true; ok; ok = batch <= limit {
		limit -= batch
		// #nosec G201
		count, err := p.GetConnection(ctx).RawQuery(fmt.Sprintf(
			"DELETE FROM `%s` WHERE `expires_at` <= ? LIMIT ?",
			new(continuity.Container).TableName(ctx),
		),
			expiresAt,
			batch,
		).ExecWithCount()
		if err != nil {
			return sqlcon.HandleError(err)
		}

		if count == 0 || limit <= 0 {
			break
		}
	}
	return nil
}
