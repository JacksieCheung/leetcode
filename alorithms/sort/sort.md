## Sort 排序 PART I

### 1. 插入排序 insertion sort / 时间复杂度 O(n2)

&#8195;插入排序就像玩扑克牌游戏。我们会从头开始排。每遍历一遍，就将该元素遍历和前面已经排好的元素进行比较，找到目标插入点，再插入。

&#8195;插入排序步骤（以下标从0开始的数组为例）(copy from geeksforgeeks)

* 1: Iterate from arr[1] to arr[n] over the array.

* 2: Compare the current element (key) to its predecessor.

* 3: If the key element is smaller than its predecessor, compare it to the elements before. Move the greater elements one position up to make space for the swapped element.

![insertsort](./picture/insertionsort.png)

&#8195;实际上很简单，每次迭代，取最末尾为关键字key。然后取一个指针指向key的前端比较。若比key大，则向后移动一位。最后把key赋值到对应位置。

* 需要注意的是插入排序移动元素不是交换，而是直接赋值。

&#8195;下面是代码：

```c
#include<iostream>

void insertsort(int* arr,int len){
	for (int j=1;j<len;j++) {
		int key = arr[j];
		int i = j-1;
		while(i>=0&&arr[i]>key) {
			arr[i+1]=arr[i];
			i--;
		}
		arr[i+1] = key;
	}
}

int main(){
	int arr[] = {5,4,1,2,5,4,8,9,7,10,4};
	int len = sizeof(arr)/sizeof(arr[0]);
	insertsort(arr,len);
	for (int i=0;i<len;i++){
		std::cout<<arr[i]<<" ";
	}
	std::cout<<std::endl;
	return 0;
}
```

* 也可看下伪代码：

```c
INSERTION-SORT(A)
  for j = 2 to A.length
    key = A[i]
    // insert A[j] into the sorted sequence A[i...j-1]
    i = j - 1
    while i > 0 and A[i] > key
      A[i+1] = A[i]
      i = i - 1
    A[i+1] = key
```

### 2. 泡沫排序 bubblesort \ 时间复杂度 O(n2)

&#8195;流行（因为简单）但低效，通过反复交换元素

&#8195;流程上，基本就是双指针，从尾巴开始遍历。每遍历一次就把最小的元素移到最前面，然后前面指针后移，长度减一。重新上面的步骤。

&#8195;代码如下：

```c
#include<iostream>

void swap(int* a,int* b){
	int tmp = *a;
	*a = *b;
	*b = tmp;
}

void bubblesort(int* arr,int n){
	for (int i=0;i<n;i++){
		for(int j=n;j>=i+1;j--){
			if (arr[j]<arr[j-1])
				swap(&arr[j],&arr[j-1]);
		}
	}
}

int main(){
	int arr[]={1,4,3,5,9,6,2,5,4,7,8};
	int n=sizeof(arr)/sizeof(arr[0]);
	
	bubblesort(arr,n-1);

	for(int i=0;i<n;i++){
		std::cout<<arr[i]<<" ";
	}
	std::cout<<std::endl;
}
```

* 伪代码如下：

```c
BUBBLESORT(A)
  for i = 1 to A.length - 1
    for j = A.length downto i + 1
	  if A[j] < A[j-1]
	    exchange A[j] with A[j-1]
```

### 3. 归并排序 merge sort  / 时间复杂度 O(nlgn)

&#8195;归并排序体现分治思想。它由两个部分组成：

* 把原数组从中间拆开，传入下层再拆开，传入下层再拆开...一直到底部

* 从底部开始回溯，一个元素的数组合并成2元素的有序数组，二元素有序数组合并成四元素有序数组...一直回溯到顶部。

&#8195;代码如下：由mergesort调用merge函数实现。

```c
#include<iostream>

void merge(int* arr,int low,int middle,int high){
	int i,j,k;
	int n1 = middle-low+1;
	int n2 = high-middle; // 这里把原有部分分成两个部分，其中middle算前半部分。
	int a[n1];
	int b[n2];
	// 对辅助数组赋值
	for (i=0;i<n1;i++) a[i]=arr[low+i];
	for (j=0;j<n2;j++) b[j]=arr[middle+1+j];

	i=0,j=0,k=low;    // 记得要对三个索引初始化
	// merge 对两个已经有序的数组采用归并排序，各走一遍
	while(i<n1&&j<n2){
		if (a[i]<=b[j]) {
			arr[k]=a[i];
			i++;
		} else {
			arr[k]=b[j];
			j++;
		}
		k++;
	}
	while(i<n1) {
		arr[k]=a[i];
		k++;
		i++;
	}
	while(j<n2){
		arr[k]=b[j];
		k++;
		j++;
	}
}

// 调用 mergesort 对数组进行拆分，拆到最后底层为一个，开始回溯
// 然后一个合并成两个，两个的片段合并成四个元素片段，以此回溯
void mergesort(int* arr,int low,int high){
	if (low<high){
		int middle=(low+high)/2;
		mergesort(arr,low,middle);
		mergesort(arr,middle+1,high);
		merge(arr,low,middle,high);
	}
}

int main(){
	int arr[]={9,8,4,5,2,1,4,5,8,7,6};
	int n = sizeof(arr)/sizeof(arr[0]);
	mergesort(arr,0,n-1);
	for (int i=0;i<n;i++){
		std::cout<<arr[i]<<" ";
	}
	std::cout<<std::endl;
	return 0;
}
```

&#8195;**merge函数是关键。**

&#8195;在merge函数中，默认调用的两个数组是有序的。我们通过传进的low，high，middle，以middle为界限拆分成两个数组。其中middle算low那边的。

