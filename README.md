This is 360 Assignment 1
Programming Part
prog 1
name: Fan Wu
NSID: faw615
Student NO.: 11204313

Help Vader out by finding an O(V +E) algorithm which identifies the leastcost  
articulation point which disconnects the graph when removed.  
We provide an input file 'test'N'.in' and its expected output file 'test'N'.exp'.
Input will be read from standard input, as a file comprising  
• line 1 contains an integer, V , telling us how many star-systems in the map  
• lines 2 . . . V + 1 each contain one star-system name (a string without  
whitespace) and, except for Scarif and Yavin, its blockade cost (a positive  
integer)  
• line V + 2 contains an integer E, telling us how many hyperspace jumps  
exist  
• lines V + 3 . . . V + E + 2 contains two star-system names, telling us we  
can jump directly between them  
Note that Scarif and Yavin will be in the list of star-systems, without blockade  
values.  

Emit one line to standard output, telling us either  
Leia escapes with the plans.  
or  
Darth blockades star-system name (blockage cost).  

For example, seven star-systems and seven hyperspace jumps might be povide input
>   7
>   Scarif
>   Coruscant 8
>   Mandalore 1
>   Tatooine 3
>   Naboo 4
>   Yavin
>   Kamino 2
>   7
>   Scarif Coruscant
>   Scarif Mandalore
>   Coruscant Tatooine
>   Tatooine Mandalore
>   Naboo Tatooine
>   Tatooine Yavin
>   Yavin Kamino

and would be expected to generate output
>   Darth blockades Tatooine (3).