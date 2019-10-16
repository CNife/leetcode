#include "list.h"

ListNode *middleNode(ListNode *head) {
    ListNode *fast = head, *slow = head;
    while (fast && fast->next && fast->next->next) {
        fast = fast->next->next;
        slow = slow->next;
    }
    if (fast->next) {
        return slow->next;
    } else {
        return slow;
    }
}