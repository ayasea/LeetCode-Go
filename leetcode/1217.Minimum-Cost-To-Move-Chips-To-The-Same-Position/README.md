# [1217. 玩筹码](https://leetcode-cn.com/problems/minimum-cost-to-move-chips-to-the-same-position/)

## 题目

数轴上放置了一些筹码，每个筹码的位置存在数组 chips 当中。

你可以对 任何筹码 执行下面两种操作之一（不限操作次数，0 次也可以）：

- 将第 i 个筹码向左或者右移动 2 个单位，代价为 0。
- 将第 i 个筹码向左或者右移动 1 个单位，代价为 1。
最开始的时候，同一位置上也可能放着两个或者更多的筹码。

返回将所有筹码移动到同一位置（任意位置）上所需要的最小代价。

 

示例 1：
```c
输入：chips = [1,2,3]
输出：1
解释：第二个筹码移动到位置三的代价是 1，第一个筹码移动到位置三的代价是 0，总代价为 1。
```
示例 2：
```c
输入：chips = [2,2,2,3,3]
输出：2
解释：第四和第五个筹码移动到位置二的代价都是 1，所以最小总代价为 2。
```

提示：

- 1 <= chips.length <= 100
- 1 <= chips[i] <= 10^9

## 解题思路

- 给出一个数组，数组的下标代表的是数轴上的坐标点，数组的元素代表的是砝码大小。砝码移动规则，左右移动 2 格，没有代价，左右移动 1 个，代价是 1 。问最终把砝码都移动到一个格子上，最小代价是多少。
- 先解读砝码移动规则：偶数位置的到偶数位置的没有代价，奇数到奇数位置的没有代价。利用这个规则，我们可以把所有的砝码无代价的摞在一个奇数的位置上和一个偶数的位置上。这样我们只用关心这两个位置了。并且这两个位置可以连续在一起。最后一步即将相邻的这两摞砝码合并到一起。由于左右移动一个代价是 1，所以最小代价的操作是移动最少砝码的那一边。奇数位置上砝码少就移动奇数位置上的，偶数位置上砝码少就移动偶数位置上的。所以这道题解法变的异常简单，遍历一次数组，找到其中有多少个奇数和偶数位置的砝码，取其中比较少的，就是最终答案。