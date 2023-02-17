import sys
import os.path as osp

FILE = sys.argv[1]

with open(FILE) as f:
    books = f.readlines()

for book in books:
    with open(f"books/{book.strip()}.pdf", "w") as f:
        f.write(f"{book}")
