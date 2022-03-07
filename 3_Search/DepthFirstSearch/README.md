#深度优先搜索

<u>**深度优先搜索**</u>（depth-first seach，DFS）在搜索到一个新的节点时，立即对该新节点进行遍历；因此遍历需要用<u>先入后出</u>的栈来实现，也可以通过与栈等价的<u>递归</u>来实现。对于树结构而言，由于总是对新节点调用遍历，因此看起来是向着“深”的方向前进。

考虑如下一颗简单的树。我们从 1 号节点开始遍历，假如遍历顺序是从左子节点到右子节点，那么按照优先向着“深”的方向前进的策略，假如我们使用递归实现，我们的遍历过程为 1（起始节点）->2（遍历更深一层的左子节点）->4（遍历更深一层的左子节点）->2（无子节点，返回父结点）->1（子节点均已完成遍历，返回父结点）->3（遍历更深一层的右子节点）->1（无子节点，返回父结点）-> 结束程序（子节点均已完成遍历）。如果我们使用栈实现，我们的栈顶元素的变化过程为 1->2->4->3。

```
    1
   / \
  2   3
 /
4
```
深度优先搜索也可以用来**检测环路**：记录每个遍历过的节点的父节点，若一个节点被再次遍历且父节点不同，则说明有环。我们也可以用之后会讲到的拓扑排序判断是否有环路，若最后存在入度不为零的点，则说明有环。

有时我们可能会需要对已经搜索过的节点进行标记，以防止在遍历时重复搜索某个节点，这种做法叫做<u>状态记录</u>或<u>记忆化</u>（memoization）。

#DFS模板

只是一个例子，不对应任何题
```
func combinationSum2(candidates []int, target int) [][]int {
    if len(candidates) == 0 {
        return [][]int{}
    }
    c, res := []int{}, [][]int{}
    sort.Ints(candidates)
    findcombinationSum2(candidates, target, 0, c, &res)
    return res
}

func findcombinationSum2(nums []int, target, index int, c []int, res *[][]int) {
    if target == 0 {
        b := make([]int, len(c))
        copy(b, c)
        *res = append(*res, b)
        return
    }
    for i := index; i < len(nums); i++ {
        if i > index && nums[i] == nums[i-1] { // 这⾥是去重的关键逻辑
            continue
        }
        if target >= nums[i] {
            c = append(c, nums[i])
            findcombinationSum2(nums, target-nums[i], i+1, c, res)
            c = c[:len(c)-1]
        }
    }
}
```