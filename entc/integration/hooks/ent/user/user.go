// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/facebook/ent"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldWorth holds the string denoting the worth field in the database.
	FieldWorth = "worth"

	// EdgeCards holds the string denoting the cards edge name in mutations.
	EdgeCards = "cards"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeBestFriend holds the string denoting the best_friend edge name in mutations.
	EdgeBestFriend = "best_friend"

	// Table holds the table name of the user in the database.
	Table = "users"
	// CardsTable is the table the holds the cards relation/edge.
	CardsTable = "cards"
	// CardsInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardsInverseTable = "cards"
	// CardsColumn is the table column denoting the cards relation/edge.
	CardsColumn = "user_cards"
	// FriendsTable is the table the holds the friends relation/edge. The primary key declared below.
	FriendsTable = "user_friends"
	// BestFriendTable is the table the holds the best_friend relation/edge.
	BestFriendTable = "users"
	// BestFriendColumn is the table column denoting the best_friend relation/edge.
	BestFriendColumn = "user_best_friend"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldName,
	FieldWorth,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"user_best_friend",
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/facebook/ent/entc/integration/hooks/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int
)
