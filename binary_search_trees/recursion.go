/* en rekursiv funktion är en funktion som anropar sig självt
Innehåller alltid två delar:
- Bas-fall 
- Rekursivt fall 
*/

func Factorial(n int) int {
    if n <= 1 {
        return 1          // bas-fall: stoppa här, utan den blir det en oändlig rekursion (stack-overflow)
    }
    return n * Factorial(n-1)   // rekursivt fall: n gånger fakulteten av (n-1)
}

