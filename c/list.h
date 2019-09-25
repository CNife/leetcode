#ifndef LEETCODE_C_LIST_H
#define LEETCODE_C_LIST_H

#include <stddef.h>

typedef struct ListNode ListNode;

struct ListNode {
    int val;
    ListNode *next;
};

ListNode *newNode(int val);

ListNode *newList(int *values, size_t len);

void printList(ListNode *head);

#endif //LEETCODE_C_LIST_H
