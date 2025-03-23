// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalLogLoginsDao is internal type for wrapping internal DAO implements.
type internalLogLoginsDao = *internal.LogLoginsDao

// logLoginsDao is the data access object for table log_logins.
// You can define custom methods on it to extend its functionality as you wish.
type logLoginsDao struct {
	internalLogLoginsDao
}

var (
	// LogLogins is globally public accessible object for table log_logins operations.
	LogLogins = logLoginsDao{
		internal.NewLogLoginsDao(),
	}
)

// Fill with you ideas below.
