
/*
 What is Constants and Variables

 - associate a name with a value of a particular type
 - The value of a constant can’t be changed once it’s set, whereas a variable can be set to a different value in the future
 
 */


// declare a new constant
let maximumNumberOfLoginAttempts = 10

// declare a new variable
var currentLoginAttempt = 0
var x = 0.0, y = 0.0, z = 0.0
let (t1, t2) = (1, 2)

/*
 Type Annotations

 - Write a type annotation by placing a colon after the constant or variable name, followed by a space, followed by the name of the type to use

 */
var welcomeMessage: String
var red, green, blue: Double


/*
 Naming Constants and Variables

 - names can’t contain whitespace characters, mathematical symbols, arrows, private-use Unicode scalar values, or line- and box-drawing characters, Nor can they begin with a number

 */
let π = 3.14159
let 你好 = "你好世界"
let 🐶🐮 = "dog&cow"


/*
 Printing Constants and Variables
 - print the current value of a constant or variable with the print(_:separator:terminator:) function
 - The separator and terminator parameter have default values, so you can omit them when you call this function
 */
var friendlyWelcome: String
friendlyWelcome = "Welcome!"
print(friendlyWelcome)
print(π)
print(你好, terminator: "\n")

print("The current value of friendlyWelcome is \(🐶🐮)")
print("The value of π is \(π)")

/*
 Semicolons
 - semicolons are required if you want to write multiple separate statements on a single line
 */

let cat = "🐱"; print(cat)

/*
 Integers
 - Integers are either signed (positive, zero, or negative) or unsigned (positive or zero)
 - Swift provides signed and unsigned integers in 8, 16, 32, and 64 bit forms
    
 Integer Bounds
 - You can access the minimum and maximum values of each integer type with its min and max properties
 
 Int
 - In most cases, you don’t need to pick a specific size of integer to use, Swift provides an additional integer type: Int
 - On a 32-bit platform, Int is the same size as Int32.
 - On a 64-bit platform, Int is the same size as Int64.
 - Int can store any value between -2,147,483,648 and 2,147,483,647
 
 UInt
 - an unsigned integer type, UInt, which has the same size as the current platform’s native word size
 - On a 32-bit platform, UInt is the same size as UInt32.
 - On a 64-bit platform, UInt is the same size as UInt64.
 */

let minValue = UInt8.min  // minValue is equal to 0, and is of type UInt8
let maxValue = UInt8.max  // maxValue is equal to 255, and is of type UInt8
print("minValue: ", minValue)
print("maxValue: ", maxValue)

/*
 Floating-Point Numbers
 - Double represents a 64-bit floating-point number.
 - Float represents a 32-bit floating-point number.
 */

var a: Double
var b: Float

a = 3.1415826
b = 0.1234567
print("doubleValue: ", a)
print("floatValue: ", b)

/*
 Type Safety and Type Inference
 - Type-checking helps you avoid errors when you’re working with different types of values
 - Swift performs type checks when compiling your code and flags any mismatched types as errors
 - If you don’t specify the type of value you need, Swift uses type inference to work out the appropriate type
 */

let meaningOfLife = 42 // meaningOfLife is inferred to be of type Int
let pi = 3.14159 // pi is inferred to be of type Double
let anotherPi = 3 + 0.14159 // anotherPi is also inferred to be of type Double

/*
 Numeric Literals
 - A decimal number, with no prefix
 - A binary number, with a 0b prefix
 - An octal number, with a 0o prefix
 - A hexadecimal number, with a 0x prefix
 */

let decimalInteger = 17
let binaryInteger = 0b10001       // 17 in binary notation
let octalInteger = 0o21           // 17 in octal notation
let hexadecimalInteger = 0x11     // 17 in hexadecimal notation

let decimalDouble = 12.1875
let exponentDouble = 1.21875e1
let hexadecimalDouble = 0xC.3p0

// Both integers and floats can be padded with extra zeros and can contain underscores to help with readability
let paddedDouble = 000123.456
let oneMillion = 1_000_000
let justOverOneMillion = 1_000_000.000_000_1

/*
 Numeric Type Conversion
 - Because each numeric type can store a different range of values, you must opt in to numeric type conversion on a case-by-case basis
 
 Integer and Floating-Point Conversion
 */

