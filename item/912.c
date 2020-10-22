// 排序数组
// 堆排序法
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

void heapify(int* arr,int n,int i){ // n 为长度，i为当前节点索引
    int largest = i;
    int left = 2*i+1;
    int right = 2*i+2;

    if (left<n&&arr[left]>arr[largest]) largest = left;
    if (right<n&&arr[right]>arr[largest]) largest = right;

    if (largest != i) {
        int tmp = arr[i];
        arr[i] = arr[largest];
        arr[largest] = tmp;
        heapify(arr,n,largest);
    }
}

void heapsort(int* arr,int n){
    for (int i=n/2 - 1;i>=0;i--) heapify(arr,n,i);

    // for (int i=0;i<n;i++) printf("%d",arr[i]);

    for (int i=n-1;i>0;i--){
        int tmp = arr[0];
        arr[0] = arr[i];
        arr[i] = tmp;

        heapify(arr,i,0);
    }
}

int* sortArray(int* nums, int numsSize, int* returnSize){
    heapsort(nums,numsSize);
    *returnSize = numsSize;
    return nums;
}
