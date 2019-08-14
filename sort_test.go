package testing

import ("testing"
       "github.com/stretchr/testify/require")

func TestSort(t *testing.T){
  c:= require.New(t)
  
  unsorted := []int{4,1,2,5,6,7}
  
  sorted := sortSlice(unsorted)
  
  expect := []int{1,2,4,5,6,7}
  
  c.Equal(expected, sorted)
}


