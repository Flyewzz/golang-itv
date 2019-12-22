# golang-itv
Welcome to the ITV Server page!

The description of the service is below:

## REST API:

* GET /request             params: method:string, url: string   => a server to do your request
* GET /requests               (without params)   => a server gives you a list of all the tasks
*	GET /requests/page/{number} (number:integer)   => a server gives you a list of all the tasks in page #number
*	GET /requests/{id}          (id:integer)       => a server gives you the task with #id
* DELETE /tasks               (without params)   => a server removes all the tasks
*	DELETE /requests/{id}       (id:integer)       => a server removes the task you want (with #id)

!This repository uses Continuous Integration system!

You can test this system. A site is available [there](https://itv.fezzo.ru) (The main page will not answer, use other API options)

https://itv.fezzo.ru/tasks (to obtain all the tasks on the server)

