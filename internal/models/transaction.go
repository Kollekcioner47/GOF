package models

import "time"

type Transaction struct {
    ID          int
    UserID      int
    CategoryID  int
    Amount      float64
    Description string
    Date        time.Time
    CreatedAt   time.Time
}
