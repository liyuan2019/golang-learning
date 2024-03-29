package sort

import "fmt"

/**

根节点/ 什么都不表示
做一个字典比如a-z 字母表
没一个节点包含这26个字母的字典表，每个位置保存下个节点的指针。
缺点：
trie树比较消耗内存：因为他没一层都保存一个字典表，就算这层的节点只有一个也要有一组表
使用的是指针类型，内存不连续对存储不友好，性能打折扣 优点：
查询效率比较高，对于一些范围较小的或者内存不敏感的应用可以使用
特别适用自动补全类应用
*/

type TrieNode struct {
	// value      int
	dictionary [26]*TrieNode
}

type TrieTree struct {
	root *TrieNode
}

func CreateTree() TrieTree {
	arrList := []string{"how", "hi", "her", "hello", "so", "see"}
	myTree := TrieTree{}
	//添加跟节点
	myTree.root = &TrieNode{}
	for _, value := range arrList {
		myTree.addWord(value)
	}
	return myTree
}
func (t *TrieTree) addWord(word string) {
	fmt.Println(word)
	nowNode := t.root
	a := int('a')
	var char int
	for i := 0; i < len(word); i++ {
		char = int(word[i])
		if nowNode.dictionary[char-a] != nil {
			nowNode = nowNode.dictionary[char-a]
			continue
		} else {
			newNode := &TrieNode{}
			nowNode.dictionary[char-a] = newNode
			nowNode = newNode
			continue
		}
	}
}

func (t *TrieTree) findWord(word string) int {
	nowNode := t.root
	a := int('a')
	var char int
	for i := 0; i < len(word); i++ {
		char = int(word[i])
		if nowNode.dictionary[char-a] != nil {
			return 0
		} else {
			nowNode = nowNode.dictionary[char-a]
		}
		if i == len(word)-1 {
			return 1
		}
	}
	return 0
}

// func main() {
// 	tree := createTree()
// 	//fmt.Println(tree)
// 	flag := tree.findWord("her")
// 	fmt.Println(flag)
// 	flag = tree.findWord("x")
// 	fmt.Println(flag)
// }
