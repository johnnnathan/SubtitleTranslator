package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/bregydoc/gtranslate"
)


var (
	destLanguage string
  sourceLanguage string 
)

func main() {
	intermediateSourceLanguage := flag.String("srclang", "en", "Source language (e.g., 'en' for English)")
	intermediateDestinationLanguage := flag.String("destlang", "gr", "Destination language (e.g., 'en' for English)")

	flag.Parse()

	sourceLanguage = strings.ToLower(*intermediateSourceLanguage)  // Dereference the pointer
	destLanguage = strings.ToLower(*intermediateDestinationLanguage) // Dereference the pointer

	// Print out the parsed language codes
	fmt.Printf("Source Language: %s\n", sourceLanguage)
	fmt.Printf("Destination Language: %s\n", destLanguage)
	// Read the file

  print("Something")
	text := readFile("/home/dimi/Desktop/Projects/SubtitleTranslator/template.srt")

  print("Something")
	// Translate the text
	translatedText := translateText(text)

	// Output the translated text
	for _, element := range translatedText {
		fmt.Println(element)
	}
}

// Read the file and extract strings for translation
func readFile(path string) []string {
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
		if previousEmpty || (text != "\n" && (len(text) >= 2 && text[2] != ':')) {
			preTranslationStrings = append(preTranslationStrings, text)
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
	return preTranslationStrings
}

// Translate the text
func translateText(text []string) []string {
	var translatedText []string
	for _, element := range text{
    translatedLine , _ := gtranslate.TranslateWithParams(element, gtranslate.TranslationParams{From: sourceLanguage , To: destLanguage},) 
    translatedText = append(translatedText, translatedLine)
	}

	return translatedText
}