//let cannotBeNegative: UInt8 = -1 // UInt8 cannot store negative numbers, and so this will report an error
//let tooBig: Int8 = Int8.max + 1 // Int8 cannot store a number larger than its maximum value, and so this will also report an error

let twoThousand: UInt16 = 2_000
let one: UInt8 = 1
let twoThousandAndOne = twoThousand + UInt16(one)

let three = 3
let pointOneFourOneFiveNine = 0.14159
let Pi = Double(three) + pointOneFourOneFiveNine
// pi equals 3.14159, and is inferred to be of type Double
let integerPi = Int(pi)
print("integerPi:", integerPi)

/*
 Type Aliases
 - define an alternative name for an existing type. You define type aliases with the typealias keyword
 */

typealias AudioSample = UInt16
var maxAmplitudeFound = AudioSample.min // maxAmplitudeFound is now 0

/*
 Booleans
 - true and false
 */
let orangesAreOrange = true
let turnipsAreDelicious = false

if turnipsAreDelicious {
    print("Mmm, tasty turnips!")
} else {
    print("Eww, turnips are horrible.")
}

// cannot convert value of type 'Int' to expected condition type 'Bool'
let i = 1
//if i {
//    // this example will not compile, and will report an error
//}

if i == 1 {
    // this example will compile successfully
}


/*
 Tuples
 - Tuples group multiple values into a single compound value. The values within a tuple can be of any type and don’t have to be of the same type as each other
 */

// http404Error is of type (Int, String), and equals (404, "Not Found")
let http404Error = (404, "Not Found")
let (statusCode, statusMessage) = http404Error
print("The status code is \(statusCode)")
print("The status message is \(statusMessage)")

let (justTheStatusCode, _) = http404Error
print("The status code is \(justTheStatusCode)")

print("The status code is \(http404Error.0)")
print("The status message is \(http404Error.1)")

let http200Status = (statusCode: 200, description: "OK")
print("The status code is \(http200Status.statusCode)")
print("The status message is \(http200Status.description)")

print((1, "zebra") < (2, "apple"))   // true because 1 is less than 2; "zebra" and "apple" are not compared
print((3, "apple") < (3, "bird"))    // true because 3 is equal to 3, and "apple" is less than "bird"
print((4, "dog") == (4, "dog"))      // true because 4 is equal to 4, and "dog" is equal to "dog"
// print(("blue", false) < ("purple", true))  // Error because < can't compare Boolean values

/*
 Optionals
 - use optionals in situations where a value may be absent
 - use an if statement to find out whether an optional contains a value by comparing the optional against nil
 - Once you’re sure that the optional does contain a value, you can access its underlying value by adding an exclamation point (!) to the end of the optional’s name
 
 Optional Binding
 - You use optional binding to find out whether an optional contains a value, and if so, to make that value available as a temporary constant or variable, Optional binding can be used with if and while statements to check for a value inside an optional
 
 Implicitly Unwrapped Optionals
 - Sometimes it’s clear from a program’s structure that an optional will always have a value
 
 nil
 - You set an optional variable to a valueless state by assigning it the special value nil
 - If you define an optional variable without providing a default value, the variable is automatically set to nil for you
 */

let possibleNumber = "123"
let convertedNumber = Int(possibleNumber)
// convertedNumber is inferred to be of type "Int?", or "optional Int"

var serverResponseCode: Int? = 404
// serverResponseCode contains an actual Int value of 404
serverResponseCode = nil
// serverResponseCode now contains no value

var surveyAnswer: String?
// surveyAnswer is automatically set to nil


/*
 If Statements and Forced Unwrapping
 */

if convertedNumber != nil {
    print("convertedNumber contains some integer value.")
    print("convertedNumber has an integer value of \(convertedNumber!).")
}
// Prints "convertedNumber contains some integer value."


// Optional binding can be used with if and while statements to check for a value inside an optional
if let actualNumber = Int(possibleNumber) {
    print("The string \"\(possibleNumber)\" has an integer value of \(actualNumber)")
} else {
    print("The string \"\(possibleNumber)\" could not be converted to an integer")
}

