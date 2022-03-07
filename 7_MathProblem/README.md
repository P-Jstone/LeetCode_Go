# 数学问题
## 7.1 公倍数与公因数
利用<u>辗转相除法</u>，我们可以很方便地求得两个数的最大公因数（greatest common divisor，gcd）；将两个数相乘再除以最大公因数即可得到最小公倍数（least common multiple, lcm）。

```
func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

func lcm(a, b int) int {
    return a*b/gcd(a, b)
}
```

进一步地，我们也可以通过<u>扩展欧几里得算法</u>（extended gcd）在求得 a 和 b 最大公因数的同时，也得到它们的系数 x 和 y，从而使 ax + by = gcd(a, b)。

```
func xGCD(a, b int, x, y *int) int {
    if !b {
        x = 1, y = 0
        return a
    }
    var x1, y1 int
    gcd := xGCD(b, a%b, x1, y1)
    x = y1, y = x1-(a/b)*y1
    return gcd
}
```

## 7.2 质数
质数又称素数，指的是指在大于 1 的自然数中，除了 1 和它本身以外不再有其他因数的自然数。值得注意的是，每一个数都可以分解成质数的乘积。

<u>埃拉托斯特尼筛法</u>（Sieve of Eratosthenes，简称埃氏筛法）是非常常用的，判断一个整数是否是质数的方法。并且它可以在判断一个整数 n 时，同时判断所小于 n 的整数。其原理也十分易懂：从 1 到 n 遍历，假设当前遍历到 m，则把所有小于 n 的、且是 m 的倍数的整数标为合数；遍历完成后，没有被标为合数的数字即为质数。

## 7.3 数字处理
504、172、415、326
