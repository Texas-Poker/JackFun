package redis_helper

import (
	"github.com/go-redis/redis"

	"time"
)

// Nil is a wrapper for redis.Nil
const Nil = redis.Nil

type (
	// Client is a wrapper for redis.Client
	Client = redis.Client

	// Pipeliner is a wrapper for redis.Pipeliner
	Pipeliner = redis.Pipeliner

	// Pipeline is a wrapper for redis.Pipeline
	Pipeline = redis.Pipeline

	// Cmder is a wrapper for redis.Cmder
	Cmder = redis.Cmder

	// Cmd is a wrapper for redis.Cmd
	Cmd = redis.Cmd

	// SliceCmd is a wrapper for redis.SliceCmd
	SliceCmd = redis.SliceCmd

	// StatusCmd is a wrapper for redis.StatusCmd
	StatusCmd = redis.StatusCmd

	// IntCmd is a wrapper for redis.IntCmd
	IntCmd = redis.IntCmd

	// DurationCmd is a wrapper for redis.DurationCmd
	DurationCmd = redis.DurationCmd

	// TimeCmd is a wrapper for redis.TimeCmd
	TimeCmd = redis.TimeCmd

	// BoolCmd is a wrapper for redis.BoolCmd
	BoolCmd = redis.BoolCmd

	// StringCmd is a wrapper for redis.StringCmd
	StringCmd = redis.StringCmd

	// FloatCmd is a wrapper for redis.FloatCmd
	FloatCmd = redis.FloatCmd

	// StringSliceCmd is a wrapper for redis.StringSliceCmd
	StringSliceCmd = redis.StringSliceCmd

	// BoolSliceCmd is a wrapper for redis.BoolSliceCmd
	BoolSliceCmd = redis.BoolSliceCmd

	// StringStringMapCmd is a wrapper for redis.StringStringMapCmd
	StringStringMapCmd = redis.StringStringMapCmd

	// StringIntMapCmd is a wrapper for redis.StringIntMapCmd
	StringIntMapCmd = redis.StringIntMapCmd

	// StringStructMapCmd is a wrapper for redis.StringStructMapCmd
	StringStructMapCmd = redis.StringStructMapCmd

	// XMessageSliceCmd is a wrapper for redis.XMessageSliceCmd
	XMessageSliceCmd = redis.XMessageSliceCmd

	// XStreamSliceCmd is a wrapper for redis.XStreamSliceCmd
	XStreamSliceCmd = redis.XStreamSliceCmd

	// XPendingCmd is a wrapper for redis.XPendingCmd
	XPendingCmd = redis.XPendingCmd

	// XPendingExtCmd is a wrapper for redis.XPendingExtCmd
	XPendingExtCmd = redis.XPendingExtCmd

	// ZSliceCmd is a wrapper for redis.ZSliceCmd
	ZSliceCmd = redis.ZSliceCmd

	// ZWithKeyCmd is a wrapper for redis.ZWithKeyCmd
	ZWithKeyCmd = redis.ZWithKeyCmd

	// ScanCmd is a wrapper for redis.ScanCmd
	ScanCmd = redis.ScanCmd

	// ClusterSlotsCmd is a wrapper for redis.ClusterSlotsCmd
	ClusterSlotsCmd = redis.ClusterSlotsCmd

	// GeoLocationCmd is a wrapper for redis.GeoLocationCmd
	GeoLocationCmd = redis.GeoLocationCmd

	// GeoPosCmd is a wrapper for redis.GeoPosCmd
	GeoPosCmd = redis.GeoPosCmd

	// CommandsInfoCmd is a wrapper for redis.CommandsInfoCmd
	CommandsInfoCmd = redis.CommandsInfoCmd

	// Z is a wrapper for redis.Z
	Z = redis.Z
)

// NewCmd is a wrapper for redis.NewCmd
func NewCmd(args ...interface{}) *redis.Cmd {
	return redis.NewCmd(args...)
}

// NewSliceCmd is a wrapper for redis.NewSliceCmd
func NewSliceCmd(args ...interface{}) *redis.SliceCmd {
	return redis.NewSliceCmd(args...)
}

// NewStatusCmd is a wrapper for redis.NewStatusCmd
func NewStatusCmd(args ...interface{}) *redis.StatusCmd {
	return redis.NewStatusCmd(args...)
}

// NewIntCmd is a wrapper for redis.NewIntCmd
func NewIntCmd(args ...interface{}) *redis.IntCmd {
	return redis.NewIntCmd(args...)
}

// NewDurationCmd is a wrapper for redis.NewDurationCmd
func NewDurationCmd(precision time.Duration, args ...interface{}) *redis.DurationCmd {
	return redis.NewDurationCmd(precision, args...)
}

