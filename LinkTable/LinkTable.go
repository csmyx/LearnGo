/*
 * @Author: SA21225357 MaYunxiu
 * @Date: 2022-04-06 17:43:59
 * @Last Modified by: SA21225357  MaYunxiu
 * @Last Modified time: 2022-04-06 19:31:00
 */
package LinkTable

// 链表节点的值类型
type ValType = int

// 链表节点
type Node struct {
	val  ValType // 节点值
	prev *Node   // 前节点
	next *Node   // 后节点
}

func NewNode(v ValType) *Node {
	return &Node{v, nil, nil}
}

func Fuck() {

}

// 链表接口设计
type LinkTable interface {
	Empty() bool
	Size() int
	Head() (*Node, error) // 获取头节点
	Tail() (*Node, error) // 获取尾节点
	ToString() string

	/* 插入节点 */
	Insert(*Node)                               // 在尾节点后插入给定节点
	InsertVal(ValType)                          // 插入一个给定值的节点
	InsertBefor(node1 *Node, node2 *Node) error // 在 node1 后插入 node2
	InsertAfter(node1 *Node, node2 *Node) error // 在 node1 前插入 node2

	/* 删除节点 */
	Delete(*Node) error
	DeleteFirst(ValType) error // 删除第一个等于给定值的节点

	/* 查找节点 */
	FindFirst(ValType) (*Node, error) // 返回第一个等于给定值的节点

	Clear() // 清空链表
}
