package main

import "testing"

// 本文件要以_test为结尾 才能被 go test -bench=.检测到
//go语言测试代码只需要放到以 _test.go 结尾的文件中即可。
//golang的测试分为单元测试和性能测试，单元测试的测试用例必须以Test开头，其后的函数名不能以小写字母开头；
//性能测试必须以Benchmark开头，其后的函数名不能以小写字母开头。
//为了测试方法和被测试方法的可读性，一般Test或Benchmark后为被测试方法的函数名。
//测试代码通常与测试对象文件在同一目录下。

// 判断逻辑正确性
func TestRepeatLimitedString(t *testing.T) {
	tests := []struct {
		s   string
		n   int
		ans string
	}{
		{"cczazcc", 3, "zzcccac"},
		{"aababab", 2, "bbabaa"},
	}
	for _, tt := range tests {
		if actual := repeatLimitedString(tt.s, tt.n); actual != tt.ans {
			t.Errorf("Add(%s,%d) got %s;expected %s", tt.s, tt.n, actual, tt.ans)
		}
	}
}

// benchmark 性能测试
//-bench regexp 性能测试，运行指定的测试函
//-bench . 运行所有的benchmark函数测试，指定名称则只执行具体测试方法而不是全部
//-benchmem 性能测试的时候显示测试函数的内存分配的统计信息
//－count n 运行测试和性能多少此，默认一次
//-run regexp 只运行特定的测试函数
//-timeout t 测试时间如果超过 t 则panic，默认10分钟
//-v 显示测试的详细信息

// 压测
func BenchmarkRepeatLimitedString(t *testing.B) {
	//重置时间点
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		repeatLimitedString("dsfadfcxxczsdacsdc", 3)
	}
}

/*
=== RUN   TestRepeatLimitedString
--- PASS: TestRepeatLimitedString (0.00s)
goos: darwin
goarch: amd64
pkg: acm
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkRepeatLimitedString
BenchmarkRepeatLimitedString-4(表示GOMAXPROCS为4)           5517868（一共执行了5517868次，即b.N的值）           （平均每次操作花费了 219。2 纳秒）219.2 ns/op
PASS
ok      acm     2.446s
*/
func splitWordsBySeparator(words []string, separator byte) []string {
	ans := []string{}
	for _, str := range words {
		lenS := len(words)
		tempStr := ""
		for i := 0; i < lenS; i++ {
			if str[i] == separator {
				if tempStr != "" {
					ans = append(ans, tempStr)
					tempStr = ""
				}
			} else {
				tempStr += string(str[i])
			}
		}
		if tempStr != "" {
			ans = append(ans, tempStr)
		}
	}
	return ans
}