// NewTimeCmd is a wrapper for redis.NewTimeCmd
func NewTimeCmd(args ...interface{}) *redis.TimeCmd {
	return redis.NewTimeCmd(args...)
}

// NewBoolCmd is a wrapper for redis.NewBoolCmd
func NewBoolCmd(args ...interface{}) *redis.BoolCmd {
	return redis.NewBoolCmd(args...)
}

// NewStringCmd is a wrapper for redis.NewStringCmd
func NewStringCmd(args ...interface{}) *redis.StringCmd {
	return redis.NewStringCmd(args...)
}

// NewFloatCmd is a wrapper for redis.NewFloatCmd
func NewFloatCmd(args ...interface{}) *redis.FloatCmd {
	return redis.NewFloatCmd(args...)
}

// NewStringSliceCmd is a wrapper for redis.NewStringSliceCmd
func NewStringSliceCmd(args ...interface{}) *redis.StringSliceCmd {
	return redis.NewStringSliceCmd(args...)
}

// NewBoolSliceCmd is a wrapper for redis.NewBoolSliceCmd
func NewBoolSliceCmd(args ...interface{}) *redis.BoolSliceCmd {
	return redis.NewBoolSliceCmd(args...)
}

// NewStringStringMapCmd is a wrapper for redis.NewStringStringMapCmd
func NewStringStringMapCmd(args ...interface{}) *redis.StringStringMapCmd {
	return redis.NewStringStringMapCmd(args...)
}

// NewStringIntMapCmd is a wrapper for redis.NewStringIntMapCmd
func NewStringIntMapCmd(args ...interface{}) *redis.StringIntMapCmd {
	return redis.NewStringIntMapCmd(args...)
}

// NewStringStructMapCmd is a wrapper for redis.NewStringStructMapCmd
func NewStringStructMapCmd(args ...interface{}) *redis.StringStructMapCmd {
	return redis.NewStringStructMapCmd(args...)
}

// NewXMessageSliceCmd is a wrapper for redis.NewXMessageSliceCmd
func NewXMessageSliceCmd(args ...interface{}) *redis.XMessageSliceCmd {
	return redis.NewXMessageSliceCmd(args...)
}

// NewXStreamSliceCmd is a wrapper for redis.NewXStreamSliceCmd
func NewXStreamSliceCmd(args ...interface{}) *redis.XStreamSliceCmd {
	return redis.NewXStreamSliceCmd(args...)
}

// NewXPendingCmd is a wrapper for redis.NewXPendingCmd
func NewXPendingCmd(args ...interface{}) *redis.XPendingCmd {
	return redis.NewXPendingCmd(args...)
}

// NewXPendingExtCmd is a wrapper for redis.NewXPendingExtCmd
func NewXPendingExtCmd(args ...interface{}) *redis.XPendingExtCmd {
	return redis.NewXPendingExtCmd(args...)
}

// NewZSliceCmd is a wrapper for redis.NewZSliceCmd
func NewZSliceCmd(args ...interface{}) *redis.ZSliceCmd {
	return redis.NewZSliceCmd(args...)
}

// NewZWithKeyCmd is a wrapper for redis.NewZWithKeyCmd
func NewZWithKeyCmd(args ...interface{}) *redis.ZWithKeyCmd {
	return redis.NewZWithKeyCmd(args...)
}

// NewScanCmd is a wrapper for redis.NewScanCmd
func NewScanCmd(process func(cmd redis.Cmder) error, args ...interface{}) *redis.ScanCmd {
	return redis.NewScanCmd(process, args...)
}

// NewClusterSlotsCmd is a wrapper for redis.NewClusterSlotsCmd
func NewClusterSlotsCmd(args ...interface{}) *redis.ClusterSlotsCmd {
	return redis.NewClusterSlotsCmd(args...)
}

// NewGeoLocationCmd is a wrapper for redis.NewGeoLocationCmd
func NewGeoLocationCmd(q *redis.GeoRadiusQuery, args ...interface{}) *redis.GeoLocationCmd {
	return redis.NewGeoLocationCmd(q, args...)
}

// NewGeoPosCmd is a wrapper for redis.NewGeoPosCmd
func NewGeoPosCmd(args ...interface{}) *redis.GeoPosCmd {
	return redis.NewGeoPosCmd(args...)
}

// NewCommandsInfoCmd is a wrapper for redis.NewCommandsInfoCmd
func NewCommandsInfoCmd(args ...interface{}) *redis.CommandsInfoCmd {
	return redis.NewCommandsInfoCmd(args...)
}
