# Agenda
15331141 赖秀娜        
15331176 李梓桥        
15331435 周长安        

## Usage
　　`agd command`

- regist
　　`agd regist -u usernaeme -p password -e e-mail -t telephone`

- login
　　`agd login -u username -p password`

- logout
　　`agd logout`

- list all users
　　`agd lsu`

- delete current account
　　`agd del -p password`

- create a meeting
　　`agd cm -t title -p participator1 -p participator2 -s start -e end`

- change the patricipators of a meeting
　　add: `agd ap -t title -p name`
　　delete: `agd dp -t title -p name`

- list meetings during a period
　　`agd lsm -s start -e end`

- cancel a meeting
　　`agd cacel -t title`

- quit a meeing
　　`agd quit -t title`

- clear all meeting
　　`agd clear`
