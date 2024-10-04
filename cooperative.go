package main
import (
  "fmt"
  "time"
)

type Member struct {
  ID int
  Name string
  JoinDate time.Time
  Savings float64
  Loan float64
}

type Cooperative struct {
  MemberList []Member
}

func (c *Cooperative) AddMember(name string){
  newMember := Member{
    ID: len(c.MemberList) +1,
    Name: name,
    JoinDate: time.Now(),
    Savings: 0.0,
    Loan: 0.0,
  }

  c.MemberList = append(c.MemberList, newMember)
  fmt.Printf("New member added with ID %d and name %s.\n", newMember.ID, newMember.Name)
}

