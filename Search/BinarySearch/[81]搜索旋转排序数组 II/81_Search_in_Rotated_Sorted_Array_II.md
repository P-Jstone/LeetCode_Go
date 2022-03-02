#题目

<p>已知存在一个按非降序排列的整数数组 <code>nums</code> ，数组中的值不必互不相同。</p>

<p>在传递给函数之前，<code>nums</code> 在预先未知的某个下标 <code>k</code>（<code>0 &lt;= k &lt; nums.length</code>）上进行了 <strong>旋转 </strong>，使数组变为 <code>[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]</code>（下标 <strong>从 0 开始</strong> 计数）。例如， <code>[0,1,2,4,4,4,5,6,6,7]</code> 在下标 <code>5</code> 处经旋转后可能变为 <code>[4,5,6,6,7,0,1,2,4,4]</code> 。</p>

<p>给你 <strong>旋转后</strong> 的数组 <code>nums</code> 和一个整数 <code>target</code> ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 <code>nums</code> 中存在这个目标值 <code>target</code> ，则返回 <code>true</code> ，否则返回 <code>false</code> 。</p>

<p>你必须尽可能减少整个操作步骤。</p>

<p><strong>示例1：</strong></p>

<pre>
<strong>输入：</strong>nums = [2<code>,5,6,0,0,1,2]</code>, target = 0
<strong>输出：</strong>true
</pre>

<p><strong>示例2：</strong></p>

<pre>
<strong>输入：</strong>nums = [2<code>,5,6,0,0,1,2]</code>, target = 3
<strong>输出：</strong>false</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 5000</code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
	<li>题目数据保证 <code>nums</code> 在预先未知的某个下标上进行了旋转</li>
	<li><code>-10<sup>4</sup> &lt;= target &lt;= 10<sup>4</sup></code></li>
</ul>

<p><strong>进阶：</strong></p>

<ul>
	<li>这是 <a href="https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/">搜索旋转排序数组</a>的延伸题目，本题中的<code>nums</code> 可能包含重复元素。</li>
	<li>这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？</li>
</ul>

<div><div>Related Topics</div><div><li>数组</li><li>二分查找</li></div></div><br><div></div>

#题解

即使数组被旋转过，我们仍然可以利用这个数组的递增性，使用二分查找。对于当前的中点，如果它指向的值小于等于右端，那么说明右区间是排好序的；反之，那么说明左区间是排好序的。如果目标值位于排好序的区间内，我们可以对这个区间继续二分查找；反之，我们对于另一半区间继续二分查找。

注意，因为数组存在重复数字，如果中点和左端的数字相同，我们并不能确定是左区间全部相同，还是右区间完全相同。在这种情况下，我们可以简单地将左端点右移一位，然后继续进行二分查找。
