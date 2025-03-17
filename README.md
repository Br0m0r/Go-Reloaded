Go-Reloaded
Overview

This project implements a text transformation tool written in Go. It reads an input file containing text with special transformation commands, applies various modifications, and writes the corrected text to an output file. The tool supports number conversions, case modifications, punctuation formatting, quote adjustments, and simple grammatical corrections.
Features

    Number Conversion:
        Converts hexadecimal numbers (using (hex)) to decimal.
        Converts binary numbers (using (bin)) to decimal.

    Case Transformation:
        Converts a single word to uppercase, lowercase, or capitalized with (up), (low), or (cap).
        Applies multi-word transformations (e.g., (up, N), (low, N), (cap, N)) to modify the preceding N words.

    Punctuation Formatting:
        Adjusts spaces around punctuation marks (,, ., !, ?, :, and ;), ensuring they are close to the preceding word.
        Preserves group punctuations such as ellipses (...) and special sequences (e.g., !?.).

    Quote Formatting:
        Corrects spacing around text enclosed in single quotes so that there are no extra spaces inside the quotes.

    Grammatical Corrections:
        Automatically changes "a" to "an" when the following word begins with a vowel or an 'h'.



Installation

    Clone the repository:
    git clone https://github.com/Br0m0r/Go-Reloaded.git



Running the Application


    go run .

    The program reads from input.txt, applies the text transformations, and writes the output to output.txt.

Running Tests

    To run the unit tests, execute:

    go test

How It Works

    Input File Creation:
    The createInputFile() function generates an input.txt file containing sample sentences with transformation commands.

    Processing the Text:
        The processText() function starts by calling the instances() function to apply multi-word transformations based on commands like (up, N), (low, N), or (cap, N).
        It then processes the text word by word:
            Applies single-word transformations for tags such as (hex), (bin), (up), (low), and (cap).
            Corrects grammatical issues by changing "a" to "an" when necessary.
        Finally, it calls helper functions to format punctuation correctly and to remove extra spaces around quoted text.

    Output File Generation:
    The modified text is written line by line to output.txt using buffered I/O for efficiency.



Special thanks to peers and auditors for their feedback and contributions to making this project robust and easy to maintain.