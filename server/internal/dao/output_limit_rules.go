// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"server/internal/dao/internal"
)

// internalOutputLimitRulesDao is internal type for wrapping internal DAO implements.
type internalOutputLimitRulesDao = *internal.OutputLimitRulesDao

// outputLimitRulesDao is the data access object for table output_limit_rules.
// You can define custom methods on it to extend its functionality as you wish.
type outputLimitRulesDao struct {
	internalOutputLimitRulesDao
}

var (
	// OutputLimitRules is globally public accessible object for table output_limit_rules operations.
	OutputLimitRules = outputLimitRulesDao{
		internal.NewOutputLimitRulesDao(),
	}
)

// Fill with you ideas below.
