# 题目

<p>给定一个整数数组 <code>temperatures</code> ，表示每天的温度，返回一个数组 <code>answer</code> ，其中 <code>answer[i]</code>&nbsp;是指在第 <code>i</code> 天之后，<span style="font-size:10.5pt"><span style="font-family:Calibri"><span style="font-size:10.5000pt"><span style="font-family:宋体"><font face="宋体">才会有更高的温度</font></span></span></span></span>。如果气温在这之后都不会升高，请在该位置用 <code>0</code> 来代替。</p>

<p><strong>示例 1:</strong></p>

<pre>
<strong>输入:</strong> <code>temperatures</code> = [73,74,75,71,69,72,76,73]
<strong>输出:</strong> [1,1,4,2,1,1,0,0]
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> temperatures = [30,40,50,60]
<strong>输出:</strong> [1,1,1,0]
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong> temperatures = [30,60,90]
<strong>输出: </strong>[1,1,0]</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= temperatures.length &lt;= 10<sup>5</sup></code></li>
	<li><code>30 &lt;= temperatures[i] &lt;= 100</code></li>
</ul>
<div><div>Related Topics</div><div><li>栈</li><li>数组</li><li>单调栈</li></div></div><br><div></div>

# 题解

本题是指从位置 i 之后第一个大于当前温度的位置与当前位置的差值。

单调栈通过维持栈内值的单调递增（递减）性，在整体 O(n) 的时间内处理需要大小比较的问题。

我们可以维持一个单调递减的栈，表示每天的温度；为了方便计算天数差，我们这里存放位置（即日期）而非温度本身。

我们从左向右遍历温度数组，对于每个日期 p，如果 p 的温度比栈顶存储位置 q 的温度高，则我们取出 q，并记录 q 需要等待的天数为 p − q；

我们重复这一过程，直到 p 的温度小于等于栈顶存储位置的温度（或空栈）时，我们将 p 插入栈顶，然后考虑下一天。

在这个过程中，栈内数组永远保持单调递减，避免了使用排序进行比较。最后若栈内剩余一些日期，则说明它们之后都没有出现更暖和的日期。

