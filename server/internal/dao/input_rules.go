// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalInputRulesDao is internal type for wrapping internal DAO implements.
type internalInputRulesDao = *internal.InputRulesDao

// inputRulesDao is the data access object for table input_rules.
// You can define custom methods on it to extend its functionality as you wish.
type inputRulesDao struct {
	internalInputRulesDao
}

var (
	// InputRules is globally public accessible object for table input_rules operations.
	InputRules = inputRulesDao{
		internal.NewInputRulesDao(),
	}
)

// Fill with you ideas below.