if let firstNumber = Int("4"), let secondNumber = Int("42"), firstNumber < secondNumber && secondNumber < 100 {
    print("\(firstNumber) < \(secondNumber) < 100")
}

if let firstNumber = Int("4") {
    if let secondNumber = Int("42") {
        if firstNumber < secondNumber && secondNumber < 100 {
            print("\(firstNumber) < \(secondNumber) < 100")
        }
    }
}

let possibleString: String? = "An optional string."
let forcedString: String = possibleString! // requires an exclamation point

let assumedString: String! = "An implicitly unwrapped optional string."
let implicitString: String = assumedString // no need for an exclamation point

// If an implicitly unwrapped optional is nil and you try to access its wrapped value, you’ll trigger a runtime error
let optionalString = assumedString
if assumedString != nil {
    print(assumedString!)
}

if let definiteString = assumedString {
    print(definiteString)
}



/*
 Error Handling
 */

enum SandwichError: Error {
    case outOfCleanDishes
    case missingIngredients(ingredients: String)
    case outOfStock
}

func makeASandwich() throws {
    // ...
}

func eatASandwich() {}
func washDishes() {}
func buyGroceries(ingredients: String) {}

let ingredients = "Coke"
do {
    try makeASandwich()
    eatASandwich()
} catch SandwichError.outOfCleanDishes {
    washDishes()
} catch SandwichError.missingIngredients(ingredients) {
    buyGroceries(ingredients: ingredients)
}

/*
 Assertions and Preconditions
 
 Debugging with Assertions
 - assert(_:_:file:line:)
 - assertionFailure(_:file:line:)
 
 Enforcing Preconditions
 - precondition(_:_:file:line:)
 - preconditionFailure(_:file:line:)
 */

var age = 3
assert(age >= 0, "A person's age can't be less than zero.")

if age > 10 {
    print("You can ride the roller-coaster or the ferris wheel.")
} else if age >= 0 {
    print("You can ride the ferris wheel.")
} else {
    assertionFailure("A person's age can't be less than zero.")
}


precondition(age > 0, "Age must be greater than zero.")

if age > 0 {
    //
} else {
    preconditionFailure("A person's age can't be less than zero.")
}


/*
 Arithmetic Operators
 - +, -, *, /, %
 - +=, -=, *=, /=
 - ==, >, < , >=, <=, !=
 */


/*
 Unary Minus Operator
*/
let four = 4
let minusFour = -four       // minusThree equals -4
let plusFour = -minusFour   // plusThree equals 4, or "minus minus four"


/*
 Unary Plus Operator
*/
let minusSix = -6
let alsoMinusSix = +minusSix  // alsoMinusSix equals -6

/*
 Ternary Conditional Operator
 - question ? answer1 : answer2
 */
let contentHeight = 40
let hasHeader = true
let rowHeight = contentHeight + (hasHeader ? 50 : 20) // rowHeight is equal to 90

/*
 Nil-Coalescing Operator
 - The nil-coalescing operator (a ?? b) unwraps an optional a if it contains a value, or returns a default value b if a is nil
 - a != nil ? a! : b
 */

let defaultColorName = "red"
var userDefinedColorName: String?   // defaults to nil

var colorNameToUse = userDefinedColorName ?? defaultColorName

userDefinedColorName = "green"
colorNameToUse = userDefinedColorName ?? defaultColorName


/*
 Range Operators
 
 Closed Range Operator
 - The closed range operator (a...b) defines [a, b]
 
 Half-Open Range Operator
 - (a..<b) defines [a, b)
 
 One-Sided Ranges
 -
 */
for index in 1...5 {
    print("\(index) times 5 is \(index * 5)")
}

for index in 1..<5 {
    print("\(index) times 5 is \(index * 5)")
}

let names = ["Anna", "Alex", "Brian", "Jack"]
for name in names[...2] {
    print(name)
}

for name in names[2...] {
    print(name)
}

for name in names[..<2] {
    print(name)
}

let range = ...5
print(range.contains(7))   // false
print(range.contains(4))   // true
print(range.contains(-1))  // true


/*
 Logical Operators
 - Logical NOT (!a)
 - Logical AND (a && b)
 - Logical OR (a || b)
 */

