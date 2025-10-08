package pkg

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorWithArgWrap(t *testing.T) {
	err1 := WrapError(sql.ErrNoRows, "hello", ErrorArgument("world", "!"))

	require.ErrorIs(t, err1, sql.ErrNoRows, "not equal origin error")
	require.NotErrorIs(t, err1, sql.ErrTxDone, "equal not origin error")
	assert.Equal(t, "hello (world=!): "+sql.ErrNoRows.Error(), err1.Error(), "not correct message")

	err2 := WrapError(err1, "double trouble", ErrorArgument("first", 1), ErrorArgument("second", [2]int{}))

	require.ErrorIs(t, err2, err1, "not equal origin error")
	require.ErrorIs(t, err2, sql.ErrNoRows, "not equal deep origin error")
	require.NotErrorIs(t, err2, sql.ErrTxDone, "equal not origin error")
	assert.Equal(t, "double trouble (first=1, second=[0 0]): "+err1.Error(), err2.Error(), "not correct message")

	err3 := WrapError(sql.ErrNoRows, "")
	assert.Equal(t, sql.ErrNoRows.Error(), err3.Error(), "not correct message")
}

func TestErrorWithArgNew(t *testing.T) {
	err1 := ErrorWithArgs("hello", ErrorArgument("world", "!"))

	require.NotErrorIs(t, err1, sql.ErrNoRows, "equal not origin error")
	require.NotErrorIs(t, err1, sql.ErrTxDone, "equal not origin error")
	assert.Equal(t, "hello (world=!)", err1.Error(), "not correct message")

	err2 := WrapError(err1, "double trouble", ErrorArgument("first", 1), ErrorArgument("second", [2]int{}))

	require.ErrorIs(t, err2, err1, "not equal origin error")
	require.NotErrorIs(t, err2, sql.ErrNoRows, "equal not origin error")
	require.NotErrorIs(t, err2, sql.ErrTxDone, "equal not origin error")
	assert.Equal(t, "double trouble (first=1, second=[0 0]): "+err1.Error(), err2.Error(), "not correct message")

	err3 := ErrorWithArgs("", ErrorArgument("hello", "world!"))
	assert.Equal(t, "(hello=world!)", err3.Error(), "not correct message")
}

func TestErrorWithArgEqualNil(t *testing.T) {
	err1 := WrapError(nil, "hello", ErrorArgument("world", "!"))

	require.NoError(t, err1, "empty is not nil")

	require.NotErrorIs(t, err1, sql.ErrNoRows, "equal not origin error")
	require.NotErrorIs(t, err1, sql.ErrTxDone, "equal not origin error")

	err2 := WrapError(err1, "double trouble", ErrorArgument("first", 1), ErrorArgument("second", [2]int{}))

	require.NoError(t, err2, "empty is not nil")
	require.NotErrorIs(t, err2, sql.ErrNoRows, "equal not origin error")
	require.NotErrorIs(t, err2, sql.ErrTxDone, "equal not origin error")

	err3 := ErrorWithArgs("")
	assert.NoError(t, err3, "empty is not nil")
}
