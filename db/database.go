package db

import (
	"errors"
	"os"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/wisepythagoras/dimoschain/utils"
)

// DB defines our database object handler.
type DB struct {
	Name string
	db   *badger.DB
}

// Open opens the database.
func (d *DB) Open() (bool, error) {
	// Get the chain's directory.
	path, err := utils.GetChainDir(true)

	if err != nil {
		return false, err
	}

	basePath := path + "/" + utils.ChainDir

	// Create the base patch for the chain directory if it doesn't exist.
	if !utils.CheckIfFileExists(basePath) {
		os.Mkdir(basePath, 0777)
	}

	// Here we define the badger options. This is also where we'll handle the horizontal
	// partitioning scheme.
	options := badger.DefaultOptions(basePath + "/" + d.Name)

	// Now try to open the database.
	db, err := badger.Open(options)

	if err != nil {
		return false, err
	}

	// Save our instance here.
	d.db = db

	return true, nil
}

// Insert creates a new entry in the database.
func (d *DB) Insert(key []byte, value []byte) (bool, error) {
	if d.db == nil {
		return false, errors.New("uninitialized database")
	}

	// Create a new transaction.
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	var err error

	// Set the data in the database.
	if err = txn.Set(key, value); err != nil {
		return false, err
	}

	// Commit the changes to the database.
	if err = txn.Commit(); err != nil {
		return false, err
	}

	return true, nil
}

// Get retrieves the contents of a specific key in the database.
func (d *DB) Get(key []byte) ([]byte, error) {
	if d.db == nil {
		return nil, errors.New("uninitialized database")
	}

	// Create a new transaction.
	txn := d.db.NewTransaction(true)
	defer txn.Discard()

	// Get the item of the entry with the hash as the key.
	item, err := txn.Get(key)

	if err != nil {
		return nil, err
	}

	// Get the vaue from the item.
	value, err := item.ValueCopy(nil)

	if err != nil {
		return nil, err
	}

	return value, nil
}
