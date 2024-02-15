package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	require.Empty(t, err)
	defer db.Close()

	clientID := 1

	client, err := selectClient(db, clientID)
	require.Empty(t, err)
	assert.Equal(t, client.ID, 1)
	assert.NotEmpty(t, client.Birthday)
	assert.NotEmpty(t, client.Email)
	assert.NotEmpty(t, client.FIO)
	assert.NotEmpty(t, client.Login)

}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	require.Empty(t, err)
	defer db.Close()

	clientID := -1

	client, err := selectClient(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)
	assert.Empty(t, client.ID)
	assert.Empty(t, client.Birthday)
	assert.Empty(t, client.Email)
	assert.Empty(t, client.FIO)
	assert.Empty(t, client.Login)

}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	cl.ID, err = insertClient(db, cl)
	require.Empty(t, err)
	require.NotEmpty(t, cl.ID)
	clFromDB, err := selectClient(db, cl.ID)
	require.Empty(t, err)
	assert.Equal(t, cl, clFromDB)

}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	cl.ID, err = insertClient(db, cl)
	require.Empty(t, err)
	require.NotEmpty(t, cl.ID)
	clFromDB, err := selectClient(db, cl.ID)
	require.Empty(t, err)
	err = deleteClient(db, clFromDB.ID)
	require.Empty(t, err)
	_, err = selectClient(db, clFromDB.ID)
	require.Equal(t, sql.ErrNoRows, err)

}
