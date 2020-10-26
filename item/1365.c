/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

void swap(int* a,int* b){
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

int partition(int* arr,int low,int high){
    int pivot = high;
    int i = low-1;

    for (int j=low;j<high;j++){
        if (arr[j]<pivot){
            i++;
            swap(&arr[i],&arr[j]);
        }
    }
    swap(&arr[i+1],&arr[high]);
    return i+1;
}

void quicksort(int* arr,int low,int high){
    if(low<high){
        int middle = partition(arr,low,high);
        quicksort(arr,low,middle-1);
        quicksort(arr,middle+1,high);
    }
}

int* smallerNumbersThanCurrent(int* nums, int numsSize, int* returnSize){
    // 建立一个数组保存索引，以数字为索引，原索引为值
    int index[100]={0};
    for (int i=0;i<numsSize;i++){
        index[nums[i]]=i;
    }
    quicksort(nums,0,numsSize-1);
    
    int* res = (int*)malloc(sizeof(int)*numsSize);
    for (int i=0;i<numsSize;i++){
        res[index[nums[i]]]=i;
    }

    *returnSize = numsSize;
    return res;
}
