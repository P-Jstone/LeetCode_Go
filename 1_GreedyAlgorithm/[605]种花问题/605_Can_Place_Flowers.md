#题目

<p>假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。</p>

<p>给你一个整数数组<code>flowerbed</code> 表示花坛，由若干 <code>0</code> 和 <code>1</code> 组成，其中 <code>0</code> 表示没种植花，<code>1</code> 表示种植了花。另有一个数<code>n</code><strong> </strong>，能否在不打破种植规则的情况下种入<code>n</code><strong></strong>朵花？能则返回 <code>true</code> ，不能则返回 <code>false</code>。</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>flowerbed = [1,0,0,0,1], n = 1
<strong>输出：</strong>true
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>flowerbed = [1,0,0,0,1], n = 2
<strong>输出：</strong>false
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= flowerbed.length <= 2 * 10<sup>4</sup></code></li>
	<li><code>flowerbed[i]</code> 为 <code>0</code> 或 <code>1</code></li>
	<li><code>flowerbed</code> 中不存在相邻的两朵花</li>
	<li><code>0 <= n <= flowerbed.length</code></li>
</ul>
<div><div>Related Topics</div><div><li>贪心</li><li>数组</li></div></div><br><div></div>

#题解

贪心的解法是在遵守规则的情况下尽可能多的种花，也就是从头遍历数组只要能够种花就直接种植，在此过程中有以下情况：  
1. 当前位置i的值是1，那么最近可以种植的位置是i+2；
2. 如果当前位置i的值不是1，那么就是0，如果当前的位置是花坛末尾，或者下一个位置的值为0，那么可以种花，此时将i+2；
3. 如果当前位置的值是0，但是下一个位置的值为1，那么i+3