class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    ## This solution loops through the linked lists seperately to retrieve the number from each
    ## Once we have each number, we add them together. Then we break down the final number format
    ## end to start, using the modulo operator to get the number and then losing that number bytearray
    ## dividing by 10. From there, we just create a new linked list.bytearray
    ## The problem with this solution is you seem to encounter integer overflow with very large numbers,
    ## but it works with normal integers.
    def addTwoNumbersGood(self, l1: ListNode, l2: ListNode) -> ListNode:
        sumL1, sumL2, power = 0, 0, 1

        while True:
            # print("node val: ", l1.val)
            sumL1 = sumL1 + (l1.val * power)
            # print("sumL1: ", sumL1)
            power *= 10
            if l1.next is None:
                break
            l1 = l1.next

        power = 1
        while True:
            # print("node val: ", l2.val)
            sumL2 = sumL2 + (l2.val * power)
            # print("sumL2: ", sumL2)
            power *= 10
            if l2.next is None:
                break
            l2 = l2.next

        answer = sumL1 + sumL2
        print(sumL1, " + ", sumL2, " = ", answer)
        if answer == 0:
            return ListNode(0, None)

        answerList = []
        while answer >= 1:
            val = answer % 10
            print("answer", answer, "val", val, "int val", int(val))
            answer = int(int(answer) / 10)
            answerList.append(int(val))

        # print("answer list", answerList)

        nodeHead = None
        for num in reversed(answerList):
            node = ListNode(num, nodeHead)
            # print(node.val, "<--")
            nodeHead = node

        return nodeHead

    def addTwoNumbersBetter(self, l1: ListNode, l2: ListNode) -> ListNode:
        head = l1
        prev = head
        firstLoop = True
        remainder = 0

        while True:
            if l1 is None and l2 is None and remainder == 0:
                break

            valOne, valTwo = 0, 0
            if l1 is not None:
                valOne = l1.val
                l1 = l1.next
            if l2 is not None:
                valTwo = l2.val
                l2 = l2.next

            val = valOne + valTwo + remainder
            remainder = int(val / 10)
            if val >= 10:
                val %= 10

            temp = ListNode(val, None)
            if firstLoop is True:
                head = temp
                prev = head
                firstLoop = False
            else:
                prev.next = temp
                prev = temp

        return head

    ## This solution handles creating the next nodes in place as we loop through each linked list together. We totally avoid
    ## creating the entire numbers, and really just focus on the math of these 2 digits. The big thing here is keeping track of
    ## the remainder, using it in the addition and making sure not to forget about it at the end. The dummy node is also really
    ## nice because it allows us to just create a new node and set the next of the dummy to what our head will be
    def addTwoNumbersBetterandPrettier(self, l1: ListNode, l2: ListNode) -> ListNode:
        dummy = ListNode()
        curr = dummy
        carry = 0

        while l1 or l2 or carry:
            v1 = l1.val if l1 else 0
            v2 = l2.val if l2 else 0

            val = v1 + v2 + carry
            carry = val // 10
            val %= 10

            curr.next = ListNode(val)
            curr = curr.next

            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None

        return dummy.next

    def addTwoNumbersJuly2026(
        self, l1: Optional[ListNode], l2: Optional[ListNode]
    ) -> Optional[ListNode]:
        numA, numB = 0, 0
        base = 1

        while l1 or l2:
            if l1 is not None:
                val = l1.val * base
                numA += val
                l1 = l1.next
            if l2 is not None:
                val = l2.val * base
                numB += val
                l2 = l2.next
            base *= 10

        total = numA + numB
        head = ListNode()
        curr = head
        while total > 0:
            val = total % 10
            curr.val = val

            total = total // 10
            if total > 0:
                nxt = ListNode()
                curr.next = nxt
                curr = nxt

        return head


## Number is 4321 1 --> 2 --> 3 --> 4
node4 = ListNode(4, None)
node3 = ListNode(3, node4)
node2 = ListNode(2, node3)
node1 = ListNode(1, node2)

## Number is 8765 5 --> 6 --> 7 --> 8
node8 = ListNode(8, None)
node7 = ListNode(7, node8)
node6 = ListNode(6, node7)
node5 = ListNode(5, node6)

## Number is 1000000000000000001
node28 = ListNode(1, None)
node27 = ListNode(0, node28)
node26 = ListNode(0, node27)
node25 = ListNode(0, node26)
node24 = ListNode(0, node25)
node23 = ListNode(0, node24)
node22 = ListNode(0, node23)
node21 = ListNode(0, node22)
node20 = ListNode(0, node21)
node19 = ListNode(0, node20)
node18 = ListNode(0, node19)
node17 = ListNode(0, node18)
node16 = ListNode(0, node17)
node15 = ListNode(0, node16)
node14 = ListNode(0, node15)
node13 = ListNode(0, node14)
node12 = ListNode(0, node13)
node11 = ListNode(0, node12)
node10 = ListNode(1, node11)

## Number is 465
node32 = ListNode(4, None)
node31 = ListNode(6, node32)
node30 = ListNode(5, node31)

s = Solution()
s.addTwoNumbersGood(node10, node30)
