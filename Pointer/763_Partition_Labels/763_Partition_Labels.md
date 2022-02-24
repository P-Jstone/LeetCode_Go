#题目

<p>字符串 <code>S</code> 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。</p>

<p><strong>示例：</strong></p>

<pre>
<strong>输入：</strong>S = "ababcbacadefegdehijhklij"
<strong>输出：</strong>[9,7,8]
<strong>解释：</strong>
划分结果为 "ababcbaca", "defegde", "hijhklij"。
每个字母最多出现在一个片段中。
像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>S</code>的长度在<code>[1, 500]</code>之间。</li>
	<li><code>S</code>只包含小写字母 <code>'a'</code> 到 <code>'z'</code> 。</li>
</ul>
<div><div>Related Topics</div><div><li>贪心</li><li>哈希表</li><li>双指针</li><li>字符串</li></div></div><br><div></div>

#题解

此题有 2 种思路：
- 第一种思路是先记录下每个字母的出现次数，然后对滑动窗口中的每个字母判断次数是否用尽为 0，如果这个窗口内的所有字母次数都为 0，这个窗口就是符合条件的窗口。时间复杂度为 O(n^2)
- 另外一种思路是记录下每个字符最后一次出现的下标，这样就不用记录次数。在每个滑动窗口中，依次判断每个字母最后一次出现的位置，如果在一个下标内，所有字母的最后一次出现的位置都包含进来了，那么这个下标就是这个满足条件的窗口大小。时间复杂度为 O(n^2)