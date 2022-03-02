#题目

<p>在给定的二维二进制数组<code>A</code>中，存在两座岛。（岛是由四面相连的 <code>1</code> 形成的一个最大组。）</p>

<p>现在，我们可以将<code>0</code>变为<code>1</code>，以使两座岛连接起来，变成一座岛。</p>

<p>返回必须翻转的<code>0</code> 的最小数目。（可以保证答案至少是 <code>1</code> 。）</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>A = [[0,1],[1,0]]
<strong>输出：</strong>1
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>A = [[0,1,0],[0,0,0],[0,0,1]]
<strong>输出：</strong>2
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>A = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
<strong>输出：</strong>1</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 <= A.length == A[0].length <= 100</code></li>
	<li><code>A[i][j] == 0</code> 或 <code>A[i][j] == 1</code></li>
</ul>
<div><div>Related Topics</div><div><li>深度优先搜索</li><li>广度优先搜索</li><li>数组</li><li>矩阵</li></div></div><br><div></div>

#题解

我们使用的方法非常直接：首先找到这两座岛，随后选择一座，将它不断向外延伸一圈，直到到达了另一座岛。

在寻找这两座岛时，我们使用深度优先搜索。在向外延伸时，我们使用广度优先搜索。

我们通过对数组 A 中的 1 进行深度优先搜索，可以得到两座岛的位置集合，分别为 source 和 target。随后我们从 source 中的所有位置开始进行广度优先搜索，当它们到达了 target 中的任意一个位置时，搜索的层数就是答案。

