package gopdf

import (
	"fmt"
	"testing"

	"github.com/tiechui1994/gopdf/core"
	"strings"
	"math/rand"
)

const (
	TABLE_IG = "IPAexG"
	TABLE_MD = "MPBOLD"
	TABLE_MY = "微软雅黑"
)

func ComplexTableReportWithData() {
	r := core.CreateReport()
	font1 := core.FontMap{
		FontName: TABLE_IG,
		FileName: "example//ttf/ipaexg.ttf",
	}
	font2 := core.FontMap{
		FontName: TABLE_MD,
		FileName: "example//ttf/mplus-1p-bold.ttf",
	}
	font3 := core.FontMap{
		FontName: TABLE_MY,
		FileName: "example//ttf/microsoft.ttf",
	}
	r.SetFonts([]*core.FontMap{&font1, &font2, &font3})
	r.SetPage("A4", "mm", "P")

	r.RegisterExecutor(core.Executor(ComplexTableReportWithDataExecutor), core.Detail)

	r.Execute("table_test_data.pdf")
	r.SaveAtomicCellText("table_test_data.txt")
	fmt.Println(r.GetCurrentPageNo())
}
func ComplexTableReportWithDataExecutor(report *core.Report) {
	unit := report.GetUnit()

	lineSpace := 0.01 * unit
	lineHeight := 2 * unit

	table := NewTable(5, 100, 80*unit, lineHeight, report)
	table.SetMargin(core.Scope{0, 0, 0, 0})

	// 先把当前的行设置完毕, 然后才能添加单元格内容.
	c00 := table.NewCellByRange(1, 1)
	c01 := table.NewCellByRange(2, 1)
	c03 := table.NewCellByRange(2, 2)
	c10 := table.NewCellByRange(3, 1)

	f1 := core.Font{Family: TABLE_MY, Size: 15, Style: ""}
	border := core.NewScope(0.8*unit, 0.8*unit, 0.8*unit, 0.8*unit)
	c00.SetElement(NewTextCell(table.GetColWidth(0, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).HorizontalCentered().SetContent("0-0"))
	c01.SetElement(NewTextCell(table.GetColWidth(0, 1), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).RightAlign().SetContent("0-1"))
	c03.SetElement(NewTextCell(table.GetColWidth(0, 3), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).RightAlign().SetContent("0-3近日，江苏苏州市姑苏区市场监督管理局向苏州定园旅游服务有限公司送达行政处罚决定书，对定园进行处罚，吊销其营业执照，此举开创了我国旅游景点因虚假宣传被吊销营业执照的先河"))
	c10.SetElement(NewTextCell(table.GetColWidth(1, 0), lineHeight, lineSpace, report).SetFont(f1).SetBorder(border).SetContent("1-0近日，江苏苏州市姑苏区市场监督管理局向苏州定园旅游服务有限公司送达行政处罚决定书了我国旅游景点因虚假宣传被吊销营业执照的先河"))

	f1 = core.Font{Family: TABLE_MY, Size: 10}
	border = core.NewScope(0.5*unit, 0.5*unit, 0, 0)

	for i := 0; i < 98; i++ {
		cells := make([]*TableCell, 5)
		for j := 0; j < 5; j++ {
			cells[j] = table.NewCell()
		}

		for j := 0; j < 5; j++ {
			str := `有限公司送达行政处罚决定书`
			s := fmt.Sprintf("%v-%v", i+2, str)
			w := table.GetColWidth(i+2, j)
			e := NewTextCell(w, lineHeight, lineSpace, report)
			e.SetFont(f1)
			if i%2 == 0 {
				e.SetBackColor("255,192,203")
			}
			e.SetBorder(border)
			e.SetContent(s)
			cells[j].SetElement(e)
		}
	}

	table.GenerateAtomicCell()
}

func ComplexTableReport() {
	r := core.CreateReport()

	font1 := core.FontMap{
		FontName: TABLE_IG,
		FileName: "example//ttf/ipaexg.ttf",
	}
	font2 := core.FontMap{
		FontName: TABLE_MD,
		FileName: "example//ttf/mplus-1p-bold.ttf",
	}
	font3 := core.FontMap{
		FontName: TABLE_MY,
		FileName: "example//ttf/microsoft.ttf",
	}
	r.SetFonts([]*core.FontMap{&font1, &font2, &font3})
	r.SetPage("A4", "mm", "P")

	r.RegisterExecutor(core.Executor(ComplexTableReportExecutor), core.Detail)

	r.Execute("table_test.pdf")
	r.SaveAtomicCellText("table_test.txt")
}
func ComplexTableReportExecutor(report *core.Report) {
	unit := report.GetUnit()
	lineSpace := 0.01 * unit
	lineHeight := 2 * unit

	form := NewTable(5, 100, 50*unit, lineHeight, report)
	form.SetMargin(core.Scope{0 * unit, 0 * unit, 0, 0})

	// todo: 先把当前的行设置完毕, 然后才能添加单元格内容.
	c00 := form.NewCellByRange(1, 1)
	c01 := form.NewCellByRange(2, 1)
	c03 := form.NewCellByRange(2, 2)
	c10 := form.NewCellByRange(3, 1)

	f1 := core.Font{Family: TABLE_MY, Size: 15}
	c00.SetElement(NewTextCell(form.GetColWidth(0, 0), lineHeight, lineSpace, report).SetFont(f1).SetContent(GetRandStr()))
	c01.SetElement(NewTextCell(form.GetColWidth(0, 1), lineHeight, lineSpace, report).SetFont(f1).SetContent(GetRandStr()))
	c03.SetElement(NewTextCell(form.GetColWidth(0, 3), lineHeight, lineSpace, report).SetFont(f1).SetContent(GetRandStr()))
	c10.SetElement(NewTextCell(form.GetColWidth(1, 0), lineHeight, lineSpace, report).SetFont(f1).SetContent(GetRandStr()))

	f1 = core.Font{Family: TABLE_MY, Size: 10}
	border := core.NewScope(0.5*unit, 0.5*unit, 0, 0)

	for i := 0; i < 98; i++ {
		cells := make([]*TableCell, 5)
		for j := 0; j < 5; j++ {
			cells[j] = form.NewCell()
		}

		for j := 0; j < 5; j++ {
			w := form.GetColWidth(i+2, j)
			// todo: div执行的严格顺序
			e := NewTextCell(w, lineHeight, lineSpace, report)
			e.SetFont(f1)
			if i%2 == 0 {
				e.SetBackColor("255,192,203")
			}
			e.SetBorder(border)
			e.SetContent(GetRandStr())
			cells[j].SetElement(e)
		}
	}

	form.GenerateAtomicCell()
}

func ManyTableReportWithData() {
	r := core.CreateReport()
	font1 := core.FontMap{
		FontName: TABLE_IG,
		FileName: "example//ttf/ipaexg.ttf",
	}
	font2 := core.FontMap{
		FontName: TABLE_MD,
		FileName: "example//ttf/mplus-1p-bold.ttf",
	}
	font3 := core.FontMap{
		FontName: TABLE_MY,
		FileName: "example//ttf/microsoft.ttf",
	}
	r.SetFonts([]*core.FontMap{&font1, &font2, &font3})
	r.SetPage("A4", "mm", "P")

	r.RegisterExecutor(core.Executor(ManyTableReportWithDataExecutor), core.Detail)

	r.Execute("many_table_data.pdf")
}
func ManyTableReportWithDataExecutor(report *core.Report) {
	unit := report.GetUnit()

	lineSpace := 0.01 * unit
	lineHeight := 2 * unit

	rows, cols := 340, 5
	table := NewTable(cols, rows, 80*unit, lineHeight, report)
	table.SetMargin(core.Scope{0, 0, 0, 0})

	for i := 0; i < rows; i += 5 {
		key := rand.Intn(3)
		//key := (i+1)%2 + 1
		f1 := core.Font{Family: TABLE_MY, Size: 10}
		border := core.NewScope(0.5*unit, 0.5*unit, 0, 0)

		switch key {
		case 0:
			for row := 0; row < 5; row++ {
				for col := 0; col < cols; col++ {
					conent := fmt.Sprintf("%v-(%v,%v)", 0, i+row, col)
					cell := table.NewCell()
					txt := NewTextCell(table.GetColWidth(i+row, col), lineHeight, lineSpace, report)
					txt.SetFont(f1).SetBorder(border).SetContent(conent + GetRandStr(1))
					cell.SetElement(txt)
				}
			}

		case 1:
			c00 := table.NewCellByRange(1, 5)
			c01 := table.NewCellByRange(2, 2)
			c03 := table.NewCellByRange(2, 3)
			c21 := table.NewCellByRange(2, 1)
			c31 := table.NewCellByRange(4, 1)
			c41 := table.NewCellByRange(4, 1)

			t00 := NewTextCell(table.GetColWidth(i+0, 0), lineHeight, lineSpace, report)
			t01 := NewTextCell(table.GetColWidth(i+0, 1), lineHeight, lineSpace, report)
			t03 := NewTextCell(table.GetColWidth(i+0, 3), lineHeight, lineSpace, report)
			t21 := NewTextCell(table.GetColWidth(i+2, 1), lineHeight, lineSpace, report)
			t31 := NewTextCell(table.GetColWidth(i+3, 1), lineHeight, lineSpace, report)
			t41 := NewTextCell(table.GetColWidth(i+4, 1), lineHeight, lineSpace, report)

			t00.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+0, 0) + GetRandStr(5))
			t01.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+0, 1) + GetRandStr(4))
			t03.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+0, 3) + GetRandStr(6))
			t21.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+2, 1) + GetRandStr(2))
			t31.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+3, 1) + GetRandStr(4))
			t41.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 1, i+4, 1) + GetRandStr(4))

			c00.SetElement(t00)
			c01.SetElement(t01)
			c03.SetElement(t03)
			c21.SetElement(t21)
			c31.SetElement(t31)
			c41.SetElement(t41)

		case 2:
			c00 := table.NewCellByRange(3, 2)
			c03 := table.NewCellByRange(2, 3)
			c20 := table.NewCellByRange(1, 2)
			c21 := table.NewCellByRange(2, 3)
			c33 := table.NewCellByRange(2, 2)
			c40 := table.NewCellByRange(1, 1)

			t00 := NewTextCell(table.GetColWidth(i+0, 0), lineHeight, lineSpace, report)
			t03 := NewTextCell(table.GetColWidth(i+0, 3), lineHeight, lineSpace, report)
			t20 := NewTextCell(table.GetColWidth(i+2, 0), lineHeight, lineSpace, report)
			t21 := NewTextCell(table.GetColWidth(i+2, 1), lineHeight, lineSpace, report)
			t33 := NewTextCell(table.GetColWidth(i+3, 3), lineHeight, lineSpace, report)
			t40 := NewTextCell(table.GetColWidth(i+4, 0), lineHeight, lineSpace, report)

			t00.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+0, 0) + GetRandStr(6))
			t03.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+0, 3) + GetRandStr(6))
			t20.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+2, 0) + GetRandStr(2))
			t21.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+2, 1) + GetRandStr(6))
			t33.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+3, 3) + GetRandStr(4))
			t40.SetFont(f1).SetBorder(border).SetContent(fmt.Sprintf("%v-(%v,%v)", 2, i+4, 0) + GetRandStr(1))

			c00.SetElement(t00)
			c03.SetElement(t03)
			c20.SetElement(t20)
			c21.SetElement(t21)
			c33.SetElement(t33)
			c40.SetElement(t40)
		}

	}

	table.GenerateAtomicCell()
}

func TestTableWithdata(t *testing.T) {
	//ComplexTableReportWithData()
}

func TestTable(t *testing.T) {
	//for i := 0; i < 1000; i++ {
	ManyTableReportWithData()
	//}
}

func GetRandStr(l ...int) string {
	str := "0123456789ABCDEFGHIGKLMNOPQRSTUVWXYZ"
	l = append(l, 5)
	data := strings.Repeat(str, l[0]*8/36+1)
	return data[:l[0]*8] + "---"
}
