- You are iterating over `runes` but you are using the index to get single `bytes` from the string. 
This is working fine as the exercise only requires you to handle `ascii` characters. 
You can go ahead and add a few more test cases containing non-ascii characters.
See what happens!