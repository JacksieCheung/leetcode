/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


struct ListNode* removeNthFromEnd(struct ListNode* head, int n){
    struct ListNode* tmp = head;
    int count = 0;
    while(tmp!=NULL) {
        count++;
        tmp = tmp->next;
    }
    if (count <= 1) return NULL;
    int index = count - n;
    if (index == 0) return head->next;
    struct ListNode* ll = head;
    for (int i=0;i<index-1;i++){
        ll = ll->next;
    }
    // tmp = ll->next;
    ll ->next= ll->next->next;
    //free(tmp);
    return head;
}
