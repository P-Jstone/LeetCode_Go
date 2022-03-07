#题目

<p>假设你正在爬楼梯。需要 <code>n</code> 阶你才能到达楼顶。</p>

<p>每次你可以爬 <code>1</code> 或 <code>2</code> 个台阶。你有多少种不同的方法可以爬到楼顶呢？</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>n = 2
<strong>输出：</strong>2
<strong>解释：</strong>有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>n = 3
<strong>输出：</strong>3
<strong>解释：</strong>有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= n &lt;= 45</code></li>
</ul>
<div><div>Related Topics</div><div><li>记忆化搜索</li><li>数学</li><li>动态规划</li></div></div><br><div></div>

#题解

本题属于斐波那契数列题。

1. 定义一个数组 dp，`dp[i]` 表示走到第 i 阶的方法数。

2. 因为我们每次可以走一步或者两步，所以第 i 阶可以从第 i-1 或 i-2 阶到达。换句话说，走到第 i 阶的方法数即为走到第 `i-1` 阶的方法数加上走到第 `i-2` 阶的方法数。这样我们就得到了状态转移方程`dp[i] = dp[i-1] + dp[i-2]`。

3. 注意边界条件的处理。初始化`dp[1] = 1,dp[2] = 2`。

4. 递推公式dp[i] = dp[i - 1] + dp[i - 2];中可以看出，遍历顺序一定是从前向后遍历的。

5. n为5时，情况如下

   | dp1  | dp2  | dp3  | dp4  | dp5  |
   | ---- | ---- | ---- | ---- | ---- |
   | 1    | 2    | 3    | 5    | 8    |

本题可以只使用两个变量来处理，以减少内存占用。

