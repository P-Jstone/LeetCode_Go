# 题目：实现 strStr()

<p>实现 <a href="https://baike.baidu.com/item/strstr/811469" target="_blank">strStr()</a> 函数。</p>

<p>给你两个字符串 <code>haystack</code> 和 <code>needle</code> ，请你在 <code>haystack</code> 字符串中找出 <code>needle</code> 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  <code>-1</code><strong> </strong>。</p>

<p><strong>说明：</strong></p>

<p>当 <code>needle</code> 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。</p>

<p>对于本题而言，当 <code>needle</code> 是空字符串时我们应当返回 0 。这与 C 语言的 <a href="https://baike.baidu.com/item/strstr/811469" target="_blank">strstr()</a> 以及 Java 的 <a href="https://docs.oracle.com/javase/7/docs/api/java/lang/String.html#indexOf(java.lang.String)" target="_blank">indexOf()</a> 定义相符。</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>haystack = "hello", needle = "ll"
<strong>输出：</strong>2
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>haystack = "aaaaa", needle = "bba"
<strong>输出：</strong>-1
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>haystack = "", needle = ""
<strong>输出：</strong>0
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 <= haystack.length, needle.length <= 5 * 10<sup>4</sup></code></li>
	<li><code>haystack</code> 和 <code>needle</code> 仅由小写英文字符组成</li>
</ul>
<div><div>Related Topics</div><div><li>双指针</li><li>字符串</li><li>字符串匹配</li></div></div><br><div><li>👍 1019</li><li>👎 0</li></div>

# 题解

使用著名的Knuth-Morris-Pratt（KMP）算法，可以在 O(m + n) 时间利用动态规划完成匹配。
KMP 解法
KMP 算法是一个快速查找匹配串的算法，它的作用其实就是本题问题：如何快速在「原字符串」中找到「匹配字符串」。

上述的朴素解法，不考虑剪枝的话复杂度是 O(m∗n) 的，而 KMP 算法的复杂度为 O(m+n)。

KMP 之所以能够在 O(m+n) 复杂度内完成查找，是因为其能在「非完全匹配」的过程中提取到有效信息进行复用，以减少「重复匹配」的消耗。

你可能不太理解，没关系，我们可以通过举 🌰 来理解 KMP。

1. 匹配过程
   在模拟 KMP 匹配过程之前，我们先建立两个概念：

- 前缀：对于字符串 abcxxxxefg，我们称 abc 属于 abcxxxxefg 的某个前缀。
- 后缀：对于字符串 abcxxxxefg，我们称 efg 属于 abcxxxxefg 的某个后缀。
然后我们假设原串为 abeababeabf，匹配串为 abeabf：

<img alt="" src="https://pic.leetcode-cn.com/1618739635-lrhElP-image.png" />

我们可以先看看如果不使用 KMP，会如何进行匹配（不使用 substring 函数的情况下）。

首先在「原串」和「匹配串」分别各自有一个指针指向当前匹配的位置。

首次匹配的「发起点」是第一个字符 a。显然，后面的 abeab 都是匹配的，两个指针会同时往右移动（黑标）。

在都能匹配上 abeab 的部分，「朴素匹配」和「KMP」并无不同。

直到出现第一个不同的位置（红标）：

<img alt="" src="https://pic.leetcode-cn.com/1618741727-pqXsfg-image.png" />

接下来，正是「朴素匹配」和「KMP」出现不同的地方：

先看下「朴素匹配」逻辑：
1. 将原串的指针移动至本次「发起点」的下一个位置（b 字符处）；匹配串的指针移动至起始位置。

2. 尝试匹配，发现对不上，原串的指针会一直往后移动，直到能够与匹配串对上位置。

如图：

<img alt="" src="https://pic.leetcode-cn.com/1618742678-lTXSgV-image.png" />

也就是说，对于「朴素匹配」而言，一旦匹配失败，将会将原串指针调整至下一个「发起点」，匹配串的指针调整至起始位置，然后重新尝试匹配。

这也就不难理解为什么「朴素匹配」的复杂度是 O(m∗n) 了。

- 然后我们再看看「KMP 匹配」过程：

首先匹配串会检查之前已经匹配成功的部分中里是否存在相同的「前缀」和「后缀」。如果存在，则跳转到「前缀」的下一个位置继续往下匹配：

