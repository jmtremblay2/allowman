package scanCSV

import "testing"

func TestCountFieldsOneLine(t *testing.T){
  var tests = []struct {
    input FileLine
    nsep int
  }{
    {"",0},
    {",,,",3},
    {"1,b,b",2},
    {",1,2,3,",4},
    {",,,1,,,",6},
    {"123,345,432",2},
    {",aaa,aaa,aaa,aaa,",5},
  }
  for _, test := range tests {
    if got := CountFieldsOneLine(test.input); got != test.nsep {
      t.Errorf(`CountFieldsOneLine(%q) != %d`,test.input, test.nsep)
     }
  }
}
