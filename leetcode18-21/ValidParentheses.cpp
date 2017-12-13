class Solution {
public:
    bool isValid(string s) {
        vector<char> stack(0);
        for (int i = 0; i < s.size(); i++) {
            char oc = otherChar(s[i]);
            if (oc == ' ') {
                stack.push_back(s[i]);
            } else {
                if (stack.empty() || stack.back() != oc) {
                    return false;
                } else {
                    stack.pop_back();
                }
            }
        }
        return stack.empty();
    }
    char otherChar(char c) {
        switch (c) {
        case ')':
            return '(';
        case ']':
            return '[';
        case '}':
            return '{';
        default:
            return ' ';
        }
    }
};
