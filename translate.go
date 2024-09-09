package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/bregydoc/gtranslate"
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
	intermediateDestinationLanguage := flag.String("destlang", "el", "Destination language (e.g., 'en' for English)")

	flag.Parse()

	sourceLanguage = strings.ToLower(*intermediateSourceLanguage)    // Dereference the pointer
	destLanguage = strings.ToLower(*intermediateDestinationLanguage) // Dereference the pointer

	// Print out the parsed language codes
	fmt.Printf("Source Language: %s\n", sourceLanguage)
	fmt.Printf("Destination Language: %s\n", destLanguage)
	// Read the file

	text , booleanMap:= readFile("/home/dimi/Desktop/Projects/SubtitleTranslator/template.srt")

	// Translate the text
	translatedText := translateText(text, booleanMap)

	// Output the translated text
	for _, element := range translatedText {
		fmt.Println(element)
	}
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
		if previousEmpty || (text != "\n" && (len(text) >= 2 && text[2] != ':')) {
      booleanMapOfLines = append(booleanMapOfLines, true)
		}else{
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

// Translate the text
func translateText(text []string, booleanmap []bool) []string {
	var translatedText []string
	for index, element := range text {
    if booleanmap[index] == true{
      if element != ""{
        translatedLine, _ := gtranslate.TranslateWithParams(element, gtranslate.TranslationParams{From: sourceLanguage, To: destLanguage})
        translatedText = append(translatedText, translatedLine)
      }
      continue
    }
    translatedText = append(translatedText, element)
	}

	return translatedText
}

func writeFile(translatedText []string) {
  f, _ := os.Create("Translation.srt")
  defer f.Close()
  for _, line := range translatedText{
    _, err := f.WriteString(line + "\n")
    if err != nil{
      log.Fatal(err)
    }
  }
}
