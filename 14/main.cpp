#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution 
{
    public:
        string longestCommonPrefix(vector<string>& strs) 
        {
            // Check if strs has only one word
            if (strs.size() == 1)
            {
                return strs[0];
            }

            string commonPrefix = "";
            int smallestWordSize = getSmallestWordSize(strs);

            for (int charIndex = 0; charIndex < smallestWordSize; ++charIndex)
            {
                char charBeingChecked = strs[0][charIndex];
                for (int wordIndex = 0; wordIndex < strs.size(); ++wordIndex)
                {
                    if (charBeingChecked != strs[wordIndex][charIndex])
                    {
                        return commonPrefix;
                    }
                }
                commonPrefix.push_back(charBeingChecked);
            }
            return commonPrefix;
        }

        int getSmallestWordSize(vector<string> strs) {
            int smallestWordSize = strs[0].size();

            for (int i = 1; i < strs.size(); ++i)
            {
                int currentWordSize = strs[i].size();
                if (currentWordSize < smallestWordSize)
                {
                    smallestWordSize = currentWordSize;
                }
            }

            return smallestWordSize;
        }
};

int main()
{
    Solution solution;

    vector<string> strs;

    strs.push_back("ab");
    strs.push_back("a");

    cout << solution.longestCommonPrefix(strs);
}