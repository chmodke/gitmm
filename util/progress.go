package util

import (
	"fmt"
	"math"
	"strings"
)

type Progress struct {
	name    string //任务名称
	percent int    //百分比
	current int    //当前进度位置
	total   int    //总进度
	rate    string //进度条
	graph   string //显示符号
}

func (p *Progress) NewOption(name string, start, total int) {
	p.name = fmt.Sprintf("%-20s", name)
	p.total = total
	if p.graph == "" {
		p.graph = "*"
	}
	p.Play(start)
}

func (p *Progress) NewOptionWithGraph(name string, start, total int, graph string) {
	p.graph = graph
	p.NewOption(name, start, total)
}

func (p *Progress) Next() {
	p.Play(p.current + 1)
}

func (p *Progress) Skip(cnt int) {
	p.Play(p.current + cnt)
}

func (p *Progress) Play(cur int) {
	if cur > p.total {
		return
	}
	p.current = cur
	p.percent = p.getPercent()
	p.rate = strings.Repeat(p.graph, p.percent/2)
	fmt.Printf("\r%s[%-50s]%3d%%  %8d/%d", p.name, p.rate, p.percent, p.current, p.total)
}

func (p *Progress) Finish(status string) {
	fmt.Printf("\r%s[%-50s]%3d%%  %8d/%d  %8s", p.name, p.rate, p.percent, p.current, p.total, status)
	fmt.Printf("\n")
}

func (p *Progress) getPercent() int {
	return int(math.Min((float64(p.current)/float64(p.total))*100, 100))
}
