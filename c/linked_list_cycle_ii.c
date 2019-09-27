#include "list.h"

#include <assert.h>

//#include <stdio.h>
//#include <stdbool.h>
//#include <uthash.h>
//
//typedef struct {
//    ListNode *node;
//    UT_hash_handle hh;
//} HashNode;
//
//void insert(HashNode **hashTable, ListNode *node) {
//    HashNode *hashNode = malloc(sizeof(HashNode));
//    if (hashNode) {
//        hashNode->node = node;
//        HASH_ADD_PTR(*hashTable, node, hashNode); // NOLINT
//    } else {
//        perror("failed to insert node");
//        exit(1);
//    }
//}
//
//bool exists(HashNode *hashTable, ListNode *node) {
//    HashNode *res;
//    HASH_FIND_PTR(hashTable, &node, res); // NOLINT
//    return res != NULL;
//}
//
//ListNode *detectCycle(ListNode *head) {
//    HashNode *hashTable = NULL;
//    ListNode *ptr = head;
//    while (ptr) {
//        if (exists(hashTable, ptr))
//            return ptr;
//        insert(&hashTable, ptr);
//        ptr = ptr->next;
//    }
//    return NULL;
//}

ListNode *findCircle(ListNode *head) {
    ListNode *fast = head, *slow = head;
    while (fast) {
        if (fast->next)
            fast = fast->next->next;
        else
            return NULL;
        slow = slow->next;
        if (fast == slow)
            return slow;
    }
    return NULL;
}

ListNode *detectCycle(ListNode *head) {
    ListNode *p1 = head, *p2 = findCircle(head);
    if (!p2)
        return NULL;
    while (p1 != p2) {
        p1 = p1->next;
        p2 = p2->next;
    }
    return p1;
}

int main(int argc, char *argv[]) {
    int values[] = {3, 2, 0, -4};
    ListNode *list = newList((int *) values, sizeof(values) / sizeof(int));
    list->next->next->next->next = list->next;
    assert(detectCycle(list) == list->next);
    return 0;
}