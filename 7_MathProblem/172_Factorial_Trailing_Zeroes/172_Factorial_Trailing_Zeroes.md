# 题目

<p>给定一个整数 <code>n</code> ，返回 <code>n!</code> 结果中尾随零的数量。</p>

<p>提示<code>n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1</code></p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>n = 3
<strong>输出：</strong>0
<strong>解释：</strong>3! = 6 ，不含尾随 0
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>n = 5
<strong>输出：</strong>1
<strong>解释：</strong>5! = 120 ，有一个尾随 0
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>n = 0
<strong>输出：</strong>0
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= n &lt;= 10<sup>4</sup></code></li>
</ul>

<p><b>进阶：</b>你可以设计并实现对数时间复杂度的算法来解决此问题吗？</p>
<div><div>Related Topics</div><div><li>数学</li></div></div><br><div></div>

# 题解

每个尾部的 0 由 2 × 5 = 10 而来，因此我们可以把阶乘的每一个元素拆成质数相乘，统计有多少个 2 和 5。

每两个数字就会多⼀个质因数 2，⽽每五个数字才多⼀个质因数 5。每 5 个数字就会多⼀个质因数 5。0~4 的阶乘⾥没有质因数 5，5~9 的阶乘⾥有 1 个质因数 5，10~14 的阶乘⾥有 2 个质因数 5，依此类推。明显的，质因子 2 的数量远多于质因子 5 的数量，因此我们可以只统计阶乘结果里有多少个质因子 5。

N! 有多少个后缀 0，即 N! 有多少个质因数 5。N! 有多少个质因数 5，即 N 可以划分成多少组 5个数字⼀组，加上划分成多少组 25 个数字⼀组，加上划分多少组成 125 个数字⼀组，等等。即 res = N/5 + N/(5^2) + N/(5^3) + ... = ((N / 5) / 5) / 5 /... 。最终算法复杂度为 O(logN)。
