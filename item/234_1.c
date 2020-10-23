// 回文链表
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
// 递归实现，递归进去回溯查询，通过外面的一个指针。
struct ListNode* ptr;

bool checkPoint(struct ListNode* point){
    if (point != NULL){
        if (!checkPoint(point->next)) return false;// 这里要对回溯的结果进行判断
        if (point->val != ptr->val) return false;
        ptr = ptr->next;
    }
    return true;
}

bool isPalindrome(struct ListNode* head){
    ptr = head;
    return checkPoint(head);
}
