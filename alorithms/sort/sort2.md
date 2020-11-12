## Sort 排序 PART II

### 1.计数排序

&#8195;先看计数排序，计数排序是通过保存数组中每个数字出现的次数，然后从前累加得出次序，再写入答案的一种排序方法。简单来说就是直接计算得出答案。

* 下面直接上代码：

```c
#include<iostream>
#include<cstring>
#define RANGE 255 // 宏定义一个范围做输入元素的范围

void countingSort(int* arr,int len){
	int output[len];
	
	int count[RANGE]; // 用来保存计数的数组
	memset(count,0,sizeof(count)); // 此数组初始化为0

	for (int i=0;i<len;i++) count[arr[i]]++; // 统计每个数字出现的次数

	for (int i=1;i<RANGE;i++) count[i]+=count[i-1];

	for (int i=0;i<len;i++){
		output[count[arr[i]]-1]=arr[i];
		count[arr[i]]--;
	}

	for (int i=0;i<len;i++){
		arr[i]=output[i];
	}
}

int main(){
	int arr[]={1,2,5,8,5,4,7,6,3,4};
	int len = sizeof(arr)/sizeof(arr[0]);
	countingSort(arr,len);
	
	for (int i=0;i<len;i++){ 
		std::cout<<arr[i];
	}
	std::cout<<std::endl;
	return 0;
}
```

&#8195;当输入数字是 `[0,2,2,1,3]` 时, count 数组是这样的:`[1,1,2,1]`

&#8195;当count数组完成累加后，是这样的：`[1,2,4,5]`

&#8195;当2的索引访问一次，4-1=3,这样就得到了我们要的次序，12345

&#8195;这也是为什么我们在写入答案时候，output的索引要减1。因为不存在第0位，最低是从第一位开始的。

&#8195;计数排序的时间复杂度是O(n+k)。但是要注意的是计数排序只适合用在一定区间内的数的排序。(1--k) 如果区间太大，到了(1-n2)，那可能还不如别比较排序算法。那大的区间能否用线性复杂度的算法来排序呢？答案就是下面的基数排序。

### 2.基数排序

&#8195;基数排序思想很简单，就是按位来排序。从最低位开始，以此往高位排序。其中每个位用一般用计数排序进行排序。

>Radix Sort is the answer. The idea of Radix Sort is to do digit by digit sort starting from least significant digit to most significant digit. Radix sort uses counting sort as a subroutine to sort.(摘自geeksforgeeks)

* 代码如下

```c
#include<iostream>

int getMax(int* arr,int len){ // 需要获取最大值来确定按位比较的次数
	int max = arr[0];
	for (int i=1;i<len;i++){
		if (max<arr[i]) max=arr[i];
	}
	return max;
}

void countingsort(int* arr,int n,int exp){
	int count[10]={0}; // k = 10
	int i=0;

	for (i=0;i<n;i++) count[(arr[i]/exp)%10]++;

	for (i=1;i<10;i++){
		count[i]=count[i]+count[i-1];
	}

	int output[n]={0};

	for (i=n-1;i>=0;i--){
		output[count[(arr[i]/exp)%10]-1]=arr[i];         // count数组保存的就是我们要的次序，用arr[i]访问得到
		count[(arr[i]/exp)%10]=count[(arr[i]/exp)%10]-1; // 特定的字符的次序。然后该次序减1,如果有重复的元
	}                                                    // 素，给他们留出位置。所以count保存的同一个字符最后
	                                                     // 出现的位置
	for (i=0;i<n;i++) arr[i] = output[i]; // output直接存了完整的数。然后再赋值来排序
}

void radixsort(int* arr,int len){
	int m=getMax(arr,len);
	for (int exp = 1;m/exp>0;exp*=10) countingsort(arr,len,exp);
}

int main(){
	int arr[]={9696,176,254,7412,4568,4523,84555};
	int len = sizeof(arr)/sizeof(arr[0]);
	radixsort(arr,len);
	for (int i=0;i<len;i++) std::cout<<arr[i]<<" ";
	std::cout<<std::endl;
	return 0;
}
```

&#8195;基数排序的时间复杂度是 O(d(n+b)) 。d是输入整数的位数（比较的次数），后面的括号其实是计数排序的时间复杂度。b为基数，十进制就是10。

&#8195;基数排序大多情况下超越不了快排。而且快排是原地修改，基数排序底层用了计数排序，会额外消耗空间。快排能更有效使用硬件缓存。只有我们每个位都是log2n时，基数排序才会优于快排。

### 3.桶排序

&#8195;如果我们要排序的数是浮点数，又想在用线性时间复杂度的算法，计数排序肯定是不符合我们要求的。因为计数排序的索引必须是整数。这里我们可以用第三种线性排序：桶排序。(bucket sort)。

&#8195;桶排序要做的事情不过是三步：

* 将数据均匀插入桶中

* 将每个桶内数据排序

* 连接每个桶的数据

&#8195;而这个“桶”其实在这里我们会使用链表数组来操作。用于存放数据。均匀插入，在这里我们用数据的个数`n`乘上每个数据得到的整数作为这个数组的索引。

&#8195;比如`n=10`,数据为`0.21`，则计算过程为`10*0.21=2`,索引为2。

* 如图为一次桶排序数组排序完的结果：


![bucketsort](./picture/BucketSort.png)

&#8195;每个桶内的排序一般会使用插入排序，插入排序的时间复杂度和每个桶内的元素个数是相关的。所以桶排序处理的数据最好是均匀分布的，使桶内的数据分散。一般来说，所有桶的大小的平方和与总的元素数呈线性关系，那么桶排序就仍能在线性时间内完成。时间复杂度为O(n)。

下面是代码示例：

```c
#include<algorithm>
#include<iostream>
#include<vector>

using namespace std;

void bucketSort(float arr[],int n){
	vector<float> b[n]; // 是用了cpp的容器，其实这里就是一个动态的二维数组

	for (int i=0;i<n;i++){
		int bi = n*arr[i];
		b[bi].push_back(arr[i]);// 插入元素到末尾
	}

	for (int i=0;i<n;i++) sort(b[i].begin(),b[i].end()); // 排序

	int index = 0;
	for (int i=0;i<n;i++){
		for (int j=0;j<b[i].size();j++){
			arr[index++] = b[i][j]; // 连接
		}
	}
}

int main(){
	float arr[] = {0.897,0.565,0.656,0.1234,0.665,0.3434};
	int n = sizeof(arr)/sizeof(arr[0]);
	bucketSort(arr,n);

	cout<<"Sorted arry is \n";
	for (int i=0;i<n;i++) cout<<arr[i]<<" ";
	return 0;
}

// 摘自(geeksforgeeks)
```

