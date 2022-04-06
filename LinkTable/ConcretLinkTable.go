/*
 * @Author: SA21225357 MaYunxiu
 * @Date: 2022-04-06 17:43:56
 * @Last Modified by: SA21225357  MaYunxiu
 * @Last Modified time: 2022-04-06 19:48:55
 */
package LinkTable

import (
	"errors"
	"fmt"
)

/* 定义具体实现 */
type ConcretLinkTable struct {
	count int   // 节点数
	head  *Node // 头节点
	tail  *Node //尾节点
}

func NewLinkTable() LinkTable {
	return &ConcretLinkTable{0, nil, nil}
}

/*------------------------------------*/

func (lt *ConcretLinkTable) Empty() bool {
	return lt.head == nil
}

func (lt *ConcretLinkTable) Size() int {
	return lt.count
}

func (lt *ConcretLinkTable) Head() (*Node, error) {
	if !lt.Empty() {
		return lt.head, nil
	}
	return nil, errors.New("LinkTable: No head node exists!")
}

func (lt *ConcretLinkTable) Tail() (*Node, error) {
	if !lt.Empty() {
		return lt.tail, nil
	}
	return nil, errors.New("LinkTable.Tail: No head node exists!")
}

func (lt *ConcretLinkTable) ToString() string {
	s := "LinkTable: [ "
	if !lt.Empty() {
		for cur := lt.head; cur != lt.tail; cur = cur.next {
			s += fmt.Sprintf("%d, ", cur.val)
		}
		s += fmt.Sprintf("%d", lt.tail.val)
	}
	s += " ]"
	return s
}

/* 插入节点 */

func (lt *ConcretLinkTable) Insert(node *Node) {
	if lt.Empty() { // 原链表为空
		node.prev = nil
		lt.head = node
	} else {
		node.prev = lt.tail
		lt.tail.next = node
	}
	lt.tail = node
	lt.tail.next = nil
	lt.count++ // 节点数自增
}

func (lt *ConcretLinkTable) InsertVal(val ValType) {
	lt.Insert(&Node{val, nil, nil})
}

// 注意：node1 应该是 lt 中的节点！
func (lt *ConcretLinkTable) InsertBefor(node1 *Node, node2 *Node) error {
	if node1 != nil && node2 != nil {
		node2.prev = node1.prev
		node1.prev = node2
		node2.next = node1
		if node2.prev != nil {
			node2.prev.next = node2
		}
		lt.count++
		return nil
	}
	return errors.New("LinkTable.InsertBefor: node is nil!")
}

// 注意：node1 应该是 lt 中的节点！
func (lt *ConcretLinkTable) InsertAfter(node1 *Node, node2 *Node) error {
	if node1 != nil && node2 != nil {
		node2.next = node1.next
		node1.next = node2
		node2.prev = node1
		if node2.next != nil {
			node2.next.prev = node2
		}
		lt.count++
		return nil
	}
	return errors.New("LinkTable.InsertAfter: node is nil!")
}

/*------------------------------------*/

/* 删除节点 */

// 注意：node 应该是 lt 中的节点！
func (lt *ConcretLinkTable) Delete(node *Node) error {
	if node != nil {
		if node.prev != nil {
			node.prev.next = node.next
		} else { // node 为头节点
			lt.head = lt.head.next
		}

		if node.next != nil {
			node.next.prev = node.prev
		} else { // node 为尾节点
			lt.tail = lt.tail.prev
		}

		lt.count--
		return nil
	}
	return errors.New("LinkTable.Delete: node is nil!")
}

func (lt *ConcretLinkTable) DeleteFirst(val ValType) error {
	cur, err := lt.FindFirst(val)
	if err == nil {
		lt.Delete(cur)
		return nil
	}
	return errors.New("LinkTable.DeleteFirst: No node with the given value!")
}

/*------------------------------------*/

/* 查找节点 */

func (lt *ConcretLinkTable) FindFirst(val ValType) (*Node, error) {
	cur := lt.head
	for i := 0; i < lt.count; i++ {
		if cur.val == val {
			return cur, nil
		}
		cur = cur.next
	}
	return nil, errors.New("LinkTable.FindFirst: No node with the given value!")
}

/*------------------------------------*/

func (lt *ConcretLinkTable) Clear() {
	lt.count = 0
	lt.head = nil
	lt.tail = nil
}
