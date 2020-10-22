// 快速排序法解决数组排序
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

void swap(int* a,int* b){
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

int partition(int* arr,int low,int high){
    int pivot = arr[high];
    int i = low - 1;

    for (int j = low;j <= high - 1;j++){
        if (arr[j]<pivot){
            i++;
            swap(&arr[i],&arr[j]);
        }
    }
    swap(&arr[i+1],&arr[high]);
    return i+1;
}

void quicksort(int* arr,int low,int high){
    if (low<high){
        int pi = partition(arr,low,high);
        quicksort(arr,low,pi-1);
        quicksort(arr,pi+1,high);
    }
}

int* sortArray(int* nums, int numsSize, int* returnSize){
    quicksort(nums,0,numsSize-1);
    *returnSize = numsSize;
    return nums;
}
