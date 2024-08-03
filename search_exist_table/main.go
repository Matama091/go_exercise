package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// CSVファイルが保存されているディレクトリ
	csvDir := "./csv/"

	// レコードが入っているテーブル名を格納するためのリスト
	var tablesWithRecords []string

	// ディレクトリ内のすべてのCSVファイルを読み込む
	err := filepath.Walk(csvDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ファイルのみを処理
		if !info.IsDir() && filepath.Ext(path) == ".csv" {
			// CSVファイルを開く
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// CSVリーダーを作成
			reader := csv.NewReader(file)

			// 全ての行を読み込む
			records, err := reader.ReadAll()
			if err != nil {
				return err
			}

			// レコードが2行以上ある場合（1行目はカラム名）はテーブルにデータがあるとみなす
			if len(records) > 1 {
				// テーブル名を追加（ファイル名から拡張子を除いた部分をテーブル名とする）
				tableName := info.Name()
				tableName = tableName[:len(tableName)-len(filepath.Ext(tableName))]
				tablesWithRecords = append(tablesWithRecords, tableName)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through files: %v\n", err)
		return
	}

	// tablesWithRecordsのテーブル名を"result.txt"にかき出す
	result_file, err := os.Create("result.txt")
	if err != nil {
		fmt.Printf("Error creating result.txt: %v\n", err)
		return
	}

	for _, tableName := range tablesWithRecords {
		_, err := result_file.WriteString(tableName + "\n")
		if err != nil {
			fmt.Printf("Error writing to result.txt: %v\n", err)
			return
		}

		defer result_file.Close()

	}
}
