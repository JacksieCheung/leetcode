// 回文链表
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

// 复制到数组双指针
bool isPalindrome(struct ListNode* head){
    struct ListNode* ptr = head;
    int count = 0;
    while(ptr!=NULL) {
        ptr = ptr->next;
        count++;
    }

    if (count<=0){
        return true;
    }

    int a[count];
    ptr = head;
    for (int i = 0;i<count;i++){
        a[i] = ptr->val;
        ptr = ptr->next;
    }

    int i = 0;
    int j = count-1;
    while(i<j){
        if (a[i]!=a[j]) return false;
        i++;j--;
    }
    return true;
}
