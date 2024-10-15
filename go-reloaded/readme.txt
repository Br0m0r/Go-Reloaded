	                              PART 1: THE MAIN FUNCTION

This is the entry point of the program. It starts by creating the input file (input.txt), reading the text from it, applying transformations to the text, and then writing the transformed content to an output file (output.txt).

STEP 1: We call createInputFile() to create a file input.txt with predefined content. This content will serve as the input to our transformations.
STEP 2: We open this file for reading and handle any errors that might occur.
STEP 3: We create an output file output.txt, which will store the transformed version of the text.
STEP 4: Using a bufio.Scanner, we read the input.txt file line by line.
STEP 5: For each line, we apply transformations via the processText() function.
STEP 6: The processed result is written to output.txt. Finally, we flush the buffer to ensure all data is written to the file.



                                     PART 2: CREATEINPUTFILE FUNCTION

This function creates an input.txt file with some sample sentences, including special transformation commands like (up, N), (low, N), (hex), and more.

PURPOSE: This function writes a predefined text into input.txt. This text includes transformations like converting text to uppercase, lowercase, or converting numbers from hexadecimal or binary to decimal.
HOW IT WORKS: The function uses os.WriteFile to create the file and write the content into it.


                                     PART 3: PROCESSTEXT FUNCTION

This function is responsible for processing each line of text by applying the necessary transformations.

TRANSFORMATION FLOW:
INITIAL TRANSFORMATION: The text is first passed to the instances function, which processes multi-word transformations like (up, 3) or (low, 2).
WORD-LEVEL TRANSFORMATION: It then iterates over each word in the text, checking for tags like (hex), (bin), (up), etc. Each tag applies a specific transformation.
GRAMMAR FIX: It ensures that "a" is changed to "an" if the next word starts with a vowel.
Final Cleanup: The processed words are joined together, and punctuation and quotes are formatted properly.



                                    PART 4: THE INSTANCES FUNCTION

This function processes transformations like (up, N), (low, N), or (cap, N) where the transformation applies to a specified number of preceding words.

PURPOSE: This function finds transformations like (up, N) and applies them to the preceding words.
HOW IT WORKS:
It uses regular expressions to find transformations.
For each match, it applies the transformation to the specified number of words and then reconstructs the text without the tag.