<img alt="" src="https://pic.leetcode-cn.com/1618845342-ydYJRp-9364346F937803F03CD1A0AE645EA0F1.jpg" />

跳转到下一匹配位置后，尝试匹配，发现两个指针的字符对不上，并且此时匹配串指针前面不存在相同的「前缀」和「后缀」，这时候只能回到匹配串的起始位置重新开始：

<img alt="" src="https://pic.leetcode-cn.com/1618755191-ddejks-image.png" />

到这里，你应该清楚 KMP 为什么相比于朴素解法更快：

因为 KMP 利用已匹配部分中相同的「前缀」和「后缀」来加速下一次的匹配。

因为 KMP 的原串指针不会进行回溯（没有朴素匹配中回到下一个「发起点」的过程）。

第一点很直观，也很好理解。

我们可以把重点放在第二点上，原串不回溯至「发起点」意味着什么？

其实是意味着：随着匹配过程的进行，原串指针的不断右移，我们本质上是在不断地在否决一些「不可能」的方案。

当我们的原串指针从 i 位置后移到 j 位置，不仅仅代表着「原串」下标范围为 [i,j)[i,j) 的字符与「匹配串」匹配或者不匹配，更是在否决那些以「原串」下标范围为 [i,j)[i,j) 为「匹配发起点」的子集。

2. 分析实现
   到这里，就结束了吗？要开始动手实现上述匹配过程了吗？

我们可以先分析一下复杂度。如果严格按照上述解法的话，最坏情况下我们需要扫描整个原串，复杂度为 O(n)。同时在每一次匹配失败时，去检查已匹配部分的相同「前缀」和「后缀」，跳转到相应的位置，如果不匹配则再检查前面部分是否有相同「前缀」和「后缀」，再跳转到相应的位置 ... 这部分的复杂度是 O(m^2) ，因此整体的复杂度是 O(n * m^2)，而我们的朴素解法是 O(m * n)的。

说明还有一些性质我们没有利用到。

显然，扫描完整原串操作这一操作是不可避免的，我们可以优化的只能是「检查已匹配部分的相同前缀和后缀」这一过程。

再进一步，我们检查「前缀」和「后缀」的目的其实是「为了确定匹配串中的下一段开始匹配的位置」。

同时我们发现，对于匹配串的任意一个位置而言，由该位置发起的下一个匹配点位置其实与原串无关。

举个 🌰，对于匹配串 abcabd 的字符 d 而言，由它发起的下一个匹配点跳转必然是字符 c 的位置。因为字符 d 位置的相同「前缀」和「后缀」字符 ab 的下一位置就是字符 c。

可见从匹配串某个位置跳转下一个匹配位置这一过程是与原串无关的，我们将这一过程称为找 next 点。

显然我们可以预处理出 next 数组，数组中每个位置的值就是该下标应该跳转的目标位置（ next 点）。

当我们进行了这一步优化之后，复杂度是多少呢？

预处理 next 数组的复杂度未知，匹配过程最多扫描完整个原串，复杂度为 O(n)。

因此如果我们希望整个 KMP 过程是 O(m + n) 的话，那么我们需要在 O(m) 的复杂度内预处理出 next数组。

所以我们的重点在于如何在 O(m) 复杂度内处理处 next 数组。

3. next 数组的构建

接下来，我们看看 next 数组是如何在 O(m) 的复杂度内被预处理出来的。

假设有匹配串 aaabbab，我们来看看对应的 next 是如何被构建出来的。

<img alt="" src="https://pic.leetcode-cn.com/1618846927-xFAEXE-010FD8AE2B79FFE03DC3735ACD224A6A.png" />

<img alt="" src="https://pic.leetcode-cn.com/1618847960-lkVIDM-B9497542844478144BED83E9ADA0C12F.png" />

<img alt="" src="https://pic.leetcode-cn.com/1618847981-wncoqJ-161584A2D930A7B91092A2C3872D9DE5.png" />

<img alt="" src="https://pic.leetcode-cn.com/1618847995-vRWimV-6127EBA37435560C20BB8B15D5B790B6.png" />

这就是整个 next 数组的构建过程，时空复杂度均为 O(m)。

至此整个 KMP 匹配过程复杂度是 O(m + n) 的。