# game = [[5, 3, 0, 0, 7, 0, 0, 0, 0],
#         [6, 0, 0, 1, 9, 5, 0, 0, 0],
#         [0, 9, 8, 0, 0, 0, 0, 6, 0],
#         [8, 0, 0, 0, 6, 0, 0, 0, 3],
#         [4, 0, 0, 8, 0, 3, 0, 0, 1],
#         [7, 0, 0, 0, 2, 0, 0, 0, 6],
#         [0, 6, 0, 0, 0, 0, 2, 8, 0],
#         [0, 0, 0, 4, 1, 9, 0, 0, 5],
#         [0, 0, 0, 0, 8, 0, 0, 7, 9]]
game = [[8, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 3, 6, 0, 0, 0, 0, 0],
        [0, 7, 0, 0, 9, 0, 2, 0, 0],
        [0, 5, 0, 0, 0, 7, 0, 0, 0],
        [0, 0, 0, 0, 4, 5, 7, 0, 0],
        [0, 0, 0, 1, 0, 0, 0, 3, 0],
        [0, 0, 1, 0, 0, 0, 0, 6, 8],
        [0, 0, 8, 5, 0, 0, 0, 1, 0],
        [0, 9, 0, 0, 0, 0, 4, 0, 0]]


class Play:
    def __init__(self, value, row, col):
        self.value = value
        self.row = row
        self.col = col


class Sudoku:
    def __init__(self, board):
        self.stack = []
        self.board = board

    def find_empty_cell(self):
        for i in range(9):
            for j in range(9):
                if self.board[i][j] == 0:
                    return i, j  # row, col
        return None, None

    """OPTIMISE"""
    def try_value(self, row, col):
        for val in range(self.board[row][col] + 1, 10):
            status = all([self.check_column(col, val),
                          self.check_row(row, val),
                          self.check_block(row, col, val)])
            if status:
                self.board[row][col] = val
                self.stack.append(Play(val, row, col))
                return True
        return False

    # speed up by exiting on false
    def check_column(self, col, val):
        return all(row[col] != val for row in self.board)

    # speed up by exiting on false
    def check_row(self, row, val):
        return all(col != val for col in self.board[row])

    def check_cells(self, row_l, row_h, col_l, col_h, val):
        for i in range(row_l, row_h):
            for j in range(col_l, col_h):
                if self.board[i][j] is val:
                    return False
        return True

    """OPTIMISE"""
    def check_block(self, row, col, val):
        if row < 3 and col < 3:  # 1
            return self.check_cells(0, 3, 0, 3, val)
        elif row < 3 <= col <= 5:  # 2
            return self.check_cells(0, 3, 3, 6, val)
        elif row < 3 and col > 5:  # 3
            return self.check_cells(0, 3, 6, 9, val)
        elif col < 3 <= row <= 5:  # 4
            return self.check_cells(3, 6, 0, 3, val)
        elif 3 <= row <= 5 and 3 <= col <= 5:  # 5
            return self.check_cells(3, 6, 3, 6, val)
        elif 3 <= row <= 5 < col:  # 6
            return self.check_cells(3, 6, 6, 9, val)
        elif row > 5 and col < 3:  # 7
            return self.check_cells(6, 9, 0, 3, val)
        elif 3 <= col <= 5 < row:  # 8
            return self.check_cells(6, 9, 3, 6, val)
        elif row > 5 and col > 5:  # 9
            return self.check_cells(6, 9, 6, 9, val)

    def backtrack(self):
        backtrack_status = False
        while not backtrack_status:
            val = self.stack.pop(-1)
            backtrack_status = self.try_value(val.row, val.col)
            if backtrack_status is False:
                self.board[val.row][val.col] = 0

    def solve(self):
        row, col = 0, 0

        while row is not None and col is not None:
            # find an empty cell
            row, col = self.find_empty_cell()
            # try a value in cell
            if row is not None and col is not None:
                if self.try_value(row, col):
                    continue
                else:
                    self.backtrack()
            else:
                break

    def print_board(self):
        for i in range(9):
            if i % 3 == 0 and i != 0:
                print('- - - - - - - - - - -')

            for j in range(9):
                if j % 3 == 0 and j != 0:
                    print('| ', end='')

                if j == 8:
                    print(self.board[i][j])
                else:
                    print(f'{self.board[i][j]} ', end='')
        print('')


if __name__ == "__main__":
    game = Sudoku(game)
    game.print_board()
    game.solve()
    game.print_board()
