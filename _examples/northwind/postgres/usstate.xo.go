package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// UsState represents a row from 'public.us_states'.
type UsState struct {
	StateID     int            `json:"state_id"`     // state_id
	StateName   sql.NullString `json:"state_name"`   // state_name
	StateAbbr   sql.NullString `json:"state_abbr"`   // state_abbr
	StateRegion sql.NullString `json:"state_region"` // state_region
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the UsState exists in the database.
func (us *UsState) Exists() bool {
	return us._exists
}

// Deleted returns true when the UsState has been marked for deletion from
// the database.
func (us *UsState) Deleted() bool {
	return us._deleted
}

// Insert inserts the UsState to the database.
func (us *UsState) Insert(ctx context.Context, db DB) error {
	switch {
	case us._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case us._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (basic)
	const sqlstr = `INSERT INTO public.us_states (` +
		`state_id, state_name, state_abbr, state_region` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`
	// run
	logf(sqlstr, us.StateID, us.StateName, us.StateAbbr, us.StateRegion)
	if err := db.QueryRowContext(ctx, sqlstr, us.StateID, us.StateName, us.StateAbbr, us.StateRegion).Scan(&us.StateID); err != nil {
		return logerror(err)
	}
	// set exists
	us._exists = true
	return nil
}

// Update updates a UsState in the database.
func (us *UsState) Update(ctx context.Context, db DB) error {
	switch {
	case !us._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case us._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.us_states SET (` +
		`state_name, state_abbr, state_region` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE state_id = $4`
	// run
	logf(sqlstr, us.StateName, us.StateAbbr, us.StateRegion, us.StateID)
	if _, err := db.ExecContext(ctx, sqlstr, us.StateName, us.StateAbbr, us.StateRegion, us.StateID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the UsState to the database.
func (us *UsState) Save(ctx context.Context, db DB) error {
	if us.Exists() {
		return us.Update(ctx, db)
	}
	return us.Insert(ctx, db)
}

// Upsert performs an upsert for UsState.
//
// NOTE: PostgreSQL 9.5+ only
func (us *UsState) Upsert(ctx context.Context, db DB) error {
	switch {
	case us._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.us_states (` +
		`state_id, state_name, state_abbr, state_region` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (state_id) DO UPDATE SET (` +
		`state_id, state_name, state_abbr, state_region` +
		`) = (` +
		`EXCLUDED.state_id, EXCLUDED.state_name, EXCLUDED.state_abbr, EXCLUDED.state_region` +
		`)`
	// run
	logf(sqlstr, us.StateID, us.StateName, us.StateAbbr, us.StateRegion)
	if _, err := db.ExecContext(ctx, sqlstr, us.StateID, us.StateName, us.StateAbbr, us.StateRegion); err != nil {
		return err
	}
	// set exists
	us._exists = true
	return nil
}

// Delete deletes the UsState from the database.
func (us *UsState) Delete(ctx context.Context, db DB) error {
	switch {
	case !us._exists: // doesn't exist
		return nil
	case us._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.us_states WHERE state_id = $1`
	// run
	logf(sqlstr, us.StateID)
	if _, err := db.ExecContext(ctx, sqlstr, us.StateID); err != nil {
		return logerror(err)
	}
	// set deleted
	us._deleted = true
	return nil
}

// UsStateByStateID retrieves a row from 'public.us_states' as a UsState.
//
// Generated from index 'us_states_pkey'.
func UsStateByStateID(ctx context.Context, db DB, stateID int) (*UsState, error) {
	// query
	const sqlstr = `SELECT ` +
		`state_id, state_name, state_abbr, state_region ` +
		`FROM public.us_states ` +
		`WHERE state_id = $1`
	// run
	logf(sqlstr, stateID)
	us := UsState{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, stateID).Scan(&us.StateID, &us.StateName, &us.StateAbbr, &us.StateRegion); err != nil {
		return nil, logerror(err)
	}
	return &us, nil
}