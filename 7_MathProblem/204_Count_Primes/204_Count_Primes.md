# 题目

<p>给定整数 <code>n</code> ，返回 <em>所有小于非负整数&nbsp;<code>n</code>&nbsp;的质数的数量</em> 。</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>n = 10
<strong>输出：</strong>4
<strong>解释：</strong>小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>n = 0
<strong>输出：</strong>0
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>n = 1
<strong>输出</strong>：0
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= n &lt;= 5 * 10<sup>6</sup></code></li>
</ul>
<div><div>Related Topics</div><div><li>数组</li><li>数学</li><li>枚举</li><li>数论</li></div></div><br><div></div>

# 题解

从 1 到 n 遍历，假设当前遍历到 m，则把所有小于 n 的、且是 m 的倍数的整数标为合数；遍历完成后，没有被标为和数的数字即为质数。

