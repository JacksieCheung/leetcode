## Sort 排序

### 1. 插入排序

### 2. 泡沫排序 流行但低效，通过反复交换元素

```
BUBBLESORT(A)
  for i = 1 to A.length - 1
    for j = A.length downto i + 1
	  if A[j] < A[j-1]
	    exchange A[j] with A[j-1]
```

### 3. 归并排序

### 4. 堆排序 heapsort \ 时间复杂度：O(nlgn)

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

### 5. 快速排序  \  时间复杂度(一般)：O(nlgn)

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

		quicksort(arr,low,pi-1);
		quicksort(arr,pi+1,high);
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

**关于快速排序的时间复杂度**
