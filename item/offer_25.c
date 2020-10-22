// 合并两个有序数组
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* mergeTwoLists(struct ListNode* l1, struct ListNode* l2){
    if (l1==NULL&&l2==NULL) return NULL;
    if (l1==NULL) return l2;
    if (l2==NULL) return l1;
    struct ListNode* res = (struct ListNode*)malloc(sizeof(struct ListNode));
    struct ListNode* l3 = res;
    while(l1!=NULL&&l2!=NULL){
        if (l1->val <= l2->val){
            l3->next = l1;
            l1 = l1->next;
        } else {
            l3->next = l2;
            l2 = l2->next;
        }
        l3 = l3->next;
    }
    while (l1!=NULL){
        l3->next = l1;
        l1 = l1->next;
        l3 = l3->next;
    }
    while (l2!=NULL){
        l3->next = l2;
        l2 = l2->next;
        l3 = l3->next;
    }
    return res->next;
}
