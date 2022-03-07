# 题目

<p>给定一个整数，写一个函数来判断它是否是 3&nbsp;的幂次方。如果是，返回 <code>true</code> ；否则，返回 <code>false</code> 。</p>

<p>整数 <code>n</code> 是 3 的幂次方需满足：存在整数 <code>x</code> 使得 <code>n == 3<sup>x</sup></code></p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>n = 27
<strong>输出：</strong>true
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>n = 0
<strong>输出：</strong>false
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>n = 9
<strong>输出：</strong>true
</pre>

<p><strong>示例 4：</strong></p>

<pre>
<strong>输入：</strong>n = 45
<strong>输出：</strong>false
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>-2<sup>31</sup> &lt;= n &lt;= 2<sup>31</sup> - 1</code></li>
</ul>

<p><strong>进阶：</strong>你能不使用循环或者递归来完成本题吗？</p>
<div><div>Related Topics</div><div><li>递归</li><li>数学</li></div></div><br><div></div>

# 题解

## 方法一：试除法

我们不断地将 n 除以 3，直到 n=1。如果此过程中 n 无法被 3 整除，就说明 n 不是 3 的幂。

本题中的 n 可以为负数或 0，可以直接提前判断该情况并返回 False，也可以进行试除，因为负数或 0 也无法通过多次除以 3 得到 1。

## 方法二：判断是否为最大 33 的幂的约数

我们还可以使用一种较为取巧的做法。

在题目给定的 32 位有符号整数的范围内，最大的 3 的幂为 <code>3<sup>19</sup> = 1162261467</code>。我们只需要判断 n 是否是 <code>3<sup>19</sup></code>的约数即可。

与方法一不同的是，这里需要特殊判断 nn 是负数或 00 的情况。


