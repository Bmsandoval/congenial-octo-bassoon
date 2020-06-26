package internal

import (
	"encoding/csv"
	"fmt"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/models"
	"github.com/bmsandoval/congenial-octo-bassoon/internal/modules"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func ParseFile(fileName string, modules ...modules.ModuleInterface) {
	// Open file
	csvfile, err := os.Open(fileName)
	if err != nil {
		log.Printf("Unable to open file %s\n", err.Error())
		os.Exit(1)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	rowCount:=0
	for {
		// Read file one row at a time
		row, err := r.Read()

		rowCount++
		// skip header
		if rowCount == 1 {
			continue
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("error on line %d: %s\n skipping record...\n", rowCount, err.Error())
			continue
		}

		operationSize, err := strconv.Atoi(row[3])
		if err != nil {
			log.Printf("error on line %d: unparsable operation size %s\n skipping record...\n", rowCount, err.Error())
			continue
		}

		timestamp, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", row[0])
		if err != nil {
			log.Printf("error on line %d: could not load timestamp %s\n skipping record...\n", row[0], err.Error())
		}

		record := models.Row{
			Timestamp: timestamp,
			Username:  row[1],
			Operation: row[2],
			Size:      operationSize,
		}

		for _, module := range modules {
			module.Read(record)
		}
	}

	// Close file
	csvfile.Close()

	for _, module := range modules {
		fmt.Print(module.Display())
	}
}
