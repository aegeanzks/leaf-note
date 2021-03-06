package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
)

var (
	skeleton = base.NewSkeleton()     //创建骨架
	ChanRPC  = skeleton.ChanRPCServer //引用骨架中的RPC服务器并导出
)

//模型类型定义
type Module struct {
	*module.Skeleton //匿名组合骨架引用
}

//初始化
func (m *Module) OnInit() {
	m.Skeleton = skeleton //将创建的骨架保存到引用中
}

//销毁
func (m *Module) OnDestroy() {

}
