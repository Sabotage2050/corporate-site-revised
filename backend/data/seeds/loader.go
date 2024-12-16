// backend/data/seeds/loader.go

package seeds

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"
)

type Environment string

const (
	Test Environment = "test"
	Prod Environment = "prod"
)

type DataLoader struct {
	dataDir     string
	environment Environment
}

type ForkliftType struct {
	Type     string
	FileName string
}

var ForkliftTypes = []ForkliftType{
	{Type: "battery", FileName: "battery"},
	{Type: "diesel", FileName: "diesel"},
	{Type: "gasoline", FileName: "gasoline"},
	// {Type: "shovelloader", FileName: "shovelloader"},
}

func NewDataLoader(dataDir string, env Environment) *DataLoader {
	return &DataLoader{
		dataDir:     dataDir,
		environment: env,
	}
}

func (l *DataLoader) LoadData(target interface{}) error {
	dir, err := os.Getwd()
	if err != nil {

	}
	fmt.Println("Current directory:", dir)
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("target must be a pointer to slice")
	}

	sliceValue := targetValue.Elem()
	elementType := sliceValue.Type().Elem()

	// Get table name from the Seed interface
	tableName := reflect.New(elementType).Interface().(Seed).TableName()

	filePath := filepath.Join(l.dataDir, tableName+".csv")

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read headers
	headers, err := reader.Read()
	if err != nil {
		return fmt.Errorf("failed to read CSV headers: %w", err)
	}

	// Create a map of field names to their indices
	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[header] = i
	}

	// Read records
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}

		// Create a new instance of the target type
		newElement := reflect.New(elementType).Elem()

		// Iterate through the fields of the struct
		for i := 0; i < elementType.NumField(); i++ {
			field := elementType.Field(i)
			csvTag := field.Tag.Get("csv")
			if csvTag == "" {
				continue
			}

			columnIndex, ok := headerMap[csvTag]
			if !ok {
				continue
			}

			value := record[columnIndex]
			if err := setField(newElement.Field(i), value); err != nil {
				return fmt.Errorf("failed to set field %s: %w", field.Name, err)
			}
		}

		// Append the new element to the slice
		sliceValue.Set(reflect.Append(sliceValue, newElement))
	}

	return nil
}

func setField(field reflect.Value, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Float32:
		f, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		field.SetFloat(f)
	case reflect.Float64:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(f)
	case reflect.Int, reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(i)
	case reflect.Struct:
		if field.Type() == reflect.TypeOf(time.Time{}) {
			t, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(t))
		}
	}
	return nil
}

func (l *DataLoader) LoadAllForkliftData() ([]ForkliftSeed, error) {
	var allSeeds []ForkliftSeed

	for _, fType := range ForkliftTypes {
		seeds, err := l.loadForkliftTypeData(fType)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s data: %w", fType.Type, err)
		}
		allSeeds = append(allSeeds, seeds...)
	}

	return allSeeds, nil
}

func (l *DataLoader) loadForkliftTypeData(fType ForkliftType) ([]ForkliftSeed, error) {
	filePath := filepath.Join(l.dataDir, fType.FileName+".csv")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV headers: %w", err)
	}

	headerMap := make(map[string]int)
	for i, header := range headers {
		headerMap[header] = i
	}

	var seeds []ForkliftSeed
	for {
		record, err := reader.Read()
		if err != nil {
			break // End of file
		}

		seed := ForkliftSeed{
			Enginetype: fType.Type, // フォークリフトタイプを設定
		}

		// レコードの値を構造体にマッピング
		val := reflect.ValueOf(&seed).Elem()
		typ := val.Type()
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			csvTag := field.Tag.Get("csv")
			if csvTag == "" || csvTag == "Enginetype" { // Enginetypeは既に設定済み
				continue
			}

			columnIndex, ok := headerMap[csvTag]
			if !ok {
				continue
			}

			if columnIndex < len(record) {
				if err := setField(val.Field(i), record[columnIndex]); err != nil {
					return nil, fmt.Errorf("failed to set field %s: %w", field.Name, err)
				}
			}
		}

		seeds = append(seeds, seed)
	}

	return seeds, nil
}
