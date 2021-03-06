package controllers

import "github.com/robfig/revel"

func init() {
	rev.RegisterPlugin(GorpPlugin{})
	rev.InterceptMethod((*GorpController).Begin, rev.BEFORE)
	rev.InterceptMethod((*GorpController).Commit, rev.AFTER)
	rev.InterceptMethod((*GorpController).Rollback, rev.FINALLY)
}
