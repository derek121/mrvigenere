# Mr. Vigenère
Implementation of the Vigenère cipher in Go.

## About

As described at [Wikipedia](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher):

> The Vigenère cipher is a method of encrypting alphabetic text by using a series of interwoven Caesar ciphers, based on the letters of a keyword. It employs a form of polyalphabetic substitution

## Example

```
$ go run main.go --usage
2019/10/02 09:15:10 usage args: {{--encode|--decode} key message} | --printtable | --usage
$
$ go run main.go --encode LEMON ATTACKATDAWN
LXFOPVEFRNHR
$ go run main.go --decode LEMON FHHFQYFHRFKB
ATTACKATDAWN
$
$ go run main.go --printtable
 |A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 
-----------------------------------------------------
A|A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 
B|B C D E F G H I J K L M N O P Q R S T U V W X Y Z A 
C|C D E F G H I J K L M N O P Q R S T U V W X Y Z A B 
D|D E F G H I J K L M N O P Q R S T U V W X Y Z A B C 
E|E F G H I J K L M N O P Q R S T U V W X Y Z A B C D 
F|F G H I J K L M N O P Q R S T U V W X Y Z A B C D E 
G|G H I J K L M N O P Q R S T U V W X Y Z A B C D E F 
H|H I J K L M N O P Q R S T U V W X Y Z A B C D E F G 
I|I J K L M N O P Q R S T U V W X Y Z A B C D E F G H 
J|J K L M N O P Q R S T U V W X Y Z A B C D E F G H I 
K|K L M N O P Q R S T U V W X Y Z A B C D E F G H I J 
L|L M N O P Q R S T U V W X Y Z A B C D E F G H I J K 
M|M N O P Q R S T U V W X Y Z A B C D E F G H I J K L 
N|N O P Q R S T U V W X Y Z A B C D E F G H I J K L M 
O|O P Q R S T U V W X Y Z A B C D E F G H I J K L M N 
P|P Q R S T U V W X Y Z A B C D E F G H I J K L M N O 
Q|Q R S T U V W X Y Z A B C D E F G H I J K L M N O P 
R|R S T U V W X Y Z A B C D E F G H I J K L M N O P Q 
S|S T U V W X Y Z A B C D E F G H I J K L M N O P Q R 
T|T U V W X Y Z A B C D E F G H I J K L M N O P Q R S 
U|U V W X Y Z A B C D E F G H I J K L M N O P Q R S T 
V|V W X Y Z A B C D E F G H I J K L M N O P Q R S T U 
W|W X Y Z A B C D E F G H I J K L M N O P Q R S T U V 
X|X Y Z A B C D E F G H I J K L M N O P Q R S T U V W 
Y|Y Z A B C D E F G H I J K L M N O P Q R S T U V W X 
Z|Z A B C D E F G H I J K L M N O P Q R S T U V W X Y 
```

## Notes

Why "Mr. Vigenère"? A former colleague had the practice of naming internal tools he wrote "Mr. Such-and-such", so it's an homage to him.
