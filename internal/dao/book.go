// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hellogf/internal/dao/internal"
)

// internalBookDao is internal type for wrapping internal DAO implements.
type internalBookDao = *internal.BookDao

// bookDao is the data access object for table book.
// You can define custom methods on it to extend its functionality as you wish.
type bookDao struct {
	internalBookDao
}

var (
	// Book is globally public accessible object for table book operations.
	Book = bookDao{
		internal.NewBookDao(),
	}
)

// Fill with you ideas below.