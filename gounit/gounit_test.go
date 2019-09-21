package gounit

import (
	"testing"
)

// 单元测试
// 命令： go test -v .
func TestGetSum(t *testing.T) {
	sum := GetSum(2,3)
	if sum == 5{
		t.Log("测试通过")
	}else{
		t.Errorf("有点问题哦！2+3应该是%d，可是等于%d\n", 5, sum)
		//  t.Fatal() = log +FailNow
		return
	}
}

// 性能压力测试，不看结果，只看内存开销
// 命令： go test -bench gounit_test.go (-count=3表示连续做多少轮球平均值)
func BenchmarkSetNum(b *testing.B) {
	b.Log("开始对SetNum进行压力测试")
	b.ReportAllocs()
	for i:= 0; i < b.N; i++{
		SetNum(3)
	}
}

// 补充命令
/*
	1.生成性能分析文件： go test -bench=.-benchtime="3s" -cpuprofile=profile.out
	2.查看性能分析文件： go tool pprof xxx_test profile.out   =>  进入后可以通过top n 查看耗时情况
	3.输出在web查看文件： go tool pprof --web xxx_test profile.out   =>  进入后可以通过top n 查看耗时情况,需要安装Graphviz
	4.输出为pdf：	go tool pprof --pdf profile.out > profile.pdf		// 需要安装Graphviz
 */

/*
	当我们遇到一个断言错误的时候，我们就会判断这个测试用例失败，就会使用到：
	Fail  : case失败，测试用例继续
	FailedNow : case失败，测试用例中断

	当我们遇到一个断言错误，只希望跳过这个错误，但是不希望标示测试用例失败，会使用到：
	SkipNow : case跳过，测试用例不继续

	当我们只希望在一个地方打印出信息，我们会用到:
	Log : 输出信息
	Logf : 输出有format的信息

	当我们希望跳过这个用例，并且打印出信息:
	Skip : Log + SkipNow
	Skipf : Logf + SkipNow

	当我们希望断言失败的时候，测试用例失败，打印出必要的信息，但是测试用例继续：
	Error : Log + Fail
	Errorf : Logf + Fail

	当我们希望断言失败的时候，测试用例失败，打印出必要的信息，测试用例中断：
	Fatal : Log + FailNow
	Fatalf : Logf + FailNow
 */