let allowedEntry = false
if !allowedEntry {
    print("ACCESS DENIED")
}

let enteredDoorCode = true
let passedRetinaScan = false
if enteredDoorCode && passedRetinaScan {
    print("Welcome!")
} else {
    print("ACCESS DENIED")
}

let hasDoorKey = false
let knowsOverridePassword = true
if hasDoorKey || knowsOverridePassword {
    print("Welcome!")
} else {
    print("ACCESS DENIED")
}

if enteredDoorCode && passedRetinaScan || hasDoorKey || knowsOverridePassword {
    print("Welcome!")
} else {
    print("ACCESS DENIED")
}


// It’s sometimes useful to include parentheses to make the intention of a complex expression easier to read
if (enteredDoorCode && passedRetinaScan) || hasDoorKey || knowsOverridePassword {
    print("Welcome!")
} else {
    print("ACCESS DENIED")
}

/*
 Multiline String Literals
 */

let quotation = """
The White Rabbit put on his spectacles.  "Where shall I begin,
please your Majesty?" he asked.

"Begin at the beginning," the King said gravely, "and go on
till you come to the end; then stop."
"""

// The string begins on the first line after the opening quotation marks (""") and ends on the line before the closing quotation marks, which means that neither of the strings below start or end with a line break
let singleLineString = "These are the same."
let multilineString = """
These are the same.
"""

// when you don’t want the line breaks to be part of the string’s value, write a backslash (\) at the end of those lines
let softWrappedQuotation = """
The White Rabbit put on his spectacles.  "Where shall I begin, \
please your Majesty?" he asked.

"Begin at the beginning," the King said gravely, "and go on \
till you come to the end; then stop."
"""

/*
 Special Characters in String Literals
 - \0 (null character), \\ (backslash), \t (horizontal tab), \n (line feed), \r (carriage return), \" (double quotation mark) and \' (single quotation mark)
 - An arbitrary Unicode scalar value, written as \u{n}, where n is a 1–8 digit hexadecimal number
 */
let wiseWords = "\"Imagination is more important than knowledge\" - Einstein"
// "Imagination is more important than knowledge" - Einstein
let dollarSign = "\u{24}"        // $,  Unicode scalar U+0024
let blackHeart = "\u{2665}"      // ♥,  Unicode scalar U+2665
let sparklingHeart = "\u{1F496}" // 💖, Unicode scalar U+1F496

/*
 Extended String Delimiters
 - place your string within quotation marks (") and surround that with number signs (#) to include special characters in a string without invoking their effect
 */

let threeMoreDoubleQuotationMarks = #"""
Here are three more double quotes: """
"""#

/*
 Initializing an Empty String
 
 String Mutability
 */

var emptyString = ""               // empty string literal
var anotherEmptyString = String()  // initializer syntax
// these two strings are both empty, and are equivalent to each other

var variableString = "Horse"
variableString += " and carriage" // variableString is now "Horse and carriage"

/*
 Working with Characters
 */

for character in "Dog!🐶" {
    print(character)
}

let exclamationMark: Character = "!"
let catCharacters: [Character] = ["C", "a", "t", "!", "🐱"]
let catString = String(catCharacters)
print(catString)


/*
 Concatenating Strings and Characters
 */
let string1 = "hello"
let string2 = " there"
var welcome = string1 + string2 // welcome now equals "hello there"
let mark: Character = "!"
welcome.append(mark)
print(welcome)


/*
 String Interpolation
 */
