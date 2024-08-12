package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	// 检查是否提供了文件路径作为命令行参数
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Printf("File path provided: %s == start \n", filePath)

	processFile(filePath)
}

func processFile(filePath string) {
	rows, headerRow, err := readExcelFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 遍历所有行，记录第一列值的出现次数
	countMap := make(map[string]int)
	for _, row := range rows {
		if len(row) == 0 {
			break
		}
		countMap[row[0]]++
	}
	groupMap := make(map[int][]string)
	for k, v := range countMap {
		groupMap[v] = append(groupMap[v], k)
	}

	// 创建一个工作表，保存数据
	targetFile := excelize.NewFile()
	groupSW := make(map[string]*excelize.StreamWriter)
	// 为每个不同的出现次数创建一个工作表
	for i := range groupMap {
		sheetNameList := []string{
			fmt.Sprintf("Sheet%d", i),
			fmt.Sprintf("Sheet_%d", i),
		}
		for _, sheetName := range sheetNameList {
			_, err := targetFile.NewSheet(sheetName)
			if err != nil {
				return
			}
			sw, err := targetFile.NewStreamWriter(sheetName)
			if err != nil {
				return
			}
			groupSW[sheetName] = sw
		}
	}

	// 根据出现次数将数据分组
	groupData := make(map[int][][]string)
	for _, row := range rows {
		if len(row) == 0 {
			break
		}
		count := countMap[row[0]]
		groupData[count] = append(groupData[count], row)
	}

	// 对分组内部数据进行排序
	for k, rows := range groupData {
		var data SortRows
		for _, row := range rows {
			data = append(data, &SortRow{SortKey: row[0], Row: row}) // 排序的依据是第一列
		}
		sort.Sort(data)
		sortRows := make([][]string, 0, 0)
		for _, item := range data {
			sortRows = append(sortRows, item.Row)
		}
		groupData[k] = sortRows
	}

	// 写入数据
	for count, groupRows := range groupData {
		sw, ok := groupSW[fmt.Sprintf("Sheet%d", count)]
		if !ok {
			break
		}
		err = sw.SetRow("A1", ConvertStringToInterfaceSlice(headerRow))

		for index, row := range groupRows {
			cell, err := excelize.CoordinatesToCellName(1, index+2)
			if err != nil {
				fmt.Println(err)
				break
			}
			err = sw.SetRow(cell, ConvertStringToInterfaceSlice(row))
			if err != nil {
				return
			}
		}
	}

	// 对数据进行拆列
	for count, groupRows := range groupData {
		sw, ok := groupSW[fmt.Sprintf("Sheet_%d", count)]
		if !ok {
			break
		}
		var rsRows [][]string
		rsRows = append(rsRows, groupRows[0])
		for i := 1; i < len(groupRows); i++ {
			pre := groupRows[i-1]
			curr := groupRows[i]
			// 第一列名称相同，进行列转行
			if pre[0] == curr[0] {
				for i := range curr {
					if curr[i] == pre[i] {
						continue
					}
					rsRows[len(rsRows)-1] = append(rsRows[len(rsRows)-1], curr[i])
					headerRow = append(headerRow, headerRow[i])
				}
			} else {
				rsRows = append(rsRows, groupRows[i])
			}
		}

		err = sw.SetRow("A1", ConvertStringToInterfaceSlice(headerRow))
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, row := range rsRows {
			cell, err := excelize.CoordinatesToCellName(1, index+2)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = sw.SetRow(cell, ConvertStringToInterfaceSlice(row))
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		//sheetName := fmt.Sprintf("Sheet_%d", count)
		//rows, err := targetFile.GetRows(sheetName)
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//
		//// 将排序后的数据写回Excel文件
		//for i, d := range data {
		//	sortedFile.SetCellValue(sortedSheetName, fmt.Sprintf("A%d", i+2), d.Value)
		//	sortedFile.SetCellValue(sortedSheetName, fmt.Sprintf("B%d", i+2), d.OtherData)
		//}
	}

	for _, sw := range groupSW {
		if err := sw.Flush(); err != nil {
			fmt.Println(err)
			return
		}
	}

	fileDir := strings.Split(filePath, ".xlsx")[0]
	if len(fileDir) == 0 {
		fmt.Println("invalid filePath: ", filePath)
	}
	err = os.Mkdir(fileDir, os.ModeDir|os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	err = targetFile.SaveAs(fmt.Sprintf("%s/数据加工结果%s.xlsx", fileDir, time.Now().Format("2006-01-02")))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("File path provided: %s == success <`.`> \n", filePath)
}

// readExcelFile 读取目标文件数据
func readExcelFile(filePath string) (body [][]string, header []string, err error) {
	sourceFile, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := sourceFile.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := sourceFile.GetRows("Sheet1")
	if err != nil {
		return nil, nil, err
	}
	if len(rows) < 2 {
		return nil, nil, err

	}
	headerRow := rows[0]
	rows = rows[1:]
	return rows, headerRow, nil
}

func ConvertStringToInterfaceSlice(strings []string) []interface{} {
	result := make([]interface{}, len(strings))
	for i, s := range strings {
		result[i] = s
	}
	return result
}

type SortRow struct {
	SortKey string
	Row     []string
}

type SortRows []*SortRow

func (r SortRows) Len() int {
	return len(r)
}

func (r SortRows) Less(i, j int) bool {
	return r[i].SortKey < r[j].SortKey
}

func (r SortRows) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
