#include "list.h"

ListNode *getNode(ListNode *head, int index) {
    for (int i = 0; i < index; i++)
        head = head->next;
    return head;
}

int goToLast(ListNode *head, ListNode **outTail) {
    if (head) {
        int len = 1;
        ListNode *tail = head->next;
        for (; tail; len++) {
            head = tail;
            tail = tail->next;
        }
        *outTail = head;
        return len;
    } else {
        *outTail = NULL;
        return 0;
    }
}

ListNode *rotateRight(ListNode *head, int k) {
    ListNode *tail;
    int len = goToLast(head, &tail);
    if (len == 0)
        return head;

    k = len - k % len;

    if (k == len)
        return head;

    ListNode *newTail = getNode(head, k - 1);
    ListNode *newHead = newTail->next;
    newTail->next = NULL;
    tail->next = head;
    return newHead;
}