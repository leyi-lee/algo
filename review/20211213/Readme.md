### 排序
#####冒泡排序 O(n * n)  
> 前后两个比较相互置换
> i 从 0 开始 i < n - 1
> j 从 0 开始 j < n - 1 - i 排过序的就不排了后面的就放好了
> 判断条件，开始置换 

#####插入排序 O(n * n) 
> 从 1 开始循环，每次j-- 判断j 大于 0 且 j 与 j - 1 比较
> j = i, j > 0 && nums[j] < nums[j - 1]

##### 归并排序 nlogn
> 先按中位数拆开
> 然后再合并两个有序数组
> 创建临时数组 大小为 r - l + 1
> i = l; j = mid + 1, k = 0
> l <= mid && j <= r 比较两个大小
> 往temp中放入 k++,判断那个小放哪个，哪个对应的下标++
> 最后把temp数组跟新到原数组中, 原数组下标nums[left + i]

##### 快排 nlogn
> 先随机找到一个数，然后把比这个数大的找到一个，然后小的找到一个，然后把两数交换，然后继续
> 最后到l 和 r 相遇到了一个下标，依次往下走
> l > r return
> 拆分 先随机找 pivot = l + rand(0, r - l + 1)
> pivotVal = nums[pivot]
> l <= r nums[l] < privotVal l++ nums[r] > privotVal r--
> l < r 交换 l++,r-- ;相等停止 l == r break 返回相交的位置$r

  
