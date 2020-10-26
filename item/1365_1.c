// 请你统计数组中比它小的所有数字的数目
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* smallerNumbersThanCurrent(int* nums, int numsSize, int* returnSize){
    int* res = (int*)malloc(sizeof(int)*numsSize);
    int count = 0;

    for (int i=0;i<numsSize;i++){
        count = 0;
        for (int j=0;j<numsSize;j++){
            if (nums[i]>nums[j]) count++;
        }
        res[i]=count;
    }

    *returnSize = numsSize;
    return res;
}
