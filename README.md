## MKDIR на максималках ~~(нет)~~
___

#### Программа, при старте, запрашивает названия файла, который нужно создать
*например:* config.json или info или 2021-01-02.log. 
#### После чего сообщит об успешном или неуспешном создании файла.

+ Если файл создан несупешно, запросит новое имя файла и так по кругу
+ Если файл создан успешно - программа попросит ввести информацию, которую нужно записать в файл.
+ Если информация запислась неуспешно - программа сообщит об этом, удалит файл, а после завершится.
+ Если информация запислась  успешно - спросит, нужно или еще что-то записать
+ Если ввести ***no*** - программа завершится. 
+ Если ввести ***yes*** - программа снова попросит ввести информацию для записи и так по кругу.