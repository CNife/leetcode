#include "list.h"

#include <stdlib.h>
#include <stdio.h>

ListNode *newNode(int val) {
    ListNode *node = malloc(sizeof(ListNode));
    if (node) {
        node->val = val;
        node->next = NULL;
    }
    return node;
}

ListNode *newList(int *values, size_t len) {
    ListNode *head = NULL, *prev = NULL;
    for (size_t i = 0; i < len; i++) {
        ListNode *next = newNode(values[i]);
        if (next) {
            if (prev) {
                prev->next = next;
                prev = next;
            } else {
                head = prev = next;
            }
        } else {
            return NULL;
        }
    }
    return head;
}

void printList(ListNode *head) {
    printf("[");
    if (head) {
        printf("%d", head->val);
    } else {
        printf("]\n");
        return;
    }

    for (ListNode *node = head->next; node; node = node->next) {
        printf(", %d", node->val);
    }
    printf("]\n");
}