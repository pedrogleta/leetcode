class Solution:
    def romanToInt(self, s: str) -> int:
        result = 0
        lastCharacter = ''

        for i in range(len(s)):
            char = s[-i-1]
            if char == 'I':
                if lastCharacter in ['V', 'X']:
                    result -= 1
                else:
                    result += 1
            if char == 'V':
                result += 5
            if char == 'X':
                if lastCharacter in ['L', 'C']:
                    result -= 10
                else:
                    result += 10
            if char == 'L':
                result += 50
            if char == 'C':
                if lastCharacter in ['D', 'M']:
                    result -= 100
                else:
                    result += 100
            if char == 'D':
                result += 500
            if char == 'M':
                result += 1000

            lastCharacter = char

        return result


print(Solution().romanToInt('MCMXCIV'))
