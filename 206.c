// 反转链表
// 用栈（递归） 还能用迭代
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

struct ListNode* func(struct ListNode* node){
    if (node->next == NULL) {
        node;
        return node;
    }
    struct ListNode* res = func(node->next);
    node->next->next = node;
    node->next = NULL;
    return res;
}


struct ListNode* reverseList(struct ListNode* head){
    if (head==NULL) return NULL;
    struct ListNode* res = NULL;
    res = func(head);
    return res;
}
