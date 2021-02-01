def longestUniqueSubsttr(string):
    seen = {}
    maximum_length = 0
    start = 0

    for end in range(len(string)):

        # Checking if we have already seen the element or not
        if string[end] in seen:
            # If we have seen the number, move the start pointer
            # to position after the last occurrence
            start = max(start, seen[string[end]] + 1)

        # Updating the last seen value of the character
        seen[string[end]] = end
        maximum_length = max(maximum_length, end-start + 1)
    return maximum_length
