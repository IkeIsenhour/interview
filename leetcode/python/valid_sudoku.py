class Solution:
    # This seems to be the most efficent solution. It's actually very simple to
    # solve, you just validate rows, then columns and finally boxes. You set up a series
    # of intricate nested for loops to check each one, and use a set to track what has been encountered.
    # tc: technically O(1) because 9 is constant. BUT, if we make this infinite it is O(n^2).
    # sc: technically O(1) because 9 is constant. BUT, if we make this infinite it is O(n).

    def isValidSudoku(self, board: List[List[str]]) -> bool:
        # Validate rows
        for row in range(9):
            s = set()
            for col in range(9):
                num = board[row][col]
                if num in s:
                    return False
                elif num != ".":
                    s.add(num)

        # Validate columns
        for row in range(9):
            s = set()
            for col in range(9):
                num = board[col][row]
                if num in s:
                    return False
                elif num != ".":
                    s.add(num)

        # Validate boxes
        starts = [
            (0, 0),
            (0, 3),
            (0, 6),
            (3, 0),
            (3, 3),
            (3, 6),
            (6, 0),
            (6, 3),
            (6, 6),
        ]

        for i, j in starts:
            s = set()
            for row in range(i, i + 3):
                for col in range(j, j + 3):
                    num = board[row][col]
                    if num in s:
                        return False
                    elif num != ".":
                        s.add(num)

        return True
