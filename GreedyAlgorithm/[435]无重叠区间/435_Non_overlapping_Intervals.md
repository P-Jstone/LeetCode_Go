#题目

<p>给定一个区间的集合<code>intervals</code>，其中 <code>intervals[i] = [start<sub>i</sub>, end<sub>i</sub>]</code>。返回 <em>需要移除区间的最小数量，使剩余区间互不重叠</em>。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> intervals = [[1,2],[2,3],[3,4],[1,3]]
<strong>输出:</strong> 1
<strong>解释:</strong> 移除 [1,3] 后，剩下的区间没有重叠。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> intervals = [ [1,2], [1,2], [1,2] ]
<strong>输出:</strong> 2
<strong>解释:</strong> 你需要移除两个 [1,2] 来使剩下的区间没有重叠。
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> intervals = [ [1,2], [2,3] ]
<strong>输出:</strong> 0
<strong>解释:</strong> 你不需要移除任何区间，因为它们已经是无重叠的了。
</pre>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= intervals.length &lt;= 10<sup>5</sup></code></li>
	<li><code>intervals[i].length == 2</code></li>
	<li><code>-5 * 10<sup>4</sup>&lt;= start<sub>i</sub>&lt; end<sub>i</sub>&lt;= 5 * 10<sup>4</sup></code></li>
</ul>
<div><div>Related Topics</div><div><li>贪心</li><li>数组</li><li>动态规划</li><li>排序</li></div></div><br><div></div>

#题解

求最少的移除区间个数，等价于尽量多保留不重叠的区间。在选择要保留区间时，区间的结尾十分重要：选择的区间结尾越小，余留给其它区间的空间就越大，就越能保留更多的区间。因此，我们采取的贪心策略为，优先保留结尾小且不相交的区间。
 
-  先将区间按照结尾的大小进行增序排序，每次选择结尾最小且和前一个选择的区间不重叠的区间
