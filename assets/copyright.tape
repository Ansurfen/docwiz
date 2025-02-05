# Copyright 2025 The DocWiz Authors. All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.
Output copyright.gif

Set TypingSpeed 75ms
Set FontSize 22
Set Width 1300
Set Height 650

Type "echo 'This is a text' > test.txt"
Enter
Type "./docwiz copyright -p 'test.txt' -c '// © 2025 DocWiz\n'"
Enter
Type "head test.txt"
Enter
Sleep 1s
Type "./docwiz copyright -p 'test.txt' -c '// © 2025 DocWiz\n'"
Enter
Type "head test.txt"
Enter
Sleep 1s
Type "./docwiz copyright -p 'test.txt' -c '// © 2025 DocWiz\n' -r"
Enter
Type "head test.txt"
Enter
Sleep 1s
Type "clear"
Enter
Type "./docwiz copyright -p 'test.txt' -c '// © 2025 DocWiz\n' -t"
Enter
Type "head test.txt"
Enter
Sleep 2s