package rocketmq5Kit

import "github.com/richelieu42/go-scales/src/core/sliceKit"

func MixTags(tags ...string) string {
	return sliceKit.Join(tags, "||")
}
