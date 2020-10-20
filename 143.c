// 重排链表
// 用线性表存储链表顺序

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */


void reorderList(struct ListNode* head){
    struct ListNode* SqList[40001];
    struct ListNode* tmp = head;
    int i = 0;
    int j = 0;
    while(tmp!=NULL){
        SqList[j] = tmp;
        tmp=tmp->next;
        j++;
    }

    j--;
    while (i < j) {
        SqList[i]->next = SqList[j];
        i++;
        if (i == j) {
            break;
        }
        SqList[j]->next = SqList[i];
        j--;
    }
    SqList[i]->next = NULL;
}
