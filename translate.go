package main

import (
    "bufio"
    "flag"
    "fmt"
    "github.com/bregydoc/gtranslate"
    "github.com/schollz/progressbar/v3"
    "log"
    "os"
    "strings"
)

var (
    destLanguage   string
    sourceLanguage string
)

func main() {
    intermediateSourceLanguage := flag.String("srclang", "en", "Source language (e.g., 'en' for English)")
    intermediateDestinationLanguage := flag.String("destlang", "el", "Destination language (e.g., 'el' for Greek)")

    flag.Parse()

    sourceLanguage = strings.ToLower(*intermediateSourceLanguage)    // Dereference the pointer
    destLanguage = strings.ToLower(*intermediateDestinationLanguage) // Dereference the pointer

    // Print out the parsed language codes
    fmt.Printf("Source Language: %s\n", sourceLanguage)
    fmt.Printf("Destination Language: %s\n", destLanguage)

    // Read the file
    text, booleanMap := readFile("/home/dimi/Desktop/Projects/SubtitleTranslator/John.Wick.2014.720p.BluRay.x264.YIFY.en.srt")

    // Translate the text
    translatedText := translateText(text, booleanMap)

    // Write the translated text to file
    writeFile(translatedText)
}

// Read the file and extract strings for translation
func readFile(path string) ([]string, []bool) {
    var booleanMapOfLines []bool
    var preTranslationStrings []string
    file, err := os.Open(path)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var previousEmpty bool
    previousEmpty = false
    for scanner.Scan() {
        text := scanner.Text()
        preTranslationStrings = append(preTranslationStrings, text)
        if previousEmpty || (text != "\n" && (len(text) > 2 && text[2] != ':')) {
            booleanMapOfLines = append(booleanMapOfLines, true)
        } else {
            booleanMapOfLines = append(booleanMapOfLines, false)
        }
        if text == "\n" {
            previousEmpty = true
            continue
        }
        previousEmpty = false
    }
    if err := scanner.Err(); err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }
    return preTranslationStrings, booleanMapOfLines
}

// Translate the text with a progress bar
func translateText(text []string, booleanmap []bool) []string {
    var translatedText []string

    // Initialize progress bar
    bar := progressbar.Default(int64(len(text)), "Translating")

    for index, element := range text {
        if booleanmap[index] == true {
            if element != "" {
                translatedLine, err := gtranslate.TranslateWithParams(element, gtranslate.TranslationParams{From: sourceLanguage, To: destLanguage})
                if err != nil {
                    log.Printf("Translation failed for text: %s, error: %v", element, err)
                    translatedText = append(translatedText, element) // Fallback to original text if translation fails
                } else {
                    translatedText = append(translatedText, translatedLine)
                }
            }
        } else {
            translatedText = append(translatedText, element)
        }
        bar.Add(1) // Update the progress bar
    }

    return translatedText
}

func writeFile(translatedText []string) {
    f, err := os.Create("Translation.srt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
    for _, line := range translatedText {
        _, err := f.WriteString(line + "\n")
        if err != nil {
            log.Fatal(err)
        }
    }
}
