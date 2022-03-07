#题目

<p>给定一个链表，返回链表开始入环的第一个节点。如果链表无环，则返回<code>null</code>。</p>

<p>为了表示给定链表中的环，我们使用整数 <code>pos</code> 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 <code>pos</code> 是 <code>-1</code>，则在该链表中没有环。<strong>注意，<code>pos</code> 仅仅是用于标识环的情况，并不会作为参数传递到函数中。</strong></p>

<p><strong>说明：</strong>不允许修改给定的链表。</p>

<p><strong>进阶：</strong></p>

<ul>
	<li>你是否可以使用 <code>O(1)</code> 空间解决此题？</li>
</ul>

<p><strong>示例 1：</strong></p>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png" style="height: 97px; width: 300px;" /></p>

<pre>
<strong>输入：</strong>head = [3,2,0,-4], pos = 1
<strong>输出：</strong>返回索引为 1 的链表节点
<strong>解释：</strong>链表中有一个环，其尾部连接到第二个节点。
</pre>

<p><strong>示例2：</strong></p>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png" style="height: 74px; width: 141px;" /></p>

<pre>
<strong>输入：</strong>head = [1,2], pos = 0
<strong>输出：</strong>返回索引为 0 的链表节点
<strong>解释：</strong>链表中有一个环，其尾部连接到第一个节点。
</pre>

<p><strong>示例 3：</strong></p>

<p><img alt="" src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png" style="height: 45px; width: 45px;" /></p>

<pre>
<strong>输入：</strong>head = [1], pos = -1
<strong>输出：</strong>返回 null
<strong>解释：</strong>链表中没有环。
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li>链表中节点的数目范围在范围 <code>[0, 10<sup>4</sup>]</code> 内</li>
	<li><code>-10<sup>5</sup> <= Node.val <= 10<sup>5</sup></code></li>
	<li><code>pos</code> 的值为 <code>-1</code> 或者链表中的一个有效索引</li>
</ul>
<div><div>Related Topics</div><div><li>哈希表</li><li>链表</li><li>双指针</li></div></div><br><div></div>

#题解

对于链表找环路的问题，有一个通用的解法——**快慢指针（Floyd 判圈法）**。给定两个指针，分别命名为 slow 和 fast，起始位置在链表的开头。每次 fast 前进两步，slow 前进一步。如果 fast 可以走到尽头，那么说明没有环路；如果 fast 可以无限走下去，那么说明一定有环路，且一定存在一个时刻 slow 和 fast 相遇。当 slow 和 fast 第一次相遇时，我们将 fast 重新移动到链表开头，并让 slow 和 fast 每次都前进一步。当 slow 和 fast 第二次相遇时，相遇的节点即为环路的开始点。

分析一下判断环的原理。fast 指针一次都 2 步，slow 指针一次走 1 步。令链表 head 到环的一个点需要 x1 步，从环的第一个点到相遇点需要 x2 步，从环中相遇点回到环的第一个点需要 x3 步。那么环的总长度是 x2 + x3 步。fast 和 slow 会相遇，说明他们走的时间是相同的，可以知道他们走的路程有以下的关系：

```
fast 的 t = (x1 + x2 + x3 + x2) / 2
slow 的 t = (x1 + x2) / 1
x1 + x2 + x3 + x2 = 2 * (x1 + x2)
所以 x1 = x3
```

所以 2 个指针相遇以后，如果 slow 继续往前走，fast 指针回到起点 head，两者都每次走一步，那么必定会在环的起点相遇，相遇以后输出这个点即是结果。