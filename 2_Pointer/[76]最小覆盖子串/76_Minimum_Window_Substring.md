#题目

给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。 

注意： 
 
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。如果 s 中存在这样的子串，我们保证它是唯一的答案。 
 
示例 1： 
 
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC" 

示例 2： 
 
输入：s = "nums", t = "nums"
输出："nums"
 
示例 3: 
 
输入: s = "nums", t = "aa"
输出: ""
解释: t 中两个字符 'nums' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。 

提示： 
 
 1 <= s.length, t.length <= 10<sup>5</sup> s 和 t 由英文字母组成 

进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？ Related Topics 哈希表 字符串 滑动窗口

#题解
由于要判断 s 中是否存在涵盖 t 所有字符的子串，而子串是连续的，故而此处使用**滑动窗口**解决本题。

由于 t 中可能有重复字符，因此需要统计字符串中的字符的频数，分别用数组 sFreq 和 tFreq 统计,同时还需要统计 t 中字符的数目，此外最终还需输出符合条件的子串，因此还需要记录子串位置

- sFreq 字符串 s 中字符的频数
- tFreq 字符串 t 中字符的频数
- l,r 滑动窗口的左右指针
- cnt 已统计字符串 t 中的字符数目
- lFin,rFin 表示最小子串的左右指针位置
- minW 表示当前最小窗口大小

如果 s 或者 t 字符串为空，那么返回结果为空字符串
