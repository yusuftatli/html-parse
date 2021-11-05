package main

import (
	"strings"

	"github.com/yusuftatli/html-parse/utils"
)


func main(){
}

func GetParsedHtml(data string, fileName string, uniqueID string) [][]string {
	var _ListHtml [][]string
	defaultColumns := utils.GetDefaultColumns()
	_ListHtml = append(_ListHtml, defaultColumns)
	html := utils.GetDataFromHtml(data, "table tr td table tr td table")
	parsedHtml := strings.Split(html, "<td class=\"BG0 S\" colspan=\"9\">")
	for _, masterRow := range parsedHtml {
			firstSplit := strings.Split(masterRow, "<td class=\"BG1 S\" colspan=\"16\">")
			for _, firstRow := range firstSplit {
					if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"15\">") {
							secondSplit := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"15\">")
							for _,
							secondRow := range secondSplit {
									if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"16\">") {
											splitData2 := strings.Split(secondRow, "<td class=\"BG0 S\" colspan=\"16\">")
											for _,
											bg0Data := range splitData2 {
													trSplitdata := strings.Split(bg0Data, "<tr>")
													for _, row := range trSplitdata {
															_index := "0"
															if strings.Contains(row, "class=\"BG0") {
																	_index = "3"
															} else if strings.Contains(row, "class=\"BG1") {
																	_index = "4"
															} else if strings.Contains(row, "class=\"BG2") {
																	_index = "2"
															} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																	_index = "3"
															}
															val := utils.ParseData(row, _index)
															if utils.FieldControl(val) {
																_ListHtml = append(_ListHtml, val)
															}
													}
											}
									} else if strings.Contains(secondRow, "<td class=\"BG2 S\" colspan=\"14\">") {
											splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
											for _,
											bg2Data := range splitData {
													trSplitdata := strings.Split(bg2Data, "<tr>")
													for _, row := range trSplitdata {
															_index := "0"
															if strings.Contains(row, "class=\"BG0") {
																	_index = "3"
															} else if strings.Contains(row, "class=\"BG1") {
																	_index = "1"
															} else if strings.Contains(row, "class=\"BG2") {
																	_index = "2"
															} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																	_index = "3"
															}
															val := utils.ParseData(row, _index)
															if utils.FieldControl(val) {
																_ListHtml = append(_ListHtml, val)
															}
													}
											}

									} else if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"14\">") {
											splitDataThird := strings.Split(secondRow, "<td class=\"BG0 S\" colspan=\"14\">")
											for _,
											thirdData := range splitDataThird {
													if strings.Contains(thirdData, "<td class=\"BG1 S\" colspan=\"16\">") {
															splitBg016 := strings.Split(thirdData, "<td class=\"BG1 S\" colspan=\"16\">")
															for _,
															bg2Data := range splitBg016 {
																	trSplitdata := strings.Split(bg2Data, "<tr>")
																	for _, row := range trSplitdata {
																			_index := "0"
																			if strings.Contains(row, "class=\"BG0") {
																					_index = "3"
																			} else if strings.Contains(row, "class=\"BG1") {
																					_index = "4"
																			} else if strings.Contains(row, "class=\"BG2") {
																					_index = "5"
																			} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																					_index = "3"
																			}
																			val := utils.ParseData(row, _index)
																			if utils.FieldControl(val) {
																				_ListHtml = append(_ListHtml, val)
																			}
																	}
															}
													} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"9\">") {
															splitbgos15 := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"9\">")
															for _,
															bgorow := range splitbgos15 {
																	trSplitdata := strings.Split(bgorow, "<tr>")
																	for _,
																	row := range trSplitdata {
																			_index := "0"
																			if strings.Contains(row, "class=\"BG0") {
																					_index = "2"
																			} else if strings.Contains(row, "class=\"BG1") {
																					_index = "4"
																			} else if strings.Contains(row, "class=\"BG2\">") {
																					_index = "5"
																			} else if strings.Contains(row, "class=\"BG2 S") {
																					_index = "2"
																			} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																					_index = "3"
																			}
																			val := utils.ParseData(row, _index)
																			if utils.FieldControl(val) {
																					_ListHtml = append(_ListHtml, val)
																			}
																	}
															}
													} else {
															trSplitdata := strings.Split(thirdData, "<tr>")
															for _,
															row := range trSplitdata {
																	_index := "0"
																	if strings.Contains(row, "class=\"BG0") {
																			_index = "3"
																	} else if strings.Contains(row, "class=\"BG1") {
																			_index = "4"
																	} else if strings.Contains(row, "class=\"BG2\">") {
																			_index = "5"
																	} else if strings.Contains(row, "class=\"BG2 S") {
																			_index = "2"
																	} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																			_index = "3"
																	}
																	val := utils.ParseData(row, _index)
																	if utils.FieldControl(val) {
																			_ListHtml = append(_ListHtml, val)
																	}
															}
													}
											}
									} else if strings.Contains(secondRow, "<td class=\"BG0 S\" colspan=\"15\">") {
											splitbgos15 := strings.Split(secondRow, "<td class=\"BG0 S\" colspan=\"15\">")
											for _,
											bgorow := range splitbgos15 {
													if strings.Contains(bgorow, "<td class=\"BG1 S\" colspan=\"16\">") {
															splitBg016 := strings.Split(bgorow, "<td class=\"BG1 S\" colspan=\"16\">")
															for _,
															bg2Data := range splitBg016 {
																	trSplitdata := strings.Split(bg2Data, "<tr>")
																	for _, row := range trSplitdata {
																			_index := "0"
																			if strings.Contains(row, "class=\"BG0") {
																					_index = "3"
																			} else if strings.Contains(row, "class=\"BG1") {
																					_index = "4"
																			} else if strings.Contains(row, "class=\"BG2") {
																					_index = "5"
																			} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																					_index = "3"
																			}
																			val := utils.ParseData(row, _index)
																			if utils.FieldControl(val) {
																					_ListHtml = append(_ListHtml, val)
																			}
																	}
															}
													} else {
															trSplitdata := strings.Split(bgorow, "<tr>")
															for _,
															row := range trSplitdata {
																	_index := "0"
																	if strings.Contains(row, "class=\"BG0") {
																			_index = "3"
																	} else if strings.Contains(row, "class=\"BG1") {
																			_index = "4"
																	} else if strings.Contains(row, "class=\"BG2\">") {
																			_index = "5"
																	} else if strings.Contains(row, "class=\"BG2 S") {
																			_index = "2"
																	} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																			_index = "3"
																	}
																	val := utils.ParseData(row, _index)
																	if utils.FieldControl(val) {
																			_ListHtml = append(_ListHtml, val)
																	}
															}
													}
											}
									} else {
											trSplitdata := strings.Split(secondRow, "<tr>")
											for _,
											row := range trSplitdata {
													_index := "0"
													if strings.Contains(row, "class=\"BG0") {
															_index = "3"
													} else if strings.Contains(row, "class=\"BG1") {
															_index = "1"
													} else if strings.Contains(row, "class=\"BG2") {
															_index = "2"
													} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
															_index = "3"
													}
													val := utils.ParseData(row, _index)
													if utils.FieldControl(val) {
															_ListHtml = append(_ListHtml, val)
													}
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"16\">") {
							splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
							for _,
							bg2Data := range splitData {
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
											_index := "0"
											if strings.Contains(row, "class=\"BG0") {
													_index = "3"
											} else if strings.Contains(row, "class=\"BG1") {
													_index = "1"
											} else if strings.Contains(row, "class=\"BG2") {
													_index = "2"
											} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
													_index = "3"
											}
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, val)
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"14\">") {
							splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"16\">")
							for _,
							bg2Data := range splitData {
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
											_index := "0"
											if strings.Contains(row, "class=\"BG0") {
													_index = "2"
											} else if strings.Contains(row, "class=\"BG1") {
													_index = "1"
											} else if strings.Contains(row, "class=\"BG2") {
													_index = "2"
											} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
													_index = "3"
											}
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, val)
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"13\">") {
							splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"13\">")
							for _,
							bg2Data := range splitData {
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
											_index := "0"
											if strings.Contains(row, "class=\"BG0") {
													_index = "2"
											} else if strings.Contains(row, "class=\"BG1") {
													_index = "1"
											} else if strings.Contains(row, "class=\"BG2") {
													_index = "2"
											} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
													_index = "3"
											}
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, val)
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"17\">") {
							splitData := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"17\">")
							for _,
							bg2Data := range splitData {
									trSplitdata := strings.Split(bg2Data, "<tr>")
									for _, row := range trSplitdata {
											_index := "0"
											if strings.Contains(row, "class=\"BG0") {
													_index = "2"
											} else if strings.Contains(row, "class=\"BG1") {
													_index = "1"
											} else if strings.Contains(row, "class=\"BG2") {
													_index = "2"
											} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
													_index = "3"
											}
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, val)
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG0 S\" colspan=\"15\">") {
							splitbgos15 := strings.Split(firstRow, "<td class=\"BG0 S\" colspan=\"15\">")
							for _,
							bgorow := range splitbgos15 {
									if strings.Contains(bgorow, "<td class=\"BG1 S\" colspan=\"16\">") {
											splitBg016 := strings.Split(bgorow, "<td class=\"BG1 S\" colspan=\"16\">")
											for _,
											bg2Data := range splitBg016 {
													trSplitdata := strings.Split(bg2Data, "<tr>")
													for _, row := range trSplitdata {
															_index := "0"
															if strings.Contains(row, "class=\"BG0") {
																	_index = "3"
															} else if strings.Contains(row, "class=\"BG1") {
																	_index = "4"
															} else if strings.Contains(row, "class=\"BG2") {
																	_index = "5"
															} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
																	_index = "3"
															}
															val := utils.ParseData(row, _index)
															if utils.FieldControl(val) {
																	_ListHtml = append(_ListHtml, val)
															}
													}
											}
									} else {
											trSplitdata := strings.Split(bgorow, "<tr>")
											for _,
											row := range trSplitdata {
													_index := "0"
													if strings.Contains(row, "class=\"BG0") {
															_index = "3"
													} else if strings.Contains(row, "class=\"BG1") {
															_index = "4"
													} else if strings.Contains(row, "class=\"BG2\">") {
															_index = "5"
													} else if strings.Contains(row, "class=\"BG2 S") {
															_index = "2"
													} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
															_index = "3"
													}
													val := utils.ParseData(row, _index)
													if utils.FieldControl(val) {
															_ListHtml = append(_ListHtml, val)
													}
											}
									}
							}
					} else if strings.Contains(firstRow, "<td class=\"BG2 S\" colspan=\"9\">") {
							splitbgos15 := strings.Split(firstRow, "<td class=\"BG2 S\" colspan=\"9\">")
							for _,
							bgorow := range splitbgos15 {
									trSplitdata := strings.Split(bgorow, "<tr>")
									for _,
									row := range trSplitdata {
											_index := "0"
											if strings.Contains(row, "class=\"BG0") {
													_index = "2"
											} else if strings.Contains(row, "class=\"BG1") {
													_index = "4"
											} else if strings.Contains(row, "class=\"BG2\">") {
													_index = "5"
											} else if strings.Contains(row, "class=\"BG2 S") {
													_index = "2"
											} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
													_index = "3"
											}
											val := utils.ParseData(row, _index)
											if utils.FieldControl(val) {
													_ListHtml = append(_ListHtml, val)
											}
									}
							}
					} else {
							trSplitdata := strings.Split(firstRow, "<tr>")
							for _,
							row := range trSplitdata {
									_index := "0"
									if strings.Contains(row, "class=\"BG0") {
											_index = "master"
									} else if strings.Contains(row, "class=\"BG1") {
											_index = "1"
									} else if strings.Contains(row, "class=\"BG2") {
											_index = "2"
									} else if strings.Contains(row, "<td class=\"BG0 R FP1\">") {
											_index = "3"
									}
									val := utils.ParseData(row, _index)
									if utils.FieldControl(val) {
											_ListHtml = append(_ListHtml, val)
									}
							}
					}
			}
	}
	return _ListHtml
}