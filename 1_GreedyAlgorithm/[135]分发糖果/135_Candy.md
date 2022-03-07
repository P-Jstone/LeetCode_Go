#题目

<p><code>n</code> 个孩子站成一排。给你一个整数数组 <code>ratings</code> 表示每个孩子的评分。</p>

<p>你需要按照以下要求，给这些孩子分发糖果：</p>

<ul>
	<li>每个孩子至少分配到 <code>1</code> 个糖果。</li>
	<li>相邻两个孩子评分更高的孩子会获得更多的糖果。</li>
</ul>

<p>请你给每个孩子分发糖果，计算并返回需要准备的 <strong>最少糖果数目</strong> 。</p>

<p><strong>示例1：</strong></p>

<pre>
<strong>输入：</strong>ratings = [1,0,2]
<strong>输出：</strong>5
<strong>解释：</strong>你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
</pre>

<p><strong>示例2：</strong></p>

<pre>
<strong>输入：</strong>ratings = [1,2,2]
<strong>输出：</strong>4
<strong>解释：</strong>你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
     第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>n == ratings.length</code></li>
	<li><code>1 &lt;= n &lt;= 2 * 10<sup>4</sup></code></li>
	<li><code>0 &lt;= ratings[i] &lt;= 2 * 10<sup>4</sup></code></li>
</ul>
<div><div>Related Topics</div><div><li>贪心</li><li>数组</li></div></div><br><div></div>

#题解
- 本题的输入数据是孩子们的评分数组，所有的孩子最少要获得一颗糖；
- 从左向右遍历，如果当前孩子的评分比前一个孩子的评分要高，那么当前孩子的糖果数目在前一个孩子的糖果数目基础上加一，否则继续遍历；
- 当遍历结束后再从右往左进行遍历，如果左边孩子的评分大于右边的孩子，并且左边孩子的糖果数不大于右边孩子的糖果数，则左边孩子孩子的糖果数在右边孩子糖果的基础上加一。  
- 经过两次遍历分配的糖果满足题目要求了。


