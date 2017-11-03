class Solution(object):
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        number = ''
        i = 0
        while i < len(str):
            if str[i] != ' ':
                if 47 < ord(str[i]) < 58:
                    number = str[i]
                    j = i + 1
                    while j < len(str) and 47 < ord(str[j]) < 58:
                        number = number + str[j]
                        j = j + 1
                    break
                elif str[i] == '+' or str[i] == '-':
                    j = i + 1
                    if j < len(str) and 47 < ord(str[j]) < 58:
                        number = str[i]
                        while j < len(str) and 47 < ord(str[j]) < 58:
                            number = number + str[j]
                            j = j + 1
                        break
                    else:
                        break
                else:
                    return 0
            i = i + 1
        if number == '':
            return 0
        else:
            if int(number) > 2147483647:
                return 2147483647
        if int(number) < -2147483648:
            return -2147483648
        return int(number)
