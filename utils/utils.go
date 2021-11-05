package utils

import (
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FieldControl(dataList []string) bool {
	if dataList[0] != "" {
			return true
	} else {
			return false
	}
}

func ParseData(value string, dept string) []string {
	var e []string
	if strings.Contains(value, "<tbody>") && strings.Contains(value, "<table>") {
			e[0] = ""
	} else {
			value = strings.ReplaceAll(value, "</td>", "</td>\n")
					//   value = strings.ReplaceAll(value,"</tr>","")
					//  value = strings.ReplaceAll(value,"class=\"BG0\"","")
					//  value = strings.ReplaceAll(value,"class=\"BG0\"","")
					//  value = strings.ReplaceAll(value,"class=\"BG0 L MLZ\"","")
					//  value = strings.ReplaceAll(value,"\u00a0</td>","")
					//  value = strings.ReplaceAll(value,"<td >","<td>")
					//  value = strings.ReplaceAll(value,"<td></td>","<td>-</td>")
					//  value = strings.ReplaceAll(value,"class=\"BG0 R FP1\"","")
			re := regexp.MustCompile(`<td.*?>(.*)</td>`)
			e[18] = dept
			_writeHeaders := os.Getenv("WRITE_HEADERS")
			if _writeHeaders == "true" && strings.Contains(value, "th") {
					value = strings.ReplaceAll(value, "</th>", "</th>\n")
					re = regexp.MustCompile(`<th.*?>(.*)</th>`)
					e[18] = dept + "_header_"
			}
			submatchall := re.FindAllStringSubmatch(value, -1)
			for i, element := range submatchall {
					switch i {
							case 0:
									e[0] = IsCheckSpace(element[1])
							case 1:
									e[1] = IsCheckSpace(element[1])
							case 2:
									e[2] = IsCheckSpace(element[1])
							case 3:
									e[3] = IsCheckSpace(element[1])
							case 4:
									e[4] = IsCheckSpace(element[1])
							case 5:
									e[5] = IsCheckSpace(element[1])
							case 6:
									e[6] = IsCheckSpace(element[1])
							case 7:
									e[7] = IsCheckSpace(element[1])
							case 8:
									e[8] = IsCheckSpace(element[1])
							case 9:
									e[9] = IsCheckSpace(element[1])
							case 10:
									e[10] = IsCheckSpace(element[1])
							case 11:
									e[11] = IsCheckSpace(element[1])
							case 12:
									e[12] = IsCheckSpace(element[1])
							case 13:
									e[13] = IsCheckSpace(element[1])
							case 14:
									e[14] = IsCheckSpace(element[1])
							case 15:
									e[15] = IsCheckSpace(element[1])
							case 16:
									e[16] = IsCheckSpace(element[1])
							case 17:
									e[17] = IsCheckSpace(element[1])
					}
			}
	}
	return e
}

func IsCheckSpace(value string) string {
	if value == "&nbsp;" {
			return "-"
	} else {
			return value
	}
}

func GetDefaultColumns() []string {
var	columns []string 
	columns[0]= "col0"
	columns[1]= "col1"
	columns[2]= "col2"
	columns[3]= "col3"
	columns[4]= "col4"
	columns[5]= "col5"
	columns[6]= "col6"
	columns[7]= "col7"
	columns[8]= "col8"
	columns[9]= "col9"
	columns[10] = "col10"
	columns[11] = "col11"
	columns[12] = "col12"
	columns[13] = "col13"
	columns[14] = "col14"
	columns[15] = "col15"
	columns[16] = "col16"
	columns[17] = "col17"
	columns[18] = "col18"
	return columns
}


func GetDataFromHtml(data string, pattern string) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(data))
	s := doc.Find(pattern)
	html, _ := s.Html()
	return html
}