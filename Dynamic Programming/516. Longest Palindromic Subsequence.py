def longestPalindromeLDP(s):
    dp = [0 for i in range(len(s))]
    for end in range(len(s)):
        dp[end], mL = 1, 0
        for start in reversed(range(0, end)):
            tmp = dp[start]
            if s[start] == s[end]:
                dp[start] = mL + 2
            mL = max(tmp, mL)
    for i in range(len(s)):
        mL = max(mL, dp[i])
    return mL
    