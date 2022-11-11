
***Запуск приложения***

``` sudo docker-compose up --build app ```

***Get all***

Метод возвращает json со всеми пользователями и их данными

![image](https://user-images.githubusercontent.com/106326324/201370184-79bdd304-2d02-42d5-b320-4e89b974545e.png)

***Post donate***

Метод возвращает json с сообщением об успешном начислении средств

![image](https://user-images.githubusercontent.com/106326324/201370708-711d0a65-4f05-44d9-8099-8088ff11383d.png)

***Post trade***

Метод возвращает json с сообщением об успешном или неуспешном переводе средств другому пользователю

![image](https://user-images.githubusercontent.com/106326324/201370951-ec967af9-44ca-4710-93ec-8b3b3ecf80a5.png)

***Post info***

Метод возвращает json с сообщением об количестве средств на счету пользователя

![image](https://user-images.githubusercontent.com/106326324/201371250-6cbfddf2-8b11-4074-aea1-cb28ff78c915.png)

***Post buy***

Метод возвращает json с сообщением об успешном или неуспешном создании заказа ( резервирования средств )

![image](https://user-images.githubusercontent.com/106326324/201371615-ccf2b5f6-7f26-4cfa-8b7e-fb197c93310e.png)

***Post accept***

Метод возвращает json с сообщением об успешном или неуспешном выполнении заказа 

![image](https://user-images.githubusercontent.com/106326324/201371916-7cdcd2f2-dac0-4169-9ea8-53c921698ac5.png)


***Состояние БД после всех манипуляций***

![image](https://user-images.githubusercontent.com/106326324/201372777-b4a5ae0c-de0d-47b6-a356-b407835b67b0.png)

***Сформированный отчет для бухгалтерии (другие входные данные)***

![image](https://user-images.githubusercontent.com/106326324/201380386-9bb5a37f-07d2-46fd-a681-3915f3d564bd.png)

***p/s***
Добавил свою реализацию обмена деньгами, так как нет ее описания.
Добавил метод чтобы узнать всю информацию о пользователях
Так же добавил некоторые поля в структуры

