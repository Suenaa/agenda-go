# Agenda

contributors:                          
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

- change the participants of a meeting                      
　　add: `agd ap -t title -p name`                                   
　　delete: `agd dp -t title -p name`

- list meetings during a period                    
　　`agd lsm -s start -e end`

- cancel a meeting                      
　　`agd cancel -t title`

- quit a meeing                   
　　`agd quit -t title`

- clear all meetings                  
　　`agd clear`

## examples

- regist
```
./agd regist -u user1 -p 123 -e email1@mail.com -t 11111
Success

./agd regist -u user2 -p 122 -e email2@mail.com -t 22222
Success

./agd regist -u user3 -p 333 -e email3@mail.com -t 33333
Success

./agd regist -u user1 -p 111 -e mail2@mail.com -t 11111
this username is aleardy exist
```

- login
```
./agd login -u usr1 -p 123
username not exist

./agd login -u user1 -p 12
incorrect password

./agd login -u user1 -p 123
Success
```

- logout
```
./agd logout
Success
```

- list all users
```
./agd lsu
user1, email1@mail.com, 11111
user2, email2@mail.com, 22222
user3, email3@mail.com, 33333
```

- delete current account
```
./agd lsu
user1, email1@mail.com, 11111
user2, email2@mail.com, 22222
user3, email3@mail.com, 33333

./agd login -u user1 -p 123
Success

./agd del -p 124
incorrect password

./agd del -p 123
Success

./agd lsu
user2, email2@mail.com, 22222
user3, email3@mail.com, 33333
```

- create a meeting
```
./agd cm -t title -p user1, user2 -s 2017-11-01T10:00 -e 2017-11-01T10:30
no user login

./agd login -u user3 -p333
Success

./agd cm -t title -p user1, user2 -s 2017-11-01T10:00 -e 2017-11-01T10:30
Success

./agd lsm -s 2017-11-01T10:00 -e 2017-11-01T10:30
title
　- sponsor: user3
　- time: 2017-11-01 10:00 - 2017-11-01 10:30
　- participators: user1
```

- list meetings during a period
```
./agd lsm -s 2017-11-01T0:00 -e 2017-11-01T20:00
title
　- sponsor: user3
　- time: 2017-11-01 10:00 - 2017-11-01 10:30
　- participators: user1

title2
　- sponsor: user3
　- time: 2017-11-01 11:00 - 2017-11-01 11:30
　- participators: user1

title3
　- sponsor: user3
　- time: 2017-11-01 11:00 - 2017-11-01 11:30
　- participators: user2
```

- change the participants of a meeting
```
./agd ap -t title2 -p user
participator user is not exist

./agd ap -t title2 -p user2
Success

./agd dp -t title1 -p user2
the title not exist

./agd dp -t title3 -p user2
Success

./agd lsm -s 2017-11-01T0:00 -e 2017-11-01T23:00
title
　- sponsor: user3
　- time: 2017-11-01 10:00 - 2017-11-01 10:30
　- participators: user1

title2
　- sponsor: user3
　- time: 2017-11-01 11:00 - 2017-11-01 11:30
　- participators: user1, user2
```

- cancel a meeting
```
./agd cancel -t title
Success

./agd lsm -s 2017-11-01T0:00 -e 2017-11-01T23:00
title2
　- sponsor: user3
　- time: 2017-11-01 11:00 - 2017-11-01 11:30
　- participators: user1, user2
```

- quit a meeting
```


```

- clear all meetings
```
./agd lsm -s 2017-11-01T0:00 -e 2017-11-01T23:00
title2
　- sponsor: user3
　- time: 2017-11-01 11:00 - 2017-11-01 11:30
　- participators: user1, user2

title
　- sponsor: user3
　- time: 2017-11-01 11:00 - 2017-11-01 11:30
　- participators: user1

title3
　- sponsor: user3
　- time: 2017-11-01 11:00 - 2017-11-01 11:30
　- participators: user2

./agd clear
Success

./agd lsm -s 2017-11-01T0:00 -e 2017-11-01T23:00


```
