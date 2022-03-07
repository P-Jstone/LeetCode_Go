# 题目

<p>按字典<code>wordList</code> 完成从单词 <code>beginWord</code> 到单词 <code>endWord</code> 转化，一个表示此过程的 <strong>转换序列</strong> 是形式上像 <code>beginWord -&gt; s<sub>1</sub> -&gt; s<sub>2</sub> -&gt; ... -&gt; s<sub>k</sub></code> 这样的单词序列，并满足：</p>

<div class="original__bRMd">
<div>
<ul>
	<li>每对相邻的单词之间仅有单个字母不同。</li>
	<li>转换过程中的每个单词 <code>s<sub>i</sub></code>（<code>1 &lt;= i &lt;= k</code>）必须是字典&nbsp;<code>wordList</code> 中的单词。注意，<code>beginWord</code> 不必是字典 <code>wordList</code> 中的单词。</li>
	<li><code>s<sub>k</sub> == endWord</code></li>
</ul>

<p>给你两个单词 <code>beginWord</code> 和 <code>endWord</code> ，以及一个字典 <code>wordList</code> 。请你找出并返回所有从 <code>beginWord</code> 到 <code>endWord</code> 的 <strong>最短转换序列</strong> ，如果不存在这样的转换序列，返回一个空列表。每个序列都应该以单词列表<em> </em><code>[beginWord, s<sub>1</sub>, s<sub>2</sub>, ..., s<sub>k</sub>]</code> 的形式返回。</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
<strong>输出：</strong>[["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
<strong>解释：</strong>存在 2 种最短的转换序列：
"hit" -&gt; "hot" -&gt; "dot" -&gt; "dog" -&gt; "cog"
"hit" -&gt; "hot" -&gt; "lot" -&gt; "log" -&gt; "cog"
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
<strong>输出：</strong>[]
<strong>解释：</strong>endWord "cog" 不在字典 wordList 中，所以不存在符合要求的转换序列。
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= beginWord.length &lt;= 5</code></li>
	<li><code>endWord.length == beginWord.length</code></li>
	<li><code>1 &lt;= wordList.length &lt;= 5000</code></li>
	<li><code>wordList[i].length == beginWord.length</code></li>
	<li><code>beginWord</code>、<code>endWord</code> 和 <code>wordList[i]</code> 由小写英文字母组成</li>
	<li><code>beginWord != endWord</code></li>
	<li><code>wordList</code> 中的所有单词 <strong>互不相同</strong></li>
</ul>
</div>
</div>
<div><div>Related Topics</div><div><li>广度优先搜索</li><li>哈希表</li><li>字符串</li><li>回溯</li></div></div><br><div></div>

# 题解

将起始字符串、终止字符串和单词表中所有的字符串视为节点，如果两个字符串只有一个字符不同，则两个节点相连。由于要输出修改次数最少的所有修改方式，因此使用广度优先搜索，求起始节点到终止节点的最短距离。

从起始节点到终止节点单向广度优先搜索，可以解决问题，而从起始节点和终止节点双向广度优先搜索可以减少搜索的节点数，搜索结束后使用回溯算法重建路径。
