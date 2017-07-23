from tex import latex2pdf
import sys

if __name__ == "__main__":
    sys.stdout.write(latex2pdf(sys.stdin.read().decode("utf-8")))
    