&#8195;接着我们用归并排序，一次遍历，两个数组各走一遍的办法将顶层数组arr的一部分进行排序。然后返回。

&#8195;**mergesort函数助拆分**

&#8195;但是一般我们要排序的数组，两边不会是有序的。所以mergesort函数的作用就是不断调用它自身，将数组不断拆分。然后回溯上来，一层层建立，直到合并整个数组。

* 下面是它的过程：

![mergesort](./picture/Merge-Sort-Tutorial.png)



### 4. 堆排序 heapsort / 时间复杂度：O(nlgn)

&#8195;堆排序以heapify为基础。要先构建heapify函数，再实现排序

&#8195;heapify函数使每次遍历时，找到最大的元素放在堆顶

&#8195;堆排序只要将堆顶元素放到堆尾，再缩减堆的长度（使堆尾最大长度不进入循环）以此重复即可

* 代码如下：

```c
#include<iostream>

void heapify(int* arr,int n,int i){
	int largest = i;
	int l = 2*i + 1;
	int r = 2*i + 2;
	
	if (l<n&&arr[l]>arr[largest])
		largest = l;

	if (r<n&&arr[r]>arr[largest])
		largest = r;

	if (largest != i){
		// swap(arr[i],arr[largest]);
		int tmp = arr[i];
		arr[i] = arr[largest];
		arr[largest] = tmp;
		
		heapify(arr,n,largest);
	}
}

void heapsort(int* arr,int n){
	for (int i = n/2-1;i>=0;i--)
		heapify(arr,n,i);

	for (int i=n-1;i>0;i--){
		int tmp = arr[0];
		arr[0] = arr[i];
		arr[i] = tmp;

		heapify(arr,i,0);
	}
}

int main(){
	int arr[]={5,1,4,6,8,3};
	int n = sizeof(arr)/sizeof(arr[0]);

	heapsort(arr,n);

	for (int i=0;i<n;i++)
		std::cout<<arr[i]<<" ";
	std::cout<<std::endl;
	return 0;
}
```

* heap 排序流程如图

![流程](./picture/heapsort.png)

### 5. 快速排序 quicksort /  时间复杂度(一般)：O(nlgn)

&#8195;快速排序最坏的情况是 O(n2)，但是一般情况非常好，甚至比堆排序还快。

&#8195;快速排序有两个步骤：

* 拆分

* 递归

&#8195;其中最重要的，就是拆分。我们会设置一个中枢点 pivot 。然后用双指针（比较快）遍历一遍数组，比 pivot 大的放 pivot 右边，比 pivot 小的放左边。接着我们把这两部分放入递归，再按上述步骤进行拆分。

* 快速排序也可以不用递归，用迭代实现

* 实现拆分的方法有很多，可以是最后一个作为中枢，可以是第一个成为中枢，也可以是中间那个是中枢。

下面上代码：

```c
#include<iostream>

// 交换函数
void swap(int* m,int* n){
	int tmp = *m;
	*m = *n;
	*n = tmp;
}

// 默认以最后一个元素作为中枢
// 比中枢大的放右边，比中枢小的放左边。一直递归求解即可。
int partition(int* arr,int low,int high){
	int pivot = arr[high];
	int i = low - 1;

	// 采用了双指针比较，一遍循环结束
	for (int j=low;j<=high-1;j++){
		if (arr[j]<pivot) {
			i++;
			swap(&arr[j],&arr[i]);
		}
	}
	swap(&arr[i+1],&arr[high]);
	return (i+1);
}

void quicksort(int* arr,int low,int high){
	if (low<high){
		int pi = partition(arr,low,high);

		quicksort(arr,low,pi-1);  // 传入的节点是pi-1和pi+1
		quicksort(arr,pi+1,high); // 不用担心越界的问题，因为下一次递归有条件限制。
	}
}

int main(){
	int a[]={5,4,1,4,4,8,6,8,4,2,6,5};
	int n = sizeof(a)/sizeof(a[0]);
	quicksort(a,0,n-1);
	
	for (int i=0;i<n;i++){
		std::cout<<a[i]<<" ";
	}
	std::cout<<std::endl;
	return 0;
}
```

过程如图：

![quickSort](./picture/QuickSort.png)

&#8195;拆分的过程，也是非常精妙的算法。分的过程有四段，小于pivot的，大于pivot的，未拆分的和pivot。

&#8195;其中i指向小于pivot和大于pivot的边界，j指向大于pivot和未拆分的边界。每当j往后移动，访问一个未拆分的元素，判断它和pivot的大小，如果大于，则不需要操作，继续往后走（j往后走的过程已经将元素分为了大于pivot边界的范围）。如果小于，则i往后走，并且交换i和j的元素。（这个过程i指向了大于pivot的元素，用小于pivot的换大于pivot的。i扩张了，j整体向后移动）。最后i++和pivot交换。

&#8195;i每次向后走，都是直接指向了大于pivot的元素，所以直接交换不用怕出什么问题，交换之后那个元素仍然在大于pivot的范围内。

**关于快速排序的时间复杂度**(以后再看)

&#8195;为了使快速排序能划分得更均衡，我们可以用随机快速排序。就是在拆分时，随机选择主元，别的都一致。

### 总结：上述的五种排序方法，都是比较排序

&#8195;因为上述排序的输出结果，都是基于排序来完成的。任何比较排序在最坏的情况下都要经过O(nlgn)次比较。所以归并排序和快速排序已经接近最优了。其中快速排序在一般情况下接近最优。