let multiplier = 3
let message = "\(multiplier) times 2.5 is \(Double(multiplier) * 2.5)" // message is "3 times 2.5 is 7.5"
print(#"Write an interpolated string in Swift using \(multiplier)."#)
print(#"6 times 7 is \#(6 * 7)."#)

/*
 Unicode
 */
let eAcute: Character = "\u{E9}"                         // é
let combinedEAcute: Character = "\u{65}\u{301}"          // e followed by
// eAcute is é, combinedEAcute is é

/*
 Counting Characters
 
 - Character values may not always affect a string’s character count
*/

let unusualMenagerie = "Koala 🐨, Snail 🐌, Penguin 🐧, Dromedary 🐪"
print("unusualMenagerie has \(unusualMenagerie.count) characters") // Prints "unusualMenagerie has 40 characters"
var word = "cafe"
word += "\u{301}"
print("the number of characters in \(word) is \(word.count)")

/*
 String Indices
 */
let greeting = "Guten Tag!"
print(greeting[greeting.startIndex]) // G
print(greeting[greeting.index(before: greeting.endIndex)]) // !
print(greeting[greeting.index(after: greeting.startIndex)]) // u
let index = greeting.index(greeting.startIndex, offsetBy: 7)
print(greeting[index]) // a

/*
 Inserting and Removing
 */
welcome = "hello"
welcome.insert("!", at: welcome.endIndex) // welcome now equals "hello!"

welcome.insert(contentsOf: " there", at: welcome.index(before: welcome.endIndex))
// welcome now equals "hello there!"
welcome.remove(at: welcome.index(before: welcome.endIndex))
// welcome now equals "hello there"

let r = welcome.index(welcome.endIndex, offsetBy: -6)..<welcome.endIndex
welcome.removeSubrange(r)

/*
 Substrings
 */

let h = "Hello, world!"
let beginning = h[..<(h.firstIndex(of: ",") ?? h.endIndex)]
// beginning is "Hello"

// Convert the result to a String for long-term storage.
let newString = String(beginning)

/*
 Comparing Strings
 - String and Character Equality
 - Prefix and Suffix Equality
 */

let str1 = "We're a lot alike, you and I."
let str2 = "We're a lot alike, you and I."
if str1 == str2 {
    print("These two strings are considered equal")
}
// Prints "These two strings are considered equal"

let eAcuteQuestion = "Voulez-vous un caf\u{E9}?"

// "Voulez-vous un café?" using LATIN SMALL LETTER E and COMBINING ACUTE ACCENT
let combinedEAcuteQuestion = "Voulez-vous un caf\u{65}\u{301}?"

if eAcuteQuestion == combinedEAcuteQuestion {
    print("These two strings are considered equal")
}
// Prints "These two strings are considered equal"

let romeoAndJuliet = [
    "Act 1 Scene 1: Verona, A public place",
    "Act 1 Scene 2: Capulet's mansion",
    "Act 1 Scene 3: A room in Capulet's mansion",
    "Act 1 Scene 4: A street outside Capulet's mansion",
    "Act 1 Scene 5: The Great Hall in Capulet's mansion",
    "Act 2 Scene 1: Outside Capulet's mansion",
    "Act 2 Scene 2: Capulet's orchard",
    "Act 2 Scene 3: Outside Friar Lawrence's cell",
    "Act 2 Scene 4: A street in Verona",
    "Act 2 Scene 5: Capulet's mansion",
    "Act 2 Scene 6: Friar Lawrence's cell"
]
var act1SceneCount = 0
for scene in romeoAndJuliet {
    if scene.hasPrefix("Act 1 ") {
        act1SceneCount += 1
    }
}
print("There are \(act1SceneCount) scenes in Act 1")

var mansionCount = 0
var cellCount = 0
for scene in romeoAndJuliet {
    if scene.hasSuffix("Capulet's mansion") {
        mansionCount += 1
    } else if scene.hasSuffix("Friar Lawrence's cell") {
        cellCount += 1
    }
}
print("\(mansionCount) mansion scenes; \(cellCount) cell scenes")

/*
 Unicode Representations of Strings
 - UTF-8 Representation
 - UTF-16 Representation
 - Unicode Scalar Representation
 */

let dogString = "Dog‼🐶"
for codeUnit in dogString.utf8 {
    print("\(codeUnit) ", terminator: "")
}
print("") // Prints "68 111 103 226 128 188 240 159 144 182 "

for codeUnit in dogString.utf16 {
    print("\(codeUnit) ", terminator: "")
}
print("") // Prints "68 111 103 8252 55357 56374 "

for scalar in dogString.unicodeScalars {
    print("\(scalar.value) ", terminator: "")
}
print("") // Prints "68 111 103 8252 128054 "

for scalar in dogString.unicodeScalars {
    print("\(scalar) ")
}

