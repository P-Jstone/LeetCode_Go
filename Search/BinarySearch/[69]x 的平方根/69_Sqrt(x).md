#题目
<p>给你一个非负整数 <code>x</code> ，计算并返回<code>x</code>的 <strong>算术平方根</strong> 。</p>

<p>由于返回类型是整数，结果只保留 <strong>整数部分 </strong>，小数部分将被 <strong>舍去 。</strong></p>

<p><strong>注意：</strong>不允许使用任何内置指数函数和算符，例如 <code>pow(x, 0.5)</code> 或者 <code>x ** 0.5</code> 。</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>x = 4
<strong>输出：</strong>2
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>x = 8
<strong>输出：</strong>2
<strong>解释：</strong>8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
</pre>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= x &lt;= 2<sup>31</sup> - 1</code></li>
</ul>
<div><div>Related Topics</div><div><li>数学</li><li>二分查找</li></div></div><br><div></div>

#题解

##解法一:二分查找

根号 x 的取值范围一定在 [0,x] 之间，这个区间内的值是递增有序的，有边界的，可以用下标访问的，满足这三点正好也就满足了二分搜索的 3 大条件。所以解题思路一，二分搜索

##解法二：牛顿迭代法

求根号 x，即求满足 <code>x<sup>2</sup> - n = 0</code> 方程的所有解。  

一个曲线方程 <code>f(x)</code> ，在它的 <code>f(x<sub>n</sub>)</code> 处画一条切线与 x 轴相交，交点为 <code>x<sub>n+1</sub></code> ，如果继续在它的 <code>f(x<sub>n+1</sub>)</code> 处画一条切线与 x 轴相交，会得到交点 <code>x<sub>n+2</sub></code> .而在这个过程中，可以发现交点 <code>x<sub>n+m</sub></code> 会无限逼近方程 <code>f(x)=0</code> 的解，最终可以得到一个与理想值无限靠近的解。  

而此处讨论的是求平方根，所以曲线方程更简单。例如要求 N 的平方根。

那么就是求方程方程 <code>f(x)=x<sup>2</sup>-N</code> ，当 <code>f(x)=0</code>的解。

函数  <code>f(x)</code> 的导函数是： <code>f'(x)=2x</code> 

那么函数 <code>f(x)</code> 的曲线在 <code>(x<sub>n</sub>,x<sub>n</sub><sup>2</sup>-N)</code> 点出切线的斜率为： <code>2x<sub>n</sub></code> 

所以切线方程为： <code>f(x)-(x<sub>n</sub><sup>2</sup>-N)=2x<sub>n</sub>(x-x<sub>n</sub>)</code> ，即： <code>f(x)=2x<sub>n</sub>x-x<sub>n</sub><sup>2</sup>-N</code> 

那么切线方程与 x 轴的交点 <code>x<sub>n+1</sub>=(x<sub>n</sub>+N/x<sub>n</sub>)/2</code>

可以将得到的交点值的平方与N比较，循环以上过程直到得到满意的值。
