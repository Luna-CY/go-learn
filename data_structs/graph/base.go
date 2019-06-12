package graph

const IntMax = int(^uint(0) >> 1)

type Graph interface {
	Create(v []Vertex, vr []Arc)
	Destroy()
	GetValue(v int) string        // 获取顶点v的值
	PutValue(v int, value string) // 设置顶点v的值
	LocateVex(u int)              // 返回u的位置
	FirstAdjVex(v int) Vertex     // 返回顶点v的一个邻接顶点，若顶点v在图G中没有邻接顶点返回空
	NextAdjVex(v, w int) Vertex
	InsertVex(v Vertex)
	DeleteVex(v Vertex)
	InsertArc(arc Arc)
	DeleteArc(arc Arc)
	DFSTraverse()
	HFSTraverse()
}

type Arc struct {
	Vi, Vj, Weight int
}

type Vertex struct {
	Name, Value string
}
