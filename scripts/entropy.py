# MIT License
#
# Copyright (c) 2024 buuhuu
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

import math
import sys

def calculate_entropy(filename):
    try:
        with open(filename, "rb") as file:
            counters = {byte: 0 for byte in range(2 ** 8)}
            for byte in file.read():
                counters[byte] += 1

            filesize = file.tell()
            probabilities = [counter / filesize for counter in counters.values() if counter > 0]
            entropy = -sum(math.log2(probability)*probability for probability in probabilities)
            return entropy

    except FileNotFoundError:
        print("Error: File not found -", filename)
    except IOError as e:
        print("Error: Unable to read file -", e)
    except ZeroDivisionError:
        print("Error: File is empty")
    except Exception as e:
        print("An unexpected error occurred:", e)
    return None

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python entropy.py <filename>")
        sys.exit(1)

    filename = sys.argv[1]
    entropy = calculate_entropy(filename)
    if entropy is not None:
        print("Entropy:", entropy)
