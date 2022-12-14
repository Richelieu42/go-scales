package strKit

import "strings"

// StartWith
/*
PS: 区分大小写.

@param s		prefix不为""的情况下，如果s为""，返回值必定为false
@param prefix 	如果prefix为""，返回值必定为true

e.g.	""的情况
("", "") 	=> true
("1", "")	=> true
("", "1")	=> false

e.g.1	区分大小写
("abc", "abc") => true
("abc", "Abc") => false
*/
func StartWith(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// RemovePrefixIfExist 去掉指定的"前缀"（如果存在的话）
/*
PS:
(1) 区分大小写；
(2) 存在多个的话，只会移除第1个.

@param s		如果为""，返回""
@param prefix	如果为""，返回传参s

e.g.	""的情况
("", "")	=> ""
("1", "")	=> "1"
("", "1") 	=> ""

e.g.1	区分大小写
("abcd", "abcd") => ""
("abcd", "Abcd") => "abcd"
*/
func RemovePrefixIfExist(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// PrependIfMissing 如果给定字符串不是以给定的字符串为开头，则在"首部"添加 起始字符串.
/*
PS: 区分大小写.

e.g.
("abc", "a")	=> "abc"
("abc", "A")) 	=> "Aabc"
("abc", "0")	=> "0abc"
*/
func PrependIfMissing(str, prefix string) string {
	if !StartWith(str, prefix) {
		str = prefix + str
	}
	return str
}
