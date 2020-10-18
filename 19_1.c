// 删除倒数元素的链表
// 用快慢指针做，一次遍历

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* removeNthFromEnd(struct ListNode* head, int n){
    struct ListNode* dummy = (struct ListNode*)malloc(sizeof(struct ListNode));
    dummy->next = head;
    dummy->val = 0;
    struct ListNode* first = dummy;
    struct ListNode* second = dummy;

    for (int i=0;i<n;i++){
        printf("hello");
        first = first->next;
    }
    while(first->next!=NULL){
        first=first->next;
        second=second->next;
    }
    printf("%d",second->val);
    second->next=second->next->next;
    head = dummy->next;
    free(dummy);
    return head;
}
